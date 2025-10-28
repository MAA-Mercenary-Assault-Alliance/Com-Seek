<script setup>
import { computed } from "vue"
import { useRoute } from "vue-router"
import LandingSiteFooter from "./components/landing/SiteFooter.vue"
import FullNavBar from "./components/FullNavBar.vue";

const route = useRoute()
const isBlankLayout = computed(() => route.meta?.layout === "blank")
</script>

<template>
  <!-- BLANK LAYOUT (e.g., Landing page) -->
  <div v-if="isBlankLayout" class="min-h-screen bg-base-100">
    <router-view />
  </div>

  <!-- DEFAULT LAYOUT -->
  <div v-else class="flex flex-col relative min-h-screen">
    <FullNavBar></FullNavBar>
    <div id="all" class="flex flex-col bg-background w-full flex-grow text-black"> <!--Default Text set to black right here -->
      <div id="main" class="flex justify-center">
        <router-view />
      </div>
    </div>
  </div>
  <LandingSiteFooter />
</template>
