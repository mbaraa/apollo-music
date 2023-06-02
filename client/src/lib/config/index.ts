const config: { [keys: string]: string } = {
	backendAddress: import.meta.env.VITE_BACKEND_ADDRESS,
	landingAddress: import.meta.env.VITE_LANDING_ADDRESS ?? "http://localhost:1413",
	defaultLocale: "en"
};

export default config;
