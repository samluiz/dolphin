<script setup lang="ts">
import { toRefs } from "vue";
import LightModeIcon from "./ui/icons/LightModeIcon.vue";
import DarkModeIcon from "./ui/icons/DarkModeIcon.vue";

const props = defineProps({
  isDarkMode: Boolean,
});

const emits = defineEmits(["toggle-theme"]);

const { isDarkMode } = toRefs(props);
</script>

<template>
  <div
    class="fixed duration-200 top-8 right-8"
    :class="{
      dark: isDarkMode,
    }"
  >
    <button
      @click="$emit('toggle-theme')"
      class="p-2 text-black dark:text-white rounded-lg"
    >
      <transition name="icon-fade" mode="out-in">
        <span v-if="isDarkMode" key="dark"><LightModeIcon /></span>
        <span v-else key="light"><DarkModeIcon /></span>
      </transition>
    </button>
  </div>
  <slot></slot>
</template>

<style scoped>
.icon-fade-enter-active,
.icon-fade-leave-active {
  transition: opacity 0.2s ease-out;
}
.icon-fade-enter,
.icon-fade-leave-to {
  opacity: 0;
}
</style>
