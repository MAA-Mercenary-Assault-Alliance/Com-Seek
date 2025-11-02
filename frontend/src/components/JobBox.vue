<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {JobTemplate} from './temp_template'
const props = defineProps<{
  jobInfo: JobTemplate,
  hR: boolean
}>();

onMounted(() => {
  console.log('JobBox received jobInfo:', props.jobInfo)
})

const applied = ref(0)
const new_app = ref(0)
const day = ref(0)

function daysAgo(dateString: string): number {
  if (!dateString) return 0;

  const date = new Date(dateString);
  const now = new Date();

  // calculate difference in milliseconds
  const diffMs = now.getTime() - date.getTime();

  // convert to days
  return Math.floor(diffMs / (1000 * 60 * 60 * 24));
}

onMounted(() => {
  day.value = daysAgo(props.jobInfo.CreatedAt)
  // applied.value = props.jobInfo.Applications ? props.jobInfo.Applications.length : 0
  // new_app.value = props.jobInfo.Applications ? props.jobInfo.Applications.filter((app) => {
  //   return daysAgo(app.CreatedAt) <= 7
  // }).length : 0
  // TODO: Fix applied and new_app
})

</script>

<template>
  <div id="job-box" class="flex relative rounded-2xl flex-row p-4 pb-7 space-x-5 box-shadow bg-white hover:bg-gray-100" @click="hR && $emit('click')"  :class="{ 'cursor-pointer': hR }">

    <img src="../assets/company.jpg" class="w-20 h-20 rounded-2xl" alt="company-logo"/>
    <div id="job-box-content" class="flex mr-2 flex-col space-y-1.5">
      <span class="underline">{{ jobInfo.Title }}</span>
      <router-link v-if="!hR" :to="{ name: 'CompanyProfilePublic', params: { id: Number(jobInfo.Company?.UserID) }}" class="mb-2">{{ jobInfo.Company?.Name }}</router-link>
      <span v-else class="mb-2">{{ jobInfo.Company?.Name }}</span>

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
    <span class="absolute bottom-3 right-5 text-gray-500">{{ day }} days ago</span>

    <div v-if=hR class="flex flex-col ml-auto mr-6 items-end space-y-1.5">
      <div class="flex flex-row gap-2 pb-3">
        <span
            class="text-xs px-2 py-1 rounded"
            :class="jobInfo.Approved
                  ? 'bg-[#EAF6EC] text-[#0A3B1F] border border-[#56A45C]'
                  : 'bg-yellow-100 text-yellow-700 border border-yellow-300'"
        >{{ jobInfo.Approved ? 'Approved' : 'Pending' }}</span>
        <span
            v-if="!jobInfo.Visibility"
            class="text-xs px-2 py-1 rounded bg-gray-200 text-gray-700"
        >Hidden</span>
      </div>
      <span class="text-gray-500">Applied: {{applied}}</span>
      <span :class="new_app == 0 ? 'text-gray-500' : 'text-green-500'">New: {{new_app}}</span>
    </div>
  </div>
</template>

<style scoped>

</style>