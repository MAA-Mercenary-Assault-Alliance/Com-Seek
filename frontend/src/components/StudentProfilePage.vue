<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Top bar -->
    <div class="navbar bg-base-200 px-6 shadow-md">
      <div class="flex-1">
        <a class="text-xl font-bold flex items-center">
          <img src="/logo.png" alt="logo" class="h-8 mr-2" />
          Com-Seek
        </a>
      </div>
      <div class="flex-none space-x-6">
        <router-link to="/" class="hover:underline">Home</router-link>
        <router-link to="/profile" class="hover:underline">Profile</router-link>
        <router-link to="/logout" class="hover:underline">Logout</router-link>
      </div>
    </div>

    <!-- Cover image -->
    <div class="relative">
      <img
        src="https://picsum.photos/1200/250"
        alt="cover"
        class="w-full h-60 object-cover"
      />
      <button
        class="absolute right-4 bottom-4 btn btn-circle btn-outline"
      >
        <i class="fas fa-cog"></i>
      </button>
    </div>

    <!-- Profile Info -->
    <div class="w-11/12 md:w-8/12 mx-auto bg-white rounded-xl shadow-md -mt-16 p-6 relative">
      <div class="flex items-center space-x-6">
        <!-- Profile avatar -->
        <div class="avatar">
          <div class="w-28 rounded-full ring ring-primary ring-offset-base-100 ring-offset-2">
            <img src="https://i.pravatar.cc/150?img=12" />
          </div>
        </div>

        <!-- User Info -->
        <div>
          <h1 class="text-2xl font-bold text-gray-800">Reed Richard</h1>
          <div class="flex items-center space-x-4 mt-1 text-xl text-primary">
            <i class="fab fa-facebook"></i>
            <i class="fab fa-x-twitter"></i>
            <i class="fab fa-instagram"></i>
            <i class="fab fa-github"></i>
          </div>
          <p class="text-sm text-gray-500 mt-2">April 13th, 2019</p>
          <p class="text-gray-600 text-sm">This is going to be fantastic!</p>
        </div>
      </div>
    </div>

    <!-- Details Section -->
    <div class="w-11/12 md:w-8/12 mx-auto mt-6 bg-white rounded-xl shadow-md p-6">
      <div class="divide-y divide-gray-300">
        <div class="py-3">
          <p class="font-semibold text-gray-700">Username</p>
          <p class="text-gray-600">Fanstatic-Username</p>
        </div>
        <div class="py-3">
          <p class="font-semibold text-gray-700">Email Address</p>
          <p class="text-gray-600">mybigthing@gmail.com</p>
        </div>
        <div class="py-3">
          <p class="font-semibold text-gray-700">First Name</p>
          <p class="text-gray-600">Reed</p>
        </div>
        <div class="py-3">
          <p class="font-semibold text-gray-700">Last Name</p>
          <p class="text-gray-600">Richard</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "StudentProfile",
};
</script>
