/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "selector",
  content: [
    "./pages/**/*.{ts,tsx,vue}",
    "./components/**/*.{ts,tsx,vue}",
    "./app/**/*.{ts,tsx,vue}",
    "./src/**/*.{ts,tsx,vue}",
  ],
  prefix: "",
  theme: {
    extend: {
      colors: {
        primary: "#30C5D5",
        secondary: "#0090A0",
        dark: "#1c1c1c",
        light: "#eaeaea",
        "light-text": "#E8E8E8",
      },
    },
  },
  plugins: [],
};
