import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'font-awesome/css/font-awesome.min.css'
import './scss/_reset.scss'

createApp(App)
  .use(store)
  .use(router)
  .mount('#app')
