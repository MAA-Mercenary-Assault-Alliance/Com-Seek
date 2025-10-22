<template>
  <section
    class="min-h-screen flex flex-col items-center justify-center text-center bg-gradient-to-b from-[#F2F6FC] via-white to-[#F2F6FC] text-[#0A3B1F]"
  >
    <div class="max-w-md">
      <h1 class="text-8xl font-extrabold text-[#56A45C] mb-4">404</h1>
      <h2 class="text-3xl font-bold mb-3">Page Not Found</h2>
      <p class="text-[#0A3B1F]/70 mb-8">
        The page you’re looking for doesn’t exist or has been moved.
      </p>

      <div class="flex flex-wrap justify-center gap-4">
        <router-link
          to="/landing-page"
          class="btn border-none text-white bg-[#56A45C] hover:bg-[#44B15B] transition-all"
        >
          Go to Landing Page
        </router-link>

        <router-link
          to="/login"
          class="btn btn-outline border-[#56A45C] text-[#56A45C] hover:bg-[#56A45C] hover:text-white transition-all"
        >
          Log In
        </router-link>
      </div>
    </div>
  </section>
</template>
