<template>
  <div class="divide-y divide-gray-300">
    <div class="py-3">
      <p class="font-semibold text-gray-700 text-l">Profile Type
        <span v-if="profile.isAlum" class="badge bg-[#EAF6EC] text-[#0A3B1F] border border-[#56A45C] px-4">Alumni</span>
        <span v-else class="badge bg-gray-100 text-gray-700 border border-gray-300 px-4">Student</span>
      </p>
    </div>
    <div class="py-3">
      <p class="font-semibold text-gray-700 text-l">Email Address</p>
      <p class="text-gray-600 break-all">{{ profile.email }}</p>
    </div>
    <div class="py-3">
      <p class="font-semibold text-gray-700 text-l">First Name</p>
      <p class="text-gray-600">{{ profile.firstName }}</p>
    </div>
    <div class="py-3">
      <p class="font-semibold text-gray-700 text-l">Last Name</p>
      <p class="text-gray-600">{{ profile.lastName }}</p>
    </div>

    <div class="mt-4 text-gray-700">
      <p class="font-semibold text-gray-700 text-l">Description</p>

      <!-- Markdown-rendered description -->
      <div
        v-if="profile.description && profile.description.length"
        class="prose prose-sm max-w-none mt-1 text-gray-700"
        v-html="renderedDescription"
      />

      <p v-else class="italic text-gray-400 mt-1">
        No description provided.
      </p>
    </div>

  </div>
</template>

<script>
import { marked } from 'marked'

export default {
  name: 'ProfileDetails',
  props: {
    profile: { type: Object, required: true },
    canEdit: { type: Boolean, default: false },
  },
  computed: {
    fullName() { return this.profile?.name || ''; },
    description() { return this.profile?.description || ''; },

    renderedDescription() {
      const md = this.profile?.description || ''
      return md ? marked.parse(md, { breaks: true }) : ''
    },

    socials() { return this.profile?.socials || {}; },
    isAlum() { return !!this.profile?.isAlum; },
    approved() { return !!this.profile?.approved; },
  },
};
</script>
