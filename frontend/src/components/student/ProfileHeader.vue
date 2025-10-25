<template>
  <div class="flex items-center justify-between">
    <div class="flex items-center gap-4">
      <div class="avatar">
        <div class="w-28 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2">
          <img :src="avatar" alt="avatar" />
        </div>
      </div>
      <div>
        <h1 class="text-3xl font-bold text-gray-900">
          {{ fullName }}
        </h1>
        <p class="text-gray-600">{{ profile?.email }}</p>
      </div>
    </div>

    <button
      v-if="canEdit"
      class="btn btn-outline btn-sm"
      @click="$emit('edit')"
    >
      <span class="material-icons text-base mr-1">edit</span>
      Edit Profile
    </button>
  </div>
</template>

<script>
export default {
  name: 'ProfileHeader',
  props: {
    profile: { type: Object, required: true },
    avatar: { type: String, required: true },
    canEdit: { type: Boolean, default: false }
  },
  computed: {
    fullName() {
      const f = this.profile?.first_name || ''
      const l = this.profile?.last_name || ''
      return `${f} ${l}`.trim() || 'Student'
    }
  }
}
</script>
