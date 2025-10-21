<template>
  <section class="relative">
    <!-- Cover -->
    <img :src="coverSrc" alt="company cover" class="w-full h-52 object-cover" />

    <!-- Header card -->
    <div class="w-full mx-auto bg-white rounded-xl shadow-md -mt-12 p-6 relative">
      <div class="px-4 md:px-12">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-4">
            <!-- Logo -->
            <div class="avatar">
              <div class="w-24 h-24 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2 overflow-hidden bg-white">
                <img :src="logoSrc" alt="company logo" class="object-cover w-full h-full" />
              </div>
            </div>

            <!-- Identity -->
            <div>
              <h1 class="text-2xl md:text-3xl font-bold">{{ profile?.Name || 'Company Name' }}</h1>
              <div class="mt-1 text-sm text-gray-600">
                <a v-if="profile?.Website" :href="profile.Website" target="_blank" class="link link-primary break-all">
                  {{ profile.Website }}
                </a>
                <span v-else class="text-gray-400">website not set</span>
              </div>
            </div>
          </div>

          <span
            class="px-2 py-1 text-xs rounded-full"
            :class="profile?.Approved ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700'"
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
const props = defineProps({ profile: { type: Object, default: null } })

const coverSrc = computed(() =>
  props.profile?.Cover || 'https://picsum.photos/1200/240?blur=2&random=1'
)
const logoSrc = computed(() =>
  props.profile?.Logo || 'https://placehold.co/128x128?text=Logo'
)
</script>
