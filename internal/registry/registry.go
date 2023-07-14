package registry

import (
	"fmt"
	"strings"

	"github.com/superproj/dtags/internal/registry/artifactregistry"
	"github.com/superproj/dtags/internal/registry/dockerhub"
	"github.com/superproj/dtags/internal/registry/ecr"
	"github.com/superproj/dtags/internal/registry/ecrpublic"
	"github.com/superproj/dtags/internal/registry/gcr"
)

type Registry interface {
	ListTags(name string) ([]string, error)
}

type Config struct {
	AWSProfile string
}

func New(domain string, cfg *Config) (Registry, error) {
	switch {
	case domain == "docker.io":
		return dockerhub.New(), nil
	case domain == "public.ecr.aws":
		return ecrpublic.New(&ecrpublic.Config{Profile: cfg.AWSProfile})
	case strings.HasSuffix(domain, "amazonaws.com"):
		// <AWS_ACCOUNT_ID>.dkr.ecr.<REGION>.amazonaws.com/<REPOSITORY_NAME>
		return ecr.New(&ecr.Config{Profile: cfg.AWSProfile, Domain: domain})
	case domain == "gcr.io":
		return gcr.New(), nil
	case strings.HasSuffix(domain, "-docker.pkg.dev"):
		// <LOCATION>-docker.pkg.dev/<PROJECT>/<REPOSITORY>/<PACKAGE>
		return artifactregistry.New(&artifactregistry.Config{Domain: domain}), nil
	default:
		return nil, fmt.Errorf("unsupported image registry: %s", domain)
	}
}
