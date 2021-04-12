// ----------------------------------------------复制到router_api ------------------------------------------------------
// {{.Path}} {{.Notes}}
api.GET("/{{.Path}}/get", sysctl.{{.Name}}Get)
api.GET("/{{.Path}}/all", sysctl.{{.Name}}All)
api.GET("/{{.Path}}/page", sysctl.{{.Name}}Page)
// ----------------------------------------------复制到router_adm ------------------------------------------------------
// {{.Path}} {{.Notes}}
adm.POST("/{{.Path}}/add", sysctl.{{.Name}}Add)
adm.POST("/{{.Path}}/edit", sysctl.{{.Name}}Edit)
adm.POST("/{{.Path}}/drop", sysctl.{{.Name}}Drop)