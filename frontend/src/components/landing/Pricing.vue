<template>
  <section id="pricing" class="container mx-auto px-4 py-16">
    <div class="text-center max-w-2xl mx-auto">
      <h2 class="text-3xl lg:text-4xl font-bold">Simple, transparent pricing</h2>
      <p class="mt-3 text-base-content/70">Start free, scale as you grow. Cancel anytime.</p>
    </div>

    <div class="mt-10 grid md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div class="card bg-base-100 border border-base-200">
        <div class="card-body">
          <h3 class="card-title">Starter</h3>
          <p class="opacity-70">For small programs and early pilots.</p>
          <div class="text-4xl font-bold mt-4">Free</div>
          <ul class="mt-4 space-y-2 text-sm">
            <li>• 1 admin seat</li><li>• 100 student profiles</li><li>• Basic search & filters</li>
          </ul>
          <div class="card-actions mt-6">
            <router-link to="/register" class="btn btn-outline w-full">Get Started</router-link>
          </div>
        </div>
      </div>

      <div class="card bg-base-100 border border-primary shadow-lg">
        <div class="card-body">
          <div class="badge badge-primary w-fit">Popular</div>
          <h3 class="card-title mt-2">Growth</h3>
          <p class="opacity-70">For faculties and career centers.</p>
          <div class="text-4xl font-bold mt-4">$49<span class="text-lg font-medium text-base-content/60">/mo</span></div>
          <ul class="mt-4 space-y-2 text-sm">
            <li>• 5 admin seats</li><li>• 5,000 profiles</li><li>• Advanced search & analytics</li>
          </ul>
          <div class="card-actions mt-6">
            <router-link to="/register" class="btn btn-primary w-full">Choose Growth</router-link>
          </div>
        </div>
      </div>

      <div class="card bg-base-100 border border-base-200">
        <div class="card-body">
          <h3 class="card-title">Enterprise</h3>
          <p class="opacity-70">Custom integrations & SLAs.</p>
          <div class="text-4xl font-bold mt-4">Let's talk</div>
          <ul class="mt-4 space-y-2 text-sm">
            <li>• Unlimited seats & profiles</li><li>• SSO & custom roles</li><li>• Priority support</li>
          </ul>
          <div class="card-actions mt-6">
            <router-link to="/contact" class="btn btn-outline w-full">Contact Sales</router-link>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
