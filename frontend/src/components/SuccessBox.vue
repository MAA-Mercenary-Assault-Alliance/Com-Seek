<script setup lang="ts">
// @ts-nocheck
import { ref } from 'vue'
const props = defineProps({ message: String });

const dialog = ref(null)

function open() {
  dialog.value.showModal();
  setTimeout(() => {
    close();
  }, 2000);
}

function close() {
  dialog.value.close();
}

defineExpose({ open, close });
</script>

<template>
  <dialog ref="dialog" class="modal">
    <div class="modal-box">
      <div class="flex flex-col justify-center items-center p-10">
        <img src="../assets/checkmark.svg" class="w-30"/>
        <p class="py-4 text-xl">{{ props.message }}</p>
      </div>

    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped>

</style>