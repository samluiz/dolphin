<script setup lang="ts">
import BrazilianFlag from "./ui/icons/BrazilianFlag.vue";
import AmericanFlag from "./ui/icons/AmericanFlag.vue";

const locales = [
  { code: "pt-BR", label: "Português", flag: BrazilianFlag },
  { code: "en-US", label: "English", flag: AmericanFlag },
];
</script>

<template>
  <div>
    <ul
      class="px-2 text-center grid place-items-center grid-flow-col gap-2 bg-transparent text-black dark:text-light-text"
    >
      <li
        class="grid place-items-center grid-flow-row cursor-pointer"
        v-for="locale in locales"
        :key="`locale-${locale.code}`"
        :value="locale.code"
        @click="$i18n.locale = locale.code"
      >
        <component
          class="text-2xl duration-200 hover:scale-110"
          :is="locale.flag"
        />
      </li>
    </ul>
  </div>
</template>
