const defaultTheme = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
export default {
	content: ["./src/**/*.{html,js,svelte,ts}"],
	theme: {
		colors: {
			dark: {
				primary: "#273d50",
				secondary: "#fff",
				accent: "#add0fb",
				neutral: "#051220"
			},
			light: {},
			...defaultTheme.colors
		},
		extend: {
			fontFamily: {
				Comfortaa: ["Comfortaa", "sans"],
				IBMPlexSans: ["IBM Plex Sans", "sans"]
			}
		}
	},
	plugins: []
};
