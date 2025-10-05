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
          <h3 id="tos-title" class="text-lg font-semibold">Terms of Service</h3>
        </div>

        <!-- Scrollable Content -->
        <div
          ref="scrollArea"
          class="px-6 py-4 max-h-[60vh] overflow-y-auto space-y-4 leading-relaxed text-sm text-gray-800"
          @scroll="checkScrolledToBottom"
        >
          <p class="text-xs text-gray-500">
            Last Updated: October 6, 2025
          </p>

          <h4 class="font-semibold text-base">1. Introduction</h4>
          <p>
            These Terms of Service (“<strong>Terms</strong>”) govern your access to and use of the
            website, application, and related services provided by <strong>[Your Organization]</strong>
            (“<strong>we</strong>,” “<strong>us</strong>,” or “<strong>our</strong>”). By creating an
            account, accessing, or using our services (collectively, the “<strong>Services</strong>”),
            you agree to be bound by these Terms. If you do not agree, do not use the Services.
          </p>

          <h4 class="font-semibold text-base">2. Eligibility</h4>
          <p>
            You must be at least the age of majority in your jurisdiction (or have verifiable parental consent)
            and capable of forming a binding contract to use the Services. You represent and warrant you meet these
            requirements and are not prohibited from using the Services under applicable law.
          </p>

          <h4 class="font-semibold text-base">3. Accounts & Security</h4>
          <ul class="list-disc pl-6 space-y-1">
            <li><strong>Registration:</strong> Provide accurate, current, and complete information and keep it updated.</li>
            <li><strong>Credentials:</strong> Keep your login credentials confidential; you’re responsible for all activity under your account.</li>
            <li><strong>Notification:</strong> Notify us immediately at <em>[support@email.com]</em> if you suspect unauthorized access.</li>
            <li><strong>One Account:</strong> Unless we permit otherwise, you may not maintain more than one account.</li>
          </ul>

          <h4 class="font-semibold text-base">4. User Content</h4>
          <p>
            “<strong>User Content</strong>” includes any documents, data, text, images, audio, video, or other materials
            you upload, submit, or make available through the Services (e.g., transcripts or resumes).
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
            The Services (excluding User Content) are owned by or licensed to <strong>[Your Organization]</strong> and
            protected by intellectual property laws. Except as expressly permitted, you may not reproduce, modify,
            distribute, or create derivative works without prior written consent.
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
            If you believe content on the Services infringes your copyright, email <em>[legal@email.com]</em> with:
            (a) identification of the copyrighted work; (b) identification of the allegedly infringing material;
            (c) your contact information; (d) a good-faith statement; (e) a statement under penalty of perjury; and
            (f) your physical or electronic signature.
          </p>

          <h4 class="font-semibold text-base">12. Suspension & Termination</h4>
          <ul class="list-disc pl-6 space-y-1">
            <li>We may suspend/terminate access for violations, suspected fraud, or to comply with legal obligations, and may discontinue Services or parts thereof with reasonable notice where required by law.</li>
            <li>You may stop using the Services at any time. Provisions that by their nature should survive (e.g., IP, disclaimers, limitation of liability) will survive termination.</li>
          </ul>

          <h4 class="font-semibold text-base">13. Disclaimers</h4>
          <p class="italic">
            TO THE MAXIMUM EXTENT PERMITTED BY LAW, THE SERVICES ARE PROVIDED “AS IS” AND “AS AVAILABLE” WITHOUT
            WARRANTIES OF ANY KIND, WHETHER EXPRESS, IMPLIED, OR STATUTORY, INCLUDING IMPLIED WARRANTIES OF
            MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, AND NON-INFRINGEMENT. WE DO NOT WARRANT THAT THE
            SERVICES WILL BE UNINTERRUPTED, ERROR-FREE, SECURE, OR FREE OF HARMFUL COMPONENTS.
          </p>

          <h4 class="font-semibold text-base">14. Limitation of Liability</h4>
          <p class="italic">
            TO THE MAXIMUM EXTENT PERMITTED BY LAW, IN NO EVENT SHALL <strong>[YOUR ORGANIZATION]</strong> OR ITS
            AFFILIATES, OFFICERS, EMPLOYEES, OR AGENTS BE LIABLE FOR ANY INDIRECT, INCIDENTAL, SPECIAL, CONSEQUENTIAL,
            EXEMPLARY, OR PUNITIVE DAMAGES, OR ANY LOSS OF PROFITS, REVENUE, DATA, OR GOODWILL, ARISING OUT OF OR
            RELATED TO YOUR USE OF THE SERVICES, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGES. OUR TOTAL
            AGGREGATE LIABILITY FOR ANY CLAIMS RELATING TO THE SERVICES SHALL NOT EXCEED THE GREATER OF (A) THE AMOUNT
            YOU PAID US IN THE 3 MONTHS PRECEDING THE EVENT GIVING RISE TO THE CLAIM OR (B) THB 5,000.
          </p>

          <h4 class="font-semibold text-base">15. Indemnification</h4>
          <p>
            You agree to defend, indemnify, and hold harmless <strong>[Your Organization]</strong> and its affiliates,
            officers, employees, and agents from and against any claims, liabilities, damages, losses, and expenses
            (including reasonable attorneys’ fees) arising out of or related to your User Content, your use of the
            Services, or your violation of these Terms or any applicable law.
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
            with a new “Last Updated” date or by other reasonable means. Your continued use after such changes constitutes
            acceptance of the revised Terms.
          </p>

          <h4 class="font-semibold text-base">18. Contact</h4>
          <p>
            If you have questions about these Terms, please contact us at:
            <br />Email: <em>[support@email.com]</em>
            <br />Address: <em>[Company Address, Bangkok, Thailand]</em>
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
            I agree to the Terms of Service.
            <span v-if="!scrolledToBottom" class="text-gray-500">
              (scroll to the end to enable)
            </span>
          </label>

          <div class="ml-auto flex gap-2">
            <button class="btn btn-ghost" @click="onClose">Cancel</button>
            <button
              class="btn btn-primary"
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
