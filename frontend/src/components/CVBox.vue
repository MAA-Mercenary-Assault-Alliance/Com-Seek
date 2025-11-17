<script setup lang="ts">
// @ts-nocheck
import { onMounted, ref } from "vue";
import { renderRecaptcha } from "../services/reCAPTCHA";
import { api } from "../../api/client";

const reCAPTCHASiteKey = import.meta.env.VITE_RECAPTCHA_SITE_KEY;

const props = defineProps({ jobID: Number });

const dialog = ref(null);

function open() {
  try {
    // eslint-disable-next-line no-undef
    grecaptcha.reset();
    // eslint-disable-next-line no-unused-vars
  } catch (error) {
    console.log("Rendering reCAPTCHA");
  }
  renderRecaptcha(this.$refs.recaptchaContainer, reCAPTCHASiteKey);
  dialog.value.show();
}

function close() {
  dialog.value.close();
}

defineExpose({ open, close });
const emit = defineEmits("success");

const role = ref(localStorage.getItem("role"));
onMounted(() => {
  window.addEventListener("storage", () => {
    role.value = localStorage.getItem("role");
  });
});

const fileInput = ref(null);
const selectedFile = ref(null);
const loading = ref(false);

async function applyJob() {
  if (role.value !== "student") {
    alert("Only students can apply to jobs.");
    return;
  }

  if (!selectedFile.value) {
    alert("Please upload your CV/Resume (PDF) before applying.");
    return;
  }

  try {
    console.log("Applying for job ID:", props.jobID);

    // loading screen
    loading.value = true;

    const fd = new FormData();

    fd.append("job_id", String(props.jobID));
    fd.append("file", selectedFile.value);
    // eslint-disable-next-line no-undef
    fd.append("recaptcha_response", grecaptcha.getResponse());

    const res = await api.post("/job/apply", fd);

    loading.value = false;

    console.log("Applied to job:", res.data);
    emit("success");
    close();
    selectedFile.value = null;
    if (fileInput.value) {
      fileInput.value.value = "";
    }
    return;
  } catch (error) {
    loading.value = false;
    close();
    if (error.response) {
      console.error("Error applying to job:", error.response.data);
      if (error.response.data.error == "job application already exists") {
        alert("You have already applied to this job.");
      } else {
        alert("Error applying to job. Please try again later.");
      }
    } else {
      console.error("Error applying to job:", error);
    }
  }
}

function handleFileUpload(event) {
  const file = event.target.files[0]; // take the first selected file
  if (file) {
    if (file.type !== "application/pdf") {
      alert("Please upload a PDF file.");
      event.target.value = ""; // Clear file input
      selectedFile.value = null;
      return;
    }

    const MAX_SIZE_MB = 10;

    if (file.size > MAX_SIZE_MB * 1024 * 1024) {
      alert(`File size exceeds the limit of ${MAX_SIZE_MB}MB.`);
      event.target.value = "";
      selectedFile.value = null;
      return;
    }

    selectedFile.value = file;
    console.log("Selected file:", file);
  }
}
</script>

<template>
  <dialog ref="dialog" class="modal">
    <div class="modal-box">
      <div
        v-if="loading"
        class="flex flex-col justify-center items-center p-10 space-y-5"
      >
        <svg
          class="animate-spin h-10 w-10 text-blue-600"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
        >
          <circle
            class="opacity-25"
            cx="12"
            cy="12"
            r="10"
            stroke="currentColor"
            stroke-width="4"
          ></circle>
          <path
            class="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
          ></path>
        </svg>
        <p class="text-lg">Submitting your application...</p>
      </div>
      <div
        v-if="!loading"
        class="flex flex-col justify-center items-center p-10 space-y-5"
      >
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
            <!-- @click is as if we click on the hidden fileInput above -->
            Upload PDF
          </button>

          <!-- Display selected file name -->
          <span v-if="selectedFile" class="ml-2 text-gray-700">{{
            selectedFile.name
          }}</span>
        </div>

        <div
          class="g-recaptcha flex justify-center"
          ref="recaptchaContainer"
          :data-sitekey="reCAPTCHASiteKey"
        ></div>

        <button
          class="btn shadow-none border-0 h-10 w-30 rounded-md text-white text-md font-extralight px-7 disabled:bg-gray-200 bg-lighter"
          :disabled="!selectedFile"
          @click="applyJob"
        >
          Apply
        </button>
      </div>
    </div>
    <form method="dialog" class="modal-backdrop">
      <button>close</button>
    </form>
  </dialog>
</template>

<style scoped></style>
