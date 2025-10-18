<script setup lang="ts">
import { ref, watch, reactive} from 'vue'
import ConfirmBoxGeneral from "@/components/ConfirmBoxGeneral.vue";
import {marked} from "marked";
import { api } from "../../api/client.js"
import {useRouter} from "vue-router";

const router = useRouter()
const confirmBox = ref(null)

const jobName = ref('')
const employmentStatus = ref('')
const location = ref('')
const salaryMin = ref()
const salaryMax = ref()
const expMin = ref()
const expMax = ref()
const jobType = ref('')
const jobDesc = ref('')
let desc_html = ref('There is no description yet')

watch(jobDesc, async (newVal) => {
  if (newVal == "") {
    desc_html.value = "There is no description yet"
    return
  }
  desc_html.value = await marked(newVal)
})

const preview = ref(false)

const errors = reactive({
  jobName: '',
  employmentStatus: '',
  location: '',
  salaryMin: '',
  salaryMax: '',
  expMin: '',
  expMax: '',
  jobType: '',
  jobDesc: ''
})

const validateForm = () => {
  let valid = true

  // Reset errors
  Object.keys(errors).forEach(key => errors[key] = '')

  if (!jobName.value.trim()) {
    errors.jobName = 'Job Name is required'
    valid = false
  }

  if (!employmentStatus.value) {
    errors.employmentStatus = 'Employment Status is required'
    valid = false
  }

  if (!location.value) {
    errors.location = 'Location is required'
    valid = false
  }

  if (!jobType.value) {
    errors.jobType = 'Job Type is required'
    valid = false
  }

  if (!jobDesc.value.trim()) {
    errors.jobDesc = 'Job description is required'
    valid = false
  }

// Salary validation
  if (!salaryMin.value) {
    errors.salaryMin = 'Minimum salary is required'
    valid = false
  } else if (isNaN(Number(salaryMin.value))) {
    errors.salaryMin = 'Minimum salary must be a number'
    valid = false
  }

  if (!salaryMax.value) {
    errors.salaryMax = 'Maximum salary is required'
    valid = false
  } else if (isNaN(Number(salaryMax.value))) {
    errors.salaryMax = 'Maximum salary must be a number'
    valid = false
  }

// Experience validation
  if (!expMin.value) {
    expMin.value = 0
  } else if (isNaN(Number(expMin.value))) {
    errors.expMin = 'Minimum experience must be a number'
    valid = false
  }

  if (!expMax.value) {
    expMax.value = 0
  } else if (isNaN(Number(expMax.value))) {
    errors.expMax = 'Maximum experience must be a number'
    valid = false
  }

// Optional: check if min > max
  if (
      salaryMin.value && salaryMax.value &&
      Number(salaryMin.value) > Number(salaryMax.value)
  ) {
    errors.salaryMin = 'Min salary cannot be greater than max'
    valid = false
  }

  if (
      expMin.value && expMax.value &&
      Number(expMin.value) > Number(expMax.value)
  ) {
    errors.expMin = 'Min experience cannot be greater than max'
    valid = false
  }


  return valid
}

const submitJob = async () => {
  try {

    if (!validateForm()) {
      alert("Please fix the errors in the form")
      return
    }

    const payload = {
      title: jobName.value,
      location: location.value,
      job_type: jobType.value,
      employment_status: employmentStatus.value,
      min_salary: Number(salaryMin.value),
      max_salary: Number(salaryMax.value),
      min_experience: Number(expMin.value),
      max_experience: Number(expMax.value),
      description: jobDesc.value
    }

    const res = await api.post(`/job`, payload)

    console.log("Job Created: ", res.data)
    alert('Job created successfully!')
    router.back()

  } catch (err) {
    if (err.response) {
      // Backend returned an error response (like 403)
      console.error('Error:', err.response.data)
      alert(err.response.data.error || 'Something went wrong')
    } else if (err.request) {
      // No response (maybe network/CORS issue)
      console.error('No response from server:', err.request)
      alert('No response from server')
    } else {
      // Something else happened before the request was sent
      console.error('Axios error:', err.message)
      alert('Request setup failed')
    }
  }
}

</script>

<template>
  <div class="flex flex-col w-full space-y-10 2xl:px-60 lg:px-42">
    <div class="flex w-full h-40 items-center">
      <span class="text-black text-5xl font-bold">Create a Job</span>
    </div>

    <div id="input-field">
      <div class="grid grid-cols-2 space-y-5">

        <div class="job-input-box">
          <div class="title-and-error-box">
            <span class="text-xl">Job Name</span>
            <label v-if="errors.jobName" class="text-red-500 text-sm">{{ errors.jobName }}</label>
          </div>
          <label class="input-box-gray">
            <input v-model="jobName" type="search" class="grow pl-3" placeholder="E.g. Account Manager (Sales Engineer)" />
          </label>
        </div>

        <div class="job-input-box">
          <div class="title-and-error-box">
            <span class="text-xl">Employment Status</span>
            <label v-if="errors.employmentStatus" class="text-red-500 text-sm">{{ errors.employmentStatus }}</label>
          </div>
          <select v-model="employmentStatus" class="input-select">
            <option disabled value="">Employment Status</option>
            <option>full-time</option>
            <option>part-time</option>
            <option>contract</option>
            <option>internship</option>
          </select>
        </div>

        <div class="job-input-box">
          <div class="title-and-error-box">
            <span class="text-xl">Location</span>
            <label v-if="errors.location" class="text-red-500 text-sm">{{ errors.location }}</label>
          </div>
          <label class="input-box-gray">
            <input v-model="location" type="search" class="grow pl-3" placeholder="E.g. กรุงเทพมหานคร" />
          </label>
        </div>

        <div class="job-input-box">
          <div class="title-and-error-box">
            <span class="text-xl">Salary Range</span>
            <span v-if="errors.salaryMin" class="text-red-500 text-sm">{{ errors.salaryMin }}</span>
            <span v-if="errors.salaryMax" class="text-red-500 text-sm">{{ errors.salaryMax }}</span>
          </div>
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
          <div class="title-and-error-box">
            <span class="text-xl">Employment Status</span>
            <label v-if="errors.jobType" class="text-red-500 text-sm">{{ errors.jobType }}</label>
          </div>
          <select v-model="jobType" class="input-select">
            <option disabled value="">Job Type</option>
            <option>Software & Application Development</option>
            <option>Data & AI</option>
            <option>Cloud & Infrastructure</option>
            <option>Cybersecurity</option>
            <option>Product & Design</option>
            <option>Testing & Quality</option>
            <option>Hardware & Electronics</option>
            <option>Management & Leadership</option>
            <option>IT Support & Operations</option>

          </select>
        </div>

        <div class="job-input-box">
          <div>
            <div class="title-and-error-box">
              <span class="text-xl">Experience Range</span>
              <span class="text-xl text-gray-400"> [Year] *leave empty if require no experience</span>
            </div>
            <div v-if="errors.expMin || errors.expMax" class="flex space-x-5 mt-3">
              <span v-if="errors.expMin" class="text-red-500 text-sm">{{ errors.expMin }}</span>
              <span v-if="errors.expMax" class="text-red-500 text-sm">{{ errors.expMax }}</span>
            </div>
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
          <div class="flex space-x-5">
            <span class="text-xl">Full Job Description</span>
            <span class="text-lg text-gray-400">*support markdown</span>
            <div id="preview-slider" class="flex justify-center">
              <div class="bg-base-200 rounded-full p-1 flex w-64 border-gray-300 border-2">
                <!-- Student -->
                <button
                    type="button"
                    @click="preview=false"
                    :class="[
                  'flex-1 py-2 rounded-full font-medium transition',
                  !preview ? 'bg-lighter text-white shadow' : 'text-gray-600'
                ]"
                >
                  Raw
                </button>

                <!-- Alumni -->
                <button
                    type="button"
                    @click="preview=true"
                    :class="[
                  'flex-1 py-2 rounded-full font-medium transition',
                  preview ? 'bg-lighter text-white shadow' : 'text-gray-600'
                ]"
                >
                  Preview
                </button>
              </div>
            </div>
          </div>
          <span v-if="errors.jobDesc" class="text-red-500 text-sm">{{ errors.jobDesc }}</span>
          <label>
            <textarea v-model="jobDesc" type="search" v-if="!preview" class="bg-white rounded-2xl border border-gray-300 placeholder-gray-300 grow pl-3 xl:w-5/6 w-full h-70 py-2" placeholder="E.g. Account Manager (Sales Engineer)" />
          </label>
        </div>
      </div>

      <div v-if="preview" class="prose">
        <div v-html="desc_html"></div>
      </div>

      <div class="flex flex-row space-x-10 mt-10 pb-20">
        <button class="btn shadow-none border-0 h-15 rounded-3xl text-white text-lg font-extralight px-7 bg-[#44B15B]" @click="submitJob ">Confirm</button>
        <button class="btn shadow-none  border-0 h-15 rounded-3xl text-white text-lg font-extralight px-7 bg-gray-300" @click="confirmBox.open()">Cancel</button>
      </div>

      <confirm-box-general
          ref="confirmBox"
          :message="'The data on this page will be lost'"
          @accept="router.back"
      ></confirm-box-general>

    </div>
  </div>
</template>

<style scoped>

</style>