/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        text: '#181818',
        text2: '#636363',
        blue: '#2A94F4',
        yellow: '#FFBE55',
        gray: '#F2F2F2',
      },
      backgroundImage: {
        'home': "url('/public/image/home.png')",
        'preference1': "url('/public/image/preference-1.png')",
        'preference2': "url('/public/image/preference-2.png')",
        'bridging': "url('/public/image/bridging.png')",
        'bridging2': "url('/public/image/bridging-2.png')",
        'bridging3': "url('/public/image/bridging-3.png')",
        'buddy': "url('/public/image/buddy.png')",
        'buddy2': "url('/public/image/buddy-2.png')",
      },
      fontFamily: {
        'sans': ['"Poppins"']
      }
    },
  },
  plugins: [],
}

