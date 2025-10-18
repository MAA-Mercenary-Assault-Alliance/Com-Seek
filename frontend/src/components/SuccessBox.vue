<script setup lang="ts">
// @ts-nocheck
import { ref } from 'vue'
const props = defineProps({ message: String });

const dialog = ref(null)

function open() {
  dialog.value.showModal();
  setTimeout(() => {
    dialog.value.close();
  }, 2000);
}

defineExpose({ open });
</script>

<template>
  <dialog ref="dialog" class="modal">
    <div class="modal-box">
      <img src="../assets/checkmark.svg"/>
      <h3 class="text-lg font-bold">Are you sure?</h3>
      <p class="py-4">{{ props.message }}</p>

    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped>

</style>