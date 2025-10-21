<template>
  <div class="w-full bg-gray-100">
    <!-- Cover -->
    <div class="relative">
      <img
        :src="company.cover"
        alt="cover"
        class="w-full h-60 object-cover"
        @error="(e)=>e.target.src='https://images.unsplash.com/photo-1500530855697-b586d89ba3ee?q=80&w=1600&auto=format&fit=crop'"
      />
      <h1
        class="absolute inset-0 flex items-center justify-center text-white font-extrabold text-4xl md:text-5xl drop-shadow-lg"
      >
        {{ company.name }}
      </h1>
    </div>

    <!-- Body -->
    <div class="max-w-[1400px] mx-auto px-4 md:px-8 py-6 grid grid-cols-1 lg:grid-cols-12 gap-6">
      <!-- LEFT: Company panel -->
      <aside class="lg:col-span-5">
        <div class="bg-white rounded-xl shadow-md p-6">
          <!-- Header row -->
          <div class="flex items-center gap-5">
            <div class="avatar">
              <div class="w-36 h-36 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2 overflow-hidden">
                <img :src="company.avatar" alt="logo"
                     @error="(e)=>e.target.src='https://via.placeholder.com/256x256.png?text=Logo'"/>
              </div>
            </div>
            <div class="flex-1">
              <h2 class="text-2xl md:text-3xl font-extrabold text-gray-900">
                {{ company.name }}
              </h2>
              <div class="mt-3 flex flex-wrap gap-2">
                <span
                  v-for="tag in company.tags"
                  :key="tag"
                  class="px-3 py-1 rounded-full bg-gray-100 text-gray-700 text-sm font-medium border"
                >
                  {{ tag }}
                </span>
              </div>
            </div>
          </div>

          <hr class="my-5" />

          <!-- About -->
          <section>
            <h3 class="font-semibold text-gray-800 mb-1">Future Foundation</h3>
            <p class="text-sm leading-relaxed text-gray-700">
              Future Foundation is dedicated to empowering innovation, fostering talent, and creating opportunities
              that shape a better tomorrow. We collaborate with organizations, educators, and communities to drive
              impactful projects that prepare individuals for the challenges of the modern world.
            </p>
          </section>

          <hr class="my-5" />

          <!-- Contact -->
          <section>
            <h4 class="font-semibold text-gray-800 mb-3">Contact Us</h4>
            <ul class="space-y-1 text-sm text-gray-700">
              <li class="flex items-start gap-2">
                <span class="material-icons text-base">location_on</span>
                <span><b>Address:</b> {{ company.contact.address || '—' }}</span>
              </li>
              <li class="flex items-start gap-2">
                <span class="material-icons text-base">call</span>
                <span><b>Phone:</b> {{ company.contact.phone || '—' }}</span>
              </li>
              <li class="flex items-start gap-2">
                <span class="material-icons text-base">mail</span>
                <span><b>Email:</b> {{ company.contact.email || '—' }}</span>
              </li>
              <li class="flex items-start gap-2">
                <span class="material-icons text-base">public</span>
                <span><b>Website:</b> {{ company.contact.website || '—' }}</span>
              </li>
            </ul>
          </section>
        </div>
      </aside>

      <!-- RIGHT: Jobs column -->
      <main class="lg:col-span-7">
        <div class="bg-white rounded-xl shadow-md p-5">
          <!-- Search -->
          <div
            class="w-full flex items-center gap-3 bg-gray-100 rounded-full px-4 py-2 focus-within:ring-2 focus-within:ring-primary/50"
          >
            <span class="material-icons text-gray-500">search</span>
            <input
              v-model.trim="keyword"
              type="text"
              placeholder="Keywords"
              class="bg-transparent outline-none w-full text-sm md:text-base"
            />
          </div>

          <!-- Tabs -->
          <div class="mt-4 flex items-center gap-4 text-sm">
            <button
              class="font-semibold"
              :class="activeTab==='active' ? 'text-gray-900' : 'text-gray-500'"
              @click="activeTab='active'"
            >Active</button>
            <span class="text-gray-300">|</span>
            <button
              class="font-semibold"
              :class="activeTab==='history' ? 'text-gray-900' : 'text-gray-500'"
              @click="activeTab='history'"
            >History</button>
          </div>

          <!-- Job list -->
          <div class="mt-4 max-h-[65vh] overflow-y-auto pr-1 space-y-4">
            <article
              v-for="job in filteredJobs"
              :key="job.id"
              class="rounded-xl border bg-white shadow-sm"
            >
              <div class="p-4 md:p-5 flex gap-4 md:gap-6">
                <!-- Logo -->
                <div class="flex-shrink-0">
                  <div class="w-16 h-16 rounded-xl overflow-hidden bg-gray-100 grid place-items-center">
                    <img :src="job.logo" alt="company logo" class="w-16 h-16 object-cover"
                         @error="(e)=>e.target.src='https://via.placeholder.com/64x64.png?text=LOGO'"/>
                  </div>
                </div>

                <!-- Middle -->
                <div class="flex-1 min-w-0">
                  <h3 class="font-semibold text-gray-800 truncate">
                    <a href="#" class="hover:underline">
                      {{ job.title }}
                    </a>
                  </h3>
                  <p class="text-sm text-gray-600">{{ job.company }}</p>

                  <div class="mt-2 grid grid-cols-2 sm:flex sm:flex-wrap gap-x-6 gap-y-1 text-sm text-gray-600">
                    <div class="flex items-center gap-1">
                      <span class="material-icons text-base">payments</span>
                      <span>{{ job.salary }}</span>
                    </div>
                    <div class="flex items-center gap-1">
                      <span class="material-icons text-base">school</span>
                      <span>{{ job.experience }}</span>
                    </div>
                    <div class="flex items-center gap-1">
                      <span class="material-icons text-base">place</span>
                      <span>{{ job.location }}</span>
                    </div>
                  </div>
                </div>

                <!-- Right stats -->
                <div class="flex flex-col items-end justify-between text-sm">
                  <div class="text-gray-600">
                    <span class="text-gray-500">Applied:</span>
                    <span class="font-semibold text-gray-800">{{ job.applied }}</span>
                  </div>
                  <div class="text-gray-600">
                    <span class="text-gray-500">New Application:</span>
                    <span class="font-semibold text-gray-800">{{ job.newApplications }}</span>
                  </div>
                </div>
              </div>
            </article>

            <p v-if="filteredJobs.length === 0" class="text-center text-gray-500 py-8">
              No jobs found.
            </p>
          </div>
        </div>
      </main>
    </div>

    <!-- FAB (edit / create) -->
    <button
      class="fixed bottom-6 right-6 btn btn-primary btn-circle shadow-xl"
      @click="onFab"
      aria-label="Create or edit"
      title="Create or edit"
    >
      <span class="material-icons">edit</span>
    </button>

    <!-- Material Icons (safe to keep here; navbar already exists) -->
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet" />
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'

const company = ref({
  name: 'Future Foundation',
  cover:
    'https://images.unsplash.com/photo-1500534314209-a25ddb2bd429?q=80&w=1600&auto=format&fit=crop',
  avatar:
    'https://images.unsplash.com/photo-1520975922284-9f8b06a43c36?q=80&w=800&auto=format&fit=crop',
  tags: ['Tech', 'Aviation', 'Material Science'],
  contact: {
    address: 'Bangkok, Thailand',
    phone: '+66-2-123-4567',
    email: 'hr@future.foundation',
    website: 'future.foundation',
  },
})

const activeTab = ref('active') // 'active' | 'history'
const keyword = ref('')

const jobs = ref([
  {
    id: 1,
    title: 'ACCOUNT MANAGER (Sales Engineer)',
    company: 'TE Connectivity',
    logo:
      'https://logo.clearbit.com/te.com',
    salary: '฿120,000–฿150,000 per month',
    experience: '1–2 years',
    location: 'ออฟิศ, กรุงเทพมหานคร',
    applied: 15,
    newApplications: 4,
    status: 'active',
  },
  {
    id: 2,
    title: 'ACCOUNT MANAGER (Sales Engineer)',
    company: 'TE Connectivity',
    logo:
      'https://logo.clearbit.com/te.com',
    salary: '฿120,000–฿150,000 per month',
    experience: '1–2 years',
    location: 'ออฟิศ, กรุงเทพมหานคร',
    applied: 15,
    newApplications: 4,
    status: 'active',
  },
  {
    id: 3,
    title: 'Mechanical Engineer (R&D)',
    company: 'Future Foundation',
    logo:
      'https://logo.clearbit.com/example.com',
    salary: '฿70,000–฿90,000 per month',
    experience: '0–1 year',
    location: 'Hybrid, Bangkok',
    applied: 42,
    newApplications: 0,
    status: 'history',
  },
])

const filteredJobs = computed(() => {
  const k = keyword.value.toLowerCase()
  return jobs.value
    .filter((j) => (activeTab.value === 'active' ? j.status === 'active' : j.status !== 'active'))
    .filter(
      (j) =>
        !k ||
        j.title.toLowerCase().includes(k) ||
        j.company.toLowerCase().includes(k) ||
        j.location.toLowerCase().includes(k)
    )
})

function onFab() {
  // hook into your modal/router here
  // e.g., router.push('/jobs/new') or open a job editor modal
  alert('Open job editor (wire this to your modal or route).')
}
</script>

<style scoped>
/* Subtle scroll styling for the job list */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
::-webkit-scrollbar-thumb {
  background: rgba(0,0,0,0.15);
  border-radius: 9999px;
}
::-webkit-scrollbar-track {
  background: transparent;
}
</style>
