<script lang="ts">
	import '../app.postcss';
	import { t } from '../lib/translations';

	import Header from '../lib/components/Header.svelte';
	import Footer from '../lib/components/Footer.svelte';
</script>

<svelte:head>
	<title>{$t('common.name')}</title>
</svelte:head>

<div class="flex flex-col h-[100vh]">
	<Header />
	<div class="relative py-6">
		<slot />
	</div>
	<Footer />
</div>
