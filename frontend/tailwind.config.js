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
        // Primary dolphin blues
        primary: {
          DEFAULT: '#1E88E5',
          light: '#42A5F5',
          dark: '#1565C0',
          50: '#E3F2FD',
          100: '#BBDEFB',
          200: '#90CAF9',
          300: '#64B5F6',
          400: '#42A5F5',
          500: '#1E88E5',
          600: '#1976D2',
          700: '#1565C0',
          800: '#0D47A1',
          900: '#0A3D91',
        },
        // Secondary ocean blues
        secondary: {
          DEFAULT: '#1565C0',
          light: '#2196F3',
          dark: '#0D47A1',
          50: '#E3F2FD',
          100: '#BBDEFB',
          200: '#90CAF9',
          300: '#64B5F6',
          400: '#42A5F5',
          500: '#1565C0',
          600: '#1976D2',
          700: '#1565C0',
          800: '#0D47A1',
          900: '#0A3D91',
        },
        // Accent colors
        accent: {
          DEFAULT: '#42A5F5',
          light: '#64B5F6',
          dark: '#1E88E5',
        },
        // Surface colors
        surface: {
          light: '#F5F9FF',
          dark: '#1A2332',
        },
        // Background colors
        bg: {
          light: '#E3F2FD',
          dark: '#0D1520',
        },
        // Border colors
        border: {
          light: '#90CAF9',
          dark: '#2D3F5F',
        },
        // Semantic colors
        success: '#14B8A6',
        danger: '#F43F5E',
        warning: '#F59E0B',
        info: '#3B82F6',
        // Text colors
        'text-primary': {
          light: '#0A3D91',
          dark: '#E3F2FD',
        },
        'text-secondary': {
          light: '#546E7A',
          dark: '#B0BEC5',
        },
        // Legacy compatibility
        dark: "#1c1c1c",
        light: "#eaeaea",
        "light-text": "#E8E8E8",
      },
      // Custom gradients
      backgroundImage: {
        'gradient-ocean': 'linear-gradient(135deg, #1E88E5 0%, #1565C0 100%)',
        'gradient-ocean-light': 'linear-gradient(135deg, #42A5F5 0%, #1E88E5 100%)',
        'gradient-surface': 'linear-gradient(180deg, rgba(255,255,255,0.05) 0%, rgba(255,255,255,0) 100%)',
        'gradient-dark-ocean': 'linear-gradient(135deg, #1A2332 0%, #0D1520 100%)',
      },
      // Blue-tinted shadows
      boxShadow: {
        'sm-blue': '0 1px 2px 0 rgba(30, 136, 229, 0.05)',
        'blue': '0 4px 6px -1px rgba(30, 136, 229, 0.1), 0 2px 4px -1px rgba(30, 136, 229, 0.06)',
        'md-blue': '0 10px 15px -3px rgba(30, 136, 229, 0.15), 0 4px 6px -2px rgba(30, 136, 229, 0.08)',
        'lg-blue': '0 20px 25px -5px rgba(30, 136, 229, 0.15), 0 10px 10px -5px rgba(30, 136, 229, 0.08)',
        'xl-blue': '0 25px 50px -12px rgba(30, 136, 229, 0.25)',
        'inner-blue': 'inset 0 2px 4px 0 rgba(30, 136, 229, 0.06)',
        'glow-blue': '0 0 20px rgba(30, 136, 229, 0.3)',
        'glow-blue-lg': '0 0 40px rgba(30, 136, 229, 0.4)',
      },
    },
  },
  plugins: [],
};
