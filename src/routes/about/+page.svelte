<script>
	import Button from '$lib/button.svelte';
</script>

<div class="pointer-events-none fixed top-0 left-0 z-10 h-dvh w-dvw">
	<div class="relative h-full w-full">
		<div
			class="bg-background-hard/90 border-gruvbox-dark3 pointer-events-auto absolute top-32 left-100 h-72 w-[450px] rounded-lg border p-6 shadow-2xl backdrop-blur-sm"
		>
			<div class="mb-4 flex items-center justify-between">
				<h2 class="font-michroma text-gruvbox-light0 text-xl">About Me</h2>
				<Button href="/" class="text-gruvbox-light4 hover:text-gruvbox-light0 text-xl leading-none"
					>x</Button
				>
			</div>
			<div class="font-fira text-gruvbox-light1 space-y-4">
				<p>
					I'm a software developer passionate about programming and creating efficient, scalable
					solutions.
				</p>
				<div class="space-y-3">
					<div>
						<h3 class="text-gruvbox-light0 mb-1 font-semibold">Technologies I work with:</h3>
						<div class="flex flex-wrap gap-2">
							<span class="bg-gruvbox-orange text-gruvbox-light0 rounded px-2 py-1 text-xs">Go</span
							>
							<span class="bg-gruvbox-blue text-gruvbox-light0 rounded px-2 py-1 text-xs"
								>TypeScript</span
							>
							<span class="bg-gruvbox-purple text-gruvbox-light0 rounded px-2 py-1 text-xs"
								>PostgreSQL</span
							>
							<span class="bg-gruvbox-aqua text-gruvbox-dark0 rounded px-2 py-1 text-xs"
								>Docker</span
							>
							<span class="bg-gruvbox-green text-gruvbox-dark0 rounded px-2 py-1 text-xs"
								>Node.js</span
							>
						</div>
					</div>
					<p class="text-sm">
						When I'm not working, you'll find me exploring new technologies or contributing to open
						source projects.
					</p>
				</div>
			</div>
		</div>
	</div>
</div>
