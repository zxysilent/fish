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


// ----------------------------------------------复制到Layout.vue------------------------------------------------------
/*


<Submenu name="{{.Path}}" v-auth="'{{.Path}}_show'">
    <template slot="title" v-auth="'{{.Path}}_show'">
        <Icon type="ios-color-fill-outline" />
        {{.Notes}}管理
    </template>
    <MenuItem name="{{.Path}}-list" to="/{{.Path}}/list" v-auth="'{{.Path}}_show'">
    <Icon type="ios-list-box-outline" />{{.Notes}}列表</MenuItem>
    <MenuItem name="{{.Path}}-add" to="/{{.Path}}/add" v-auth="'{{.Path}}_add'">
    <Icon type="ios-add-circle-outline" />添加{{.Notes}}</MenuItem>
</Submenu>


*/