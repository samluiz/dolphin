<script setup lang="ts">
import { defineProps } from "vue";

interface Props {
  name: string;
  title: string;
  type: string;
  required: boolean;
  checked: boolean;
  disabled: boolean;
  modelValue: any;
}

defineProps<Props>();
</script>

<template>
  <div class="flex flex-row justify-start items-center gap-2">
    <input
      class="rounded-sm shadow-lg text-light-text bg-secondary dark:text-black dark:bg-primary"
      :disabled="disabled"
      :required="required"
      :type="type"
      :id="name"
      :checked="checked"
      :value="modelValue"
      @input="
        $emit('update:modelValue', ($event.target as HTMLInputElement)?.value)
      "
    />
    <label class="w-full text-sm text-black dark:text-light-text" :for="name">{{
      title
    }}</label>
  </div>
</template>
