<script setup lang="ts">
import { Transition, toRefs, watch } from "vue";

const props = defineProps({
  isOpen: Boolean,
});

const { isOpen } = toRefs(props);

watch(isOpen, (open: boolean) => {
  document.body.classList.toggle("modal-open", open);
});
</script>

<template>
  <Teleport to="body">
    <Transition
      enter-active-class="transition ease-out duration-300"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition ease-in duration-200"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div
        v-if="isOpen"
        class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/60 pointer-events-auto"
      >
        <Transition
          enter-active-class="transition ease-out duration-300 delay-100"
          enter-from-class="opacity-0 scale-95 translate-y-4"
          enter-to-class="opacity-100 scale-100 translate-y-0"
          leave-active-class="transition ease-in duration-200"
          leave-from-class="opacity-100 scale-100 translate-y-0"
          leave-to-class="opacity-0 scale-95 translate-y-4"
        >
          <div
            v-if="isOpen"
            class="card w-full max-w-md p-8 bg-white dark:bg-surface-dark shadow-2xl max-h-[90vh] overflow-y-auto"
          >
            <slot></slot>
          </div>
        </Transition>
      </div>
    </Transition>
  </Teleport>
</template>
