<script setup lang="ts">
import { ref } from 'vue'
const faqOpen = ref<number | null>(0)
</script>

<template>
  <section id="faq" class="container mx-auto px-4 py-16">
    <div class="text-center max-w-2xl mx-auto">
      <h2 class="text-3xl lg:text-4xl font-bold">Frequently asked questions</h2>
      <p class="mt-3 text-base-content/70">Quick answers to common questions.</p>
    </div>

    <div class="mt-8 max-w-3xl mx-auto space-y-3">
      <div class="collapse collapse-arrow border border-base-300 rounded-box"
           :class="{ 'collapse-open': faqOpen===0 }"
           @click="faqOpen = faqOpen===0 ? null : 0">
        <input type="checkbox" :checked="faqOpen===0" />
        <div class="collapse-title text-lg font-medium">Do I need a backend to try this?</div>
        <div class="collapse-content">
          <p>You can start with mock data and wire the API later (Flask/Go/Node). Components are decoupled and typed.</p>
        </div>
      </div>

      <div class="collapse collapse-arrow border border-base-300 rounded-box"
           :class="{ 'collapse-open': faqOpen===1 }"
           @click="faqOpen = faqOpen===1 ? null : 1">
        <input type="checkbox" :checked="faqOpen===1" />
        <div class="collapse-title text-lg font-medium">Is it compatible with DaisyUI themes?</div>
        <div class="collapse-content">
          <p>Yes. It uses DaisyUI primitives, so switching themes is one line in your tailwind config or HTML <code>data-theme</code>.</p>
        </div>
      </div>

      <div class="collapse collapse-arrow border border-base-300 rounded-box"
           :class="{ 'collapse-open': faqOpen===2 }"
           @click="faqOpen = faqOpen===2 ? null : 2">
        <input type="checkbox" :checked="faqOpen===2" />
        <div class="collapse-title text-lg font-medium">Can I reuse the Admin/Company/Alumni tabs?</div>
        <div class="collapse-content">
          <p>Absolutely. The layout mirrors your example with tabs and a search bar, ready to drop in.</p>
        </div>
      </div>
    </div>
  </section>
</template>
