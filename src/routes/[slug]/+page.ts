import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

interface Lyrics {
	title: string;
	album: string;
	artist: string;
	cover: string;
	editors: LyricsLink[];
	author?: string;
	relations: LyricsLink[];
	scripts: LyricsScript[];
}
interface LyricsLink {
	name: string;
	href: string;
}
interface LyricsScript {
	time: string;
	text: string;
}

function parseLrc(text: string) {
	const lryics: Lyrics = {
		title: '',
		album: '',
		artist: '',
		cover: '',
		editors: [],
		author: '',
		relations: [],
		scripts: []
	};

	const lines = text.split('\n');

	for (const line of lines) {
		const parsed = parseLrcLine(line);
		if (parsed) {
			if (parsed.key === false) {
				lryics.scripts.push({
					time: parsed.value,
					text: parsed.text
				});
			} else {
				let link: LyricsLink | null;
				switch (parsed.key) {
					case 'ti':
						lryics.title = parsed.value;
						break;
					case 'al':
						lryics.album = parsed.value;
						break;
					case 'ar':
						lryics.artist = parsed.value;
						break;
					case 'cv':
						lryics.cover = parsed.value;
						break;
					case 're':
						link = parseLrcLink(parsed.value);
						if (link) lryics.editors.push(link);
						break;
					case 'by':
						lryics.author = parsed.value;
						break;
					case 'rel':
						link = parseLrcLink(parsed.value);
						if (link) {
							link.href = new URL(link.href).pathname.replace(/\.lrc/, '');
							lryics.relations.push(link);
						}
						break;
				}
			}
		}
	}

	return lryics;
}

function parseLrcLine(line: string):
	| {
			key: string;
			value: string;
	  }
	| {
			key: false;
			value: string;
			text: string;
	  }
	| null {
	line = line.trim();

	if (line === '') return null;

	const meta = line.match(/^\[([a-z]+):\s*(.*)\]$/);
	if (meta) {
		return { key: meta[1], value: meta[2] };
	}

	const script = line.match(/^\[([\d:.-]*)\]\s*(.*)$/);
	if (script) {
		return { key: false, value: script[1], text: script[2] };
	}

	return null;
}

function parseLrcLink(text: string): LyricsLink | null {
	const link = text.match(/^(.+):\s*(https?:\/\/.*)$/);
	if (link) {
		return { name: link[1], href: link[2] };
	}
	return null;
}

export interface WikiMoreLikeResponse {
	batchcomplete: string;
	continue: WikiMoreLikeContinue;
	query: WikiMoreLikeQuery;
}

export interface WikiMoreLikeContinue {
	gsroffset: number;
	continue: string;
}

export interface WikiMoreLikeQuery {
	pages: { [key: string]: WikiMoreLikePage };
}

export interface WikiMoreLikePage {
	pageid: number;
	ns: number;
	title: string;
	index: number;
	thumbnail: WikiMoreLikeThumbnail;
}

export interface WikiMoreLikeThumbnail {
	source: string;
	width: number;
	height: number;
}

export const load = (async ({ params, fetch }) => {
	const response = await fetch(`https://lyrics.thwiki.cc/${params.slug}.lrc`);

	if (!response.ok) {
		throw error(404, await response.text());
	}

	const lrc = await response.text();
	const lyrics = parseLrc(lrc);

	return {
		lyrics
	};
}) satisfies PageLoad;
