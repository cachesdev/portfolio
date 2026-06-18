<script lang="ts">
  import * as Accordion from '$lib/components/ui/accordion/index.js';
  import { Badge } from '$lib/components/ui/badge/index.js';
  import * as Breadcrumb from '$lib/components/ui/breadcrumb/index.js';
  import * as Collapsible from '$lib/components/ui/collapsible/index.js';
  import * as NavigationMenu from '$lib/components/ui/navigation-menu/index.js';
  import * as Pagination from '$lib/components/ui/pagination/index.js';
  import * as Resizable from '$lib/components/ui/resizable/index.js';
  import { ScrollArea } from '$lib/components/ui/scroll-area/index.js';
  import { Separator } from '$lib/components/ui/separator/index.js';
  import * as Sidebar from '$lib/components/ui/sidebar/index.js';
  import * as Tabs from '$lib/components/ui/tabs/index.js';

  type NavItem = {
    title: string;
    href: string;
    eyebrow: string;
  };

  type SidebarItem = {
    label: string;
    count: string;
    active?: boolean;
  };

  const navItems: NavItem[] = [
    { title: 'Grid Shell', href: '#layout-showcase-grid', eyebrow: 'Responsive frame' },
    { title: 'Flow Stack', href: '#layout-showcase-flow', eyebrow: 'Progressive disclosure' },
    { title: 'Review Queue', href: '#layout-showcase-review', eyebrow: 'Paged scans' }
  ];

  const sidebarItems: SidebarItem[] = [
    { label: 'Overview', count: '12', active: true },
    { label: 'Routes', count: '08' },
    { label: 'Panels', count: '04' },
    { label: 'Logs', count: '19' }
  ];

  const logLines = [
    '00:04 layout:grid template columns locked',
    '00:08 nav:menu viewport measured at 416px',
    '00:11 tabs:value changed to canvas',
    '00:17 resizable:pane stored 62/38 split',
    '00:22 scroll-area kept the latest entry visible',
    '00:29 accordion:review notes opened',
    '00:31 pagination:page window recalculated',
    '00:38 collapsible:notes expanded'
  ];

  let activeTab = $state('canvas');
  let activeGate = $state('motion');
  let notesOpen = $state(true);
  let reviewPage = $state(2);
  let sidebarOpen = $state(true);

  const activeReview = $derived(`Batch ${String(reviewPage).padStart(2, '0')}`);
</script>

{#snippet NavCard({ item }: { item: NavItem })}
  <li>
    <NavigationMenu.Link
      href={item.href}
      class="grid min-w-0 gap-1 border-border/80 bg-background/80 p-3 hover:bg-[#18181b]"
    >
      <span class="text-zinc-100">{item.title}</span>
      <span class="text-[11px] tracking-normal normal-case text-terminal-muted">{item.eyebrow}</span
      >
    </NavigationMenu.Link>
  </li>
{/snippet}

{#snippet StatusRow({ label, value }: { label: string; value: string })}
  <div
    class="grid grid-cols-[1fr_auto] items-center gap-3 border-b border-border/70 px-3 py-2 last:border-b-0"
  >
    <span class="truncate font-mono text-[11px] tracking-[0.08em] text-zinc-400 uppercase"
      >{label}</span
    >
    <span class="font-mono text-xs font-semibold text-terminal-violet-soft">{value}</span>
  </div>
{/snippet}

<section
  id="layout-showcase"
  aria-labelledby="layout-showcase-title"
  class="grid gap-5 text-foreground"
>
  <div class="grid gap-5">
    <div class="grid gap-4 border border-border bg-background/90 p-4 sm:p-5">
      <div class="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
        <div class="grid gap-3">
          <Breadcrumb.Root>
            <Breadcrumb.List class="font-mono text-xs">
              <Breadcrumb.Item>
                <Breadcrumb.Link href="/">Portfolio</Breadcrumb.Link>
              </Breadcrumb.Item>
              <Breadcrumb.Separator />
              <Breadcrumb.Item>
                <Breadcrumb.Link href="#layout-showcase">Showcase</Breadcrumb.Link>
              </Breadcrumb.Item>
              <Breadcrumb.Separator />
              <Breadcrumb.Item>
                <Breadcrumb.Page>Layout</Breadcrumb.Page>
              </Breadcrumb.Item>
            </Breadcrumb.List>
          </Breadcrumb.Root>

          <div class="grid gap-2">
            <h2
              id="layout-showcase-title"
              class="text-2xl font-semibold tracking-normal text-zinc-50 sm:text-3xl"
            >
              Navigation inside a cramped workspace
            </h2>
            <p class="max-w-2xl text-sm leading-6 text-muted-foreground">
              Breadcrumbs, menu content, tabs, a collapsible sidebar, split panes, scroll areas, and
              pagination share one frame.
            </p>
          </div>
        </div>

        <NavigationMenu.Root viewport={false} class="z-20 max-w-none justify-start lg:justify-end">
          <NavigationMenu.List
            class="justify-start rounded-full border border-border bg-background p-1"
          >
            <NavigationMenu.Item>
              <NavigationMenu.Trigger>Systems</NavigationMenu.Trigger>
              <NavigationMenu.Content class="right-0! left-auto! w-[min(34rem,calc(100vw-3rem))]!">
                <ul class="grid gap-2 p-1 sm:grid-cols-3">
                  {#each navItems as item (item.title)}
                    {@render NavCard({ item })}
                  {/each}
                </ul>
              </NavigationMenu.Content>
            </NavigationMenu.Item>
            <NavigationMenu.Item>
              <NavigationMenu.Link
                href="#layout-showcase-review"
                class="h-9 rounded-full border-transparent px-4.5 py-2.5 hover:border-transparent hover:bg-zinc-50 hover:text-zinc-950 focus:border-transparent focus:bg-zinc-50 focus:text-zinc-950"
              >
                Review
              </NavigationMenu.Link>
            </NavigationMenu.Item>
          </NavigationMenu.List>
        </NavigationMenu.Root>
      </div>

      <Separator />

      <div class="grid gap-4">
        <div
          id="layout-showcase-grid"
          class="min-h-105 overflow-hidden border border-border bg-background"
        >
          <Sidebar.Provider
            bind:open={sidebarOpen}
            class="relative h-full! min-h-105! bg-transparent"
            style="--sidebar-width: 17rem; --sidebar-width-icon: 3rem;"
          >
            <Sidebar.Sidebar
              collapsible="icon"
              class="absolute! z-30! h-full! min-h-0! border-r border-border bg-[#09090b]"
            >
              <Sidebar.Header class="group-data-[collapsible=icon]:items-center">
                <div
                  class="flex items-center justify-between gap-3 group-data-[collapsible=icon]:justify-center"
                >
                  <div class="grid gap-1 group-data-[collapsible=icon]:hidden">
                    <span class="font-mono text-sm font-semibold text-zinc-100">Project files</span>
                  </div>
                  <Sidebar.Trigger aria-label="Toggle showcase sidebar" />
                </div>
              </Sidebar.Header>
              <Sidebar.Content>
                <Sidebar.Group>
                  <Sidebar.GroupLabel>Navigation</Sidebar.GroupLabel>
                  <Sidebar.GroupContent>
                    <Sidebar.Menu>
                      {#each sidebarItems as item (item.label)}
                        <Sidebar.MenuItem>
                          <Sidebar.MenuButton
                            isActive={item.active}
                            tooltipContent={item.label}
                            class="group-data-[collapsible=icon]:justify-center"
                          >
                            <span class="size-1.5 shrink-0 bg-terminal-violet-soft"></span>
                            <span class="group-data-[collapsible=icon]:hidden">{item.label}</span>
                          </Sidebar.MenuButton>
                          <Sidebar.MenuBadge>{item.count}</Sidebar.MenuBadge>
                        </Sidebar.MenuItem>
                      {/each}
                    </Sidebar.Menu>
                  </Sidebar.GroupContent>
                </Sidebar.Group>
              </Sidebar.Content>
              <Sidebar.Footer>
                <div
                  class="grid grid-cols-2 gap-2 font-mono text-[11px] text-zinc-500 group-data-[collapsible=icon]:flex group-data-[collapsible=icon]:justify-center"
                >
                  <span class="group-data-[collapsible=icon]:hidden">uptime</span>
                  <span class="text-right text-success group-data-[collapsible=icon]:hidden"
                    >99.41%</span
                  >
                  <span
                    aria-label="Uptime healthy"
                    class="hidden size-1.5 bg-success group-data-[collapsible=icon]:block"
                  ></span>
                </div>
              </Sidebar.Footer>
            </Sidebar.Sidebar>
            <Sidebar.Inset class="min-h-105 border-0 bg-background">
              <Tabs.Root id="layout-showcase-flow" bind:value={activeTab} class="min-w-0">
                <Tabs.List
                  variant="line"
                  class="w-full justify-start overflow-x-auto border-border bg-background"
                >
                  <Tabs.Trigger value="canvas">Canvas</Tabs.Trigger>
                  <Tabs.Trigger value="signals">Rules</Tabs.Trigger>
                  <Tabs.Trigger value="review">Review</Tabs.Trigger>
                </Tabs.List>

                <Tabs.Content value="canvas" class="min-h-93.5 p-0">
                  <Resizable.PaneGroup direction="horizontal" class="min-h-93.5">
                    <Resizable.Pane defaultSize={62} minSize={42}>
                      <div class="grid h-full grid-rows-[auto_1fr]">
                        <div
                          class="flex items-center justify-between border-b border-border px-4 py-3"
                        >
                          <span
                            class="font-mono text-xs font-semibold tracking-widest text-muted-foreground uppercase"
                            >Resizable panes</span
                          >
                          <Badge
                            variant="outline"
                            class="border-terminal-violet-soft/60 text-terminal-violet-soft"
                          >
                            Live
                          </Badge>
                        </div>
                        <div class="terminal-grid-sm grid content-center gap-4 p-4">
                          <div class="grid gap-3 border border-border bg-background/80 p-4">
                            <div class="h-2 w-24 bg-terminal-violet"></div>
                            <div class="grid gap-2 sm:grid-cols-3">
                              <div class="h-20 border border-border bg-[#18181b]"></div>
                              <div class="h-20 border border-border bg-[#09090b]"></div>
                              <div
                                class="h-20 border border-terminal-violet-soft/50 bg-[#18181b]"
                              ></div>
                            </div>
                          </div>
                          <div class="grid grid-cols-3 gap-2 font-mono text-[11px] text-zinc-500">
                            <span>cols: 12</span>
                            <span>gap: 16</span>
                            <span>snap: on</span>
                          </div>
                        </div>
                      </div>
                    </Resizable.Pane>
                    <Resizable.Handle withHandle />
                    <Resizable.Pane defaultSize={38} minSize={28}>
                      <div class="grid h-full grid-rows-[auto_1fr]">
                        <div class="border-b border-border px-4 py-3">
                          <span
                            class="font-mono text-xs font-semibold tracking-widest text-muted-foreground uppercase"
                            >Activity log</span
                          >
                        </div>
                        <ScrollArea class="h-79.5" orientation="vertical">
                          <div class="grid gap-2 p-3">
                            {#each logLines as line (line)}
                              <div
                                class="border border-border bg-[#09090b] px-3 py-2 font-mono text-xs text-zinc-400"
                              >
                                {line}
                              </div>
                            {/each}
                          </div>
                        </ScrollArea>
                      </div>
                    </Resizable.Pane>
                  </Resizable.PaneGroup>
                </Tabs.Content>

                <Tabs.Content value="signals" class="min-h-93.5">
                  <div class="grid gap-4 p-4 lg:grid-cols-[minmax(0,1fr)_16rem]">
                    <Accordion.Root type="single" bind:value={activeGate}>
                      <Accordion.Item value="motion">
                        <Accordion.Trigger>Motion</Accordion.Trigger>
                        <Accordion.Content>
                          <p>Transitions stay short, directional, and tied to component state.</p>
                        </Accordion.Content>
                      </Accordion.Item>
                      <Accordion.Item value="density">
                        <Accordion.Trigger>Density</Accordion.Trigger>
                        <Accordion.Content>
                          <p>Panels keep rows readable without stretching the whole workspace.</p>
                        </Accordion.Content>
                      </Accordion.Item>
                      <Accordion.Item value="access">
                        <Accordion.Trigger>Keyboard paths</Accordion.Trigger>
                        <Accordion.Content>
                          <p>
                            Keyboard navigation remains native through the shadcn-svelte primitives.
                          </p>
                        </Accordion.Content>
                      </Accordion.Item>
                    </Accordion.Root>

                    <Collapsible.Root bind:open={notesOpen}>
                      <Collapsible.Trigger>
                        <span>Release notes</span>
                        <span class="text-terminal-violet-soft"
                          >{notesOpen ? 'Open' : 'Closed'}</span
                        >
                      </Collapsible.Trigger>
                      <Collapsible.Content>
                        <div class="grid gap-2">
                          {@render StatusRow({ label: 'Sidebar', value: 'Collapsible' })}
                          {@render StatusRow({ label: 'Tabs', value: 'Keyboard' })}
                          {@render StatusRow({ label: 'Panels', value: 'Resizable' })}
                        </div>
                      </Collapsible.Content>
                    </Collapsible.Root>
                  </div>
                </Tabs.Content>

                <Tabs.Content id="layout-showcase-review" value="review" class="min-h-93.5 p-4">
                  <div class="grid h-full content-between gap-5">
                    <div class="grid gap-3 sm:grid-cols-3">
                      <div class="border border-border bg-[#09090b] p-4">
                        <p
                          class="font-mono text-xs font-semibold tracking-widest text-muted-foreground uppercase"
                        >
                          Page
                        </p>
                        <p class="mt-3 font-mono text-2xl text-zinc-50">{activeReview}</p>
                      </div>
                      <div class="border border-border bg-[#09090b] p-4">
                        <p
                          class="font-mono text-xs font-semibold tracking-widest text-muted-foreground uppercase"
                        >
                          Total pages
                        </p>
                        <p class="mt-3 font-mono text-2xl text-zinc-50">06</p>
                      </div>
                      <div class="border border-border bg-[#09090b] p-4">
                        <p
                          class="font-mono text-xs font-semibold tracking-widest text-muted-foreground uppercase"
                        >
                          Rows
                        </p>
                        <p class="mt-3 font-mono text-2xl text-zinc-50">72</p>
                      </div>
                    </div>

                    <Pagination.Root
                      count={72}
                      perPage={12}
                      siblingCount={1}
                      bind:page={reviewPage}
                      class="justify-start"
                    >
                      {#snippet children({ pages, range })}
                        <div class="grid gap-3">
                          <Pagination.Content class="mx-0 w-fit rounded-none">
                            <Pagination.Item>
                              <Pagination.Previous />
                            </Pagination.Item>
                            {#each pages as page (page.key)}
                              <Pagination.Item>
                                {#if page.type === 'ellipsis'}
                                  <Pagination.Ellipsis />
                                {:else}
                                  <Pagination.Link {page} isActive={page.value === reviewPage} />
                                {/if}
                              </Pagination.Item>
                            {/each}
                            <Pagination.Item>
                              <Pagination.Next />
                            </Pagination.Item>
                          </Pagination.Content>
                          <p class="font-mono text-xs text-zinc-500">
                            Showing {range.start}-{range.end} of 72 layout checks.
                          </p>
                        </div>
                      {/snippet}
                    </Pagination.Root>
                  </div>
                </Tabs.Content>
              </Tabs.Root>
            </Sidebar.Inset>
          </Sidebar.Provider>
        </div>
      </div>
    </div>
  </div>
</section>
