/**
 * dolphin\frontend\src\stores\SettingsStore.ts
 *
 * Pinia store to manage and persist UI settings (zoom, locale, currency, date format,
 * items per page, dark mode). Based on the style of ProfileStore.ts and configured
 * to persist automatically via Pinia persistence plugin.
 */

import { defineStore } from "pinia";
import { Ref, ref } from "vue";

export interface SettingsState {
  zoom: number;
  locale: string;
  currency: string;
  dateFormat: string;
  itemsPerPage: number;
  darkMode: boolean;
}

const DEFAULTS: SettingsState = {
  zoom: 100,
  locale: "en-US",
  currency: "USD",
  dateFormat: "MM/DD/YYYY",
  itemsPerPage: 5,
  darkMode: false,
};

export const useSettingsStore = defineStore(
  "settings",
  () => {
    // state
    const zoom: Ref<number> = ref(DEFAULTS.zoom);
    const locale: Ref<string> = ref(DEFAULTS.locale);
    const currency: Ref<string> = ref(DEFAULTS.currency);
    const dateFormat: Ref<string> = ref(DEFAULTS.dateFormat);
    const itemsPerPage: Ref<number> = ref(DEFAULTS.itemsPerPage);
    const darkMode: Ref<boolean> = ref(DEFAULTS.darkMode);

    // Load from localStorage (useful if Pinia persist plugin isn't available
    // or for initial population before plugin restores state)
    const fetchSettings = () => {
      const storedZoom = localStorage.getItem("zoomLevel");
      const storedLocale = localStorage.getItem("locale");
      const storedCurrency = localStorage.getItem("currencyFormat");
      const storedDateFormat = localStorage.getItem("dateFormat");
      const storedItemsPerPage = localStorage.getItem("itemsPerPage");
      const storedDarkMode = localStorage.getItem("darkMode");

      if (storedZoom) zoom.value = parseInt(storedZoom);
      if (storedLocale) locale.value = storedLocale;
      if (storedCurrency) currency.value = storedCurrency;
      if (storedDateFormat) dateFormat.value = storedDateFormat;
      if (storedItemsPerPage) itemsPerPage.value = parseInt(storedItemsPerPage);
      if (storedDarkMode) darkMode.value = storedDarkMode === "true";
    };

    // Persist to localStorage (Pinia persist may already handle this, but keeping
    // an explicit method is handy when you want to force-write immediately)
    const saveSettings = () => {
      localStorage.setItem("zoomLevel", zoom.value.toString());
      localStorage.setItem("locale", locale.value);
      localStorage.setItem("currencyFormat", currency.value);
      localStorage.setItem("dateFormat", dateFormat.value);
      localStorage.setItem("itemsPerPage", itemsPerPage.value.toString());
      localStorage.setItem("darkMode", darkMode.value ? "true" : "false");
    };

    // Individual setters that also persist and apply side-effects where appropriate
    const setZoom = (value: number) => {
      zoom.value = value;
      applyZoom();
      saveSettings();
    };

    const applyZoom = () => {
      // Apply zoom by changing root font-size (%)
      // Keep this safe in environments where document may be undefined (SSR)
      if (typeof document !== "undefined" && document.documentElement) {
        document.documentElement.style.fontSize = `${zoom.value}%`;
      }
    };

    const setLocale = (value: string) => {
      locale.value = value;
      saveSettings();
    };

    /**
     * If you want this store to directly update a vue-i18n locale ref, pass it in.
     * Example from a component:
     *   const { locale } = useI18n();
     *   settingsStore.applyLocaleTo(i18nLocaleRef);
     */
    const applyLocaleTo = (i18nLocaleRef: Ref<string> | { value: string } | null) => {
      if (i18nLocaleRef && typeof i18nLocaleRef === "object" && "value" in i18nLocaleRef) {
        i18nLocaleRef.value = locale.value;
      }
    };

    const setCurrency = (value: string) => {
      currency.value = value;
      saveSettings();
    };

    const setDateFormat = (value: string) => {
      dateFormat.value = value;
      saveSettings();
    };

    const setItemsPerPage = (value: number) => {
      itemsPerPage.value = value;
      saveSettings();
    };

    const toggleDarkMode = () => {
      darkMode.value = !darkMode.value;
      saveSettings();
    };

    const setDarkMode = (value: boolean) => {
      darkMode.value = value;
      saveSettings();
    };

    const resetToDefaults = () => {
      zoom.value = DEFAULTS.zoom;
      locale.value = DEFAULTS.locale;
      currency.value = DEFAULTS.currency;
      dateFormat.value = DEFAULTS.dateFormat;
      itemsPerPage.value = DEFAULTS.itemsPerPage;
      darkMode.value = DEFAULTS.darkMode;
      applyZoom();
      saveSettings();
    };

    // Return state + actions
    return {
      // state
      zoom,
      locale,
      currency,
      dateFormat,
      itemsPerPage,
      darkMode,

      // actions
      fetchSettings,
      saveSettings,
      setZoom,
      applyZoom,
      setLocale,
      applyLocaleTo,
      setCurrency,
      setDateFormat,
      setItemsPerPage,
      toggleDarkMode,
      setDarkMode,
      resetToDefaults,
    };
  },
  {
    persist: true,
  },
);
