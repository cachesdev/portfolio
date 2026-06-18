<script lang="ts">
  import { resolve } from '$app/paths';
  import { Button } from '$lib/components/ui/button/index.js';
  import { showcaseSections } from '$lib/showcase/sections.js';

  const signalBlocks = [
    'col-start-7 col-span-1 row-start-1',
    'col-start-3 col-span-2 row-start-2',
    'col-start-5 col-span-3 row-start-2',
    'col-start-5 col-span-1 row-start-3',
    'col-start-6 col-span-2 row-start-3',
    'col-start-5 col-span-3 row-start-4',
    'col-start-3 col-span-2 row-start-5',
    'col-start-5 col-span-3 row-start-5',
    'col-start-5 col-span-1 row-start-6',
    'col-start-6 col-span-2 row-start-6',
    'col-start-1 col-span-1 row-start-7',
    'col-start-2 col-span-6 row-start-7',
    'col-start-3 col-span-2 row-start-8',
    'col-start-5 col-span-3 row-start-8'
  ];
</script>

<svelte:head>
  <title>Portfolio components</title>
  <meta
    name="description"
    content="A SvelteKit component showcase for forms, overlays, navigation, and data-heavy product screens."
  />
</svelte:head>

<main id="main-content" class="min-h-[100dvh] bg-background text-foreground">
  <section class="terminal-grid border-b border-border">
    <div class="mx-auto grid max-w-7xl gap-8 px-4 py-8 sm:px-6 lg:px-8 lg:py-10">
      <nav class="flex items-center justify-between gap-4" aria-label="Primary">
        <a
          href={resolve('/')}
          class="terminal-focus font-mono text-sm font-semibold tracking-widest"
        >
          portfolio.sys
        </a>
        <Button href={resolve('/showcase')} size="sm">Browse components</Button>
      </nav>

      <div
        class="grid overflow-hidden border border-border bg-background/70 lg:grid-cols-[minmax(0,0.93fr)_minmax(30rem,1.07fr)]"
      >
        <div
          class="grid content-center gap-7 border-border p-5 sm:p-8 lg:min-h-[30rem] lg:border-r lg:p-10"
        >
          <div class="grid gap-4">
            <h1
              class="max-w-4xl text-balance text-5xl leading-[0.95] font-semibold tracking-tight text-foreground sm:text-6xl lg:text-7xl"
            >
              Components for dense product screens.
            </h1>
            <p class="max-w-2xl text-lg leading-8 text-muted-foreground text-pretty">
              Forms, menus, navigation, and data surfaces shown in real states, with enough pressure
              around them to reveal whether they hold up.
            </p>
          </div>

          <div class="flex flex-wrap gap-3">
            <Button href={resolve('/showcase')} size="lg">Browse components</Button>
            <Button href={resolve('/showcase/forms')} variant="outline" size="lg"
              >Review inputs</Button
            >
          </div>

          <div
            class="grid max-w-xl gap-2 border-t border-border pt-5 text-sm leading-6 text-muted-foreground"
          >
            <p>
              The examples are intentionally compact: labels wrap, menus collide with edges, error
              copy appears inline, and selected states share one accent.
            </p>
          </div>
        </div>

        <div
          class="signal-map relative min-h-80 overflow-hidden border-t border-border lg:min-h-[30rem] lg:border-t-0"
          aria-label="Grid-aligned signal blocks"
        >
          <div class="grid h-full min-h-80 grid-cols-7 grid-rows-8 lg:min-h-[30rem]">
            {#each signalBlocks as block, index (index)}
              <div
                class={[
                  'signal-block m-2 h-auto w-auto border border-primary/20',
                  block,
                  index % 5 === 0 && 'opacity-80',
                  index % 4 === 1 && 'opacity-95'
                ]}
                aria-hidden="true"
              ></div>
            {/each}
          </div>
        </div>
      </div>
    </div>
  </section>

  <section class="mx-auto grid max-w-7xl gap-6 px-4 py-10 sm:px-6 lg:px-8 lg:py-14">
    <div class="grid gap-4 md:grid-cols-[0.72fr_1.28fr] md:items-end">
      <div>
        <h2 class="mt-3 text-3xl font-semibold tracking-tight text-foreground">
          Choose the kind of component you want to stress.
        </h2>
      </div>
      <p class="max-w-3xl text-base leading-7 text-muted-foreground text-pretty">
        Each section is grouped around the jobs these pieces do in product screens: collecting
        input, moving through a layout, opening temporary UI, or reading dense information.
      </p>
    </div>

    <div class="grid gap-4 lg:grid-cols-4">
      {#each showcaseSections as section (section.href)}
        <a
          href={resolve(section.href)}
          class="terminal-focus group grid min-h-56 content-between border border-border bg-card transition-colors hover:border-primary/55 hover:bg-primary/10"
        >
          <div class="grid gap-3 p-5">
            <div>
              <h3
                class="font-mono text-sm font-semibold tracking-[0.12em] text-foreground uppercase"
              >
                {section.title}
              </h3>
              <p class="mt-3 text-sm leading-6 text-muted-foreground group-hover:text-foreground">
                {section.description}
              </p>
            </div>
          </div>
          <div class="border-t border-border p-5">
            <span class="font-mono text-xs font-semibold tracking-widest text-primary uppercase">
              Open section
            </span>
          </div>
        </a>
      {/each}
    </div>

    <div class="grid gap-4 lg:grid-cols-[1.15fr_0.85fr]">
      <div class="hairline-frame terminal-grid-sm grid gap-5 p-5 lg:col-span-2">
        <div>
          <h3 class="text-xl font-semibold tracking-tight">Useful things to look for</h3>
        </div>
        <div class="grid gap-3 md:grid-cols-3">
          <div class="border-t border-border pt-4">
            <h4 class="font-mono text-xs font-semibold tracking-widest text-foreground uppercase">
              Errors beside fields
            </h4>
            <p class="mt-2 text-sm leading-6 text-muted-foreground">
              Forms show inline validation without moving the whole page around.
            </p>
          </div>
          <div class="border-t border-border pt-4">
            <h4 class="font-mono text-xs font-semibold tracking-widest text-foreground uppercase">
              Menus under pressure
            </h4>
            <p class="mt-2 text-sm leading-6 text-muted-foreground">
              Dialogs, sheets, command menus, and tooltips keep focus paths visible.
            </p>
          </div>
          <div class="border-t border-border pt-4">
            <h4 class="font-mono text-xs font-semibold tracking-widest text-foreground uppercase">
              Dense data
            </h4>
            <p class="mt-2 text-sm leading-6 text-muted-foreground">
              Tables, calendars, charts, empty states, and skeletons share the same rhythm.
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>
