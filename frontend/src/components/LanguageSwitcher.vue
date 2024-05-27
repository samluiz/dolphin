<script setup lang="ts"></script>

<template>
  <div>
    <select
      class="p-2 text-center text-sm bg-transparent text-black dark:text-white border-0 appearance-none outline-none focus:ring-1 placeholder:text-center"
      v-model="$i18n.locale"
    >
      <option
        class="text-black"
        v-for="locale in $i18n.availableLocales"
        :key="`locale-${locale}`"
        :value="locale"
      >
        {{ locale !== "pt-BR" ? "English" : "Português" }}
      </option>
    </select>
  </div>
</template>
