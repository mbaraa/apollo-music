const defaultTheme = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
export default {
  mode: "jit",
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    colors: {
      dark: {
        primary: "#273d50",
        secondary: "#fff",
        accent: "#add8fb",
        accent2: "#add8fb66",
        neutral: "#051220",
      },
      light: {},
      ...defaultTheme.colors,
    },
  },
  extend: {
    fontFamily: {
      Comfortaa: ["Comfortaa", "sans"],
      IBMPlexSans: ["IBM Plex Sans", "sans"],
    },
  },
  plugins: [],
};
