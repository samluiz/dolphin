<script setup lang="ts">
import { defineProps } from "vue";
import { types } from "../../../wailsjs/go/models";

interface Props<T> {
  options: T[];
  name: string;
  title: string;
  required: boolean;
  modelValue: any;
}

defineProps<Props<types.Category>>();
</script>

<template>
  <div class="flex flex-col justify-start items-center gap-2">
    <label class="w-full text-sm" :for="name">Category</label>
    <select
      :value="modelValue"
      :name="name"
      :required="required"
      class="w-full p-1 rounded-sm shadow-lg text-white bg-secondary dark:text-black dark:bg-primary border-0 appearance-none outline-none focus:ring-1 placeholder:text-center"
      placeholder="Select a category"
      @change="
        $emit('update:modelValue', ($event.target as HTMLInputElement)?.value)
      "
    >
      <option value="" disabled selected>Select an option</option>
      <option v-for="option in options" :key="option.id" :value="option.id">
        {{ option.description }}
      </option>
    </select>
  </div>
</template>
