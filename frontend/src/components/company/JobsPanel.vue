<template>
  <div class="bg-white rounded-xl shadow-md p-6 h-[20vh] md:h-[70vh] flex flex-col">
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-xl font-semibold text-[#0A3B1F]">{{ companyName || 'Company Name' }}'s Job Posts</h2>

      <div class="flex gap-2">
        <input
          class="input input-bordered input-sm w-48 focus:outline-none focus:border-[#44B15B] focus:ring-2 focus:ring-[#44B15B]/30"
          placeholder="Keyword"
          v-model="keyword"
        />
        <input
          class="input input-bordered input-sm w-40 focus:outline-none focus:border-[#44B15B] focus:ring-2 focus:ring-[#44B15B]/30"
          placeholder="Location"
          v-model="location"
        />
        <select
          v-model="jobType"
          class="select select-bordered select-sm w-48 focus:outline-none focus:border-[#44B15B] focus:ring-2 focus:ring-[#44B15B]/30"
        >
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
        <button
          class="btn btn-sm bg-[#44B15B] border-[#44B15B] hover:bg-[#56A45C] hover:border-[#56A45C] text-white"
          @click="applyFilter"
        >
          Filter
        </button>
      </div>
    </div>

    <!-- Make content fill remaining height; allow children to shrink -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 flex-1 min-h-0">
      <!-- List -->
      <div class="md:col-span-1 h-full overflow-auto space-y-3 rounded-lg p-3 shadow-sm bg-background">
        <div v-if="isLoading" class="text-gray-500">Loading jobs…</div>
        <template v-else>
          <JobFullEmpty v-if="jobs.length === 0" />
          <div
            v-for="job in filteredJobs"
            :key="job.ID"
            class="p-3 rounded-lg shadow-sm cursor-pointer transition-all duration-150 bg-white hover:shadow-md hover:bg-[#EAF6EC]"
            :class="modelValue && modelValue.ID === job.ID ? 'ring-2 ring-[#44B15B] bg-[#EAF6EC]' : ''"
            @click="$emit('update:modelValue', job)"
          >
            <div class="font-semibold text-[#0A3B1F]">{{ job.Title }}</div>
            <div class="text-sm text-gray-600">{{ job.Location }} • {{ job.EmploymentStatus }}</div>
            <div class="text-xs mt-1">
              <span class="px-2 py-0.5 rounded border border-[#56A45C] text-[#0A3B1F] bg-[#EAF6EC]">
                {{ job.JobType }}
              </span>
            </div>
          </div>
        </template>
      </div>

      <!-- Detail: fills to end and scrolls without overflowing parent -->
      <div class="md:col-span-2 h-full min-h-0 flex">
        <div
          v-if="jobs.length === 0"
          class="flex flex-col flex-1 items-center justify-center -translate-y-10"
        >
          <img src="../../assets/leaf.svg" class="w-30" alt="leaf" />
          <span class="text-2xl text-gray-500 mt-10">
            No jobs posted yet
          </span>
        </div>

        <div
          v-else-if="modelValue"
          class="rounded-lg p-5 bg-white shadow-md flex-1 overflow-auto"
        >
          <div class="flex items-start justify-between">
            <div>
              <h3 class="text-lg font-bold text-[#0A3B1F]">
                {{ modelValue.Title }}
              </h3>
              <div class="text-sm text-gray-600">
                {{ modelValue.Location }} • {{ modelValue.EmploymentStatus }}
              </div>
            </div>
            <div class="flex items-center gap-2">
              <span
                class="text-xs px-2 py-1 rounded"
                :class="modelValue.Approved
                  ? 'bg-[#EAF6EC] text-[#0A3B1F] border border-[#56A45C]'
                  : 'bg-yellow-100 text-yellow-700 border border-yellow-300'"
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

          <!-- Markdown-rendered description -->
          <div
            class="prose prose-sm max-w-none mt-3 text-gray-800"
            v-html="renderedMarkdown"
          />

          <div class="mt-4 text-sm text-gray-700">
            <div>
              <span class="font-medium text-[#0A3B1F]">Experience:</span>
              {{ modelValue.MinExperience }}–{{ modelValue.MaxExperience }} yrs
            </div>
            <div class="mt-1">
              <span class="font-medium text-[#0A3B1F]">Salary:</span>
              {{ modelValue.MinSalary?.toLocaleString?.() || modelValue.MinSalary }} –
              {{ modelValue.MaxSalary?.toLocaleString?.() || modelValue.MaxSalary }}
            </div>
            <div class="mt-2">
              <span class="font-medium text-[#0A3B1F]">Type:</span>
              <span class="ml-2 px-2 py-0.5 rounded border border-[#56A45C] text-[#0A3B1F] bg-[#EAF6EC]">
                {{ modelValue.JobType }}
              </span>
            </div>
          </div>
        </div>

        <!-- ℹ️ Jobs exist, but none selected -->
        <div
          v-else
          class="flex flex-1 justify-center items-center"
        >
          <p class="text-gray-500 text-lg">
            Select a job to see details.
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { marked } from 'marked' // ← add this
import JobFullEmpty from '@/components/JobFullEmpty.vue'

const props = defineProps({
  jobs: { type: Array, default: () => [] },
  isLoading: { type: Boolean, default: false },
  modelValue: { type: Object, default: null },
  companyName: { type: String, default: 'Company Name' },
  hR: { type: Boolean, default: false }
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

const renderedMarkdown = computed(() => {
  const md = props.modelValue?.Description || ''
  return md ? marked.parse(md, { breaks: true }) : ''
})
</script>
