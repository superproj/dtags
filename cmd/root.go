package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/distribution/distribution/reference"
	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
	"github.com/superproj/dtags/internal/printers"
	"github.com/superproj/dtags/internal/registry"
)

var cliPlugin string = ""

var (
	output     string
	withName   bool
	awsProfile string
	latest     bool
)

var rootCmd = &cobra.Command{
	Use:   "dtags [IMAGE]",
	Short: "Command line tool to get a list of tags for docker images.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		img := args[0]

		prtr, err := printers.Get(output)
		if err != nil {
			return err
		}

		named, err := reference.ParseNormalizedNamed(img)
		if err != nil {
			return err
		}

		d := reference.Domain(named)
		p := reference.Path(named)

		r, err := registry.New(d, &registry.Config{AWSProfile: awsProfile})
		if err != nil {
			return err
		}

		tags, err := r.ListTags(p)
		if err != nil {
			return err
		}

		if err := prtr.Print(os.Stdout, &printers.PrintParameters{
			Image:    img,
			Tags:     tags,
			Latest:   latest,
			WithName: withName,
		}); err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	if cliPlugin == "true" {
		plugin.Run(
			func(c command.Cli) *cobra.Command {
				return rootCmd
			},
			manager.Metadata{
				SchemaVersion:    "0.1.0",
				Vendor:           "Koki Sato",
				Version:          version,
				ShortDescription: "Command line tool to get a list of tags for docker images",
				URL:              "https://github.com/superproj/dtags",
			},
		)
	} else {
		if err := rootCmd.Execute(); err != nil {
			os.Exit(1)
		}
	}
}

func init() {
	if cliPlugin == "true" {
		rootCmd.Use = "tags [IMAGE]"
	}

	rootCmd.Flags().StringVarP(&output, "output", "o", "text", fmt.Sprintf("output format (%s)", strings.Join(printers.List(), "|")))
	rootCmd.Flags().BoolVarP(&withName, "with-name", "n", false, "print with image name.")
	rootCmd.Flags().StringVar(&awsProfile, "aws-profile", "", "aws profile.")
	rootCmd.Flags().BoolVar(&latest, "latest", false, "return only the latest version.")
}
