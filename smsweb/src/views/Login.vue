<template>
  <div class="login-container">
    <h1>用户登录</h1>
    <form @submit.prevent="handleLogin">
      <div class="form-group">
        <label>用户名:</label>
        <input
            type="text"
            v-model="form.username"
            required
            placeholder="请输入用户名"
        />
      </div>

      <div class="form-group">
        <label>密码:</label>
        <input
            type="password"
            v-model="form.password"
            required
            placeholder="请输入密码"
        />
      </div>

      <div v-if="errorMessage" class="error-message">
        {{ errorMessage }}
      </div>

      <button type="submit" :disabled="loading">
        {{ loading ? '登录中...' : '登录' }}

      </button>
      <!--      <index-view></index-view>-->
    </form>
  </div>
</template>

<script setup>
import {ref, reactive, nextTick} from 'vue'
import requests from "../api/request.ts";
import {useRouter} from "vue-router";
import { userInfoStore } from '../store/user.ts'
const router = useRouter();
// const store = userInfoStore
const form = reactive({
  username: '',
  password: ''
})

const loading = ref(false)
const errorMessage = ref('')

const handleLogin = async () => {
  // 清除错误信息
  errorMessage.value = ''

  // 表单验证
  if (!form.username || !form.password) {
    errorMessage.value = '请输入用户名和密码'
    return
  }

  // loading.value = true
  const response = await requests.post('/api/login', {
        username: form.username,
        password: form.password
      }
  )
  if (response.status === 200) {
    // 登录成功处理
    console.log("登录成功")
    const store = userInfoStore()
    store.isLogin = true
    store.username = form.username
    await  nextTick()
    await router.push({name: 'index'})
  }else{
    console.log("登录失败")
  }

}

</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 2rem auto;
  padding: 2rem;
  box-shadow: 0 0 10px rgba(0,0,0,0.1);
  border-radius: 8px;
}

.form-group {
  margin-bottom: 1.5rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
}

input {
  width: 100%;
  padding: 0.8rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  box-sizing: border-box;
}

button {
  width: 100%;
  padding: 0.8rem;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
}

.error-message {
  color: #dc3545;
  margin-bottom: 1rem;
  font-size: 0.9rem;
}
</style>