// store/auth.ts
import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: null as string | null,
        tokenExpiry: null as number | null
    }),
    actions: {
        setToken(token: string) {
            this.token = token
            // 设置10分钟后过期 (10分钟 × 60秒 × 1000毫秒)
            this.tokenExpiry = Date.now() + 10 * 60 * 1000
        },
        getToken() {
            if (this.tokenExpiry && Date.now() > this.tokenExpiry) {
                this.token = null
                this.tokenExpiry = null
                return null
            }
            return this.token
        }
    },
    persist: true // 如果需要持久化存储
})

