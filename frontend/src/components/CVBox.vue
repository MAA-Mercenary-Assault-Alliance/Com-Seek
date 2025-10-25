<script setup lang="ts">
// @ts-nocheck
import {onMounted, ref, defineEmits} from 'vue'
import {api} from "../../api/client";
const props = defineProps({ jobID: Number });

const dialog = ref(null)

function open() {
  dialog.value.showModal();
}

function close() {
  dialog.value.close();
}

defineExpose({ open, close });
const emit = defineEmits("success")

const role = ref(localStorage.getItem('role'));
onMounted(() => {
  window.addEventListener('storage', () => {
    role.value = localStorage.getItem('role');
  });
});

async function applyJob() {
  if (role.value !== 'student') {
    alert("Only students can apply to jobs.")
    return
  }
  try {
    console.log("This job id:", props.jobID)
    const res = await api.post("/job/apply", {
      job_id: Number(props.jobID)
    })
    console.log("Applied to job:", res.data)
    emit("success")

    return
  } catch (error) {
    console.error("Error applying to job:", error.response.data)
    if (error.response.data.error == "job application already exists") {
      alert("You have already applied to this job.")
    } else {
      alert("Error applying to job. Please try again later.")
    }
  }
}
</script>

<template>
  <dialog ref="dialog" class="modal">
    <div class="modal-box">
      <div class="flex flex-col justify-center items-center p-10">
        <p class="py-4 text-xl">Upload your CV/Resume here</p>
      </div>

    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped>

</style>