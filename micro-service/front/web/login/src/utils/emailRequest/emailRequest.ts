import axios from "axios";

const service = axios.create({
    baseURL: import.meta.env.VITE_VUE_APP_EMAIL_URL,
    timeout: 5 * 1000,
    headers:{
        "Content-Type":"application/json",
    }
})
service.interceptors.request.use(function(config){
    Object.assign(config.headers,{
        "Content-Type":"application/json",
    });
    return config;
},function(error){
    return Promise.reject(error);
})

export default service