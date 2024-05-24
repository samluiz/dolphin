<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import Table from "./components/Table.vue";
import DolphinIcon from "./components/ui/icons/DolphinIcon.vue";
import ThemeSwitcher from "./components/ThemeSwitcher.vue";
import { GetAppVersion } from "../wailsjs/go/main/App";
import CreateProfileModal from "./components/CreateProfileModal.vue";
import IconButton from "./components/ui/IconButton.vue";
import AddIcon from "./components/ui/icons/AddIcon.vue";
import { useProfileStore } from "./stores/ProfileStore";
import ProfileSelector from "./components/ProfileSelector.vue";

const profileStore = useProfileStore();

const isCreateProfileModalOpen = ref(false);
const appVersion = ref("");

const fetchData = async () => {
  await fetchProfiles();
};

const fetchProfiles = async () => {
  await profileStore.fetchProfiles();
};

const fetchAppVersion = async () => {
  const response = await GetAppVersion();
  appVersion.value = response;
};

const isDarkMode = ref(false);

const toggleTheme = () => {
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
});

const openCreateProfileModal = () => {
  isCreateProfileModalOpen.value = true;
};

const closeCreateProfileModal = () => {
  isCreateProfileModalOpen.value = false;
  fetchData();
};

fetchAppVersion();
</script>

<template>
  <ThemeSwitcher @toggle-theme="toggleTheme" :is-dark-mode="isDarkMode">
    <div
      class="h-screen w-screen bg-light dark:bg-dark p-8"
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
      <div
        class="w-full grid place-items-center p-8 text-center text-black dark:text-white"
        v-if="!profileStore.hasProfileCreated()"
      >
        <span class="text-lg font-semibold"
          >No profiles found. Create one right now!</span
        >
        <IconButton @click="openCreateProfileModal">
          <AddIcon />
        </IconButton>
      </div>

      <div
        class="grid place-items-center p-8 text-center text-black dark:text-white"
        v-if="profileStore.hasProfileCreated() && !profileStore.profile"
      >
        <ProfileSelector />
      </div>

      <div v-else>
        <div class="w-full p-2">
          <Table />
        </div>
      </div>
      <CreateProfileModal
        v-if="isCreateProfileModalOpen"
        :has-profile-created="!profileStore.hasProfileCreated()"
        :is-open="isCreateProfileModalOpen"
        @on-cancel="closeCreateProfileModal"
      />

      <div class="absolute bottom-0 left-0 p-4">
        <span class="text-xs opacity-70 text-black dark:text-white">{{
          appVersion
        }}</span>
      </div>
    </div>
  </ThemeSwitcher>
</template>
