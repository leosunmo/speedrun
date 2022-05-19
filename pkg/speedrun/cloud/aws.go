package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type AWSClient struct {
	ec2Client *ec2.Client
}

func NewAWSClient() (*AWSClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	ec2Client := ec2.NewFromConfig(cfg)
	return &AWSClient{ec2Client}, nil
}

func (c *AWSClient) GetInstances(ctx context.Context) ([]Instance, error) {
	descInput := &ec2.DescribeInstancesInput{
		MaxResults: aws.Int32(100),
	}

	p := ec2.NewDescribeInstancesPaginator(
		c.ec2Client,
		descInput,
		func(dipo *ec2.DescribeInstancesPaginatorOptions) {
			dipo.Limit = 20
		},
	)
	instances := []Instance{}

	for p.HasMorePages() {
		page, err := p.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, reservation := range page.Reservations {
			for _, instance := range reservation.Instances {
				kvTags := map[string]string{}
				for _, tag := range instance.Tags {
					kvTags[*tag.Key] = *tag.Value
				}

				i := Instance{
					Name:           *instance.InstanceId,
					PrivateAddress: *instance.PrivateIpAddress,
					PublicAddress:  *instance.PublicIpAddress,
					Labels:         kvTags,
				}
				instances = append(instances, i)
			}
		}
	}
	return instances, nil
}
