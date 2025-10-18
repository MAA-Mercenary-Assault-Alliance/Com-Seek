<script setup lang="ts">
import Applicant from "./Applicant.vue";

const prop = defineProps({
  applicants: { type: Array, required: true }
});
</script>

<template>
  <div class="flex flex-col w-1/3">
    <Applicant v-for="applicant in applicants" :key="applicant.ID" :applicant="applicant"/>
  </div>
</template>

<style scoped>

</style>