package cloud

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

type GoogleClient struct {
	project string
	*compute.Service
}

var (
	// ErrProjectMissing is returned when the project ID is missing
	ErrProjectMissing = errors.New("missing Google Cloud project ID")
)

func newGCPClient() (*GoogleClient, error) {
	var err error
	ctx := context.Background()

	gce, err := compute.NewService(ctx,
		option.WithCredentialsFile(viper.GetString("gcp.credentials")))
	if err != nil {
		err = fmt.Errorf("couldn't initialize GCP client: %v", err)
		return nil, err
	}

	return &GoogleClient{
		project: viper.GetString("gcp.project"),
		Service: gce,
	}, nil
}

// GetInstances returns a list of external IP addresses used for the SSH connection
func (c *GoogleClient) getInstances(ctx context.Context) ([]Instance, error) {
	instances := []Instance{}
	listCall := c.Instances.AggregatedList(c.project).Fields("nextPageToken", "items(Name,NetworkInterfaces,Labels)")
	listCall.Pages(ctx, func(list *compute.InstanceAggregatedList) error {
		for _, item := range list.Items {
			for _, instance := range item.Instances {
				i := Instance{
					Name:           instance.Name,
					PrivateAddress: instance.NetworkInterfaces[0].NetworkIP,
					PublicAddress:  instance.NetworkInterfaces[0].AccessConfigs[0].NatIP,
					platform:       ProviderGCP,
					Labels:         instance.Labels,
				}
				instances = append(instances, i)
			}
		}
		return nil
	})
	_, err := listCall.Do()
	if err != nil {
		return nil, err
	}

	return instances, nil
}
