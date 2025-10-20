<script setup lang="ts">
import { ref } from 'vue'
const open = ref<number | null>(0)
</script>

<template>
  <section id="faq" class="container mx-auto px-6 py-16">
    <div class="text-center max-w-2xl mx-auto">
      <h2 class="text-3xl lg:text-4xl font-bold text-[#0A3B1F]">Frequently asked questions</h2>
      <p class="mt-3 text-[#0A3B1F]/70">
        How Com-Seek actually helps KU Computer Engineering students/alumni and employers.
      </p>
    </div>

    <div class="mt-8 max-w-3xl mx-auto space-y-3">
      <div class="collapse collapse-arrow rounded-box border border-[#0A3B1F]/15"
           :class="{ 'collapse-open': open===0 }"
           @click="open = open===0 ? null : 0">
        <input type="checkbox" :checked="open===0" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          Who can use Com-Seek?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Com-Seek is designed for the <strong>Kasetsart University Computer Engineering</strong> community.
            Students and alumni sign up with KU verification, and employers use approved company accounts.
            This keeps the network focused and trustworthy for CE roles.
          </p>
        </div>
      </div>

      <div class="collapse collapse-arrow rounded-box border border-[#0A3B1F]/15"
           :class="{ 'collapse-open': open===1 }"
           @click="open = open===1 ? null : 1">
        <input type="checkbox" :checked="open===1" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          How does Com-Seek help students and alumni get hired?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Create a CE-focused profile, upload your resume/portfolio, and browse <strong>Computer Engineering–related</strong> openings.
            Filters (tech stacks, experience, salary range) help you quickly find roles that fit your background.
            Apply with your documents in one place.
          </p>
        </div>
      </div>

      <div class="collapse collapse-arrow rounded-box border border-[#0A3B1F]/15"
           :class="{ 'collapse-open': open===2 }"
           @click="open = open===2 ? null : 2">
        <input type="checkbox" :checked="open===2" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          What do employers do on Com-Seek?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Publish CE-relevant job postings with clear requirements and deadlines, then review submitted applications
            (resume, cover letter, and supporting files up to 10MB). You’ll receive email alerts when new applications arrive.
          </p>
        </div>
      </div>

      <div class="collapse collapse-arrow rounded-box border border-[#0A3B1F]/15"
           :class="{ 'collapse-open': open===3 }"
           @click="open = open===3 ? null : 3">
        <input type="checkbox" :checked="open===3" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          What happens after I submit an application?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            You’ll see a submission confirmation. Employers receive an email notification and can review your materials.
            Note: Com-Seek does <strong>not</strong> provide real-time messaging or detailed post-submission status tracking.
            Employers may contact you directly if you’re shortlisted.
          </p>
        </div>
      </div>

      <div class="collapse collapse-arrow rounded-box border border-[#0A3B1F]/15"
           :class="{ 'collapse-open': open===4 }"
           @click="open = open===4 ? null : 4">
        <input type="checkbox" :checked="open===4" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          Are the jobs really CE-specific?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Yes. Listings are limited to roles relevant to Computer Engineering (e.g., software engineering, embedded systems,
            data/ML engineering, DevOps, security). Non-CE roles are not accepted to keep the board focused and valuable.
          </p>
        </div>
      </div>

      <div class="collapse collapse-arrow rounded-box border border-[#0A3B1F]/15"
           :class="{ 'collapse-open': open===5 }"
           @click="open = open===5 ? null : 5">
        <input type="checkbox" :checked="open===5" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          How is my data protected?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            Sensitive data is encrypted in transit and at rest, and we align with Thailand’s Personal Data Protection Act (PDPA).
            Anti-bot checks help prevent spam submissions. You control what appears on your public profile.
          </p>
        </div>
      </div>


      <div class="collapse collapse-arrow rounded-box border border-[#0A3B1F]/15"
           :class="{ 'collapse-open': open===7 }"
           @click="open = open===7 ? null : 7">
        <input type="checkbox" :checked="open===7" />
        <div class="collapse-title text-lg font-medium text-[#0A3B1F]">
          What languages are supported and how do I get help?
        </div>
        <div class="collapse-content text-[#0A3B1F]/80">
          <p>
            The portal supports Thai and English. For help with job listings or applications, use the contact details on each
            company page or refer to the department’s guidance channels.
          </p>
        </div>
      </div>
    </div>

    <div class="mt-8 text-center text-sm text-[#0A3B1F]/60">
      Com-Seek focuses on verified KU CE users and CE-relevant roles. No in-app chat; employers will reach out directly if you advance.
    </div>
  </section>
</template>
