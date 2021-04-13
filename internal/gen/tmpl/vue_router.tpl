	// ----------------------------------------------复制到router.js ------------------------------------------------------
    {
		path: "/{{.Path}}",
		name: "{{.Path}}",
		meta: { title: "文章管理" },
		component: Layout,
		children: [
			{
				path: "list",
				name: "{{.Path}}-list",
				meta: { title: "{{.Notes}}列表" },
				component: () => import("@/views/{{.Path}}/list.vue")
			},
			{
				path: "add",
				name: "{{.Path}}-add",
				meta: { title: "添加{{.Notes}}" },
				component: () => import("@/views/{{.Path}}/add.vue")
			},
			{
				path: "edit/:id(\\d+)",
				name: "{{.Path}}-edit",
				meta: { title: "修改{{.Notes}}" },
				component: () => import("@/views/{{.Path}}/edit.vue")
			}
		]
	},