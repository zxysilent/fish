package cmds

import (
	"os"
	"text/template"

	. "github.com/zxysilent/fish/logger"
	"github.com/zxysilent/fish/logger/colors"
)

var (
	Fishs     []*Command         // 命令集🧑
	fishTmpls *template.Template // 模板集合
)

// Regcmd 注册命令
func Regcmd(cmd *Command) {
	Fishs = append(Fishs, cmd)
}

//帮助模板
var usageTmpl = `Fish is a tool for managing your go application.

{{"Usage:" | head}}

    {{"fish command [arguments]" | cyanbold}}

{{"The commands are:" | head}}
{{range .}}
    {{.Name | printf "%-10s" | cyanbold}} {{.Short}}{{end}}

Use {{"fish help <command>" | cyanbold}} for more information about a command.
`

// 帮助模板
var helpTmpl = `{{"Usage" | head}}
  {{.UsageLine | printf "fish %s" | cyanbold}}
{{if .Args}}
{{"Arguments" | head}}{{range $k,$v := .Args}}
  {{$k | printf "-%s" | cyanbold}}
 	 {{$v}}
  {{end}}{{end}}
{{.Long}}
`

// 错误模板
var errorTmpl = `fish: {{.}}.
Use {{"fish help" | cyanbold}} for more information.
`

func init() {
	funcs := template.FuncMap{
		"cyanbold": colors.CyanBold,
		"head":     colors.MagentaBold,
	}
	fishTmpls = template.New("fish").Funcs(funcs)
	fishTmpls.New("usage").Parse(usageTmpl)
	fishTmpls.New("help").Parse(helpTmpl)
	fishTmpls.New("error").Parse(errorTmpl)
}

// 帮助信息
func Usage() {
	Tips("usage", Fishs)
}

// Help 帮助命令
func Help(args []string) {
	if len(args) == 0 {
		Usage()
		return
	}
	if len(args) != 1 {
		Tips("error", "Too many arguments")
		return
	}
	arg := args[0]
	for _, cmd := range Fishs {
		if cmd.Name() == arg {
			Tips("help", cmd)
			return
		}
	}
	Tips("error", "Unknown help command")
}

// 提示输出
func Tips(name string, data interface{}) {
	err := fishTmpls.ExecuteTemplate(colors.NewColorWriter(os.Stdout), name, data)
	if err != nil {
		Flog.Error(err.Error())
	}
}
