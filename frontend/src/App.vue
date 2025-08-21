<script setup>
import { ref } from "vue"
import AppNavbar from "./components/AppNavbar.vue";
const message = ref("")

async function callBackend() {
  const res = await fetch("http://localhost:8080/api/hello")
  const data = await res.json()
  message.value = data.message
}
</script>

<template>
  <AppNavbar></AppNavbar>
  <div>
    <h1>Vue + Go Boilerplate</h1>
    <button @click="callBackend">Call Go Backend</button>
    <p>{{ message }}</p>
  </div>
</template>
