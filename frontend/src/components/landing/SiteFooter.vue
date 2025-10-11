<template>
  <footer class="border-t border-base-300 ">
    <div class="footer container mx-auto px-4 py-10 text-base-content">
      <aside>
        <img src="/logo.png" alt="logo" class="h-10 w-10 rounded-md" />
        <p>Com-Seek<br/>Connecting talent with opportunity.</p>
      </aside>
      <nav>
        <h6 class="footer-title">Legal</h6>
        <a class="link link-hover">Terms</a>
        <a class="link link-hover">Privacy</a>
        <a class="link link-hover">Cookies</a>
      </nav>
    </div>
    <div class="container mx-auto px-4 pb-8 text-sm text-base-content/60">
      © {{ new Date().getFullYear() }} Com-Seek. All rights reserved.
    </div>
  </footer>
</template>
