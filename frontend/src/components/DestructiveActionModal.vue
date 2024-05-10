<script setup lang="ts">
import { toRefs } from "vue";
import Dialog from "./ui/Dialog.vue";
import CancelButton from "./ui/CancelButton.vue";
import DangerButton from "./ui/DangerButton.vue";

const props = defineProps({
  isOpen: Boolean,
});

const { isOpen } = toRefs(props);
</script>

<template>
  <div class="fixed z-50">
    <Dialog :isOpen="isOpen">
      <div class="flex flex-col gap-4">
        <h1 class="text-lg text-black dark:text-white">
          Are you sure you want to delete this item? This action cannot be
          undone.
        </h1>
        <div class="flex gap-4">
          <CancelButton @click="$emit('on-cancel')">Cancel</CancelButton>
          <DangerButton @click="$emit('on-confirm')">Delete</DangerButton>
        </div>
      </div>
    </Dialog>
  </div>
</template>
