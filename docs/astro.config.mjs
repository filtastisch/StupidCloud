// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	// This site uses a custom domain, so it is served from the domain root.
	base: '/',
	integrations: [
		starlight({
			title: {
				en: 'StupidCloud Documentation',
				de: 'StupidCloud Dokumentation',
			},
			defaultLocale: 'root',
			locales: {
				root: { label: 'English', lang: 'en' },
				de: { label: 'Deutsch', lang: 'de' },
			},
			social: [],
		}),
	],
});
