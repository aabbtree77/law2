/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./index.html", "./**/*.html"],
  theme: {
    extend: {},
  },
  //plugins: [require('@tailwindcss/typography'),],
  plugins: [require('@tailwindcss/typography'),],
}

