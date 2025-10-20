<script setup lang="ts">
import Applicant from "./Applicant.vue";

const props = defineProps({
  applicants: { type: Array, required: true }
});

const emit = defineEmits(["refresh"]);

function onEmit() {
  emit("refresh");
}
</script>

<template>
  <div class="flex flex-col w-1/3">
    <Applicant v-for="applicant in applicants" :key="applicant.id" :applicant="applicant" @refresh="onEmit"/>
  </div>
</template>

<style scoped>

</style>