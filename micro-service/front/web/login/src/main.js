
import { createApp } from 'vue'

import axios from 'axios'
import App from './App.vue'
import router from './router'
import Vuex from 'vuex'
import store from './store'
import ElemrntPLus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'



const app = createApp(App)
app.config.globalProperties.$axios = axios
app.use(router)
app.use(Vuex)
app.use(store)
app.use(ElemrntPLus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
app.mount('#app')
