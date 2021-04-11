package model

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
	return int(count)
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

// {{.Name}}Ids 通过id集合返回{{.Notes}}信息
func {{.Name}}Ids(ids []int) map[int]*{{.Name}} {
	mods := make([]{{.Name}}, 0, len(ids))
	Db.In("id", ids).Find(&mods)
	mapSet := make(map[int]*{{.Name}}, len(mods))
	for idx := range mods {
		mapSet[mods[idx].Id] = &mods[idx]
	}
	return mapSet
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