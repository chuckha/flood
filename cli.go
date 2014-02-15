package main

import (
	"flag"
	"github.com/chuckha/flood/commands"
	"io"
	"os"
	"text/template"
)

var cmds = []*commands.Command{
	commands.Droplet,
	commands.Size,
	commands.Image,
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
	}
	// print help of the command if args[0] == help
	for _, cmd := range cmds {
		if cmd.Name == args[0] {
			cmd.Run(cmd, args[1:])
			return
		}
	}
	usage()
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	//	t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": capitalize})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}
func printUsage(w io.Writer) {
	tmpl(w, usageTemplate, cmds)
}
func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

var apis = `
flood droplet list
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
flood droplet destroy 1234

flood region list

flood image list
flood image show 1234
flood image destroy 1234
flood image transfer 1234

flood ssh_key list
flood ssh_key create
flood ssh_key edit 1234
flood ssh_key destroy 1234

flood size list

flood domain list
flood domain create
flood domain show 1234
flood domain destroy 1234
flood domain 1234 records list
flood domain 1234 records create
flood domain 1234 records show 1234
flood domain 1234 records edit 1234
flood domain 1234 records destroy 1234

flood event show 1234
`

var usageTemplate = `flood is a tool for interfacing with the Digital Ocean API.

General usage pattern:

        flood resource action [id]

The actions by resource are:
{{range .}}
        {{.Name | printf "%-11s"}} {{.Short}}{{end}}
`
