<script setup lang="ts">
import { defineProps, ref } from "vue";

interface Props {
  name: string;
  title: string;
  type: string;
  required: boolean;
  modelValue: any;
}

defineProps<Props>();

const isFocused = ref(false);
</script>

<template>
  <div class="relative w-full">
    <input
      class="peer w-full px-4 py-3 rounded-lg
             bg-white dark:bg-surface-dark
             border-2 border-border-light dark:border-border-dark
             text-text-primary-light dark:text-text-primary-dark
             placeholder-transparent
             focus:border-primary-500 dark:focus:border-primary-400
             focus:ring-4 focus:ring-primary-100 dark:focus:ring-primary-900/30
             transition-all duration-200 ease-out
             shadow-sm focus:shadow-md-blue"
      :required="required"
      :type="type"
      :id="name"
      :value="modelValue"
      :placeholder="title"
      step="0.01"
      @focus="isFocused = true"
      @blur="isFocused = false"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement)?.value)"
    />
    <label
      :for="name"
      class="absolute left-4 -top-2.5 px-1
             text-sm font-medium
             text-text-secondary-light dark:text-text-secondary-dark
             bg-white dark:bg-surface-dark
             transition-all duration-200 ease-out
             peer-placeholder-shown:text-base peer-placeholder-shown:top-3 peer-placeholder-shown:text-text-secondary-light/60
             peer-focus:-top-2.5 peer-focus:text-sm peer-focus:text-primary-500 dark:peer-focus:text-primary-400"
    >
      {{ title }}
    </label>
  </div>
</template>
