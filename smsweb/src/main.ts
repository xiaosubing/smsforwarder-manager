import { createApp } from 'vue'

import App from './App.vue'
import router from  './router/router.ts'
import pinia from './store'



createApp(App).use(router).use(pinia).mount("#app")

