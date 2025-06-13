<script lang="ts">
	import { getContext } from 'svelte';
	import type { ApplicationState } from '../state.svelte';
	import Button from '$lib/button.svelte';
</script>

<div class="pointer-events-none fixed top-0 left-0 z-10 h-dvh w-dvw">
	<div class="relative h-full w-full">
		<div
			class="bg-background-hard/90 border-gruvbox-dark3 pointer-events-auto absolute top-16 left-100 h-80 w-96 rounded-lg border p-6 shadow-2xl backdrop-blur-sm"
		>
			<div class="mb-4 flex items-center justify-between">
				<h2 class="font-michroma text-gruvbox-light0 text-xl">Blog</h2>
				<Button class="text-gruvbox-light4 hover:text-gruvbox-light0 text-xl leading-none" href="/"
					>x</Button
				>
			</div>
			<div class="font-fira text-gruvbox-light1 space-y-4">
				<p>Empty! I haven't post any blogs yet.</p>
				<div class="space-y-2">
					<div class="bg-gruvbox-dark1 border-gruvbox-blue rounded border-l-4 p-3">
						<h3 class="text-gruvbox-light0 font-semibold">Latest Post</h3>
						<p class="text-gruvbox-light2 text-sm">No post!</p>
					</div>
					<div class="bg-gruvbox-dark1 border-gruvbox-green rounded border-l-4 p-3">
						<h3 class="text-gruvbox-light0 font-semibold">Featured</h3>
						<p class="text-gruvbox-light2 text-sm">No post!</p>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
