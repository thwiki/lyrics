<script lang="ts">
	/**
	 * Specify the input value
	 */
	export let value = '';

	/**
	 * Set to `true` to auto focus the input on mount
	 */
	export let autofocus = false;

	/**
	 * Specify the debounce value in milliseconds (ms)
	 */
	export let debounce = 0;

	/**
	 * Specify the input placeholder text
	 */
	export let placeholder = 'Search...';

	/**
	 * Specify an `id` for the `input`
	 */
	export let id = 'search' + Math.random().toString(36);

	/**
	 * Obtain a reference to the `input` element
	 */
	export let ref: HTMLInputElement | null = null;

	/**
	 * Set to `true` to omit the form `role="search"` attribute
	 */
	export let removeFormAriaAttributes = false;

	import { createEventDispatcher, onMount, afterUpdate } from 'svelte';

	const dispatch = createEventDispatcher();

	let prevValue = value;
	let timeout: ReturnType<typeof setTimeout>;
	let calling = false;

	function debounceFn(fn: () => any) {
		if (calling) return;
		calling = true;
		timeout = setTimeout(() => {
			fn();
			calling = false;
		}, debounce);
	}

	onMount(() => {
		if (autofocus) window.requestAnimationFrame(() => ref?.focus());
		return () => clearTimeout(timeout);
	});

	afterUpdate(() => {
		if (value.length > 0 && value !== prevValue) {
			dispatch('key', value);
			if (debounce > 0) {
				debounceFn(() => dispatch('type', value));
			} else {
				dispatch('type', value);
			}
		}

		if (value.length === 0 && prevValue.length > 0) dispatch('clear');

		prevValue = value;
	});
</script>

<form
	data-svelte-search
	role={removeFormAriaAttributes ? null : 'search'}
	aria-labelledby={removeFormAriaAttributes ? null : id}
	on:submit|preventDefault
>
	<input
		bind:this={ref}
		name="search"
		type="search"
		{placeholder}
		autocomplete="off"
		spellcheck="false"
		{...$$restProps}
		{id}
		bind:value
		on:input
		on:change
		on:focus
		on:blur
		on:keydown
	/>
</form>
