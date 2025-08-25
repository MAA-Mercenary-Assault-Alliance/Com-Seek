<script setup lang="ts">
import { ref } from 'vue'
import { marked } from "marked"

const jobMarkdown = `
The Account Manager is responsible for building and maintaining strong client relationships while driving sales growth. This role involves identifying customer needs, presenting tailored solutions, and ensuring long-term satisfaction. The Account Manager acts as the main point of contact between the company and its clients, balancing technical knowledge with strong communication skills to deliver value.

**Key Responsibilities:**

- Develop and manage client accounts to achieve sales targets.
- Identify opportunities for upselling and cross-selling.
- Collaborate with technical teams to provide tailored solutions.
- Maintain regular communication to ensure customer satisfaction.
- Prepare reports, proposals, and sales forecasts.

**Skills & Qualifications:**

- Strong interpersonal and negotiation skills.
- Technical knowledge relevant to the industry.
- Ability to manage multiple accounts and priorities.
- Excellent problem-solving and communication abilities.
`

const LIST_JOBOBJ = ref('')
const selectedJobID = ref('')

const CHOSEN_JOB_TEMPLATE = {
  name: "ACCOUNT MANAGER (Sales Engineer)",
  company: "TE Connectivity",
  location: "กรุงเทพมหานคร",
  salary_range: "฿12000-฿15000 per month", // MAKE THIS FLOOR AND CEILING
  type: "full-time",
  experience: "1-2 years",
  desc: jobMarkdown,
}

const desc_html = marked(CHOSEN_JOB_TEMPLATE.desc)

</script>

<template>
 <div class="flex w-full flex-row px-42 py-10 bg-[#f2f6fc]">

   <div id="job-box-column" class="flex border-1 w-1/3 flex-col space-y-10 mr-10">

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

     <div class="flex border-1 rounded-2xl h-40 bg-white">

     </div>

   </div>


   <div id="job full" class="flex relative border-1 w-2/3 ml-7 px-25 py-10 flex-col rounded-2xl shadow-lg bg-white">

     <div id="title-box" class="flex row items-center mt-3">
       <img src="../assets/company.jpg" class="w-20 h-20 rounded-2xl" alt="company-logo"/>
       <div id="title" class="flex flex-col ml-7 space-y-4">
         <span class="text-2xl">{{ CHOSEN_JOB_TEMPLATE.company }}</span>
         <span class="underline">{{ CHOSEN_JOB_TEMPLATE.name }}</span>
       </div>
       <button class="btn shadow-none bg-[#1F7AB9] border-0 h-12 rounded-2xl text-white text-xl font-extralight ml-auto mt-6">Apply Now</button>
     </div>
     <span class="absolute top-5 right-5 text-gray-500">16/07/2025</span>

     <span class="mt-10 text-lg">Basic Requirements</span>
     <div id="requirements-box" class="grid grid-cols-2 gap-4 w-full max-w-4xl mt-6 space-y-2">
       <div class="requirements-div">
         <img src="../assets/location.svg" class="requirements-icon" alt="location-icon"/>
         <span>{{ CHOSEN_JOB_TEMPLATE.location }}</span>
       </div>
       <div class="requirements-div">
         <img src="../assets/money.svg" class="requirements-icon" alt="money-icon"/>
         <span>{{ CHOSEN_JOB_TEMPLATE.salary_range }}</span>
       </div>
       <div class="requirements-div">
         <img src="../assets/time.svg" class="requirements-icon" alt="time-icon"/>
         <span>{{ CHOSEN_JOB_TEMPLATE.type }}</span>
       </div>
       <div class="requirements-div">
         <img src="../assets/cap.svg" class="requirements-icon" alt="cap-icon"/>
         <span>{{ CHOSEN_JOB_TEMPLATE.experience }}</span>
       </div>
     </div>

     <span class="mt-12 text-lg">Full Job Description</span>
     <div class="prose">
       <div v-html="desc_html"></div>
     </div>

   </div>
 </div>
</template>

<style scoped>

</style>