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

const fileInput = ref(null)
const selectedFile = ref(null)

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
    close()
    if (error.response.data.error == "job application already exists") {
      alert("You have already applied to this job.")
    } else {
      alert("Error applying to job. Please try again later.")
    }
  }
}

function handleFileUpload(event) {
  const file = event.target.files[0] // take the first selected file
  if (file) {
    selectedFile.value = file
    console.log('Selected file:', file)
  }
}
</script>

<template>
  <dialog ref="dialog" class="modal">
    <div class="modal-box">
      <div class="flex flex-col justify-center items-center p-10 space-y-5">
        <p class="py-4 text-xl">Upload your CV/Resume here</p>
        <div class="relative inline-flex items-center">
          <!-- Hidden file input -->
          <input
              ref="fileInput"
              type="file"
              accept=".pdf"
              @change="handleFileUpload"
              class="absolute w-0 h-0 opacity-0 overflow-hidden"
          />

          <!-- Styled button -->
          <button
              type="button"
              @click="fileInput.click()"
              class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md"
          >
            Upload PDF
          </button>

          <!-- Display selected file name -->
          <span v-if="selectedFile" class="ml-2 text-gray-700">{{ selectedFile.name }}</span>
        </div>

        <button class="btn shadow-none border-0 h-10 w-30 rounded-md text-white text-md font-extralight px-7 disabled:bg-gray-200 bg-lighter"
                :disabled="!selectedFile"
                @click="applyJob">
          Apply
        </button>
      </div>

    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped>

</style>