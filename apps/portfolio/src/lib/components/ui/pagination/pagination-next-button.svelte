<script lang="ts">
	import { Pagination as PaginationPrimitive } from 'bits-ui';
	import CaretRightIcon from 'phosphor-svelte/lib/CaretRight';
	import { cn } from '$lib/utils.js';
	import { buttonVariants } from '../button/index.js';

	let {
		ref = $bindable(null),
		class: className,
		children,
		...restProps
	}: PaginationPrimitive.NextButtonProps = $props();
</script>

{#snippet Fallback()}
	<span>Next</span>
	<CaretRightIcon class={cn('size-4', className)} />
{/snippet}

<PaginationPrimitive.NextButton
	bind:ref
	aria-label="Go to next page"
	class={cn(
		buttonVariants({ variant: 'ghost' }),
		'rounded-full border border-zinc-800 bg-black pr-2! font-mono text-xs font-semibold tracking-[0.08em] text-zinc-400 uppercase hover:border-[#556af3] hover:bg-[#3526ff] hover:text-white',
		className
	)}
	{...restProps}
>
	{#if children}
		{@render children?.()}
	{:else}
		{@render Fallback()}
	{/if}
</PaginationPrimitive.NextButton>
