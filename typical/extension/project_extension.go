package extension

import (
	"fmt"

	"github.com/typical-go/typical-rest-server/typical/task/generate"
	"github.com/typical-go/typical-rest-server/typical/task/project"
	"gopkg.in/urfave/cli.v1"
)

// ProjectExtension provide standard command to see project context and configuration
type ProjectExtension struct {
	Extension
	ActionTrigger
}

// Setup go extension
func (e *ProjectExtension) Setup() error {
	return fmt.Errorf("not implement")
}

//Command for go extension
func (e *ProjectExtension) Command() cli.Command {
	return cli.Command{
		Name:      "project",
		ShortName: "proj",
		Subcommands: []cli.Command{
			{Name: "config", Usage: "Config details", Action: e.Print(project.ConfigDetail)},
			{Name: "context", Usage: "Context details", Action: e.Print(project.ContextDetail)},
			{Name: "readme", Usage: "Generate readme", Action: e.Run(generate.Readme)},
		},
	}
}