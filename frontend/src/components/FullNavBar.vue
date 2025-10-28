<script setup lang="ts">
import { ref, onMounted } from "vue";

const role = ref(localStorage.getItem("role"));
onMounted(() => {
  window.addEventListener("storage", () => {
    role.value = localStorage.getItem("role");
  });
});
</script>

<template>
  <div
    class="flex bg-background justify-between items-center top-0 w-full min-h-18 p-4 text-black px-42 text-xl flex-wrap"
  >
    <!-- Left section: logo + title -->
    <ul class="flex space-x-4 items-center">
      <li>
        <router-link to="/home">
          <img src="../assets/logo.png" alt="logo" class="w-12" />
        </router-link>
      </li>

      <li class="font-[righteous] text-3xl">
        <router-link
          to="/home"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          Com-Seek
        </router-link>
      </li>

      <li
        v-if="role == 'admin'"
        class="flex ml-4 box-shadow text-xl bg-[#56A45C] rounded-4xl w-48 h-9 text-center items-center justify-center"
      >
        Admin User
      </li>
    </ul>

    <!-- Right section: navigation links -->
    <ul class="flex space-x-15 items-center">
      <!-- Home -->
      <li>
        <router-link
          to="/home"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          Home
        </router-link>
      </li>

      <!-- Company Dashboard -->
      <li v-if="role == 'company'">
        <router-link
          to="/hr-dashboard"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          HR Dashboard
        </router-link>
      </li>

      <!-- Admin Dashboard -->
      <li v-if="role == 'admin'">
        <router-link
          to="/admin"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          Admin Dashboard
        </router-link>
      </li>

      <!-- Login / Register Group -->
      <li v-if="role == null" class="flex space-x-3 items-center">
        <router-link
          to="/login"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          Login
        </router-link>

        <img
          src="../assets/vertical-line.png"
          alt="vertical line"
          class="block h-12 w-15 -ml-6 -mr-4"
        />

        <router-link
          to="/register"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          Register
        </router-link>
      </li>

      <!-- For Company -->
      <li v-if="role == null">
        <router-link
          to="/register-company"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          For Company
        </router-link>
      </li>

      <!-- Student Profile -->
      <li v-if="role == 'student'">
        <router-link
          to="/student-profile"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          Profile
        </router-link>
      </li>

      <!-- Company Profile -->
      <li v-if="role == 'company'">
        <router-link
          to="/company-profile"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          Profile
        </router-link>
      </li>

      <!-- Logout -->
      <li v-if="role">
        <router-link
          to="/logout"
          class="pb-1 border-b-2 border-transparent hover:border-green-500 transition"
          active-class="border-green-700 text-green-700"
        >
          Logout
        </router-link>
      </li>
    </ul>
  </div>
</template>
