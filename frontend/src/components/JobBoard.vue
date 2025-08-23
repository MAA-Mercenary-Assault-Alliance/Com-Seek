<script setup lang="ts">
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

const jobHtml = marked(jobMarkdown)

</script>

<template>
 <div class="flex w-full flex-row px-42 my-10">

   <div id="job-box-column" class="flex border-1 w-1/3 flex-col space-y-10 mr-10">

     <div id="job-box" class="flex relative border-1 rounded-2xl flex-row p-4 pb-7 space-x-5">

       <img src="../assets/company.jpg" class="w-20 h-20 rounded-2xl" alt="company-logo"/>
       <div id="job-box-content" class="flex mr-2 flex-col space-y-1.5">
         <span class="underline">ACCOUNT MANAGER (Sales Engineer)</span>
         <span class="mb-2">TE Connectivity</span>

         <div class="in-line-stat">
           <img src="../assets/money.svg" class="in-line-icon" alt="money-icon"/>
           <span>฿12000-฿15000 per month</span>
         </div>
         <div class="in-line-stat">
           <img src="../assets/cap.svg" class="in-line-icon" alt="cap-icon"/>
           <span>1-2 years</span>
         </div>
         <div class="in-line-stat">
           <img src="../assets/location.svg" class="in-line-icon" alt="location-icon"/>
           <span>จตุจักร, กรุงเทพมหานคร</span>
         </div>
       </div>
       <span class="absolute bottom-3 right-5 text-gray-500">24 days ago</span>
     </div>

     <div class="flex border-1 rounded-2xl h-40">

     </div>
   </div>


   <div id="job full" class="flex relative border-1 w-2/3 ml-7 px-25 py-10 flex-col rounded-2xl">

     <div id="title-box" class="flex row items-center mt-3">
       <img src="../assets/company.jpg" class="w-20 h-20 rounded-2xl" alt="company-logo"/>
       <div id="title" class="flex flex-col ml-7 space-y-4">
         <span class="text-2xl">TE Connectivity</span>
         <span class="underline">ACCOUNT MANAGER (Sales Engineer)</span>
       </div>
       <button class="btn shadow-none bg-[#1F7AB9] border-0 h-12 rounded-2xl text-white text-xl font-extralight ml-auto mt-5">Apply Now</button>
     </div>
     <span class="absolute top-5 right-5 text-gray-500">16/07/2025</span>

     <span class="mt-10 text-lg">Basic Requirements</span>
     <div id="requirements-box" class="grid grid-cols-2 gap-4 w-full max-w-4xl mt-6 space-y-2">
       <div class="requirements-div">
         <img src="../assets/location.svg" class="requirements-icon" alt="location-icon"/>
         <span>จตุจักร, กรุงเทพมหานคร</span>
       </div>
       <div class="requirements-div">
         <img src="../assets/money.svg" class="requirements-icon" alt="money-icon"/>
         <span>฿12000-฿15000 per month</span>
       </div>
       <div class="requirements-div">
         <img src="../assets/time.svg" class="requirements-icon" alt="time-icon"/>
         <span>Full-time</span>
       </div>
       <div class="requirements-div">
         <img src="../assets/cap.svg" class="requirements-icon" alt="cap-icon"/>
         <span>1-2 years</span>
       </div>
     </div>

     <span class="mt-12 text-lg">Full Job Description</span>
     <div class="prose">
       <div v-html="jobHtml"></div>
     </div>

   </div>
 </div>
</template>

<style scoped>

</style>