<script setup lang="ts">

import {TE_Info} from "../components/temp_template";
import JobFull from "../components/JobFull.vue";
import ApplicantsColumn from "../components/ApplicantsColumn.vue";
</script>

<template>
  <div class="flex flex-col w-full space-y-10">
    <div class="flex bg-title w-full h-40 px-42 items-center">
      <span class="text-white text-5xl font-bold">Applicants</span>
    </div>
    <div class="flex flex-row w-full xl:px-42 bg-[#f2f6fc] space-x-5 xl:space-x-20 pb-20">
      <JobFull :job-info="TE_Info" />
      <ApplicantsColumn />
    </div>
  </div>
</template>

<style scoped>

</style>