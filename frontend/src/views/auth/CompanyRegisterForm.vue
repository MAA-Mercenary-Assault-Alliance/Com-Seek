<template>
  <div class="pt-10 pb-32">
    <!-- Watermark Icon -->
    <div
      class="absolute bottom-4 right-6 text-green-900 opacity-10 pointer-events-none select-none"
    >
      <i class="fas fa-briefcase text-[15rem]"></i>
    </div>

    <div class="w-120 mx-auto p-6 bg-white rounded-xl shadow-md overflow-y-auto">
      <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">
        <span class="text-3xl text-green-700 font-black">Company</span> Register
      </h2>

      <!-- Alert -->
      <div v-if="alert.message" class="mb-4">
        <div :class="`alert ${alert.type === 'error' ? 'alert-error' : 'alert-success'}`">
          <span>{{ alert.message }}</span>
        </div>
      </div>

      <form @submit.prevent="handleRegister" class="space-y-4">

        <!-- Account Email -->
        <div class="form-control">
          <label class="label"><span class="label-text">Account Email</span></label>
          <input
            v-model.trim="form.email"
            type="email"
            placeholder="Login email"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.email }"
            required
          />
          <label v-if="errors.email" class="label text-error text-sm">{{ errors.email }}</label>
        </div>

        <!-- Passwords (PasswordField component) -->
        <div class="form-control">
          <PasswordField
            id="password"
            label="Password"
            placeholder="Create a password"
            v-model="form.password"
          />
          <label v-if="errors.password" class="label text-error text-sm">{{ errors.password }}</label>
        </div>

        <div class="form-control">
          <PasswordField
            id="confirm"
            label="Confirm Password"
            placeholder="Repeat password"
            v-model="form.confirmPassword"
          />
          <label v-if="errors.confirmPassword" class="label text-error text-sm">{{ errors.confirmPassword }}</label>
        </div>

        <div class="divider my-2"></div>

        <!-- Company Name -->
        <div class="form-control">
          <label class="label"><span class="label-text">Company Name</span></label>
          <input
            v-model.trim="form.companyName"
            type="text"
            placeholder="Acme Robotics"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.companyName }"
            required
          />
          <label v-if="errors.companyName" class="label text-error text-sm">{{ errors.companyName }}</label>
        </div>

        <!-- Location -->
        <div class="form-control">
          <label class="label"><span class="label-text">Location</span></label>
          <input
            v-model.trim="form.location"
            type="text"
            placeholder="Bangkok, Thailand"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.location }"
            required
          />
          <label v-if="errors.location" class="label text-error text-sm">{{ errors.location }}</label>
        </div>

        <!-- Contact Email -->
        <div class="form-control">
          <label class="label"><span class="label-text">Contact Email</span></label>
          <input
            v-model.trim="form.contactEmail"
            type="email"
            placeholder="hr@acme.co"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.contactEmail }"
            required
          />
          <label v-if="errors.contactEmail" class="label text-error text-sm">{{ errors.contactEmail }}</label>
        </div>

        <!-- Contact Number -->
        <div class="form-control">
          <label class="label"><span class="label-text">Contact Number</span></label>
          <input
            v-model.trim="form.contactNumber"
            type="text"
            placeholder="02-123-4567"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.contactNumber }"
            required
          />
          <label v-if="errors.contactNumber" class="label text-error text-sm">{{ errors.contactNumber }}</label>
        </div>

        <!-- Submit -->
        <div class="form-control mt-4 flex justify-center">
          <button
            type="submit"
            class="btn btn-primary w-full"
            :disabled="loading"
            :title="!tosAccepted ? 'You must accept the Terms of Service before registering' : ''"
          >
            {{ loading
              ? 'CREATING ACCOUNT...'
              : (tosAccepted ? 'REGISTER' : 'READ & ACCEPT TERMS TO REGISTER') }}
          </button>
        </div>
      </form>
      <!-- Switch to Login -->
      <p class="text-center text-sm mt-4 border-t pt-4">
        Already have an account?
        <router-link to="/login" class="text-blue-500 hover:underline">
          Login here
        </router-link>
      </p>

      <p class="text-center text-sm">
        or
      </p>

      <p class="text-center text-sm">
        Register as a student.
        <router-link to="/register" class="text-blue-500 hover:underline">
          Register here
        </router-link>
      </p>
    </div>

    <!-- Terms Modal: uses showTos flag and emits 'accept' / 'close' -->
    <TermsModal
      :show="showTos"
      @close="showTos = false"
      @accept="onAcceptTos"
    />
  </div>
</template>

<script>
import { api } from '../../../api/client';
import { defineAsyncComponent } from 'vue';

export default {
  name: 'CompanyRegister',
  components: {
    PasswordField: defineAsyncComponent(() => import('@/components/auth/PasswordField.vue')),
    TermsModal: defineAsyncComponent(() => import('@/components/auth/TermsModal.vue')),
  },
  emits: ['switch-to-login'],
  data() {
    return {
      form: {
        companyName: '',
        location: '',
        contactEmail: '',
        contactNumber: '',
        email: '',
        password: '',
        confirmPassword: '',
      },
      errors: {},
      alert: { message: '', type: '' },
      loading: false,

      // ToS state
      tosAccepted: false,
      showTos: false,
      continueAfterTos: false,
    };
  },
  methods: {
    validateEmail(email) {
      return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
    },
    validatePassword(password) {
      return password.length >= 8;
    },

    validateForm() {
      this.errors = {};

      if (!this.form.companyName.trim()) this.errors.companyName = 'Required';
      if (!this.form.location.trim()) this.errors.location = 'Required';

      if (!this.validateEmail(this.form.contactEmail)) {
        this.errors.contactEmail = 'Invalid email';
      }
      if (!this.form.contactNumber.trim()) this.errors.contactNumber = 'Required';

      if (!this.validateEmail(this.form.email)) this.errors.email = 'Invalid email';
      if (!this.validatePassword(this.form.password)) this.errors.password = 'Minimum 8 characters';
      if (this.form.password !== this.form.confirmPassword) {
        this.errors.confirmPassword = 'Passwords do not match';
      }

      return Object.keys(this.errors).length === 0;
    },

    // Public submit handler (ToS-gated)
    async handleRegister() {
      if (!this.tosAccepted) {
        this.continueAfterTos = true;
        this.showTos = true;
        return;
      }

      if (!this.validateForm()) return;
      await this.performRegister();
    },

    // Called after ToS accept or direct submit
    async performRegister() {
      this.loading = true;
      this.alert = { message: 'Creating account...', type: 'success' };

      try {
        const body = {
          email: this.form.email,
          password: this.form.password,
          user_type: 'company',
          company_profile: {
            name: this.form.companyName,
            location: this.form.location,
            contact_email: this.form.contactEmail,
            contact_number: this.form.contactNumber,
          },
        };

        const { data } = await api.post('/auth/register', body);
        console.log('Registered:', data);

        this.alert = { message: 'Registration successful! Pending admin approval.', type: 'success' };
        setTimeout(() => {
          this.resetForm();
          this.$router.push('/login');
        }, 2000);
      } catch (error) {
        const msg = error?.response?.data?.error || error?.message || 'Registration failed.';
        this.alert = { message: msg, type: 'error' };
      } finally {
        this.loading = false;
      }
    },

    onAcceptTos() {
      this.tosAccepted = true;
      this.showTos = false;

      if (this.continueAfterTos) {
        this.continueAfterTos = false;
        if (this.validateForm()) {
          this.performRegister();
        }
      }
    },

    resetForm() {
      this.form = {
        companyName: '',
        location: '',
        contactEmail: '',
        contactNumber: '',
        email: '',
        password: '',
        confirmPassword: '',
      };
      this.errors = {};
      this.alert = { message: '', type: '' };
      // keep tosAccepted true so user doesn't need to accept again in this session
    },
  },
};
</script>

