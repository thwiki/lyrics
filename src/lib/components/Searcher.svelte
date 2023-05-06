<script lang="ts">
	import { copy } from '../../utils/copy';
	import Select from './Select.svelte';
	import Typeahead from './Typeahead.svelte';

	export let prefix: string;
	export let suffix: string;

	export let input: HTMLInputElement;
	let name: string = '';
	let lang: string = '.all';

	let full: string | undefined = undefined;
	let page: string | undefined = undefined;

	$: {
		page = name ? `https://${prefix}${name}` : undefined;
		full = name ? `https://${prefix}${name}${lang}${suffix}` : undefined;
	}

	$: {
		name;
		if (input) {
			input.style.width = 'auto';
			input.style.width = `${input.scrollWidth}px`;
		}
	}

	async function handleCopy() {
		if (full) {
			await copy(full);
		}
	}

	type Result = {
		value: string;
		label: string;
		disabled?: boolean;
	}[];

	const searchCache = new Map<string, Result>();

	async function handleSearch(term: string): Promise<Result> {
		if (term.trim() === '') return [];

		if (searchCache.has(term)) {
			return searchCache.get(term) ?? [];
		}

		const url = new URL('https://cache.thwiki.cc/api.php');
		const { searchParams } = url;
		searchParams.set('action', 'opensearch');
		searchParams.set('format', 'json');
		searchParams.set('namespace', '512');
		searchParams.set('limit', '8');
		searchParams.set('search', '歌词:' + term);
		const data: {
			text: string;
			link: string;
		}[] = await (await fetch(url.href)).json();

		const result = data.map((item) => ({
			value: item.text.substring(3),
			label: item.text.substring(3)
		}));

		searchCache.set(term, result);

		return result;
	}
</script>

<div class="w-full">
	<div class="flex flex-col items-center space-y-8 lg:space-y-12">
		<div class="text-base text-amber-500 font-semibold"><slot name="header" /></div>

		<div class="flex flex-wrap justify-center text-base lg:text-lg font-mono space-x-2 max-w-full">
			<div class="py-1">{prefix}</div>
			<Typeahead
				bind:searchRef={input}
				bind:value={name}
				search={handleSearch}
				placeholder="name"
				showDropdownOnFocus
				inputAfterSelect="update"
				size="20"
				debounce={300}
			>
				<slot name="no-results" slot="no-results" />
			</Typeahead>
			<Select bind:value={lang} />
			<div class="py-1">{suffix}</div>
		</div>

		<div
			class="w-full flex flex-col sm:items-center justify-center space-y-2 sm:flex-row sm:space-y-0 sm:space-x-4 text-sm text-center"
		>
			<a
				href={page}
				target="_blank"
				class="bg-blue-500 text-white rounded py-2 px-6 select-none w-full sm:w-auto {name
					? ''
					: 'cursor-not-allowed opacity-50'}"><slot name="page" /></a
			>
			<div><slot name="or" /></div>
			<a
				href={full}
				target="_blank"
				class="bg-orange-500 text-white rounded py-2 px-6 select-none w-full sm:w-auto {name
					? ''
					: 'cursor-not-allowed opacity-50'}"><slot name="view" /></a
			>
			<div><slot name="or" /></div>
			<button
				class="bg-orange-100 text-orange-500 rounded py-2 px-6 select-none w-full sm:w-auto {name
					? ''
					: 'cursor-not-allowed opacity-50'}"
				on:click={handleCopy}><slot name="copy" /></button
			>
		</div>
	</div>
</div>
