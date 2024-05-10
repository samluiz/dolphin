<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import Table from "./components/Table.vue";
import DolphinIcon from "./components/ui/icons/DolphinIcon.vue";
import ThemeSwitcher from "./components/ThemeSwitcher.vue";
import { types } from "../wailsjs/go/models";
import {
  Create as createProfile,
  Update as updateProfile,
  FindAll as findAllProfiles,
  Delete as deleteProfile,
} from "../wailsjs/go/profile/service";

const profiles = ref<types.Profile[]>([]);
const hasProfileCreated = ref(false);

const fetchData = async () => {
  const response = await findAllProfiles();
  profiles.value = response;
};

const isDarkMode = ref(false);

const toggleTheme = () => {
  console.log("toggleTheme");
  isDarkMode.value = !isDarkMode.value;
};

onMounted(() => {
  const isDark = localStorage.getItem("isDarkMode");
  if (isDark) {
    isDarkMode.value = JSON.parse(isDark);
  }

  watch(isDarkMode, (newVal) => {
    localStorage.setItem("isDarkMode", JSON.stringify(newVal));
  });

  fetchData();
  if (profiles.value.length !== 0) {
    hasProfileCreated.value = true;
  }
});
</script>

<template>
  <ThemeSwitcher @toggle-theme="toggleTheme" :is-dark-mode="isDarkMode">
    <div
      class="h-screen w-screen bg-[#eaeaea] dark:bg-[#1c1c1c] p-8"
      :class="{
        dark: isDarkMode,
      }"
    >
      <div class="w-full grid place-items-start p-8">
        <div class="flex items-center gap-4">
          <DolphinIcon />
          <h1 class="text-black dark:text-white text-5xl">Dolphin</h1>
        </div>
      </div>
      <div class="w-full p-2">
        <Table />
      </div>
    </div>
  </ThemeSwitcher>
</template>
