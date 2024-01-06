/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["**/*.{html,templ}"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["dark", "nord"],
  },
};
