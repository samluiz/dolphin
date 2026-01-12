<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import Table from "./components/Table.vue";
import DolphinIcon from "./components/ui/icons/DolphinIcon.vue";
import ThemeSwitcher from "./components/ThemeSwitcher.vue";
import { GetAppVersion } from "../wailsjs/go/main/App";
import CreateProfileModal from "./components/CreateProfileModal.vue";
import SettingsModal from "./components/SettingsModal.vue";
import IconButton from "./components/ui/IconButton.vue";
import AddIcon from "./components/ui/icons/AddIcon.vue";
import SettingsIcon from "./components/ui/icons/SettingsIcon.vue";
import { useProfileStore } from "./stores/ProfileStore";
import ProfileSelector from "./components/ProfileSelector.vue";
import { isUserPortugueseSpeaker } from "./shared/utils";
import { useI18n } from "vue-i18n";

const { t, locale } = useI18n();

const profileStore = useProfileStore();

const isCreateProfileModalOpen = ref(false);
const isSettingsModalOpen = ref(false);
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

const setTheme = (isDark: boolean) => {
  isDarkMode.value = isDark;
};

const getThemeFromLocalStorage = () => {
  const isDark = localStorage.getItem("isDarkMode");
  if (isDark) {
    isDarkMode.value = JSON.parse(isDark);
  }
};

const setThemeOnLocalStorage = (darkMode: boolean) => {
  localStorage.setItem("isDarkMode", JSON.stringify(darkMode));
};

const getLocaleFromLocalStorage = () => {
  const localeLs = localStorage.getItem("locale");
  if (localeLs) {
    locale.value = localeLs;
  } else {
    locale.value = isUserPortugueseSpeaker() ? "pt-BR" : "en-US";
  }
};

const setLocaleOnLocalStorage = (locale: string) => {
  localStorage.setItem("locale", locale);
};

const applyZoomLevel = () => {
  const zoom = localStorage.getItem("zoomLevel");
  if (zoom) {
    document.documentElement.style.fontSize = `${zoom}%`;
  }
};

onMounted(() => {
  getThemeFromLocalStorage();

  watch(isDarkMode, (newVal) => {
    setThemeOnLocalStorage(newVal);
  });

  getLocaleFromLocalStorage();

  watch(locale, (newVal) => {
    setLocaleOnLocalStorage(newVal);
  });

  // Apply zoom level on mount
  applyZoomLevel();

  fetchData();
});

const openCreateProfileModal = () => {
  isCreateProfileModalOpen.value = true;
};

const closeCreateProfileModal = () => {
  isCreateProfileModalOpen.value = false;
  fetchData();
};

const openSettingsModal = () => {
  isSettingsModalOpen.value = true;
};

const closeSettingsModal = () => {
  isSettingsModalOpen.value = false;
};

const handleSettingsChange = (settings: any) => {
  console.log("Settings changed:", settings);
};

fetchAppVersion();
</script>

<template>
  <ThemeSwitcher @toggle-theme="toggleTheme" :is-dark-mode="isDarkMode">
    <div class="min-h-screen w-screen relative" :class="{ dark: isDarkMode }">
      <div
        class="absolute inset-0 bg-gradient-to-br from-primary-50 via-blue-50 to-primary-100 dark:from-bg-dark dark:via-surface-dark dark:to-bg-dark"
      ></div>

      <div
        class="absolute inset-0 overflow-hidden opacity-30 dark:opacity-20 pointer-events-none"
      >
        <div
          class="absolute -top-24 -right-24 w-96 h-96 bg-primary-300 dark:bg-primary-800 rounded-full mix-blend-multiply dark:mix-blend-soft-light filter blur-3xl opacity-50 animate-blob"
        ></div>
        <div
          class="absolute -bottom-24 -left-24 w-96 h-96 bg-secondary-300 dark:bg-secondary-800 rounded-full mix-blend-multiply dark:mix-blend-soft-light filter blur-3xl opacity-50 animate-blob animation-delay-2000"
        ></div>
        <div
          class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-96 h-96 bg-accent-300 dark:bg-accent-800 rounded-full mix-blend-multiply dark:mix-blend-soft-light filter blur-3xl opacity-50 animate-blob animation-delay-4000"
        ></div>
      </div>

      <div class="relative z-10 p-4 md:p-6">
        <div class="max-w-7xl mx-auto mb-4">
          <div class="card p-4 flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div class="p-2 bg-gradient-ocean rounded-xl shadow-glow-blue">
                <DolphinIcon class="w-8 h-8 text-white" />
              </div>
              <div>
                <h1 class="text-3xl font-bold text-gradient">Dolphin</h1>
                <p
                  class="text-xs text-text-secondary-light dark:text-text-secondary-dark mt-0.5"
                >
                  {{ $t("finance_manager") }}
                </p>
              </div>
            </div>

            <div class="flex items-center gap-2">
              <IconButton @click="openSettingsModal" title="Settings">
                <SettingsIcon class="w-5 h-5" />
              </IconButton>
            </div>
          </div>
        </div>

        <div class="max-w-7xl mx-auto">
          <div
            class="card p-8 text-center"
            v-if="!profileStore.hasProfileCreated()"
          >
            <div class="max-w-md mx-auto">
              <div
                class="w-16 h-16 bg-gradient-ocean rounded-full flex items-center justify-center mx-auto mb-4 shadow-glow-blue"
              >
                <AddIcon class="w-8 h-8 text-white" />
              </div>
              <h2
                class="text-xl font-semibold text-text-primary-light dark:text-text-primary-dark mb-3"
              >
                {{ $t("no_profile") }}
              </h2>
              <p
                class="text-sm text-text-secondary-light dark:text-text-secondary-dark mb-4"
              >
                {{ $t("prompt_profile") }}
              </p>
              <IconButton @click="openCreateProfileModal">
                <AddIcon />
              </IconButton>
            </div>
          </div>

          <div
            class="card p-8"
            v-else-if="
              profileStore.hasProfileCreated() && !profileStore.profile
            "
          >
            <ProfileSelector
              :profile="profileStore.profile!"
              @on-select="profileStore.setActiveProfile"
              @on-profile-create="fetchData()"
              @on-cancel="fetchData()"
            />
          </div>

          <div v-else class="card p-4">
            <Table />
          </div>
        </div>
      </div>

      <CreateProfileModal
        v-if="isCreateProfileModalOpen"
        :has-profile-created="!profileStore.hasProfileCreated()"
        :is-open="isCreateProfileModalOpen"
        @on-cancel="closeCreateProfileModal"
      />

      <SettingsModal
        v-if="isSettingsModalOpen"
        :is-open="isSettingsModalOpen"
        :is-dark-mode="isDarkMode"
        @on-cancel="closeSettingsModal"
        @on-set-theme="setTheme"
        @on-settings-change="handleSettingsChange"
      />

      <div
        class="absolute bottom-0 left-0 right-0 p-4 flex justify-between items-center"
      >
        <span
          class="text-xs opacity-50 text-text-secondary-light dark:text-text-secondary-dark"
        >
          {{ appVersion }}
        </span>
      </div>
    </div>
  </ThemeSwitcher>
</template>
