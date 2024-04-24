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
      <div class="bg-gray-700 p-8 rounded-md w-96">
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
