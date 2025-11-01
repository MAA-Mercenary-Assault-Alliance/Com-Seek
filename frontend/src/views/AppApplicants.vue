<script setup lang="ts">
import JobFull from "../components/JobFull.vue";
import ApplicantsColumn from "../components/ApplicantsColumn.vue";
import { api } from '../../api/client.js';

import { useRoute } from 'vue-router'
import {onMounted, ref} from "vue";
const route = useRoute()
const jobId = route.params.id
const isLoading = ref(true)
const job = ref(null)
const applicants = ref(null)

async function findThisJob() {
  try {
    isLoading.value = true;
    const res = await api.get(`/job/${jobId}`)

    job.value = res.data.job
    applicants.value = res.data.applicants
    console.log("Fetched this job:", job.value)
  } catch (error) {
    console.error("Error fetching jobs:", error)
  } finally {
    isLoading.value = false;
  }
}

onMounted(() => {
  findThisJob();
})

</script>

<template>
  <div class="flex flex-col w-full space-y-10">
    <div class="flex bg-title w-full h-40 px-42 items-center">
      <span class="text-white text-5xl font-bold">Applicants</span>
    </div>
    <div v-if="!isLoading && job" class="flex flex-row w-full xl:px-42 bg-[#f2f6fc] space-x-5 xl:space-x-20 pb-20">
      <JobFull :job-info="job" />
      <ApplicantsColumn v-if="applicants" :applicants="applicants" @refresh="findThisJob"/>
      <div v-else class="flex flex-col flex-grow w-1/3 justify-center items-center h-60 text-2xl mt-20 text-gray-500 space-y-5">
        <img src="../../src/assets/leaf2.svg" class="w-50" alt="leaf"/>
        <span>No applicants yet</span>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>