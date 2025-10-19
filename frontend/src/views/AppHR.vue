<script setup lang="ts">
import JobBox from "../components/JobBox.vue";
import {useRouter} from 'vue-router'
import {onMounted, ref} from "vue";
import { api } from '../../api/client.js';

const router = useRouter()
let isLoading = ref(true)
const jobs = ref([])

async function findMyJobs() {
  try {
    isLoading.value = true;
    const res = await api.get("/company/jobs", {
    })

    jobs.value = res.data.jobs
    console.log("Fetched jobs:", jobs.value)
    return;
  } catch (error) {
    console.error("Error fetching jobs:", error)
  } finally {
    isLoading.value = false;
  }
}

onMounted(() => {
  findMyJobs();
})

function goToJob(id) {
  console.log('Clicked job ID:', id)
  router.push({ name: 'Applicants', params: {id}})
}

function goToCreateJob() {
  router.push({ name: 'CreateJob'})
}

</script>

<template>
  <div class="flex flex-col w-full space-y-10">
    <div class="flex flew-row bg-title w-full h-40 px-42 items-center">
      <span class="text-white text-5xl font-bold">Your Job Offer</span>
      <button @click="goToCreateJob" class="btn shadow-none bg-lighter border-0 h-12 rounded-2xl text-white text-xl font-extralight w-[150px] ml-auto">Create Job</button>
    </div>

    <div v-if="!isLoading" class="grid grid-cols-2 xl:px-42 gap-x-5 md:gap-x-20 gap-y-10 pb-20">
      <job-box v-for="job in jobs" :key="job" :job-info=job :h-r="true" @click="goToJob(job.ID)"/>
    </div>
  </div>

</template>

<style scoped>

</style>