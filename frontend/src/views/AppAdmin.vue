<script setup lang="ts" xmlns:hover="http://www.w3.org/1999/xhtml">
import { ref, onMounted, watch } from 'vue'
import CompanyBox from "../components/CompanyBox.vue";
import { api } from '../../api/client.js';

import StudentBox from "../components/StudentBox.vue";
import JobBoxAdmin from "../components/JobBoxAdmin.vue";
import PureSearchBar from "../components/PureSearchBar.vue";
const activeTab = ref('company')

const company_list = ref([])
const student_list = ref([])
const job_list = ref([])
let isLoading = ref(true)

async function getCompanies() {
  try {
    const res = await api.get("/admin/companies")
    company_list.value = res.data.companies
    console.log("RES", res.data)
    console.log("Fetched companies:", company_list.value)
    isLoading.value = false
    return company_list.value
  } catch (error) {
    console.error("Error fetching companies:", error)
    return []
  }

}

async function getStudents() {
  try {
    const res = await api.get("/admin/students")
    student_list.value = res.data.students
    console.log("Fetched students:", student_list.value)
    return student_list.value
  } catch (error) {
    console.error("Error fetching students:", error)
    return []
  }
}

async function getJobs() {
  try {
    const res = await api.get("/admin/jobs")
    job_list.value = res.data.jobs
    console.log("Fetched jobs:", job_list.value)
    return job_list.value
  } catch (error) {
    console.error("Error fetching jobs:", error)
    return []
  }
}

onMounted( async() => {
  await getCompanies()
})

async function selectTab(tab: string) {
  activeTab.value = tab

  if (tab === 'companies') await getCompanies()
  else if (tab === 'students') await getStudents()
  else if (tab === 'jobs') await getJobs()
}

</script>

<template>
  <div class="flex flex-col w-full">

    <div class="flex w-full h-40 px-42 items-center bg-[url('/adminBar.png')] bg-right">
      <span class="text-white text-5xl font-bold">Admin Dashboard</span>
    </div>

    <div id="content" class="flex flex-row space-x-10">
      <div id="sidebar" class="hidden sm:flex flex-col w-1/5">
        <div class="sidebar-properties" :class="[activeTab === 'companies' ? 'sidebar-highlight' : 'sidebar-normal']" @click="selectTab('companies')">
          <span class="ml-4 xl:ml-15 text-xl xl:text-3xl">Company</span>
        </div>
        <div class="sidebar-properties" :class="[activeTab === 'students' ? 'sidebar-highlight' : 'sidebar-normal']" @click="selectTab('students')">
          <span class="ml-4 xl:ml-15 text-xl xl:text-3xl">Student</span>
        </div>
        <div class="sidebar-properties" :class="[activeTab === 'jobs' ? 'sidebar-highlight' : 'sidebar-normal']" @click="selectTab('jobs')">
          <span class="ml-4 xl:ml-15 text-xl xl:text-3xl">Job Posting</span>
        </div>
      </div>
      <div v-if="!isLoading" id="boxes" class="flex relative flex-col mt-10 space-y-6 w-4/5 xl:pr-42 pb-10">
        <PureSearchBar></PureSearchBar>
        <template v-if="activeTab === 'companies'">
          <CompanyBox v-for="company in company_list" :key="company.id" :company-info="company"/>
        </template>
        <template v-if="activeTab === 'students'">
          <StudentBox v-for="student in student_list" :key="student.id" :student-info="student"/>
        </template>
        <template v-if="activeTab === 'jobs'">
          <JobBoxAdmin v-for="job in job_list" :key="job.id" :job-info="job"/>
        </template>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>