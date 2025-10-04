<template>
  <footer class="border-t border-base-200">
    <div class="footer container mx-auto px-4 py-10 text-base-content">
      <aside>
        <img src="/logo.png" alt="logo" class="h-10 w-10 rounded-md" />
        <p>Com-Seek<br/>Connecting talent with opportunity.</p>
      </aside>
      <nav>
        <h6 class="footer-title">Product</h6>
        <a class="link link-hover">Features</a>
        <a class="link link-hover">Pricing</a>
        <a class="link link-hover">Changelog</a>
      </nav>
      <nav>
        <h6 class="footer-title">Company</h6>
        <a class="link link-hover">About</a>
        <a class="link link-hover">Careers</a>
        <a class="link link-hover">Contact</a>
      </nav>
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
