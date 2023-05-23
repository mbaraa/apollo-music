const defaultTheme = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
export default {
	content: ["./src/**/*.{html,js,svelte,ts}"],
	theme: {
		colors: {
			dark: {
				primary: "#051220",
				secondary: "#fff",
				accent: "#add8fb",
				accent2: "#add8fb66",
				neutral: "#273d50"
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
