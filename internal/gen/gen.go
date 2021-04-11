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
	name  string //名称
	notes string //注释
)

func init() {
	CmdGen.Flag.StringVar(&name, "n", "Demo", "main name")
	CmdGen.Flag.StringVar(&notes, "r", "示例", "注释信息")
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
	fileName := strings.ToLower(name) + ".go"
	os.MkdirAll("gens/model/", 0666)
	os.MkdirAll("gens/control/", 0666)
	// 渲染model
	buf := &bytes.Buffer{}
	genTmpls.ExecuteTemplate(buf, "model", mod)
	ioutil.WriteFile("gens/model/"+fileName, buf.Bytes(), 0666)
	// 渲染 control
	buf.Reset()
	genTmpls.ExecuteTemplate(buf, "control", mod)
	ioutil.WriteFile("gens/control/"+fileName, buf.Bytes(), 0666)
}
