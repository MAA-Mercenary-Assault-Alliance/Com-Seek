<script setup lang="ts">
import JobFull from "../components/JobFull.vue";
import JobBox from "../components/JobBox.vue";
import { api } from '../../api/client.js';

import {onMounted, ref, watch} from 'vue'
import JobFullEmpty from "../components/JobFullEmpty.vue";

const jobs = ref([])

const keyword = ref('')
const jobType = ref('')
const location = ref('')
const isLoading = ref(true)
const selectedJob = ref(null) // need to pass this value to JobBoard
const verified = ref(false)

watch(keyword, (newVal) => {
  console.log('keyword changed:', newVal)
})

watch(jobType, (newVal) => {
  console.log('jobType changed:', newVal)
})

watch(location, (newVal) => {
  console.log('location changed:', newVal)
})

async function findJobs() {
  try {
    isLoading.value = true;
    const params = {};
    if (keyword.value) params.keyword = keyword.value
    if (jobType.value) params.job_type = jobType.value
    if (location.value) params.location = location.value
    const res = await api.get("/job", {
      params
    })
    const job_filter = res.data.jobs.filter(job => job.Approved === true && job.Visibility === true)

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

async function getMyStatus() {
  try {
    const res = await api.get("/student")
    console.log("User status:", res.data.profile.approved)
    verified.value = res.data.profile.approved
  } catch (error) {
    console.error("Error fetching user status:", error)
  }
}

onMounted(() => {
  findJobs();
  getMyStatus();
});

console.log("SelectedJob: ", selectedJob)

</script>

<template>
  <div class="flex flex-grow flex-col w-full">
    <div
        class="flex bg-title w-full h-40"
        style="background-image: url('/images/banner.png');"
    >
      <div class="flex px-42 w-full flex-col justify-center items-center space-y-7 text-xl">
        <div class="text-white min-w-300 w-1/2">
          Search for your jobs now
        </div>
        <div class="relative flex flex-row space-x-10 min-w-300 w-1/2">
          <label class="input-box w-2/7">
            <input v-model="keyword" type="search" class="grow pl-3 placeholder-black" placeholder="Keywords" />
          </label>
          <div class="relative w-2/7">
            <img src="../assets/case.svg" class="absolute px-2 w-12 left-3 top-2 z-10" alt="case"/>
            <select v-model="jobType" class="select rounded-2xl select-lg pl-18 z-0 w-full">
              <option disabled value="">Job Type</option>
              <option value="">Any</option>
              <option value="Software & Application Development">Software & Application Development&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</option>
              <option>Data & AI</option>
              <option>Cloud & Infrastructure</option>
              <option>Cybersecurity</option>
              <option>Product & Design</option>
              <option>Testing & Quality</option>
              <option>Hardware & Electronics</option>
              <option>Management & Leadership</option>
              <option>IT Support & Operations</option>
            </select>
          </div>
          <div class="relative w-2/7 mr-10">
            <img src="../assets/location.svg" class="absolute px-2 w-12 left-3 top-2 z-10" alt="case"/>
            <label class="input-box w-2/7">
              <input v-model="location" type="search" class="grow pl-13 placeholder-black" placeholder="Location" />
            </label>
          </div>
          <button @click="findJobs" class="btn shadow-none bg-lighter border-0 h-12 rounded-2xl text-white text-xl font-extralight w-1/7">Find Now</button>
        </div>
      </div>

    </div>

    <div v-if="jobs.length==0" class="flex flex-col flex-grow -translate-y-10 w-full justify-center items-center h-60 text-2xl text-gray-500 space-y-5">
      <img src="../../src/assets/leaf2.svg" class="w-50" alt="leaf"/>
      <span>Jobs not Found</span>
    </div>

    <div v-if="!isLoading && jobs.length>0" class="flex w-full flex-row px-42 py-10 bg-background space-x-20 scroll">

      <div id="job-box-column" class="flex w-1/3 flex-col space-y-5 mr-10">
        <JobBox v-for="job in jobs" :key=job.ID :jobInfo=job :h-r="false" @click="selectedJob=job"/>
      </div>

      <div class="relative w-2/3">
        <div class="flex sticky top-10">
        <JobFull v-if="selectedJob" :job-info="selectedJob" :h-r="false" :verified="verified"/>
        <JobFullEmpty v-if="!selectedJob"/>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>

</style>