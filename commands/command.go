package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"text/template"
)

type Command struct {
	Run      func(*Command, []string) error
	Name     string
	Short    string
	Long     string
	Template string
	Commands []*Command
}

func NewCommand(name, short string, cmds ...*Command) *Command {
	template := fmt.Sprintf(`flood %v %v

Usage:

        flood %v <command>

Commands:
{{range .}}
        {{.Name | printf "%%-11s"}} {{.Short}}{{end}}
`, name, short, name)
	command := &Command{
		Name:     name,
		Short:    short,
		Commands: cmds,
		Run:      cmdRun,
		Template: template,
	}
	command.Long = BuildLong(command.Template, command.Commands)
	return command
}

func cmdRun(cmd *Command, args []string) error {
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

func BuildLong(template string, data interface{}) string {
	var b bytes.Buffer
	tmpl(&b, template, data)
	return b.String()
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func PrintResponse(resp []byte) error {
	var b bytes.Buffer
	err := json.Indent(&b, resp, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(b.String())
	return nil
}
