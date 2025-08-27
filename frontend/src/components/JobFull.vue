<script setup lang="ts">
import {JobTemplate, TE_Info} from './temp_template'
import {marked} from "marked";
const props = defineProps<{
  jobInfo: JobTemplate
}>();

const desc_html = marked(props.jobInfo.desc)
</script>

<template>
  <div id="job full" class="flex relative border-1 w-2/3 ml-7 px-25 py-10 flex-col rounded-2xl shadow-lg bg-white">

    <div id="title-box" class="flex row items-center mt-3">
      <img src="../assets/company.jpg" class="w-20 h-20 rounded-2xl" alt="company-logo"/>
      <div id="title" class="flex flex-col ml-7 space-y-4">
        <span class="text-2xl">{{ jobInfo.company }}</span>
        <span class="underline">{{ jobInfo.name }}</span>
      </div>
      <button class="btn shadow-none bg-[#1F7AB9] border-0 h-12 rounded-2xl text-white text-xl font-extralight ml-auto mt-6">Apply Now</button>
    </div>
    <span class="absolute top-5 right-5 text-gray-500">16/07/2025</span>

    <span class="mt-10 text-lg">Basic Requirements</span>
    <div id="requirements-box" class="grid grid-cols-2 gap-4 w-full max-w-4xl mt-6 space-y-2">
      <div class="requirements-div">
        <img src="../assets/location.svg" class="requirements-icon" alt="location-icon"/>
        <span>{{ jobInfo.location }}</span>
      </div>
      <div class="requirements-div">
        <img src="../assets/money.svg" class="requirements-icon" alt="money-icon"/>
        <span>{{ jobInfo.salary_range }}</span>
      </div>
      <div class="requirements-div">
        <img src="../assets/time.svg" class="requirements-icon" alt="time-icon"/>
        <span>{{ jobInfo.type }}</span>
      </div>
      <div class="requirements-div">
        <img src="../assets/cap.svg" class="requirements-icon" alt="cap-icon"/>
        <span>{{ jobInfo.experience }}</span>
      </div>
    </div>

    <span class="mt-12 text-lg">Full Job Description</span>
    <div class="prose">
      <div v-html="desc_html"></div>
    </div>

  </div>
</template>

<style scoped>

</style>