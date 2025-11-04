<script setup lang="ts">
import { ref } from 'vue'
import { TE_Info } from './temp_template'
import JobBox from "./JobBox.vue";
import JobFull from "./JobFull.vue";

const LIST_JOBOBJ = ref('')
const selectedJobID = ref('')

</script>

<template>
 <div class="flex w-full flex-row px-42 py-10 bg-background space-x-20">

   <div id="job-box-column" class="flex w-1/3 flex-col space-y-10 mr-10">
     <JobBox :jobInfo=TE_Info :h-r="false"/>

     <JobBox :jobInfo=TE_Info :h-r="false"/>
   </div>

   <JobFull :job-info="TE_Info" class="ml-7"/>

 </div>
</template>

<style scoped>

</style>