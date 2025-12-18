<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const currentUser = ref('')

onMounted(() => {
  const user = localStorage.getItem('user')
  
  const token = localStorage.getItem('token')

  if (user && token) {
    currentUser.value = user
  } else {
    router.push('/')
  }
})

const handleLogout = () => {
  localStorage.removeItem('user')
  localStorage.removeItem('token')
  router.push('/')
}
</script>

<template>
  <div class="container">
    <div class="card">
      <div class="card-header">IT 02-3</div>
      <div class="card-body">
        <h2>Welcome User : {{ currentUser }}</h2>
        <br>
        <button @click="handleLogout">Logout</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container { display: flex; justify-content: center; align-items: center; height: 100vh; }
.card { border: 2px solid #000; width: 500px; height: 350px; display: flex; flex-direction: column; }
.card-header { background-color: #00b050; color: black; padding: 10px; text-align: center; border-bottom: 2px solid #000; font-weight: bold; }
.card-body { padding: 40px; display: flex; flex-direction: column; align-items: center; justify-content: center; height: 100%; }
button { margin-top: 20px; padding: 5px 10px; cursor: pointer; }
</style>