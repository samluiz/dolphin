<script setup lang="ts">
import { Transition, toRefs } from "vue";

const props = defineProps({
  isOpen: Boolean,
});

const { isOpen } = toRefs(props);
</script>

<template>
  <Transition name="fade">
    <div
      v-show="isOpen"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50"
    >
      <div class="bg-white dark:bg-[#2d2d2d] p-8 rounded-sm w-96">
        <slot></slot>
      </div>
    </div>
  </Transition>
</template>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
