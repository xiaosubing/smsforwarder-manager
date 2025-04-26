
// store/user.ts
// import { defineStore } from 'pinia'
// import { ref } from 'vue'
//
// export const userInfoStore = defineStore('user', () => {
//     const isLogin = ref(false)
//     const username = ref('')
//
//     return { isLogin, username }
// }, {
//     persist: true // 需要安装 pinia-plugin-persistedstate
// })


// store/user.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'

export const userInfoStore = defineStore('user', () => {
    const isLogin = ref(false)
    const username = ref('')
    const token = ref('')
    const tokenExpiration = ref(0) // 存储过期时间戳

    // 设置 token 和过期时间（10分钟后）
    function setToken(newToken: string) {
        token.value = newToken
        // 当前时间 + 10分钟（600000毫秒）
        tokenExpiration.value = Date.now() + 6000
        isLogin.value = true
    }

    // 检查 token 是否有效
    function isTokenValid() {
        return token.value && Date.now() < tokenExpiration.value
    }

    // 清除 token
    function clearToken() {
        token.value = ''
        tokenExpiration.value = 0
        isLogin.value = false
    }

    return {
        isLogin,
        username,
        token,
        tokenExpiration,
        setToken,
        isTokenValid,
        clearToken
    }
}, {
    persist: true
})