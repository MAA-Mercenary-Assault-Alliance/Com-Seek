<script setup lang="ts">
import {JobTemplate} from './temp_template'
import {marked} from "marked";
import {onMounted, ref, watch} from "vue";
import {api} from "../../api/client.js";
import DateConverter from './dateConverter';
import {useRouter} from 'vue-router'
import SuccessBox from "@/components/SuccessBox.vue";

const props = defineProps<{
  jobInfo: JobTemplate
}>();

const router = useRouter();
const successBox = ref(null);

let desc_html = marked(props.jobInfo?.Description || "")
const date = ref("")

watch(() => props.jobInfo, (newJobInfo) => {
  console.log('JobFull received jobInfo:', newJobInfo)
  desc_html = marked(newJobInfo.Description || "")
})

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
    console.log("This job id:", props.jobInfo.ID)
    const res = await api.post("/job/apply", {
      job_id: Number(props.jobInfo.ID)
    })
    console.log("Applied to job:", res.data)
    successBox.value.open()

    setTimeout(() => {
      successBox.value.close();
    }, 2000);

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

function goToCompany(companyID: number) {
  router.push({ name: 'CompanyProfile', params: {id: companyID}})
  // TODO: integrate with the real company profile
}

onMounted(() => {
  date.value = DateConverter(props.jobInfo?.CreatedAt || "");
})
</script>

<template>
  <div id="job full" class="flex w-full px-25 py-10 flex-col rounded-2xl box-shadow bg-white sticky">

    <div id="title-box" class="flex row items-center mt-3">
      <img src="../assets/company.jpg" class="w-20 h-20 rounded-2xl" alt="company-logo"/>
      <div id="title" class="flex flex-col ml-7 space-y-4">
        <span class="text-2xl" @click="goToCompany">{{ jobInfo.Company?.Name }}</span>
        <span class="underline">{{ jobInfo.Title }}</span>
      </div>
      <button v-if="role === 'student'" class="btn shadow-none bg-[#44b15b] border-0 h-12 rounded-2xl text-white text-xl font-extralight ml-auto mt-6" @click="applyJob">Apply Now</button>
    </div>
    <span class="absolute top-5 right-5 text-gray-500">{{ date }}</span>

    <span class="mt-10 text-lg">Basic Requirements</span>
    <div id="requirements-box" class="grid grid-cols-2 gap-4 w-full max-w-4xl mt-6 space-y-2">
      <div class="requirements-div">
        <img src="../assets/location.svg" class="requirements-icon" alt="location-icon"/>
        <span>{{ jobInfo.Location }}</span>
      </div>
      <div class="requirements-div">
        <img src="../assets/money.svg" class="requirements-icon" alt="money-icon"/>
        <span>฿{{ jobInfo.MinSalary }}-{{ jobInfo.MaxSalary}}</span>
      </div>
      <div class="requirements-div">
        <img src="../assets/time.svg" class="requirements-icon" alt="time-icon"/>
        <span>{{ jobInfo.JobType }}</span>
      </div>
      <div class="requirements-div">
        <img src="../assets/cap.svg" class="requirements-icon" alt="cap-icon"/>
        <span>{{ jobInfo.MinExperience }}-{{ jobInfo.MaxExperience }} years</span>
      </div>
    </div>

    <span class="mt-12 text-lg">Full Job Description</span>
    <div class="prose">
      <div v-html="desc_html"></div>
    </div>

  </div>

  <success-box
      ref="successBox"
      :message="'Successfully applied for the job!'"
  ></success-box>

</template>

<style scoped>

</style>