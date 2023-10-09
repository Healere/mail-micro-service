import logService from '../../service/logService'

const logModule = {
    namespaced: true,
    actions:{
        log(context,{email,logData}){
            return new Promise((resolve, reject)=>{
                logData = email + " logged in..."
                logService.log({email,logData})
                .then(res =>{
                    resolve(res);
                })
                .then(err => {
                    reject(err)
                })
            })
        }
    }
}

export default logModule;