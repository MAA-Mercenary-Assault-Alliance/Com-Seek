<template>
  <div>
    <h2 class="form-title">Register</h2>
    <div v-if="alert.message" :class="`alert alert-${alert.type}`">
      {{ alert.message }}
    </div>
    
    <!-- User Type Selector -->
    <div class="user-type-selector">
      <button 
        type="button" 
        :class="['user-type-btn', { active: userType === 'student' }]"
        @click="setUserType('student')"
      >
        KU Student
      </button>
      <button 
        type="button" 
        :class="['user-type-btn', { active: userType === 'alumni' }]"
        @click="setUserType('alumni')"
      >
        Alumni
      </button>
    </div>

    <form @submit.prevent="handleRegister">
      <!-- Common Fields -->
      <div class="form-group">
        <label for="firstName">First Name</label>
        <input 
          type="text" 
          id="firstName"
          v-model="form.firstName"
          :class="{ 'error': errors.firstName }"
          required 
          placeholder="Enter your first name"
        >
        <div v-if="errors.firstName" class="error-message">{{ errors.firstName }}</div>
      </div>
      
      <div class="form-group">
        <label for="lastName">Last Name</label>
        <input 
          type="text" 
          id="lastName"
          v-model="form.lastName"
          :class="{ 'error': errors.lastName }"
          required 
          placeholder="Enter your last name"
        >
        <div v-if="errors.lastName" class="error-message">{{ errors.lastName }}</div>
      </div>
      
      <div class="form-group">
        <label for="email">Email</label>
        <input 
          type="email" 
          id="email"
          v-model="form.email"
          :class="{ 'error': errors.email }"
          required 
          placeholder="Enter your email address"
        >
        <div v-if="errors.email" class="error-message">{{ errors.email }}</div>
      </div>
      
      <div class="form-group">
        <label for="password">Password</label>
        <input 
          type="password" 
          id="password"
          v-model="form.password"
          :class="{ 'error': errors.password }"
          required 
          placeholder="Enter a strong password"
          minlength="8"
        >
        <div v-if="errors.password" class="error-message">{{ errors.password }}</div>
      </div>
      
      <div class="form-group">
        <label for="confirmPassword">Confirm Password</label>
        <input 
          type="password" 
          id="confirmPassword"
          v-model="form.confirmPassword"
          :class="{ 'error': errors.confirmPassword }"
          required 
          placeholder="Confirm your password"
        >
        <div v-if="errors.confirmPassword" class="error-message">{{ errors.confirmPassword }}</div>
      </div>
      
      <div class="form-group">
        <label for="description">About Me</label>
        <textarea 
          id="description"
          v-model="form.description"
          placeholder="Tell us about yourself, your skills, and career interests..."
        ></textarea>
      </div>

      <!-- Alumni-specific fields -->
      <template v-if="userType === 'alumni'">
        <div class="form-group">
          <label for="graduationYear">Graduation Year</label>
          <select 
            id="graduationYear"
            v-model="form.graduationYear"
            :class="{ 'error': errors.graduationYear }"
          >
            <option value="">Select graduation year</option>
            <option v-for="year in graduationYears" :key="year" :value="year">
              {{ year }}
            </option>
          </select>
          <div v-if="errors.graduationYear" class="error-message">{{ errors.graduationYear }}</div>
        </div>
        
        <div class="form-group">
          <label for="transcript">Official Transcript (PDF)</label>
          <div class="file-upload">
            <input 
              type="file" 
              id="transcript"
              ref="transcriptInput"
              accept=".pdf"
              @change="handleFileUpload"
            >
            <div class="file-upload-btn">
              {{ transcriptFileName || '📄 Click to upload transcript (PDF only)' }}
            </div>
          </div>
          <div v-if="errors.transcript" class="error-message">{{ errors.transcript }}</div>
        </div>
      </template>

      <!-- Student-specific fields -->
      <template v-if="userType === 'student'">
        <div class="form-group">
          <label for="studentId">Student ID</label>
          <input 
            type="text" 
            id="studentId"
            v-model="form.studentId"
            :class="{ 'error': errors.studentId }"
            required 
            placeholder="Enter your KU student ID (e.g., 6610545391)"
          >
          <div v-if="errors.studentId" class="error-message">{{ errors.studentId }}</div>
        </div>
        
        <div class="form-group">
          <label for="currentYear">Current Year</label>
          <select 
            id="currentYear"
            v-model="form.currentYear"
            :class="{ 'error': errors.currentYear }"
            required
          >
            <option value="">Select your current year</option>
            <option value="1">1st Year</option>
            <option value="2">2nd Year</option>
            <option value="3">3rd Year</option>
            <option value="4">4th Year</option>
            <option value="5">5th Year</option>
          </select>
          <div v-if="errors.currentYear" class="error-message">{{ errors.currentYear }}</div>
        </div>
      </template>
      
      <button type="submit" class="btn" :disabled="loading">
        {{ loading ? 'CREATING ACCOUNT...' : 'SIGN UP' }}
      </button>
    </form>
    
    <div class="form-switch">
      <p>Already have an account? <a href="#" @click.prevent="$emit('switch-to-login')">Login here</a></p>
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
        description: '',
        // Student fields
        studentId: '',
        currentYear: '',
        // Alumni fields
        graduationYear: '',
        transcript: null
      },
      errors: {},
      alert: {
        message: '',
        type: ''
      },
      loading: false,
      transcriptFileName: ''
    }
  },
  computed: {
    graduationYears() {
      const currentYear = new Date().getFullYear();
      const years = [];
      for (let year = currentYear + 1; year >= 1990; year--) {
        years.push(year);
      }
      return years;
    }
  },
  methods: {
    setUserType(type) {
      this.userType = type;
      this.errors = {}; // Clear errors when switching
    },
    
    validateEmail(email) {
      const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      return re.test(email);
    },
    
    validatePassword(password) {
      return password.length >= 8;
    },
    
    validateStudentId(studentId) {
      const re = /^66\d{8}$/;
      return re.test(studentId);
    },
    
    handleFileUpload(event) {
      const file = event.target.files[0];
      
      if (file) {
        if (file.type !== 'application/pdf') {
          this.errors.transcript = 'Please upload a PDF file only.';
          this.$refs.transcriptInput.value = '';
          this.transcriptFileName = '';
          return;
        }
        
        if (file.size > 10 * 1024 * 1024) { // 10MB limit
          this.errors.transcript = 'File size must be less than 10MB.';
          this.$refs.transcriptInput.value = '';
          this.transcriptFileName = '';
          return;
        }
        
        this.form.transcript = file;
        this.transcriptFileName = `📄 ${file.name} (${(file.size / 1024 / 1024).toFixed(2)}MB)`;
        delete this.errors.transcript;
      }
    },
    
    validateForm() {
      this.errors = {};
      
      // Common validations
      if (!this.form.firstName.trim()) {
        this.errors.firstName = 'First name is required.';
      }
      
      if (!this.form.lastName.trim()) {
        this.errors.lastName = 'Last name is required.';
      }
      
      if (!this.validateEmail(this.form.email)) {
        this.errors.email = 'Please enter a valid email address.';
      }
      
      if (!this.validatePassword(this.form.password)) {
        this.errors.password = 'Password must be at least 8 characters long.';
      }
      
      if (this.form.password !== this.form.confirmPassword) {
        this.errors.confirmPassword = 'Passwords do not match.';
      }
      
      // User type specific validations
      if (this.userType === 'student') {
        if (!this.validateStudentId(this.form.studentId)) {
          this.errors.studentId = 'Please enter a valid KU student ID (10 digits starting with 66).';
        }
        
        if (!this.form.currentYear) {
          this.errors.currentYear = 'Please select your current year.';
        }
      } else if (this.userType === 'alumni') {
        if (!this.form.graduationYear) {
          this.errors.graduationYear = 'Please select your graduation year.';
        }
        
        if (!this.form.transcript) {
          this.errors.transcript = 'Please upload your official transcript.';
        }
      }
      
      return Object.keys(this.errors).length === 0;
    },
    
    async handleRegister() {
      if (!this.validateForm()) return;
      
      this.loading = true;
      this.alert = { message: 'Creating your account...', type: 'success' };
      
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 2000));
        
        // In a real app, you would call your API here:
        // const formData = new FormData();
        // Object.keys(this.form).forEach(key => {
        //   if (this.form[key]) formData.append(key, this.form[key]);
        // });
        // formData.append('userType', this.userType);
        // const response = await this.$api.register(formData);
        
        if (this.userType === 'alumni') {
          this.alert = { 
            message: 'Registration successful! Your account is pending approval. You will receive an email notification once approved.', 
            type: 'success' 
          };
        } else {
          this.alert = { 
            message: 'Registration successful! You can now login with your credentials.', 
            type: 'success' 
          };
        }
        
        // Reset form and switch to login after delay
        setTimeout(() => {
          this.resetForm();
          this.$emit('switch-to-login');
        }, 3000);
        
      } catch (error) {
        this.alert = { message: 'Registration failed. Please try again.', type: 'error' };
      } finally {
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
        description: '',
        studentId: '',
        currentYear: '',
        graduationYear: '',
        transcript: null
      };
      this.transcriptFileName = '';
      this.errors = {};
      this.alert = { message: '', type: '' };
    }
  }
}
</script>

<style scoped>
.form-title {
  text-align: center;
  color: #333;
  font-size: 24px;
  margin-bottom: 30px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  color: #555;
  font-size: 14px;
  margin-bottom: 5px;
  font-weight: 500;
}

.form-group input, 
.form-group textarea, 
.form-group select {
  width: 100%;
  padding: 12px 15px;
  border: 2px solid #e1e1e1;
  border-radius: 8px;
  font-size: 14px;
  transition: all 0.3s ease;
  background-color: #fafafa;
}

.form-group input:focus, 
.form-group textarea:focus, 
.form-group select:focus {
  outline: none;
  border-color: #4682B4;
  background-color: white;
  box-shadow: 0 0 0 3px rgba(70, 130, 180, 0.1);
}

.form-group input.error,
.form-group textarea.error,
.form-group select.error {
  border-color: #dc3545;
}

.form-group textarea {
  min-height: 80px;
  resize: vertical;
}

.file-upload {
  position: relative;
  display: inline-block;
  cursor: pointer;
  width: 100%;
}

.file-upload input[type=file] {
  position: absolute;
  opacity: 0;
  width: 100%;
  height: 100%;
  cursor: pointer;
}

.file-upload-btn {
  display: block;
  width: 100%;
  padding: 12px 15px;
  background: #f8f9fa;
  border: 2px dashed #ccc;
  border-radius: 8px;
  text-align: center;
  color: #666;
  transition: all 0.3s ease;
}

.file-upload:hover .file-upload-btn {
  border-color: #4682B4;
  color: #4682B4;
}

.btn {
  width: 100%;
  padding: 12px;
  background: linear-gradient(135deg, #4682B4 0%, #2c5aa0 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 10px;
}

.btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(70, 130, 180, 0.4);
}

.btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.form-switch {
  text-align: center;
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid #eee;
}

.form-switch a {
  color: #4682B4;
  text-decoration: none;
  font-weight: 500;
}

.form-switch a:hover {
  text-decoration: underline;
}

.user-type-selector {
  display: flex;
  background: #f8f9fa;
  border-radius: 8px;
  padding: 4px;
  margin-bottom: 30px;
}

.user-type-btn {
  flex: 1;
  padding: 10px;
  background: transparent;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-weight: 500;
}

.user-type-btn.active {
  background: white;
  color: #4682B4;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.error-message {
  color: #dc3545;
  font-size: 12px;
  margin-top: 5px;
}

.alert {
  padding: 12px;
  border-radius: 8px;
  margin-bottom: 20px;
}

.alert-success {
  background: #d4edda;
  color: #155724;
  border: 1px solid #c3e6cb;
}

.alert-error {
  background: #f8d7da;
  color: #721c24;
  border: 1px solid #f5c6cb;
}
</style>