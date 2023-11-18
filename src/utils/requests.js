//axios 封装

import axios from "axios";
import {ElMessage} from "element-plus";

import {store} from "./store.js";

// 全局配置
const service = axios.create({
    baseURL:"/api",
    timeout:10000,  //请求超时
})


// 响应拦截
service.interceptors.response.use(res=>{
        // res为获取的所有数据
        // console.log(res)
        const {code,data,msg} = res.data
        // code,data,msg是后端返回json里的三个字段（后端随便约定）
        // code是后端随便约定的状态标识码（不是http状态码）
        if (code === "200" || code === "201" ){
            // 后端返回成功
            if(code === "200"){
                ElMessage.success(msg)
            }
            const newToken = res.headers['new_token'];
            if(newToken!==localStorage.getItem('token') && typeof newToken !== 'undefined'){
                // ElMessage.success("token续签")
                // console.log(res.data.data.token)
                // 将新的token存储到浏览器的localStorage中
                localStorage.setItem('token', newToken)
            }
            store.commit('setName', localStorage.getItem("name"))
            return data
        }else if(code === "400"){
            // 后端返回失败
            ElMessage.error(msg)
            console.log(res.data)
            return data
        }else if(code === "233"){
            //登陆成功，签发了token
            ElMessage.success(msg)
            const newToken = res.headers['new_token'];
            // console.log(newToken)
            const newLongToken = res.headers['new_long_token'];
            // console.log(res.data.data.token)
            // 将新的token存储到浏览器的localStorage中
            localStorage.setItem('token', newToken)
            localStorage.setItem('long_token', newLongToken)
            localStorage.setItem('name', res.headers['name'])
            //name 存到vuex
            store.commit('setName', res.headers['name'])  // 设置 name
            window.location.href = '#/'
            return data
        } else if(code === "234") {
            //注册成功
            ElMessage.success(msg)
            window.location.href = '#/user/login'
            return data
        }else if(code === "235"){
            //登录注销成功
            localStorage.setItem('token', "")
            localStorage.setItem('long_token', "")
            localStorage.setItem('name', "")
            store.commit('setName', "")
            ElMessage.success(msg)
            window.location.href = '#/user/login'
            return data
        } else if(code === "444"){
            //无效登录状态
            ElMessage.error(msg)
            console.log(res.data)
            window.location.href = '#/user/login'
            return data
        }
    },
    error => {
        ElMessage.error("网络连接超时")
        return Promise.reject(error);
    }
)

// 请求方法配置
function request(options){
    if(options.method.toLowerCase() === "get"){
        options.params = options.data
    }
    // 从localStorage中获取token
    const token = localStorage.getItem('token')
    const longToken = localStorage.getItem('long_token')
    const name = localStorage.getItem('name')
    // 将token添加到请求头中
    options.headers = {...options.headers, 'token': token,'name':name,"long_token":longToken}
    return service({
        ...options,
        // headers: options.headers,  // 在这里设置headers
    })
}

// 给request添加request.post()和.get(),从而给request传参（这一步作用是封装）
["get","post"].forEach(item=>{
    request[item] = (url,data,headers) =>{
        return request({
            url,
            data,
            headers,
            method:item
        })
    }
})

// 导出才能用
export {request}