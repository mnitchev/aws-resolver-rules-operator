package aws_test

import (
	"context"

	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capa "sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"

	"github.com/aws-resolver-rules-operator/pkg/aws"
	"github.com/aws-resolver-rules-operator/pkg/resolver"
)

var _ = Describe("Prefix Lists", func() {
	var (
		ctx context.Context

		cluster *capa.AWSCluster

		prefixLists resolver.PrefixListClient
	)

	createPrefixList := func() string {
		input := &ec2.CreateManagedPrefixListInput{
			AddressFamily:  awssdk.String("IPv4"),
			MaxEntries:     awssdk.Int64(4),
			PrefixListName: awssdk.String(aws.GetPrefixListName(cluster.Name)),
		}
		out, err := rawEC2Client.CreateManagedPrefixList(input)
		Expect(err).NotTo(HaveOccurred())

		return *out.PrefixList.PrefixListArn
	}

	BeforeEach(func() {
		ctx = context.Background()
		var err error
		cluster = &capa.AWSCluster{
			ObjectMeta: metav1.ObjectMeta{
				Name: uuid.NewString(),
			},
			Spec: capa.AWSClusterSpec{
				AdditionalTags: additionalTags,
			},
		}

		prefixLists, err = awsClients.NewPrefixListClient(Region, AwsIamArn)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Apply", func() {
		It("creates a prefix list", func() {
			arn, err := prefixLists.Apply(ctx, cluster.Name, cluster.Spec.AdditionalTags)
			Expect(err).NotTo(HaveOccurred())

			prefixListID, err := aws.GetARNResourceID(arn)
			Expect(err).NotTo(HaveOccurred())

			out, err := rawEC2Client.DescribeManagedPrefixLists(&ec2.DescribeManagedPrefixListsInput{
				PrefixListIds: awssdk.StringSlice([]string{prefixListID}),
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(out.PrefixLists).To(HaveLen(1))
		})

		When("the prefix list already exists", func() {
			var originalID string

			BeforeEach(func() {
				arn, err := prefixLists.Apply(ctx, cluster.Name, cluster.Spec.AdditionalTags)
				Expect(err).NotTo(HaveOccurred())

				originalID, err = aws.GetARNResourceID(arn)
				Expect(err).NotTo(HaveOccurred())
			})

			It("does not return an error", func() {
				arn, err := prefixLists.Apply(ctx, cluster.Name, cluster.Spec.AdditionalTags)
				Expect(err).NotTo(HaveOccurred())

				actualID, err := aws.GetARNResourceID(arn)
				Expect(err).NotTo(HaveOccurred())

				Expect(actualID).To(Equal(originalID))

				out, err := rawEC2Client.DescribeManagedPrefixLists(&ec2.DescribeManagedPrefixListsInput{
					PrefixListIds: awssdk.StringSlice([]string{actualID}),
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(out.PrefixLists).To(HaveLen(1))
			})

			When("multiple prefix lists exist for the same cluster", func() {
				BeforeEach(func() {
					createPrefixList()
				})

				It("returns an error", func() {
					_, err := prefixLists.Apply(ctx, cluster.Name, cluster.Spec.AdditionalTags)
					Expect(err).To(MatchError(ContainSubstring(
						"found unexpected number: 2 of prefix lists",
					)))
				})
			})
		})
	})

	Describe("ApplyEntry", func() {
		var (
			prefixListARN string
			prefixListID  string
			entry         resolver.PrefixListEntry
		)

		BeforeEach(func() {
			prefixListARN = createPrefixList()
			var err error
			prefixListID, err = aws.GetARNResourceID(prefixListARN)
			Expect(err).NotTo(HaveOccurred())

			entry = resolver.PrefixListEntry{
				PrefixListARN: prefixListARN,
				CIDR:          "10.1.0.0/24",
				Description:   "the-description",
			}
		})

		It("adds the entry to the prefix list", func() {
			err := prefixLists.ApplyEntry(ctx, entry)
			Expect(err).NotTo(HaveOccurred())

			out, err := rawEC2Client.GetManagedPrefixListEntries(&ec2.GetManagedPrefixListEntriesInput{
				PrefixListId: awssdk.String(prefixListID),
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(out.Entries).To(HaveLen(1))
		})

		When("the entry already exists", func() {
			BeforeEach(func() {
				err := prefixLists.ApplyEntry(ctx, entry)
				Expect(err).NotTo(HaveOccurred())
			})

			It("does not return an error", func() {
				err := prefixLists.ApplyEntry(ctx, entry)
				Expect(err).NotTo(HaveOccurred())
			})

			When("the entry's description doesn't match the desired one", func() {
				It("returns an error", func() {
					entry.Description = "something-else"
					err := prefixLists.ApplyEntry(ctx, entry)
					Expect(err).To(HaveOccurred())
				})
			})
		})

		When("the cidr is not valid", func() {
			It("returns an error", func() {
				entry.CIDR = "not-a-valid-cidr"
				err := prefixLists.ApplyEntry(ctx, entry)
				Expect(err).To(HaveOccurred())
			})
		})

		When("the arn is invalid", func() {
			It("returns an error", func() {
				entry.PrefixListARN = "invalid-arn"
				err := prefixLists.ApplyEntry(ctx, entry)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("DeleteEntry", func() {
		var (
			prefixListARN string
			prefixListID  string
			entry         resolver.PrefixListEntry
		)

		BeforeEach(func() {
			prefixListARN = createPrefixList()
			var err error
			prefixListID, err = aws.GetARNResourceID(prefixListARN)
			Expect(err).NotTo(HaveOccurred())

			entry = resolver.PrefixListEntry{
				PrefixListARN: prefixListARN,
				CIDR:          "10.1.0.0/24",
				Description:   "the-description",
			}
			err = prefixLists.ApplyEntry(ctx, entry)
			Expect(err).NotTo(HaveOccurred())
		})

		It("removes the entry to the prefix list", func() {
			err = prefixLists.DeleteEntry(ctx, entry)
			Expect(err).NotTo(HaveOccurred())

			out, err := rawEC2Client.GetManagedPrefixListEntries(&ec2.GetManagedPrefixListEntriesInput{
				PrefixListId: awssdk.String(prefixListID),
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(out.Entries).To(HaveLen(0))
		})

		When("the entry is already deleted", func() {
			BeforeEach(func() {
				err := prefixLists.DeleteEntry(ctx, entry)
				Expect(err).NotTo(HaveOccurred())
			})

			It("does not return an error", func() {
				err = prefixLists.DeleteEntry(ctx, entry)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		When("the cidr is not valid", func() {
			It("returns an error", func() {
				entry.CIDR = "not-a-valid-cidr"
				err := prefixLists.ApplyEntry(ctx, entry)
				Expect(err).To(HaveOccurred())
			})
		})

		When("the arn is invalid", func() {
			It("returns an error", func() {
				entry.PrefixListARN = "not-a-valid-arn"
				err := prefixLists.ApplyEntry(ctx, entry)
				Expect(err).To(HaveOccurred())
			})
		})

		When("the entry's description doesn't match the desired one", func() {
			BeforeEach(func() {
				err := prefixLists.ApplyEntry(ctx, entry)
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns an error", func() {
				entry.Description = "something-else"
				err := prefixLists.DeleteEntry(ctx, entry)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete", func() {
		var prefixListID string

		BeforeEach(func() {
			prefixListARN := createPrefixList()
			var err error
			prefixListID, err = aws.GetARNResourceID(prefixListARN)
			Expect(err).NotTo(HaveOccurred())
		})

		It("deletes the prefix list", func() {
			err := prefixLists.Delete(ctx, cluster.Name)
			Expect(err).NotTo(HaveOccurred())

			out, err := rawEC2Client.DescribeManagedPrefixLists(&ec2.DescribeManagedPrefixListsInput{
				PrefixListIds: awssdk.StringSlice([]string{prefixListID}),
			})
			Expect(err).NotTo(HaveOccurred())
			for _, prefixList := range out.PrefixLists {
				Expect(*prefixList.State).To(Equal("delete-complete"))
			}
		})

		When("the prefix list has already been deleted", func() {
			BeforeEach(func() {
				_, err := rawEC2Client.DeleteManagedPrefixList(&ec2.DeleteManagedPrefixListInput{
					PrefixListId: awssdk.String(prefixListID),
				})
				Expect(err).NotTo(HaveOccurred())
			})

			It("does not return an error", func() {
				err := prefixLists.Delete(ctx, cluster.Name)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		When("there are multiple prefix lists for the same cluster", func() {
			BeforeEach(func() {
				createPrefixList()
			})

			It("returns an error", func() {
				err := prefixLists.Delete(ctx, cluster.Name)
				Expect(err).To(MatchError(ContainSubstring(
					"found unexpected number: 2 of prefix lists",
				)))
			})
		})
	})
})
