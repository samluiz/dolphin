<script setup lang="ts">
import { computed, ref, toRefs } from "vue";
import { maskCurrency } from "@/shared/utils";

const props = defineProps<{
  earnings_subtotal: number;
  expenses_subtotal: number;
}>();

const { earnings_subtotal, expenses_subtotal } = toRefs(props);

const isDropdownOpen = ref(false);

const balance = computed(() => {
  return earnings_subtotal.value - expenses_subtotal.value;
});

const openDropdown = () => {
  isDropdownOpen.value = true;
};

const closeDropdown = () => {
  setTimeout(() => (isDropdownOpen.value = false), 1000);
};
</script>

<template>
  <div
    class="flex flex-col justify-self-end gap-2 text-sm min-w-44"
    @mouseover="openDropdown"
    @mouseleave="closeDropdown"
  >
    <div
      class="p-2 rounded-md"
      :class="{
        'bg-red-600': balance < 0,
        'bg-green-600': balance > 0,
        'bg-gray-600': balance == 0,
      }"
    >
      <span>Balance: {{ maskCurrency(balance) }}</span>
      <span
        class="ml-2 opacity-50 duration-200"
        :class="{
          'opacity-100': isDropdownOpen,
        }"
        >&#128712;</span
      >
    </div>
    <transition name="fade">
      <div class="grid grid-flow-row gap-2 min-w-44" v-if="isDropdownOpen">
        <span
          class="p-2 rounded-md bg-green-600 bg-opacity-50 hover:bg-opacity-100 duration-200"
          >Total Earnings: {{ maskCurrency(earnings_subtotal) }}</span
        >

        <span
          class="p-2 rounded-md bg-red-600 bg-opacity-50 hover:bg-opacity-100 duration-200"
          >Total Expenses: {{ maskCurrency(expenses_subtotal) }}</span
        >
      </div>
    </transition>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
