const colors = require("tailwindcss/colors");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["components/*.templ"],
  theme: {
    fontFamily: {
      "geist-light": ["Geist-Light", "sans-serif"],
      "geist-mono-light": ["GeistMono-Light", "sans-serif"],
    },
    keyframes: {
      "fade-in": {
        "0%": { opacity: "0" },
        "100%": { opacity: "1" },
      },
    },
    animation: {
      "fade-in": "fade-in 500ms ease-in-out forwards",
    },
    extend: {
      colors: {},
    },
  },
  plugins: [require("@tailwindcss/forms"), require("@tailwindcss/typography")],
};
