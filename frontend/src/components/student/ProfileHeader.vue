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
      <h1 class="text-5xl font-bold text-gray-800 truncate">
        {{ profile.first_name || profile.firstName }} {{ profile.last_name || profile.lastName }}
      </h1>

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
        id="edit-profile"
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
    }
  }

};
</script>
