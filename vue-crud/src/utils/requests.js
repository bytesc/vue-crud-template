//axios 封装

import axios from "axios";
import {ElMessage} from "element-plus";


// 全局配置
const service = axios.create({
    baseURL:"/api",
    timeout:300,  //请求超时
})

// 响应拦截
service.interceptors.response.use(res=>{
    // res为获取的所有数据
    // console.log(res)
    const {code,data,msg} = res.data
    // code,data,msg是后端返回json里的三个字段（后端随便约定）
    // code是后端随便约定的状态标识码（不是http状态码）
    if (code === "200"){
        // 后端返回成功
        ElMessage.success(msg)
        return data
    }else if(code === "400"){
        // 后端返回失败
        ElMessage.error(msg)
        console.log(res.data)
    }
})

// 请求方法配置
function request(options){
    if(options.method.toLowerCase() === "get"){
        options.params = options.data
    }
    return service(options)
}

// 给request添加request.post()和.get(),从而给request传参（这一步作用是封装）
["get","post"].forEach(item=>{
    request[item] = (url,data) =>{
        return request({
            url,
            data,
            method:item
        })
    }
})

// 导出才能用
export {request}