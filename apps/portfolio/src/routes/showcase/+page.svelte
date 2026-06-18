<script lang="ts">
  import { resolve } from '$app/paths';
  import ShowcaseShell from '$lib/showcase/ShowcaseShell.svelte';
  import { showcaseSections } from '$lib/showcase/sections.js';
  import { Button } from '$lib/components/ui/button/index.js';
  import * as Card from '$lib/components/ui/card/index.js';
</script>

<ShowcaseShell
  title="Component groups"
  description="Pick a section by the kind of UI you need to test: forms, layout, overlays, or data display."
>
  <div class="grid gap-6">
    <section class="grid gap-4 md:grid-cols-2">
      {#each showcaseSections as section (section.href)}
        <Card.Root class="terminal-grid-sm">
          <Card.Header>
            <p class="font-mono text-xs font-semibold tracking-widest text-primary uppercase">
              {section.short}
            </p>
            <Card.Title>{section.title}</Card.Title>
            <Card.Description>{section.description}</Card.Description>
          </Card.Header>
          <Card.Footer>
            <Button href={resolve(section.href)} size="sm">Open section</Button>
          </Card.Footer>
        </Card.Root>
      {/each}
    </section>
  </div>
</ShowcaseShell>
