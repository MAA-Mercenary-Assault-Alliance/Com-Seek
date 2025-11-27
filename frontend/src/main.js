import { createApp } from 'vue'
import App from './App.vue'
import Router from './router'
import '@fortawesome/fontawesome-free/css/all.min.css'

import VueTelInput from 'vue-tel-input'
import 'vue-tel-input/vue-tel-input.css'   // <-- FIXED PATH

const app = createApp(App)

app.use(VueTelInput)
app.use(Router)

app.mount('#app')
