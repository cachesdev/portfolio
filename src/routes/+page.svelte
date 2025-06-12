<script lang="ts">
	import { Canvas } from '@threlte/core';
	import Scene from './scene.svelte';
	import PortalTarget from '$lib/tweakpane/portal-target.svelte';
	import SidePanel from './side-panel.svelte';
</script>

<div class="bg-background text-foreground flex h-dvh w-dvw">
	<PortalTarget />
	<SidePanel />
	<Canvas>
		<Scene />
	</Canvas>
</div>
