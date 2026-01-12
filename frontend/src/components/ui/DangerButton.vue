<template>
  <button
    @click="$emit('click')"
    type="button"
    class="group relative px-4 py-2 rounded-lg font-medium text-sm text-white overflow-hidden
           bg-gradient-to-r from-danger to-red-600
           hover:from-red-600 hover:to-red-700
           shadow-md hover:shadow-lg
           transform hover:scale-105 active:scale-95
           transition-all duration-200 ease-out
           focus:outline-none focus:ring-2 focus:ring-danger focus:ring-offset-2"
  >
    <!-- Shimmer effect on hover -->
    <span class="absolute inset-0 w-full h-full bg-gradient-to-r from-transparent via-white to-transparent opacity-0 group-hover:opacity-20 transform -skew-x-12 group-hover:translate-x-full transition-all duration-700 ease-out"></span>

    <!-- Content -->
    <span class="relative z-10 flex items-center justify-center gap-2">
      <slot></slot>
    </span>
  </button>
</template>
