<script setup lang="ts">
import { ref, watch, computed } from "vue";
import Dialog from "./ui/Dialog.vue";
import ConfirmButton from "./ui/ConfirmButton.vue";
import CancelButton from "./ui/CancelButton.vue";
import SettingsIcon from "./ui/icons/SettingsIcon.vue";
import { useI18n } from "vue-i18n";
import { useSettingsStore } from "../stores/SettingsStore";

const { t, locale } = useI18n();
const settings = useSettingsStore();

console.log("SettingsModal: store initial state:", {
  zoom: settings.zoom,
  locale: settings.locale,
  currency: settings.currency,
  dateFormat: settings.dateFormat,
  itemsPerPage: settings.itemsPerPage,
  darkMode: settings.darkMode,
});

interface Props {
  isOpen: boolean;
  isDarkMode: boolean;
}

const props = defineProps<Props>();
const emit = defineEmits([
  "on-cancel",
  "on-set-theme", // used to ask parent to set specific theme for preview
  "on-settings-change",
]);

// Options (keys use i18n)
const zoomLevels = [
  { value: 75, labelKey: "zoom_75" },
  { value: 90, labelKey: "zoom_90" },
  { value: 100, labelKey: "zoom_100_default" },
  { value: 110, labelKey: "zoom_110" },
  { value: 125, labelKey: "zoom_125" },
];

const languages = [
  { value: "en-US", labelKey: "language_en" },
  { value: "pt-BR", labelKey: "language_pt_br" },
];

const currencyFormats = [
  { value: "USD", labelKey: "currency_USD", symbol: "$" },
  { value: "EUR", labelKey: "currency_EUR", symbol: "€" },
  { value: "GBP", labelKey: "currency_GBP", symbol: "£" },
  { value: "BRL", labelKey: "currency_BRL", symbol: "R$" },
  { value: "JPY", labelKey: "currency_JPY", symbol: "¥" },
  { value: "CAD", labelKey: "currency_CAD", symbol: "C$" },
  { value: "AUD", labelKey: "currency_AUD", symbol: "A$" },
];

const dateFormats = [
  { value: "MM/DD/YYYY", labelKey: "dateformat_MMDD", example: "12/31/2024" },
  { value: "DD/MM/YYYY", labelKey: "dateformat_DDMM", example: "31/12/2024" },
  { value: "YYYY-MM-DD", labelKey: "dateformat_ISO", example: "2024-12-31" },
  {
    value: "DD.MM.YYYY",
    labelKey: "dateformat_DD_DOT_MM",
    example: "31.12.2024",
  },
];

const itemsPerPageOptions = [5, 10, 15, 20];

// Preview state (local to modal)
const previewZoom = ref<number>(100);
const previewLanguage = ref<string>("en-US");
const previewCurrency = ref<string>("USD");
const previewDateFormat = ref<string>("MM/DD/YYYY");
const previewItemsPerPage = ref<number>(5);
const previewDarkMode = ref<boolean>(false);

// snapshot for revert on cancel
let snapshot: {
  zoom: number;
  language: string;
  currency: string;
  dateFormat: string;
  itemsPerPage: number;
  darkMode: boolean;
} | null = null;

// apply immediate preview side-effects
function applyZoom(z: number) {
  if (typeof document !== "undefined" && document.documentElement) {
    document.documentElement.style.fontSize = `${z}%`;
  }
}

// watchers to visualize changes live (but not persist)
watch(previewZoom, (z) => {
  applyZoom(z);
});

watch(previewLanguage, (lang) => {
  locale.value = lang;
});

// Theme preview: explicitly set to dark or light mode
function setThemePreview(isDark: boolean) {
  if (previewDarkMode.value !== isDark) {
    previewDarkMode.value = isDark;
    // ask parent to set the theme
    emit("on-set-theme", isDark);
  }
}

// Initialize preview state when modal opens
watch(
  () => props.isOpen,
  (isOpen) => {
    if (isOpen) {
      console.log("Settings from store (Pinia persist should have loaded):", {
        zoom: settings.zoom,
        locale: settings.locale,
        currency: settings.currency,
        dateFormat: settings.dateFormat,
        itemsPerPage: settings.itemsPerPage,
        darkMode: settings.darkMode,
      });

      // initialize previews from store (Pinia auto-unwraps refs) with defaults
      previewZoom.value = settings.zoom ?? 100;
      previewLanguage.value = settings.locale ?? "en-US";
      previewCurrency.value = settings.currency ?? "USD";
      previewDateFormat.value = settings.dateFormat ?? "MM/DD/YYYY";
      previewItemsPerPage.value = settings.itemsPerPage ?? 5;
      previewDarkMode.value = props.isDarkMode ?? false;

      console.log("Preview initialized:", {
        zoom: previewZoom.value,
        language: previewLanguage.value,
        currency: previewCurrency.value,
        dateFormat: previewDateFormat.value,
        itemsPerPage: previewItemsPerPage.value,
        darkMode: previewDarkMode.value,
      });

      // snapshot current values AFTER initialization (capture what user sees)
      snapshot = {
        zoom: previewZoom.value,
        language: previewLanguage.value,
        currency: previewCurrency.value,
        dateFormat: previewDateFormat.value,
        itemsPerPage: previewItemsPerPage.value,
        darkMode: previewDarkMode.value,
      };

      console.log("Snapshot created:", snapshot);

      // apply preview side-effects
      applyZoom(previewZoom.value);
      locale.value = previewLanguage.value;
      // if previewDarkMode differs from parent-provided prop, set it explicitly
      if (props.isDarkMode !== previewDarkMode.value) {
        emit("on-set-theme", previewDarkMode.value);
      }
    } else {
      // nothing upon close; Save/Cancel handle persistence/revert
    }
  },
  { immediate: true }
);

// Track if any changes have been made from snapshot
const hasChanges = computed(() => {
  if (!snapshot) {
    console.log("hasChanges: no snapshot");
    return false;
  }
  const changed = (
    previewZoom.value !== snapshot.zoom ||
    previewLanguage.value !== snapshot.language ||
    previewCurrency.value !== snapshot.currency ||
    previewDateFormat.value !== snapshot.dateFormat ||
    previewItemsPerPage.value !== snapshot.itemsPerPage ||
    previewDarkMode.value !== snapshot.darkMode
  );
  console.log("hasChanges:", changed, {
    preview: {
      zoom: previewZoom.value,
      language: previewLanguage.value,
      currency: previewCurrency.value,
      dateFormat: previewDateFormat.value,
      itemsPerPage: previewItemsPerPage.value,
      darkMode: previewDarkMode.value,
    },
    snapshot,
  });
  return changed;
});

// Save: persist preview values to store and ensure side-effects are applied
function saveSettings() {
  console.log("Saving settings:", {
    zoom: previewZoom.value,
    locale: previewLanguage.value,
    currency: previewCurrency.value,
    dateFormat: previewDateFormat.value,
    itemsPerPage: previewItemsPerPage.value,
    darkMode: previewDarkMode.value,
  });

  // commit to store
  settings.setZoom(previewZoom.value);
  settings.setLocale(previewLanguage.value);
  settings.setCurrency(previewCurrency.value);
  settings.setDateFormat(previewDateFormat.value);
  settings.setItemsPerPage(previewItemsPerPage.value);
  settings.setDarkMode(previewDarkMode.value);

  console.log("Settings after save:", {
    zoom: settings.zoom,
    locale: settings.locale,
    currency: settings.currency,
    dateFormat: settings.dateFormat,
    itemsPerPage: settings.itemsPerPage,
    darkMode: settings.darkMode,
  });

  // ensure zoom and locale applied consistently
  settings.applyZoom();
  // keep the i18n locale in sync with the store
  locale.value = settings.locale;

  emit("on-settings-change", {
    zoom: settings.zoom,
    language: settings.locale,
    currency: settings.currency,
    dateFormat: settings.dateFormat,
    itemsPerPage: settings.itemsPerPage,
  });

  emit("on-cancel"); // reuse on-cancel to close the modal
}

// Cancel: revert all preview values to snapshot and ask parent to revert theme if needed
function handleCancel() {
  console.log("Cancelling, reverting to snapshot:", snapshot);
  if (snapshot) {
    // revert all preview values to snapshot
    previewZoom.value = snapshot.zoom;
    previewLanguage.value = snapshot.language;
    previewCurrency.value = snapshot.currency;
    previewDateFormat.value = snapshot.dateFormat;
    previewItemsPerPage.value = snapshot.itemsPerPage;
    previewDarkMode.value = snapshot.darkMode;

    console.log("Preview after revert:", {
      zoom: previewZoom.value,
      language: previewLanguage.value,
      currency: previewCurrency.value,
      dateFormat: previewDateFormat.value,
      itemsPerPage: previewItemsPerPage.value,
      darkMode: previewDarkMode.value,
    });

    // revert side effects
    applyZoom(snapshot.zoom);
    locale.value = snapshot.language;

    // revert theme: if current theme differs from snapshot, set it back
    if (props.isDarkMode !== snapshot.darkMode) {
      emit("on-set-theme", snapshot.darkMode);
    }
  }

  emit("on-cancel");
}

// local helpers for UI binding
function setPreviewZoom(z: number) {
  previewZoom.value = z;
}
function setPreviewLanguage(l: string) {
  previewLanguage.value = l;
}
function setPreviewCurrency(c: string) {
  previewCurrency.value = c;
}
function setPreviewDateFormat(f: string) {
  previewDateFormat.value = f;
}
function setPreviewItemsPerPage(n: number) {
  previewItemsPerPage.value = n;
}
</script>

<template>
  <Dialog :isOpen="isOpen">
    <div class="space-y-6">
      <!-- Header -->
      <div class="text-center">
        <div
          class="w-16 h-16 bg-gradient-ocean rounded-full flex items-center justify-center mx-auto mb-4 shadow-glow-blue"
        >
          <SettingsIcon class="w-8 h-8 text-white" />
        </div>
        <h2 class="text-2xl font-bold text-gradient">
          {{ t("settings") }}
        </h2>
        <p
          class="text-sm text-text-secondary-light dark:text-text-secondary-dark mt-2"
        >
          {{ t("settings_description") }}
        </p>
        <p
          v-if="hasChanges"
          class="text-xs text-primary-600 dark:text-primary-400 mt-1 font-medium"
        >
          {{ t("preview_not_saved") }}
        </p>
      </div>

      <!-- Settings Form -->
      <div class="space-y-4 max-h-96 overflow-y-auto pr-2">
        <!-- Zoom Level -->
        <div class="space-y-2">
          <label
            class="block text-sm font-medium text-text-primary-light dark:text-text-primary-dark"
          >
            {{ t("zoom_level") }}
          </label>
          <div class="grid grid-cols-5 gap-2">
            <button
              v-for="zoom in zoomLevels"
              :key="zoom.value"
              @click="setPreviewZoom(zoom.value)"
              class="px-3 py-2 rounded-lg text-xs font-medium transition-all duration-200"
              :class="{
                'bg-gradient-ocean text-white shadow-md-blue':
                  previewZoom === zoom.value,
                'bg-gray-100 dark:bg-surface-dark text-text-primary-light dark:text-text-primary-dark hover:bg-gray-200 dark:hover:bg-opacity-80':
                  previewZoom !== zoom.value,
              }"
            >
              {{ t(zoom.labelKey) }}
            </button>
          </div>
        </div>

        <!-- Theme -->
        <div class="space-y-2">
          <label
            class="block text-sm font-medium text-text-primary-light dark:text-text-primary-dark"
          >
            {{ t("theme") }}
          </label>
          <div class="grid grid-cols-2 gap-2">
            <button
              @click="setThemePreview(false)"
              class="px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200 flex items-center justify-center gap-2"
              :class="{
                'bg-gradient-ocean text-white shadow-md-blue':
                  !previewDarkMode,
                'bg-gray-100 dark:bg-surface-dark text-text-primary-light dark:text-text-primary-dark hover:bg-gray-200 dark:hover:bg-opacity-80':
                  previewDarkMode,
              }"
            >
              <svg
                class="w-5 h-5"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z"
                  clip-rule="evenodd"
                ></path>
              </svg>
              {{ t("light_mode") }}
            </button>
            <button
              @click="setThemePreview(true)"
              class="px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200 flex items-center justify-center gap-2"
              :class="{
                'bg-gradient-ocean text-white shadow-md-blue':
                  previewDarkMode,
                'bg-gray-100 dark:bg-surface-dark text-text-primary-light dark:text-text-primary-dark hover:bg-gray-200 dark:hover:bg-opacity-80':
                  !previewDarkMode,
              }"
            >
              <svg
                class="w-5 h-5"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"
                ></path>
              </svg>
              {{ t("dark_mode") }}
            </button>
          </div>
        </div>

        <!-- Language -->
        <div class="space-y-2">
          <label
            class="block text-sm font-medium text-text-primary-light dark:text-text-primary-dark"
          >
            {{ t("language_label") }}
          </label>
          <div class="grid grid-cols-2 gap-2">
            <button
              v-for="lang in languages"
              :key="lang.value"
              @click="setPreviewLanguage(lang.value)"
              class="px-4 py-3 rounded-lg text-sm font-medium transition-all duration-200"
              :class="{
                'bg-gradient-ocean text-white shadow-md-blue':
                  previewLanguage === lang.value,
                'bg-gray-100 dark:bg-surface-dark text-text-primary-light dark:text-text-primary-dark hover:bg-gray-200 dark:hover:bg-opacity-80':
                  previewLanguage !== lang.value,
              }"
            >
              {{ t(lang.labelKey) }}
            </button>
          </div>
        </div>

        <!-- Currency Format -->
        <div class="space-y-2">
          <label
            class="block text-sm font-medium text-text-primary-light dark:text-text-primary-dark"
          >
            {{ t("currency_format_label") }}
          </label>
          <select
            v-model="previewCurrency"
            class="w-full px-4 py-3 rounded-lg text-sm bg-white dark:bg-surface-dark border-2 border-border-light dark:border-border-dark text-text-primary-light dark:text-text-primary-dark focus:border-primary-500 dark:focus:border-primary-400 focus:ring-4 focus:ring-primary-100 dark:focus:ring-primary-900/30 transition-all duration-200 ease-out shadow-sm focus:shadow-md-blue cursor-pointer"
          >
            <option
              v-for="curr in currencyFormats"
              :key="curr.value"
              :value="curr.value"
            >
              {{ t(curr.labelKey) }}
            </option>
          </select>
        </div>

        <!-- Date Format -->
        <div class="space-y-2">
          <label
            class="block text-sm font-medium text-text-primary-light dark:text-text-primary-dark"
          >
            {{ t("date_format_label") }}
          </label>
          <select
            v-model="previewDateFormat"
            class="w-full px-4 py-3 rounded-lg text-sm bg-white dark:bg-surface-dark border-2 border-border-light dark:border-border-dark text-text-primary-light dark:text-text-primary-dark focus:border-primary-500 dark:focus:border-primary-400 focus:ring-4 focus:ring-primary-100 dark:focus:ring-primary-900/30 transition-all duration-200 ease-out shadow-sm focus:shadow-md-blue cursor-pointer"
          >
            <option
              v-for="format in dateFormats"
              :key="format.value"
              :value="format.value"
            >
              {{ t(format.labelKey) }} - {{ format.example }}
            </option>
          </select>
        </div>

        <!-- Items Per Page -->
        <div class="space-y-2">
          <label
            class="block text-sm font-medium text-text-primary-light dark:text-text-primary-dark"
          >
            {{ t("default_items_per_page") }}
          </label>
          <div class="grid grid-cols-4 gap-2">
            <button
              v-for="items in itemsPerPageOptions"
              :key="items"
              @click="setPreviewItemsPerPage(items)"
              class="px-3 py-2 rounded-lg text-sm font-medium transition-all duration-200"
              :class="{
                'bg-gradient-ocean text-white shadow-md-blue':
                  previewItemsPerPage === items,
                'bg-gray-100 dark:bg-surface-dark text-text-primary-light dark:text-text-primary-dark hover:bg-gray-200 dark:hover:bg-opacity-80':
                  previewItemsPerPage !== items,
              }"
            >
              {{ items }}
            </button>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div
        class="flex gap-3 pt-4 border-t border-border-light dark:border-border-dark"
      >
        <CancelButton @click="handleCancel" class="flex-1">{{
          t("cancel")
        }}</CancelButton>
        <ConfirmButton @click="saveSettings" :disabled="!hasChanges" class="flex-1">{{
          t("save_settings")
        }}</ConfirmButton>
      </div>
    </div>
  </Dialog>
</template>
