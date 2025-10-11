<script setup>
import FullNavBar from "./components/FullNavBar.vue";

</script>

<template>
  <div class="flex flex-col relative min-h-screen">
  <FullNavBar></FullNavBar>
  <div id="all" class="flex flex-col bg-[#F2F6FC] w-full flex-grow text-black"> <!--Default Text set to black right here -->

    <div id="main" class="flex justify-center">
      <router-view></router-view>
    </div>

  </div>
  </div>

</template>