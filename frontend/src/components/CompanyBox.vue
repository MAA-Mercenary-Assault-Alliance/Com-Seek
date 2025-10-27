<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {CompanyTemplate} from './temp_template'
import ConfirmBox from "./ComfirmBox.vue";
import { api } from "../../api/client.js";
import DateConverter from './dateConverter';
import {useRouter} from "vue-router";

const props = defineProps<{
  companyInfo: CompanyTemplate,
}>();

console.log("Company Info:", props.companyInfo);

const confirmBox = ref(null);
const emit = defineEmits(["refresh"]);
const date = ref()
const router = useRouter()

async function rejectCompany() {
  try {
    await api.patch(`/admin/review-company/${props.companyInfo.UserID}`, {
      approved: false
    });
    console.log("Company rejected successfully");
    emit("refresh");
  } catch (error) {
    console.error("Error rejecting company:", error);
  }
}

async function acceptCompany() {
  try {
    await api.patch(`/admin/review-company/${props.companyInfo.UserID}`, {
      approved: true
    });
    console.log("Company accepted successfully");
    emit("refresh");
  } catch (error) {
    console.error("Error accepting company:", error);
  }
}

function goToCompany(newtab: boolean = false) {
  const route = { name: 'CompanyProfilePublic', params: {id: props.companyInfo.UserID}}
  if (newtab) {
    const url = router.resolve(route).href
    window.open(url, '_blank')
    return
  }
  router.push(route)
}


onMounted(() => {
  date.value = DateConverter(props.companyInfo.CreatedAt);
})

</script>

<template>
  <div id="company-box" class="flex relative rounded-2xl flex-row p-4 space-x-7 h-[180px] box-shadow bg-white" @click="goToCompany(false)">
    <div class="flex items-center justify-center flex-shrink-0">
      <img src="../assets/company.jpg" class="company-admin-logo" alt="company-logo"/>
    </div>
    <div id="job-box-content" class="flex mr-2 flex-col space-y-1.5">
      <span class="text-2xl">{{ companyInfo.Name }}</span>

      <div class="in-line-stat mt-1">
        <img src="../assets/phone.svg" class="in-line-icon" alt="phone-icon"/>
        <span>{{ companyInfo.ContactNumber }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/email.svg" class="in-line-icon" alt="email-icon"/>
        <span>{{ companyInfo.ContactEmail }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/internet.svg" class="in-line-icon" alt="internet-icon"/>
        <span>{{ companyInfo.Website }}</span>
      </div>
      <div class="in-line-stat">
        <img src="../assets/location.svg" class="in-line-icon" alt="location-icon"/>
        <span>{{ companyInfo.Location }}</span>
      </div>
    </div>

    <div class="flex flex-col ml-auto items-end space-y-3 mr-4">
      <span class="text-gray-500">{{ date }}</span>
      <img src="../assets/newTab.svg" class="w-5 h-5 cursor-pointer" alt="new-tab-icon" @click.stop="goToCompany(true)"/>
      <div class="flex flex-row mt-auto space-x-6">
        <button class="btn shadow-none bg-[#1F7AB9] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click="acceptCompany">Accept</button>
        <button class="btn shadow-none bg-[#9A0000] border-0 h-8 rounded-4xl text-white text-md font-extralight px-7" @click="confirmBox.open()">Reject</button>
      </div>
    </div>

    <confirm-box
        ref="confirmBox"
        :company-name="companyInfo.Name"
        @reject="rejectCompany"
    ></confirm-box>

  </div>
</template>

<style scoped>

</style>