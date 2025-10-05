<template>
  <div class="pt-10 pb-32">
  <div class="w-120 mx-auto p-6 bg-white rounded-xl shadow-md overflow-y-auto">
    <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">Register</h2>

    <!-- Alert -->
    <div v-if="alert.message" class="mb-4">
      <div :class="`alert ${alert.type === 'success' ? 'alert-success' : 'alert-error'}`">
        {{ alert.message }}
      </div>
    </div>

    <div class="flex justify-center mb-6">
      <div class="bg-base-200 rounded-full p-1 flex w-64 border-gray-300 border-2">
        <!-- Student -->
        <button
          type="button"
          @click="setUserType('student')"
          :class="[
            'flex-1 py-2 rounded-full font-medium transition',
            userType === 'student' ? 'bg-primary text-white shadow' : 'text-gray-600'
          ]"
        >
          KU Student
        </button>

        <!-- Alumni -->
        <button
          type="button"
          @click="setUserType('alumni')"
          :class="[
            'flex-1 py-2 rounded-full font-medium transition',
            userType === 'alumni' ? 'bg-primary text-white shadow' : 'text-gray-600'
          ]"
        >
          Alum
        </button>
      </div>
    </div>

    <form @submit.prevent="handleRegister" class="flex flex-col gap-4">
      <!-- User Type Specific Fields -->
        <div class="form-control flex flex-col">
          <label class="label" for="email"><span class="label-text">Email</span></label>
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

        <div class="form-control flex flex-col">
          <label class="label" for="password"><span class="label-text">Password</span></label>
          <input
            type="password"
            id="password"
            v-model="form.password"
            placeholder="Enter a strong password"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.password }"
          />
          <label v-if="errors.password" class="label">
            <span class="label-text-alt text-error">{{ errors.password }}</span>
          </label>
        </div>

        <div class="form-control flex flex-col">
          <label class="label" for="confirmPassword"><span class="label-text">Confirm Password</span></label>
          <input
            type="password"
            id="confirmPassword"
            v-model="form.confirmPassword"
            placeholder="Confirm your password"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.confirmPassword }"
          />
          <label v-if="errors.confirmPassword" class="label">
            <span class="label-text-alt text-error">{{ errors.confirmPassword }}</span>
          </label>
        </div>

        <!-- Common Fields -->
        <div class="form-control flex flex-col">
          <label class="label" for="firstName"><span class="label-text">First Name</span></label>
          <input
            type="text"
            id="firstName"
            v-model="form.firstName"
            placeholder="Enter your first name"
            class="input input-bordered w-full"
            :class="{ 'input-error': errors.firstName }"
          />
          <label v-if="errors.firstName" class="label">
            <span class="label-text-alt text-error">{{ errors.firstName }}</span>
          </label>
        </div>

        <div class="form-control flex flex-col">
          <label class="label" for="lastName"><span class="label-text">Last Name</span></label>
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

      <!-- Transcript Fields -->
        <div class="form-control flex flex-col">
          <label class="label"><span class="label-text">Official Transcript (PDF)</span></label>
          <input type="file" ref="transcriptInput" accept=".pdf" @change="handleFileUpload" class="file-input file-input-bordered w-full" />
          <span class="text-sm mt-1">{{ transcriptFileName }}</span>
          <label v-if="errors.transcript" class="label"><span class="label-text-alt text-error">{{ errors.transcript }}</span></label>
        </div>


      <!-- Submit Button -->
      <div class="form-control mt-4 flex justify-center">
        <button type="submit" class="btn btn-primary w-full" :disabled="loading">
          {{ loading ? 'CREATING ACCOUNT...' : 'REGISTER' }}
        </button>
      </div>
    </form>

    <!-- Switch to Login -->
<p class="text-center text-sm mt-4 border-t pt-4">
  Already have an account? 
  <router-link 
    to="/login" 
    class="text-blue-500 hover:underline"
  >
    Login here
  </router-link>
</p>
  </div>
  </div>
</template>

<script>
import { api } from '../../../api/client';

export default {
  name: "RegisterForm",
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
    };
  },
  methods: {
    // Toggle student/alumni
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

      if (this.userType === "alumni" && !this.form.transcript) {
        this.errors.transcript = "Transcript required for alumni.";
      }

      return Object.keys(this.errors).length === 0;
    },

    async handleRegister() {
      if (!this.validateForm()) return;
      this.loading = true;
      this.alert = { message: "Creating account...", type: "success" };

      try {
        // prepare payload
        const isAlum = this.userType === "alumni";
        const body = {
          email: this.form.email,
          password: this.form.password,
          user_type: "student", // backend expects 'student' or 'company'
          student_profile: {
            first_name: this.form.firstName,
            last_name: this.form.lastName,
            is_alum: isAlum,
          },
        };

        // Send request
        const { data } = await api.post("/auth/register", body);
        console.log("Registered:", data);

        this.alert = {
          message: "Registration successful! Pending admin approval.",
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
    },
  },
};
</script>
