package aws_test

import (
	"context"
	"fmt"
	"strings"
	"time"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/aws-resolver-rules-operator/pkg/resolver"
)

var _ = Describe("Route53 Resolver client", func() {
	BeforeEach(func() {
		ctx = context.Background()

		route53Client, err = awsClients.NewRoute53Client(Region, AwsIamArn)
		Expect(err).NotTo(HaveOccurred())
	})

	When("creating hosted zones", func() {
		var hostedZoneId string
		tags := map[string]string{
			"Name":      "jose",
			"something": "else",
		}

		AfterEach(func() {
			_, err = rawRoute53Client.DeleteHostedZoneWithContext(ctx, &route53.DeleteHostedZoneInput{Id: awssdk.String(hostedZoneId)})
			Expect(err).NotTo(HaveOccurred())
		})

		Context("we want a public hosted zone", func() {
			It("creates a public hosted zone successfully", func() {
				hostedZoneId, err = route53Client.CreateHostedZone(ctx, logger, resolver.BuildPublicHostedZone("apublic.test.example.com", tags))
				Expect(err).NotTo(HaveOccurred())

				var publicListHostedZoneResponse *route53.ListHostedZonesByNameOutput
				Eventually(func() (int, error) {
					publicListHostedZoneResponse, err = rawRoute53Client.ListHostedZonesByNameWithContext(ctx, &route53.ListHostedZonesByNameInput{
						DNSName:  awssdk.String("apublic.test.example.com"),
						MaxItems: awssdk.String("1"),
					})
					return len(publicListHostedZoneResponse.HostedZones), err
				}, "2s", "100ms").Should(Equal(1))

				actualTags, err := rawRoute53Client.ListTagsForResourceWithContext(ctx, &route53.ListTagsForResourceInput{
					ResourceId:   publicListHostedZoneResponse.HostedZones[0].Id,
					ResourceType: awssdk.String("hostedzone"),
				})
				Expect(err).NotTo(HaveOccurred())

				Expect(actualTags.ResourceTagSet.Tags).To(ContainElement(&route53.Tag{
					Key:   awssdk.String("Name"),
					Value: awssdk.String("jose"),
				}))
				Expect(actualTags.ResourceTagSet.Tags).To(ContainElement(&route53.Tag{
					Key:   awssdk.String("something"),
					Value: awssdk.String("else"),
				}))
			})

			When("the hosted zone already exists", func() {
				BeforeEach(func() {
					now := time.Now()
					_, err = rawRoute53Client.CreateHostedZone(&route53.CreateHostedZoneInput{
						CallerReference: awssdk.String(fmt.Sprintf("1%d", now.UnixNano())),
						Name:            awssdk.String("already.public.exists.test.example.com"),
					})
					Expect(err).NotTo(HaveOccurred())
				})

				It("doesn't return error", func() {
					hostedZoneId, err = route53Client.CreateHostedZone(ctx, logger, resolver.BuildPublicHostedZone("already.public.exists.test.example.com", tags))
					Expect(err).NotTo(HaveOccurred())

					hostedZoneResponse, err := rawRoute53Client.ListHostedZonesByNameWithContext(ctx, &route53.ListHostedZonesByNameInput{
						DNSName:  awssdk.String("already.public.exists.test.example.com"),
						MaxItems: awssdk.String("1"),
					})
					Expect(err).NotTo(HaveOccurred())
					Expect(len(hostedZoneResponse.HostedZones)).To(Equal(1))
				})
			})
		})

		Context("we want a private hosted zone", func() {
			var cluster resolver.Cluster
			var dnsZone resolver.DnsZone

			BeforeEach(func() {
				cluster = resolver.Cluster{
					Name:   "aprivate",
					Region: Region,
					VPCId:  VPCId,
				}
				dnsZone = resolver.BuildPrivateHostedZone("aprivate.test.example.com", cluster, tags, []string{MCVPCId})
			})

			It("creates a private hosted zone successfully", func() {
				hostedZoneId, err = route53Client.CreateHostedZone(ctx, logger, dnsZone)
				Expect(err).NotTo(HaveOccurred())

				var privateHostedZoneResponse *route53.ListHostedZonesByNameOutput
				Eventually(func() (int, error) {
					privateHostedZoneResponse, err = rawRoute53Client.ListHostedZonesByNameWithContext(ctx, &route53.ListHostedZonesByNameInput{
						DNSName:  awssdk.String("aprivate.test.example.com"),
						MaxItems: awssdk.String("1"),
					})
					return len(privateHostedZoneResponse.HostedZones), err
				}, "2s", "100ms").Should(Equal(1))
				Expect(*privateHostedZoneResponse.HostedZones[0].Name).To(Equal("aprivate.test.example.com."))

				actualTags, err := rawRoute53Client.ListTagsForResourceWithContext(ctx, &route53.ListTagsForResourceInput{
					ResourceId:   privateHostedZoneResponse.HostedZones[0].Id,
					ResourceType: awssdk.String("hostedzone"),
				})
				Expect(err).NotTo(HaveOccurred())

				Expect(actualTags.ResourceTagSet.Tags).To(ContainElement(&route53.Tag{
					Key:   awssdk.String("Name"),
					Value: awssdk.String("jose"),
				}))

				associatedHostedZones, err := rawRoute53Client.ListHostedZonesByVPCWithContext(ctx, &route53.ListHostedZonesByVPCInput{
					VPCId:     awssdk.String(MCVPCId),
					VPCRegion: awssdk.String(Region),
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(associatedHostedZones.HostedZoneSummaries).To(ContainElement(&route53.HostedZoneSummary{
					HostedZoneId: awssdk.String(strings.TrimPrefix(*privateHostedZoneResponse.HostedZones[0].Id, "/hostedzone/")),
					Name:         awssdk.String("aprivate.test.example.com."),
					Owner: &route53.HostedZoneOwner{
						OwningAccount: awssdk.String("000000000000"),
					},
				}))
			})

			When("the hosted zone already exists", func() {
				BeforeEach(func() {
					now := time.Now()
					_, err = rawRoute53Client.CreateHostedZone(&route53.CreateHostedZoneInput{
						CallerReference: awssdk.String(fmt.Sprintf("1%d", now.UnixNano())),
						HostedZoneConfig: &route53.HostedZoneConfig{
							Comment:     awssdk.String("Zone for CAPI cluster"),
							PrivateZone: awssdk.Bool(true),
						},
						Name: awssdk.String("already.private.exists.test.example.com"),
						VPC: &route53.VPC{
							VPCId:     awssdk.String(VPCId),
							VPCRegion: awssdk.String(Region),
						},
					})
					Expect(err).NotTo(HaveOccurred())
				})

				It("doesn't return error", func() {
					hostedZoneId, err = route53Client.CreateHostedZone(ctx, logger, dnsZone)
					Expect(err).NotTo(HaveOccurred())

					hostedZoneResponse, err := rawRoute53Client.ListHostedZonesByNameWithContext(ctx, &route53.ListHostedZonesByNameInput{
						DNSName:  awssdk.String("already.private.exists.test.example.com"),
						MaxItems: awssdk.String("1"),
					})
					Expect(err).NotTo(HaveOccurred())
					Expect(len(hostedZoneResponse.HostedZones)).To(Equal(1))
				})
			})
		})
	})

	When("fetching hosted zone id by name", func() {
		var hostedZoneToFind, differentHostedZoneToFind *route53.CreateHostedZoneOutput
		BeforeEach(func() {
			now := time.Now()
			hostedZoneToFind, err = rawRoute53Client.CreateHostedZone(&route53.CreateHostedZoneInput{
				CallerReference: awssdk.String(fmt.Sprintf("1%d", now.UnixNano())),
				Name:            awssdk.String("findid.test.example.com"),
			})
			Expect(err).NotTo(HaveOccurred())

			differentHostedZoneToFind, err = rawRoute53Client.CreateHostedZone(&route53.CreateHostedZoneInput{
				CallerReference: awssdk.String(fmt.Sprintf("1%d", now.UnixNano())),
				Name:            awssdk.String("different.test.example.com"),
			})
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			_, err = rawRoute53Client.DeleteHostedZoneWithContext(ctx, &route53.DeleteHostedZoneInput{Id: hostedZoneToFind.HostedZone.Id})
			Expect(err).NotTo(HaveOccurred())
			_, err = rawRoute53Client.DeleteHostedZoneWithContext(ctx, &route53.DeleteHostedZoneInput{Id: differentHostedZoneToFind.HostedZone.Id})
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns the id", func() {
			hostedZoneId, err := route53Client.GetHostedZoneIdByName(ctx, logger, "findid.test.example.com")
			Expect(err).NotTo(HaveOccurred())
			Expect(hostedZoneId).To(Equal(*hostedZoneToFind.HostedZone.Id))
		})

		It("returns error when hosted zone does not exist", func() {
			_, err = route53Client.GetHostedZoneIdByName(ctx, logger, "nonexisting.test.example.com.")
			Expect(err).To(HaveOccurred())
		})
	})

	When("adding delegation to parent hosted zone", func() {
		var parentHostedZoneToFind *route53.CreateHostedZoneOutput
		var hostedZoneToFind *route53.CreateHostedZoneOutput

		BeforeEach(func() {
			now := time.Now()
			parentHostedZoneToFind, err = rawRoute53Client.CreateHostedZone(&route53.CreateHostedZoneInput{
				CallerReference: awssdk.String(fmt.Sprintf("1%d", now.UnixNano())),
				Name:            awssdk.String("test.example.com"),
			})
			Expect(err).NotTo(HaveOccurred())

			hostedZoneToFind, err = rawRoute53Client.CreateHostedZone(&route53.CreateHostedZoneInput{
				CallerReference: awssdk.String(fmt.Sprintf("1%d", now.UnixNano())),
				Name:            awssdk.String("different.test.example.com"),
			})
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			_, err = rawRoute53Client.DeleteHostedZoneWithContext(ctx, &route53.DeleteHostedZoneInput{Id: parentHostedZoneToFind.HostedZone.Id})
			Expect(err).NotTo(HaveOccurred())
			_, err = rawRoute53Client.DeleteHostedZoneWithContext(ctx, &route53.DeleteHostedZoneInput{Id: hostedZoneToFind.HostedZone.Id})
			Expect(err).NotTo(HaveOccurred())
		})

		It("creates the dns records", func() {
			listRecordSets, err := rawRoute53Client.ListResourceRecordSets(&route53.ListResourceRecordSetsInput{
				HostedZoneId: hostedZoneToFind.HostedZone.Id,
				MaxItems:     awssdk.String("1"), // First entry is always NS record
			})
			Expect(err).NotTo(HaveOccurred())

			err = route53Client.AddDelegationToParentZone(ctx, logger, *parentHostedZoneToFind.HostedZone.Id, *hostedZoneToFind.HostedZone.Id)
			Expect(err).NotTo(HaveOccurred())

			listParentRecordSets, err := rawRoute53Client.ListResourceRecordSets(&route53.ListResourceRecordSetsInput{
				HostedZoneId: parentHostedZoneToFind.HostedZone.Id,
			})
			Expect(err).NotTo(HaveOccurred())

			found := false
			for _, recordSet := range listParentRecordSets.ResourceRecordSets {
				if *recordSet.Name == *hostedZoneToFind.HostedZone.Name {
					found = true
					Expect(recordSet.ResourceRecords).To(Equal(listRecordSets.ResourceRecordSets[0].ResourceRecords))
				}
			}
			Expect(found).To(BeTrue())
		})
	})

	When("deleting a hosted zone", func() {
		When("the zone exists", func() {
			var hostedZoneToFind *route53.CreateHostedZoneOutput

			BeforeEach(func() {
				now := time.Now()

				hostedZoneToFind, err = rawRoute53Client.CreateHostedZone(&route53.CreateHostedZoneInput{
					CallerReference: awssdk.String(fmt.Sprintf("1%d", now.UnixNano())),
					Name:            awssdk.String("deleting.test.example.com"),
				})
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				// We don't check error here because we actually expect an error since we just removed the hosted zone.
				_, err = rawRoute53Client.DeleteHostedZoneWithContext(ctx, &route53.DeleteHostedZoneInput{Id: hostedZoneToFind.HostedZone.Id})
			})

			It("deletes the zone", func() {
				err = route53Client.DeleteHostedZone(ctx, logger, *hostedZoneToFind.HostedZone.Id)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() (int, error) {
					listHostedZoneResponse, err := rawRoute53Client.ListHostedZonesByName(&route53.ListHostedZonesByNameInput{
						DNSName:  awssdk.String("deleting.test.example.com"),
						MaxItems: awssdk.String("1"),
					})
					return len(listHostedZoneResponse.HostedZones), err
				}, "3s", "500ms").Should(BeZero())
			})
		})

		When("the zone doesn't exist", func() {
			It("deletes the zone", func() {
				err = route53Client.DeleteHostedZone(ctx, logger, "non-existing-zone.example.com")
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	When("deleting delegation to parent hosted zone", func() {
		var parentHostedZoneToFind *route53.CreateHostedZoneOutput
		var hostedZoneToFind *route53.CreateHostedZoneOutput

		BeforeEach(func() {
			now := time.Now()
			parentHostedZoneToFind, err = rawRoute53Client.CreateHostedZone(&route53.CreateHostedZoneInput{
				CallerReference: awssdk.String(fmt.Sprintf("1%d", now.UnixNano())),
				Name:            awssdk.String("test.example.com"),
			})
			Expect(err).NotTo(HaveOccurred())

			hostedZoneToFind, err = rawRoute53Client.CreateHostedZone(&route53.CreateHostedZoneInput{
				CallerReference: awssdk.String(fmt.Sprintf("1%d", now.UnixNano())),
				Name:            awssdk.String("different.test.example.com"),
			})
			Expect(err).NotTo(HaveOccurred())

			err = route53Client.AddDelegationToParentZone(ctx, logger, *parentHostedZoneToFind.HostedZone.Id, *hostedZoneToFind.HostedZone.Id)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			_, err = rawRoute53Client.DeleteHostedZoneWithContext(ctx, &route53.DeleteHostedZoneInput{Id: parentHostedZoneToFind.HostedZone.Id})
			Expect(err).NotTo(HaveOccurred())
			_, err = rawRoute53Client.DeleteHostedZoneWithContext(ctx, &route53.DeleteHostedZoneInput{Id: hostedZoneToFind.HostedZone.Id})
			Expect(err).NotTo(HaveOccurred())
		})

		It("deletes the delegation from the parent zone", func() {
			err = route53Client.DeleteDelegationFromParentZone(ctx, logger, *parentHostedZoneToFind.HostedZone.Id, *hostedZoneToFind.HostedZone.Id)
			Expect(err).NotTo(HaveOccurred())

			listParentRecordSets, err := rawRoute53Client.ListResourceRecordSets(&route53.ListResourceRecordSetsInput{
				HostedZoneId: parentHostedZoneToFind.HostedZone.Id,
			})
			Expect(err).NotTo(HaveOccurred())

			found := false
			for _, recordSet := range listParentRecordSets.ResourceRecordSets {
				if *recordSet.Name == *hostedZoneToFind.HostedZone.Name {
					found = true
				}
			}
			Expect(found).To(BeFalse())
		})
	})
})