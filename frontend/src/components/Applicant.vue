<script setup lang="ts">
import DateConverter from "../services/dateConverter";
import {onMounted, ref} from "vue";
import { api } from '../../api/client.js';
import ConfirmBox from "./ComfirmBox.vue";
import {useRouter} from "vue-router";
import {getFileUrl} from "@/services/fileUpload";

const props = defineProps({
  applicant: { type: Object, required: true }
});
const router = useRouter();

console.log("Applicant Data: ", props.applicant)

const studentStatus = () => {
  if (props.applicant.is_alum) {
    return "Alumni";
  }
    return "Student";
};

const confirmBox = ref(null);
const emit = defineEmits(["refresh"]);
const date = ref("");
const DEFAULT_AVATAR = "/images/avatar.png";
const student_logo_url = ref("");

async function deleteApplicant() {
  try {
    await api.delete(`/job/application/${props.applicant.id}`);
    console.log("Applicant deleted successfully");
    emit("refresh");
  } catch (error) {
    console.error("Error deleting applicant:", error.response.data);
  }
}

function gotoStudent(studentID) {
  router.push({ name: 'StudentProfilePublic', params: {id: studentID}})
}

onMounted(() => {
  date.value = DateConverter(props.applicant.created_at);
  student_logo_url.value = getFileUrl(props.applicant.profile_image_id, DEFAULT_AVATAR)
})

</script>

<template>
  <div class="flex w-full cursor-pointer" @click="gotoStudent(applicant.student_id)">
    <div id="applicant-box" class="flex relative rounded-2xl flex-row p-4 pb-4 space-x-5 box-shadow bg-white w-full">
      <img :src=student_logo_url class="w-20 h-20 rounded-full" alt="company-logo"/>
      <div class="flex flex-col">
        <span>{{ applicant.first_name }} {{ applicant.last_name }}</span>
        <span class="text-gray-500">{{ studentStatus() }}</span>
      </div>
      <div class="flex flex-col ml-auto">
        <span class="text-gray-500 ml-auto">{{ date }}</span>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-9 w-25 rounded-4xl text-white text-md font-extralight mt-auto" @click.stop="confirmBox.open">Delete</button>
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