<template>
  <footer class="border-t border-base-300 ">
    <div class="footer container mx-auto px-4 py-10 text-base-content">
      <aside>
        <img src="/logo.png" alt="logo" class="h-10 w-10 rounded-md" />
        <p>Com-Seek<br/>Connecting talent with opportunity.</p>
      </aside>
      <nav>
        <router-link to="/docs/terms" class="link link-hover text-[#0A3B1F] hover:text-[#56A45C]">
          Terms
        </router-link>
        <router-link to="/docs/privacy" class="link link-hover text-[#0A3B1F] hover:text-[#56A45C]">
          Privacy
        </router-link>
        <router-link to="/docs/cookies" class="link link-hover text-[#0A3B1F] hover:text-[#56A45C]">
          Cookies
        </router-link>
      </nav>
    </div>
    <div class="container mx-auto px-4 pb-8 text-sm text-base-content/60">
      © {{ new Date().getFullYear() }} Com-Seek. All rights reserved.
    </div>
  </footer>
</template>
