<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { locale } from '../translations';
	import { browser } from '$app/environment';
	export let time: Date | string | number;
	export let format: 'long' | 'short' | 'narrow' = 'long';
	export let numeric: 'always' | 'auto' = 'auto';
	export let sync: boolean = false;

	let relativeTime = '';
	let isoTime = '';
	let updateTimeout: number;
	let update = 0;

	onMount(() => {
		update++;
	});

	onDestroy(() => {
		clearTimeout(updateTimeout);
	});

	interface UnitConfig {
		max: number;
		value: number;
		unit: Intl.RelativeTimeFormatUnit;
	}

	const availableUnits: UnitConfig[] = [
		{ max: 2760000, value: 60000, unit: 'minute' }, // max 46 minutes
		{ max: 72000000, value: 3600000, unit: 'hour' }, // max 20 hours
		{ max: 518400000, value: 86400000, unit: 'day' }, // max 6 days
		{ max: 2419200000, value: 604800000, unit: 'week' }, // max 28 days
		{ max: 28512000000, value: 2592000000, unit: 'month' }, // max 11 months
		{ max: Infinity, value: 31536000000, unit: 'year' }
	];

	function getTimeUntilNextUnit(unit: 'second' | 'minute' | 'hour' | 'day') {
		const units = { second: 1000, minute: 60000, hour: 3600000, day: 86400000 };
		const value = units[unit];
		return value - (Date.now() % value);
	}

	function clearSync() {
		clearTimeout(updateTimeout);
	}

	function setSync(unit: Intl.RelativeTimeFormatUnit) {
		let nextInterval: number;

		// NOTE: this could be optimized to determine when the next update should actually occur, but the size and cost of
		// that logic probably isn't worth the performance benefit
		if (unit === 'minute') {
			nextInterval = getTimeUntilNextUnit('second');
		} else if (unit === 'hour') {
			nextInterval = getTimeUntilNextUnit('minute');
		} else if (unit === 'day') {
			nextInterval = getTimeUntilNextUnit('hour');
		} else {
			// Cap updates at once per day. It's unlikely a user will reach this value, plus setTimeout has a limit on the
			// value it can accept. https://stackoverflow.com/a/3468650/567486
			nextInterval = getTimeUntilNextUnit('day'); // next day
		}

		updateTimeout = window.setTimeout(() => update++, nextInterval);
	}

	$: {
		update;
		const now = new Date();
		const then = new Date(time);

		// Check for an invalid date
		if (isNaN(then.getMilliseconds())) {
			relativeTime = '';
			isoTime = '';
		} else {
			const diff = then.getTime() - now.getTime();
			const { unit, value } = availableUnits.find((singleUnit) => Math.abs(diff) < singleUnit.max)!;

			isoTime = then.toISOString();
			relativeTime = new Intl.RelativeTimeFormat($locale, {
				numeric: numeric ?? 'auto',
				style: format ?? 'long'
			}).format(Math.round(diff / value), unit);

			// If sync is enabled, update as time passes
			if (browser) {
				clearSync();

				if (sync) {
					setSync(unit);
				}
			}
		}
	}
</script>

<time datetime={isoTime}>{relativeTime}</time>
