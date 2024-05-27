import { createI18n } from "vue-i18n";
import { messages } from "./messages";
import { datetimeFormats } from "./datetime";
import { numberFormats } from "./number";

export const i18n = createI18n({
  legacy: false,
  locale: "pt-BR",
  fallbackLocale: "en-US",
  messages,
  datetimeFormats,
  numberFormats,
});
