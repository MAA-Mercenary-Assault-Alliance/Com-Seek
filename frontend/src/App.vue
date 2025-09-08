<script setup>
import { ref } from "vue"
import { RouterView, useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
import GuestNavbar from "./components/GuestNavbar.vue";
const message = ref("")

async function callBackend() {
  const res = await fetch("http://localhost:8080/api/hello")
  const data = await res.json()
  message.value = data.message
}

function logout() {
  localStorage.removeItem('userToken')
  localStorage.removeItem('userData')
  router.push('/auth')
}
</script>

<template>
  <div class="flex flex-col relative min-h-screen">
  <GuestNavbar></GuestNavbar>
  <div id="all" class="flex flex-col bg-[#F2F6FC] w-full flex-grow text-black"> <!--Default Text set to black right here -->

    <div id="main" class="flex justify-center">
      <router-view></router-view>
    </div>

    <!-- Main content -->
    <main>
      <RouterView />
    </main>
  </div>
  </div>

</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  min-height: 100vh;
}

.navbar {
  background: white;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  padding: 1rem 0;
  position: sticky;
  top: 0;
  z-index: 100;
}

.nav-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 1rem;
}

.nav-brand {
  font-size: 1.5rem;
  font-weight: bold;
  color: #2c5aa0;
  text-decoration: none;
}

.nav-links {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.nav-links a {
  color: #333;
  text-decoration: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: all 0.2s ease;
}

.nav-links a:hover,
.nav-links a.router-link-active {
  background-color: #f0f0f0;
  color: #2c5aa0;
}

.logout-btn {
  background: #dc3545;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.logout-btn:hover {
  background: #c82333;
}
</style>