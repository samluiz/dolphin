<script setup lang="ts">
import { computed, ref, toRefs } from "vue";
import { maskCurrency } from "@/shared/utils";
import SelectArrowIcon from "./ui/icons/SelectArrowIcon.vue";
import { useI18n } from "vue-i18n";

const i18n = useI18n();

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
    class="relative min-w-52"
    @click="toggleDropdown"
    v-click-outside="() => (isDropdownOpen = false)"
  >
    <!-- Main balance card -->
    <div
      class="card p-3 cursor-pointer
             transform hover:scale-105 active:scale-95
             transition-all duration-200 ease-out
             hover:shadow-xl-blue"
      :class="{
        'ring-2 ring-danger': balance < 0,
        'ring-2 ring-success': balance > 0,
      }"
    >
      <div class="flex items-center justify-between">
        <div class="flex-1">
          <p class="text-xs font-medium text-text-secondary-light dark:text-text-secondary-dark mb-0.5">
            {{ $t("balance") }}
          </p>
          <p class="text-lg font-bold" :class="{
            'text-danger': balance < 0,
            'text-success': balance > 0,
            'text-gray-500 dark:text-gray-400': balance === 0,
          }">
            {{ maskCurrency(balance) }}
          </p>
        </div>
        <div
          class="p-2 rounded-full transition-transform duration-200"
          :class="{
            'bg-danger/10': balance < 0,
            'bg-success/10': balance > 0,
            'bg-gray-200 dark:bg-gray-700': balance === 0,
            'rotate-180': isDropdownOpen,
          }"
        >
          <SelectArrowIcon
            class="w-4 h-4"
            :class="{
              'text-danger': balance < 0,
              'text-success': balance > 0,
              'text-gray-500': balance === 0,
            }"
          />
        </div>
      </div>
    </div>

    <!-- Dropdown details -->
    <transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="opacity-0 transform -translate-y-2"
      enter-to-class="opacity-100 transform translate-y-0"
      leave-active-class="transition ease-in duration-150"
      leave-from-class="opacity-100 transform translate-y-0"
      leave-to-class="opacity-0 transform -translate-y-2"
    >
      <div
        v-show="isDropdownOpen"
        class="absolute z-50 right-0 mt-2 w-full space-y-2"
      >
        <!-- Earnings card -->
        <div class="card p-3 bg-gradient-to-br from-success to-teal-600">
          <div class="flex items-center justify-between text-white">
            <div>
              <p class="text-xs opacity-90 mb-0.5">
                {{ $t("total") }} {{ i18n.t("earning", 2) }}
              </p>
              <p class="text-base font-bold">
                {{ maskCurrency(earnings_subtotal) }}
              </p>
            </div>
            <div class="p-1.5 bg-white/20 rounded-lg">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11l5-5m0 0l5 5m-5-5v12"/>
              </svg>
            </div>
          </div>
        </div>

        <!-- Expenses card -->
        <div class="card p-3 bg-gradient-to-br from-danger to-red-600">
          <div class="flex items-center justify-between text-white">
            <div>
              <p class="text-xs opacity-90 mb-0.5">
                {{ $t("total") }} {{ i18n.t("expense", 2) }}
              </p>
              <p class="text-base font-bold">
                {{ maskCurrency(expenses_subtotal) }}
              </p>
            </div>
            <div class="p-1.5 bg-white/20 rounded-lg">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13l-5 5m0 0l-5-5m5 5V6"/>
              </svg>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>
