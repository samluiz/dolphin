<template>
  <button
    @click="$emit('click')"
    type="button"
    class="bg-slate-300 text-black px-6 py-2 rounded-sm shadow-lg hover:bg-opacity-80 focus:outline-none focus:ring-1 focus:ring-primary-500 focus:ring-opacity-50 transition-all duration-200 ease-in-out"
  >
    <slot></slot>
  </button>
</template>
