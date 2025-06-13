<script lang="ts">
	import { ModeWatcher } from 'mode-watcher';
	import '../app.css';
	import { ApplicationState, SidebarState } from './state.svelte';
	import { setContext } from 'svelte';
	import { page } from '$app/state';
	import { Canvas } from '@threlte/core';
	import Scene from './scene.svelte';
	import PortalTarget from '$lib/tweakpane/portal-target.svelte';
	import SidePanel from './side-panel.svelte';

	let { children } = $props();

	const application = new ApplicationState(
		new SidebarState({
			items: [
				{
					title: 'Blog',
					active: false,
					href: '/blog'
				},
				{
					title: 'Projects',
					active: false,
					href: '/projects'
				},
				{
					title: 'About',
					active: false,
					href: '/about'
				}
			]
		})
	);

	function handleNavigation(pathname: string) {
		if (pathname === '') {
			return;
		}

		// split into ["", "mypath", ...]
		const path = pathname.split('/');
		console.log(path);
		application.sidebar.setActive(`/${path[1]}`);
	}

	$effect(() => {
		handleNavigation(page.url.pathname);
	});

	setContext('application', application);
</script>

<ModeWatcher />
<PortalTarget />

<div class="bg-background text-foreground flex h-dvh w-dvw">
	<SidePanel />
	{@render children()}
	<Canvas>
		<Scene />
	</Canvas>
</div>
