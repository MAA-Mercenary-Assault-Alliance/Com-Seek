<template>
  <div class="flex bg-[#F2F6FC] justify-between items-center top-0 w-full h-18 p-4 text-black px-42 text-xl">
    <ul class="flex space-x-4 items-center">
      <li>
        <router-link to="/">
          <img src="../assets/logo.png" alt="logo" class="w-12"/>
        </router-link>
      </li>
      <li class="font-[righteous] text-3xl"><router-link to="/">Com-Seek</router-link></li>
    </ul>
    <ul class="flex space-x-15 items-center">
      <li><router-link to="/">Home</router-link></li>
      <li><router-link to="/hr-dashboard">HR TEST</router-link></li>
      <li><router-link to="/applicants">Applicants TEST</router-link></li>
      <li><router-link to="/admin">Admin Test</router-link></li>
      <!-- Grouped links inside one li -->
      <li class="flex space-x-3 items-center">
        <router-link to="/sign-in">Sign in</router-link>
        <img src="../assets/vertical-line.png" alt="vertical line" class="block h-12 w-15 -ml-6 -mr-4"/>
        <router-link to="/sign-up">Sign-up</router-link>
      </li>
      <li><router-link to="/for-company">For Company</router-link></li>
    </ul>

  </div>
</template>

<script setup>

</script>