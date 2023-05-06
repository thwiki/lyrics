<script lang="ts">
	export let href: string;
	export let name: string | undefined = undefined;
	export let target: string | undefined = undefined;
</script>

<a
	class="text-blue-500 hover:text-blue-900 transition-colors"
	{href}
	target={target ?? (href.startsWith('/') ? undefined : '_blank')}
	{...$$restProps}
	>{#if name == null}<slot />{:else}{name}{/if}</a
>
