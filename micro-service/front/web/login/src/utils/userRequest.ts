import axios from "axios";
import storageService from "../service/storageService";
const service = axios.create({
  baseURL: import.meta.env.VITE_VUE_APP_USER_URL,
  timeout: 1000 * 5,
  headers: {
    Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}`,
  },

});

service.interceptors.request.use(function (config) {
  // Do something before request is sent
  Object.assign(config.headers, {
    Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}`
  });
  return config;
}, function (error) {
  // Do something with request error
  return Promise.reject(error);
});

export default service;
