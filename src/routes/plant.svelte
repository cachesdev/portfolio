<script lang="ts">
	import * as THREE from 'three';

	import type { Snippet } from 'svelte';
	import { T, type Props } from '@threlte/core';
	import { useGltf, useDraco, TransformControls, interactivity } from '@threlte/extras';

	interactivity();

	let {
		fallback,
		error,
		children,
		ref = $bindable(),
		...props
	}: Props<THREE.Group> & {
		ref?: THREE.Group<THREE.Object3DEventMap> | undefined;
		children?: Snippet<[{ ref: THREE.Group<THREE.Object3DEventMap> | undefined }]>;
		fallback?: Snippet;
		error?: Snippet<[{ error: Error }]>;
	} = $props();

	type GLTFResult = {
		nodes: {
			Plant_Flower_pot__0: THREE.Mesh;
			Plant_Material001_0: THREE.Mesh;
			Plant_Land_0: THREE.Mesh;
		};
		materials: {
			Flower_pot: THREE.MeshStandardMaterial;
			['Material.001']: THREE.MeshStandardMaterial;
			Land: THREE.MeshStandardMaterial;
		};
	};

	const styles = getComputedStyle(document.documentElement);
	let color = styles.getPropertyValue('--color-foreground');

	const wireframe = new THREE.MeshBasicMaterial({ wireframe: true, color });

	const gltf = useGltf<GLTFResult>('/plant-transformed.glb', { dracoLoader: useDraco() });

	let controlsState = $state({
		visible: false,
		moving: false
	});
</script>

{#if controlsState.visible || controlsState.moving}
	<TransformControls object={ref} />
{/if}
<T.Group
	bind:ref
	dispose={false}
	{...props as any}
	onpointerenter={() => {
		controlsState.visible = true;
	}}
	onpointerleave={() => {
		if (!controlsState.moving) {
			controlsState.visible = false;
		}
	}}
	onpointerdown={() => {
		controlsState.moving = true;
	}}
	onpointerup={() => {
		controlsState.moving = false;
	}}
>
	{#await gltf}
		{@render fallback?.()}
	{:then gltf}
		<T.Group scale={2}>
			<T.Group rotation={[-Math.PI / 2, 0, 0]} scale={100}>
				<T.Mesh geometry={gltf.nodes.Plant_Flower_pot__0.geometry} material={wireframe} />
				<T.Mesh geometry={gltf.nodes.Plant_Material001_0.geometry} material={wireframe} />
				<T.Mesh geometry={gltf.nodes.Plant_Land_0.geometry} material={wireframe} />
			</T.Group>
		</T.Group>
	{:catch err}
		{@render error?.({ error: err })}
	{/await}

	{@render children?.({ ref })}
</T.Group>
