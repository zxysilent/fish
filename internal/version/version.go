package version

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"text/template"
	"time"

	"github.com/zxysilent/fish/internal/cmds"
	. "github.com/zxysilent/fish/logger"
	"github.com/zxysilent/fish/logger/colors"
)

const version = "0.3.1"
const versionTmpl string = `%s%s
   ____  _         __ 
  / __/ (_) ___   / / 
 / _/  / / (_-<  / _ \
/_/   /_/ /___/ /_//_/ v{{ .Version }}%s
%s%s
├── Go      : {{.Go    }}
├── GOOS    : {{.GOOS  }}
├── GOARCH  : {{.GOARCH}}
├── NumCPU  : {{.NumCPU}}
├── GOPATH  : {{.GOPATH}}
├── GOROOT  : {{.GOROOT}}
└── Date    : {{.Date  }}%s
`

var CmdVersion = &cmds.Command{
	UsageLine: "version [-t=json]",
	Short:     "print fish version",
	Long: `
Prints the current Fish 、Go version and platform information.`,
	Run: runVersion,
}
var enctype string

func init() {
	CmdVersion.Flag.StringVar(&enctype, "t", "", "Set the output type. eg: json")
	cmds.Regcmd(CmdVersion)
}

func runVersion(cmd *cmds.Command, args []string) {
	fishEnvs := struct {
		Go      string
		GOOS    string
		GOARCH  string
		NumCPU  int
		GOPATH  string
		GOROOT  string
		Version string
		Date    string
	}{
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.NumCPU(),
		os.Getenv("GOPATH"),
		runtime.GOROOT(),
		version,
		time.Now().Format("2006-01-02 15:04:05"),
	}
	if enctype == "json" {
		b, err := json.MarshalIndent(fishEnvs, "", "    ")
		if err != nil {
			Flog.Error(err.Error())
		}
		fmt.Println(string(b))
		return
	}
	tmpl, err := template.New("version").Parse(fmt.Sprintf(versionTmpl, "\x1b[35m", "\x1b[1m", "\x1b[0m", "\x1b[36m", "\x1b[1m", "\x1b[0m"))
	if err != nil {
		Flog.Fatalf("Cannot parse the version template: %s", err)
	}
	err = tmpl.Execute(colors.NewColorWriter(os.Stdout), fishEnvs)
	if err != nil {
		Flog.Error(err.Error())
	}
}
