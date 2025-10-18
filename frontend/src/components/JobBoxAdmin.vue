<script setup lang="ts">
import {ref} from 'vue'
import {JobTemplate} from './temp_template'
import {api} from "../../api/client.js";
import DateConverter from "./dateConverter";
import ConfirmBox from "./ComfirmBox.vue";
const props = defineProps<{
  jobInfo: JobTemplate
}>();

const confirmBox = ref(null);
const emit = defineEmits(["refresh"]);

async function rejectJob() {
  try {
    await api.patch(`/admin/review-job/${props.jobInfo.ID}`, {
      approved: false
    });
    console.log("Job rejected successfully");
    emit("refresh");
  } catch (error) {
    console.error("Error rejecting Job:", error);
  }
}

async function acceptJob() {
  try {
    await api.patch(`/admin/review-job/${props.jobInfo.ID}`, {
      approved: true
    });
    console.log("Job accepted successfully");
    emit("refresh");
  } catch (error) {
    console.error("Error accepting Job:", error);
  }
}

const date = DateConverter(props.jobInfo.CreatedAt);
</script>

<template>
  <div id="job-box" class="flex relative rounded-2xl flex-row p-4 h-[180px] pr-4 pb-4 space-x-7 box-shadow bg-white">
    <div class="flex items-center justify-center flex-shrink-0">
      <img src="../assets/company.jpg" class="company-admin-logo" alt="company-logo"/>
    </div>
    <div id="job-box-content" class="flex mr-2 flex-col space-y-1.5">
      <span class="underline">{{ jobInfo.Title }}</span>
      <span class="mb-2">{{ jobInfo.Company.Name }}</span>

      <div class="in-line-stat">
        <img src="../assets/money.svg" class="in-line-icon" alt="money-icon"/>
        <span>฿{{ jobInfo.MinSalary }}-{{ jobInfo.MaxSalary}}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/cap.svg" class="in-line-icon" alt="cap-icon"/>
        <span>{{ jobInfo.MinExperience }}-{{ jobInfo.MaxExperience }} years</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/location.svg" class="in-line-icon" alt="location-icon"/>
        <span>{{ jobInfo.Location }}</span>
      </div>
    </div>

    <div class="flex flex-col ml-auto items-end space-y-3 mr-4">
      <span class="text-gray-500">{{ date }}</span>
      <img src="../assets/newTab.svg" class="w-5 h-5" alt="new-tab-icon"/>
      <div class="flex flex-row mt-auto space-x-6">
        <button class="btn shadow-none bg-[#1F7AB9] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click="acceptJob">Accept</button>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click="confirmBox.open()">Reject</button>
      </div>
    </div>

    <confirm-box
        ref="confirmBox"
        :company-name="jobInfo.Title"
        @reject="rejectJob"
    ></confirm-box>

  </div>
</template>

<style scoped>

</style>