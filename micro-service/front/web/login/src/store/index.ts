
import { createStore } from 'vuex'
import logModule from './module/log';
import userModule from './module/user';
import emailModule from './module/email/email';

const store = createStore({
    // eslint-disable-next-line no-undef
    strict: import.meta.env.NODE_ENV !== 'production',
    state(){
        return{
           
        }
    },
    mutations:{
        
    },
    modules:{
        userModule: userModule,
        logModule: logModule,
        emailModule: emailModule,
    },
})

export default store;
