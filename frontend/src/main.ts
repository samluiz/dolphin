import { createApp } from "vue";
import App from "./App.vue";
import "./index.css";
import { createPinia } from "pinia";
import piniaPluginPersistedState from "pinia-plugin-persistedstate";
import { clickOutside } from "./shared/directives/ClickOutside";

const pinia = createPinia();
pinia.use(piniaPluginPersistedState);

createApp(App)
  .directive("click-outside", clickOutside)
  .use(pinia)
  .mount("#app");
