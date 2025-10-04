<script setup lang="ts">
// Page-only helpers (sections self-handle their internal state)
function scrollToId(id: string) {
  const el = document.getElementById(id)
  if (el) el.scrollIntoView({ behavior: "smooth", block: "start" })
}
</script>

<template>
  <div class="min-h-screen flex flex-col bg-base-100">
    <LandingNavBar @goto="scrollToId" />

    <LandingHero @goto="scrollToId" />
    <LandingStats />
    <LandingFeatures />
    <LandingHowItWorks />
    <LandingPricing />
    <LandingFAQ />
    <LandingCTA />

    <LandingSiteFooter />
  </div>
</template>

<script lang="ts">
import LandingNavBar from '../components/landing/NavBar.vue'
import LandingHero from '../components/landing/Hero.vue'
import LandingStats from '../components/landing/Stats.vue'
import LandingFeatures from '../components/landing/Features.vue'
import LandingHowItWorks from '../components/landing/HowItWorks.vue'
import LandingPricing from '../components/landing/Pricing.vue'
import LandingFAQ from '../components/landing/FAQ.vue'
import LandingCTA from '../components/landing/CTA.vue'
import LandingSiteFooter from '../components/landing/SiteFooter.vue'
export default {}
</script>

<style scoped>
html:has(#features) { scroll-behavior: smooth; }
</style>
