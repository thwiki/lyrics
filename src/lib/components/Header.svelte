<script lang="ts">
	import Logo from '../../assets/logo.png';
	import { locale, locales, t } from '../translations';
	import GithubIcon from '../../assets/icons/github.svg?component';

	function languageDisplayName(lang: string) {
		return new Intl.DisplayNames([lang], { type: 'language' }).of(lang);
	}
</script>

<header class="relative z-30">
	<nav class="border-b border-gray-300 shadow h-full min-h-[60px]">
		<div class="flex items-center h-full container !px-3">
			<div class="relative flex items-center justify-between">
				<a href="/" class="my-3 mr-6"
					><img
						width="36"
						height="36"
						title={$t('common.name')}
						src={Logo}
						alt={$t('common.logo')}
					/>
				</a>
			</div>
			<div class="ml-auto pl-3">
				<ul
					class="flex items-center space-x-4 sm:space-x-8 text-sm [&>li]:block [&>li]:relative [&>li]:transition-all"
				>
					<li>
						<a
							class="flex items-center justify-center h-10 font-semibold text-gray-700 hover:text-black"
							target="_blank"
							href="https://thwiki.cc"
							title={$t('common.author')}
							><span class="sm:hidden">THB</span><span class="hidden sm:block"
								>{$t('common.author')}</span
							></a
						>
					</li>
					<li>
						<a
							class="flex items-center justify-center h-10 font-semibold text-gray-700 hover:text-black"
							target="_blank"
							href="https://github.com/thwiki/lyrics"
							title={$t('common.github')}
							><GithubIcon class="h-6 sm:hidden" /><span class="hidden sm:block"
								>{$t('common.github')}</span
							></a
						>
					</li>
					<li>
						<select
							class="h-10 pb-0.5 font-semibold text-gray-700 hover:text-black"
							bind:value={$locale}
						>
							{#each $locales as lang}
								<option value={lang}>{languageDisplayName(lang)}</option>
							{/each}
						</select>
					</li>
				</ul>
			</div>
		</div>
	</nav>
</header>
