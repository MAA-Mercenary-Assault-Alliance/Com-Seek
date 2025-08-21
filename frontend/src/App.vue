<script setup>
import { ref } from "vue"
import AppNavbar from "./components/AppNavbar.vue";
import JobBoard from "./components/JobBoard.vue";
const message = ref("")

async function callBackend() {
  const res = await fetch("http://localhost:8080/api/hello")
  const data = await res.json()
  message.value = data.message
}
</script>

<template>
  <AppNavbar></AppNavbar>
  <div id="all" class="bg-[#F2F6FC] text-black"> <!--Default Text set to black right here -->

    <div id="main" class="flex justify-center">
      <JobBoard></JobBoard>
    </div>

  </div>

</template>
