<script setup lang="ts">
const emit = defineEmits<{ (e: 'goto', id: string): void }>()
</script>

<template>
  <section class="relative overflow-hidden bg-gradient-to-b from-[#F2F6FC] via-white to-[#F2F6FC]">
    <!-- Decorative gradient overlay -->
    <div class="absolute inset-0 -z-10 bg-gradient-to-br from-[#56A45C]/10 via-[#44B15B]/10 to-transparent"></div>

    <div
      class="container mx-auto px-6 py-20 lg:py-28 grid lg:grid-cols-2 gap-12 items-center text-[#0A3B1F]"
    >
      <!-- LEFT CONTENT -->
      <div class="space-y-6">
        <div
          class="inline-flex items-center bg-[#56A45C]/10 text-[#0A3B1F] px-4 py-1 rounded-full text-sm font-medium"
        >
          <span class="w-2.5 h-2.5 bg-[#56A45C] rounded-full mr-2"></span>
          Introducing the new Com-Seek platform
        </div>

        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">
          Empowering <span class="text-[#56A45C]">Students</span>,
          <span class="text-[#44B15B]">Alumni</span>, and
          <span class="text-[#1F7AB9]">Employers</span> to connect.
        </h1>

        <p class="text-lg text-[#0A3B1F]/70 max-w-xl">
          Build your profile, discover verified companies, and manage your hiring workflow —
          all in one green, modern, and easy-to-use platform designed for the KU community.
        </p>

        <!-- CTA BUTTONS -->
        <div class="mt-8 flex flex-wrap gap-4">
          <router-link
            to="/register"
            class="btn border-none text-white text-lg px-8 bg-[#56A45C] hover:bg-[#44B15B] transition-all duration-200 shadow-md"
          >
            Get Started
          </router-link>
          <button
            class="btn btn-outline text-lg border-[#56A45C] text-[#56A45C] hover:bg-[#56A45C] hover:text-white transition-all duration-200"
            @click="emit('goto', 'howitworks')"
          >
            See How It Works
          </button>
        </div>

        <!-- TRUST SECTION -->
        <div class="mt-10 flex items-center gap-6">
          <div class="avatar-group -space-x-4 rtl:space-x-reverse">
            <div class="avatar">
              <div class="w-10 ring ring-[#56A45C]/20">
                <img src="https://i.pravatar.cc/100?img=12" alt="user" />
              </div>
            </div>
            <div class="avatar">
              <div class="w-10 ring ring-[#56A45C]/20">
                <img src="https://i.pravatar.cc/100?img=32" alt="user" />
              </div>
            </div>
            <div class="avatar">
              <div class="w-10 ring ring-[#56A45C]/20">
                <img src="https://i.pravatar.cc/100?img=48" alt="user" />
              </div>
            </div>
            <div class="avatar placeholder">
              <div class="w-10 bg-[#56A45C] text-white">+9k</div>
            </div>
          </div>
          <p class="text-sm text-[#0A3B1F]/70 max-w-[200px]">
            Trusted by KU students, alumni, and leading employers.
          </p>
        </div>
      </div>

      <!-- RIGHT MOCKUP -->
      <div class="relative">
        <div
          class="mockup-window border border-[#56A45C]/20 bg-white shadow-xl rounded-xl overflow-hidden"
        >
          <div class="bg-[#F2F6FC] px-8 py-10">
            <div class="grid md:grid-cols-3 gap-4">
              <div
                v-for="(card, i) in [
                  { title: 'Company', text: 'Profiles, verification, and openings.' },
                  { title: 'Alumni', text: 'Career paths and mentorship.' },
                  { title: 'Jobs', text: 'Search, filter, and apply instantly.' },
                ]"
                :key="i"
                class="card bg-white border border-[#0A3B1F]/10 shadow-sm hover:shadow-md transition-all"
              >
                <div class="card-body">
                  <h3 class="card-title text-[#0A3B1F]">{{ card.title }}</h3>
                  <p class="text-sm text-[#0A3B1F]/70">{{ card.text }}</p>
                  <div class="card-actions">
                    <button
                      class="btn btn-sm border-[#56A45C] text-[#56A45C] hover:bg-[#56A45C] hover:text-white transition-all"
                    >
                      View
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div class="mt-6">
              <input
                type="text"
                placeholder="Search roles, skills, or companies…"
                class="input input-bordered w-full border-[#0A3B1F]/20 focus:border-[#56A45C] focus:ring-[#56A45C]/30"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
