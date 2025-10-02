<script setup lang="ts">
import { ref } from 'vue'
const keyword = ref('')

</script>

<template>
  <div class="flex min-w-70 max-w-1/3 p-2">
    <label class="input input-lg rounded-2xl placeholder-black w-full">
      <input v-model="keyword" type="search" class="grow pl-3 placeholder-black" placeholder="Keywords" />
      <img src="../assets/search.svg" class="w-7 h-7 rounded-2xl" alt="company-logo"/>
    </label>
  </div>
</template>

<style scoped>

</style>