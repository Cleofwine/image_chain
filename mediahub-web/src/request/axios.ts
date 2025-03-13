import axios,{AxiosError} from "axios"
import { ElMessage } from "element-plus";


const service = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
});

// 请求拦截器
service.interceptors.request.use(function(config){
    return config;
},function(error){
    return Promise.reject(error);
});

// 响应拦截器
service.interceptors.response.use(function(response){
    return response;
},function(error){
    const axiosErr = error as AxiosError
    // console.log(axios.response?.status)
    if (!axiosErr.response?.status) {
        // 消息提示
        ElMessage({
            showClose:true,
            message:axiosErr.message,
            type:'error',
        })
    }else if (axiosErr.response?.status == 500) {
        // 提示服务器内部错误
        ElMessage({
            showClose:true,
            message:"服务器内部错误",
            type:'error',
        })
    }else if (axiosErr.response?.status == 504) { 
        // 提示服务器网关超时
        ElMessage({
            showClose:true,
            message:"网关超时",
            type:'error',
        })
    }else if (axiosErr.response?.status == 413) { 
        // 提示请求过大，提示最大限制
        ElMessage({
            showClose:true,
            message:"仅支持上传20M以内的图片",
            type:'error',
        })
    }
    console.log(axiosErr.message)
    console.log(axiosErr.response?.status)
    return Promise.reject(error);
});

export default service