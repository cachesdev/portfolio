<script lang="ts">
	import { DropdownMenu as DropdownMenuPrimitive } from 'bits-ui';
	import CaretRightIcon from 'phosphor-svelte/lib/CaretRight';
	import { cn } from '$lib/utils.js';

	let {
		ref = $bindable(null),
		class: className,
		inset,
		children,
		...restProps
	}: DropdownMenuPrimitive.SubTriggerProps & {
		inset?: boolean;
	} = $props();
</script>

<DropdownMenuPrimitive.SubTrigger
	bind:ref
	data-slot="dropdown-menu-sub-trigger"
	data-inset={inset}
	class={cn(
		"flex cursor-default items-center gap-2 border border-transparent px-3 py-2 font-mono text-xs font-medium uppercase tracking-[0.04em] text-zinc-300 outline-hidden select-none focus:border-[#3526ff]/50 focus:bg-[#18181b] focus:text-zinc-50 not-data-[variant=destructive]:focus:**:text-[#556af3] data-open:border-[#3526ff]/50 data-open:bg-[#18181b] data-open:text-zinc-50 data-inset:pl-9.5 data-[inset]:pl-8 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-3.5",
		className
	)}
	{...restProps}
>
	{@render children?.()}
	<CaretRightIcon class="ml-auto text-[#556af3]" />
</DropdownMenuPrimitive.SubTrigger>
