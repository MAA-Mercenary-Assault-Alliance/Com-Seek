<script setup lang="ts">
import {StudentTemplate} from '../services/model_template'
import {onMounted, ref} from "vue";
import { api } from "../../api/client.js";
import ConfirmBox from "./ComfirmBox.vue";
import DateConverter from "../services/dateConverter";
import { useRouter } from 'vue-router';
import {getFileUrl} from "@/services/fileUpload";
const props = defineProps<{
  studentInfo: StudentTemplate,
}>();

const confirmBox = ref(null);
const emit = defineEmits(["refresh"]);
const date = ref()
const router = useRouter();
const DEFAULT_AVATAR = "/images/avatar.png";
const student_logo_url = ref("");

async function rejectStudent() {
  try {
    await api.patch(`/admin/review-student/${props.studentInfo.UserID}`, {
      approved: false
    });
    console.log("Student rejected successfully");
    emit("refresh");
  } catch (error) {
    console.error("Error rejecting Student:", error);
  }
}

async function acceptStudent() {
  try {
    await api.patch(`/admin/review-student/${props.studentInfo.UserID}`, {
      approved: true
    });
    console.log("Student accepted successfully");
    emit("refresh");
  } catch (error) {
    console.error("Error accepting Student:", error);
  }
}

function gotoStudent(newTab: boolean = false) {
  const route = { name: 'StudentProfilePublic', params: {id: props.studentInfo.UserID}}
  if (newTab) {
    const url = router.resolve(route).href
    window.open(url, '_blank')
    return
  }
  router.push(route)
}

onMounted(() => {
  date.value = DateConverter(props.studentInfo.CreatedAt);
  student_logo_url.value = getFileUrl(props.studentInfo.profile_image_id, DEFAULT_AVATAR)
})
</script>

<template>
  <div id="student-box" class="flex relative rounded-2xl flex-row p-4 space-x-7 box-shadow bg-white" @click="gotoStudent(false)">
    <div class="items-center flex-shrink-0">
      <img :src=student_logo_url class="w-30 h-30 rounded-full ml-2" alt="company-logo"/>
    </div>
    <div id="student-box-content" class="flex mr-2 flex-col space-y-1.5 max-w-1/3">
      <span class="text-2xl">{{ studentInfo.FirstName }} {{ studentInfo.LastName }}</span>
      <span class="text-md"> {{ studentInfo.Description }} </span>
    </div>

    <div class="flex flex-col ml-auto items-end space-y-3 mr-4">
      <span class="text-gray-500">{{ date }}</span>
      <img src="../assets/newTab.svg" class="w-5 h-5 cursor-pointer" alt="new-tab-icon" @click.stop="gotoStudent(true)"/>
      <div class="flex flex-row mt-auto space-x-6">
        <button class="btn shadow-none bg-[#1F7AB9] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click.stop="acceptStudent">Accept</button>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click="confirmBox.open()">Reject</button>
      </div>
    </div>

    <confirm-box
        ref="confirmBox"
        :company-name="studentInfo.FirstName"
        @reject="rejectStudent"
    ></confirm-box>

  </div>
</template>

<style scoped>

</style>