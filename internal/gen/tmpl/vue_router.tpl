	// ----------------------------------------------复制到router.js ------------------------------------------------------
    {
		path: "/{{.Path}}",
		name: "{{.Path}}",
		meta: { title: "{{.Notes}}管理" },
		component: Layout,
		children: [
            {
                path: "/{{.Path}}/index",
                name: "{{.Path}}-index",
                meta: { title: "{{.Notes}}管理", auth: "{{.Path}}_view" },
                component: () => import("@/views/{{.Path}}/index.vue"),
            },
            {
                path: "/{{.Path}}/add",
                name: "{{.Path}}-add",
                meta: { title: "添加{{.Notes}}", auth: "{{.Path}}_add", hidden: true, active: "{{.Path}}-index" },
                component: () => import("@/views/{{.Path}}/add.vue"),
            },
            {
                path: "/{{.Path}}/edit/:id(\\d+)",
                name: "{{.Path}}-edit",
                meta: { title: "修改{{.Notes}}", auth: "{{.Path}}_edit", hidden: true, active: "{{.Path}}-index" },
                component: () => import("@/views/{{.Path}}/edit.vue"),
            },
            {
                path: "/{{.Path}}/detail/:id(\\d+)",
                name: "{{.Path}}-detail",
                meta: { title: "查看{{.Notes}}", auth: "{{.Path}}_view", hidden: true, active: "{{.Path}}-index" },
                component: () => import("@/views/{{.Path}}/detail.vue"),
            },
		]
	},


	// ----------------------------------------------复制到db ------------------------------------------------------

INSERT INTO `app_grant` (`guid`, `name`, `group`, `path`, `method`, `sort`, `inner`, `updated`, `created`) VALUES ('{{.Path}}_view', '显示', '{{.Notes}}', '', '', '1000', '1', REPLACE(unix_timestamp(current_timestamp(3)),'.',''), REPLACE(unix_timestamp(current_timestamp(3)),'.',''));
INSERT INTO `app_grant` (`guid`, `name`, `group`, `path`, `method`, `sort`, `inner`, `updated`, `created`) VALUES ('{{.Path}}_add', '新增', '{{.Notes}}', '', '', '1000', '1', REPLACE(unix_timestamp(current_timestamp(3)),'.',''), REPLACE(unix_timestamp(current_timestamp(3)),'.',''));
INSERT INTO `app_grant` (`guid`, `name`, `group`, `path`, `method`, `sort`, `inner`, `updated`, `created`) VALUES ('{{.Path}}_edit', '编辑', '{{.Notes}}', '', '', '1000', '1', REPLACE(unix_timestamp(current_timestamp(3)),'.',''), '1664519152014');
INSERT INTO `app_grant` (`guid`, `name`, `group`, `path`, `method`, `sort`, `inner`, `updated`, `created`) VALUES ('{{.Path}}_drop', '删除', '{{.Notes}}', '', '', '1000', '1', REPLACE(unix_timestamp(current_timestamp(3)),'.',''), REPLACE(unix_timestamp(current_timestamp(3)),'.',''));
