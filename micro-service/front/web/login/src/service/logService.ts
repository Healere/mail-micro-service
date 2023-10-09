import logRequestequest from '../utils/logRequest'

//发送日志文件
const log = ({email,logData}) =>{
    return logRequestequest.post('log',{email,logData})
}

export default{
    log,
}