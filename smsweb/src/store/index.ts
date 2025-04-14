import { createPinia } from "pinia";

// import {userKey } from './key.ts'

import piniaPluginExtend from  'pinia-plugin-persistedstate'


const pinia = createPinia()
pinia.use(piniaPluginExtend)


export default pinia
// export default useUserStore