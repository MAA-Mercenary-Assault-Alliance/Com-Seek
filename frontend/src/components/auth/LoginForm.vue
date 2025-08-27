<template>
  <div>
    <h2 class="form-title">Login</h2>
    <div v-if="alert.message" :class="`alert alert-${alert.type}`">
      {{ alert.message }}
    </div>
    
    <form @submit.prevent="handleLogin">
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
          placeholder="Enter your password"
        >
        <div v-if="errors.password" class="error-message">{{ errors.password }}</div>
      </div>
      
      <button type="submit" class="btn" :disabled="loading">
        {{ loading ? 'LOGGING IN...' : 'LOGIN' }}
      </button>
    </form>
    
    <div class="form-switch">
      <p>Don't have an account? <a href="#" @click.prevent="$emit('switch-to-register')">Sign up here</a></p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LoginForm',
  emits: ['switch-to-register'],
  data() {
    return {
      form: {
        email: '',
        password: ''
      },
      errors: {},
      alert: {
        message: '',
        type: ''
      },
      loading: false
    }
  },
  methods: {
    validateEmail(email) {
      const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      return re.test(email);
    },
    
    validateForm() {
      this.errors = {};
      
      if (!this.validateEmail(this.form.email)) {
        this.errors.email = 'Please enter a valid email address.';
      }
      
      if (!this.form.password) {
        this.errors.password = 'Password is required.';
      }
      
      return Object.keys(this.errors).length === 0;
    },
    
    async handleLogin() {
      if (!this.validateForm()) return;
      
      this.loading = true;
      this.alert = { message: 'Logging in...', type: 'success' };
      
      try {
        // Simulate API call
        await new Promise(resolve => setTimeout(resolve, 1500));
        
        this.alert = { message: 'Login successful! Redirecting...', type: 'success' };
        
        setTimeout(() => {
          this.$router.push('/dashboard');
        }, 1000);
        
      } catch (error) {
        this.alert = { message: 'Invalid credentials. Please try again.', type: 'error' };
      } finally {
        this.loading = false;
      }
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

.form-group input {
  width: 100%;
  padding: 12px 15px;
  border: 2px solid #e1e1e1;
  border-radius: 8px;
  font-size: 14px;
  transition: all 0.3s ease;
  background-color: #fafafa;
}

.form-group input:focus {
  outline: none;
  border-color: #4682B4;
  background-color: white;
  box-shadow: 0 0 0 3px rgba(70, 130, 180, 0.1);
}

.form-group input.error {
  border-color: #dc3545;
}

.error-message {
  color: #dc3545;
  font-size: 12px;
  margin-top: 5px;
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
