// {{.Name}}Get doc
// @Tags {{.Path}}
// @Summary 通过id获取单条{{.Notes}}信息
// @Param id path int true "pk id" default(1)
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
// @Param cid path int true "分类id" default(1)
// @Param pi query int true "分页数" default(1)
// @Param ps query int true "每页条数[5,20]" default(5)
// @Router /api/{{.Path}}/page/{cid} [get]
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
// @Param token query string true "token"
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
// @Summary 修改{{.Notes}}信息
// @Param token query string true "token"
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
// @Summary 通过id删除单条{{.Notes}}信息
// @Param id path int true "pk id" default(1)
// @Param token query string true "token"
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