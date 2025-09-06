<script setup lang="ts">
import {ref} from 'vue'
import {CompanyTemplate} from './temp_template'
const props = defineProps<{
  companyInfo: CompanyTemplate,
}>();
</script>

<template>
  <div id="company-box" class="flex relative rounded-2xl flex-row p-4 space-x-7 box-shadow bg-white cursor-pointer" @click="$emit('click')">
    <div class="flex items-center">
      <img src="../assets/company.jpg" class="w-30 h-30 rounded-2xl ml-2" alt="company-logo"/>
    </div>
    <div id="job-box-content" class="flex mr-2 flex-col space-y-1.5">
      <span class="text-2xl">{{ companyInfo.name }}</span>

      <div class="in-line-stat">
        <img src="../assets/money.svg" class="in-line-icon" alt="money-icon"/>
        <span>{{ companyInfo.number }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/cap.svg" class="in-line-icon" alt="cap-icon"/>
        <span>{{ companyInfo.email }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/location.svg" class="in-line-icon" alt="location-icon"/>
        <span>{{ companyInfo.website }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/location.svg" class="in-line-icon" alt="location-icon"/>
        <span>{{ companyInfo.location }}</span>
      </div>
    </div>

    <div class="flex flex-col ml-auto items-end space-y-3 mr-4">
      <span>16/07/2025</span>
      <img src="../assets/company.jpg" class="w-5 h-5" alt="company-logo"/>
      <div class="flex flex-row mt-auto space-x-6">
        <button class="btn shadow-none bg-[#1F7AB9] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7">Accept</button>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7">Reject</button>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>