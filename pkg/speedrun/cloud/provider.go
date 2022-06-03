package cloud

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/spf13/viper"
)

type PlatformProvider interface {
	getInstances(context.Context) ([]Instance, error)
}

type ProviderType string

const (
	ProviderGCP = ProviderType("gcp")
	ProviderAWS = ProviderType("aws")
)

func GetProviders() ([]PlatformProvider, error) {

	providers := []PlatformProvider{}
	if (viper.IsSet("gcp.credentials") || viper.IsSet("gcp.project")) ||
		viper.GetBool("gcp.enabled") {

		gc, err := newGCPClient()
		if err != nil {
			log.WithError(err).Error("could not initialize GCP client")
		} else {
			providers = append(providers, gc)
		}
	}

	if (viper.IsSet("aws.credentials") || viper.IsSet("aws.region")) ||
		viper.GetBool("aws.enabled") {
		aws, err := newAWSClient()
		if err != nil {
			log.WithError(err).Error("could not initialize AWS client")
		} else {
			providers = append(providers, aws)
		}
	}

	if len(providers) == 0 {
		return nil, fmt.Errorf("no cloud providers enabled")
	}
	return providers, nil
}
