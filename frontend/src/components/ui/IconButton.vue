<template>
  <button
    @click.stop="$emit('click')"
    type="button"
    class="duration-200 p-2 hover:bg-opacity-20 dark:hover:bg-opacity-20 hover:bg-black dark:hover:bg-white [clip-path:circle(50%)]"
  >
    <slot></slot>
  </button>
</template>
