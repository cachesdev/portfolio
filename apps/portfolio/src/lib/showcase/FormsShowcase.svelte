<script lang="ts">
	import ArrowClockwise from 'phosphor-svelte/lib/ArrowClockwise';
	import ArrowRight from 'phosphor-svelte/lib/ArrowRight';
	import Check from 'phosphor-svelte/lib/Check';
	import Command from 'phosphor-svelte/lib/Command';
	import EnvelopeSimple from 'phosphor-svelte/lib/EnvelopeSimple';
	import Sparkle from 'phosphor-svelte/lib/Sparkle';

	import { Badge } from '$lib/components/ui/badge/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as ButtonGroup from '$lib/components/ui/button-group/index.js';
	import { Checkbox } from '$lib/components/ui/checkbox/index.js';
	import * as Field from '$lib/components/ui/field/index.js';
	import * as Form from '$lib/components/ui/form/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as InputGroup from '$lib/components/ui/input-group/index.js';
	import * as InputOTP from '$lib/components/ui/input-otp/index.js';
	import * as Kbd from '$lib/components/ui/kbd/index.js';
	import { Label } from '$lib/components/ui/label/index.js';
	import * as NativeSelect from '$lib/components/ui/native-select/index.js';
	import { Progress } from '$lib/components/ui/progress/index.js';
	import * as RadioGroup from '$lib/components/ui/radio-group/index.js';
	import * as Select from '$lib/components/ui/select/index.js';
	import { Separator } from '$lib/components/ui/separator/index.js';
	import { Slider } from '$lib/components/ui/slider/index.js';
	import { Spinner } from '$lib/components/ui/spinner/index.js';
	import { Switch } from '$lib/components/ui/switch/index.js';
	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import { Toggle } from '$lib/components/ui/toggle/index.js';
	import * as ToggleGroup from '$lib/components/ui/toggle-group/index.js';

	const channels = [
		{ value: 'email', label: 'Email' },
		{ value: 'console', label: 'Console' },
		{ value: 'webhook', label: 'Webhook' }
	];

	const priorityOptions = [
		{ value: 'low', label: 'Low' },
		{ value: 'normal', label: 'Normal' },
		{ value: 'urgent', label: 'Urgent' }
	];

	let title = $state('Launch telemetry');
	let endpoint = $state('api.cluster.dev');
	let notes = $state('Track conversion spikes and deploy drift.');
	let remember = $state(true);
	let syncEnabled = $state(true);
	let channel = $state('email');
	let priority = $state('normal');
	let environment = $state('production');
	let format = $state('compact');
	let activeTool = $state('preview');
	let otp = $state('4821');
	let confidence = $state([72]);
	let submitted = $state(false);

	const selectedChannel = $derived(
		channels.find((item) => item.value === channel)?.label ?? 'Email'
	);
	const progress = $derived(Math.min(100, Math.max(0, confidence[0] ?? 0)));

	function submitDemo(event: SubmitEvent) {
		event.preventDefault();
		submitted = true;
	}

	function resetDemo() {
		title = 'Launch telemetry';
		endpoint = 'api.cluster.dev';
		notes = 'Track conversion spikes and deploy drift.';
		remember = true;
		syncEnabled = true;
		channel = 'email';
		priority = 'normal';
		environment = 'production';
		format = 'compact';
		activeTool = 'preview';
		otp = '4821';
		confidence = [72];
		submitted = false;
	}
</script>

<section class="grid gap-5 text-zinc-50">
	<div class="terminal-surface overflow-hidden">
		<div class="flex items-center justify-between border-b border-border bg-zinc-950/70 px-4 py-3">
			<div>
				<p class="terminal-label">Forms Stack</p>
				<h2 class="mt-1 text-xl font-semibold tracking-normal text-zinc-50">Signal Console</h2>
			</div>
			<div class="flex items-center gap-2">
				<Badge variant={submitted ? 'default' : 'secondary'}>
					{#if submitted}
						Sent
					{:else}
						Draft
					{/if}
				</Badge>
				<Spinner class={syncEnabled ? 'opacity-100' : 'opacity-30'} aria-label="Sync status" />
			</div>
		</div>

		<form class="p-4 sm:p-5" onsubmit={submitDemo}>
			<Field.Group class="gap-5">
				<Field.Set class="space-y-5">
					<Field.Legend>Request</Field.Legend>
					<div class="grid gap-4 md:grid-cols-2">
						<Field.Field>
							<Label for="signal-title">Title</Label>
							<Input
								id="signal-title"
								name="title"
								bind:value={title}
								placeholder="Brief request title"
								autocomplete="off"
							/>
						</Field.Field>

						<Field.Field>
							<Label for="signal-endpoint">Endpoint</Label>
							<InputGroup.InputGroup>
								<InputGroup.InputGroupAddon>
									<EnvelopeSimple aria-hidden />
								</InputGroup.InputGroupAddon>
								<InputGroup.InputGroupInput
									id="signal-endpoint"
									name="endpoint"
									bind:value={endpoint}
									placeholder="service.domain"
									autocomplete="off"
								/>
								<InputGroup.InputGroupAddon align="inline-end">
									<Kbd.Kbd>API</Kbd.Kbd>
								</InputGroup.InputGroupAddon>
							</InputGroup.InputGroup>
						</Field.Field>
					</div>

					<Field.Field>
						<Label for="signal-notes">Notes</Label>
						<Textarea id="signal-notes" name="notes" bind:value={notes} rows={3} />
					</Field.Field>
				</Field.Set>

				<Separator />

				<div class="grid gap-5 md:grid-cols-[1fr_0.95fr]">
					<Field.Set class="space-y-4">
						<Field.Legend>Routing</Field.Legend>
						<div class="grid gap-4 sm:grid-cols-2">
							<Field.Field>
								<Label id="channel-label">Channel</Label>
								<Select.Select type="single" bind:value={channel} name="channel">
									<Select.SelectTrigger
										class="w-full"
										aria-labelledby="channel-label"
										aria-label="Channel"
									>
										<span class="truncate">{selectedChannel}</span>
									</Select.SelectTrigger>
									<Select.SelectContent>
										{#each channels as item (item.value)}
											<Select.SelectItem value={item.value} label={item.label} />
										{/each}
									</Select.SelectContent>
								</Select.Select>
							</Field.Field>

							<Field.Field>
								<Label for="priority">Priority</Label>
								<NativeSelect.NativeSelect
									id="priority"
									name="priority"
									bind:value={priority}
									class="w-full"
								>
									{#each priorityOptions as item (item.value)}
										<NativeSelect.NativeSelectOption value={item.value}>
											{item.label}
										</NativeSelect.NativeSelectOption>
									{/each}
								</NativeSelect.NativeSelect>
							</Field.Field>
						</div>

						<Field.Field>
							<Label id="environment-label">Environment</Label>
							<RadioGroup.RadioGroup
								bind:value={environment}
								aria-labelledby="environment-label"
								class="grid gap-3 sm:grid-cols-3"
							>
								{#each ['staging', 'preview', 'production'] as option (option)}
									<div class="flex items-center gap-2">
										<RadioGroup.RadioGroupItem id={`env-${option}`} value={option} />
										<Label for={`env-${option}`}>{option}</Label>
									</div>
								{/each}
							</RadioGroup.RadioGroup>
						</Field.Field>
					</Field.Set>

					<Field.Set class="space-y-4">
						<Field.Legend>Security</Field.Legend>
						<Field.Field>
							<Label for="otp-code">OTP</Label>
							<InputOTP.InputOTP
								id="otp-code"
								maxlength={4}
								bind:value={otp}
								aria-label="One-time password"
							>
								{#snippet children({ cells })}
									<InputOTP.InputOTPGroup>
										{#each cells as cell, index (index)}
											<InputOTP.InputOTPSlot {cell} />
										{/each}
									</InputOTP.InputOTPGroup>
								{/snippet}
							</InputOTP.InputOTP>
						</Field.Field>

						<div class="grid gap-3">
							<Field.Field orientation="horizontal" class="justify-between gap-4">
								<Field.Content>
									<Field.Title>Remember device</Field.Title>
									<Field.Description>Skip this prompt next session.</Field.Description>
								</Field.Content>
								<Checkbox
									id="remember-device"
									bind:checked={remember}
									aria-label="Remember device"
								/>
							</Field.Field>

							<Field.Field orientation="horizontal" class="justify-between gap-4">
								<Field.Content>
									<Field.Title>Live sync</Field.Title>
									<Field.Description>Push updates as they arrive.</Field.Description>
								</Field.Content>
								<Switch id="live-sync" bind:checked={syncEnabled} aria-label="Live sync" />
							</Field.Field>
						</div>
					</Field.Set>
				</div>

				<Separator />

				<Field.Set class="space-y-5">
					<div class="flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
						<div>
							<Field.Legend>Output</Field.Legend>
							<p class="mt-2 text-sm text-zinc-500">Confidence {progress}%</p>
						</div>
						<div class="flex flex-wrap items-center gap-2">
							<Toggle bind:pressed={submitted} variant="outline" aria-label="Mark as submitted">
								<Check aria-hidden />
								Ready
							</Toggle>
							<ToggleGroup.ToggleGroup
								type="single"
								bind:value={format}
								variant="outline"
								size="sm"
								aria-label="Output format"
							>
								<ToggleGroup.ToggleGroupItem value="compact">Compact</ToggleGroup.ToggleGroupItem>
								<ToggleGroup.ToggleGroupItem value="verbose">Verbose</ToggleGroup.ToggleGroupItem>
							</ToggleGroup.ToggleGroup>
						</div>
					</div>

					<Field.Field>
						<Label for="confidence">Confidence</Label>
						<Slider
							id="confidence"
							type="multiple"
							bind:value={confidence}
							min={0}
							max={100}
							step={1}
							aria-label="Confidence"
						/>
						<Progress value={progress} aria-label="Confidence progress" />
					</Field.Field>

					<div class="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
						<ButtonGroup.ButtonGroup aria-label="Preview actions">
							<Button
								type="button"
								variant={activeTool === 'preview' ? 'default' : 'outline'}
								onclick={() => (activeTool = 'preview')}
							>
								<Sparkle aria-hidden />
								Preview
							</Button>
							<Button
								type="button"
								variant={activeTool === 'diff' ? 'default' : 'outline'}
								onclick={() => (activeTool = 'diff')}
							>
								Diff
							</Button>
							<ButtonGroup.ButtonGroupSeparator />
							<Button
								type="button"
								variant="outline"
								size="icon"
								aria-label="Reset"
								onclick={resetDemo}
							>
								<ArrowClockwise aria-hidden />
							</Button>
						</ButtonGroup.ButtonGroup>

						<div class="flex items-center gap-3">
							<Kbd.KbdGroup>
								<Kbd.Kbd><Command aria-hidden /></Kbd.Kbd>
								<Kbd.Kbd>Enter</Kbd.Kbd>
							</Kbd.KbdGroup>
							<Form.Button class="w-auto">
								Send
								<ArrowRight aria-hidden />
							</Form.Button>
						</div>
					</div>
				</Field.Set>
			</Field.Group>
		</form>
	</div>

	<aside class="hairline-frame terminal-grid-sm grid gap-6 p-5 lg:grid-cols-[minmax(0,1fr)_auto]">
		<div class="space-y-4 lg:min-w-80">
			<div class="flex items-center justify-between gap-3">
				<p class="terminal-label">Payload</p>
				<Badge variant="outline">{environment}</Badge>
			</div>

			<div class="space-y-3 font-mono text-sm">
				<div class="flex items-center justify-between gap-4 border-b border-border pb-3">
					<span class="text-zinc-500">title</span>
					<span class="truncate text-right text-zinc-100">{title}</span>
				</div>
				<div class="flex items-center justify-between gap-4 border-b border-border pb-3">
					<span class="text-zinc-500">route</span>
					<span class="truncate text-right text-zinc-100">{selectedChannel}</span>
				</div>
				<div class="flex items-center justify-between gap-4 border-b border-border pb-3">
					<span class="text-zinc-500">priority</span>
					<span class="text-right text-zinc-100">{priority}</span>
				</div>
				<div class="flex items-center justify-between gap-4 border-b border-border pb-3">
					<span class="text-zinc-500">format</span>
					<span class="text-right text-zinc-100">{format}</span>
				</div>
			</div>
		</div>

		<div class="space-y-4">
			<Separator />
			<div class="grid grid-cols-3 gap-2 text-center font-mono text-xs uppercase tracking-widest">
				<div class="hairline-frame-muted p-3">
					<p class="text-zinc-500">OTP</p>
					<p class="mt-2 text-zinc-50">{otp || '----'}</p>
				</div>
				<div class="hairline-frame-muted p-3">
					<p class="text-zinc-500">Sync</p>
					<p class="mt-2 text-zinc-50">{syncEnabled ? 'On' : 'Off'}</p>
				</div>
				<div class="hairline-frame-muted p-3">
					<p class="text-zinc-500">Tool</p>
					<p class="mt-2 text-zinc-50">{activeTool}</p>
				</div>
			</div>
		</div>
	</aside>
</section>
