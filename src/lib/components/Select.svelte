<script lang="ts">
	export let value = '';
</script>

<select
	class="border-b-2 py-1 transition-colors border-gray-300 focus:border-pink-300 focus:outline-none {value ===
	''
		? 'select-none text-gray-400'
		: ''}"
	bind:value
	placeholder="Lang"
>
	<option value="" />
	<option value=".all">.all</option>
	<option value=".ja">.ja</option>
	<option value=".zh">.zh</option>
</select>
