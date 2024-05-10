<template>
  <button
    @click.stop="$emit('click')"
    type="button"
    class="duration-200 p-2 text-black dark:text-white hover:bg-opacity-30 hover:bg-white [clip-path:circle(50%)]"
  >
    <slot></slot>
  </button>
</template>
