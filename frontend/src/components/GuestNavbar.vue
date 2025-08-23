<template>
  <div class="flex bg-[#F2F6FC] justify-between items-center top-0 w-full h-18 p-4 text-black px-42 text-xl">
    <ul class="flex space-x-4">
      <li><router-link to="/">LOGO</router-link></li>
    </ul>
    <ul class="flex space-x-15">
      <li><router-link to="/">Home</router-link></li>
      <!-- Grouped links inside one li -->
      <li class="flex space-x-3 items-center">
        <router-link to="/sign-in">Sign in</router-link>
        <span>|</span>
        <router-link to="/sign-up">Sign-up</router-link>
      </li>
      <li><router-link to="/for-company">For Company</router-link></li>
    </ul>

  </div>
</template>

<script setup>

</script>