import { loadTranslations, locale } from '../lib/translations';
import type { LayoutLoad } from './$types';

export const load = (async ({ url }) => {
	const { pathname } = url;
	const defaultLocale = window.navigator.language.startsWith('zh') ? 'zh' : 'en';
	const initLocale = locale.get() || defaultLocale;
	await loadTranslations(initLocale, pathname);

	return {};
}) satisfies LayoutLoad;

export const ssr = false;

export const prerender = true;
