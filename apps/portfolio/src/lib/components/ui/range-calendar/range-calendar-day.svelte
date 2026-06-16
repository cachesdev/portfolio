<script lang="ts">
	import { RangeCalendar as RangeCalendarPrimitive } from 'bits-ui';
	import { cn } from '$lib/utils.js';

	let {
		ref = $bindable(null),
		class: className,
		...restProps
	}: RangeCalendarPrimitive.DayProps = $props();
</script>

<RangeCalendarPrimitive.Day
	bind:ref
	class={cn(
		'flex size-(--cell-size) flex-col items-center justify-center gap-1 rounded-(--cell-radius) border border-transparent p-0 font-mono text-xs leading-none font-medium whitespace-nowrap text-zinc-300 select-none',
		'not-data-selected:hover:border-zinc-700 not-data-selected:hover:bg-zinc-950 not-data-selected:hover:text-white',
		'[&[data-today]:not([data-selected])]:border-[#556af3]/60 [&[data-today]:not([data-selected])]:bg-zinc-950 [&[data-today]:not([data-selected])]:text-[#aab4ff] [&[data-today][data-disabled]]:text-zinc-500 data-[range-middle]:rounded-none',
		// range Start
		'data-[range-start]:border-[#3526ff] data-[range-start]:bg-[#3526ff] data-[range-start]:text-white data-[range-start]:hover:text-white',
		// range End
		'data-[range-end]:border-[#3526ff] data-[range-end]:bg-[#3526ff] data-[range-end]:text-white data-[range-end]:hover:text-white',
		// Outside months
		'[&[data-outside-month]:not([data-selected])]:text-zinc-700 [&[data-outside-month]:not([data-selected])]:hover:text-zinc-400',
		// Disabled
		'data-[disabled]:pointer-events-none data-[disabled]:text-zinc-600 data-[disabled]:opacity-50',
		// Unavailable
		'data-[unavailable]:line-through',
		'data-[range-middle]:bg-[#3526ff]/20 data-[range-middle]:text-zinc-100 dark:data-[range-middle]:hover:bg-[#3526ff]/20',
		// focus
		'focus:relative focus:border-zinc-300 focus:ring-2 focus:ring-zinc-300/30',
		// inner spans
		'[&>span]:text-xs [&>span]:opacity-70',
		className
	)}
	{...restProps}
/>
