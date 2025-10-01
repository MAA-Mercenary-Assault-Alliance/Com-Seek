<template>
  <div class="flex items-center space-x-6">
    <!-- Avatar -->
    <div class="avatar">
      <div class="w-28 h-28 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2 overflow-hidden">
        <img :src="profile.avatar" alt="avatar" class="object-cover w-28 h-28" />
      </div>
    </div>

    <!-- Identity -->
    <div class="min-w-0">
      <h1 class="text-5xl font-bold text-gray-800 truncate">
        {{ profile.firstName }} {{ profile.lastName }}
      </h1>

        <!-- Socials -->
        <div class="flex items-center space-x-4 mt-1 text-3xl text-gray-700">
        <a
            v-if="profile.socials?.facebook"
            :href="profile.socials.facebook"
            target="_blank"
            rel="noopener"
            aria-label="Facebook"
            class="hover:opacity-80 transition"
            title="Facebook"
        >
            <i class="fa-brands fa-facebook"></i>
        </a>

        <a
            v-if="profile.socials?.twitter"
            :href="profile.socials.twitter"
            target="_blank"
            rel="noopener"
            aria-label="X (Twitter)"
            class="hover:opacity-80 transition"
            title="X (Twitter)"
        >
            <!-- FA6 name -->
            <i class="fa-brands fa-x-twitter"></i>
        </a>

        <a
            v-if="profile.socials?.instagram"
            :href="profile.socials.instagram"
            target="_blank"
            rel="noopener"
            aria-label="Instagram"
            class="hover:opacity-80 transition"
            title="Instagram"
        >
            <i class="fa-brands fa-instagram"></i>
        </a>

        <a
            v-if="profile.socials?.github"
            :href="profile.socials.github"
            target="_blank"
            rel="noopener"
            aria-label="GitHub"
            class="hover:opacity-80 transition"
            title="GitHub"
        >
            <i class="fa-brands fa-github"></i>
        </a>
        </div>


      <p class="text-sm text-gray-500 mt-2">
        {{ formatDate(profile.joinDate) }}
      </p>
      <p class="text-gray-600 text-sm whitespace-pre-line">
        {{ profile.bio }}
      </p>
    </div>

    <!-- Spacer pushes gear right -->
    <div class="ml-auto">
      <button
        class="p-2 rounded-full hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-primary transition"
        aria-label="Edit profile"
        title="Edit profile"
        @click="$emit('edit')"
      >
        <span class="material-icons text-3xl text-gray-600 hover:rotate-90 transition">settings</span>
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProfileHeader",
  props: {
    profile: { type: Object, required: true },
  },
  methods: {
    formatDate(dateStr) {
      const options = { year: "numeric", month: "long", day: "numeric" };
      return new Date(dateStr).toLocaleDateString(undefined, options);
    },
  },
};
</script>
