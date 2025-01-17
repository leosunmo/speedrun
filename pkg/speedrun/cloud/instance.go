package cloud

import (
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
}

func (i Instance) GetAddress(private bool) string {
	if private {
		return i.PrivateAddress
	}

	return i.PublicAddress
}

func GetInstances(target string) ([]Instance, error) {
	project := viper.GetString("gcp.projectid")
	if project == "" {
		return nil, fmt.Errorf("missing Google Cloud project ID, consider adding it to your config file: %s", viper.ConfigFileUsed())
	}

	gcpClient, err := NewGCPClient()
	if err != nil {
		return nil, err
	}

	log.Info("Fetching instance list")
	instances, err := gcpClient.GetInstances(project)
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

func filter(instnces []Instance, target string) ([]Instance, error) {
	if target == "" {
		return instnces, nil
	}

	var subset []Instance

	program, err := expr.Compile(target, expr.Env(Instance{}), expr.AsBool())
	if err != nil {
		return nil, err
	}

	for _, instance := range instnces {
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
