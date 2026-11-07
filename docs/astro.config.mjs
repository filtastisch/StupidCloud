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
			sidebar: [
				{
					label: 'Guides',
					translations: { de: 'Leitfäden' },
					items: [
						{ label: 'Getting started', translations: { de: 'Erste Schritte' }, slug: 'guides/getting-started' },
						{ label: 'Architecture', translations: { de: 'Architektur' }, slug: 'guides/architecture' },
						{ label: 'Development', translations: { de: 'Entwicklung' }, slug: 'guides/development' },
						{ label: 'Configuration', translations: { de: 'Konfiguration' }, slug: 'guides/configuration' },
						{ label: 'Security', translations: { de: 'Sicherheit' }, slug: 'guides/security' },
					],
				},
				{
					label: 'Reference',
					translations: { de: 'Referenz' },
					items: [
						{ label: 'Project status', translations: { de: 'Projektstatus' }, slug: 'reference/project-status' },
						{ label: 'Roadmap', translations: { de: 'Roadmap' }, slug: 'reference/roadmap' },
						{ label: 'Glossary', translations: { de: 'Glossar' }, slug: 'reference/glossary' },
					],
				},
			],
		}),
	],
});
