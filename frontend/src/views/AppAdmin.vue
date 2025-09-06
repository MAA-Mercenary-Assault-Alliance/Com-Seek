<script setup lang="ts">

import CompanyBox from "../components/CompanyBox.vue";
import { TE_Info_Company } from "../components/temp_template";
</script>

<template>
  <div class="flex flex-col w-full">

    <div class="flex bg-[#153243] w-full h-40 px-42 items-center">
      <span class="text-white text-5xl font-bold">Admin Dashboard CHANGE BACKGROUND</span>
    </div>

    <div id="content" class="flex flex-row space-x-10">
      <div id="sidebar" class="flex flex-col w-1/5">
        <div class="flex items-center border-1 h-35">
          <span class="ml-15 text-3xl">Company</span>
        </div>
        <div class="flex items-center border-1 h-35">
          <span class="ml-15 text-3xl">Company</span>
        </div>
      </div>
      <div id="boxes" class="flex flex-col mt-10 space-y-6 w-4/5 pr-42">
        <CompanyBox :company-info="TE_Info_Company"/>
        <CompanyBox :company-info="TE_Info_Company"/>
      </div>
    </div>

  </div>
</template>

<style scoped>

</style>