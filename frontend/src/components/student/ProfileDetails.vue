<template>
  <div class="divide-y divide-gray-300">
    <div class="py-3">
      <p class="font-semibold text-gray-700">Username</p>
      <p class="text-gray-600 break-all">{{ profile.username }}</p>
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
  </div>
</template>

<script>
export default {
  name: "ProfileDetails",
  props: {
    profile: { type: Object, required: true },
  },
};
</script>
