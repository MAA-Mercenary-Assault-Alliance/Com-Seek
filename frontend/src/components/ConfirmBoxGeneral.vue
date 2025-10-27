<script setup lang="ts">
// @ts-nocheck
import { ref } from 'vue'
const props = defineProps({ message: String });
const emit = defineEmits(["accept", "reject"]);

const dialog = ref(null)

function open() {
  dialog.value.showModal();
}
function accept() {
  emit("accept");
  dialog.value.close();
}
function reject() {
  emit("reject");
  dialog.value.close();
}

defineExpose({ open });
</script>

<template>
  <dialog ref="dialog" class="modal">
    <div class="modal-box">
      <h3 class="text-lg font-bold">Are you sure?</h3>
      <p class="py-4">{{ props.message }}</p>
      <div class="flex flex-row justify-end space-x-4 mt-5">
        <button class="btn shadow-none bg-[#1F7AB9] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click="accept">Yes</button>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click="reject">No</button>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped>

</style>