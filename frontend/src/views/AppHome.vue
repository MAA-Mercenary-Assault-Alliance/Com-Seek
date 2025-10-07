<script setup lang="ts">
import JobFull from "../components/JobFull.vue";
import JobBox from "../components/JobBox.vue";
import { api } from '../../../api/client';

import { ref, watch } from 'vue'

const jobs = ref([])

const keyword = ref('')
const jobType = ref('')
const location = ref('')

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
    const res = await api.get("/job", {
      params: {
        keyword: keyword,
        jobType: jobType,
        location: location
      }
    })

    jobs.value = res.data
    console.log("Fetched jobs:", jobs.value)
    return;
  } catch (error) {
    console.error("Error fetching jobs:", error)
  }
}

const selectedJob = ref() // need to pass this value to JobBoard
</script>

<template>
  <div class="flex flex-col w-full">
    <div class="flex bg-title w-full h-40">
      <div class="flex px-42 w-full flex-col justify-center items-center space-y-7 text-xl">
        <div class="text-white min-w-300 w-1/2">
          Search for your jobs now
        </div>
        <div class="flex flex-row space-x-10 min-w-300 w-1/2">
          <label class="input-box">
            <input v-model="keyword" type="search" class="grow pl-3 placeholder-black" placeholder="Keywords" />
          </label>
          <div class="relative">
            <img src="../assets/case.svg" class="absolute px-2 w-12 left-3 top-2 z-10" alt="case"/>
            <select v-model="jobType" class="select rounded-2xl select-lg pl-18 z-0">
              <option disabled value="">Job Type</option>
              <option>Software & Application Development</option>
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
          <div class="relative">
            <img src="../assets/location.svg" class="absolute px-2 w-12 left-3 top-2 z-10" alt="case"/>
            <select v-model="location" class="select rounded-2xl select-lg pl-18 z-0 w-70">
              <option disabled value="">Job Type</option>
              <option>กรุงเทพมหานคร</option>
              <option>สมุทรสาคร</option>
              <option>นครปฐม</option>
              <option>ปทุมธานี</option>
              <option>นนทบุรี</option>
            </select>
          </div>
          <button @click="findJobs" class="btn shadow-none bg-lighter border-0 h-12 rounded-2xl text-white text-xl font-extralight">Find Now</button>
        </div>
      </div>

    </div>

    <div class="flex w-full flex-row px-42 py-10 bg-[#f2f6fc] space-x-20">

      <div id="job-box-column" class="flex w-1/3 flex-col space-y-10 mr-10">
        <JobBox v-for="job in jobs" :key=job.id :jobInfo=job :h-r="false" @click="selectedJob=job"/>
      </div>

      <JobFull :job-info="selectedJob" class="ml-7"/>

    </div>
  </div>
</template>

<style scoped>

</style>