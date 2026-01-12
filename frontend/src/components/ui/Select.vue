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
  <div class="relative w-full">
    <label
      :for="name"
      class="block mb-2 text-sm font-medium text-text-primary-light dark:text-text-primary-dark"
    >
      {{ title }}
    </label>
    <select
      :value="modelValue"
      :name="name"
      :required="required"
      class="w-full px-4 py-3 rounded-lg
             bg-white dark:bg-surface-dark
             border-2 border-border-light dark:border-border-dark
             text-text-primary-light dark:text-text-primary-dark
             focus:border-primary-500 dark:focus:border-primary-400
             focus:ring-4 focus:ring-primary-100 dark:focus:ring-primary-900/30
             transition-all duration-200 ease-out
             shadow-sm focus:shadow-md-blue
             cursor-pointer
             appearance-none
             bg-no-repeat bg-right
             pr-10"
      style="background-image: url('data:image/svg+xml;charset=UTF-8,%3csvg xmlns=%27http://www.w3.org/2000/svg%27 viewBox=%270 0 24 24%27 fill=%27none%27 stroke=%27currentColor%27 stroke-width=%272%27 stroke-linecap=%27round%27 stroke-linejoin=%27round%27%3e%3cpolyline points=%276 9 12 15 18 9%27%3e%3c/polyline%3e%3c/svg%3e'); background-size: 1.5rem; background-position: right 0.75rem center;"
      @change="$emit('update:modelValue', ($event.target as HTMLInputElement)?.value)"
    >
      <option value="" disabled selected>
        {{ $t("select_an_option") }}
      </option>
      <option
        v-for="option in options"
        :key="option.id"
        :value="option.id"
      >
        {{ $t(option.description.toLowerCase()) }}
      </option>
    </select>
  </div>
</template>
