<script lang="ts">
	let className: string | undefined = undefined;
	export { className as class };
	export let reverse = false;
	export let windowed = false;
</script>

<section
	class="relative flex {reverse
		? 'flex-col-reverse md:flex-row-reverse space-x-reverse md:space-x-reverse'
		: 'flex-col md:flex-row'} justify-between items-center py-6 md:py-10 lg:py-12 space-y-6 md:space-y-0 md:space-x-6 {className ??
		''}"
>
	<div class="basis-0 grow max-w-lg flex flex-col items-start">
		<h3 class="text-3xl font-semibold text-slate-800 tracking-wide">
			<slot name="heading" />
		</h3>
		<div class="mt-4 text-base text-slate-600">
			<slot name="content" />
		</div>
		<div class="mt-6 md:mt-10">
			<slot name="action" />
		</div>
	</div>
	{#if windowed}
		<div class="basis-0 grow-[2] max-w-2xl flex flex-col bg-white shadow-md rounded-lg">
			<div class="h-4 md:h-6 {reverse ? 'bg-pink-300' : 'bg-amber-300'} rounded-t-lg" />
			<div class="overflow-hidden grow rounded-b-lg">
				<slot name="display" />
			</div>
		</div>
	{:else}
		<div class="basis-0 grow-[2] max-w-2xl">
			<div class="overflow-hidden grow rounded-b-lg">
				<slot name="display" />
			</div>
		</div>
	{/if}
</section>
