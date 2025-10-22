<script setup lang="ts">
import { ref } from 'vue'
const mobileOpen = ref(false)

// Emit navigation requests to parent (so it can smooth scroll)
const emit = defineEmits<{ (e:'goto', id:string): void }>()
</script>

<template>
  <header class="sticky top-0 z-50 bg-base-100/90 backdrop-blur border-b border-base-200">
    <div class="navbar container mx-auto px-4">
      <div class="flex-1">
        <a class="btn btn-ghost normal-case text-xl flex items-center gap-3">
          <img src="/logo.png" alt="logo" class="h-8 w-8 rounded-md" />
          Com-Seek
        </a>
      </div>

      <!-- Desktop -->
      <nav class="hidden lg:flex items-center gap-2">
        <button class="btn btn-ghost" @click="emit('goto','features')">Features</button>
        <button class="btn btn-ghost" @click="emit('goto','howitworks')">How it Works</button>
        <button class="btn btn-ghost" @click="emit('goto','pricing')">Pricing</button>
        <button class="btn btn-ghost" @click="emit('goto','faq')">FAQ</button>
        <div class="divider divider-horizontal mx-1"></div>
        <router-link to="/login" class="btn btn-ghost">Log in</router-link>
        <router-link to="/register-company" class="btn btn-ghost">For Company</router-link>
        <router-link to="/register" class="btn border-none text-white bg-[#56A45C] hover:bg-[#44B15B]">Get Started</router-link>
      </nav>

      <!-- Mobile toggle -->
      <div class="lg:hidden">
        <button class="btn btn-ghost" @click="mobileOpen = !mobileOpen" aria-label="Open Menu">
          <span class="material-icons">menu</span>
        </button>
      </div>
    </div>

    <!-- Mobile drawer -->
    <div v-if="mobileOpen" class="lg:hidden border-t border-base-200 bg-base-100">
      <div class="container mx-auto px-4 py-2 flex flex-col gap-2">
        <button class="btn btn-ghost justify-start" @click="emit('goto','features'); mobileOpen=false">Features</button>
        <button class="btn btn-ghost justify-start" @click="emit('goto','howitworks'); mobileOpen=false">How it Works</button>
        <button class="btn btn-ghost justify-start" @click="emit('goto','pricing'); mobileOpen=false">Pricing</button>
        <button class="btn btn-ghost justify-start" @click="emit('goto','faq'); mobileOpen=false">FAQ</button>
        <div class="divider my-1"></div>
        <router-link to="/login" class="btn btn-ghost justify-start" @click="mobileOpen=false">Log in</router-link>
        <router-link to="/register" class="btn btn-primary" @click="mobileOpen=false">Get Started</router-link>
      </div>
    </div>
  </header>
</template>
