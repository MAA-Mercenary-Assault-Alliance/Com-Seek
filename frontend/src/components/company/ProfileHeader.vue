<template>
  <section class="relative">
    <!-- Cover -->
    <div class="relative">
      <img :src="coverSrc" alt="company cover" class="w-full h-52 object-cover" />
      <div class="absolute inset-x-0 bottom-0 h-10 bg-gradient-to-t from-black/15 to-transparent"></div>
    </div>

    <!-- Header card -->
    <div class="w-full mx-auto bg-white rounded-xl shadow-md -mt-12 p-6 relative">
      <div class="px-4 md:px-35">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <!-- Logo -->
            <div class="avatar">
              <div
                class="w-30 h-30 rounded-full ring ring-offset-2 ring-[#44B15B] ring-offset-white overflow-hidden bg-white"
              >
                <img :src="logoSrc" alt="company logo" class="object-cover w-full h-full" />
              </div>
            </div>

            <!-- Identity -->
            <div>
              <h1 class="text-2xl md:text-4xl font-bold text-[#0A3B1F]">
                {{ profile?.Name || 'Company Name' }}
              </h1>

              <!-- Website -->
              <div class="mt-1 text-sm">
                <a
                  v-if="profile?.Website"
                  :href="profile.Website"
                  target="_blank"
                  class="break-all text-[#0A3B1F] hover:text-[#44B15B] underline underline-offset-2"
                >
                  {{ profile.Website }}
                </a>
                <span v-else class="text-gray-400">website not set</span>
              </div>

              <!-- Tags -->
              <div v-if="tagsList.length" class="flex flex-wrap gap-2 mt-2">
                <button
                  v-for="(tag, index) in tagsList"
                  :key="index"
                  type="button"
                  class="px-2 py-0.5 rounded-full border text-sm font-medium
                         border-[#56A45C] text-[#0A3B1F] bg-[#EAF6EC]
                         hover:bg-[#DDF0E1] transition"
                  @click="$emit('tag-click', tag)"
                >
                  {{ tag }}
                </button>
              </div>
              <div v-else class="text-xs text-gray-400 mt-2">No tags available</div>
            </div>
          </div>

          <!-- Approval badge -->
          <span
            class="px-2 py-1 text-xs rounded-full border"
            :class="profile?.Approved
              ? 'bg-[#EAF6EC] text-[#0A3B1F] border-[#56A45C]'
              : 'bg-yellow-100 text-yellow-700 border-yellow-300'"
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
