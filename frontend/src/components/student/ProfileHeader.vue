<template>
  <div class="flex items-center space-x-6">
    <!-- Avatar -->
    <div class="avatar">
      <div class="w-40 h-40 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2 overflow-hidden">
        <img 
        :src="profile.avatar || '/images/avatar.png'" 
        @error="onImgError"
        alt="avatar" 
        class="object-cover w-28 h-28" />
      </div>
    </div>

    <!-- Identity -->
    <div class="min-w-0">
      <div class="flex items-center gap-3">
        <h1 class="text-5xl font-bold text-gray-800 truncate flex-1 min-w-0">
          {{ profile.first_name || profile.firstName }} {{ profile.last_name || profile.lastName }}
        </h1>
        <div
          v-if="profile.transcript || profile.cv"
          class="flex items-center gap-3 text-gray-600 text-2xl"
        >
          <a
            v-if="profile.transcript"
            :href="profile.transcript"
            target="_blank"
            rel="noopener"
            :download="fileDownloadName('transcript')"
            class="group relative inline-flex items-center justify-center w-10 h-10 rounded-full bg-gray-200 hover:bg-gray-300 transition"
            aria-label="View transcript"
            title="View transcript"
          >
            <i class="fa-regular fa-file-lines"></i>
            <span
              class="pointer-events-none absolute -bottom-8 left-1/2 -translate-x-1/2 whitespace-nowrap rounded-md bg-gray-800 px-2 py-1 text-xs text-white opacity-0 group-hover:opacity-100 transition"
            >
              {{ fileDownloadName('transcript') }}
            </span>
          </a>
          <a
            v-if="profile.cv"
            :href="profile.cv"
            target="_blank"
            rel="noopener"
            :download="fileDownloadName('cv')"
            class="group relative inline-flex items-center justify-center w-10 h-10 rounded-full bg-gray-200 hover:bg-gray-300 transition"
            aria-label="View CV"
            title="View CV"
          >
            <i class="fa-regular fa-file"></i>
            <span
              class="pointer-events-none absolute -bottom-8 left-1/2 -translate-x-1/2 whitespace-nowrap rounded-md bg-gray-800 px-2 py-1 text-xs text-white opacity-0 group-hover:opacity-100 transition"
            >
              {{ fileDownloadName('cv') }}
            </span>
          </a>
        </div>
      </div>

      <!-- Socials (supports either nested socials.* or top-level fields from API) -->
      <div class="flex items-center space-x-4 mt-1 text-3xl text-gray-700">
        <a
          v-if="(profile.socials && profile.socials.facebook) || profile.facebook"
          :href="(profile.socials && profile.socials.facebook) || profile.facebook"
          target="_blank" rel="noopener" aria-label="Facebook" class="hover:opacity-80 transition" title="Facebook"
        >
          <i class="fa-brands fa-facebook"></i>
        </a>

        <a
          v-if="(profile.socials && profile.socials.twitter) || profile.twitter"
          :href="(profile.socials && profile.socials.twitter) || profile.twitter"
          target="_blank" rel="noopener" aria-label="X (Twitter)" class="hover:opacity-80 transition" title="X (Twitter)"
        >
          <i class="fa-brands fa-x-twitter"></i>
        </a>

        <a
          v-if="(profile.socials && profile.socials.instagram) || profile.instagram"
          :href="(profile.socials && profile.socials.instagram) || profile.instagram"
          target="_blank" rel="noopener" aria-label="Instagram" class="hover:opacity-80 transition" title="Instagram"
        >
          <i class="fa-brands fa-instagram"></i>
        </a>

        <a
          v-if="(profile.socials && profile.socials.github) || profile.github"
          :href="(profile.socials && profile.socials.github) || profile.github"
          target="_blank" rel="noopener" aria-label="GitHub" class="hover:opacity-80 transition" title="GitHub"
        >
          <i class="fa-brands fa-github"></i>
        </a>

        <a
          v-if="(profile.socials && profile.socials.linkedin) || profile.linkedin"
          :href="(profile.socials && profile.socials.linkedin) || profile.linkedin"
          target="_blank" rel="noopener" aria-label="LinkedIn" class="hover:opacity-80 transition" title="LinkedIn"
        >
          <i class="fa-brands fa-linkedin"></i>
        </a>
      </div>
    </div>

    <!-- Spacer pushes gear right -->
    <div class="ml-auto" v-if="canEdit">
      <button
        class="w-10 h-10 flex items-center justify-center rounded-full hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-[#44B15B] transition"
        aria-label="Edit profile"
        title="Edit profile"
        @click="handleEdit"
      >
        <span class="material-icons text-3xl text-gray-600 hover:rotate-90 transition">
          settings
        </span>
      </button>
    </div>
  </div>
</template>


<script>
export default {
  name: 'ProfileHeader',
  props: {
    profile: { type: Object, required: true },
    canEdit: { type: Boolean, default: false },
  },
  emits: ['edit'],
  methods: {
    handleEdit() {
      if (this.canEdit) this.$emit('edit');
    },
    onImgError(event) {
      event.target.src = '/images/avatar.png';
    },
    fileDownloadName(type) {
      const first = this.profile.first_name || this.profile.firstName || 'Student';
      const last = this.profile.last_name || this.profile.lastName || 'Profile';
      const suffix = type === 'cv' ? 'CV' : 'Transcript';
      return `${first}_${last}_${suffix}`.replace(/\s+/g, '_');
    }
  }
};
</script>
