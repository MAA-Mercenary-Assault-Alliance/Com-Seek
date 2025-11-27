  <template>
  <div class="py-32 relative">
  <div class="max-w-md mx-auto p-6 bg-base-100 shadow-xl rounded-xl">
    <h2 class="text-2xl font-bold text-center mb-6">Login</h2>

    <div v-if="alert.message" class="mb-4">
      <div :class="`alert ${alert.type === 'error' ? 'alert-error' : 'alert-success'}`">
        <span>{{ alert.message }}</span>
      </div>
    </div>

    <form @submit.prevent="handleLogin" class="space-y-4">
      <!-- Email -->
      <div class="form-control">
        <label for="email" class="label">
          <span class="label-text">Email</span>
        </label>
        <input
          type="email"
          id="email"
          v-model="form.email"
          placeholder="Enter your email address"
          class="input input-bordered w-full"
          :class="{ 'input-error': errors.email }"
          required
        />
        <label v-if="errors.email" class="label text-error text-sm">
          {{ errors.email }}
        </label>
      </div>

      <!-- Password -->
      <PasswordField
        id="login-password"
        label="Password"
        v-model="form.password"
        :error="errors.password"
        autocomplete="current-password"
        required
      />

      <!-- Submit button -->
      <button
        type="submit"
        class="btn bg-[#56A45C] text-white hover:bg-[#44B15B] w-full"
        :disabled="loading"
      >
        {{ loading ? 'LOGGING IN...' : 'LOGIN' }}
      </button>
    </form>

    <!-- Switch to register -->
<div class="text-center mt-6 border-t pt-4">
  <p>
    Don't have an account?
    <router-link to="/register" class="link link-primary">
      Sign up here
    </router-link>
  </p>
</div>

  </div>
  <!-- Inactivity timeout modal -->
  <div
    v-if="showTimeoutModal"
    class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 px-4"
  >
    <div class="bg-white rounded-xl shadow-2xl p-6 max-w-sm w-full space-y-4 text-center">
      <h3 class="text-xl font-semibold text-[#0A3B1F]">Login timed out</h3>
      <p class="text-sm text-[#0A3B1F]/80">
        You were inactive for 2 minutes on the login page. You&rsquo;ll be sent back to the landing page.
      </p>
      <button
        class="btn bg-[#56A45C] text-white hover:bg-[#44B15B] w-full"
        @click="redirectToLanding"
      >
        Go to landing page now
      </button>
    </div>
  </div>
  </div>
</template>

<script>
import { api } from '../../../api/client';
import { defineAsyncComponent } from 'vue';

export default {
  name: "LoginForm",
  components: {
    PasswordField: defineAsyncComponent(() => import('@/components/auth/PasswordField.vue')), // ✅
  },
  emits: ["switch-to-register"],
  data() {
    return {
      form: {
        email: "",
        password: ""
      },
      errors: {},
      alert: {
        message: "",
        type: ""
      },
      loading: false,
      showPassword: false,
      inactivityTimer: null,
      redirectTimer: null,
      activityEvents: null,
      timeoutDelayMs: 2 * 60 * 1000, // 2 minutes
      showTimeoutModal: false,
    };
  },
  methods: {
    toggleShowPassword() {
      this.showPassword = !this.showPassword;
    },
    
    validateEmail(email) {
      const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      return re.test(email);
    },

    validateForm() {
      this.errors = {};

      if (!this.validateEmail(this.form.email)) {
        this.errors.email = "Please enter a valid email address.";
      }

      if (!this.form.password) {
        this.errors.password = "Password is required.";
      }

      return Object.keys(this.errors).length === 0;
    },

    async handleLogin() {
      if (!this.validateForm()) return;

      this.loading = true;
      this.alert = { message: "Logging in...", type: "success" };

      try {
        const response = await api.post("/auth/login", {
          email: this.form.email,
          password: this.form.password,
        });

        localStorage.setItem('email', response.data.email);
        localStorage.setItem('role', response.data.role);

        this.alert = { message: "Login successful! Redirecting...", type: "success" };

        // if (response.data.role === "student") {
        //   this.$router.push({ name: "StudentProfile" });
        // } else if (response.data.role === "company") {
        //   this.$router.push({ path: '/' });
        // } else if (response.data.role === "admin") {
        //   this.$router.push({ path: '/' });
        // } else {
        //   this.$router.push("/");
        // }

        this.$router.push("/home").then(() => {
          window.location.reload();
        });

      } catch (error) {
        const msg =
          error?.response?.data?.error ||
          error?.message ||
          "Invalid credentials. Please try again.";
        this.alert = { message: msg, type: "error" };
      } finally {
        this.loading = false;
      }
    },
    handleUserActivity() {
      if (this.showTimeoutModal) return;
      this.resetInactivityTimer();
    },
    startInactivityTimer() {
      this.clearInactivityTimer();
      this.inactivityTimer = setTimeout(this.handleTimeout, this.timeoutDelayMs);
    },
    resetInactivityTimer() {
      this.startInactivityTimer();
    },
    clearInactivityTimer() {
      if (this.inactivityTimer) {
        clearTimeout(this.inactivityTimer);
        this.inactivityTimer = null;
      }
    },
    handleTimeout() {
      this.showTimeoutModal = true;
      this.redirectTimer = setTimeout(() => {
        this.redirectToLanding();
      }, 2000);
    },
    redirectToLanding() {
      if (this.redirectTimer) {
        clearTimeout(this.redirectTimer);
        this.redirectTimer = null;
      }
      this.$router.push("/landing-page");
    },
    attachActivityListeners() {
      const events = ["mousemove", "keydown", "click", "touchstart"];
      this.activityEvents = events;
      events.forEach((evt) => window.addEventListener(evt, this.handleUserActivity));
    },
    removeActivityListeners() {
      if (!this.activityEvents) return;
      this.activityEvents.forEach((evt) =>
        window.removeEventListener(evt, this.handleUserActivity)
      );
      this.activityEvents = null;
    },
  },
  mounted() {
    this.attachActivityListeners();
    this.startInactivityTimer();
  },
  beforeUnmount() {
    this.removeActivityListeners();
    this.clearInactivityTimer();
    if (this.redirectTimer) {
      clearTimeout(this.redirectTimer);
      this.redirectTimer = null;
    }
  }
};
</script>
