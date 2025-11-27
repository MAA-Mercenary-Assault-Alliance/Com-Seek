<template>
  <transition name="fade">
    <div v-if="show" class="fixed inset-0 z-50 flex items-center justify-center">
      <!-- Backdrop -->
      <div class="absolute inset-0 bg-black/50" @click="onClose"></div>

      <!-- Modal -->
      <div
        class="relative z-10 w-full max-w-2xl bg-white rounded-xl shadow-xl overflow-hidden"
        role="dialog"
        aria-modal="true"
        aria-labelledby="tos-title"
      >
        <!-- Header -->
        <div class="px-6 py-4 border-b">
          <h3 id="tos-title" class="text-lg font-semibold">Terms of Service & Privacy Policy</h3>
        </div>

        <!-- Scrollable Content -->
        <div
          ref="scrollArea"
          class="px-6 py-4 max-h-[60vh] overflow-y-auto space-y-4 leading-relaxed text-sm text-gray-800"
          @scroll="checkScrolledToBottom"
        >
          <p class="text-xs text-gray-500">
            Terms Last Updated: October 6, 2025
          </p>

          <h4 class="font-semibold text-base">1. Introduction</h4>
          <p>
            These Terms of Service ("<strong>Terms</strong>") govern your access to and use of the website, application,
            and related services provided by <strong>Com-Seek</strong> ("<strong>we</strong>", "<strong>us</strong>", or
            "<strong>our</strong>"). By creating an account, accessing, or using our services (collectively, the
            "<strong>Services</strong>"), you agree to be bound by these Terms. If you do not agree, do not use the Services.
          </p>

          <h4 class="font-semibold text-base">2. Eligibility</h4>
          <p>
            You must be at least the age of majority in your jurisdiction (or have verifiable parental consent) and
            capable of forming a binding contract to use the Services. You represent and warrant you meet these
            requirements and are not prohibited from using the Services under applicable law.
          </p>

          <h4 class="font-semibold text-base">3. Accounts & Security</h4>
          <ul class="list-disc pl-6 space-y-1">
            <li><strong>Registration:</strong> Provide accurate, current, and complete information and keep it updated.</li>
            <li><strong>Credentials:</strong> Keep your login credentials confidential; you are responsible for all activity under your account.</li>
            <li><strong>Notification:</strong> Notify us immediately at <em>support@email.com</em> if you suspect unauthorized access.</li>
            <li><strong>One Account:</strong> Unless we permit otherwise, you may not maintain more than one account.</li>
          </ul>

          <h4 class="font-semibold text-base">4. User Content</h4>
          <p>
            "<strong>User Content</strong>" includes any documents, data, text, images, audio, video, or other materials you
            upload, submit, or make available through the Services (e.g., transcripts or resumes).
          </p>
          <ul class="list-disc pl-6 space-y-1">
            <li><strong>Ownership:</strong> You retain your rights. You grant us a worldwide, non-exclusive, royalty-free license to host, store, process, transmit, display, and perform your User Content as reasonably necessary to operate and improve the Services.</li>
            <li><strong>Rights & Permissions:</strong> You have all rights to upload and share User Content; it does not violate third-party rights or laws.</li>
            <li><strong>Prohibited Content:</strong> No illegal, infringing, harmful, deceptive, defamatory, obscene, malware, or privacy-invading content.</li>
          </ul>

          <h4 class="font-semibold text-base">5. Acceptable Use</h4>
          <p>You agree not to, and not to allow others to:</p>
          <ul class="list-disc pl-6 space-y-1">
            <li>Reverse engineer, decompile, or attempt to derive source code of the Services.</li>
            <li>Interfere with or disrupt the integrity or performance of the Services.</li>
            <li>Access the Services using automated means without written consent.</li>
            <li>Use the Services unlawfully or in violation of applicable regulations.</li>
            <li>Harvest or collect information about other users without consent.</li>
          </ul>

          <h4 class="font-semibold text-base">6. Documents & Verification</h4>
          <p>
            If the Services require submission of official documents (e.g., transcripts), you confirm they are authentic,
            unaltered, and lawfully obtained. We may use third-party tools to process or verify documents and may suspend
            or terminate access for suspected fraud or misuse.
          </p>

          <h4 class="font-semibold text-base">7. Intellectual Property</h4>
          <p>
            The Services (excluding User Content) are owned by or licensed to <strong>Com-Seek</strong> and protected by
            intellectual property laws. Except as expressly permitted, you may not reproduce, modify, distribute, or create
            derivative works without prior written consent.
          </p>

          <h4 class="font-semibold text-base">8. Third-Party Services & Links</h4>
          <p>
            The Services may include links to or integrations with third-party services. We are not responsible for
            third-party content, policies, or practices. Your use of third-party services is subject to their terms.
          </p>

          <h4 class="font-semibold text-base">9. Privacy</h4>
          <p>
            Our collection and use of personal data are described in our <strong>Privacy Policy</strong>, which is
            incorporated by reference. By using the Services, you consent to our data practices as described therein.
          </p>

          <h4 class="font-semibold text-base">10. Fees & Payments (If Applicable)</h4>
          <ul class="list-disc pl-6 space-y-1">
            <li><strong>Pricing:</strong> We may offer paid features or plans; prices/features may change with reasonable notice where required by law.</li>
            <li><strong>Billing:</strong> You authorize charges to your selected payment method for fees and taxes.</li>
            <li><strong>Refunds:</strong> Except where required by law or explicitly stated, fees are non-refundable.</li>
          </ul>

          <h4 class="font-semibold text-base">11. DMCA / IP Complaints</h4>
          <p>
            If you believe content on the Services infringes your copyright, email <em>legal@email.com</em> with: (a)
            identification of the copyrighted work; (b) identification of the allegedly infringing material; (c) your
            contact information; (d) a good-faith statement; (e) a statement under penalty of perjury; and (f) your
            physical or electronic signature.
          </p>

          <h4 class="font-semibold text-base">12. Suspension & Termination</h4>
          <ul class="list-disc pl-6 space-y-1">
            <li>We may suspend or terminate access for violations, suspected fraud, or to comply with legal obligations, and may discontinue Services or parts thereof with reasonable notice where required by law.</li>
            <li>You may stop using the Services at any time. Provisions that by their nature should survive (e.g., IP, disclaimers, limitation of liability) will survive termination.</li>
          </ul>

          <h4 class="font-semibold text-base">13. Disclaimers</h4>
          <p class="italic">
            TO THE MAXIMUM EXTENT PERMITTED BY LAW, THE SERVICES ARE PROVIDED "AS IS" AND "AS AVAILABLE" WITHOUT
            WARRANTIES OF ANY KIND, WHETHER EXPRESS, IMPLIED, OR STATUTORY, INCLUDING IMPLIED WARRANTIES OF
            MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, AND NON-INFRINGEMENT. WE DO NOT WARRANT THAT THE SERVICES
            WILL BE UNINTERRUPTED, ERROR-FREE, SECURE, OR FREE OF HARMFUL COMPONENTS.
          </p>

          <h4 class="font-semibold text-base">14. Limitation of Liability</h4>
          <p class="italic">
            TO THE MAXIMUM EXTENT PERMITTED BY LAW, IN NO EVENT SHALL <strong>Kasetsart University</strong> OR ITS
            AFFILIATES, OFFICERS, EMPLOYEES, OR AGENTS BE LIABLE FOR ANY INDIRECT, INCIDENTAL, SPECIAL, CONSEQUENTIAL,
            EXEMPLARY, OR PUNITIVE DAMAGES, OR ANY LOSS OF PROFITS, REVENUE, DATA, OR GOODWILL, ARISING OUT OF OR
            RELATED TO YOUR USE OF THE SERVICES, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGES. OUR TOTAL AGGREGATE
            LIABILITY FOR ANY CLAIMS RELATING TO THE SERVICES SHALL NOT EXCEED THE GREATER OF (A) THE AMOUNT YOU PAID US
            IN THE 3 MONTHS PRECEDING THE EVENT GIVING RISE TO THE CLAIM OR (B) THB 5,000.
          </p>

          <h4 class="font-semibold text-base">15. Indemnification</h4>
          <p>
            You agree to defend, indemnify, and hold harmless <strong>Com-Seek</strong> and its affiliates, officers,
            employees, and agents from and against any claims, liabilities, damages, losses, and expenses (including
            reasonable attorneys' fees) arising out of or related to your User Content, your use of the Services, or your
            violation of these Terms or any applicable law.
          </p>

          <h4 class="font-semibold text-base">16. Governing Law & Dispute Resolution</h4>
          <p>
            These Terms are governed by the laws of <strong>Thailand</strong>, without regard to conflict-of-laws rules.
            You agree to the exclusive jurisdiction and venue of the courts located in <strong>Bangkok, Thailand</strong>,
            for any disputes arising out of or relating to these Terms or the Services. Where required by law, you and we
            will first attempt informal resolution before initiating legal proceedings.
          </p>

          <h4 class="font-semibold text-base">17. Changes to the Terms</h4>
          <p>
            We may modify these Terms from time to time. Material changes will be notified by posting an updated version
            with a new "Last Updated" date or by other reasonable means. Your continued use after such changes constitutes
            acceptance of the revised Terms.
          </p>

          <h4 class="font-semibold text-base">18. Contact</h4>
          <p>
            If you have questions about these Terms, please contact us at:
            <br />Email: <em>support@email.com</em>
            <br />Address: <em>Bagkhen, Bangkok, Thailand</em>
          </p>

          <h3 class="text-base font-semibold pt-2">Privacy Policy</h3>
          <p class="text-xs text-gray-500">Privacy Policy Last Updated: November 22, 2025</p>
          <p>
            <strong>Com-Seek</strong> is an online job platform designed for Kasetsart University. This privacy policy
            explains how we use the personal data we collect from you when you use our website.
          </p>

          <h4 class="font-semibold text-base">General Data Protection Regulation (GDPR)</h4>
          <p>We are a Data Controller of your information.</p>

          <h4 class="font-semibold text-base">Data We Collect</h4>
          <ul class="list-disc pl-6 space-y-1">
            <li>Personal identification information (name, email address, phone number, etc.).</li>
            <li>CV / Portfolio.</li>
            <li>Transcript.</li>
            <li>Job post information.</li>
            <li>Job application information.</li>
          </ul>

          <h4 class="font-semibold text-base">How We Collect Your Data</h4>
          <p>You directly provide most of the data we collect when you:</p>
          <ul class="list-disc pl-6 space-y-1">
            <li>Register for any type of account.</li>
            <li>Edit personal data in your account.</li>
            <li>Post any job.</li>
            <li>Apply for any job.</li>
            <li>Use or view our website via your browser's cookies.</li>
          </ul>
          <p>
            We may also receive your data indirectly from Google reCAPTCHA, a security verification service that collects
            browser information, IP address, and interaction data to prevent spam and abuse.
          </p>

          <h4 class="font-semibold text-base">How We Use Your Data</h4>
          <ul class="list-disc pl-6 space-y-1">
            <li>Process your job application or job post.</li>
            <li>Manage your account.</li>
            <li>Email you with important notifications (for example, approving your account for use).</li>
            <li>Email the corresponding company your job application.</li>
          </ul>

          <h4 class="font-semibold text-base">Your Data Protection Rights</h4>
          <p>Every user is entitled to the following rights (subject to applicable law):</p>
          <ul class="list-disc pl-6 space-y-1">
            <li><strong>The right to access:</strong> You can request copies of your personal data.</li>
            <li><strong>The right to rectification:</strong> You can request that we correct or complete information you believe is inaccurate.</li>
            <li><strong>The right to erasure:</strong> You can request that we erase your personal data under certain conditions.</li>
            <li><strong>The right to restrict processing:</strong> You can request that we restrict the processing of your personal data under certain conditions.</li>
            <li><strong>The right to object to processing:</strong> You can object to our processing of your personal data under certain conditions.</li>
            <li><strong>The right to data portability:</strong> You can request that we transfer the data we have collected to another organization, or directly to you, under certain conditions.</li>
          </ul>

          <h4 class="font-semibold text-base">Cookies</h4>
          <p>
            Cookies are text files placed on your computer to collect standard Internet log information and visitor behavior
            information. When you visit our website, we may collect information from you automatically through cookies or
            similar technology. For further information, visit allaboutcookies.org.
          </p>

          <h4 class="font-semibold text-base">How We Use Cookies</h4>
          <ul class="list-disc pl-6 space-y-1">
            <li>Keeping you signed in.</li>
            <li>Understanding how you use our website.</li>
            <li>Keeping your role in the session.</li>
          </ul>

          <h4 class="font-semibold text-base">Types of Cookies We Use</h4>
          <ul class="list-disc pl-6 space-y-1">
            <li><strong>Functionality:</strong> We use these cookies so that we recognize you on our website and remember your preferences.</li>
            <li><strong>Advertising:</strong> We use these cookies to collect information about your visit, the content you viewed, the links you followed, and information about your browser, device, and IP address. We sometimes share limited aspects of this data with third parties for advertising purposes, including so you may see advertising based on your browsing patterns.</li>
          </ul>

          <h4 class="font-semibold text-base">How to Manage Cookies</h4>
          <p>
            It is necessary to accept cookies to use our website. We use cookies to store session information that helps
            us verify your identity and keep you logged in.
          </p>

          <h4 class="font-semibold text-base">Privacy Policies of Other Websites</h4>
          <p>
            Our website contains links to other websites. Our privacy policy applies only to our website; if you click on
            a link to another website, you should read their privacy policy.
          </p>

          <h4 class="font-semibold text-base">Changes to Our Privacy Policy</h4>
          <p>
            We keep this privacy policy under regular review and place any updates on this web page. This privacy policy
            was last updated on 22 November 2025.
          </p>
        </div>

        <!-- Footer -->
        <div class="px-6 py-4 border-t flex items-center gap-3">
          <input
            id="agree"
            type="checkbox"
            class="checkbox checkbox-primary"
            v-model="agreeChecked"
            :disabled="!scrolledToBottom"
          />
          <label for="agree" class="text-sm">
            I have read and agree to the Terms of Service & Privacy Policy.
            <br>
            <span v-if="!scrolledToBottom" class="text-gray-500">
              (scroll to the end to enable)
            </span>
          </label>

          <div class="ml-auto flex gap-2">
            <button class="btn btn-ghost" @click="onClose">Cancel</button>
            <button
              class="btn bg-[#56A45C] text-white hover:bg-[#44B15B]"
              :disabled="!canAccept"
              @click="onAccept"
            >
              Accept & Continue
            </button>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "TermsModal",
  props: {
    show: { type: Boolean, default: false },
  },
  emits: ["close", "accept"],
  data() {
    return {
      scrolledToBottom: false,
      agreeChecked: false,
    };
  },
  computed: {
    canAccept() {
      return this.scrolledToBottom && this.agreeChecked;
    },
  },
  mounted() {
    if (this.show) this.resetState();
  },
  watch: {
    show(val) {
      if (val) this.$nextTick(() => this.resetState());
    },
  },
  methods: {
    resetState() {
      this.scrolledToBottom = false;
      this.agreeChecked = false;
      const el = this.$refs.scrollArea;
      if (el) el.scrollTop = 0;
    },
    checkScrolledToBottom(e) {
      const el = e.target;
      const nearBottom = el.scrollTop + el.clientHeight >= el.scrollHeight - 8;
      if (nearBottom) this.scrolledToBottom = true;
    },
    onClose() {
      this.$emit("close");
    },
    onAccept() {
      if (this.canAccept) this.$emit("accept");
    },
  },
};
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity .15s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
