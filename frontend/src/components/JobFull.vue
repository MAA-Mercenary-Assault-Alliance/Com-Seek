<script setup lang="ts">
import {JobTemplate} from './temp_template'
import {marked} from "marked";
import {computed, onMounted, ref, watch} from "vue";
import {api} from "../../api/client.js";
import DateConverter from './dateConverter';
import {useRouter} from 'vue-router'
import SuccessBox from "@/components/SuccessBox.vue";
import CVBox from "@/components/CVBox.vue";
import ConfirmBoxGeneral from "@/components/ConfirmBoxGeneral.vue";

const props = defineProps<{
  jobInfo: JobTemplate,
  hR: boolean
}>();

const router = useRouter();
const successBox = ref(null);
const successBoxTerminate = ref(null);
const cvBox = ref(null);
const confirmBox = ref(null);

let desc_html = marked(props.jobInfo?.Description || "")
const date = computed(() => {
  return DateConverter(props.jobInfo?.CreatedAt || "")
})

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

async function terminateJob() {
  try {
    const res = await api.patch(`/job/${props.jobInfo.ID}`, {
      Visibility: false
    })
    console.log("Terminated Job:", props.jobInfo.Title)
    successBoxTerminate.value.open()
    setTimeout(() => {
      router.push({ name: "HRDashboard" })
    }, 2000) // 2 seconds

    return
  } catch (error) {
    alert("Error terminating job. Please try again later.")
    console.error("Error terminating job:", error.response.data)
  }
}

// function goToCompany(companyID: number) {
//   router.push({ name: 'CompanyProfile', params: {id: companyID}})
//
// }

function goToEdit() {
  router.push({ name: 'EditJob', params: {id: props.jobInfo.ID}})
}

</script>

<template>
  <div id="job full" class="flex w-full px-25 py-10 flex-col rounded-2xl box-shadow bg-white sticky">

    <div id="title-box" class="flex row items-center mt-3">
      <img src="../assets/company.jpg" class="w-20 h-20 rounded-2xl" alt="company-logo"/>
      <div id="title" class="flex flex-col ml-7 space-y-4">
        <router-link :to="{ name: 'CompanyProfilePublic', params: { id: Number(jobInfo.Company?.UserID) }}" class="text-2xl" >{{ jobInfo.Company?.Name }}</router-link>
        <span class="underline">{{ jobInfo.Title }}</span>
      </div>
      <div class="flex flex-col ml-auto">
        <button v-if="role === 'student'" class="btn shadow-none bg-[#44b15b] border-0 h-12 rounded-2xl text-white text-xl font-extralight ml-auto mt-6" @click="cvBox.open()">Apply Now</button>
        <button v-if="role === 'company' && hR" class="btn shadow-none bg-[#DB0000] border-0 h-18 w-45 rounded-4xl text-white text-2xl font-extralight ml-auto mt-6" @click="confirmBox.open()">Terminate</button>
        <button v-if="role === 'company' && hR" class="btn shadow-none bg-[#44b15b] border-0 h-18 w-45 rounded-4xl text-white text-2xl font-extralight ml-auto mt-6" @click="goToEdit">Edit Job</button>
      </div>
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
        <span>{{ jobInfo.EmploymentStatus }}</span>
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

  <success-box
      ref="successBoxTerminate"
      :message="'Successfully terminated the job'"
  ></success-box>

  <c-v-box
    ref="cvBox"
    :jobID="jobInfo.ID"
    @success="successBox.open()"
  ></c-v-box>

  <confirm-box-general
    ref="confirmBox"
    :message="'Are you sure you want to terminate this job posting? This action cannot be undone.'"
    @accept="terminateJob"
  ></confirm-box-general>

</template>

<style scoped>

</style>