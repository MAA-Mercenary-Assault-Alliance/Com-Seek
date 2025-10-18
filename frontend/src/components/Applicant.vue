<script setup lang="ts">
import DateConverter from "../components/dateConverter";
import {ref} from "vue";
import { api } from '../../api/client.js';
import ConfirmBox from "./ComfirmBox.vue";

const props = defineProps({
  applicant: { type: Object, required: true }
});

console.log("Applicant Data: ", props.applicant)

const studentStatus = () => {
  if (props.applicant.is_alum) {
    return "Alumni";
  }
    return "Student";
};

const confirmBox = ref(null);
const emit = defineEmits(["refresh"]);

async function deleteApplicant() {
  try {
    await api.patch(`/delete-applicant/somethingsomething`, {
      // TODO: replace path with actual api
      approved: false
    });
    console.log("Applicant deleted successfully");
    emit("refresh");
  } catch (error) {
    console.error("Error deleting applicant:", error);
  }
}

const date = DateConverter(props.applicant.created_at);
</script>

<template>
  <div class="flex w-full">
    <div id="job-box" class="flex relative rounded-2xl flex-row p-4 pb-4 space-x-5 box-shadow bg-white w-full">
      <img src="../assets/company.jpg" class="w-20 h-20 rounded-full" alt="company-logo"/>
      <div class="flex flex-col">
        <span>{{ applicant.first_name }} {{ applicant.last_name }}</span>
        <span>{{ studentStatus() }}</span>
      </div>
      <div class="flex flex-col ml-auto">
        <span class="text-gray-500 ml-auto">{{ date }}</span>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-9 w-25 rounded-4xl text-white text-md font-extralight mt-auto">Delete</button>
      </div>
    </div>

    <confirm-box
        ref="confirmBox"
        :company-name="applicant.first_name"
        @reject="deleteApplicant"
    ></confirm-box>
  </div>
</template>

<style scoped>

</style>