<script lang="ts">
  import type { Pathname } from '$app/types';
  import { resolve } from '$app/paths';
  import { page } from '$app/state';
  import { locales, localizeHref } from '$lib/paraglide/runtime';
  import './layout.css';
  import favicon from '$lib/assets/favicon.svg';

  let { children } = $props();
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
  <meta
    name="description"
    content="A SvelteKit portfolio and component showcase for dense product screens."
  />
  <meta property="og:title" content="Portfolio components" />
  <meta
    property="og:description"
    content="A working Svelte component showcase tuned for dense product interfaces."
  />
</svelte:head>
<a
  href="#main-content"
  class="terminal-focus fixed top-3 left-3 z-70 -translate-y-16 border border-border bg-card px-3 py-2 font-mono text-xs font-semibold tracking-widest text-foreground uppercase transition-transform focus-visible:translate-y-0"
>
  Skip to content
</a>
{@render children()}
<div class="noise-overlay" aria-hidden="true"></div>

<div class="hidden">
  {#each locales as locale (locale)}
    <a href={resolve(localizeHref(page.url.pathname, { locale }) as Pathname)}>{locale}</a>
  {/each}
</div>
