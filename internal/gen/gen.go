// +build gen

package gen

import (
	"bytes"
	_ "embed"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/zxysilent/fish/internal/cmds"
)

//go:embed tmpl/model.tpl
var modelTmpl string

//go:embed tmpl/control.tpl
var controlTmpl string

var CmdGen = &cmds.Command{
	UsageLine: "gen [-t=m/c] [-n=name] [-r=notes] [-o=std/file]",
	Short:     "generate code",
	Long: `
Generate custom code echo/xorm.`,
	Run: runGen,
}
var genTmpls *template.Template // 模板集合
var (
	class string
	name  string
	notes string
	out   string
)

func init() {
	CmdGen.Flag.StringVar(&class, "t", "m", "code type c=control,m=model")
	CmdGen.Flag.StringVar(&name, "n", "Demo", "main name")
	CmdGen.Flag.StringVar(&notes, "r", "示例", "注释信息")
	CmdGen.Flag.StringVar(&out, "o", "std", "输出路径 std/file")
	genTmpls = template.New("gen")
	genTmpls.New("model").Parse(modelTmpl)
	genTmpls.New("control").Parse(controlTmpl)
	cmds.Regcmd(CmdGen)
}
func runGen(cmd *cmds.Command, args []string) {
	mod := map[string]string{
		"Name":  name,
		"Notes": notes,
		"Path":  strings.ToLower(name),
	}
	switch class {
	case "c":
		renderControl(mod)
	case "m":
		renderModel(mod)
	}
}
func renderControl(data interface{}) {
	if out == "std" {
		genTmpls.ExecuteTemplate(os.Stdout, "control", data)
	} else {
		buf := &bytes.Buffer{}
		genTmpls.ExecuteTemplate(buf, "control", data)
		ioutil.WriteFile(strings.ToLower(name)+".go", buf.Bytes(), 0666)
	}
}
func renderModel(data interface{}) {
	if out == "std" {
		genTmpls.ExecuteTemplate(os.Stdout, "model", data)
	} else {
		buf := &bytes.Buffer{}
		genTmpls.ExecuteTemplate(buf, "model", data)
		ioutil.WriteFile(strings.ToLower(name)+".go", buf.Bytes(), 0666)
	}
}
