// ----------------------------------------------复制到router_api ------------------------------------------------------

// {{.Path}} {{.Notes}}
api.GET("/{{.Path}}/get", appctl.{{.Name}}Get)
api.GET("/{{.Path}}/all", appctl.{{.Name}}All)
api.GET("/{{.Path}}/page", appctl.{{.Name}}Page)


// ----------------------------------------------复制到router_adm ------------------------------------------------------

// {{.Path}} {{.Notes}}
adm.POST("/{{.Path}}/add", appctl.{{.Name}}Add)
adm.POST("/{{.Path}}/edit", appctl.{{.Name}}Edit)
adm.POST("/{{.Path}}/drop", appctl.{{.Name}}Drop)