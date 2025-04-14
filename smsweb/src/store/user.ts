
// store/user.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const userInfoStore = defineStore('user', () => {
    const isLogin = ref(false)
    const username = ref('')

    return { isLogin, username }
}, {
    persist: true // 需要安装 pinia-plugin-persistedstate
})
