package applet
// {{.Name}}Get doc
// @Tags {{.Path}}
// @Summary 通过id获取单条{{.Notes}}
// @Param id query int true "id"
// @Success 200 {object} model.Reply{data=model.{{.Name}}} "返回数据"
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
// @Success 200 {object} model.Reply{data=[]model.{{.Name}}} "返回数据"
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
// @Param query query model.Page true "请求数据"
// @Success 200 {object} model.Reply{data=[]model.{{.Name}}} "返回数据"
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
	if err = ipt.Stat(); err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
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
// @Success 200 {object} model.Reply{data=string} "返回数据"
// @Router /api/{{.Path}}/add [post]
func {{.Name}}Add(ctx echo.Context) error {
	ipt := &model.{{.Name}}{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	ipt.Updated = time.Now().UnixMilli()
	ipt.Created = ipt.Updated
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
// @Success 200 {object} model.Reply{data=string} "返回数据"
// @Router /api/{{.Path}}/edit [post]
func {{.Name}}Edit(ctx echo.Context) error {
	ipt := &model.{{.Name}}{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	ipt.Updated = time.Now().UnixMilli()
	err = model.{{.Name}}Edit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// {{.Name}}Drop doc
// @Tags {{.Path}}
// @Summary 通过id删除单条{{.Notes}}
// @Param token query string true "token"
// @Param query query model.IptId true "请求数据"
// @Success 200 {object} model.Reply{data=string} "返回数据"
// @Router /api/{{.Path}}/drop [post]
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