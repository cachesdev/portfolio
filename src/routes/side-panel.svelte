<script lang="ts">
	import { getContext } from 'svelte';
	import type { ApplicationState } from './state.svelte';
	import Button from '$lib/button.svelte';

	let application = getContext<ApplicationState>('application');
</script>

<div class="bg-background-hard font-michroma relative z-20 flex h-dvh flex-col gap-2 px-2">
	<a href="/" class="hover:text-gruvbox-light0 text-5xl">
		Gustavo<br /><span class="underline decoration-5"> Domínguez</span>
	</a>
	<div class="font-fira">Software Developer focused on backend technologies.</div>
	<div class="flex flex-col">
		{#each application.sidebar.items as item}
			<Button
				href={item.href}
				class="hover:bg-background w-full p-2 pb-3 text-left {item.active ? 'bg-background' : ''}"
				{...item.props}>{item.title}</Button
			>
		{/each}
	</div>
	<div class="font-fira mt-auto flex flex-col gap-1 pb-2">
		<span>Socials: </span>
		<div class="flex">
			<a class="hover:text-gruvbox-light0" href="https://github.com/cachesdev">Github</a>
		</div>
	</div>
</div>
