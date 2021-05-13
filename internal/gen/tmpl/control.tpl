package sysctl
// {{.Name}}Get doc
// @Tags {{.Path}}
// @Summary 通过id获取单条{{.Notes}}
// @Param id query int true "id"
// @Success 200 {object} model.Reply{data=model.{{.Name}}} "成功数据"
// @Router /api/{{.Path}}/get [get]
func {{.Name}}Get(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, has := model.{{.Name}}Get(ipt.Id)
	if !has {
		return ctx.JSON(utils.ErrOpt("未查询到{{.Notes}}"))
	}
	return ctx.JSON(utils.Succ("succ", mod))
}

// {{.Name}}All doc
// @Tags {{.Path}}
// @Summary 获取所有{{.Notes}}
// @Success 200 {object} model.Reply{data=[]model.{{.Name}}} "成功数据"
// @Router /api/{{.Path}}/all [get]
func {{.Name}}All(ctx echo.Context) error {
	mods, err := model.{{.Name}}All()
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到{{.Notes}}", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mods))
}

// {{.Name}}Page doc
// @Tags {{.Path}}
// @Summary 获取{{.Notes}}分页
// @Param cid path int true "分类id" default(1)
// @Param pi query int true "分页数" default(1)
// @Param ps query int true "每页条数[5,30]" default(5)
// @Success 200 {object} model.Reply{data=[]model.{{.Name}}} "成功数据"
// @Router /api/{{.Path}}/page [get]
func {{.Name}}Page(ctx echo.Context) error {
	// cid, err := strconv.Atoi(ctx.Param("cid"))
	// if err != nil {
	//  return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
	// }
	ipt := &model.Page{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	if ipt.Ps > 30 || ipt.Ps < 1 {
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
// @Summary 添加{{.Notes}}
// @Param token query string true "token"
// @Param body body model.{{.Name}} true "请求数据"
// @Success 200 {object} model.Reply{data=string} "成功数据"
// @Router /adm/{{.Path}}/add [post]
func {{.Name}}Add(ctx echo.Context) error {
	ipt := &model.{{.Name}}{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	ipt.Ctime = time.Now()
	err = model.{{.Name}}Add(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// {{.Name}}Edit doc
// @Tags {{.Path}}
// @Summary 修改{{.Notes}}
// @Param token query string true "token"
// @Param body body model.{{.Name}} true "请求数据"
// @Success 200 {object} model.Reply{data=string} "成功数据"
// @Router /adm/{{.Path}}/edit [post]
func {{.Name}}Edit(ctx echo.Context) error {
	ipt := &model.{{.Name}}{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	ipt.Ctime = time.Now()
	err = model.{{.Name}}Edit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// {{.Name}}Drop doc
// @Tags {{.Path}}
// @Summary 通过id删除单条{{.Notes}}
// @Param id query int true "id"
// @Param token query string true "token"
// @Success 200 {object} model.Reply{data=string} "成功数据"
// @Router /adm/{{.Path}}/drop [post]
func {{.Name}}Drop(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	err = model.{{.Name}}Drop(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}