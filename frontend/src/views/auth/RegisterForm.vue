<template>
  <div class="pt-10 pb-32">
    <!-- Watermark Icon -->
    <div
      class="absolute bottom-4 right-6 text-red-900 opacity-10 pointer-events-none select-none"
    >
      <i class="fas fa-user-graduate text-[15rem]"></i>
    </div>
    <div
      class="w-120 mx-auto p-6 bg-white rounded-xl shadow-md overflow-y-auto"
    >
      <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">
        <span class="text-3xl text-red-700 font-black">Student</span> Register
      </h2>

      <!-- Alert -->
      <div v-if="alert.message" class="mb-4">
        <div
          :class="`alert ${alert.type === 'success' ? 'alert-success' : 'alert-error'}`"
        >
          {{ alert.message }}
        </div>
      </div>

      <!-- User type switch -->
      <div class="flex justify-center mb-6">
        <div
          class="bg-base-200 rounded-full p-1 flex w-64 border-gray-300 border-2"
        >
          <button
            type="button"
            @click="setUserType('student')"
            :class="[
              'flex-1 py-2 rounded-full font-medium transition',
              userType === 'student'
                ? 'bg-primary text-white shadow'
                : 'text-gray-600',
            ]"
          >
            KU Student
          </button>
          <button
            type="button"
            @click="setUserType('alumni')"
            :class="[
              'flex-1 py-2 rounded-full font-medium transition',
              userType === 'alumni'
                ? 'bg-primary text-white shadow'
                : 'text-gray-600',
            ]"
          >
            Alum
          </button>
        </div>
      </div>

      <form @submit.prevent="handleRegister" class="flex flex-col gap-4">
        <!-- Email -->
        <div class="form-control flex flex-col">
          <label class="label" for="email"
            ><span class="label-text">Email</span></label
          >
          <input
            type="email"
            id="email"
            v-model="form.email"
            placeholder="Enter your email"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.email }"
          />
          <label v-if="errors.email" class="label">
            <span class="label-text-alt text-error">{{ errors.email }}</span>
          </label>
        </div>

        <!-- Passwords -->
        <div class="form-control flex flex-col">
          <PasswordField
            id="reg-password"
            label="Password"
            v-model="form.password"
            :error="errors.password"
            autocomplete="new-password"
            required
          />
        </div>

        <div class="form-control flex flex-col">
          <PasswordField
            id="reg-confirm"
            label="Confirm Password"
            v-model="form.confirmPassword"
            :error="errors.confirmPassword"
            autocomplete="new-password"
            required
          />
        </div>

        <!-- Names -->
        <div class="form-control flex flex-col">
          <label class="label" for="firstName"
            ><span class="label-text">First Name</span></label
          >
          <input
            type="text"
            id="firstName"
            v-model="form.firstName"
            placeholder="Enter your first name"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.firstName }"
          />
          <label v-if="errors.firstName" class="label">
            <span class="label-text-alt text-error">{{
              errors.firstName
            }}</span>
          </label>
        </div>

        <div class="form-control flex flex-col">
          <label class="label" for="lastName"
            ><span class="label-text">Last Name</span></label
          >
          <input
            type="text"
            id="lastName"
            v-model="form.lastName"
            placeholder="Enter your last name"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.lastName }"
          />
          <label v-if="errors.lastName" class="label">
            <span class="label-text-alt text-error">{{ errors.lastName }}</span>
          </label>
        </div>

        <!-- Transcript (alumni only) -->
        <div class="form-control flex flex-col">
          <label class="label"
            ><span class="label-text">Official Transcript (PDF)</span></label
          >
          <input
            type="file"
            ref="transcriptInput"
            accept=".pdf"
            @change="handleFileUpload"
            class="file-input file-input-bordered w-full"
          />
          <span class="text-sm mt-1">{{ transcriptFileName }}</span>
          <label v-if="errors.transcript" class="label">
            <span class="label-text-alt text-error">{{
              errors.transcript
            }}</span>
          </label>
        </div>

        <!-- Submit -->
        <div class="form-control mt-4 flex justify-center">
          <button
            type="submit"
            class="btn bg-[#56A45C] text-white hover:bg-[#44B15B] w-full"
            :disabled="loading"
            :title="
              !tosAccepted
                ? 'You must accept the Terms of Service before registering'
                : ''
            "
          >
            {{
              loading
                ? "CREATING ACCOUNT..."
                : tosAccepted
                  ? "REGISTER"
                  : "READ & ACCEPT TERMS TO REGISTER"
            }}
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

      <p class="text-center text-sm">or</p>

      <p class="text-center text-sm">
        Register as a company.
        <router-link
          to="/register-company"
          class="text-blue-500 hover:underline"
        >
          Register here
        </router-link>
      </p>
    </div>
  </div>

  <!-- Terms Modal (no slot; internal ToS) -->
  <TermsModal :show="showTos" @close="showTos = false" @accept="onAcceptTos" />
</template>

<script>
import { api } from "../../../api/client";
import { defineAsyncComponent } from "vue";

export default {
  name: "RegisterForm",
  components: {
    PasswordField: defineAsyncComponent(
      () => import("@/components/auth/PasswordField.vue"),
    ),
    TermsModal: defineAsyncComponent(
      () => import("@/components/auth/TermsModal.vue"),
    ),
  },
  emits: ["switch-to-login"],
  data() {
    return {
      userType: "student", // or 'alumni'
      form: {
        firstName: "",
        lastName: "",
        email: "",
        password: "",
        confirmPassword: "",
        transcript: null,
      },
      errors: {},
      alert: { message: "", type: "" },
      loading: false,
      transcriptFileName: "",
      // ToS state
      tosAccepted: false,
      showTos: false,
      // internal flag to know we should continue right after accept
      continueAfterTos: false,
    };
  },
  methods: {
    setUserType(type) {
      this.userType = type;
      this.errors = {};
    },

    validateEmail(email) {
      return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
    },
    validatePassword(password) {
      return password.length >= 8;
    },

    handleFileUpload(e) {
      const file = e.target.files[0];
      if (!file) return;

      if (file.type !== "application/pdf") {
        this.errors.transcript = "Only PDF files are allowed.";
        this.$refs.transcriptInput.value = "";
        return;
      }
      if (file.size > 10 * 1024 * 1024) {
        this.errors.transcript = "File must be under 10MB.";
        this.$refs.transcriptInput.value = "";
        return;
      }

      this.form.transcript = file;
      this.transcriptFileName = file.name;
      delete this.errors.transcript;
    },

    validateForm() {
      this.errors = {};
      if (!this.form.firstName.trim()) this.errors.firstName = "Required";
      if (!this.form.lastName.trim()) this.errors.lastName = "Required";
      if (!this.validateEmail(this.form.email))
        this.errors.email = "Invalid email";
      if (!this.validatePassword(this.form.password))
        this.errors.password = "Minimum 8 characters";
      if (this.form.password !== this.form.confirmPassword)
        this.errors.confirmPassword = "Passwords do not match";

      if (!this.form.transcript)
        this.errors.transcript = "Transcript is required";

      return Object.keys(this.errors).length === 0;
    },

    // Public submit handler
    async handleRegister() {
      // Ensure ToS is accepted first
      if (!this.tosAccepted) {
        this.continueAfterTos = true;
        this.showTos = true;
        return;
      }

      // Validate and proceed
      if (!this.validateForm()) return;
      await this.performRegister();
    },

    // Called after ToS accept or direct submit
    async performRegister() {
      this.loading = true;
      this.alert = { message: "Creating account...", type: "success" };

      try {
        const isAlum = this.userType === "alumni";

        const fd = new FormData();
        fd.append("email", this.form.email);
        fd.append("password", this.form.password);
        fd.append("first_name", this.form.firstName);
        fd.append("last_name", this.form.lastName);
        fd.append("is_alum", isAlum ? "true" : "false");
        if (this.form.transcript) fd.append("transcript", this.form.transcript);
        await api.post("/auth/register/student", fd, {
          headers: {},
        });

        this.alert = {
          message:
            "Registration successful! Pending admin approval for full access. Redirecting to login...",
          type: "success",
        };

        setTimeout(() => {
          this.resetForm();
          this.$router.push("/login");
        }, 2000);
      } catch (error) {
        const msg =
          error?.response?.data?.error ||
          error?.message ||
          "Registration failed.";
        this.alert = { message: msg, type: "error" };
      } finally {
        this.loading = false;
      }
    },

    onAcceptTos() {
      this.tosAccepted = true;
      this.showTos = false;

      // If the user clicked Register before accepting, continue now:
      if (this.continueAfterTos) {
        this.continueAfterTos = false;
        if (this.validateForm()) {
          this.performRegister();
        }
      }
    },

    resetForm() {
      this.form = {
        firstName: "",
        lastName: "",
        email: "",
        password: "",
        confirmPassword: "",
        transcript: null,
      };
      this.transcriptFileName = "";
      this.errors = {};
      this.alert = { message: "", type: "" };
      // keep tosAccepted true so user doesn't need to accept again in this session
    },
  },
};
</script>
