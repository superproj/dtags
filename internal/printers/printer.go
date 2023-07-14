package printers

import (
	"fmt"
	"io"

	"github.com/Masterminds/semver"
)

type PrintParameters struct {
	Image    string
	Tags     []string
	Latest   bool
	WithName bool
}

type Printer interface {
	Print(w io.Writer, params *PrintParameters) error
}

var printers = map[string]Printer{}

func Register(name string, printer Printer) Printer {
	printers[name] = printer
	return printer
}

func (params *PrintParameters) GetTags() []string {
	if params.Latest {
		return []string{findHighestSemverTag(params.Tags)}
	}

	return params.Tags
}

func Get(name string) (Printer, error) {
	p, ok := printers[name]
	if !ok {
		return nil, fmt.Errorf("unsupported output format: %s", name)
	}
	return p, nil
}

func List() []string {
	var names []string
	for name := range printers {
		names = append(names, name)
	}
	return names
}

func formatTags(params *PrintParameters) []string {
	tags := params.GetTags()

	if params.WithName {
		for i, t := range tags {
			tags[i] = fmt.Sprintf("%s:%s", params.Image, t)
		}
	}

	return tags
}

func findHighestSemverTag(tags []string) string {
	// Find the highest semver tag
	var highestSemver *semver.Version
	for _, tag := range tags {
		version, err := semver.NewVersion(tag)
		if err != nil {
			continue
		}
		if highestSemver == nil || version.GreaterThan(highestSemver) {
			highestSemver = version
		}
	}

	return highestSemver.Original()
}
