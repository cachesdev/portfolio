<script lang="ts">
	import { cn, type WithElementRef } from '$lib/utils.js';
	import type { HTMLAttributes } from 'svelte/elements';

	let {
		ref = $bindable(null),
		class: className,
		children,
		...restProps
	}: WithElementRef<HTMLAttributes<HTMLDivElement>> = $props();
</script>

<div
	bind:this={ref}
	data-slot="empty-description"
	class={cn(
		'mt-0.5 text-sm/relaxed text-zinc-500 [&>a]:text-[#556af3] [&>a]:underline [&>a]:underline-offset-4 [&>a:hover]:text-[#aab4ff]',
		className
	)}
	{...restProps}
>
	{@render children?.()}
</div>
