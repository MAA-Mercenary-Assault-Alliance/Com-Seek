<template>
  <div class="min-h-screen bg-gray-100" v-if="profile">
    <!-- Top bar -->
    <!-- <div class="navbar bg-base-200 px-6 shadow-md">
      <div class="flex-1">
        <a class="text-xl font-bold flex items-center">
          <img src="@/assets/logo.png" alt="logo" class="h-8 mr-2" />
          Com-Seek
        </a>
      </div>
      <div class="flex-none space-x-6">
        <router-link to="/" class="hover:underline">Home</router-link>
        <router-link to="/profile" class="hover:underline">Profile</router-link>
        <router-link to="/logout" class="hover:underline">Logout</router-link>
      </div>
    </div> -->

    <!-- Cover image -->
    <div class="relative">
      <img
        :src="profile.cover"
        alt="cover"
        class="w-full h-60 object-cover"
      />
      <button class="absolute right-4 bottom-4 btn btn-circle btn-outline">
        <i class="fas fa-cog"></i>
      </button>
    </div>

    <!-- Profile Info -->
    <div
      class="w-11/12 md:w-8/12 mx-auto bg-white rounded-xl shadow-md -mt-16 p-6 relative"
    >
      <div class="flex items-center space-x-6">
        <!-- Profile avatar -->
        <div class="avatar">
          <div
            class="w-28 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2"
          >
            <img :src="profile.avatar" />
          </div>
        </div>

        <!-- User Info -->
        <div>
          <h1 class="text-2xl font-bold text-gray-800">
            {{ profile.firstName }} {{ profile.lastName }}
          </h1>
          <div class="flex items-center space-x-4 mt-1 text-xl text-primary">
            <a v-if="profile.socials.facebook" :href="profile.socials.facebook" target="_blank"><i class="fab fa-facebook"></i></a>
            <a v-if="profile.socials.twitter" :href="profile.socials.twitter" target="_blank"><i class="fab fa-x-twitter"></i></a>
            <a v-if="profile.socials.instagram" :href="profile.socials.instagram" target="_blank"><i class="fab fa-instagram"></i></a>
            <a v-if="profile.socials.github" :href="profile.socials.github" target="_blank"><i class="fab fa-github"></i></a>
          </div>
          <p class="text-sm text-gray-500 mt-2">
            {{ formatDate(profile.joinDate) }}
          </p>
          <p class="text-gray-600 text-sm">{{ profile.bio }}</p>
        </div>
      </div>
    </div>

    <!-- Details Section -->
    <div
      class="w-11/12 md:w-8/12 mx-auto mt-6 bg-white rounded-xl shadow-md p-6"
    >
      <div class="divide-y divide-gray-300">
        <div class="py-3">
          <p class="font-semibold text-gray-700">Username</p>
          <p class="text-gray-600">{{ profile.username }}</p>
        </div>
        <div class="py-3">
          <p class="font-semibold text-gray-700">Email Address</p>
          <p class="text-gray-600">{{ profile.email }}</p>
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
    </div>
  </div>
</template>

<script>
import { fetchStudentProfile } from '../services/profileStubApi';

export default {
  name: "StudentProfile",
  data() {
    return {
      profile: null,
    };
  },
  async created() {
    this.profile = await fetchStudentProfile();
  },
  methods: {
    formatDate(dateStr) {
      const options = { year: "numeric", month: "long", day: "numeric" };
      return new Date(dateStr).toLocaleDateString(undefined, options);
    },
  },
};
</script>
