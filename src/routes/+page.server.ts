import {
	RANKING_POPULAR_LIMIT,
	RANKING_POPULAR_REQUEST_LIMIT,
	RANKING_RECENT_LIMIT,
	RANKING_RECENT_NAMESPACE,
	VERCEL_OWNER_ID,
	VERCEL_PROJECT_ID,
	VERCEL_TEAM_ID,
	VERCEL_TOKEN
} from '$env/static/private';
import type { Config } from '@sveltejs/adapter-vercel';
import type { PageServerLoad } from './$types';

interface WikiRecentChangesResponse {
	batchcomplete: boolean;
	continue: WikiContinue;
	query: WikiQuery;
}

interface WikiContinue {
	grccontinue: string;
	continue: string;
}

interface WikiQuery {
	pages: WikiPage[];
}

interface WikiPage {
	pageid: number;
	ns: number;
	title: string;
	revisions: WikiRevision[];
}

interface WikiRevision {
	timestamp: string;
}

interface VercelUsageTopResponse {
	data: VercelUsageTopDatum[];
}

interface VercelUsageTopDatum {
	title: string;
	target_path: string;
	host: string;
	type: string;
	bandwidth: number;
	execution_time: number;
	requests: number;
	invocations: number;
	bandwidth_percentage: number;
	execution_time_percentage: number;
	requests_percentage: number;
	invocations_percentage: number;
}

enum Language {
	ALL = 'all',
	ZH = 'zh',
	JA = 'ja'
}

interface LrcRequest {
	title: string;
	index: number;
	language: Language;
	extension: string;
}

function parseLanguage(text: string): Language | undefined {
	text = text.toUpperCase().trim();
	if (Object.prototype.hasOwnProperty.call(Language, text)) {
		return text as Language;
	}
	return undefined;
}

function parseName(name: string) {
	const request: LrcRequest = { title: '', index: 0, language: Language.JA, extension: '' };
	const parts = name.split('.');
	let lastIndex = parts.length - 1;

	request.index = 1;
	request.language = Language.JA;
	request.extension = 'lrc';

	if (lastIndex > 0) {
		if (parts[lastIndex] == 'lrc') {
			request.extension = parts[lastIndex];
			lastIndex--;
		}
	}

	if (lastIndex > 0) {
		const language = parseLanguage(parts[lastIndex]);
		if (language) {
			request.language = language;
			lastIndex--;
		}
	}

	if (lastIndex > 0 && parts[lastIndex].length <= 2) {
		const index = parseInt(parts[lastIndex], 10);
		if (!isNaN(index)) {
			request.index = index;
			lastIndex--;
		}
	}

	request.title = parts.slice(0, lastIndex + 1).join('.');

	return request;
}

export const load = (async ({ fetch }) => {
	const recentChangesUrl = new URL('https://thwiki.cc/api.php');
	const recentChangesUrlSearchParams = recentChangesUrl.searchParams;
	recentChangesUrlSearchParams.set('action', 'query');
	recentChangesUrlSearchParams.set('format', 'json');
	recentChangesUrlSearchParams.set('prop', 'revisions');
	recentChangesUrlSearchParams.set('generator', 'recentchanges');
	recentChangesUrlSearchParams.set('formatversion', '2');
	recentChangesUrlSearchParams.set('rvprop', 'timestamp');
	recentChangesUrlSearchParams.set('grcdir', 'older');
	recentChangesUrlSearchParams.set('grcnamespace', RANKING_RECENT_NAMESPACE);
	recentChangesUrlSearchParams.set('grcprop', '');
	recentChangesUrlSearchParams.set('grcshow', '');
	recentChangesUrlSearchParams.set('grclimit', RANKING_RECENT_LIMIT);
	recentChangesUrlSearchParams.set('grctype', 'edit|new');
	recentChangesUrlSearchParams.set('grctoponly', '1');

	const recentChanges: WikiRecentChangesResponse = await (
		await fetch(recentChangesUrl.href)
	).json();

	const toDate = new Date();
	const fromDate = new Date(new Date().setMonth(toDate.getMonth() - 2));
	const vercelUsageTopUrl = new URL('https://vercel.com/api/v4/usage/top');
	const vercelUsageTopUrlSearchParams = vercelUsageTopUrl.searchParams;
	vercelUsageTopUrlSearchParams.set('owner_id', VERCEL_OWNER_ID);
	vercelUsageTopUrlSearchParams.set('teamId', VERCEL_TEAM_ID);
	vercelUsageTopUrlSearchParams.set('projectId', VERCEL_PROJECT_ID);
	vercelUsageTopUrlSearchParams.set('from', fromDate.toISOString());
	vercelUsageTopUrlSearchParams.set('to', toDate.toISOString());
	vercelUsageTopUrlSearchParams.set('limit', RANKING_POPULAR_REQUEST_LIMIT);
	vercelUsageTopUrlSearchParams.set('sortKey', 'requests');
	vercelUsageTopUrlSearchParams.set('pathType', 'request_path');

	const popularItems: VercelUsageTopResponse = await (
		await fetch(vercelUsageTopUrl.href, {
			method: 'GET',
			headers: [['Authorization', `Bearer ${VERCEL_TOKEN}`]],
			redirect: 'follow'
		})
	).json();

	const members = [
		14817, 7355, 23839, 13659, 6793, 29475, 19038, 16099, 15610, 21774, 16179, 8879, 23283, 33173,
		27072, 15455, 42900, 61758, 2036, 19707, 5151, 234, 7154, 17698, 6933, 49179, 9324, 5146, 16384,
		49167
	];

	return {
		latest: recentChanges.query.pages
			.map((page) => ({
				title: page.title.substring(3),
				value: page.revisions[0].timestamp
			}))
			.sort((a, b) => b.value.localeCompare(a.value)),
		popular: popularItems.data
			.filter((item) => {
				if (item.type !== 'func') return false;

				try {
					item.title = decodeURIComponent(item.target_path);
					if (/[\u007F-\u00FF]/.test(item.title)) return false;

					item.title = parseName(item.title.substring(1)).title;
					return item.title !== '';
				} catch (_: unknown) {
					return false;
				}
			})
			.filter((item, _, a) => {
				const firstItem = a.find((search) => search.title === item.title);
				if (firstItem == null) return false;

				if (firstItem === item) {
					return true;
				}
				firstItem.requests += item.requests;
				return false;
			})
			.map((item) => ({
				title: item.title,
				value: item.requests
			}))
			.sort((a, b) => b.value - a.value)
			.slice(0, parseInt(RANKING_POPULAR_LIMIT)),
		members,
		timestamp: Date.now()
	};
}) satisfies PageServerLoad;

export const config: Config = {
	isr: {
		expiration: 600
	}
};
