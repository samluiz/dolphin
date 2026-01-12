<script setup lang="ts">
import { toRefs, watch } from "vue";

const props = defineProps({
  isDarkMode: Boolean,
});

const emits = defineEmits(["toggle-theme"]);

const { isDarkMode } = toRefs(props);

watch(
  () => props.isDarkMode,
  (newVal) => {
    document.documentElement.classList.toggle("dark", newVal);
  },
  { immediate: true },
);
</script>

<template>
  <slot></slot>
</template>
