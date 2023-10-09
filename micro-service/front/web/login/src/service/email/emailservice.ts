import emailRequest from '../../utils/emailRequest/emailRequest'
// 发送邮件
const send = ({fromAddress,fromName, to, subject, text}) => {
    return emailRequest.post('/email', {fromAddress, fromName, to, subject, text})
}

export default{
    send
}
