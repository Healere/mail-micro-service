import emailService from "../../../service/email/emailservice"
const emailModule = {
    namespaced: true,
    state:{

    },
    mutations:{

    },
    actions:{
        send(context,{fromAddress,fromName, to , subject, text}){
            return new Promise((resolve,reject)=>{
                emailService.send({fromAddress,fromName, to, subject, text})
                .then(res =>{
                    resolve(res)
                })
                .catch(err => {
                    reject(err)
                })
            })
        }
    }
}

export default emailModule
