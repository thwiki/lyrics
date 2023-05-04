<script lang="ts">
	import { page } from '$app/stores';
</script>

<div class="container">
	<div class="text-xl">{$page.error?.message ?? ''}</div>
</div>
