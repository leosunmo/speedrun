package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/spf13/viper"
)

type AWSClient struct {
	ec2Client *ec2.Client
}

func newAWSClient() (*AWSClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(viper.GetString("aws.region")),
		config.WithSharedConfigProfile(viper.GetString("aws.profile")))
	if err != nil {
		return nil, err
	}

	ec2Client := ec2.NewFromConfig(cfg)
	return &AWSClient{ec2Client}, nil
}

func (c *AWSClient) getInstances(ctx context.Context) ([]Instance, error) {
	descInput := &ec2.DescribeInstancesInput{
		MaxResults: aws.Int32(20),
	}

	p := ec2.NewDescribeInstancesPaginator(c.ec2Client, descInput)
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
					platform:       ProviderAWS,
					Labels:         kvTags,
				}
				instances = append(instances, i)
			}
		}
	}
	return instances, nil
}
