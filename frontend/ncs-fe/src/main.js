import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import {store} from './store/store.js'

const app = createApp(App)
app.use(router)
console.log(store)
app.use(store)


app.mount('#app')