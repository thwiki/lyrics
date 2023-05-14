<script lang="ts">
	import { t } from '../lib/translations';
	import Ranking from '../lib/components/Ranking.svelte';
	import RelativeTime from '../lib/components/RelativeTime.svelte';
	import Searcher from '../lib/components/Searcher.svelte';
	import type { PageData } from './$types';
	import Section from '../lib/components/Section.svelte';
	import UiSearch from '../assets/images/ui-search.png';
	import UiExport from '../assets/images/ui-export.png';
	import RecentIcon from '../assets/icons/recent.svg?component';
	import PopularIcon from '../assets/icons/popular.svg?component';

	export let data: PageData;

	let search: HTMLInputElement;
</script>

<div
	class="absolute top-0 left-0 right-0 w-full h-[760px] -z-10 overflow-hidden bg-gradient-to-b from-amber-500 via-pink-200 py-6 pointer-events-none"
/>
<div>
	<div class="container">
		<div class="flex flex-col items-center space-y-4 mt-12">
			<h1 class="text-white font-bold text-4xl">
				{$t('common.name')}
			</h1>
			<h2 class="text-white text-2xl">{$t('home.main.head')}</h2>
			<div class="text-amber-700 font-semibold text-center max-w-2xl">
				{$t('home.main.desc')}
			</div>
		</div>
	</div>

	<div class="container">
		<div id="search" class="mt-10 mb-10 p-12 shadow-md rounded-lg bg-white">
			<Searcher
				prefix="lyrics.thwiki.cc/"
				suffix=".lrc"
				placeholder={$t('home.searcher.placeholder')}
				bind:input={search}
			>
				<svelte:fragment slot="header">{$t('home.searcher.action')}</svelte:fragment>
				<svelte:fragment slot="or">{$t('home.searcher.or')}</svelte:fragment>
				<svelte:fragment slot="page">{$t('home.searcher.page')}</svelte:fragment>
				<svelte:fragment slot="view">{$t('home.searcher.view')}</svelte:fragment>
				<svelte:fragment slot="copy">{$t('home.searcher.copy')}</svelte:fragment>
				<svelte:fragment slot="no-results">{$t('home.searcher.empty')}</svelte:fragment>
			</Searcher>
		</div>
	</div>

	<div class="container">
		<div
			class="flex space-y-6 sm:space-y-0 space-x-4 lg:space-x-0 justify-center sm:justify-around items-center sm:items-start flex-col sm:flex-row py-10 mb-6"
		>
			<Ranking items={data.latest}>
				<div slot="header" class="flex items-center">
					<RecentIcon class="h-8 -my-2 mr-2 fill-green-500" />{$t('home.ranking.recent')}
				</div>
				<RelativeTime slot="item" let:value time={value} sync />
			</Ranking>
			<Ranking items={data.popular}>
				<div slot="header" class="flex items-center">
					<PopularIcon class="h-8 -my-2 mr-2 fill-amber-500" />{$t('home.ranking.popular')}
				</div>
				<span slot="item" let:value>{value}</span>
			</Ranking>
		</div>
	</div>

	<Section windowed class="container">
		<svelte:fragment slot="heading">{$t('home.search.head')}</svelte:fragment>
		<svelte:fragment slot="content">
			{$t('home.search.desc')}
		</svelte:fragment>
		<button
			slot="action"
			class="text-base font-bold text-amber-500 hover:text-amber-700 transition-colors"
			on:click={() => {
				search?.focus();
			}}>{$t('home.search.action')}</button
		>
		<img
			slot="display"
			src={UiSearch}
			alt={$t('home.search.display')}
			class="min-h-[240px] md:min-h-[360px] h-full w-full object-cover object-center"
		/>
	</Section>

	<div class="relative overflow-hidden md:-mt-20 md:pt-20">
		<div class="absolute inset-0 -z-10 pointer-events-none">
			<div
				class="absolute -top-16 left-0 md:top-1/4 md:left-0 md:-mt-[7%] w-[200%] h-[200%] bg-slate-100 -rotate-6 -skew-x-3"
			>
				<img
					src={UiExport}
					alt={$t('home.export.display')}
					class="absolute left-1/3 md:left-[20%] max-w-none -translate-x-1/2 w-[640px] h-[320px] md:w-[1280px] md:h-[640px] md:blur-[1px]"
					width="1280"
					height="640"
				/>
			</div>
		</div>

		<Section reverse class="container">
			<svelte:fragment slot="heading">{$t('home.export.head')}</svelte:fragment>
			<svelte:fragment slot="content">
				{$t('home.export.desc')}
			</svelte:fragment>
			<blockquote slot="action" class="text-sm bg-pink-200 px-6 py-4 rounded-lg -mx-4">
				{$t('home.export.action')}
			</blockquote>
			<div
				slot="display"
				class="min-h-[240px] md:min-h-[360px] h-full w-full object-cover object-center"
			/>
		</Section>
	</div>

	<div class="bg-amber-100">
		<Section class="container">
			<svelte:fragment slot="heading">{$t('home.community.head')}</svelte:fragment>
			<svelte:fragment slot="content">
				{$t('home.community.desc')}
			</svelte:fragment>
			<a
				slot="action"
				class="text-base font-bold text-amber-500 hover:text-amber-700 transition-colors"
				href="https://thwiki.cc/%E5%B8%AE%E5%8A%A9:%E6%AD%8C%E8%AF%8D"
				target="_blank">{$t('home.community.action')}</a
			>
			<div slot="display" class="min-h-[240px] md:min-h-[360px] h-full w-full flex items-center">
				<div class="flex flex-wrap justify-center pl-14 pr-8 py-4">
					{#each data.members as member}
						<img
							class="rounded-full border-2 border-white -ml-6 hover:scale-110 transition-transform"
							src="https://avatar.thwiki.cc/thwikicc_wiki_{member}_l.png"
							crossorigin="anonymous"
							alt={$t('home.community.display')}
						/>
					{/each}
				</div>
			</div>
		</Section>
	</div>

	<div class="container">
		<div class="mt-12 md:mt-20 lg:mt-30">
			<ul class="text-sm text-slate-700">
				<li>{$t('home.footnote.translation')}</li>
			</ul>
		</div>
	</div>
</div>
