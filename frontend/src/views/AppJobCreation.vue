<script setup lang="ts">
import { ref, watch } from 'vue'
import axios from 'axios'

const jobName = ref('')
const employmentStatus = ref('')
const location = ref('')
const salaryMin = ref('')
const salaryMax = ref('')
const expMin = ref('')
const expMax = ref('')
const jobType = ref('')
const jobDesc = ref('')

</script>

<template>
  <div class="flex flex-col w-full space-y-10 2xl:px-60 lg:px-42">
    <div class="flex w-full h-40 items-center">
      <span class="text-black text-5xl font-bold">Create a Job</span>
    </div>

    <div id="input-field">
      <div class="grid grid-cols-2 space-y-5">

        <div class="job-input-box">
          <span class="text-xl">Job Name</span>
          <label class="input-box-gray">
            <input v-model="jobName" type="search" class="grow pl-3" placeholder="E.g. Account Manager (Sales Engineer)" />
          </label>
        </div>

        <div class="job-input-box">
          <span class="text-xl">Employment Status</span>
          <select v-model="employmentStatus" class="input-select">
            <option disabled value="">Employment Status</option>
            <option>Full-Time</option>
            <option>Part-Time</option>
          </select>
        </div>

        <div class="job-input-box">
          <span class="text-xl">Location</span>
          <select v-model="location" class="input-select">
            <option disabled value="">Location</option>
            <option>กรุงเทพมหานคร</option>
            <option>สมุทรสาคร</option>
            <option>นครปฐม</option>
            <option>ปทุมธานี</option>
            <option>นนทบุรี</option>
          </select>
        </div>

        <div class="job-input-box">
          <span class="text-xl">Salary Range</span>
          <div class="flex flex-row justify-start">
            <label class="input-box-half">
              <input v-model="salaryMin" type="search" class="grow pl-3" placeholder="E.g. 20000฿" />
            </label>
            <span class="inline-flex items-center justify-center px-5 text-xl"> to </span>
            <label class="input-box-half">
              <input v-model="salaryMax" type="search" class="grow pl-3" placeholder="E.g. 30000฿" />
            </label>
          </div>
        </div>

        <div class="job-input-box">
          <span class="text-xl">Job Type</span>
          <select v-model="jobType" class="input-select">
            <option disabled value="">Job Type</option>
            <option>Full-Stack Developer</option>
            <option>Front-End Developer</option>
            <option>Back-End Developer</option>
            <option>DevOps Engineer</option>
            <option>Add more later</option>
          </select>
        </div>

        <div class="job-input-box">
          <div>
            <span class="text-xl">Experience Range</span>
            <span class="text-xl text-gray-400"> [Year] *leave empty if require no experience</span>
          </div>
          <div class="flex flex-row justify-start">
            <label class="input-box-half">
              <input v-model="expMin" type="search" class="grow pl-3" placeholder="E.g. 0" />
            </label>
            <span class="inline-flex items-center justify-center px-5 text-xl"> to </span>
            <label class="input-box-half">
              <input v-model="expMax" type="search" class="grow pl-3" placeholder="E.g. 2" />
            </label>
          </div>
        </div>

      </div>

      <div id="desc" class="mt-10">
        <div class="job-input-box">
          <div class="space-x-5">
            <span class="text-xl">Full Job Description</span>
            <span class="text-lg text-gray-400">*support markdown</span>
          </div>
          <label>
            <textarea v-model="jobDesc" type="search" class="bg-white rounded-2xl border border-gray-300 placeholder-gray-300 grow pl-3 xl:w-5/6 w-full h-70 py-2" placeholder="E.g. Account Manager (Sales Engineer)" />
          </label>
        </div>
      </div>

      <div class="flex flex-row space-x-10 mt-10">
        <button class="btn shadow-none border-0 h-15 rounded-3xl text-white text-lg font-extralight px-7 bg-[#44B15B]" >Confirm</button>
        <button class="btn shadow-none  border-0 h-15 rounded-3xl text-white text-lg font-extralight px-7 bg-gray-300" >Cancel</button>
      </div>

    </div>
  </div>
</template>

<style scoped>

</style>