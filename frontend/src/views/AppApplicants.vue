<script setup lang="ts">
import JobFull from "../components/JobFull.vue";
import { TE_Info } from '../components/temp_template';
import ApplicantsColumn from "../components/ApplicantsColumn.vue";

import { useRoute } from 'vue-router'
const route = useRoute()
const jobId = route.params.id
</script>

<template>
  <div class="flex flex-col w-full space-y-10">
    <div class="flex bg-[#153243] w-full h-40 px-42 items-center">
      <span class="text-white text-5xl font-bold">Applicants</span>
    </div>
    <div class="flex flex-row w-full px-42 bg-[#f2f6fc] space-x-20">
      <JobFull :job-info="TE_Info" />
      <ApplicantsColumn />
    </div>
  </div>
</template>

<style scoped>

</style>