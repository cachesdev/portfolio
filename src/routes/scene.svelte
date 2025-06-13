<script lang="ts">
	import { T, useTask } from '@threlte/core';
	import { Spring } from 'svelte/motion';
	import { Grid, Stars } from '@threlte/extras';
	import { mode, toggleMode } from 'mode-watcher';
	import Portal from '$lib/tweakpane/portal.svelte';
	import { Button, Pane, Slider } from 'svelte-tweakpane-ui';

	let scale = new Spring(0.2);

	let rotation = $state<[number, number, number]>([0, 0, 0]);
	useTask((delta) => {
		rotation[0] += delta / 95;
		rotation[1] += delta / 100;
		rotation[2] += delta / 80;
	});

	type Vec2 = {
		x: number;
		z: number;
	};

	let plantCount = $state(20);

	// Poisson disk sampling for even distribution
	const minDistance = 2; // Minimum distance between points
	const maxAttempts = 30;
	const areaSize = 80;

	function generatePoints(count: number): Vec2[] {
		const newPoints: Vec2[] = [];

		// First point
		newPoints.push({
			x: (Math.random() - 0.5) * areaSize,
			z: (Math.random() - 0.5) * areaSize
		});

		// Generate remaining points
		while (newPoints.length < count) {
			let placed = false;

			for (let attempt = 0; attempt < maxAttempts && !placed; attempt++) {
				const candidate = {
					x: (Math.random() - 0.5) * areaSize,
					z: (Math.random() - 0.5) * areaSize
				};

				// Check if far enough from existing points
				const tooClose = newPoints.some(
					(p) => Math.sqrt((p.x - candidate.x) ** 2 + (p.z - candidate.z) ** 2) < minDistance
				);

				if (!tooClose) {
					newPoints.push(candidate);
					placed = true;
				}
			}

			if (!placed) break;
		}

		return newPoints;
	}

	let points = $derived(generatePoints(plantCount));
</script>

{#key mode.current}
	<Stars {rotation} factor={8} lightness={mode.current === 'dark' ? 0.8 : -100} />
{/key}
<Portal>
	<Pane title="Tweakpane">
		<Slider label="Scale" bind:value={scale.target} min={0} max={10} />
		<Slider label="Plant Count" bind:value={plantCount} min={1} max={50} step={1} />
		<Button label="Toggle theme" on:click={toggleMode} />
	</Pane>
</Portal>

<!-- {#each points as point (point.x + point.z)}
	<Plant position.x={point.x} position.z={point.z} rotation.y={rotation} scale={scale.current} />
{/each} -->

<Grid sectionThickness={0} infiniteGrid cellColor="#dddddd" cellSize={2} />

<T.OrthographicCamera
	makeDefault
	zoom={20}
	near={-100}
	position={[10, 10, 10]}
	oncreate={(ref) => {
		ref.lookAt(0, 1, 0);
	}}
>
	<!-- <OrbitControls enableDamping enablePan={false} enableZoom /> -->
</T.OrthographicCamera>
