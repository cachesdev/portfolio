<script lang="ts" module>
	import type { HTMLAnchorAttributes, HTMLButtonAttributes } from 'svelte/elements';

	export type ButtonProps = HTMLButtonAttributes &
		HTMLAnchorAttributes & {
			href?: string;
		};
</script>

<script lang="ts">
	let {
		href = undefined,
		type = 'button',
		disabled,
		children,
		...restProps
	}: ButtonProps = $props();
</script>

{#if href}
	<a
		data-slot="button"
		href={disabled ? undefined : href}
		aria-disabled={disabled}
		role={disabled ? 'link' : undefined}
		tabindex={disabled ? -1 : undefined}
		{...restProps}
	>
		{@render children?.()}
	</a>
{:else}
	<button data-slot="button" {type} {disabled} {...restProps}>
		{@render children?.()}
	</button>
{/if}
