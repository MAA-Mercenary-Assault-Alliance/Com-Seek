<template>
  <div class="min-h-screen pt-10">
  <div class="w-120 mx-auto p-6 bg-white rounded-xl shadow-md overflow-y-auto">
    <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">Register</h2>

    <!-- Alert -->
    <div v-if="alert.message" class="mb-4">
      <div :class="`alert ${alert.type === 'success' ? 'alert-success' : 'alert-error'}`">
        {{ alert.message }}
      </div>
    </div>

    <div class="flex justify-center mb-6">
      <label class="flex items-center cursor-pointer">
        <span class="mr-3 font-medium" :class="{ 'text-primary': userType === 'student' }">KU Student</span>
        <input type="checkbox" class="toggle toggle-primary" 
              :checked="userType === 'alumni'" 
              @change="setUserType($event.target.checked ? 'alumni' : 'student')" />
        <span class="ml-3 font-medium" :class="{ 'text-primary': userType === 'alumni' }">Alumni</span>
      </label>
    </div>



    <form @submit.prevent="handleRegister" class="flex flex-col gap-4">

      <!-- Always on top -->
      <div class="form-control flex flex-col">
        <label class="label" for="email"><span class="label-text">Email</span></label>
        <input type="email" id="email" v-model="form.email" placeholder="Enter your email" 
               :class="errors.email ? 'input input-bordered input-error' : 'input input-bordered w-full'" />
        <label v-if="errors.email" class="label"><span class="label-text-alt text-error">{{ errors.email }}</span></label>
      </div>

      <div class="form-control flex flex-col">
        <label class="label" for="password"><span class="label-text">Password</span></label>
        <input type="password" id="password" v-model="form.password" placeholder="Enter a strong password" 
               :class="errors.password ? 'input input-bordered input-error' : 'input input-bordered w-full'" />
        <label v-if="errors.password" class="label"><span class="label-text-alt text-error">{{ errors.password }}</span></label>
      </div>

      <div class="form-control flex flex-col">
        <label class="label" for="confirmPassword"><span class="label-text">Confirm Password</span></label>
        <input type="password" id="confirmPassword" v-model="form.confirmPassword" placeholder="Confirm your password" 
               :class="errors.confirmPassword ? 'input input-bordered input-error' : 'input input-bordered w-full'" />
        <label v-if="errors.confirmPassword" class="label"><span class="label-text-alt text-error">{{ errors.confirmPassword }}</span></label>
      </div>

      <!-- Common Fields -->
      <div class="form-control flex flex-col">
        <label class="label" for="firstName"><span class="label-text">First Name</span></label>
        <input type="text" id="firstName" v-model="form.firstName" placeholder="Enter your first name" 
               :class="errors.firstName ? 'input input-bordered input-error' : 'input input-bordered w-full'" />
        <label v-if="errors.firstName" class="label"><span class="label-text-alt text-error">{{ errors.firstName }}</span></label>
      </div>

      <div class="form-control flex flex-col">
        <label class="label" for="lastName"><span class="label-text">Last Name</span></label>
        <input type="text" id="lastName" v-model="form.lastName" placeholder="Enter your last name" 
               :class="errors.lastName ? 'input input-bordered input-error' : 'input input-bordered w-full'" />
        <label v-if="errors.lastName" class="label"><span class="label-text-alt text-error">{{ errors.lastName }}</span></label>
      </div>

      <!-- Transcript Fields -->
        <div class="form-control flex flex-col">
          <label class="label"><span class="label-text">Official Transcript (PDF)</span></label>
          <input type="file" ref="transcriptInput" accept=".pdf" @change="handleFileUpload" class="file-input file-input-bordered w-full" />
          <span class="text-sm mt-1">{{ transcriptFileName }}</span>
          <label v-if="errors.transcript" class="label"><span class="label-text-alt text-error">{{ errors.transcript }}</span></label>
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
        transcript: null
      },
      errors: {},
      alert: { message: '', type: '' },
      loading: false,
      transcriptFileName: ''
    }
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

      if (file.type !== 'application/pdf') { 
        this.errors.transcript = 'Only PDF allowed'; 
        this.$refs.transcriptInput.value = ''; 
        return; 
      }
      if (file.size > 10 * 1024 * 1024) { 
        this.errors.transcript = 'Max 10MB'; 
        this.$refs.transcriptInput.value = ''; 
        return; 
      }

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

      // Alumni still requires transcript
      if (this.userType === 'alumni' && !this.form.transcript) {
        this.errors.transcript = 'Upload transcript';
      }

      return Object.keys(this.errors).length === 0;
    },
    async handleRegister() {
      if (!this.validateForm()) return;
      this.loading = true;
      this.alert = { message: 'Creating account...', type: 'success' };
      try {
        await new Promise(r => setTimeout(r, 2000));
        this.alert = { 
          message: this.userType === 'alumni' 
            ? 'Registration successful! Pending approval.' 
            : 'Registration successful! You can login now.', 
          type: 'success' 
        };
        setTimeout(() => { 
          this.resetForm(); 
          this.$emit('switch-to-login'); 
        }, 2000);
      } catch { 
        this.alert = { message: 'Registration failed', type: 'error' }; 
      }
      finally { 
        this.loading = false; 
      }
    },
    resetForm() {
      this.form = { 
        firstName: '', 
        lastName: '', 
        email: '', 
        password: '', 
        confirmPassword: '', 
        transcript: null 
      };
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
