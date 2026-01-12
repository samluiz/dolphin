<template>
  <button
    @click.stop="$emit('click')"
    type="button"
    class="group relative p-2 rounded-full
           hover:bg-primary-50 dark:hover:bg-surface-dark
           active:bg-primary-100 dark:active:bg-opacity-90
           transition-all duration-200 ease-out
           hover:shadow-md-blue
           transform hover:scale-110 active:scale-95
           focus:outline-none focus:ring-2 focus:ring-primary-400"
  >
    <span class="relative z-10 text-text-primary-light dark:text-text-primary-dark group-hover:text-primary-500 dark:group-hover:text-primary-400 transition-colors">
      <slot></slot>
    </span>
  </button>
</template>
