<template>
  <div class="form-control">
    <label v-if="label" class="label" :for="id">
      <span class="label-text">{{ label }}</span>
    </label>

    <div class="relative">
      <input
        :id="id"
        :name="name"
        :type="isVisible ? 'text' : 'password'"
        :value="modelValue"
        :placeholder="placeholder"
        :autocomplete="autocomplete"
        :required="required"
        :disabled="disabled"
        :class="['input input-bordered w-full pr-12', inputClass, { 'input-error': !!error }]"
        :aria-invalid="!!error"
        :aria-describedby="error ? describedById : undefined"
        @input="$emit('update:modelValue', $event.target.value)"
        @blur="$emit('blur', $event)"
        @focus="$emit('focus', $event)"
      />

      <button
        type="button"
        class="btn btn-ghost btn-xs absolute right-2 top-1/2 -translate-y-1/2"
        @click="toggle"
        :aria-pressed="isVisible ? 'true' : 'false'"
        :aria-label="isVisible ? 'Hide password' : 'Show password'"
        :title="isVisible ? 'Hide password' : 'Show password'"
      >
        <!-- Eye -->
        <svg v-if="!isVisible" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="h-5 w-5">
          <path fill="currentColor" d="M12 5c-5 0-9 4-10 7 1 3 5 7 10 7s9-4 10-7c-1-3-5-7-10-7Zm0 12a5 5 0 1 1 0-10 5 5 0 0 1 0 10Zm0-2.5a2.5 2.5 0 1 0 0-5 2.5 2.5 0 0 0 0 5Z"/>
        </svg>
        <!-- Eye off -->
        <svg v-else xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" class="h-5 w-5">
          <path fill="currentColor" d="M2.81 2.81 1.39 4.22l3.2 3.2C2.9 8.55 1.67 10 1 12c1 3 5 7 11 7 2.11 0 3.98-.56 5.6-1.5l2.78 2.78 1.41-1.41L2.81 2.81ZM12 17c-3.86 0-6.88-2.55-8.19-5 .44-.87 1.15-1.8 2.09-2.6l2.2 2.2A5 5 0 0 0 12 17Zm0-10c3.86 0 6.88 2.55 8.19 5a9.8 9.8 0 0 1-3.02 3.49l-2.16-2.16A5 5 0 0 0 10.67 8.2l-1.49-1.5C10.12 6.24 11.03 6 12 6Zm0 3a3 3 0 0 1 3 3c0 .37-.07.73-.2 1.06l-3.86-3.86c.33-.13.69-.2 1.06-.2Z"/>
        </svg>
      </button>
    </div>

    <label v-if="error" class="label">
      <span class="label-text-alt text-error" :id="describedById">{{ error }}</span>
    </label>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';

const props = defineProps({
  modelValue: { type: String, default: '' },
  label: { type: String, default: '' },
  id: { type: String, default: '' },
  name: { type: String, default: '' },
  placeholder: { type: String, default: 'Enter your password' },
  error: { type: String, default: '' },
  autocomplete: { type: String, default: 'current-password' }, // or 'new-password'
  required: { type: Boolean, default: false },
  disabled: { type: Boolean, default: false },
  inputClass: { type: [String, Array, Object], default: '' },
  defaultVisible: { type: Boolean, default: false },
});

const emit = defineEmits(['update:modelValue', 'blur', 'focus']);

const isVisible = ref(!!props.defaultVisible);
const describedById = computed(() => (props.id ? `${props.id}-desc` : undefined));

function toggle() {
  isVisible.value = !isVisible.value;
}
</script>
