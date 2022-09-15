package model
// {{.Name}} {{.Notes}}
type {{.Name}} struct {
	Id     int       `xorm:"INT(11) PK AUTOINCR comment('主键')"`
	String string    `xorm:"VARCHAR(255) comment('注释')"`
	Float  float64   `xorm:"DOUBLE comment('注释')"`
	Int    int       `xorm:"INT(11) DEFAULT 0 comment('注释')"`
	{{.Name}}   *{{.Name}} `xorm:"-" swaggerignore:"true"` //附加
	Updated int64    `xorm:"BIGINT" json:"updated"`              //修改时间
	Created int64    `xorm:"BIGINT" json:"created"`              //创建时间
}
// {{.Name}}Get 单条{{.Notes}}
func {{.Name}}Get(id int) (*{{.Name}}, bool) {
	mod := &{{.Name}}{}
	has, _ := db.ID(id).Get(mod)
	return mod, has
}

// {{.Name}}All 所有{{.Notes}}
func {{.Name}}All() ([]{{.Name}}, error) {
	mods := make([]{{.Name}}, 0, 8)
	err := db.Find(&mods)
	return mods, err
}

// {{.Name}}Page {{.Notes}}分页
func {{.Name}}Page(pi int, ps int, cols ...string) ([]{{.Name}}, error) {
	mods := make([]{{.Name}}, 0, ps)
	sess := db.NewSession()
	defer sess.Close()
	if len(cols) > 0 {
		sess.Cols(cols...)
	}
	err := sess.Desc("Ctime").Limit(ps, (pi-1)*ps).Find(&mods)
	return mods, err
}

// {{.Name}}Count {{.Notes}}分页总数
func {{.Name}}Count() int {
	mod := &{{.Name}}{}
	sess := db.NewSession()
	defer sess.Close()
	count, _ := sess.Count(mod)
	return int(count)
}

// {{.Name}}Add 添加{{.Notes}}
func {{.Name}}Add(mod *{{.Name}}) error {
	sess := db.NewSession()
	defer sess.Close()
	sess.Begin()
	if _, err := sess.InsertOne(mod); err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}

// {{.Name}}Edit 编辑{{.Notes}}
func {{.Name}}Edit(mod *{{.Name}}, cols ...string) error {
	sess := db.NewSession()
	defer sess.Close()
	sess.Begin()
	if _, err := sess.ID(mod.Id).Cols(cols...).Update(mod); err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}

// {{.Name}}Ids 通过id集合返回{{.Notes}}
func {{.Name}}Ids(ids []int) map[int]*{{.Name}} {
	mods := make([]{{.Name}}, 0, len(ids))
	db.In("id", ids).Find(&mods)
	mapSet := make(map[int]*{{.Name}}, len(mods))
	for idx := range mods {
		mapSet[mods[idx].Id] = &mods[idx]
	}
	return mapSet
}

// {{.Name}}Drop 删除单条{{.Notes}}
func {{.Name}}Drop(id int) error {
	sess := db.NewSession()
	defer sess.Close()
	sess.Begin()
	if _, err := sess.ID(id).Delete(&{{.Name}}{}); err != nil {
		sess.Rollback()
		return err
	}
	sess.Commit()
	return nil
}