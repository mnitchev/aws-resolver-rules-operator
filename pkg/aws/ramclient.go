package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
)

type AWSRAM struct {
	client *ram.RAM
}

func (a *AWSRAM) CreateResourceShareWithContext(ctx context.Context, resourceShareName string, allowExternalPrincipals bool, resourceArns, principals []string) (string, error) {
	response, err := a.client.CreateResourceShareWithContext(ctx, &ram.CreateResourceShareInput{
		AllowExternalPrincipals: aws.Bool(allowExternalPrincipals),
		Name:                    aws.String(resourceShareName),
		Principals:              aws.StringSlice(principals),
		ResourceArns:            aws.StringSlice(resourceArns),
	})
	if err != nil {
		return "", errors.WithStack(err)
	}

	return *response.ResourceShare.ResourceShareArn, nil
}

func (a *AWSRAM) DeleteResourceShareWithContext(ctx context.Context, logger logr.Logger, resourceShareName string) error {
	logger.Info("Trying to find RAM resource share", "resourceShareName", resourceShareName)
	resourceShare, err := a.client.GetResourceShares(&ram.GetResourceSharesInput{
		Name:          aws.String(resourceShareName),
		ResourceOwner: aws.String("SELF"),
	})
	if err != nil {
		return errors.WithStack(err)
	}

	if len(resourceShare.ResourceShares) < 1 {
		return nil
	}

	logger.Info("Deleting RAM resource share", "resourceShareName", resourceShareName)
	_, err = a.client.DeleteResourceShareWithContext(ctx, &ram.DeleteResourceShareInput{
		ResourceShareArn: resourceShare.ResourceShares[0].ResourceShareArn,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}