package buildtool

import (
	"fmt"
	"io"

	"github.com/iancoleman/strcase"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typicmd/buildtool/markdown"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typiobj"
)

const (
	configTemplate = `| Key | Type | Default | Required | Description |	
|---|---|---|---|---|{{range .}}
|{{usage_key .}}|{{usage_type .}}|{{usage_default .}}|{{usage_required .}}|{{usage_description .}}|{{end}}`
)

// Readme detail
type Readme struct {
	*typictx.Context
}

// Markdown to return the markdown
func (r Readme) Markdown(w io.Writer) *markdown.Markdown {
	md := &markdown.Markdown{Writer: w}
	md.Comment("Autogenerated by Typical-Go. DO NOT EDIT.")
	if r.Name != "" {
		md.Heading1(r.Name)
	} else {
		md.Heading1("Typical-Go Project")
	}
	if r.Description != "" {
		md.Writeln(r.Description)
	}
	r.prerequisite(md)
	r.runInstruction(md)
	r.configuration(md)
	r.releaseDistribution(md)
	return md
}

func (r Readme) prerequisite(md *markdown.Markdown) {
	md.Heading2("Prerequisite")
	md.OrderedList(
		"Install [Go](https://golang.org/doc/install) or `brew install go`",
	)
}

func (r Readme) runInstruction(md *markdown.Markdown) {
	md.Heading2("Run")
	md.Writeln("Use `./typicalw run` to compile and run local development. [Learn more](https://typical-go.github.io/learn-more/wrapper.html)")
}

func (r Readme) releaseDistribution(md *markdown.Markdown) (err error) {
	md.Heading2("Release Distribution")
	md.Writeln("Use `./typicalw release` to make the release. You can find the binary at `release` folder. [Learn more](https://typical-go.github.io/learn-more/release.html)")
	return
}

func (r Readme) configuration(md *markdown.Markdown) {
	md.Heading2("Configurations")
	if configurer, ok := r.Application.(typiobj.Configurer); ok {
		configTable(md, configurer.Configure().ConfigFields())
	}
	for _, module := range r.Modules {
		if name := typiobj.Name(module); name != "" {
			md.Heading3(strcase.ToCamel(name))
		}
		if description := typiobj.Description(module); description != "" {
			md.Writeln(description)
		}
		if configurer, ok := module.(typiobj.Configurer); ok {
			configTable(md, configurer.Configure().ConfigFields())
		}
	}
}

func configTable(md *markdown.Markdown, fields []typiobj.ConfigField) {
	md.WriteString("| Name | Type | Default | Required |\n")
	md.WriteString("|---|---|---|---|\n")
	for _, field := range fields {
		var required string
		if field.Required {
			required = "Yes"
		}
		md.WriteString(fmt.Sprintf("|%s|%s|%s|%s|\n",
			field.Name, field.Type, field.Default, required))
	}
	md.WriteString("\n")
}
