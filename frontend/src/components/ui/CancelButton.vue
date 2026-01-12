<template>
  <button
    @click="$emit('click')"
    type="button"
    class="px-4 py-2 rounded-lg font-medium text-sm
           text-text-secondary-light dark:text-text-secondary-dark
           bg-gray-100 dark:bg-surface-dark
           border border-border-light dark:border-border-dark
           hover:bg-gray-200 dark:hover:bg-opacity-80
           shadow-sm hover:shadow-md
           transform hover:scale-105 active:scale-95
           transition-all duration-200 ease-out
           focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2"
  >
    <slot></slot>
  </button>
</template>
