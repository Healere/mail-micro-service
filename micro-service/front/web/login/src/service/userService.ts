import userRequest from '../utils/userRequest'

//用户注册
const register = ({name,email,password}) =>{
    return userRequest.post('auth/register',{name,email,password})
}

//用户登录
const login = ({name,email, password}) => {
    return userRequest.post('auth/login',{name,email,password})
}

// 获取用户信息
const info = () =>{
    return userRequest.get('auth/info')
}

export default{
    register,
    login,
    info,
}