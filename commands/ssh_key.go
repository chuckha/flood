package commands

import (
	"github.com/chuckha/flood/api"
)

var SSHKey = NewCommand("sshkey", "interact with ssh keys associated sizes",
	sshKeyList)

var sshKeyList = &Command{
	Name:  "list",
	Short: "list available ssh keys",
	Run: func(cmd *Command, args []string) error {
		jsonBytes, err := api.Call("ssh_keys", "", "list")
		if err != nil {
			return err
		}
		PrintResponse(jsonBytes)
		return nil
	},
}
