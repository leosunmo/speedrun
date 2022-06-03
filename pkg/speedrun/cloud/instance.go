package cloud

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/antonmedv/expr"
	"github.com/apex/log"
	"github.com/dpogorzelski/speedrun/pkg/common/cryptoutil"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Instance struct {
	PublicAddress  string
	PrivateAddress string
	Name           string
	Labels         map[string]string
	platform       ProviderType
}

func (i Instance) GetAddress(private bool) string {
	if private {
		return i.PrivateAddress
	}

	return i.PublicAddress
}

func GetInstancesFromProvider(ctx context.Context, pp PlatformProvider, target string) ([]Instance, error) {
	log.Info("Fetching instance list")

	instances, err := pp.getInstances(ctx)
	if err != nil {
		return nil, err
	}

	subset, err := filter(instances, target)
	if len(subset) == 0 {
		return nil, fmt.Errorf("no instances found")
	}

	if err != nil {
		err = errors.Wrap(err, "could not filter instance list")
		return nil, err
	}

	return subset, nil
}

func SetupTLS() (*tls.Config, error) {
	insecure := viper.GetBool("tls.insecure")
	caPath := viper.GetString("tls.ca")
	certPath := viper.GetString("tls.cert")
	keyPath := viper.GetString("tls.key")

	if insecure {
		log.Warn("Using insecure TLS configuration, this should be avoided in production environments")
		return cryptoutil.InsecureTLSConfig()
	} else {
		return cryptoutil.ClientTLSConfig(caPath, certPath, keyPath)
	}

}

func filter(instances []Instance, target string) ([]Instance, error) {
	if target == "" {
		return instances, nil
	}

	var subset []Instance

	program, err := expr.Compile(target, expr.Env(Instance{}), expr.AsBool())
	if err != nil {
		return nil, err
	}

	for _, instance := range instances {
		output, err := expr.Run(program, instance)
		if err != nil {
			continue
		}

		match, ok := output.(bool)
		if !ok {
			continue
		}

		if match {
			subset = append(subset, instance)
		}
	}
	return subset, nil
}
