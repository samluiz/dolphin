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
      <span v-if="isDarkMode"><LightModeIcon /></span>
      <span v-else><DarkModeIcon /></span>
    </button>
  </div>
  <slot></slot>
</template>
