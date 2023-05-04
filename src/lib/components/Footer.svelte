<script lang="ts">
	import { t } from '../translations';
	import RelativeTime from './RelativeTime.svelte';

	export let date: Date | undefined = undefined;
</script>

<footer class="relative mt-auto pt-10 md:pt-3">
	<div class="bg-slate-800 w-full py-10 text-white">
		<div class="h-full container !px-3">
			<div class="flex flex-wrap items-start">
				<div class="mr-10 text-sm">
					<a rel="license" href="https://creativecommons.org/licenses/by-nc-sa/3.0/"
						><img
							class="inline"
							alt="Creative Commons License"
							src="https://i.creativecommons.org/l/by-nc-sa/3.0/80x15.png"
							width="80"
							height="15"
						/>
						{$t('common.license')}</a
					>
				</div>
				<div class="ml-auto text-center font-semibold">
					<a href="https://thwiki.cc/" target="_blank"
						><img
							src="https://thwiki.cc/banner/banner_b1.png"
							width="200"
							height="40"
							alt={$t('common.author_desc')}
						/></a
					>
					<div>
						Product of
						<a
							class="text-amber-500 hover:text-pink-300 transition-colors"
							href="https://thwiki.cc/THBWiki:%E5%85%B3%E4%BA%8E"
							target="_blank">{$t('common.author')}</a
						>
					</div>
				</div>
			</div>
			<hr class="my-4 border-slate-500" />
			<div class="flex space-x-3 flex-wrap text-sm">
				<a
					class="block text-blue-500 hover:text-blue-300 transition-colors whitespace-nowrap"
					href="https://thwiki.cc/THBWiki:%E9%9A%90%E7%A7%81%E6%9D%83%E6%94%BF%E7%AD%96"
					target="_blank"
					rel="noopener">{$t('common.privacy_policy')}</a
				>
				<a
					class="block text-blue-500 hover:text-blue-300 transition-colors whitespace-nowrap"
					href="https://thwiki.cc/THBWiki:%E5%85%B3%E4%BA%8E"
					target="_blank"
					rel="noopener">{$t('common.about_author')}</a
				>
				<a
					class="block text-blue-500 hover:text-blue-300 transition-colors whitespace-nowrap"
					href="https://thwiki.cc/THBWiki:%E5%85%8D%E8%B4%A3%E5%A3%B0%E6%98%8E"
					target="_blank"
					rel="noopener">{$t('common.disclaimer')}</a
				>
				<a
					class="block text-blue-500 hover:text-blue-300 transition-colors whitespace-nowrap"
					href="https://thwiki.cc/THBWiki:%E6%8D%90%E6%AC%BE"
					target="_blank"
					rel="noopener">{$t('common.donation')}</a
				>
				{#if date}
					<div class="!ml-auto"><RelativeTime time={date} sync /></div>
				{/if}
			</div>
		</div>
	</div>
</footer>
