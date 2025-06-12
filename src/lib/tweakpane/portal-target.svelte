<script lang="ts" module>
	import type { Snippet } from 'svelte';
	import { SvelteSet } from 'svelte/reactivity';
	let snippets = new SvelteSet<Snippet>();

	export const addSnippet = (snippet: Snippet) => {
		snippets.add(snippet);
	};

	export const removeSnippet = (snippet: Snippet) => {
		snippets.delete(snippet);
	};
</script>

{#each snippets as snippet}
	{@render snippet()}
{/each}
