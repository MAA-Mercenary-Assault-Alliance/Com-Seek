<script setup lang="ts">
import JobFull from "../components/JobFull.vue";
import { TE_Info } from '../components/temp_template';
import ApplicantsColumn from "../components/ApplicantsColumn.vue";
import { api } from '../../api/client.js';

import { useRoute } from 'vue-router'
import {ref} from "vue";
const route = useRoute()
const jobId = route.params.id
const isLoading = ref(true)
const job = ref(null)

async function findThisJob() {
  try {
    isLoading.value = true;
    const res = await api.get("/job", {
      params: {
        id: jobId
      }
    })

    job.value = res.data.jobs[0]
    console.log("Fetched this job:", job.value)
  } catch (error) {
    console.error("Error fetching jobs:", error)
  } finally {
    isLoading.value = false;
  }
}

</script>

<template>
  <div class="flex flex-col w-full space-y-10">
    <div class="flex bg-title w-full h-40 px-42 items-center">
      <span class="text-white text-5xl font-bold">Applicants</span>
    </div>
    <div class="flex flex-row w-full xl:px-42 bg-[#f2f6fc] space-x-5 xl:space-x-20 pb-20">
      <JobFull :job-info="job" />
      <ApplicantsColumn :applicants="job"/>
<!--      TODO: pass applicants not job-->
    </div>
  </div>
</template>

<style scoped>

</style>