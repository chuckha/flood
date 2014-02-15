package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chuckha/flood/api"
)

var Size = NewCommand("size", "interact with available droplet sizes",
	sizeList)

var sizeList = &Command{
	Name:  "list",
	Short: "list available droplet sizes",
	Run: func(cmd *Command, args []string) error {
		jsonBytes, err := api.Call("sizes", "", "list")
		if err != nil {
			return err
		}
		//resp := &api.SizeListResponse{}
		var b bytes.Buffer
		err = json.Indent(&b, jsonBytes, "", "    ")
		if err != nil {
			return err
		}
		fmt.Println(b.String())
		return nil
	},
}
