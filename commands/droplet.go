package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chuckha/flood/api"
	"os"
)

var Droplet = NewCommand("droplet", "manages droplets",
	list, create, show, reboot)

const dropletResource = "droplets"

var list = &Command{
	Name:  "list",
	Short: "list active droplets",
	Run: func(cmd *Command, args []string) error {
		bytes, err := api.Call(dropletResource, "", "list")
		if err != nil {
			return err
		}
		resp := &api.DropletListResponse{}
		json.Unmarshal(bytes, resp)
		indented, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			return err
		}
		fmt.Println(string(indented))
		return nil
	},
}
var create = &Command{
	Name:  "create",
	Short: "create a new droplet",
	Long: `Usage:

flood droplet create <name> <size_id> <image_id> <region_id> [<ssh_key_ids>, <private_networking>, <backups_enabled>]

To see various pieces of data:

flood size list # for sizes
flood image list # for images

`,
	Run: Require(func(cmd *Command, args []string) error {
		bytes, err := api.Call(dropletResource, "", "create")
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
		return nil
	}, RequireArgsErr),
}
var show = &Command{
	Name:  "show",
	Short: "get information on one droplet",
	Long: `Usage:

flood droplet show <id>
`,
	Run: Require(func(cmd *Command, args []string) error {
		bytes, err := api.Call(dropletResource, "", "show")
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
		return nil
	}, RequireIdErr),
}
var reboot = &Command{
	Name:  "reboot",
	Short: "reboot a droplet",
	Long: `Usage:

flood droplet reboot <id>
`,
	Run: Require(func(cmd *Command, args []string) error {
		bytes, err := api.Call(dropletResource, "", "reboot")
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
		return nil
	}, RequireIdErr),
}

func Require(fn func(*Command, []string) error, err error) func(*Command, []string) error {
	return func(cmd *Command, args []string) error {
		if len(args) < 1 {
			Error(cmd.Long, err)
		}
		return fn(cmd, args)
	}
}

var (
	RequireIdErr   = errors.New("please supply an id")
	RequireArgsErr = errors.New("more arguments are required")
)

func Error(help string, err error) {
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stderr, help)
	os.Exit(2)
}

var apis = `                                                                              flood droplet list
flood droplet create
flood droplet show 1234
flood droplet reboot 1234
flood droplet power_cycle 1234
flood droplet shutdown 1234
flood droplet power_off 1234
flood droplet power_on 1234
flood droplet password_reset 1234
flood droplet resize 1234
flood droplet snapshot 1234
flood droplet restore 1234
flood droplet rebuild 1234
flood droplet rename 1234
flood droplet destroy 1234`
