<template>
  <div class="container">
    <!-- 在容器内添加加载提示 -->
    <div v-if="isLoading" class="loading-tip">
      <i class="el-icon-loading"></i> 正在加载手机号...
    </div>

    <!-- 空状态提示 -->
    <div v-else-if="!phoneNumbers.length" class="empty-tip">
      暂无可用手机号
    </div>

    <!-- 电话号码列表 -->
    <div v-for="num in phoneNumbers" :key="num"
         class="scheme5"
         @click="toggleSelection(num)"
         :data-phone="num">
      📱 {{ num }}
      <!-- 聚焦指示器 -->
      <div v-if="focused === num" class="focus-indicator"></div>

      <!-- 展开面板动画 -->
      <transition name="slide">
        <div v-if="selectedNumbers.has(num)" class="new-element">
          <!-- 短信发送按钮 -->
          <button @click="sendMessage(num)" class="copy-btn">发送短信</button>
          <!-- 获取短信按钮 -->
          <button @click="getMessage(num)" class="copy-btn">获取短信</button>
        </div>
      </transition>
    </div>

    <!-- ▋短信发送弹窗 -->
    <div v-if="showMessageDialog" class="message-dialog">
      <div class="dialog-content">
        <h3>发送短信</h3>
        <form @submit.prevent="confirmSend">
          <div class="form-group">
            <label>发送给：</label>
            <input
                type="text"
                placeholder="10010"
                required
            >
          </div>
          <div class="form-group">
            <label>短信内容：</label>
            <textarea
                v-model="messageData.content"
                placeholder="请输入短信内容..."
                required
            ></textarea>
          </div>
          <div class="button-group">
            <button type="button" @click="cancelSend">取消</button>
            <button type="submit">确认发送</button>
          </div>
        </form>
      </div>
    </div>

    <!-- ▋短信获取弹窗 -->
    <!-- 根据短信发送弹窗的样式补充短信获取弹窗 -->
    <div v-if="showGetMessageDialog" class="get-message-dialog">
      <div class="dialog-content">
        <h3>短信收件箱</h3>
        <div class="message-list">
          <table>
            <thead>
            <tr>
              <th>短信发信人</th>
              <th>短信内容</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="(msg, index) in messages" :key="index">
              <td>{{ formatPhoneNumber(msg.sender) }}</td>
              <td class="message-content">{{ msg.content }}</td>
            </tr>
            </tbody>
          </table>
        </div>
        <div class="button-group">
          <button @click="showGetMessageDialog = false">关闭窗口</button>
        </div>
      </div>
    </div>
  </div>

</template>

<script setup>
import { ref, computed } from 'vue'
import requests from "../api/request.ts";
import { onMounted } from 'vue';  // 关键导入

// ▋响应式数据
const selected = ref(null)
const focused = ref(null)
const selectedNumbers = ref(new Set())
const phoneNumbers = ref([])
const isLoading = ref(false)

// ▋短信相关状态
const showMessageDialog = ref(false)
const selectedRecipient = ref('')
const messageData = ref({
  recipient: '',
  content: ''
})


onMounted(() => {
  fetchPhoneNumbers()
})

// get phone number
const fetchPhoneNumbers = async () => {
  try {
    isLoading.value = true
    const response = await requests.post('/api/getPhones')
    phoneNumbers.value = response.data || []
  } catch (error) {
    console.error('failed:')
  } finally {
    isLoading.value = false
  }
}


// ▋切换选中状态
const toggleSelection = (num) => {
  focused.value = focused.value === num ? null : num
  selectedNumbers.value.has(num)
      ? selectedNumbers.value.delete(num)
      : selectedNumbers.value.add(num)
}

// ▋打开短信发送弹窗
const sendMessage = (num) => {
  selectedRecipient.value = num
  // messageData.value.recipient = num
  showMessageDialog.value = true
}

// ▋短信获取状态
const showGetMessageDialog = ref(false)
const messages = ref([]);


// 获取短信（模拟数据）
const getMessage = async (num) => {
  try {
    const response = await requests.post('/api/getMessages');
    messages.value = response.data;
  } catch (error) {
    console.error('请求出错:', error);
  }
  showGetMessageDialog.value = true
};


// ▋确认发送短信
const confirmSend = async () => {
  console.log('发送短信:', messageData.value)
  try {
    const response = await requests.post('/api/sendMessages', messageData.value)
    messages.value = response.value
    // 发送成功后清空短信内容
  }catch (error) {
    console.error('请求出错:', error);
  }
}

// ▋取消发送
const cancelSend = () => {
  showMessageDialog.value = false
  messageData.value.content = ''
}
</script>

<style scoped>

.container {
  display: grid;
  gap: 1rem;
  padding: 20px;
}

.scheme5 {
  position: relative;
  background: white;
  color: #1E90FF;
  border: 1px solid #1E90FF;
  box-shadow: 0 0 8px #1E90FF40;
  padding: 12px 20px;
  border-radius: 8px;
  transition: all 0.3s;
  cursor: pointer;
}

.focus-indicator {
  position: absolute;
  inset: -2px;
  border: 2px solid #1E90FF;
  border-radius: 10px;
  animation: pulse 1.5s infinite;
}

.new-element {
  background: rgba(30, 144, 255, 0.08); /* 基于主色的5%透明度背景[1](@ref) */
  border: 1px solid rgba(30, 144, 255, 0.2); /* 半透明边框 */
  color: #1E90FF; /* 继承主色文字 */
  padding: 12px;
  border-radius: 6px;
  margin-top: 8px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  backdrop-filter: blur(2px); /* 毛玻璃效果 */
  box-shadow: 0 2px 8px rgba(30, 144, 255, 0.05); /* 柔和投影 */
}

.copy-btn {
  background: #1E90FF;
  color: white;
  border: none;
  padding: 4px 8px;
  border-radius: 4px;
  cursor: pointer;
  transition: opacity 0.2s;
}

.copy-btn:hover {
  opacity: 0.8;
}

/* 动画效果 */
.slide-enter-active {
  transition: all 0.3s ease-out;
}

.slide-leave-active {
  transition: all 0.2s ease-in;
}

.slide-enter-from,
.slide-leave-to {
  transform: translateY(-10px);
  opacity: 0;
}

@keyframes pulse {
  0% { opacity: 0.5; }
  50% { opacity: 1; }
  100% { opacity: 0.5; }
}
/* ▋短信弹窗样式 */
.message-dialog {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 999;
}

.dialog-content {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  width: 400px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.form-group {
  margin-bottom: 1.5rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  color: #333;
}

input, textarea {
  width: 100%;
  padding: 0.8rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 14px;
}

textarea {
  height: 120px;
  resize: vertical;
}

.button-group {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
}

.button-group button {
  padding: 0.6rem 1.2rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: opacity 0.2s;
}

.button-group button[type="submit"] {
  background: #1E90FF;
  color: white;
}

.button-group button[type="button"] {
  background: #f0f0f0;
  color: #666;
}

.button-group button:hover {
  opacity: 0.9;
}

/* 获取短信弹窗 */
.get-message-dialog {
  /*composes: message-dialog;*/
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 999;
}

.message-list {
  max-height: 400px;
  overflow-y: auto;
  margin: 1rem 0;
}

table {
  width: 100%;
  border-collapse: collapse;
  background: #f8f9fa;
}

th, td {
  padding: 12px;
  border-bottom: 1px solid #eee;
  text-align: left;
}

th {
  background: #1E90FF;
  color: white;
  position: sticky;
  top: 0;
}

.message-content {
  max-width: 300px;
  white-space: normal;
  word-break: break-all;
}

/* 响应式适配 */
@media (max-width: 480px) {
  .dialog-content {
    width: 90%;
    padding: 1rem;
  }

  th, td {
    padding: 8px;
    font-size: 0.9em;
  }
}
</style>