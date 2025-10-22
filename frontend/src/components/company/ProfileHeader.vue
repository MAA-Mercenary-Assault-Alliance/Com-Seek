<template>
  <section class="relative">
    <!-- Cover -->
    <img :src="coverSrc" alt="company cover" class="w-full h-52 object-cover" />

    <!-- Header card -->
    <div class="w-full mx-auto bg-white rounded-xl shadow-md -mt-12 p-6 relative">
      <div class="px-4 md:px-30">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <!-- Logo -->
            <div class="avatar">
              <div
                class="w-30 h-30 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2 overflow-hidden bg-white"
              >
                <img :src="logoSrc" alt="company logo" class="object-cover w-full h-full" />
              </div>
            </div>

            <!-- Identity -->
            <div>
              <h1 class="text-2xl md:text-4xl font-bold">{{ profile?.Name || 'Company Name' }}</h1>

              <!-- Website -->
              <div class="mt-1 text-sm text-gray-600">
                <a
                  v-if="profile?.Website"
                  :href="profile.Website"
                  target="_blank"
                  class="link link-primary break-all"
                >
                  {{ profile.Website }}
                </a>
                <span v-else class="text-gray-400">website not set</span>
              </div>

              <!-- Tags -->
              <div v-if="tagsList.length" class="flex flex-wrap gap-2 mt-2">
                <span
                  v-for="(tag, index) in tagsList"
                  :key="index"
                  class="badge badge-outline text-sm font-medium border-green-600 text-green-700 bg-green-50"
                >
                  {{ tag }}
                </span>
              </div>
              <div v-else class="text-xs text-gray-400 mt-2">No tags available</div>
            </div>
          </div>

          <!-- Approval badge -->
          <span
            class="px-2 py-1 text-s rounded-full"
            :class="
              profile?.Approved
                ? 'bg-green-100 text-green-700'
                : 'bg-yellow-100 text-yellow-700'
            "
          >
            {{ profile?.Approved ? 'Approved' : 'Pending Review' }}
          </span>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  profile: { type: Object, default: null },
})

const coverSrc = computed(
  () => props.profile?.Cover || 'https://picsum.photos/1200/240?blur=2&random=1'
)
const logoSrc = computed(
  () => props.profile?.Logo || 'https://placehold.co/128x128?text=Logo'
)

const tagsList = computed(() => {
  const raw = props.profile?.Tags || props.profile?.tags || ''
  if (Array.isArray(raw)) return raw
  if (typeof raw === 'string' && raw.trim() !== '')
    return raw.split(',').map(t => t.trim()).filter(Boolean)
  return []
})
</script>
