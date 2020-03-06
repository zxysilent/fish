package cmds

import (
	"flag"
	"os"
	"strings"
)

// Command is the unit of execution
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// UsageLine is the one-line Usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'go help' output.
	Short string

	// Long is the long message shown in the 'go help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// DiyFlags 自定义解析 未使用
	DiyFlags bool
}

var AvailableCommands = []*Command{}

// Name 名称
func (cmd *Command) Name() string {
	name := cmd.UsageLine
	idx := strings.Index(name, " ")
	if idx >= 0 {
		name = name[:idx]
	}
	return name
}

// Usage 帮助信息
func (cmd *Command) Usage() {
	Tips("cmd", cmd)
	os.Exit(0)
}

// Args  参数
func (cmd *Command) Args() map[string]string {
	args := make(map[string]string, 4)
	cmd.Flag.VisitAll(func(flg *flag.Flag) {
		defs := flg.DefValue
		if len(defs) > 0 {
			args[flg.Name+"="+defs] = flg.Usage
		} else {
			args[flg.Name] = flg.Usage
		}
	})
	return args
}
