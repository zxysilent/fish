import fetch, { Reply } from '@/utils/fetch';
// 通过id获取单条{{.Notes}}
export const api{{.Name}}Get = (params): Promise<Reply> => {
    return fetch.request({
		url: "/api/{{.Path}}/get",
		method: "get",
		params: params,
	});
};
// 获取所有{{.Notes}}
export const api{{.Name}}All = (params): Promise<Reply> => {
    return fetch.request({
		url: "/api/{{.Path}}/all",
		method: "get",
		params: params,
	});
};
// 获取{{.Notes}}分页
export const api{{.Name}}Page = (data) : Promise<Reply> => {
    return fetch.request({
		url: "/api/{{.Path}}/page",
		method: "get",
		params: data,
	});
};
// 添加{{.Notes}}
export const api{{.Name}}Add = (data) : Promise<Reply> => {
    return fetch.request({
		url: "/api/{{.Path}}/add",
		method: "post",
		data: data,
	});
};
// 修改{{.Notes}}
export const api{{.Name}}Edit = (data): Promise<Reply> =>{
	return fetch.request({
		url: "/api/{{.Path}}/edit",
		method: "post",
		data: data,
	});
};
// 通过id删除单条{{.Notes}}
export const api{{.Name}}Drop = (data) : Promise<Reply> => {
	return fetch.request({
		url: "/api/{{.Path}}/drop",
		method: "post",
		data: data,
	});
};
