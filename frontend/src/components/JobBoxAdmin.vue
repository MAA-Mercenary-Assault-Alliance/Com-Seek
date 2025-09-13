<script setup lang="ts">
import {ref} from 'vue'
import {JobTemplate} from './temp_template'
const props = defineProps<{
  jobInfo: JobTemplate
}>();

const applied = ref(0)
const new_app = ref(0)
</script>

<template>
  <div id="job-box" class="flex relative rounded-2xl flex-row p-5 pb-4 space-x-5 box-shadow bg-white" @click="hR && $emit('click')"  :class="{ 'cursor-pointer': hR }">

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

    <div class="flex flex-col ml-auto items-end space-y-3 mr-4">
      <span class="text-gray-500">16/07/2025</span>
      <img src="../assets/newTab.svg" class="w-5 h-5" alt="new-tab-icon"/>
      <div class="flex flex-row mt-auto space-x-6">
        <button class="btn shadow-none bg-[#1F7AB9] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7">Accept</button>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7">Reject</button>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>