import fetch from "./fetch";
// 通过id获取单条{{.Notes}}
export const api{{.Name}}Get = (data) => {
    return fetch.request({
		url: "/api/{{.Path}}/get",
		method: "get",
		params: data,
	});
};
// 获取所有{{.Notes}}
export const api{{.Name}}All = (data) => {
    return fetch.request({
		url: "/api/{{.Path}}/all",
		method: "get",
		params: data,
	});
};
// 获取{{.Notes}}分页
export const adm{{.Name}}Page = (data) => {
    return fetch.request({
		url: "/api/{{.Path}}/page",
		method: "get",
		params: data,
	});
};
// 添加{{.Notes}}
export const adm{{.Name}}Add = (data) => {
    return fetch.request({
		url: "/adm/{{.Path}}/add",
		method: "post",
		data: data,
	});
};
// 修改{{.Notes}}
export const adm{{.Name}}Edit = (data) => {
	return fetch.request({
		url: "/adm/{{.Path}}/edit",
		method: "post",
		data: data,
	});
};
// 通过id删除单条{{.Notes}}
export const adm{{.Name}}Drop = (data) => {
	return fetch.request({
		url: "/adm/{{.Path}}/drop",
		method: "post",
		data: data,
	});
};
