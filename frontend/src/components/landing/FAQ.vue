<script setup lang="ts">
import { ref } from 'vue'
const faqOpen = ref<number | null>(0)
</script>

<template>
  <section id="faq" class="container mx-auto px-4 py-16">
    <div class="text-center max-w-2xl mx-auto">
      <h2 class="text-3xl lg:text-4xl font-bold text-[#0A3B1F]">
        Frequently asked questions
      </h2>
      <p class="mt-3 text-[#0A3B1F]/70">
        Learn how Com-Seek helps students, alumni, and employers connect effectively.
      </p>
    </div>

    <div class="mt-8 max-w-3xl mx-auto space-y-3">
      <!-- 1 -->
      <div class="collapse collapse-arrow border border-[#0A3B1F]/15 rounded-box"
           :class="{ 'collapse-open': faqOpen===0 }"
           @click="faqOpen = faqOpen===0 ? null : 0">
        <input type="checkbox" :checked="faqOpen===0" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          How does Com-Seek help students find jobs or internships?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Com-Seek gives students a personal profile where they can showcase their skills,
            upload portfolios, and search for verified job postings from trusted companies.
            Our smart filters match students with roles that suit their major, experience, and goals.
          </p>
        </div>
      </div>

      <!-- 2 -->
      <div class="collapse collapse-arrow border border-[#0A3B1F]/15 rounded-box"
           :class="{ 'collapse-open': faqOpen===1 }"
           @click="faqOpen = faqOpen===1 ? null : 1">
        <input type="checkbox" :checked="faqOpen===1" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          What can employers do on Com-Seek?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Employers can create company profiles, post openings, and review applicant details
            in one place. Com-Seek’s dashboard lets HR teams track applications, manage interviews,
            and connect directly with potential hires from the KU student and alumni community.
          </p>
        </div>
      </div>

      <!-- 3 -->
      <div class="collapse collapse-arrow border border-[#0A3B1F]/15 rounded-box"
           :class="{ 'collapse-open': faqOpen===2 }"
           @click="faqOpen = faqOpen===2 ? null : 2">
        <input type="checkbox" :checked="faqOpen===2" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          How does Com-Seek support alumni?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Alumni can share career paths, mentor current students, and offer referrals
            within the KU network. This helps strengthen real connections between past
            graduates and new talent entering the workforce.
          </p>
        </div>
      </div>

      <!-- 4 -->
      <div class="collapse collapse-arrow border border-[#0A3B1F]/15 rounded-box"
           :class="{ 'collapse-open': faqOpen===3 }"
           @click="faqOpen = faqOpen===3 ? null : 3">
        <input type="checkbox" :checked="faqOpen===3" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          Is Com-Seek only for KU students?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Com-Seek was built for Kasetsart University’s Computer Engineering community,
            but it’s expanding to include other faculties and partner organizations.
            The goal is to grow into a broader professional ecosystem for Thai universities.
          </p>
        </div>
      </div>

      <!-- 5 -->
      <div class="collapse collapse-arrow border border-[#0A3B1F]/15 rounded-box"
           :class="{ 'collapse-open': faqOpen===4 }"
           @click="faqOpen = faqOpen===4 ? null : 4">
        <input type="checkbox" :checked="faqOpen===4" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          How is Com-Seek different from LinkedIn or other job platforms?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Unlike public job boards, Com-Seek is focused on verified university connections.
            Every student and company is validated by KU, creating a trusted environment
            where employers know they’re hiring real, qualified talent — not random applicants.
          </p>
        </div>
      </div>

      <!-- 6 -->
      <div class="collapse collapse-arrow border border-[#0A3B1F]/15 rounded-box"
           :class="{ 'collapse-open': faqOpen===5 }"
           @click="faqOpen = faqOpen===5 ? null : 5">
        <input type="checkbox" :checked="faqOpen===5" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          Can admins or faculty use Com-Seek to track engagement?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Yes — faculty admins can approve company profiles, monitor student participation,
            and gather insights into job outcomes. This data helps the department measure
            real employment impact and improve internship programs.
          </p>
        </div>
      </div>
    </div>
  </section>
</template>
