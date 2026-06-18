<script lang="ts">
  import { toast } from 'svelte-sonner';
  import BellIcon from 'phosphor-svelte/lib/Bell';
  import CommandIcon from 'phosphor-svelte/lib/Command';
  import CopyIcon from 'phosphor-svelte/lib/Copy';
  import CursorClickIcon from 'phosphor-svelte/lib/CursorClick';
  import DotsThreeVerticalIcon from 'phosphor-svelte/lib/DotsThreeVertical';
  import GearIcon from 'phosphor-svelte/lib/Gear';
  import LightningIcon from 'phosphor-svelte/lib/Lightning';
  import RowsIcon from 'phosphor-svelte/lib/Rows';
  import ShieldWarningIcon from 'phosphor-svelte/lib/ShieldWarning';
  import SidebarSimpleIcon from 'phosphor-svelte/lib/SidebarSimple';
  import TrashIcon from 'phosphor-svelte/lib/Trash';

  import * as AlertDialog from '$lib/components/ui/alert-dialog/index.js';
  import { Button } from '$lib/components/ui/button/index.js';
  import * as Command from '$lib/components/ui/command/index.js';
  import * as ContextMenu from '$lib/components/ui/context-menu/index.js';
  import * as Dialog from '$lib/components/ui/dialog/index.js';
  import * as Drawer from '$lib/components/ui/drawer/index.js';
  import * as DropdownMenu from '$lib/components/ui/dropdown-menu/index.js';
  import * as HoverCard from '$lib/components/ui/hover-card/index.js';
  import * as Menubar from '$lib/components/ui/menubar/index.js';
  import * as Popover from '$lib/components/ui/popover/index.js';
  import * as Sheet from '$lib/components/ui/sheet/index.js';
  import { Toaster } from '$lib/components/ui/sonner/index.js';
  import * as Tooltip from '$lib/components/ui/tooltip/index.js';

  let commandOpen = $state(false);

  const quickActions = ['Open command menu', 'Inspect build', 'Preview sheet', 'Send toast'];
  const commandItems = [
    { label: 'Deploy preview', shortcut: 'D P' },
    { label: 'Copy endpoint', shortcut: 'C E' },
    { label: 'Open audit trail', shortcut: 'A T' }
  ];

  function signalToast() {
    toast.success('Toast sent', {
      description: 'This uses the same toaster that app screens use.'
    });
  }
</script>

<section
  class="relative isolate grid gap-5 text-foreground"
  aria-labelledby="overlays-showcase-title"
>
  <Toaster richColors position="bottom-right" />

  <div class="flex w-full flex-col gap-5">
    <div class="grid gap-4 lg:grid-cols-[minmax(0,1fr)_20rem] lg:items-stretch">
      <div class="hairline-frame terminal-grid-sm space-y-3 p-5">
        <h2
          id="overlays-showcase-title"
          class="max-w-3xl text-2xl font-semibold tracking-normal text-zinc-50 sm:text-3xl"
        >
          Temporary UI without losing context
        </h2>
        <p class="max-w-2xl text-sm leading-6 text-zinc-400">
          Open the examples and check how focus, escape behavior, hover content, and toast feedback
          feel in the same page.
        </p>
      </div>

      <div class="hairline-frame grid content-center gap-3 p-5">
        <h3 class="font-mono text-xs font-semibold tracking-widest text-foreground uppercase">
          Try this
        </h3>
        <p class="text-sm leading-6 text-muted-foreground">
          Open a dialog, move through a menu with the keyboard, then trigger a toast. The page
          should stay readable behind each layer.
        </p>
      </div>
    </div>

    <div class="grid items-start gap-4 lg:grid-cols-[minmax(0,1.12fr)_minmax(21rem,0.88fr)]">
      <div class="terminal-surface">
        <div class="flex items-center justify-between border-b border-border px-4 py-3">
          <div>
            <h3 class="text-base font-semibold text-zinc-50">Dialogs, drawer, sheet, and toast</h3>
          </div>
          <span class="h-2 w-20 bg-[#b9d765]"></span>
        </div>

        <div class="grid gap-3 p-4 sm:grid-cols-2">
          <AlertDialog.Root>
            <AlertDialog.Trigger>
              {#snippet child({ props })}
                <Button {...props} variant="destructive" class="w-full justify-start">
                  <ShieldWarningIcon data-icon="inline-start" />
                  Reset key
                </Button>
              {/snippet}
            </AlertDialog.Trigger>
            <AlertDialog.Content>
              <AlertDialog.Header>
                <AlertDialog.Title>Rotate production key?</AlertDialog.Title>
                <AlertDialog.Description>
                  Existing sessions continue, but new deploys must use the replacement token.
                </AlertDialog.Description>
              </AlertDialog.Header>
              <AlertDialog.Footer>
                <AlertDialog.Cancel>Cancel</AlertDialog.Cancel>
                <AlertDialog.Action>Rotate key</AlertDialog.Action>
              </AlertDialog.Footer>
            </AlertDialog.Content>
          </AlertDialog.Root>

          <Dialog.Root>
            <Dialog.Trigger>
              {#snippet child({ props })}
                <Button {...props} variant="outline" class="w-full justify-start">
                  <RowsIcon data-icon="inline-start" />
                  Open dialog
                </Button>
              {/snippet}
            </Dialog.Trigger>
            <Dialog.Content>
              <Dialog.Header>
                <Dialog.Title>Release snapshot</Dialog.Title>
                <Dialog.Description>
                  Review the current overlay bundle before it is promoted.
                </Dialog.Description>
              </Dialog.Header>
              <div class="grid gap-2 font-mono text-xs uppercase tracking-[0.06em] text-zinc-400">
                <div
                  class="flex items-center justify-between border border-border bg-background px-3 py-2"
                >
                  <span>Command palette</span>
                  <span class="text-[#28c840]">opens with focus</span>
                </div>
                <div
                  class="flex items-center justify-between border border-border bg-background px-3 py-2"
                >
                  <span>Menu surfaces</span>
                  <span class="text-[#d0e891]">keyboardable</span>
                </div>
              </div>
              <Dialog.Footer>
                <Dialog.Close>
                  {#snippet child({ props })}
                    <Button {...props} variant="outline">Close</Button>
                  {/snippet}
                </Dialog.Close>
                <Button onclick={signalToast}>Show toast</Button>
              </Dialog.Footer>
            </Dialog.Content>
          </Dialog.Root>

          <Drawer.Root shouldScaleBackground={false}>
            <Drawer.Trigger>
              {#snippet child({ props })}
                <Button {...props} variant="outline" class="w-full justify-start">
                  <RowsIcon data-icon="inline-start" />
                  Open drawer
                </Button>
              {/snippet}
            </Drawer.Trigger>
            <Drawer.Content class="px-4 pb-4">
              <Drawer.Header>
                <Drawer.Title>Queued actions</Drawer.Title>
                <Drawer.Description
                  >Swipe or press escape to dismiss this bottom drawer.</Drawer.Description
                >
              </Drawer.Header>
              <div class="grid gap-2 py-2">
                {#each quickActions as action (action)}
                  <div
                    class="flex items-center gap-3 border border-border bg-background px-3 py-2 text-sm text-zinc-300"
                  >
                    <LightningIcon class="size-4 text-[#d0e891]" />
                    <span>{action}</span>
                  </div>
                {/each}
              </div>
              <Drawer.Footer class="px-0">
                <Drawer.Close>
                  {#snippet child({ props })}
                    <Button {...props} variant="secondary">Done</Button>
                  {/snippet}
                </Drawer.Close>
              </Drawer.Footer>
            </Drawer.Content>
          </Drawer.Root>

          <Sheet.Root>
            <Sheet.Trigger>
              {#snippet child({ props })}
                <Button {...props} variant="outline" class="w-full justify-start">
                  <SidebarSimpleIcon data-icon="inline-start" />
                  Open sheet
                </Button>
              {/snippet}
            </Sheet.Trigger>
            <Sheet.Content side="right" class="gap-5 p-5">
              <Sheet.Header>
                <Sheet.Title>Inspector sheet</Sheet.Title>
                <Sheet.Description
                  >Side panels keep context visible while details update.</Sheet.Description
                >
              </Sheet.Header>
              <div class="grid gap-3">
                <div class="border border-border bg-background p-3">
                  <p
                    class="font-mono text-xs font-semibold tracking-widest text-muted-foreground uppercase"
                  >
                    Latency
                  </p>
                  <p class="mt-2 text-2xl font-semibold text-zinc-50">42ms</p>
                </div>
                <div class="border border-border bg-background p-3">
                  <p
                    class="font-mono text-xs font-semibold tracking-widest text-muted-foreground uppercase"
                  >
                    Focus trap
                  </p>
                  <p class="mt-2 text-sm text-zinc-300">Handled by the dialog primitive.</p>
                </div>
              </div>
              <Sheet.Footer>
                <Sheet.Close>
                  {#snippet child({ props })}
                    <Button {...props} variant="outline">Close sheet</Button>
                  {/snippet}
                </Sheet.Close>
              </Sheet.Footer>
            </Sheet.Content>
          </Sheet.Root>

          <Button onclick={signalToast} variant="secondary" class="w-full justify-start">
            <BellIcon data-icon="inline-start" />
            Toast
          </Button>
        </div>
      </div>

      <div class="terminal-surface">
        <div class="flex items-center justify-between border-b border-border px-4 py-3">
          <div>
            <h3 class="text-base font-semibold text-zinc-50">Command menu</h3>
          </div>
          <Button variant="outline" size="sm" onclick={() => (commandOpen = true)}>
            <CommandIcon data-icon="inline-start" />
            Open
          </Button>
        </div>

        <div class="p-4">
          <Command.Root>
            <Command.Input placeholder="Search overlays..." />
            <Command.List>
              <Command.Empty>No overlay found.</Command.Empty>
              <Command.Group heading="Actions">
                {#each commandItems as item (item.label)}
                  <Command.Item value={item.label}>
                    <CommandIcon class="size-4" />
                    <span>{item.label}</span>
                    <Command.Shortcut>{item.shortcut}</Command.Shortcut>
                  </Command.Item>
                {/each}
              </Command.Group>
            </Command.List>
          </Command.Root>
        </div>

        <Command.Dialog
          bind:open={commandOpen}
          title="Command menu"
          description="Jump to a showcase action."
        >
          <Command.Input placeholder="Type a command..." />
          <Command.List>
            <Command.Empty>No command matched.</Command.Empty>
            <Command.Group heading="Jump">
              <Command.Item value="Dialog">
                <RowsIcon class="size-4" />
                <span>Open dialog sample</span>
                <Command.Shortcut>J D</Command.Shortcut>
              </Command.Item>
              <Command.Item value="Toast">
                <BellIcon class="size-4" />
                <span>Send toast</span>
                <Command.Shortcut>J T</Command.Shortcut>
              </Command.Item>
            </Command.Group>
          </Command.List>
        </Command.Dialog>
      </div>
    </div>

    <div class="grid items-start gap-4 lg:grid-cols-[minmax(0,1fr)_minmax(0,1fr)]">
      <div class="terminal-surface">
        <div class="border-b border-border px-4 py-3">
          <h3 class="text-base font-semibold text-zinc-50">Dropdown, menubar, and context menu</h3>
        </div>

        <div class="grid gap-4 p-4">
          <Menubar.Root>
            <Menubar.Menu>
              <Menubar.Trigger>File</Menubar.Trigger>
              <Menubar.Content>
                <Menubar.Item>
                  New overlay
                  <Menubar.Shortcut>Cmd N</Menubar.Shortcut>
                </Menubar.Item>
                <Menubar.Item>
                  Save preset
                  <Menubar.Shortcut>Cmd S</Menubar.Shortcut>
                </Menubar.Item>
                <Menubar.Separator />
                <Menubar.Item>Export log</Menubar.Item>
              </Menubar.Content>
            </Menubar.Menu>
            <Menubar.Menu>
              <Menubar.Trigger>View</Menubar.Trigger>
              <Menubar.Content>
                <Menubar.CheckboxItem checked>Grid</Menubar.CheckboxItem>
                <Menubar.CheckboxItem>Focus rings</Menubar.CheckboxItem>
              </Menubar.Content>
            </Menubar.Menu>
          </Menubar.Root>

          <div class="flex flex-wrap gap-3">
            <DropdownMenu.Root>
              <DropdownMenu.Trigger>
                {#snippet child({ props })}
                  <Button {...props} variant="outline">
                    <DotsThreeVerticalIcon data-icon="inline-start" />
                    Dropdown
                  </Button>
                {/snippet}
              </DropdownMenu.Trigger>
              <DropdownMenu.Content align="start">
                <DropdownMenu.Label>Overlay tools</DropdownMenu.Label>
                <DropdownMenu.Item>
                  <CopyIcon class="size-4" />
                  Copy config
                  <DropdownMenu.Shortcut>Cmd C</DropdownMenu.Shortcut>
                </DropdownMenu.Item>
                <DropdownMenu.Item>
                  <GearIcon class="size-4" />
                  Preferences
                </DropdownMenu.Item>
                <DropdownMenu.Separator />
                <DropdownMenu.Item class="text-destructive">
                  <TrashIcon class="size-4" />
                  Clear queue
                </DropdownMenu.Item>
              </DropdownMenu.Content>
            </DropdownMenu.Root>

            <ContextMenu.Root>
              <ContextMenu.Trigger>
                {#snippet child({ props })}
                  <div
                    {...props}
                    aria-label="Open contextual actions with right click, long press, or keyboard menu key"
                    class="flex min-h-24 flex-1 select-none items-center justify-center gap-3 border border-dashed border-[#b9d765]/50 bg-background px-4 text-center font-mono text-xs uppercase tracking-[0.08em] text-zinc-400 outline-none transition-colors focus-visible:ring-2 focus-visible:ring-zinc-300/30"
                    role="button"
                    tabindex="0"
                  >
                    <CursorClickIcon class="size-5 text-[#d0e891]" />
                    Context target
                  </div>
                {/snippet}
              </ContextMenu.Trigger>
              <ContextMenu.Content>
                <ContextMenu.Label>Context actions</ContextMenu.Label>
                <ContextMenu.Item>
                  Inspect node
                  <ContextMenu.Shortcut>I</ContextMenu.Shortcut>
                </ContextMenu.Item>
                <ContextMenu.Item>Duplicate sample</ContextMenu.Item>
                <ContextMenu.Separator />
                <ContextMenu.Item class="text-destructive">Delete sample</ContextMenu.Item>
              </ContextMenu.Content>
            </ContextMenu.Root>
          </div>
        </div>
      </div>

      <Tooltip.Provider>
        <div class="terminal-surface">
          <div class="border-b border-border px-4 py-3">
            <h3 class="text-base font-semibold text-zinc-50">Hover card, popover, and tooltip</h3>
          </div>

          <div class="grid gap-3 p-4 sm:grid-cols-3">
            <HoverCard.Root>
              <HoverCard.Trigger>
                {#snippet child({ props })}
                  <Button {...props} variant="outline" class="w-full">Hover card</Button>
                {/snippet}
              </HoverCard.Trigger>
              <HoverCard.Content>
                <p
                  class="font-mono text-xs font-semibold tracking-widest text-muted-foreground uppercase"
                >
                  Hover content
                </p>
                <p class="mt-2 text-sm text-zinc-300">
                  Hover cards reveal lightweight context without committing focus.
                </p>
              </HoverCard.Content>
            </HoverCard.Root>

            <Popover.Root>
              <Popover.Trigger>
                {#snippet child({ props })}
                  <Button {...props} variant="outline" class="w-full">Popover</Button>
                {/snippet}
              </Popover.Trigger>
              <Popover.Content>
                <Popover.Header>
                  <Popover.Title>Density</Popover.Title>
                  <Popover.Description>Compact controls fit best in popovers.</Popover.Description>
                </Popover.Header>
                <div class="grid grid-cols-3 gap-2">
                  {#each ['S', 'M', 'L'] as size (size)}
                    <button
                      type="button"
                      class="border border-border bg-background py-2 font-mono text-xs text-zinc-300 transition-colors hover:border-[#d0e891] hover:text-zinc-50"
                    >
                      {size}
                    </button>
                  {/each}
                </div>
              </Popover.Content>
            </Popover.Root>

            <Tooltip.Root>
              <Tooltip.Trigger>
                {#snippet child({ props })}
                  <Button {...props} variant="outline" class="w-full">Tooltip</Button>
                {/snippet}
              </Tooltip.Trigger>
              <Tooltip.Content sideOffset={8}>Keyboard focus opens this too</Tooltip.Content>
            </Tooltip.Root>
          </div>
        </div>
      </Tooltip.Provider>
    </div>
  </div>
</section>
