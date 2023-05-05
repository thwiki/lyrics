<script lang="ts">
	import Link from '../../lib/components/Link.svelte';
	import { t } from '../../lib/translations';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<svelte:head>
	<title>{data.lyrics.title} - {$t('common.name')}</title>
	<meta property="og:title" itemprop="name" content={data.lyrics.title} />
	<meta property="og:type" content="article" />
	{#if data.lyrics.cover}
		<meta property="og:image" itemprop="image" content={data.lyrics.cover} />
	{/if}
	<meta property="og:site_name" content={$t('common.name')} />
</svelte:head>

<div class="container">
	<div
		class="grid gap-3 md:grid-cols-[320px_1fr] md:grid-rows-[repeat(3,auto)] items-start justify-stretch text-sm"
	>
		<div class="shadow-md rounded-xl">
			<div class="bg-pink-300 rounded-t-lg py-1 pl-5 pr-3">{$t('lyrics.basic')}</div>
			{#if data.lyrics.cover}
				<div class="w-full bg-slate-50">
					<img
						class="w-auto max-h-[200px] md:max-h-[400px] mx-auto"
						src={data.lyrics.cover}
						alt={data.lyrics.album}
					/>
				</div>
			{/if}
			<div
				class="grid grid-cols-[auto_1fr] place-items-start gap-y-1 gap-x-2 bg-slate-100 rounded-b-lg p-3"
			>
				<div class="bg-amber-300 rounded-lg px-2">{$t('lyrics.title')}</div>
				<div class="break-all">{data.lyrics.title}</div>
				<div class="bg-amber-300 rounded-lg px-2">{$t('lyrics.album')}</div>
				<div class="break-all">{data.lyrics.album}</div>
				<div class="bg-amber-300 rounded-lg px-2">{$t('lyrics.artist')}</div>
				<div class="break-all">{data.lyrics.artist}</div>
				<div class="bg-amber-300 rounded-lg px-2">{$t('lyrics.editor')}</div>
				<div class="break-all">
					{#each data.lyrics.editors as editor, i}{i === 0 ? ' ' : ', '}<Link {...editor} />{/each}
				</div>
				{#if data.lyrics.author}
					<div class="bg-amber-300 rounded-lg px-2">{$t('lyrics.author')}</div>
					<div class="break-all">{data.lyrics.author}</div>
				{/if}
			</div>
		</div>
		<div class="shadow-md rounded-xl md:col-start-2 md:col-span-2 md:row-span-full">
			<div
				class="grid grid-cols-[auto_1fr] items-start justify-stretch gap-y-1 gap-x-2 rounded-b-xl p-3 font-mono overflow-auto"
			>
				{#each data.lyrics.scripts as script}
					<div class="bg-slate-300 text-slate-700 rounded-lg px-2 text-center">
						{script.time || '-'}
					</div>
					<div>{script.text}</div>
				{/each}
			</div>
		</div>
		{#if data.lyrics.relations.length > 0}
			<div class="shadow-md rounded-xl md:col-start-1 md:row-start-2">
				<div class="bg-blue-300 rounded-t-lg py-1 pl-5 pr-3">{$t('lyrics.other')}</div>
				<div class="bg-slate-100 rounded-lg py-3 pl-5 pr-3">
					<ul class="space-y-1">
						{#each data.lyrics.relations as relation, i}
							<li>{i + 1}. <Link {...relation} /></li>
						{/each}
					</ul>
				</div>
			</div>
		{:else}
			<div class="md:col-start-1 md:row-start-2" />
		{/if}
	</div>
</div>
