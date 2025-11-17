<script setup lang="ts">
// @ts-nocheck
import { ref } from 'vue'
import JobFull from "@/components/JobFull.vue";
const props = defineProps({ jobInfo: Object });
const emit = defineEmits(["accept", "reject"]);

const dialog = ref(null)

function open() {
  dialog.value.showModal();
}

function close() {
  dialog.value.close();
}

defineExpose({ open, close });
</script>

<template>
  <dialog ref="dialog" class="modal">
    <div class="modal-box p-0 max-w-none w-2/3 my-20">
        <job-full :job-info="jobInfo"/>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped>

</style>