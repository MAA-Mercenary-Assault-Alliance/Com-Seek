<script setup lang="ts">
import JobBox from "../components/JobBox.vue";
import {TE_Info} from "../components/temp_template";
import {useRouter} from 'vue-router'

const router = useRouter()
const job_list = [1,2,3]

function goToJob(id) {
  console.log('Clicked job ID:', id)
  router.push({ name: 'Applicants', params: {id}})
}

</script>

<template>
  <div class="flex flex-col w-full space-y-10">
    <div class="flex bg-title w-full h-40 px-42 items-center">
      <span class="text-white text-5xl font-bold">Your Job Offer</span>
    </div>

    <div class="grid grid-cols-2 px-42 gap-x-20 gap-y-10">
      <job-box v-for="job in job_list" :key="job" :job-info=TE_Info :HR="true" @click="goToJob(TE_Info.id)"/>
    </div>
  </div>

</template>

<style scoped>

</style>