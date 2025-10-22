<template>
  <div class="w-full h-full mx-auto bg-white rounded-xl shadow-md p-4 md:p-8">
    <!-- Title + Edit -->
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-xl font-bold text-center text-gray-800 mb-6">
        <span class="text-2xl text-green-700 font-black">Company</span> Details
      </h2>
      <button
        v-if="canEdit && !localEdit"
        class="btn btn-sm btn-outline"
        @click="startEdit"
      >
        Edit
      </button>
    </div>

    <!-- READ VIEW — ALWAYS VISIBLE -->
    <div class="grid md:grid-cols-2 gap-6">
      <div>
        <div class="mb-3">
          <div class="text-sm text-gray-500">Name</div>
          <div class="font-medium">{{ info.name || '-' }}</div>
        </div>

        <div class="mb-3">
          <div class="text-sm text-gray-500">Website</div>
          <div class="font-medium">
            <a
              v-if="info.website"
              :href="info.website"
              target="_blank"
              rel="noopener"
              class="link link-primary break-all"
            >{{ info.website }}</a>
            <span v-else>-</span>
          </div>
        </div>

        <div class="mb-3">
          <div class="text-sm text-gray-500">Location</div>
          <div class="font-medium">{{ info.location || '-' }}</div>
        </div>
      </div>

      <div>
        <div class="mb-3">
          <div class="text-sm text-gray-500">Contact Email</div>
          <div class="font-medium">{{ info.contact_email || '-' }}</div>
        </div>

        <div class="mb-3">
          <div class="text-sm text-gray-500">Contact Number</div>
          <div class="font-medium whitespace-pre-wrap">{{ info.contact_number || '-' }}</div>
        </div>

        <div class="mb-3">
          <div class="text-sm text-gray-500">Description</div>
          <div class="font-medium whitespace-pre-wrap">{{ info.description || '-' }}</div>
        </div>
      </div>
    </div>

    <!-- EDIT FORM — SHOWS BELOW -->
    <form
      v-show="localEdit"
      class="grid md:grid-cols-2 gap-6 mt-6 border-t pt-6"
      @submit.prevent="submit"
      novalidate
    >
      <div>
        <label class="label"><span class="label-text">Name</span></label>
        <input v-model.trim="form.name" class="input input-bordered w-full" placeholder="Acme Robotics" />

        <label class="label mt-4"><span class="label-text">Website</span></label>
        <input v-model.trim="form.website" class="input input-bordered w-full" placeholder="https://…" />

        <label class="label mt-4"><span class="label-text">Location</span></label>
        <input v-model.trim="form.location" class="input input-bordered w-full" placeholder="Bangkok" />
      </div>

      <div>
        <!-- EMAIL -->
        <label class="label"><span class="label-text">Contact Email</span></label>
        <input
          v-model.trim="form.contact_email"
          type="email"
          class="input input-bordered w-full"
          :class="emailError ? 'input-error' : ''"
          placeholder="hr@acme.co"
          @input="validateEmail"
          @blur="validateEmail"
        />
        <label v-if="emailError" class="label">
          <span class="label-text-alt text-error">{{ emailError }}</span>
        </label>

        <!-- PHONE -->
        <label class="label mt-4"><span class="label-text">Contact Number</span></label>
        <input
          v-model.trim="form.contact_number"
          type="tel"
          inputmode="numeric"
          class="input input-bordered w-full"
          :class="phoneError ? 'input-error' : ''"
          placeholder="0804326754"
          @input="validatePhone"
          @blur="validatePhone"
        />
        <label v-if="phoneError" class="label">
          <span class="label-text-alt text-error">{{ phoneError }}</span>
        </label>

        <label class="label mt-4"><span class="label-text">Description</span></label>
        <textarea
          v-model.trim="form.description"
          class="textarea textarea-bordered w-full"
          rows="5"
          placeholder="About the company…"
        ></textarea>
      </div>

      <div class="md:col-span-2 flex items-center gap-2 mt-2">
        <button
          :disabled="isSaveDisabled"
          class="btn btn-primary bg-green-600 border-green-600 hover:bg-green-700 hover:border-green-700"
        >

          <span v-if="isSaving" class="loading loading-spinner loading-sm mr-2"></span>
          Save
        </button>
        <button type="button" class="btn btn-ghost" @click="cancelEdit">Cancel</button>

        <span
          v-if="alert?.message"
          :class="alert?.type === 'error' ? 'text-red-600' : 'text-green-700'"
        >
          {{ alert.message }}
        </span>
      </div>
    </form>
  </div>
</template>

<script setup>
import { reactive, watch, computed, ref } from 'vue'

const props = defineProps({
  profile: { type: Object, default: null },
  isSaving: { type: Boolean, default: false },
  canEdit: { type: Boolean, default: false },
  alert: { type: Object, default: () => ({ type: '', message: '' }) }
})
const emit = defineEmits(['save'])

function pick(obj, ...keys) {
  for (const k of keys) if (obj && obj[k] != null && obj[k] !== '') return obj[k]
  return ''
}
const info = computed(() => {
  const p = props.profile || {}
  return {
    name: pick(p, 'name', 'Name'),
    website: pick(p, 'website', 'Website'),
    location: pick(p, 'location', 'Location'),
    contact_email: pick(p, 'contact_email', 'ContactEmail'),
    contact_number: pick(p, 'contact_number', 'ContactNumber'),
    description: pick(p, 'description', 'Description'),
  }
})

const localEdit = ref(false)
const form = reactive({
  name: '',
  website: '',
  location: '',
  contact_email: '',
  contact_number: '',
  description: ''
})

const phoneError = ref('')
const emailError = ref('')

function validatePhone() {
  const val = form.contact_number.trim()
  if (!val) {
    phoneError.value = 'Phone number is required.'
  } else if (!/^[0-9]+$/.test(val)) {
    phoneError.value = 'Phone number must contain only digits.'
  } else {
    phoneError.value = ''
  }
}

function validateEmail() {
  const val = form.contact_email.trim()
  if (!val) {
    emailError.value = 'Email is required.'
  } else if (!val.includes('@')) {
    emailError.value = 'Email must contain "@" symbol.'
  } else {
    emailError.value = ''
  }
}

watch(info, (i) => {
  if (!localEdit.value) {
    Object.assign(form, {
      name: i.name || '',
      website: i.website || '',
      location: i.location || '',
      contact_email: i.contact_email || '',
      contact_number: i.contact_number || '',
      description: i.description || ''
    })
    validatePhone()
    validateEmail()
  }
}, { immediate: true })

watch(() => form.contact_number, validatePhone)
watch(() => form.contact_email, validateEmail)

const isSaveDisabled = computed(() =>
  props.isSaving ||
  !!phoneError.value ||
  !!emailError.value ||
  !form.contact_number.trim() ||
  !form.contact_email.trim()
)

function startEdit() { localEdit.value = true }
function cancelEdit() {
  localEdit.value = false
  const i = info.value
  Object.assign(form, {
    name: i.name || '',
    website: i.website || '',
    location: i.location || '',
    contact_email: i.contact_email || '',
    contact_number: i.contact_number || '',
    description: i.description || ''
  })
  validatePhone()
  validateEmail()
}

function submit() {
  validatePhone()
  validateEmail()
  if (phoneError.value || emailError.value) return
  emit('save', { ...form })
}
</script>
