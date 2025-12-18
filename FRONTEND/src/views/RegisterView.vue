<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const username = ref('')
const password = ref('')
const confirmPassword = ref('')

const API_URL = 'http://localhost:8080/auth'

const handleRegister = async () => {
    if (!username.value || !password.value || !confirmPassword.value) {
        alert('กรุณากรอกข้อมูลให้ครบ')
        return
    }
    if (password.value !== confirmPassword.value) {
        alert('รหัสผ่านไม่ตรงกัน')
        return
    }

    try {
        await axios.post(`${API_URL}/register`, {
            username: username.value,
            password: password.value
        })

        alert('สมัครสมาชิกสำเร็จ!')
        router.push('/')
        
    } catch (error) {
        console.error(error)
        const msg = error.response?.data?.message || 'การสมัครสมาชิกผิดพลาด'
        alert(msg)
    }
}
</script>

<template>
    <div class="container">
        <div class="card">
            <div class="card-header">IT 02-2</div>
            <div class="card-body">
                <div class="form-group">
                    <label>User</label>
                    <input type="text" v-model="username" />
                </div>
                <div class="form-group">
                    <label>Password</label>
                    <input type="password" v-model="password" />
                </div>
                <div class="form-group">
                    <label>Confirm Password</label>
                    <input type="password" v-model="confirmPassword" />
                </div>
                
                <div class="actions">
                    <button @click="handleRegister">สมัครสมาชิก</button>
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
.form-group label { width: 120px; text-align: right; }
input { padding: 5px; width: 150px; border: 1px solid #000; }
button { padding: 8px 15px; background-color: #e7e6e6; border: 1px solid #000; cursor: pointer; border-radius: 5px;}
.actions { text-align: center; margin-top: 10px; }
</style>
