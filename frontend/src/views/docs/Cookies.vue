<script setup lang="ts">
import GuestNavbar from '../../components/GuestNavbar.vue'
</script>


<template>
    <!-- <LandingNavBar @goto="scrollToId" /> -->
    <GuestNavbar />
  <section class="container mx-auto px-6 py-8 text-[#0A3B1F]  my-5">
    <h1 class="text-4xl font-bold mb-6">Cookie Policy</h1>
    <p class="mb-6 text-[#0A3B1F]/70">
      Last updated: August 27, 2025  
      This policy explains how <strong>Com-Seek</strong> uses cookies to
      improve functionality and security on our website.
    </p>

    <h2 class="text-2xl font-semibold mt-8 mb-3">1. What Are Cookies?</h2>
    <p class="mb-4">
      Cookies are small text files stored on your device to help websites remember
      your preferences and activity. They do not contain personal data by themselves.
    </p>

    <h2 class="text-2xl font-semibold mt-8 mb-3">2. How We Use Cookies</h2>
    <ul class="list-disc ml-6 mb-6 text-[#0A3B1F]/80">
      <li><strong>Session Cookies:</strong> Keep you logged in securely while using the platform.</li>
      <li><strong>Preference Cookies:</strong> Store language and theme settings for better usability.</li>
      <li><strong>Analytics Cookies:</strong> Track aggregated, anonymous usage data to improve performance.</li>
      <li><strong>Security Cookies:</strong> Detect suspicious or automated login attempts (anti-bot).</li>
    </ul>

    <h2 class="text-2xl font-semibold mt-8 mb-3">3. Managing Cookies</h2>
    <p class="mb-4">
      You can control or delete cookies through your browser settings.
      Disabling essential cookies may limit your ability to log in or
      access certain features on Com-Seek.
    </p>

    <h2 class="text-2xl font-semibold mt-8 mb-3">4. Third-Party Cookies</h2>
    <p class="mb-4">
      Some external integrations (e.g., analytics or email services) may use
      their own cookies under PDPA-compliant agreements. These cookies never
      collect identifiable personal data from Com-Seek users.
    </p>

    <h2 class="text-2xl font-semibold mt-8 mb-3">5. Updates</h2>
    <p>
      We may update this cookie policy occasionally. Updates will be reflected
      on this page with the “last updated” date revised accordingly.
    </p>
  </section>
</template>
