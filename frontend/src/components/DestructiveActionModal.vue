<script setup lang="ts">
import { toRefs } from "vue";
import Dialog from "./ui/Dialog.vue";
import CancelButton from "./ui/CancelButton.vue";
import DangerButton from "./ui/DangerButton.vue";

const props = defineProps({
  isOpen: Boolean,
  title: String,
});

const { isOpen, title } = toRefs(props);
</script>

<template>
  <div class="fixed z-50">
    <Dialog :isOpen="isOpen">
      <h1 class="text-2xl mb-2 text-black dark:text-white">
        {{ title ? title : $t("delete") }}
      </h1>
      <div class="flex flex-col gap-4">
        <h1 class="text-lg text-black dark:text-white">
          {{ $t("warning") }}
        </h1>
        <div class="flex gap-4">
          <CancelButton @click="$emit('on-cancel')">{{
            $t("cancel")
          }}</CancelButton>
          <DangerButton @click="$emit('on-confirm')">{{
            $t("delete")
          }}</DangerButton>
        </div>
      </div>
    </Dialog>
  </div>
</template>
