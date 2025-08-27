<script setup lang="ts">
import desc_html from "./JobBoard.vue"
</script>

<template>
  <div id="job-box" class="flex relative border-1 rounded-2xl flex-row p-4 pb-7 space-x-5 shadow-lg bg-white">

    <img src="../assets/company.jpg" class="w-20 h-20 rounded-2xl" alt="company-logo"/>
    <div id="job-box-content" class="flex mr-2 flex-col space-y-1.5">
      <span class="underline">{{ CHOSEN_JOB_TEMPLATE.name }}</span>
      <span class="mb-2">{{ CHOSEN_JOB_TEMPLATE.company }}</span>

      <div class="in-line-stat">
        <img src="../assets/money.svg" class="in-line-icon" alt="money-icon"/>
        <span>{{ CHOSEN_JOB_TEMPLATE.salary_range }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/cap.svg" class="in-line-icon" alt="cap-icon"/>
        <span>{{ CHOSEN_JOB_TEMPLATE.experience }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/location.svg" class="in-line-icon" alt="location-icon"/>
        <span>{{ CHOSEN_JOB_TEMPLATE.location }}</span>
      </div>
    </div>
    <span class="absolute bottom-3 right-5 text-gray-500">24 days ago</span>
  </div>
</template>

<style scoped>

</style>