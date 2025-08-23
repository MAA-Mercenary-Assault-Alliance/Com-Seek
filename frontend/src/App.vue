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
  <GuestNavbar></GuestNavbar>
  <div id="all" class="bg-[#F2F6FC] text-black"> <!--Default Text set to black right here -->

    <div id="main" class="flex justify-center">
      <router-view></router-view>
    </div>

  </div>

</template>
