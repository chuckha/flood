package commands

import (
	"encoding/json"
	"fmt"
	"github.com/chuckha/flood/api"
)

var Image *Command

func init() {
	Image = &Command{
		Name:  "image",
		Short: "manage droplet images",
		Commands: []*Command{
			listImages,
		},
		Run: imageRun,
		Template: `
flood image manages droplets

Usage:

        flood image <command>

Commands:
{{range .}}
        {{.Name | printf "%-11s"}} {{.Short}}{{end}}
`,
	}
	Image.Long = BuildLong(Image.Template, Image.Commands)
}

func imageRun(cmd *Command, args []string) error {
	if len(args) < 1 {
		fmt.Println(cmd.Long)
		return nil
	}
	for _, subcmd := range cmd.Commands {
		if subcmd.Name == args[0] && subcmd.Run != nil {
			return subcmd.Run(subcmd, args[1:])
		}
	}
	return nil
}

const imageResource = "images"

var listImages = &Command{
	Name:  "list",
	Short: "List all available images",
	Run: func(cmd *Command, args []string) error {
		bytes, err := api.Call(imageResource, "", "list")
		if err != nil {
			return err
		}
		resp := &api.ImageListResponse{}
		json.Unmarshal(bytes, resp)
		indented, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			return err
		}
		fmt.Println(string(indented))
		return nil
	},
}

var showImage = &Command{}

var destroyImage = &Command{}

// TODO: implement this one
var transferImage = &Command{}
