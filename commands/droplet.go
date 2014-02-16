package commands

import (
	"errors"
	"flag"
	"fmt"
	"github.com/chuckha/flood/api"
	"os"
)

func init() {
	dropletCreate.Flag.String("name", "test", "sets the name of the droplet to create")
	dropletCreate.Flag.Int("size_id", 66, "sets the size of the droplet. Defaults to smallest droplet")
	dropletCreate.Flag.Int("region_id", 4, "sets the region to create the droplet in, defaults to NYC2")
	dropletCreate.Flag.Int("image_id", 1505699, "sets the image of the droplet. Defaults to Ubuntu 13.10 x64")
	dropletCreate.Flag.String("ssh_key_ids", "", "A comma separated list of ssh key ids (found via sshkey list). These keys will be installed on the droplet. Defaults to none.")
	dropletCreate.Flag.Bool("private_networking", false, "enable private networking. Defaults to false")
	dropletCreate.Flag.Bool("backups_enabled", false, "enable backups. Defaults to false")
}

var Droplet = NewCommand("droplet", "manages droplets",
	dropletList, dropletCreate, dropletShow, dropletReboot)

const dropletResource = "droplets"

var dropletList = &Command{
	Name:  "list",
	Short: "list active droplets",
	Run: func(cmd *Command, args []string) error {
		resp, err := api.Call(dropletResource, "", "list")
		if err != nil {
			return err
		}
		return PrintResponse(resp)
	},
}
var dropletCreate = &Command{
	Name:  "create",
	Short: "create a new droplet",
	Flag:  flag.NewFlagSet("create", flag.ContinueOnError),
	Long: `Usage:

flood droplet create -name name -size_id id -image_id id -region_id id [<ssh_key_ids>, <private_networking>, <backups_enabled>]

Options:

    -name Sets the name of the flag, defaults to "test"

To see various pieces of data:

flood size list # for sizes
flood image list # for images

`,
	Run: func(cmd *Command, args []string) error {
		if len(args) >= 1 && args[0] == "help" {
			fmt.Println(cmd.Long)
			return nil
		}
		cmd.Flag.Parse(args)
		name := cmd.Flag.Lookup("name").Value.String()
		sshKeyIds := cmd.Flag.Lookup("ssh_key_ids").Value.String()
		regionId := cmd.Flag.Lookup("region_id").Value.String()
		sizeId := cmd.Flag.Lookup("size_id").Value.String()
		imageId := cmd.Flag.Lookup("image_id").Value.String()
		privateNetworking := cmd.Flag.Lookup("private_networking").Value.String()
		backupsEnabled := cmd.Flag.Lookup("backups_enabled").Value.String()

		url := api.GetUrl(dropletResource, "", "create")
		fullUrl := fmt.Sprintf("%v&%v", url, api.CreateDropletParams(
			name, sshKeyIds, sizeId, imageId,
			regionId, privateNetworking, backupsEnabled))
		resp, err := api.MakeRequest(fullUrl)
		if err != nil {
			return err
		}
		return PrintResponse(resp)
	},
}
var dropletShow = &Command{
	Name:  "show",
	Short: "get information on one droplet",
	Long: `Usage:

flood droplet show <id>
`,
	Run: Require(func(cmd *Command, args []string) error {
		resp, err := api.Call(dropletResource, "", "show")
		if err != nil {
			return err
		}
		return PrintResponse(resp)
	}, RequireIdErr),
}
var dropletReboot = &Command{
	Name:  "reboot",
	Short: "reboot a droplet",
	Long: `Usage:

flood droplet reboot <id>
`,
	Run: Require(func(cmd *Command, args []string) error {
		resp, err := api.Call(dropletResource, "", "reboot")
		if err != nil {
			return err
		}
		return PrintResponse(resp)
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
