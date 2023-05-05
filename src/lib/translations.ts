import i18n, { type Config } from 'sveltekit-i18n';

const config = {
	loaders: [
		{
			locale: 'zh',
			key: 'common',
			loader: async () => (await import('./zh/common.json')).default
		},
		{
			locale: 'zh',
			key: 'home',
			routes: ['/'],
			loader: async () => (await import('./zh/home.json')).default
		},
		{
			locale: 'zh',
			key: 'lyrics',
			routes: [/^\/.+/],
			loader: async () => (await import('./zh/lyrics.json')).default
		},
		{
			locale: 'en',
			key: 'common',
			loader: async () => (await import('./en/common.json')).default
		},
		{
			locale: 'en',
			key: 'home',
			routes: ['/'],
			loader: async () => (await import('./en/home.json')).default
		},
		{
			locale: 'en',
			key: 'lyrics',
			routes: [/^\/.+/],
			loader: async () => (await import('./en/lyrics.json')).default
		}
	]
} satisfies Config;

export const { t, locale, locales, loading, loadTranslations } = new i18n(config);
