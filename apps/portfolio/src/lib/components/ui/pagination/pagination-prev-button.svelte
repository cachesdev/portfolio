<script lang="ts">
	import { Pagination as PaginationPrimitive } from 'bits-ui';
	import CaretLeftIcon from 'phosphor-svelte/lib/CaretLeft';
	import { cn } from '$lib/utils.js';
	import { buttonVariants } from '../button/index.js';

	let {
		ref = $bindable(null),
		class: className,
		children,
		...restProps
	}: PaginationPrimitive.PrevButtonProps = $props();
</script>

{#snippet Fallback()}
	<CaretLeftIcon class={cn('size-4', className)} />
	<span>Previous</span>
{/snippet}

<PaginationPrimitive.PrevButton
	bind:ref
	aria-label="Go to previous page"
	class={cn(
		buttonVariants({ variant: 'ghost' }),
		'rounded-full border border-zinc-800 bg-black pl-2! font-mono text-xs font-semibold tracking-[0.08em] text-zinc-400 uppercase hover:border-[#556af3] hover:bg-[#3526ff] hover:text-white',
		className
	)}
	{...restProps}
>
	{#if children}
		{@render children?.()}
	{:else}
		{@render Fallback()}
	{/if}
</PaginationPrimitive.PrevButton>
