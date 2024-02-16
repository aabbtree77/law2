/** @type {import('tailwindcss').Config} */

const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  content: ["./**/*.{html,js}"],
  plugins: [require("daisyui"), require("@tailwindcss/typography")],
  theme: {
    extend: {
      fontFamily: {
        sourcecodepro: ['"Source Code Pro"', ...defaultTheme.fontFamily.sans],
        caveat: ['"Caveat"', ...defaultTheme.fontFamily.sans],
      },
    },
  },
};
