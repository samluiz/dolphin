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
    <label class="w-full text-sm text-black dark:text-white" :for="name">{{
      title
    }}</label>
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
      <option class="text-black dark:text-white" value="" disabled selected>
        {{ $t("select_an_option") }}
      </option>
      <option
        class="hover:bg-dark dark:hover:bg-light hover:bg-opacity-30 dark:hover:bg-opacity-30"
        v-for="option in options"
        :key="option.id"
        :value="option.id"
      >
        {{ $t(option.description.toLowerCase()) }}
      </option>
    </select>
  </div>
</template>
