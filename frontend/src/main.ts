import { createApp } from "vue";
import App from "./App.vue";
import "./index.css";
import { createPinia } from "pinia";
import piniaPluginPersistedState from "pinia-plugin-persistedstate";
import { clickOutside } from "./shared/directives/ClickOutside";
import { i18n } from "./i18n";

const pinia = createPinia();
pinia.use(piniaPluginPersistedState);

createApp(App)
  .directive("click-outside", clickOutside)
  .use(pinia)
  .use(i18n)
  .mount("#app");
