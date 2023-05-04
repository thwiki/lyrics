<script lang="ts">
	import Search from './Search.svelte';
	import { tick, createEventDispatcher, afterUpdate } from 'svelte';

	interface Result {
		value: string;
		label: string;
		disabled?: boolean;
	}

	export let id = 'typeahead-' + Math.random().toString(36);
	export let value = '';

	/** Set to `false` to prevent the first result from being selected */
	export let autoselect = true;

	/**
	 * Set to `keep` to keep the search field unchanged after select, set to `clear` to auto-clear search field
	 */
	export let inputAfterSelect: 'update' | 'clear' | 'keep' = 'update';

	export let results: Result[] = [];

	/** Set to `true` to re-focus the input after selecting a result */
	export let focusAfterSelect = false;

	/** Set to `true` to only show results when the input is focused */
	export let showDropdownOnFocus = false;

	export let search: (term: string) => Promise<Result[]>;

	const dispatch = createEventDispatcher();

	export let comboboxRef: HTMLElement | null = null;
	export let searchRef: HTMLInputElement | null = null;
	let hideDropdown = false;
	let selectedIndex = -1;
	let prevResults = '';
	let isFocused = false;

	afterUpdate(() => {
		if (prevResults !== resultsId && autoselect) {
			selectedIndex = getNextNonDisabledIndex();
		}

		if (prevResults !== resultsId && !$$slots['no-results']) {
			hideDropdown = results.length === 0;
		}

		prevResults = resultsId;
	});

	async function select() {
		const result = results[selectedIndex];

		if (result.disabled) return;

		const selectedValue = result.value;
		const searchedValue = value;

		if (inputAfterSelect == 'clear') value = '';
		if (inputAfterSelect == 'update') value = selectedValue;

		dispatch('select', {
			selectedIndex,
			searched: searchedValue,
			selected: selectedValue,
			value: result.value
		});

		await tick();

		if (focusAfterSelect) searchRef?.focus();
		close();
	}

	function getNextNonDisabledIndex(): number {
		let index = 0;
		let disabled = results[index]?.disabled ?? false;

		while (disabled) {
			if (index === results.length) {
				index = 0;
			} else {
				index += 1;
			}

			disabled = results[index]?.disabled ?? false;
		}

		return index;
	}

	function change(direction: -1 | 1): void {
		let index =
			direction === 1 && selectedIndex === results.length - 1 ? 0 : selectedIndex + direction;
		if (index < 0) index = results.length - 1;

		let disabled = results[index].disabled;

		while (disabled) {
			if (index === results.length) {
				index = 0;
			} else {
				index += direction;
			}

			disabled = results[index].disabled;
		}

		selectedIndex = index;
	}

	let loading = false;

	async function getResults() {
		const searchValue = value;
		const searchResults = await search(searchValue);
		if (searchValue === value) {
			loading = false;
			results = searchResults;
		}
	}

	const open = () => {
		hideDropdown = false;
		if (showDropdownOnFocus) {
			showResults = true;
			isFocused = true;
		}
	};
	const close = () => {
		hideDropdown = true;
		isFocused = false;
	};

	$: resultsId = results.map((result) => result.value).join('');
	$: showResults = !hideDropdown && results.length > 0;
	$: if (showDropdownOnFocus) {
		showResults = showResults && isFocused;
	}
</script>

<svelte:window
	on:mousedown={({ target }) => {
		if (!hideDropdown && target instanceof Node && !comboboxRef?.contains(target)) {
			close();
		}
	}}
/>

<div
	class="relative bg-white"
	bind:this={comboboxRef}
	role="combobox"
	aria-haspopup="listbox"
	aria-controls="{id}-listbox"
	aria-expanded={showResults}
	id="{id}-typeahead"
>
	<Search
		{id}
		removeFormAriaAttributes={true}
		{...$$restProps}
		bind:ref={searchRef}
		role="combobox"
		aria-autocomplete="list"
		aria-controls="{id}-listbox"
		aria-labelledby="{id}-label"
		aria-activedescendant={selectedIndex >= 0 && !hideDropdown && results.length > 0
			? `${id}-result-${selectedIndex}`
			: null}
		bind:value
		on:key
		on:key={() => {
			loading = true;
		}}
		on:type
		on:type={() => {
			open();
			getResults();
		}}
		on:input
		on:change
		on:focus
		on:focus={open}
		on:clear
		on:clear={() => {
			open();
			getResults();
		}}
		on:blur
		on:keydown
		on:keydown={(e) => {
			if (results.length === 0) return;

			switch (e.key) {
				case 'Enter':
					select();
					break;
				case 'ArrowDown':
					e.preventDefault();
					change(1);
					break;
				case 'ArrowUp':
					e.preventDefault();
					change(-1);
					break;
				case 'Escape':
					e.preventDefault();
					value = '';
					searchRef?.focus();
					close();
					break;
			}
		}}
	/>
	<ul
		class="absolute top-full left-0 min-w-max shadow-sm text-sm bg-white border-gray-300 divide-y"
		class:svelte-typeahead-list={true}
		class:z-10={showResults}
		role="listbox"
		aria-labelledby="{id}-label"
		id="{id}-listbox"
	>
		{#if showResults}
			{#each results as result, index}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<li
					id="{id}-result-{index}"
					class="py-2 px-4 cursor-pointer select-none"
					class:bg-amber-500={selectedIndex === index}
					class:!cursor-not-allowed={result.disabled}
					class:hover:bg-amber-500={!result.disabled}
					class:opacity-40={result.disabled}
					role="option"
					aria-selected={selectedIndex === index}
					on:click={() => {
						if (result.disabled) return;
						selectedIndex = index;
						select();
					}}
					on:mouseenter={() => {
						if (result.disabled) return;
						selectedIndex = index;
					}}
				>
					<slot {result} {index} {value}>
						{result.label}
					</slot>
				</li>
			{/each}
		{/if}
		{#if $$slots['no-results'] && !loading && !hideDropdown && value.length > 0 && results.length === 0}
			<div class="py-2 px-4">
				<slot name="no-results" {value} />
			</div>
		{/if}
	</ul>
</div>

<style lang="postcss">
	:global([data-svelte-search] input) {
		@apply border-b-2 py-1 transition-colors border-gray-300 focus:border-pink-300 focus:outline-none;
	}
</style>
