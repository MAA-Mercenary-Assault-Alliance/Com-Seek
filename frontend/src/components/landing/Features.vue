<template>
  <section id="features" class="container mx-auto px-6 py-16">
    <!-- Header -->
    <div class="text-center max-w-2xl mx-auto">
      <h2 class="text-3xl lg:text-4xl font-bold text-[#0A3B1F]">
        Everything you need to run a modern career portal
      </h2>
      <p class="mt-3 text-[#0A3B1F]/70">
        Modular building blocks for students, alumni, employers, and admins—without the bloat.
      </p>
    </div>

    <!-- Cards -->
    <div class="mt-10 grid md:grid-cols-2 lg:grid-cols-3 gap-6">
      <!-- Admin -->
      <div
        class="card bg-white shadow-md border border-[#0A3B1F]/10 hover:shadow-lg transition-all"
      >
        <div class="card-body">
          <div class="badge w-fit border-none text-white bg-[#56A45C]">Admin</div>
          <h3 class="card-title mt-2 text-[#0A3B1F]">Role-Based Dashboard</h3>
          <p class="text-[#0A3B1F]/70">
            Switch between Company, Alumni, and Job views instantly; approve profiles with one click.
          </p>
        </div>
      </div>

      <!-- Search -->
      <div
        class="card bg-white shadow-md border border-[#0A3B1F]/10 hover:shadow-lg transition-all"
      >
        <div class="card-body">
          <div class="badge w-fit border-none text-white bg-[#44B15B]">Search</div>
          <h3 class="card-title mt-2 text-[#0A3B1F]">Fast Search & Filters</h3>
          <p class="text-[#0A3B1F]/70">
            Type-ahead, skill tags, and salary ranges help students and HR find the right match faster.
          </p>
        </div>
      </div>

      <!-- UX -->
      <div
        class="card bg-white shadow-md border border-[#0A3B1F]/10 hover:shadow-lg transition-all"
      >
        <div class="card-body">
          <div class="badge w-fit border-none text-white bg-[#0A3B1F]">UX</div>
          <h3 class="card-title mt-2 text-[#0A3B1F]">Clean Components</h3>
          <p class="text-[#0A3B1F]/70">
            Consistent cards, stats, and modals with DaisyUI—customized to your green theme.
          </p>
        </div>
      </div>

      <!-- Security -->
      <div
        class="card bg-white shadow-md border border-[#0A3B1F]/10 hover:shadow-lg transition-all"
      >
        <div class="card-body">
          <div class="badge w-fit border border-[#0A3B1F] text-[#0A3B1F] bg-transparent">Security</div>
          <h3 class="card-title mt-2 text-[#0A3B1F]">Auth & Verification</h3>
          <p class="text-[#0A3B1F]/70">
            Verified companies, approved postings, and role-based access keep your community safe.
          </p>
        </div>
      </div>

      <!-- Insights -->
      <div
        class="card bg-white shadow-md border border-[#0A3B1F]/10 hover:shadow-lg transition-all"
      >
        <div class="card-body">
          <div class="badge w-fit border-none text-white bg-[#1F7AB9]">Insights</div>
          <h3 class="card-title mt-2 text-[#0A3B1F]">Actionable Metrics</h3>
          <p class="text-[#0A3B1F]/70">
            Track applicants, conversion, and time-to-hire; see which skills are in demand.
          </p>
        </div>
      </div>

      <!-- API -->
      <div
        class="card bg-white shadow-md border border-[#0A3B1F]/10 hover:shadow-lg transition-all"
      >
        <div class="card-body">
          <div class="badge w-fit border-none text-white bg-[#56A45C]">API</div>
          <h3 class="card-title mt-2 text-[#0A3B1F]">API-First</h3>
          <p class="text-[#0A3B1F]/70">
            Connect to your Flask/Go/Node backend with typed DTOs and clean endpoints.
          </p>
        </div>
      </div>
    </div>

    <!-- Sub-note -->
    <div class="mt-8 text-center text-sm text-[#0A3B1F]/60">
      Need something custom? We support private endpoints and faculty-specific workflows.
    </div>
  </section>
</template>
