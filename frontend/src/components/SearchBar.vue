<script setup lang="ts">
import { ref, watch } from 'vue'

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

// BELOW are to do something if anything changes (probably Search)
// watch([keyword, jobType, location], ([newKeyword, newJobType, newLocation]) => {
//   console.log('keyword:', newKeyword)
//   console.log('jobType:', newJobType)
//   console.log('location:', newLocation)
// })

const selectedJobID = ref('') // need to pass this value to JobBoard
</script>

<template>
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
            <option>Full-Stack Developer</option>
            <option>Front-End Developer</option>
            <option>Back-End Developer</option>
            <option>DevOps Engineer</option>
            <option>Add more later</option>
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
        <button @click="something" class="btn shadow-none bg-lighter border-0 h-12 rounded-2xl text-white text-xl font-extralight">Find Now</button>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>