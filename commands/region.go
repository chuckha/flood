package commands

import (
	"github.com/chuckha/flood/api"
)

var Region = NewCommand("region", "interact with available droplet regions",
	regionList)

var regionList = &Command{
	Name:  "list",
	Short: "list available droplet regions",
	Run: func(cmd *Command, args []string) error {
		jsonBytes, err := api.Call("regions", "", "list")
		if err != nil {
			return err
		}
		PrintResponse(jsonBytes)
		return nil
	},
}
