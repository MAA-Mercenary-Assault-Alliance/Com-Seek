<template>
  <div class="w-full h-full mx-auto bg-white rounded-xl shadow-md overflow-hidden">
    <!-- Top accent bar -->
    <div class="h-2 w-full bg-gradient-to-r from-[#0A3B1F] via-[#44B15B] to-[#56A45C]"></div>

    <div class="p-4 md:p-8">
      <!-- Title + Edit -->
      <div class="flex items-center justify-between mb-6">
        <h2 class="text-xl font-bold text-gray-800">
          <span class="text-2xl text-[#0A3B1F] font-black">Company</span>
          <span class="ml-1">Details</span>
        </h2>

        <button
          v-if="canEdit && !localEdit"
          class="btn btn-sm btn-outline border-[#44B15B] text-[#0A3B1F] hover:bg-[#EAF6EC] hover:border-[#56A45C]"
          @click="startEdit"
        >
          Edit
        </button>
      </div>

      <!-- READ VIEW — ALWAYS VISIBLE -->
      <div class="grid md:grid-cols-2 gap-6">
        <div>
          <div class="mb-4">
            <div class="text-xs uppercase tracking-wide text-[#0A3B1F]/70">Name</div>
            <div class="font-medium text-gray-900">{{ info.name || '-' }}</div>
          </div>

          <div class="mb-4">
            <div class="text-xs uppercase tracking-wide text-[#0A3B1F]/70">Website</div>
            <div class="font-medium">
              <a
                v-if="info.website"
                :href="info.website"
                target="_blank"
                rel="noopener"
                class="break-all text-[#0A3B1F] hover:text-[#44B15B] underline underline-offset-2"
              >{{ info.website }}</a>
              <span v-else class="text-gray-500">-</span>
            </div>
          </div>

          <div class="mb-4">
            <div class="text-xs uppercase tracking-wide text-[#0A3B1F]/70">Location</div>
            <div class="font-medium text-gray-900">{{ info.location || '-' }}</div>
          </div>
        </div>

        <div>
          <div class="mb-4">
            <div class="text-xs uppercase tracking-wide text-[#0A3B1F]/70">Contact Email</div>
            <div class="font-medium text-gray-900">{{ info.contact_email || '-' }}</div>
          </div>

          <div class="mb-4">
            <div class="text-xs uppercase tracking-wide text-[#0A3B1F]/70">Contact Number</div>
            <div class="font-medium text-gray-900 whitespace-pre-wrap">{{ info.contact_number || '-' }}</div>
          </div>

          <div class="mb-4">
            <div class="text-xs uppercase tracking-wide text-[#0A3B1F]/70">Description</div>
            <div class="font-medium text-gray-900 whitespace-pre-wrap">{{ info.description || '-' }}</div>
          </div>
        </div>
      </div>

      <!-- Divider -->
      <div class="mt-4 border-t border-dashed border-[#0A3B1F]/15"></div>

      <!-- EDIT FORM — SHOWN BELOW -->
      <form
        v-show="localEdit"
        class="grid md:grid-cols-2 gap-6 mt-6"
        @submit.prevent="submit"
        novalidate
      >
        <div>
          <label class="label"><span class="label-text text-[#0A3B1F]">Name</span></label>
          <input
            v-model.trim="form.name"
            class="input input-bordered w-full focus:outline-none focus:border-[#44B15B] focus:ring-2 focus:ring-[#44B15B]/30"
            placeholder="Acme Robotics"
          />

          <label class="label mt-4"><span class="label-text text-[#0A3B1F]">Website</span></label>
          <input
            v-model.trim="form.website"
            class="input input-bordered w-full focus:outline-none focus:border-[#44B15B] focus:ring-2 focus:ring-[#44B15B]/30"
            placeholder="https://…"
          />

          <label class="label mt-4"><span class="label-text text-[#0A3B1F]">Location</span></label>
          <input
            v-model.trim="form.location"
            class="input input-bordered w-full focus:outline-none focus:border-[#44B15B] focus:ring-2 focus:ring-[#44B15B]/30"
            placeholder="Bangkok"
          />
        </div>

        <div>
          <!-- EMAIL -->
          <label class="label"><span class="label-text text-[#0A3B1F]">Contact Email</span></label>
          <input
            v-model.trim="form.contact_email"
            type="email"
            class="input input-bordered w-full focus:outline-none focus:border-[#44B15B] focus:ring-2 focus:ring-[#44B15B]/30"
            :class="emailError ? 'input-error' : ''"
            placeholder="hr@acme.co"
            @input="validateEmail"
            @blur="validateEmail"
          />
          <label v-if="emailError" class="label">
            <span class="label-text-alt text-error">{{ emailError }}</span>
          </label>

          <!-- PHONE -->
          <label class="label mt-4"><span class="label-text text-[#0A3B1F]">Contact Number</span></label>
          <input
            v-model.trim="form.contact_number"
            type="tel"
            inputmode="numeric"
            class="input input-bordered w-full focus:outline-none focus:border-[#44B15B] focus:ring-2 focus:ring-[#44B15B]/30"
            :class="phoneError ? 'input-error' : ''"
            placeholder="0804326754"
            @input="validatePhone"
            @blur="validatePhone"
          />
          <label v-if="phoneError" class="label">
            <span class="label-text-alt text-error">{{ phoneError }}</span>
          </label>

          <label class="label mt-4"><span class="label-text text-[#0A3B1F]">Description</span></label>
          <textarea
            v-model.trim="form.description"
            class="textarea textarea-bordered w-full focus:outline-none focus:border-[#44B15B] focus:ring-2 focus:ring-[#44B15B]/30"
            rows="5"
            placeholder="About the company…"
          ></textarea>
        </div>

        <div class="md:col-span-2 flex items-center gap-2 mt-2">
          <button
            :disabled="isSaveDisabled"
            class="btn btn-primary bg-[#44B15B] border-[#44B15B] hover:bg-[#56A45C] hover:border-[#56A45C]"
          >
            <span v-if="isSaving" class="loading loading-spinner loading-sm mr-2"></span>
            Save
          </button>
          <button
            type="button"
            class="btn btn-ghost text-[#0A3B1F] hover:bg-[#EAF6EC]"
            @click="cancelEdit"
          >
            Cancel
          </button>

          <span
            v-if="alert?.message"
            :class="alert?.type === 'error' ? 'text-red-600' : 'text-[#0A3B1F]'"
          >
            {{ alert.message }}
          </span>
        </div>
      </form>
    </div>
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
