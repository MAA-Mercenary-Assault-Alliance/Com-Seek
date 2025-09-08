<template>
  <div class="w-120 h-[1000px] mx-auto p-6 bg-white rounded-xl shadow-md overflow-y-auto">
    <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">Register</h2>

    <!-- Alert -->
    <div v-if="alert.message" class="mb-4">
      <div :class="`alert ${alert.type === 'success' ? 'alert-success' : 'alert-error'}`">
        {{ alert.message }}
      </div>
    </div>

    <!-- User Type Selector -->
    <div class="btn-group w-full mb-6">
      <button 
        type="button"
        :class="['btn', userType === 'student' ? 'btn-active' : '']"
        @click="setUserType('student')"
      >
        KU Student
      </button>
      <button 
        type="button"
        :class="['btn', userType === 'alumni' ? 'btn-active' : '']"
        @click="setUserType('alumni')"
      >
        Alumni
      </button>
    </div>

    <form @submit.prevent="handleRegister" class="flex flex-col gap-4">

      <!-- Always on top -->
      <div class="form-control flex flex-col">
        <label class="label" for="email"><span class="label-text">Email</span></label>
        <input type="email" id="email" v-model="form.email" placeholder="Enter your email" 
               :class="errors.email ? 'input input-bordered input-error' : 'input input-bordered'" />
        <label v-if="errors.email" class="label"><span class="label-text-alt text-error">{{ errors.email }}</span></label>
      </div>

      <div class="form-control flex flex-col">
        <label class="label" for="password"><span class="label-text">Password</span></label>
        <input type="password" id="password" v-model="form.password" placeholder="Enter a strong password" 
               :class="errors.password ? 'input input-bordered input-error' : 'input input-bordered'" />
        <label v-if="errors.password" class="label"><span class="label-text-alt text-error">{{ errors.password }}</span></label>
      </div>

      <div class="form-control flex flex-col">
        <label class="label" for="confirmPassword"><span class="label-text">Confirm Password</span></label>
        <input type="password" id="confirmPassword" v-model="form.confirmPassword" placeholder="Confirm your password" 
               :class="errors.confirmPassword ? 'input input-bordered input-error' : 'input input-bordered'" />
        <label v-if="errors.confirmPassword" class="label"><span class="label-text-alt text-error">{{ errors.confirmPassword }}</span></label>
      </div>

      <!-- Common Fields -->
      <div class="form-control flex flex-col">
        <label class="label" for="firstName"><span class="label-text">First Name</span></label>
        <input type="text" id="firstName" v-model="form.firstName" placeholder="Enter your first name" 
               :class="errors.firstName ? 'input input-bordered input-error' : 'input input-bordered'" />
        <label v-if="errors.firstName" class="label"><span class="label-text-alt text-error">{{ errors.firstName }}</span></label>
      </div>

      <div class="form-control flex flex-col">
        <label class="label" for="lastName"><span class="label-text">Last Name</span></label>
        <input type="text" id="lastName" v-model="form.lastName" placeholder="Enter your last name" 
               :class="errors.lastName ? 'input input-bordered input-error' : 'input input-bordered'" />
        <label v-if="errors.lastName" class="label"><span class="label-text-alt text-error">{{ errors.lastName }}</span></label>
      </div>

      <div class="form-control flex flex-col">
        <label class="label" for="description"><span class="label-text">About Me</span></label>
        <textarea id="description" v-model="form.description" placeholder="Tell us about yourself..." 
                  class="textarea textarea-bordered"></textarea>
      </div>

      <!-- Student Fields -->
      <div v-show="userType === 'student'" class="flex flex-col gap-4 mt-4 p-4 border rounded-lg bg-gray-50">
        <div class="form-control flex flex-col">
          <label class="label" for="studentId"><span class="label-text">Student ID</span></label>
          <input type="text" id="studentId" v-model="form.studentId" placeholder="Enter your KU student ID" 
                 :class="errors.studentId ? 'input input-bordered input-error' : 'input input-bordered'" />
          <label v-if="errors.studentId" class="label"><span class="label-text-alt text-error">{{ errors.studentId }}</span></label>
        </div>

        <div class="form-control flex flex-col">
          <label class="label" for="currentYear"><span class="label-text">Current Year</span></label>
          <select id="currentYear" v-model="form.currentYear" 
                  :class="errors.currentYear ? 'select select-bordered select-error' : 'select select-bordered'">
            <option value="">Select your current year</option>
            <option value="1">1st Year</option>
            <option value="2">2nd Year</option>
            <option value="3">3rd Year</option>
            <option value="4">4th Year</option>
            <option value="5">5th Year or More</option>
          </select>
          <label v-if="errors.currentYear" class="label"><span class="label-text-alt text-error">{{ errors.currentYear }}</span></label>
        </div>
      </div>

      <!-- Alumni Fields -->
      <div v-show="userType === 'alumni'" class="flex flex-col gap-4 mt-4 p-4 border rounded-lg bg-gray-50">
        <div class="form-control flex flex-col">
          <label class="label" for="graduationYear"><span class="label-text">Graduation Year</span></label>
          <select id="graduationYear" v-model="form.graduationYear" 
                  :class="errors.graduationYear ? 'select select-bordered select-error' : 'select select-bordered'">
            <option value="">Select graduation year</option>
            <option v-for="year in graduationYears" :key="year" :value="year">{{ year }}</option>
          </select>
          <label v-if="errors.graduationYear" class="label"><span class="label-text-alt text-error">{{ errors.graduationYear }}</span></label>
        </div>

        <div class="form-control flex flex-col">
          <label class="label"><span class="label-text">Official Transcript (PDF)</span></label>
          <input type="file" ref="transcriptInput" accept=".pdf" @change="handleFileUpload" class="file-input file-input-bordered w-full" />
          <span class="text-sm mt-1">{{ transcriptFileName }}</span>
          <label v-if="errors.transcript" class="label"><span class="label-text-alt text-error">{{ errors.transcript }}</span></label>
        </div>
      </div>

      <!-- Submit Button -->
      <div class="form-control mt-4">
        <button type="submit" class="btn btn-primary" :disabled="loading">
          {{ loading ? 'CREATING ACCOUNT...' : 'SIGN UP' }}
        </button>
      </div>
    </form>

    <!-- Switch to Login -->
<p class="text-center mt-4 text-sm">
  Already have an account? 
  <router-link 
    to="/login" 
    class="text-blue-500 hover:underline"
  >
    Login here
  </router-link>
</p>
  </div>
</template>

<script>
export default {
  name: 'RegisterForm',
  emits: ['switch-to-login'],
  data() {
    return {
      userType: 'student',
      form: {
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        confirmPassword: '',
        description: '',
        studentId: '',
        currentYear: '',
        graduationYear: '',
        transcript: null
      },
      errors: {},
      alert: { message: '', type: '' },
      loading: false,
      transcriptFileName: ''
    }
  },
  computed: {
    graduationYears() {
      const currentYear = new Date().getFullYear();
      const years = [];
      for (let y = currentYear + 1; y >= 1990; y--) years.push(y);
      return years;
    }
  },
  methods: {
    setUserType(type) { this.userType = type; this.errors = {}; },
    validateEmail(email) { return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email); },
    validatePassword(password) { return password.length >= 8; },
    validateStudentId(studentId) { return /^66\d{8}$/.test(studentId); },
    handleFileUpload(e) {
      const file = e.target.files[0];
      if (!file) return;
      if (file.type !== 'application/pdf') { this.errors.transcript = 'Only PDF allowed'; this.$refs.transcriptInput.value = ''; return; }
      if (file.size > 10 * 1024 * 1024) { this.errors.transcript = 'Max 10MB'; this.$refs.transcriptInput.value = ''; return; }
      this.form.transcript = file;
      this.transcriptFileName = file.name;
      delete this.errors.transcript;
    },
    validateForm() {
      this.errors = {};
      if (!this.form.firstName.trim()) this.errors.firstName = 'Required';
      if (!this.form.lastName.trim()) this.errors.lastName = 'Required';
      if (!this.validateEmail(this.form.email)) this.errors.email = 'Invalid email';
      if (!this.validatePassword(this.form.password)) this.errors.password = 'Min 8 chars';
      if (this.form.password !== this.form.confirmPassword) this.errors.confirmPassword = 'Passwords do not match';

      if (this.userType === 'student') {
        if (!this.validateStudentId(this.form.studentId)) this.errors.studentId = 'Invalid KU ID';
        if (!this.form.currentYear) this.errors.currentYear = 'Select your year';
      } else if (this.userType === 'alumni') {
        if (!this.form.graduationYear) this.errors.graduationYear = 'Select year';
        if (!this.form.transcript) this.errors.transcript = 'Upload transcript';
      }
      return Object.keys(this.errors).length === 0;
    },
    async handleRegister() {
      if (!this.validateForm()) return;
      this.loading = true;
      this.alert = { message: 'Creating account...', type: 'success' };
      try {
        await new Promise(r => setTimeout(r, 2000));
        this.alert = { message: this.userType === 'alumni' 
          ? 'Registration successful! Pending approval.' 
          : 'Registration successful! You can login now.', 
          type: 'success' 
        };
        setTimeout(() => { this.resetForm(); this.$emit('switch-to-login'); }, 2000);
      } catch { this.alert = { message: 'Registration failed', type: 'error' }; }
      finally { this.loading = false; }
    },
    resetForm() {
      this.form = { firstName:'', lastName:'', email:'', password:'', confirmPassword:'', description:'', studentId:'', currentYear:'', graduationYear:'', transcript:null };
      this.transcriptFileName = '';
      this.errors = {};
      this.alert = { message:'', type:'' };
    }
  }
}
</script>

<style>
.form-control { transition: all 0.3s ease; }
</style>
