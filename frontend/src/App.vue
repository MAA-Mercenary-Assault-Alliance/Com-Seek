<script setup>
import { ref, computed } from "vue"
import { useRoute } from "vue-router"
import GuestNavbar from "./components/GuestNavbar.vue"
import LandingSiteFooter from "./components/landing/SiteFooter.vue"

const message = ref("")

async function callBackend() {
  const res = await fetch("http://localhost:8080/api/hello")
  const data = await res.json()
  message.value = data.message
}

const route = useRoute()
const isBlankLayout = computed(() => route.meta?.layout === "blank")
</script>

<template>
  <!-- BLANK LAYOUT (e.g., Landing page) -->
  <div v-if="isBlankLayout" class="min-h-screen bg-base-100">
    <router-view />
  </div>

  <!-- DEFAULT LAYOUT -->
  <div v-else class="flex flex-col relative min-h-screen">
    <GuestNavbar />
    <div id="all" class="flex flex-col bg-[#F2F6FC] w-full flex-grow text-black">
      <div id="main" class="flex justify-center">
        <router-view />
      </div>
    </div>
  </div>
  <LandingSiteFooter />
</template>
