<script setup lang="ts">
import JobBox from "../components/JobBox.vue";
import {useRouter} from 'vue-router'
import {onMounted, ref} from "vue";
import { api } from '../../api/client.js';

const router = useRouter()
let isLoading = ref(true)
const jobs = ref([])
const verified = ref(false)

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

async function getMyStatus() {
  try {
    const res = await api.get("/company")
    console.log("User status:", res.data.profile)
    verified.value = res.data.profile.Approved
  } catch (error) {
    console.error("Error fetching user status:", error)
  }
}

onMounted(() => {
  getMyStatus();
  findMyJobs();
})

function goToJob(id) {
  console.log('Clicked job ID:', id)
  router.push({ name: 'Applicants', params: {id}})
}

function goToCreateJob() {
  router.push({ name: 'CreateJob'})
}

function alertNotVerifiedCompany() {
  alert('The company is not verified yet.')
}

</script>

<template>
  <div class="flex flex-col w-full space-y-10">
    <div class="flex flew-row bg-title w-full h-40 px-42 items-center">
      <span class="text-white text-5xl font-bold">Your Job Offer</span>
      <button v-if="verified" @click="goToCreateJob" class="btn shadow-none bg-lighter border-0 h-12 rounded-2xl text-white text-xl font-extralight w-[150px] ml-auto">Create Job</button>
      <button v-else @click="alertNotVerifiedCompany" class="btn shadow-none bg-gray-300 border-0 h-12 rounded-2xl text-white text-xl font-extralight w-[150px] ml-auto">Create Job</button>
    </div>

    <div v-if="!isLoading && jobs.length > 0" class="grid grid-cols-2 xl:px-42 gap-x-5 md:gap-x-20 gap-y-5 pb-20">
      <job-box v-for="job in jobs" :key="job" :job-info=job :h-r="true" @click="goToJob(job.ID)"/>
    </div>

    <div v-else class="flex flex-col flex-grow -translate-y-10 w-full justify-center items-center h-60 text-2xl text-gray-500 space-y-5">
      <img src="../../src/assets/leaf2.svg" class="w-50" alt="leaf"/>
      <span>You haven't posted any job offer yet</span>
    </div>
  </div>

</template>

<style scoped>

</style>