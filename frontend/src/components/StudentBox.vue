<script setup lang="ts">
import {StudentTemplate} from './temp_template'
const props = defineProps<{
  studentInfo: StudentTemplate,
}>();

</script>

<template>
  <div id="company-box" class="flex relative rounded-2xl flex-row p-4 space-x-7 box-shadow bg-white cursor-pointer" @click="$emit('click')">
    <div class="flex items-center">
      <img src="/nagi.png" class="w-30 h-30 rounded-full ml-2" alt="company-logo"/>
    </div>
    <div id="student-box-content" class="flex mr-2 flex-col space-y-1.5 max-w-1/3">
      <span class="text-2xl">{{ studentInfo.firstName }} {{ studentInfo.lastName }}</span>
      <span class="text-md"> It’s that pleasure.. ...of feeling like a god If it’s for that.. ..I can lose... everything else.. I could even destroy Blue Lock. </span>
    </div>

    <div class="flex flex-col ml-auto items-end space-y-3 mr-4">
      <span class="text-gray-500">16/07/2025</span>
      <img src="../assets/newTab.svg" class="w-5 h-5" alt="new-tab-icon"/>
      <div class="flex flex-row mt-auto space-x-6">
        <button class="btn shadow-none bg-[#1F7AB9] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7">Accept</button>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7">Reject</button>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>