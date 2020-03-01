package main

import (
	"flag"
	"os"

	"github.com/zxysilent/fish/internal/cmds"
	"github.com/zxysilent/fish/internal/run"
	"github.com/zxysilent/fish/internal/version"
)

func init() {
	// 注册所有支持的命令
	cmds.Fishs = []*cmds.Command{version.CmdVersion, run.CmdRun}
}
func main() {
	flag.Usage = cmds.Usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		cmds.Usage()
		os.Exit(0)
		return
	}
	if args[0] == "help" {
		cmds.Help(args[1:])
		os.Exit(0)
		return
	}
	for _, cmd := range cmds.Fishs {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = cmd.Usage
			// 自定义命令参数解析
			if cmd.DiyFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			os.Exit(0)
			return
		}
	}
	cmds.Tips("error", "Unknown subcommand")
}
