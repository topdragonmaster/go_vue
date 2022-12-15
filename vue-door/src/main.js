import { createApp } from 'vue'
import App from './App.vue'
import './style.css';
import { BootstrapIconsPlugin } from 'bootstrap-icons-vue';

import store from './vuex/store'
import router from './router'
import  { sync }  from 'vuex-router-sync'

sync(store, router)

const app = createApp(App)
app.use(router)
app.use(store)
app.use(BootstrapIconsPlugin);
app.mount('#app')
