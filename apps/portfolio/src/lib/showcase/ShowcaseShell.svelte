<script lang="ts">
  import { resolve } from '$app/paths';
  import { page } from '$app/state';
  import { Button } from '$lib/components/ui/button/index.js';
  import { Separator } from '$lib/components/ui/separator/index.js';
  import { showcaseSections } from './sections.js';
  import type { Snippet } from 'svelte';

  let {
    title = 'Component showcase',
    description = 'A working set of Svelte components shown in realistic product states.',
    children
  }: {
    title?: string;
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

<main id="main-content" class="min-h-[100dvh] bg-background text-foreground">
  <section class="terminal-grid border-b border-border">
    <div
      class="mx-auto grid max-w-7xl gap-8 px-4 py-8 sm:px-6 lg:grid-cols-[17rem_1fr] lg:px-8 lg:py-10"
    >
      <aside class="order-2 lg:sticky lg:top-6 lg:order-1 lg:self-start">
        <div class="hairline-frame terminal-grid-sm p-4">
          <a
            href={resolve('/showcase')}
            class="group block border-b border-border pb-4 outline-none focus-visible:ring-2 focus-visible:ring-ring/30"
          >
            <h1 class="text-2xl font-semibold tracking-tight text-foreground">Components</h1>
            <p class="mt-2 text-sm leading-6 text-muted-foreground">
              Jump between examples for forms, navigation, overlays, and data-heavy surfaces.
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
                    {section.short}
                  </span>
                </span>
                <span class="text-xs text-muted-foreground">{section.summary}</span>
              </a>
            {/each}
          </nav>

          <Separator class="my-4" />
          <p class="text-xs leading-5 text-muted-foreground">
            The TanStack data table has a separate model, so this pass keeps it out of the general
            component tour.
          </p>
        </div>
      </aside>

      <div class="order-1 min-w-0 lg:order-2">
        <header class="hairline-frame terminal-grid-sm mb-6 p-6 md:p-8">
          <div class="flex flex-wrap items-start justify-between gap-4">
            <div class="max-w-3xl">
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
