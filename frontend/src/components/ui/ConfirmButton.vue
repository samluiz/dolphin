<script setup lang="ts">
interface Props {
  disabled?: boolean;
}

defineProps<Props>();
defineEmits(['click']);
</script>

<template>
  <button
    @click="$emit('click')"
    type="button"
    :disabled="disabled"
    class="group relative px-4 py-2 rounded-lg font-medium text-sm text-white overflow-hidden
           bg-gradient-to-r from-primary-500 to-secondary-500
           transition-all duration-200 ease-out
           focus:outline-none focus:ring-2 focus:ring-primary-400 focus:ring-offset-2"
    :class="disabled
      ? 'opacity-50 cursor-not-allowed'
      : 'hover:from-primary-600 hover:to-secondary-600 shadow-md-blue hover:shadow-lg-blue transform hover:scale-105 active:scale-95 cursor-pointer'"
  >
    <!-- Shimmer effect on hover -->
    <span
      v-if="!disabled"
      class="absolute inset-0 w-full h-full bg-gradient-to-r from-transparent via-white to-transparent opacity-0 group-hover:opacity-20 transform -skew-x-12 group-hover:translate-x-full transition-all duration-700 ease-out"
    ></span>

    <!-- Content -->
    <span class="relative z-10 flex items-center justify-center gap-2">
      <slot></slot>
    </span>
  </button>
</template>
