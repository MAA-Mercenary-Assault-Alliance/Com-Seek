<script setup lang="ts">
import JobFull from "../components/JobFull.vue";
import JobBox from "../components/JobBox.vue";
import { api } from '../../api/client.js';

import {onMounted, ref, watch} from 'vue'
import JobFullEmpty from "../components/JobFullEmpty.vue";
import {useRoute} from "vue-router";

const route = useRoute();
const jobs = ref([])

const isLoading = ref(true)
const selectedJob = ref(null) // need to pass this value to JobBoard
const companyID = route.params.id

async function findCompanyJobs() {
  try {
    isLoading.value = true;
    const res = await api.get("/admin/jobs")
    const job_filter = res.data.jobs.filter(job => job.Approved === false && job.CompanyID == companyID)

    jobs.value = job_filter
    console.log("Fetched jobs:", jobs.value)
    selectedJob.value = jobs.value[0] || null; // Select the first job by default or null if no jobs
    return;
  } catch (error) {
    console.error("Error fetching jobs:", error)
  } finally {
    isLoading.value = false;
  }
}

onMounted(() => {
  findCompanyJobs();
});

console.log("SelectedJob: ", selectedJob)

</script>

<template>
  <div class="flex flex-col w-full">
    <div class="flex bg-title w-full h-40 px-42 items-center">
      <span v-if="!isLoading" class="text-white text-5xl font-bold">{{ jobs[0].Company.Name }} pending jobs</span>
    </div>
    <div v-if="jobs.length==0" class="flex w-full justify-center items-center h-60 text-2xl text-gray-500">
      Jobs not Found
    </div>

    <div v-if="!isLoading && jobs.length>0" class="flex w-full flex-row px-42 py-10 bg-[#f2f6fc] space-x-20 scroll">

      <div id="job-box-column" class="flex w-1/3 flex-col space-y-10 mr-10">
        <JobBox v-for="job in jobs" :key=job.ID :jobInfo=job :h-r="false" @click="selectedJob=job"/>
      </div>

      <div class="relative w-2/3">
        <div class="flex sticky top-10">
          <JobFull v-if="selectedJob" :job-info="selectedJob"/>
          <JobFullEmpty v-if="!selectedJob"/>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>

</style>