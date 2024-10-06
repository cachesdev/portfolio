import catppuccin from "@catppuccin/daisyui";

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.templ",
    "./views/**/*.go",
    "./posts/**/*.md",
    "./pkg/**/*.go",
  ],
  theme: {
    extend: {
      fontFamily: {
        cascadia: ['"Cascadia Code"', "monospace"],
        lato: ["Lato", "sans-serif"],
        sourcesans: ['"Source Sans 3"', "sans-serif"],
        poppins: ["Poppins", "sans-serif"],
      },
    },
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  safelist: [],
  daisyui: {
    themes: [
      {
        gruvbox: {
          primary: "oklch(72.5597% 0.084558 141.092689 /1)",

          "primary-content": "#070c06",

          secondary: "oklch(0.75 0.11 116.51)",

          "secondary-content": "#0b0c03",

          accent: "oklch(0.71 0.1 2.2)",

          "accent-content": "#100609",

          // neutral: "oklch(0.34 0.01 0)",
          neutral: "oklch(0.30 0.01 0)",

          "neutral-content": "oklch(0.81 0.06 81.19)",

          "base-100": "#3c3836",

          "base-200": "#2a292e",

          "base-300": "#252429",

          "base-content": "oklch(0.81 0.06 81.19)",

          info: "oklch(0.71 0.05 179.17)",

          "info-content": "#060b0a",

          success: "oklch(0.75 0.11 116.51)",

          "success-content": "#0b0c03",

          warning: "oklch(0.76 0.11 77.03)",

          "warning-content": "#110a03",

          error: "oklch(0.68 0.16 25.42)",

          "error-content": "#111827",
        },
      },
      "retro",
      "luxury",
    ],
    darkTheme: "gruvbox",
  },
  darkMode: ["class", '[data-theme="gruvbox"]'],
};
