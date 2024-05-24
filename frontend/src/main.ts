import { createApp } from "vue";
import App from "./App.vue";
import "./index.css";
import { createPinia } from "pinia";
import piniaPluginPersistedState from "pinia-plugin-persistedstate";

const pinia = createPinia();
pinia.use(piniaPluginPersistedState);

createApp(App).use(pinia).mount("#app");
