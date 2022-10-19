// {{.Path}} {{.Notes}}
authGET("/{{.Path}}/get", applet.{{.Name}}Get)// 单条{{.Notes}}
authGET("/{{.Path}}/all", applet.{{.Name}}All)//所有{{.Notes}}
authGET("/{{.Path}}/page", applet.{{.Name}}Page)// {{.Notes}}分页
authPOST("/{{.Path}}/add", applet.{{.Name}}Add)// 添加用户
authPOST("/{{.Path}}/edit", applet.{{.Name}}Edit)// 编辑{{.Notes}}
authPOST("/{{.Path}}/drop", applet.{{.Name}}Drop)// 删除{{.Notes}}

