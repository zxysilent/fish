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

//go:embed tmpl/router.tpl
var routerTmpl string

//go:embed tmpl/vue_api.tpl
var vueApiTmpl string

//go:embed tmpl/vue_router.tpl
var vueRouterTmpl string

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
	genTmpls.New("router").Parse(routerTmpl)
	genTmpls.New("control").Parse(controlTmpl)
	genTmpls.New("vue_api").Parse(vueApiTmpl)
	genTmpls.New("vue_router").Parse(vueRouterTmpl)
	cmds.Regcmd(CmdGen)
}
func runGen(cmd *cmds.Command, args []string) {
	mod := map[string]string{
		"Name":  name,
		"Notes": notes,
		"Path":  strings.ToLower(name),
	}
	lName := strings.ToLower(name)
	os.MkdirAll("gens/model/", 0666)
	os.MkdirAll("gens/router/", 0666)
	os.MkdirAll("gens/control/", 0666)
	os.MkdirAll("gens/vue/api/", 0666)
	os.MkdirAll("gens/vue/router/", 0666)
	// 渲染model
	buf := &bytes.Buffer{}
	genTmpls.ExecuteTemplate(buf, "model", mod)
	ioutil.WriteFile("gens/model/"+lName+".go", buf.Bytes(), 0666)

	// 渲染 router
	buf.Reset()
	genTmpls.ExecuteTemplate(buf, "router", mod)
	ioutil.WriteFile("gens/router/"+lName+".go", buf.Bytes(), 0666)

	// 渲染 control
	buf.Reset()
	genTmpls.ExecuteTemplate(buf, "control", mod)
	ioutil.WriteFile("gens/control/"+lName+".go", buf.Bytes(), 0666)
	// 渲染 control
	buf.Reset()
	genTmpls.ExecuteTemplate(buf, "vue_api", mod)
	ioutil.WriteFile("gens/vue/api/"+lName+".js", buf.Bytes(), 0666)
	// 渲染 router
	buf.Reset()
	genTmpls.ExecuteTemplate(buf, "vue_router", mod)
	ioutil.WriteFile("gens/vue/router/"+lName+".js", buf.Bytes(), 0666)
}
