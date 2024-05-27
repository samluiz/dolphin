<script setup lang="ts">
import { computed, ref, toRefs } from "vue";
import { maskCurrency } from "@/shared/utils";
import SelectArrowIcon from "./ui/icons/SelectArrowIcon.vue";

const props = defineProps<{
  earnings_subtotal: number;
  expenses_subtotal: number;
}>();

const { earnings_subtotal, expenses_subtotal } = toRefs(props);

const isDropdownOpen = ref(false);

const balance = computed(() => {
  return earnings_subtotal.value - expenses_subtotal.value;
});

const toggleDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value;
};
</script>

<template>
  <div
    class="flex flex-col justify-self-end gap-2 text-sm min-w-44 relative"
    @click="toggleDropdown"
    v-click-outside="() => (isDropdownOpen = false)"
  >
    <div
      class="p-2 rounded-md grid place-items-center grid-flow-col cursor-pointer duration-200"
      :class="{
        'bg-red-600 hover:bg-red-700': balance < 0,
        'bg-green-600 hover:bg-green-700': balance > 0,
        'bg-gray-600 hover:bg-gray-700': balance == 0,
      }"
    >
      <span>Balance: {{ maskCurrency(balance) }}</span>
      <SelectArrowIcon />
    </div>
    <transition name="fade">
      <div
        class="grid grid-flow-row text-center gap-2 min-w-44 mt-11 absolute z-50"
        v-show="isDropdownOpen"
      >
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
