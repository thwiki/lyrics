<script lang="ts">
	import { locale } from '../translations';
	export let time: Date | string | number;
	export let format: 'long' | 'short' | 'narrow' = 'long';
	export let numeric: 'always' | 'auto' = 'auto';

	let relativeTime = '';
	let isoTime = '';

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

	function formatRelativeTime(
		value: number,
		unit: Intl.RelativeTimeFormatUnit,
		options?: Intl.RelativeTimeFormatOptions
	): string {
		return new Intl.RelativeTimeFormat($locale, options).format(value, unit);
	}

	$: {
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
		}
	}
</script>

<time datetime={isoTime}>{relativeTime}</time>
