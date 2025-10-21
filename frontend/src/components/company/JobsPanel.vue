<template>
  <div class="bg-white rounded-xl shadow-md p-6">
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-xl font-semibold">Your Job Posts</h2>

      <div class="flex gap-2">
        <input class="input input-bordered input-sm w-48" placeholder="Keyword" v-model="keyword" />
        <input class="input input-bordered input-sm w-40" placeholder="Location" v-model="location" />
        <select v-model="jobType" class="select select-bordered select-sm w-48">
          <option value="">All Types</option>
          <option>Software & Application Development</option>
          <option>Data & AI</option>
          <option>Cloud & Infrastructure</option>
          <option>Cybersecurity</option>
          <option>Product & Design</option>
          <option>Testing & Quality</option>
          <option>Hardware & Electronics</option>
          <option>Management & Leadership</option>
          <option>IT Support & Operations</option>
        </select>
        <button class="btn btn-sm" @click="applyFilter">Filter</button>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      <!-- List -->
      <div class="md:col-span-1 max-h-[540px] overflow-auto space-y-3 border rounded-lg p-3">
        <div v-if="isLoading" class="text-gray-500">Loading jobs…</div>
        <template v-else>
          <JobFullEmpty v-if="jobs.length === 0" />
          <div
            v-for="job in filteredJobs"
            :key="job.ID"
            class="p-3 rounded-lg border cursor-pointer hover:bg-gray-50"
            :class="modelValue && modelValue.ID === job.ID ? 'ring-2 ring-primary' : ''"
            @click="$emit('update:modelValue', job)"
          >
            <div class="font-semibold">{{ job.Title }}</div>
            <div class="text-sm text-gray-500">{{ job.Location }} • {{ job.EmploymentStatus }}</div>
            <div class="text-xs mt-1">
              <span class="px-2 py-0.5 rounded bg-gray-100">{{ job.JobType }}</span>
            </div>
          </div>
        </template>
      </div>

      <!-- Detail -->
      <div class="md:col-span-2">
        <div v-if="modelValue" class="border rounded-lg p-5">
          <div class="flex items-start justify-between">
            <div>
              <h3 class="text-lg font-bold">{{ modelValue.Title }}</h3>
              <div class="text-sm text-gray-500">
                {{ modelValue.Location }} • {{ modelValue.EmploymentStatus }}
              </div>
            </div>
            <div class="flex items-center gap-2">
              <span
                class="text-xs px-2 py-1 rounded"
                :class="modelValue.Approved ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700'"
              >
                {{ modelValue.Approved ? 'Approved' : 'Pending' }}
              </span>
              <span
                v-if="!modelValue.Visibility"
                class="text-xs px-2 py-1 rounded bg-gray-200 text-gray-700"
              >
                Hidden
              </span>
            </div>
          </div>

          <div class="mt-3 whitespace-pre-wrap">{{ modelValue.Description }}</div>

          <div class="mt-4 text-sm text-gray-600">
            <div>Experience: {{ modelValue.MinExperience }}–{{ modelValue.MaxExperience }} yrs</div>
            <div>
              Salary:
              {{ modelValue.MinSalary?.toLocaleString?.() || modelValue.MinSalary }} –
              {{ modelValue.MaxSalary?.toLocaleString?.() || modelValue.MaxSalary }}
            </div>
            <div class="mt-2">
              Type: <span class="px-2 py-0.5 rounded bg-gray-100">{{ modelValue.JobType }}</span>
            </div>
          </div>
        </div>

        <div v-else class="text-gray-500">Select a job to see details.</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import JobFullEmpty from '@/components/JobFullEmpty.vue'

const props = defineProps({
  jobs: { type: Array, default: () => [] },
  isLoading: { type: Boolean, default: false },
  modelValue: { type: Object, default: null }
})
defineEmits(['update:modelValue'])

const keyword = ref('')
const jobType = ref('')
const location = ref('')

const filteredJobs = computed(() => {
  return props.jobs.filter(j => {
    const kw = keyword.value?.toLowerCase() || ''
    const kwOk = !kw || j.Title?.toLowerCase().includes(kw) || j.Description?.toLowerCase().includes(kw)
    const jtOk = !jobType.value || j.JobType === jobType.value
    const locOk = !location.value || j.Location?.toLowerCase().includes(location.value.toLowerCase())
    return kwOk && jtOk && locOk
  })
})

watch([keyword, jobType, location], () => {
  console.log('Filters:', { keyword: keyword.value, jobType: jobType.value, location: location.value })
})

function applyFilter() {
  console.log('Applied local filters.')
}
</script>
