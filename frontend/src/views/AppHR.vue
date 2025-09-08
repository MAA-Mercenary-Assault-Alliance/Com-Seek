<script setup lang="ts">
import JobBox from "../components/JobBox.vue";
import {TE_Info} from "../components/temp_template";

const job_list = [1,2,3]
</script>

<template>
  <div class="flex flex-col w-full space-y-10">
    <div class="flex bg-[#153243] w-full h-40 px-42 items-center">
      <span class="text-white text-5xl font-bold">Your Job Offer</span>
    </div>

    <div class="grid grid-cols-2 px-42 gap-x-20 gap-y-10">
      <job-box v-for="job in job_list" :key="job" :job-info=TE_Info :HR="true" />
    </div>
  </div>

</template>

<style scoped>

</style>