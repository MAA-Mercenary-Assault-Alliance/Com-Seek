<script setup lang="ts" xmlns:hover="http://www.w3.org/1999/xhtml">
import { ref } from 'vue'
import CompanyBox from "../components/CompanyBox.vue";
import { TE_Info_Company, Student_Info, TE_Info } from "../components/temp_template";
import StudentBox from "../components/StudentBox.vue";
import JobBoxAdmin from "../components/JobBoxAdmin.vue";
const activeTab = ref('company')

const template_list = [1,2,3]
</script>

<template>
  <div class="flex flex-col w-full">

    <div class="flex w-full h-40 px-42 items-center bg-[url('/adminBar.png')] bg-right">
      <span class="text-white text-5xl font-bold">Admin Dashboard</span>
    </div>

    <div id="content" class="flex flex-row space-x-10">
      <div id="sidebar" class="flex flex-col w-1/5">
        <div class="sidebar-properties" :class="[activeTab === 'company' ? 'sidebar-highlight' : 'sidebar-normal']" @click="activeTab = 'company'">
          <span class="ml-15 text-3xl">Company</span>
        </div>
        <div class="sidebar-properties" :class="[activeTab === 'alumni' ? 'sidebar-highlight' : 'sidebar-normal']" @click="activeTab = 'alumni'">
          <span class="ml-15 text-3xl">Alumni</span>
        </div>
        <div class="sidebar-properties" :class="[activeTab === 'jobPosting' ? 'sidebar-highlight' : 'sidebar-normal']" @click="activeTab = 'jobPosting'">
          <span class="ml-15 text-3xl">Job Posting</span>
        </div>
      </div>
      <div id="boxes" class="flex flex-col mt-10 space-y-6 w-4/5 pr-42">
        <template v-if="activeTab === 'company'">
          <CompanyBox v-for="stuff in template_list" :key="stuff" :company-info="TE_Info_Company"/>
        </template>
        <template v-if="activeTab === 'alumni'">
          <StudentBox v-for="stuff in template_list" :key="stuff" :student-info="Student_Info"/>
        </template>
        <template v-if="activeTab === 'jobPosting'">
          <JobBoxAdmin v-for="stuff in template_list" :key="stuff" :job-info="TE_Info"/>
        </template>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>