<template>
  <div class="min-h-screen bg-gradient-to-br from-sky-300 to-blue-700 flex items-center justify-center p-4">
    <div class="bg-white rounded-xl shadow-xl w-full max-w-sm text-center p-8">
      <!-- Logo -->
      <div class="mb-6">
        <h1 class="text-3xl font-bold text-blue-800">Com-Seek</h1>
        <p class="text-gray-500 text-sm">KU Computer Engineering Job Portal</p>
      </div>

      <!-- 404 Content -->
      <div>
        <h2 class="text-xl font-semibold text-blue-800 mb-2">404 - Page Not Found</h2>
        <p class="text-gray-600 mb-4">
          Sorry, the page you’re looking for doesn’t exist.
        </p>
        <router-link to="/" class="btn btn-primary">
          Go Back Home
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NotFound'
}
</script>
