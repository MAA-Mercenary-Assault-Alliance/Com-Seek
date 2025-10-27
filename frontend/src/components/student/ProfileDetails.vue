<template>
  <div class="divide-y divide-gray-300">
    <div class="py-3">
      <p class="font-semibold text-gray-700">Profile Type
        <span v-if="profile.isAlum" class="badge bg-[#EAF6EC] text-[#0A3B1F] border border-[#56A45C] px-4">Alumni</span>
        <span v-else class="badge bg-gray-100 text-gray-700 border border-gray-300 px-4">Student</span>
      </p>
    </div>
    <div class="py-3">
      <p class="font-semibold text-gray-700">Email Address</p>
      <p class="text-gray-600 break-all">{{ profile.email }}</p>
    </div>
    <div class="py-3">
      <p class="font-semibold text-gray-700">First Name</p>
      <p class="text-gray-600">{{ profile.firstName }}</p>
    </div>
    <div class="py-3">
      <p class="font-semibold text-gray-700">Last Name</p>
      <p class="text-gray-600">{{ profile.lastName }}</p>
    </div>
    <div class="mt-4 text-gray-700">
      <p class="font-semibold text-gray-700">Description</p>
      <p v-if="profile.description && profile.description.length">{{ profile.description }}</p>
      <p v-else class="italic text-gray-400">No description provided.</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ProfileDetails',
  props: {
    profile: { type: Object, required: true },
    canEdit: { type: Boolean, default: false },
  },
  computed: {
    fullName() { return this.profile?.name || ''; },
    description() { return this.profile?.description || ''; },
    socials() { return this.profile?.socials || {}; },
    isAlum() { return !!this.profile?.isAlum; },
    approved() { return !!this.profile?.approved; },
  },
};
</script>

