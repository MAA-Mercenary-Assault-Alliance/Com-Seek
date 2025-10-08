<script setup lang="ts">
import { ref, onMounted } from "vue";

const role = ref(localStorage.getItem('role'));
onMounted(() => {
  window.addEventListener('storage', () => {
    role.value = localStorage.getItem('role');
  });
});
</script>

<template>
  <div class="flex bg-[#F2F6FC] justify-between items-center top-0 w-full min-h-18 p-4 text-black px-42 text-xl flex-wrap">
    <ul class="flex space-x-4 items-center">
      <li>
        <router-link to="/">
          <img src="../assets/logo.png" alt="logo" class="w-12"/>
        </router-link>
      </li>
      <li class="font-[righteous] text-3xl"><router-link to="/">Com-Seek</router-link></li>
      <li v-if="role=='admin'" class="flex ml-4 box-shadow text-xl bg-[#56A45C] rounded-4xl w-48 h-9 text-center items-center justify-center">Admin User</li>
    </ul>
    <ul class="flex space-x-15 items-center">
      <li><router-link to="/">Home</router-link></li>
      <li v-if="role=='company'"><router-link to="/hr-dashboard">HR Dashboard</router-link></li>
      <li v-if="role=='admin'"><router-link to="/admin">Admin Dashboard</router-link></li>
      <!-- Grouped links inside one li -->
      <li v-if="role==null" class="flex space-x-3 items-center">
        <router-link to="/login">Login</router-link>
        <img src="../assets/vertical-line.png" alt="vertical line" class="block h-12 w-15 -ml-6 -mr-4"/>
        <router-link to="/register">Register</router-link>
      </li>
      <li v-if="role==null"><router-link to="/for-company">For Company</router-link></li>
      <li v-if="role=='student' || role=='company'"><router-link to="/student-profile">Profile</router-link></li>
      <li v-if="role"><router-link to="/logout">Logout</router-link></li>
    </ul>

  </div>
</template>
