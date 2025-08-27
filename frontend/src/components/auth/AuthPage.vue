<template>
  <div class="auth-page">
    <div class="auth-container">
      <div class="logo">
        <h1>Com-Seek</h1>
        <p>KU Computer Engineering Job Portal</p>
      </div>

      <!-- Login Form -->
      <div v-if="!showRegister" class="form-container">
        <LoginForm @switch-to-register="showRegister = true" />
      </div>

      <!-- Registration Form -->
      <div v-else class="form-container">
        <RegisterForm @switch-to-login="showRegister = false" />
      </div>
    </div>
  </div>
</template>

<script>
import LoginForm from './LoginForm.vue'
import RegisterForm from './RegisterForm.vue'

export default {
  name: 'AuthPage',
  components: {
    LoginForm,
    RegisterForm
  },
  data() {
    return {
      showRegister: false
    }
  }
}
</script>

<style scoped>
.auth-page {
  background: linear-gradient(135deg, #87CEEB 0%, #4682B4 100%);
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.auth-container {
  background: white;
  border-radius: 12px;
  padding: 40px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  width: 100%;
  max-width: 400px;
  position: relative;
}

.logo {
  text-align: center;
  margin-bottom: 30px;
}

.logo h1 {
  color: #2c5aa0;
  font-size: 28px;
  font-weight: bold;
  margin-bottom: 5px;
}

.logo p {
  color: #666;
  font-size: 14px;
}

@media (max-width: 480px) {
  .auth-container {
    padding: 30px 20px;
  }
  
  .logo h1 {
    font-size: 24px;
  }
}
</style>