<script setup>
import { ref } from "vue"
import GuestNavbar from "./components/GuestNavbar.vue";
const message = ref("")

async function callBackend() {
  const res = await fetch("http://localhost:8080/api/hello")
  const data = await res.json()
  message.value = data.message
}
</script>

<template>
  <div class="flex flex-col relative min-h-screen">
  <GuestNavbar></GuestNavbar>
  <div id="all" class="flex flex-col bg-[#F2F6FC] w-full flex-grow text-black"> <!--Default Text set to black right here -->

    <div id="main" class="flex justify-center">
      <router-view></router-view>
    </div>

  </div>
  </div>

</template>