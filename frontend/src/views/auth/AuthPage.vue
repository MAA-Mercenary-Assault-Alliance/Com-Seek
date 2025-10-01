<template>
  <div class="min-h-screen bg-gradient-to-br from-sky-300 to-blue-700 flex items-center justify-center p-4">
    <div class="card bg-base-100 shadow-2xl w-full max-w-md">
      <div class="card-body">
        <!-- Logo -->
        <div class="text-center mb-6">
          <h1 class="text-3xl font-bold text-primary">Com-Seek</h1>
          <p class="text-sm text-gray-500">KU Computer Engineering Job Portal</p>
        </div>

        <!-- Login Form -->
        <div v-if="!showRegister">
          <LoginForm @switch-to-register="showRegister = true" />
        </div>

        <!-- Registration Form -->
        <div v-else>
          <RegisterForm @switch-to-login="showRegister = false" />
        </div>
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
