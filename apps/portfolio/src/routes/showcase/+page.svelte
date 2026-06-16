<script lang="ts">
	import { resolve } from '$app/paths';
	import ShowcaseShell from '$lib/showcase/ShowcaseShell.svelte';
	import { representedFamilies, showcaseSections } from '$lib/showcase/sections.js';
	import { Badge } from '$lib/components/ui/badge/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
</script>

<ShowcaseShell
	title="Component showcase"
	label="registry"
	description="A set of working views for the local shadcn-svelte component library, grouped so the volume stays navigable."
>
	<div class="grid gap-6">
		<section class="grid gap-4 md:grid-cols-2">
			{#each showcaseSections as section (section.href)}
				<Card.Root class="terminal-grid-sm">
					<Card.Header>
						<div class="flex items-center justify-between gap-3">
							<Badge variant="outline">{section.label}</Badge>
							<span class="font-mono text-xs text-muted-foreground">{section.count}</span>
						</div>
						<Card.Title>{section.title}</Card.Title>
						<Card.Description>{section.description}</Card.Description>
					</Card.Header>
					<Card.Footer>
						<Button href={resolve(section.href)} size="sm">Open section</Button>
					</Card.Footer>
				</Card.Root>
			{/each}
		</section>

		<section class="hairline-frame p-6">
			<div class="flex flex-wrap items-end justify-between gap-4">
				<div>
					<p class="terminal-label">// coverage</p>
					<h3 class="mt-2 text-xl font-semibold">Represented component families</h3>
				</div>
				<Badge>{representedFamilies.length} families</Badge>
			</div>
			<Separator class="my-5" />
			<div class="flex flex-wrap gap-2">
				{#each representedFamilies as family (family)}
					<Badge variant="secondary">{family}</Badge>
				{/each}
			</div>
			<p class="mt-5 text-sm leading-6 text-muted-foreground">
				Skipped: <span class="font-mono text-foreground">data-table</span>, because the requested
				scope calls out the TanStack-powered table as intentionally out of bounds.
			</p>
		</section>
	</div>
</ShowcaseShell>
