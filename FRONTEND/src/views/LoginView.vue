<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { jwtDecode } from 'jwt-decode'

const router = useRouter()
const username = ref('')
const password = ref('')

const API_URL = 'http://localhost:8080/auth'

const handleLogin = async () => {
  if (!username.value || !password.value) {
    alert('กรุณากรอกข้อมูลให้ครบถ้วน')
    return
  }

  try {
    const response = await axios.post(`${API_URL}/login`, {
      username: username.value,
      password: password.value
    })

    const token = response.data.result?.token 

    if (token) {
      localStorage.setItem('token', token)

      const decoded = jwtDecode(token)
      console.log('Decoded Token:', decoded)

      const displayName = decoded.username || username.value
      localStorage.setItem('user', displayName)

      router.push('/welcome')
    } else {
      throw new Error('ไม่ได้รับ Token จากระบบ')
    }

  } catch (error) {
    console.error(error)
    
    if (error.response) {
       alert('Username หรือ Password ผิด')
    } else {
       alert('ไม่สามารถเชื่อมต่อกับเซิร์ฟเวอร์ได้')
    }
  }
}
</script>

<template>
  <div class="container">
    <div class="card">
      <div class="card-header">IT 02-1</div>
      <div class="card-body">
        <div class="form-group">
          <label>User</label>
          <input type="text" v-model="username" />
        </div>
        <div class="form-group">
          <label>Password</label>
          <input type="password" v-model="password" />
        </div>
        
        <div class="actions">
          <button @click="handleLogin">ลงชื่อเข้าใช้งาน</button>
          <br>
          <RouterLink to="/register" class="link">สมัครสมาชิก</RouterLink>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container { display: flex; justify-content: center; align-items: center; height: 100vh; }
.card { border: 2px solid #000; width: 500px; height: 350px; display: flex; flex-direction: column; }
.card-header { background-color: #00b050; color: black; padding: 10px; text-align: center; border-bottom: 2px solid #000; font-weight: bold; }
.card-body { padding: 40px; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 20px; height: 100%; }
.form-group { display: flex; width: 100%; justify-content: center; align-items: center; gap: 20px; }
.form-group label { width: 80px; text-align: right; }
input { padding: 5px; width: 150px; border: 1px solid #000; }
button { padding: 8px 15px; background-color: #e7e6e6; border: 1px solid #000; cursor: pointer; border-radius: 5px;}
.link { color: #007bff; text-decoration: underline; margin-top: 10px; cursor: pointer; }
.actions { text-align: center; display: flex; flex-direction: column; align-items: center; gap: 10px;}
</style>