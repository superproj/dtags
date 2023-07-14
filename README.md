# dtags

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/koki-develop/dtags)](https://github.com/superproj/dtags/releases/latest)
[![GitHub](https://img.shields.io/github/license/koki-develop/dtags)](./LICENSE)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/koki-develop/dtags/ci.yml?logo=github)](https://github.com/superproj/dtags/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/superproj/dtags)](https://goreportcard.com/report/github.com/superproj/dtags)

Command line tool to get a list of tags for docker images.
It can also be used as a docker cli plugin.

# Supported Registry

> **Note**  
> For the [Amazon ECR](https://aws.amazon.com/ecr/) and [ECR Public](https://docs.aws.amazon.com/AmazonECR/latest/public/index.html), an AWS Profile must be configured.  
> See [documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html) for details.

> **Note**  
> For the private [Google Container Registry](https://cloud.google.com/container-registry) and [Google Artifact Registry](https://cloud.google.com/artifact-registry), you must set Google's Application Default Credentials.  
> See [documentation](https://cloud.google.com/docs/authentication/application-default-credentials) for details.

- [Docker Hub](https://hub.docker.com/)
- [Amazon ECR](https://aws.amazon.com/ecr/)
- [Amazon ECR Public](https://docs.aws.amazon.com/AmazonECR/latest/public/index.html)
- [Google Container Registry](https://cloud.google.com/container-registry)
- [Google Artifact Registry](https://cloud.google.com/artifact-registry)

# Installation

## go install

```sh
$ go install github.com/superproj/dtags@latest
```

## Docker CLI Plugin

```sh
$ git clone https://github.com/superproj/dtags.git
$ cd dtags
$ make
$ mkdir -p $HOME/.docker/cli-plugins/
$ mv ./dist/dtags $HOME/.docker/cli-plugins/docker-tags
$ docker tags --help
```

## Release

Download the binary from the [releases page](https://github.com/superproj/dtags/releases/latest).

# Usage

```sh
Command line tool to get a list of tags for docker images.

Usage:
  dtags [IMAGE] [flags]

Flags:
      --aws-profile string   aws profile.
  -h, --help                 help for dtags
      --latest               return only the latest version.
  -o, --output string        output format (json|text|yaml) (default "text")
  -v, --version              version for dtags
  -n, --with-name            print with image name.
```

```sh
$ dtags alpine
latest
edge
3.9.6
3.9.5
...
```

```sh
# json format
$ dtags alpine -o json
[
  "latest",
  "edge",
  "3.9.6",
  "3.9.5",
  ...
]
```

# LICENSE

[LICENSE](./LICENSE)
