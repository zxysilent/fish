// +build gen

package gen

import (
	"os"
	"strings"
	"text/template"

	"github.com/zxysilent/fish/internal/cmds"
)

var CmdGen = &cmds.Command{
	UsageLine: "gen [-t=m/c] [-n=name] [-r=notes]",
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
)

func init() {
	CmdGen.Flag.StringVar(&class, "t", "m", "code type c=control,m=model")
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
	switch class {
	case "c":
		renderControl(mod)
	case "m":
		renderModel(mod)
	}
}
func renderControl(data interface{}) {
	genTmpls.ExecuteTemplate(os.Stdout, "control", data)
}
func renderModel(data interface{}) {
	genTmpls.ExecuteTemplate(os.Stdout, "model", data)
}

var modelTmpl = `
// {{.Name}}Get 单条{{.Notes}}信息
func {{.Name}}Get(id int) (*{{.Name}}, bool) {
	mod := &{{.Name}}{}
	has, _ := Db.ID(id).Get(mod)
	return mod, has
}

// {{.Name}}All 所有{{.Notes}}信息
func {{.Name}}All() ([]{{.Name}}, error) {
	mods := make([]{{.Name}}, 0, 8)
	err := Db.Find(&mods)
	return mods, err
}

// {{.Name}}Page {{.Notes}}分页信息
func {{.Name}}Page(pi int, ps int, cols ...string) ([]{{.Name}}, error) {
	mods := make([]{{.Name}}, 0, ps)
	sess := Db.NewSession()
	defer sess.Close()
	if len(cols) > 0 {
		sess.Cols(cols...)
	}
	err := sess.Desc("Utime").Limit(ps, (pi-1)*ps).Find(&mods)
	return mods, err
}

// {{.Name}}Count {{.Notes}}分页信息总数
func {{.Name}}Count() int {
	mod := &{{.Name}}{}
	sess := Db.NewSession()
	defer sess.Close()
	count, _ := sess.Count(mod)
	return count
}

// {{.Name}}Add 添加{{.Notes}}信息
func {{.Name}}Add(mod *{{.Name}}) error {
	sess := Db.NewSession()
	defer sess.Close()
	sess.Begin()
	if _, err := sess.InsertOne(mod); err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}

// {{.Name}}Edit 编辑{{.Notes}}信息
func {{.Name}}Edit(mod *{{.Name}}, cols ...string) error {
	sess := Db.NewSession()
	defer sess.Close()
	sess.Begin()
	if _, err := sess.ID(mod.Id).Cols(cols...).Update(mod); err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}

// {{.Name}}Ids 返回{{.Notes}}信息-ids
func {{.Name}}Ids(ids []int) map[int]*{{.Name}} {
	mods := make([]{{.Name}}, 0, len(ids))
	Db.In("id", ids).Find(&mods)
	if len(mods) > 0 {
		mapMods := make(map[int]*{{.Name}}, len(mods))
		for idx := range mods {
			mapMods[mods[idx].Id] = &mods[idx]
		}
		return mapMods
	}
	return nil
}

// {{.Name}}Drop 删除单条{{.Notes}}信息
func {{.Name}}Drop(id int) error {
	sess := Db.NewSession()
	defer sess.Close()
	sess.Begin()
	if _, err := sess.ID(id).Delete(&{{.Name}}{}); err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}
`

var controlTmpl = `
// {{.Name}}Get doc
// @Tags {{.Path}}
// @Summary 通过id获取单条{{.Notes}}信息
// @Param id path int64 true "pk id" default(1)
// @Success 200 {object} utils.Reply "成功数据"
// @Router /api/{{.Path}}/get/{id} [get]
func {{.Name}}Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
	}
	mod, has := model.{{.Name}}Get(id)
	if !has {
		return ctx.JSON(utils.ErrOpt("未查询到{{.Notes}}信息"))
	}
	return ctx.JSON(utils.Succ("succ", mod))
}

// {{.Name}}All doc
// @Tags {{.Path}}
// @Summary 获取所有{{.Notes}}信息
// @Success 200 {object} utils.Reply "成功数据"
// @Router /api/{{.Path}}/all [get]
func {{.Name}}All(ctx echo.Context) error {
	mods, err := model.{{.Name}}All()
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到{{.Notes}}信息", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mods))
}

// {{.Name}}Page doc
// @Tags {{.Path}}
// @Summary 获取{{.Notes}}分页信息
// @Param cid path int64 true "分类id" default(1)
// @Param pi query int true "分页数" default(1)
// @Param ps query int true "每页条数[5,20]" default(5)
// @Success 200 {object} utils.Reply "成功数据"
// @Router /api/{{.Path}}/page/{cid} [get]
func {{.Name}}Page(ctx echo.Context) error {
	// cid, err := strconv.Atoi(ctx.Param("cid"))
	// if err != nil {
	//      ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
	//      return
	// }
	ipt := &model.Page{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	if ipt.Ps > 20 || ipt.Ps < 5 {
		return ctx.JSON(utils.ErrIpt("分页大小输入错误", ipt.Ps))
	}
	count := model.{{.Name}}Count()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", " count < 1"))
	}
	mods, err := model.{{.Name}}Page(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("查询数据错误", err.Error()))
	}
	if len(mods) < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", "len(mods) < 1"))
	}
	return ctx.JSON(utils.Page("succ", mods, int(count)))
}

// {{.Name}}Add doc
// @Tags {{.Path}}
// @Summary 添加{{.Notes}}信息
// @Param token query string true "jwt" default(jwt-token)
// @Param body body model.{{.Name}} true "Req"
// @Success 200 {object} utils.Reply "成功数据"
// @Router /adm/{{.Path}}/add [post]
func {{.Name}}Add(ctx echo.Context) error {
	ipt := &model.{{.Name}}{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	ipt.Utime = time.Now()
	err = model.{{.Name}}Add(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// {{.Name}}Edit doc
// @Tags {{.Path}}
// @Summary 修改{{.Notes}}信息
// @Param token query string true "jwt" default(jwt-token)
// @Param body body model.{{.Name}} true "Req"
// @Success 200 {object} utils.Reply "成功数据"
// @Router /adm/{{.Path}}/edit [post]
func {{.Name}}Edit(ctx echo.Context) error {
	ipt := &model.{{.Name}}{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	ipt.Utime = time.Now()
	err = model.{{.Name}}Edit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// {{.Name}}Drop doc
// @Tags {{.Path}}
// @Summary 通过id删除单条{{.Notes}}信息
// @Param id path int64 true "pk id" default(1)
// @Param token query string true "jwt" default(jwt-token)
// @Success 200 {object} utils.Reply "成功数据"
// @Router /adm/{{.Path}}/drop/{id} [get]
func {{.Name}}Drop(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
	}
	err = model.{{.Name}}Drop(id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

`
