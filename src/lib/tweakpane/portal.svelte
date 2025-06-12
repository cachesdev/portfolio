<script lang="ts">
	import { onDestroy, type Snippet } from 'svelte';
	import { addSnippet, removeSnippet } from './portal-target.svelte';

	let { children }: { children: Snippet } = $props();

	addSnippet(children);

	onDestroy(() => {
		removeSnippet(children);
	});
</script>
