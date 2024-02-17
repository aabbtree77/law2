/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./index.html", "./**/*.{html,js}"],
  theme: {
    extend: {},
  },
  //plugins: [require('@tailwindcss/typography'),],
  plugins: [require("daisyui"),require('@tailwindcss/typography'),],
  /*
  daisyui: {
    themes: []
  },
  */
}

