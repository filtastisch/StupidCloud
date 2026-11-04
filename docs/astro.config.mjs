// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	// GitHub Pages serves this project from /StupidCloud; local development uses /.
	base: process.env.GITHUB_ACTIONS === 'true' ? '/StupidCloud' : '/',
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
