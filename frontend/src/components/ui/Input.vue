<script setup lang="ts">
import { defineProps } from "vue";

interface Props {
  name: string;
  title: string;
  type: string;
  required: boolean;
  modelValue: any;
}

defineProps<Props>();
</script>

<template>
  <div class="flex flex-col justify-start items-center gap-2">
    <label class="w-full text-sm text-black dark:text-white" :for="name">{{
      title
    }}</label>
    <input
      class="w-full p-1 rounded-sm shadow-lg text-white bg-secondary dark:text-black dark:bg-primary border-0 appearance-none outline-none focus:ring-1"
      :required="required"
      :type="type"
      :id="name"
      :value="modelValue"
      step="0.01"
      @input="
        $emit('update:modelValue', ($event.target as HTMLInputElement)?.value)
      "
    />
  </div>
</template>
