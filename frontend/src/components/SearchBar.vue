<script setup lang="ts">

</script>

<template>
  <div class="flex bg-[#153243] w-full h-40">
    <div class="flex px-42 w-full flex-col justify-center items-center space-y-7 text-xl">
      <div class="text-white min-w-300 w-1/2">
        Search for your jobs now
      </div>
      <div class="flex flex-row space-x-10 min-w-300 w-1/2">
        <label class="input-box">
          <input type="search" class="grow pl-3 placeholder-black" placeholder="Keywords" />
        </label>
        <div class="relative">
          <img src="../assets/case.svg" class="absolute px-2 w-12 left-3 top-2 z-10" alt="case"/>
          <select class="select rounded-2xl select-lg pl-18 z-0">
            <option disabled selected>Job Type</option>
            <option>Full-Stack Developer</option>
            <option>Front-End Developer</option>
            <option>Back-End Developer</option>
            <option>DevOps Engineer</option>
            <option>Add more later</option>
          </select>
        </div>
        <label class="input-box">
          <img src="../assets/location.svg" class="px-2 w-13" alt="location"/>
          <input type="text" class="grow placeholder-black" placeholder="Location" />
        </label>
        <button class="btn shadow-none bg-[#1F7AB9] border-0 h-12 rounded-2xl text-white text-xl font-extralight">Find Now</button>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>