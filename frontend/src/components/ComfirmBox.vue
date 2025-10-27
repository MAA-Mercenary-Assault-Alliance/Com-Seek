<script setup lang="ts">
// @ts-nocheck
import { ref } from 'vue'
const { companyName } = defineProps({ companyName: String });
const emit = defineEmits(["reject", "notReject"]);

const dialog = ref(null)

function open() {
  dialog.value.showModal();
}
function yes() {
  emit("reject");
  dialog.value.close();
}
function no() {
  emit("notReject");
  dialog.value.close();
}

defineExpose({ open });
</script>

<template>
  <dialog ref="dialog" class="modal">
    <div class="modal-box" @click.stop>
      <h3 class="text-lg font-bold">Are you sure?</h3>
      <p class="py-4">You are going to reject {{ companyName }}</p>
      <div class="flex flex-row justify-end space-x-4 mt-5">
        <button class="btn shadow-none bg-[#1F7AB9] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click="yes">Yes</button>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click="no">No</button>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop" @click.stop>
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped>

</style>