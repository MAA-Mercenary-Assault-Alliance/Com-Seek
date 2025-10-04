<script setup lang="ts">
const emit = defineEmits<{ (e:'goto', id:string): void }>()
</script>

<template>
  <section class="relative overflow-hidden">
    <div class="absolute inset-0 -z-10 bg-gradient-to-b from-primary/10 via-base-100 to-base-100"></div>

    <div class="container mx-auto px-4 py-16 lg:py-24 grid lg:grid-cols-2 gap-10 items-center">
      <div>
        <div class="badge badge-primary badge-lg mb-4">New</div>
        <h1 class="text-4xl lg:text-6xl font-bold leading-tight">
          Connect Students, Alumni, and Employers—seamlessly.
        </h1>
        <p class="mt-4 text-base-content/70 text-lg">
          Showcase opportunities, discover talent, and manage job postings with a clean,
          admin-friendly workflow built on Vue + Tailwind + DaisyUI.
        </p>

        <div class="mt-8 flex flex-wrap gap-3">
          <router-link to="/register" class="btn btn-primary btn-lg">Start for Free</router-link>
          <button class="btn btn-ghost btn-lg" @click="emit('goto','howitworks')">See Demo</button>
        </div>

        <!-- Trust row -->
        <div class="mt-8 flex items-center gap-6">
          <div class="avatar-group -space-x-4 rtl:space-x-reverse">
            <div class="avatar"><div class="w-10"><img src="https://i.pravatar.cc/100?img=12" alt="user" /></div></div>
            <div class="avatar"><div class="w-10"><img src="https://i.pravatar.cc/100?img=32" alt="user" /></div></div>
            <div class="avatar"><div class="w-10"><img src="https://i.pravatar.cc/100?img=48" alt="user" /></div></div>
            <div class="avatar placeholder"><div class="w-10 bg-neutral-focus text-neutral-content">+9k</div></div>
          </div>
          <p class="text-sm text-base-content/60">Trusted by universities and startups across SEA.</p>
        </div>
      </div>

      <!-- Mock -->
      <div class="relative">
        <div class="mockup-window border bg-base-300 shadow-xl">
          <div class="bg-base-200 px-6 py-10">
            <div class="grid md:grid-cols-3 gap-4">
              <div class="card bg-base-100 shadow">
                <div class="card-body">
                  <h3 class="card-title">Company</h3>
                  <p class="text-sm opacity-70">Profiles, verifications, and openings.</p>
                  <div class="card-actions"><button class="btn btn-outline btn-sm">Explore</button></div>
                </div>
              </div>
              <div class="card bg-base-100 shadow">
                <div class="card-body">
                  <h3 class="card-title">Alumni</h3>
                  <p class="text-sm opacity-70">Career paths and mentorship loops.</p>
                  <div class="card-actions"><button class="btn btn-outline btn-sm">Browse</button></div>
                </div>
              </div>
              <div class="card bg-base-100 shadow">
                <div class="card-body">
                  <h3 class="card-title">Jobs</h3>
                  <p class="text-sm opacity-70">Search, filter, and apply instantly.</p>
                  <div class="card-actions"><button class="btn btn-outline btn-sm">Find jobs</button></div>
                </div>
              </div>
            </div>

            <div class="mt-6">
              <input type="text" placeholder="Search roles, skills, or companies…" class="input input-bordered w-full" />
            </div>
          </div>
        </div>
        <div class="absolute -bottom-6 -left-6 blur-2xl opacity-25 w-56 h-56 bg-primary rounded-full"></div>
        <div class="absolute -top-8 -right-8 blur-2xl opacity-25 w-56 h-56 bg-secondary rounded-full"></div>
      </div>
    </div>
  </section>
</template>
