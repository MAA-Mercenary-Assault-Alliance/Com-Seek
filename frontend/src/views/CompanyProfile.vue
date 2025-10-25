<template>
  <div class="w-full bg-[F2F6FC]" v-if="loadedOnce">
    <ProfileHeader :profile="profile" />

    <section class="w-full bg-[#0a3b1f] py-5"></section>

    <div
      class="absolute bottom-4 right-6 text-green-900 opacity-10 pointer-events-none select-none"
    >
      <i class="fas fa-building text-[10rem]"></i>
    </div>

    <div class="max-w-6xl mx-auto py-6 px-4">
      <ProfileDetails
        :profile="profile"
        :is-loading="isLoadingProfile"
        :is-saving="isSaving"
        :can-edit="canEdit"
        :alert="alert"
        @save="saveProfile"
      />
    </div>

    <div v-if="canEdit" class="max-w-6xl mx-auto pb-12 px-4">
      <JobsPanel 
      :jobs="jobs" 
      :is-loading="isLoadingJobs" 
      v-model="selectedJob"
      :company-name="profile?.Name"
      />
    </div>
  </div>

  <div v-else class="min-h-[60vh] flex items-center justify-center text-gray-500">
    Loading…
  </div>
</template>

<script setup>
import { api } from '../../api/client.js'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import ProfileHeader from '@/components/company/ProfileHeader.vue'
import ProfileDetails from '@/components/company/ProfileDetails.vue'
import JobsPanel from '@/components/company/JobsPanel.vue'

const route = useRoute()
const profile = ref(null)
const isLoadingProfile = ref(true)
const isSaving = ref(false)
const canEdit = ref(false)
const alert = ref({ type: '', message: '' })
const loadedOnce = ref(false)

const jobs = ref([])
const isLoadingJobs = ref(true)
const selectedJob = ref(null)

async function getCompanyProfile() {
  try {
    isLoadingProfile.value = true
    const id = route.params.id
    const url = id ? `/company/${id}` : '/company'
    const res = await api.get(url)
    profile.value = res.data.profile
    canEdit.value = !id
  } catch (err) {
    console.error('Error fetching company profile:', err)
    alert.value = { type: 'error', message: 'Unable to load company profile.' }
  } finally {
    isLoadingProfile.value = false
    loadedOnce.value = true
  }
}

async function saveProfile(payload) {
  try {
    isSaving.value = true
    alert.value = { type: '', message: '' }
    const res = await api.patch('/company', payload)
    console.log('Profile updated:', res.data)
    alert.value = { type: 'success', message: 'Profile updated ✅' }
    await getCompanyProfile()
  } catch (err) {
    console.error('Error saving profile:', err)
    alert.value = { type: 'error', message: 'Failed to update profile.' }
  } finally {
    isSaving.value = false
  }
}

async function getCompanyJobs() {
  try {
    isLoadingJobs.value = true
    if (!route.params.id) {
      const res = await api.get('/company/jobs')
      jobs.value = res.data.jobs || []
      selectedJob.value = jobs.value[0] || null
    } else {
      jobs.value = []
      selectedJob.value = null
    }
  } catch (err) {
    console.error('Error fetching jobs:', err)
  } finally {
    isLoadingJobs.value = false
  }
}

onMounted(async () => {
  await getCompanyProfile()
  await getCompanyJobs()
})
</script>
