<template>
  <div class="w-full h-full mx-auto bg-white rounded-xl shadow-md p-4 md:p-8">
    <!-- Title + Edit -->
    <div class="flex items-center justify-between mb-4">
      <h2 class="text-xl font-semibold">Company Details</h2>
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
          <div class="font-medium">{{ info.contact_number || '-' }}</div>
        </div>

        <div class="mb-3">
          <div class="text-sm text-gray-500">Description</div>
          <div class="font-medium whitespace-pre-wrap">{{ info.description || '-' }}</div>
        </div>
      </div>
    </div>

    <!-- EDIT FORM — SHOWS BELOW; READ VIEW STAYS -->
    <form
      v-show="localEdit"
      class="grid md:grid-cols-2 gap-6 mt-6 border-t pt-6"
      @submit.prevent="submit"
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
        <label class="label"><span class="label-text">Contact Email</span></label>
        <input v-model.trim="form.contact_email" type="email" class="input input-bordered w-full" placeholder="hr@acme.co" />

        <label class="label mt-4"><span class="label-text">Contact Number</span></label>
        <input v-model.trim="form.contact_number" class="input input-bordered w-full" placeholder="+66-2-123-4567" />

        <label class="label mt-4"><span class="label-text">Description</span></label>
        <textarea v-model.trim="form.description" class="textarea textarea-bordered w-full" rows="5" placeholder="About the company…"></textarea>
      </div>

      <div class="md:col-span-2 flex items-center gap-2 mt-2">
        <button :disabled="isSaving" class="btn btn-primary">
          <span v-if="isSaving" class="loading loading-spinner loading-sm mr-2"></span>
          Save
        </button>
        <button type="button" class="btn btn-ghost" @click="cancelEdit">Cancel</button>

        <span v-if="alert?.message" :class="alert?.type === 'error' ? 'text-red-600' : 'text-green-700'">
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

/* Normalize backend keys (supports Name/name, etc.) */
function pick(obj, ...keys) {
  for (const k of keys) if (obj && obj[k] != null && obj[k] !== '') return obj[k]
  return ''
}
const info = computed(() => {
  const p = props.profile || {}
  return {
    name:           pick(p, 'name', 'Name'),
    website:        pick(p, 'website', 'Website'),
    location:       pick(p, 'location', 'Location'),
    contact_email:  pick(p, 'contact_email', 'ContactEmail'),
    contact_number: pick(p, 'contact_number', 'ContactNumber'),
    description:    pick(p, 'description', 'Description'),
  }
})

/* Local edit state + form; read view remains visible */
const localEdit = ref(false)
const form = reactive({ name:'', website:'', location:'', contact_email:'', contact_number:'', description:'' })

watch(info, (i) => {
  if (!localEdit.value) {
    form.name = i.name || ''
    form.website = i.website || ''
    form.location = i.location || ''
    form.contact_email = i.contact_email || ''
    form.contact_number = i.contact_number || ''
    form.description = i.description || ''
  }
}, { immediate: true })

function startEdit() { localEdit.value = true }
function cancelEdit() {
  localEdit.value = false
  const i = info.value
  form.name = i.name || ''
  form.website = i.website || ''
  form.location = i.location || ''
  form.contact_email = i.contact_email || ''
  form.contact_number = i.contact_number || ''
  form.description = i.description || ''
}
function submit() { emit('save', { ...form }) }
</script>
