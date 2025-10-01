  <template>
  <div class="py-32">
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
      <div class="form-control">
        <label for="password" class="label">
          <span class="label-text">Password</span>
        </label>
        <input
          type="password"
          id="password"
          v-model="form.password"
          placeholder="Enter your password"
          class="input input-bordered w-full"
          :class="{ 'input-error': errors.password }"
          required
        />
        <label v-if="errors.password" class="label text-error text-sm">
          {{ errors.password }}
        </label>
      </div>

      <!-- Submit button -->
      <button
        type="submit"
        class="btn btn-primary w-full"
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
  </div>
</template>

<script>
export default {
  name: "LoginForm",
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
      loading: false
    };
  },
  methods: {
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
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 1500));

        this.alert = { message: "Login successful! Redirecting...", type: "success" };

        setTimeout(() => {
          this.$router.push("/dashboard");
        }, 1000);
      } catch (error) {
        this.alert = { message: "Invalid credentials. Please try again.", type: "error" };
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>
