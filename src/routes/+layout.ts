import { loadTranslations, locale } from '../lib/translations';
import type { LayoutLoad } from './$types';

export const load = (async ({ url }) => {
	const { pathname } = url;

	const defaultLocale = 'zh';
	const initLocale = locale.get() || defaultLocale;
	await loadTranslations(initLocale, pathname);

	return {};
}) satisfies LayoutLoad;
