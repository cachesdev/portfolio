<script lang="ts">
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { Badge } from '$lib/components/ui/badge/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { showcaseSections, representedFamilies } from './sections.js';
	import type { Snippet } from 'svelte';

	let {
		title = 'Component showcase',
		label = 'showcase',
		description = 'A working index of the local shadcn-svelte component library after the terminal-grid design pass.',
		children
	}: {
		title?: string;
		label?: string;
		description?: string;
		children: Snippet;
	} = $props();

	const currentPath = $derived(page.url.pathname);

	function isActive(href: string) {
		return currentPath === href || (href !== '/showcase' && currentPath.startsWith(`${href}/`));
	}
</script>

<svelte:head>
	<title>{title} | Component showcase</title>
	<meta name="description" content={description} />
</svelte:head>

<main class="min-h-screen bg-background text-foreground">
	<section class="terminal-grid border-b border-border">
		<div
			class="mx-auto grid max-w-7xl gap-8 px-4 py-8 sm:px-6 lg:grid-cols-[17rem_1fr] lg:px-8 lg:py-10"
		>
			<aside class="lg:sticky lg:top-6 lg:self-start">
				<div class="hairline-frame terminal-grid-sm p-4">
					<a
						href={resolve('/showcase')}
						class="group block border-b border-border pb-4 outline-none focus-visible:ring-2 focus-visible:ring-ring/30"
					>
						<p class="terminal-label">// local registry</p>
						<h1 class="mt-2 text-2xl font-semibold tracking-tight text-foreground">Showcase</h1>
						<p class="mt-2 text-sm leading-6 text-muted-foreground">
							{representedFamilies.length} component families represented. TanStack data-table intentionally
							skipped.
						</p>
					</a>

					<nav class="mt-4 grid gap-1" aria-label="Showcase sections">
						{#each showcaseSections as section (section.href)}
							<a
								href={resolve(section.href)}
								aria-current={isActive(section.href) ? 'page' : undefined}
								class="group grid gap-1 border border-transparent px-3 py-2 text-sm text-muted-foreground outline-none transition-colors hover:border-border hover:bg-muted/70 hover:text-foreground focus-visible:ring-2 focus-visible:ring-ring/30 aria-[current=page]:border-primary/40 aria-[current=page]:bg-primary/15 aria-[current=page]:text-foreground"
							>
								<span class="flex items-center justify-between gap-3">
									<span class="font-medium">{section.title}</span>
									<span
										class="font-mono text-[0.625rem] uppercase tracking-widest text-muted-foreground"
									>
										{section.label}
									</span>
								</span>
								<span class="text-xs text-muted-foreground">{section.count}</span>
							</a>
						{/each}
					</nav>

					<Separator class="my-4" />

					<div class="flex flex-wrap gap-2">
						<Badge variant="secondary">Svelte 5</Badge>
						<Badge variant="outline">Bits UI</Badge>
						<Badge variant="secondary">Tailwind v4</Badge>
					</div>
				</div>
			</aside>

			<div class="min-w-0">
				<header class="hairline-frame terminal-grid-sm mb-6 p-6 md:p-8">
					<div class="flex flex-wrap items-start justify-between gap-4">
						<div class="max-w-3xl">
							<p class="terminal-label">// {label}</p>
							<h2 class="mt-3 text-3xl font-semibold tracking-tight text-foreground md:text-5xl">
								{title}
							</h2>
							<p class="mt-4 max-w-2xl text-base leading-7 text-muted-foreground">{description}</p>
						</div>
						<Button href={resolve('/showcase')} variant="outline" size="sm">Index</Button>
					</div>
				</header>

				{@render children()}
			</div>
		</div>
	</section>
</main>
