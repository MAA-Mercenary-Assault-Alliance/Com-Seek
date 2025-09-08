<script setup lang="ts">
import {ref} from 'vue'
import {JobTemplate} from './temp_template'
const props = defineProps<{
  jobInfo: JobTemplate,
  HR: boolean
}>();

const applied = ref(0)
const new_app = ref(0)
</script>

<template>
  <div id="job-box" class="flex relative rounded-2xl flex-row p-4 pb-7 space-x-5 box-shadow bg-white">

    <img src="../assets/company.jpg" class="w-20 h-20 rounded-2xl" alt="company-logo"/>
    <div id="job-box-content" class="flex mr-2 flex-col space-y-1.5">
      <span class="underline">{{ jobInfo.name }}</span>
      <span class="mb-2">{{ jobInfo.company }}</span>

      <div class="in-line-stat">
        <img src="../assets/money.svg" class="in-line-icon" alt="money-icon"/>
        <span>{{ jobInfo.salary_range }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/cap.svg" class="in-line-icon" alt="cap-icon"/>
        <span>{{ jobInfo.experience }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/location.svg" class="in-line-icon" alt="location-icon"/>
        <span>{{ jobInfo.location }}</span>
      </div>
    </div>
    <span class="absolute bottom-3 right-5 text-gray-500">24 days ago</span>

    <div v-if=HR>
      <span class="absolute top-5 right-10 text-gray-500">Applied: {{applied}}</span>
      <span class="absolute top-12   right-10 text-gray-500">New: {{new_app}}</span>
    </div>
  </div>
</template>

<style scoped>

</style>