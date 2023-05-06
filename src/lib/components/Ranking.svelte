<script lang="ts">
	import Link from './Link.svelte';

	export let items: {
		title: string;
		value: string | number;
	}[];
</script>

<div
	class="text-sm flex-auto w-full sm:max-w-md lg:min-w-[384px] shadow-md rounded-lg overflow-hidden"
>
	<div class="font-bold py-3 pl-2 pr-5 bg-amber-200 text-base">
		<slot name="header" />
	</div>
	<div class="divide-y bg-white">
		{#each items as item, index}
			<div class="py-3 pl-3 pr-5 flex items-center justify-start">
				<div
					class="mr-3 w-full max-w-[24px] leading-6 bg-pink-200 text-center rounded font-semibold select-none"
				>
					{index + 1}
				</div>
				<div class="mr-3 whitespace-nowrap overflow-hidden text-ellipsis">
					<Link href="/{item.title}" target="_blank">{item.title}</Link>
					<Link href="https://lyrics.thwiki.cc/{item.title}.lrc">.lrc</Link>
				</div>
				<div class="ml-auto text-amber-500 whitespace-nowrap">
					<slot name="item" title={item.title} value={item.value} />
				</div>
			</div>
		{/each}
	</div>
</div>
