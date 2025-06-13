<script>
	import Button from '$lib/button.svelte';
</script>

<div class="pointer-events-none fixed top-0 left-0 z-10 h-dvh w-dvw">
	<div class="relative h-full w-full">
		<div
			class="bg-background-hard/90 border-gruvbox-dark3 pointer-events-auto absolute top-20 left-100 h-96 w-[500px] rounded-lg border p-6 shadow-2xl backdrop-blur-sm"
		>
			<div class="mb-4 flex items-center justify-between">
				<h2 class="font-michroma text-gruvbox-light0 text-xl">Projects</h2>
				<Button href="/" class="text-gruvbox-light4 hover:text-gruvbox-light0 text-xl leading-none"
					>x</Button
				>
			</div>
			<div class="font-fira text-gruvbox-light1 space-y-4">
				<p>Check out some of my recent work and side projects.</p>
				<div class="grid grid-cols-1 gap-4">
					<div class="bg-gruvbox-dark1 border-gruvbox-dark2 rounded-lg border p-4">
						<h3 class="text-gruvbox-light0 mb-2 font-semibold">Portfolio Website</h3>
						<p class="text-gruvbox-light2 mb-2 text-sm">
							Interactive 3D portfolio built with Threlte and SvelteKit
						</p>
						<div class="flex gap-2">
							<span class="bg-gruvbox-blue text-gruvbox-light0 rounded px-2 py-1 text-xs"
								>Svelte</span
							>
							<span class="bg-gruvbox-green text-gruvbox-light0 rounded px-2 py-1 text-xs"
								>Three.js</span
							>
						</div>
					</div>
					<div class="bg-gruvbox-dark1 border-gruvbox-dark2 rounded-lg border p-4">
						<a
							href="https://github.com/cachesdev/souuup"
							class="text-gruvbox-light0 mb-2 font-semibold">Validation Library (Souuup)</a
						>
						<p class="text-gruvbox-light2 mb-2 text-sm">Type-Safe validation library.</p>
						<div class="flex gap-2">
							<span class="bg-gruvbox-orange text-gruvbox-light0 rounded px-2 py-1 text-xs">Go</span
							>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
