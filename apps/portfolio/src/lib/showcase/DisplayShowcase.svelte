<script lang="ts">
	import { CalendarDate, type DateValue } from '@internationalized/date';
	import {
		ArrowRight,
		CalendarBlank,
		CheckCircle,
		CircleNotch,
		Database,
		Pulse
	} from 'phosphor-svelte';
	import * as Alert from '$lib/components/ui/alert/index.js';
	import { AspectRatio } from '$lib/components/ui/aspect-ratio/index.js';
	import * as Avatar from '$lib/components/ui/avatar/index.js';
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Card from '$lib/components/ui/card/index.js';
	import * as Carousel from '$lib/components/ui/carousel/index.js';
	import { ChartContainer, type ChartConfig } from '$lib/components/ui/chart/index.js';
	import * as Empty from '$lib/components/ui/empty/index.js';
	import * as Item from '$lib/components/ui/item/index.js';
	import { Skeleton } from '$lib/components/ui/skeleton/index.js';
	import * as Table from '$lib/components/ui/table/index.js';
	import { Calendar } from '$lib/components/ui/calendar/index.js';
	import { RangeCalendar } from '$lib/components/ui/range-calendar/index.js';

	const metrics = [
		{ id: 'latency', label: 'Latency', value: '82ms', trend: '-18%', tone: 'text-success' },
		{
			id: 'uptime',
			label: 'Uptime',
			value: '99.98%',
			trend: '+0.3%',
			tone: 'text-terminal-violet-soft'
		},
		{ id: 'queue', label: 'Queue', value: '12', trend: 'steady', tone: 'text-warning' }
	];

	const pipeline = [
		{ id: 'ingest', label: 'Ingest', owner: 'AV', status: 'Live', cost: '$84' },
		{ id: 'cache', label: 'Cache', owner: 'NR', status: 'Warm', cost: '$21' },
		{ id: 'search', label: 'Search', owner: 'KM', status: 'Indexing', cost: '$39' }
	];

	const releases = [
		{ id: 'edge', title: 'Edge preview', copy: 'Synthetic checks are green across all regions.' },
		{
			id: 'core',
			title: 'Core deploy',
			copy: 'Traffic shifted with no elevated error budget burn.'
		},
		{ id: 'audit', title: 'Audit trail', copy: 'Calendar windows and artifact hashes are pinned.' }
	];

	const chartConfig = {
		requests: { label: 'Requests', color: 'var(--terminal-violet-soft)' },
		errors: { label: 'Errors', color: 'var(--warning)' }
	} satisfies ChartConfig;

	const chartBars = [
		{ id: 'mon', day: 'M', requests: 42, errors: 7 },
		{ id: 'tue', day: 'T', requests: 58, errors: 5 },
		{ id: 'wed', day: 'W', requests: 50, errors: 9 },
		{ id: 'thu', day: 'T', requests: 71, errors: 4 },
		{ id: 'fri', day: 'F', requests: 64, errors: 6 },
		{ id: 'sat', day: 'S', requests: 78, errors: 3 }
	];

	let selectedMetric = $state(metrics[0].id);
	let selectedDate = $state<DateValue>(new CalendarDate(2026, 6, 18));
	let selectedRange = $state({
		start: new CalendarDate(2026, 6, 22),
		end: new CalendarDate(2026, 6, 26)
	});

	const activeMetric = $derived(
		metrics.find((metric) => metric.id === selectedMetric) ?? metrics[0]
	);
</script>

{#snippet calendarDay({ day, outsideMonth }: { day: DateValue; outsideMonth: boolean })}
	<span
		class={[
			'grid size-8 place-items-center text-xs',
			selectedDate.compare(day) === 0 &&
				!outsideMonth &&
				'outline outline-1 outline-terminal-violet-soft',
			outsideMonth && 'opacity-30'
		]}
	>
		{day.day}
	</span>
{/snippet}

{#snippet rangeDay({ day, outsideMonth }: { day: DateValue; outsideMonth: boolean })}
	<span
		class={[
			'grid size-8 place-items-center text-xs',
			!outsideMonth &&
				selectedRange.start.compare(day) <= 0 &&
				selectedRange.end.compare(day) >= 0 &&
				'bg-terminal-violet/25 text-foreground',
			outsideMonth && 'opacity-30'
		]}
	>
		{day.day}
	</span>
{/snippet}

<section class="relative grid gap-5 text-foreground" aria-labelledby="display-showcase-title">
	<div class="grid gap-5">
		<div
			class="flex flex-col gap-3 border border-border bg-black/80 p-4 sm:flex-row sm:items-end sm:justify-between"
		>
			<div class="grid gap-2">
				<p class="terminal-label">Display system</p>
				<h2 id="display-showcase-title" class="text-2xl font-semibold tracking-normal sm:text-3xl">
					Terminal-grade component previews
				</h2>
			</div>
			<div class="flex items-center gap-2 text-sm text-muted-foreground">
				<span class="size-2 bg-success" aria-hidden="true"></span>
				<span>12 components online</span>
			</div>
		</div>

		<Alert.Alert class="bg-terminal/90">
			<CheckCircle class="size-4" aria-hidden="true" />
			<Alert.AlertTitle>Release channel is stable</Alert.AlertTitle>
			<Alert.AlertDescription>
				Edge traffic is mirrored and the rollback window closes at 18:00 UTC.
			</Alert.AlertDescription>
			<Alert.AlertAction>
				<Button size="icon-xs" variant="outline" aria-label="Open release details">
					<ArrowRight aria-hidden="true" />
				</Button>
			</Alert.AlertAction>
		</Alert.Alert>

		<div class="grid gap-5 lg:grid-cols-[1.2fr_0.8fr]">
			<Card.Card class="bg-card/95" size="sm">
				<Card.CardHeader>
					<Card.CardTitle>Operations cockpit</Card.CardTitle>
					<Card.CardDescription
						>Concise cards, avatars, aspect media, and live status controls.</Card.CardDescription
					>
					<Card.CardAction>
						<Avatar.AvatarGroup>
							{#each pipeline as row (row.id)}
								<Avatar.Avatar size="sm" title={row.owner}>
									<Avatar.AvatarFallback>{row.owner}</Avatar.AvatarFallback>
								</Avatar.Avatar>
							{/each}
						</Avatar.AvatarGroup>
					</Card.CardAction>
				</Card.CardHeader>
				<Card.CardContent class="grid gap-4 lg:grid-cols-[0.9fr_1.1fr]">
					<AspectRatio ratio={16 / 10} class="terminal-grid-sm border-terminal-line bg-terminal">
						<div class="flex h-full flex-col justify-between p-4">
							<div class="flex items-center justify-between text-xs text-muted-foreground">
								<span class="font-mono uppercase">trace://primary</span>
								<span>{activeMetric.trend}</span>
							</div>
							<div>
								<p class="text-4xl font-semibold">{activeMetric.value}</p>
								<p class={['font-mono text-xs uppercase', activeMetric.tone]}>
									{activeMetric.label}
								</p>
							</div>
							<div class="grid grid-cols-12 items-end gap-1" aria-hidden="true">
								{#each { length: 12 }, index (index)}
									<div
										class="bg-terminal-violet-soft/70"
										style={`height: ${18 + ((index * 17) % 48)}px`}
									></div>
								{/each}
							</div>
						</div>
					</AspectRatio>

					<div class="grid content-start gap-3">
						<div class="grid grid-cols-3 gap-2">
							{#each metrics as metric (metric.id)}
								<button
									type="button"
									class={[
										'border border-border bg-black px-3 py-2 text-left transition-colors hover:border-terminal-violet-soft',
										selectedMetric === metric.id && 'border-terminal-violet-soft bg-terminal'
									]}
									aria-pressed={selectedMetric === metric.id}
									onclick={() => (selectedMetric = metric.id)}
								>
									<span class="block font-mono text-[0.65rem] text-muted-foreground uppercase">
										{metric.label}
									</span>
									<span class="text-lg font-semibold">{metric.value}</span>
								</button>
							{/each}
						</div>

						<Item.ItemGroup class="grid gap-2">
							{#each pipeline as row (row.id)}
								<Item.Item variant="outline" size="sm">
									<Item.ItemMedia variant="icon">
										<Database aria-hidden="true" />
									</Item.ItemMedia>
									<Item.ItemContent>
										<Item.ItemTitle>{row.label}</Item.ItemTitle>
										<Item.ItemDescription>{row.status} by {row.owner}</Item.ItemDescription>
									</Item.ItemContent>
									<Item.ItemActions class="font-mono text-xs text-muted-foreground">
										{row.cost}
									</Item.ItemActions>
								</Item.Item>
							{/each}
						</Item.ItemGroup>
					</div>
				</Card.CardContent>
				<Card.CardFooter class="gap-3">
					<Skeleton class="h-2 flex-1 bg-terminal-violet-soft/25" />
					<Skeleton class="h-2 w-20 bg-terminal-glow/25" />
				</Card.CardFooter>
			</Card.Card>

			<Card.Card class="bg-card/95" size="sm">
				<Card.CardHeader>
					<Card.CardTitle>Release carousel</Card.CardTitle>
					<Card.CardDescription
						>Keyboard and pointer navigation over compact status slides.</Card.CardDescription
					>
				</Card.CardHeader>
				<Card.CardContent>
					<Carousel.Carousel class="bg-terminal">
						<Carousel.CarouselContent>
							{#each releases as release (release.id)}
								<Carousel.CarouselItem>
									<div class="grid min-h-44 content-between gap-5 p-5">
										<div class="flex items-center justify-between">
											<span class="terminal-label">{release.id}</span>
											<Pulse class="size-4 text-terminal-violet-soft" aria-hidden="true" />
										</div>
										<div class="grid gap-2">
											<h3 class="text-xl font-semibold">{release.title}</h3>
											<p class="max-w-sm text-sm text-muted-foreground">{release.copy}</p>
										</div>
									</div>
								</Carousel.CarouselItem>
							{/each}
						</Carousel.CarouselContent>
						<Carousel.CarouselPrevious class="left-3" />
						<Carousel.CarouselNext class="right-3" />
					</Carousel.Carousel>
				</Card.CardContent>
			</Card.Card>
		</div>

		<div class="grid gap-5 xl:grid-cols-[0.9fr_1.1fr]">
			<Card.Card class="bg-card/95" size="sm">
				<Card.CardHeader>
					<Card.CardTitle>Load shape</Card.CardTitle>
					<Card.CardDescription>Chart container with terminal-series styling.</Card.CardDescription>
				</Card.CardHeader>
				<Card.CardContent>
					<ChartContainer
						config={chartConfig}
						class="terminal-grid-sm aspect-[4/3]"
						role="img"
						aria-label="Requests and errors by day"
					>
						<div class="flex h-full w-full flex-col justify-between gap-4">
							<div class="flex items-center justify-between text-xs text-muted-foreground">
								<span>Requests</span>
								<span>Errors</span>
							</div>
							<div class="grid flex-1 grid-cols-6 items-end gap-3">
								{#each chartBars as bar (bar.id)}
									<div class="grid gap-2">
										<div class="flex h-40 items-end gap-1">
											<div
												class="w-full bg-terminal-violet-soft"
												style={`height: ${bar.requests}%;`}
												aria-label={`${bar.day} requests ${bar.requests}`}
											></div>
											<div
												class="w-full bg-warning"
												style={`height: ${bar.errors * 5}%;`}
												aria-label={`${bar.day} errors ${bar.errors}`}
											></div>
										</div>
										<span class="text-center text-xs text-muted-foreground">{bar.day}</span>
									</div>
								{/each}
							</div>
						</div>
					</ChartContainer>
				</Card.CardContent>
			</Card.Card>

			<Card.Card class="bg-card/95" size="sm">
				<Card.CardHeader>
					<Card.CardTitle>Pipeline ledger</Card.CardTitle>
					<Card.CardDescription>Responsive table with terse operational rows.</Card.CardDescription>
				</Card.CardHeader>
				<Card.CardContent>
					<Table.Table>
						<Table.TableCaption>Current deploy lane ownership.</Table.TableCaption>
						<Table.TableHeader>
							<Table.TableRow>
								<Table.TableHead>Service</Table.TableHead>
								<Table.TableHead>Owner</Table.TableHead>
								<Table.TableHead>Status</Table.TableHead>
								<Table.TableHead class="text-right">Spend</Table.TableHead>
							</Table.TableRow>
						</Table.TableHeader>
						<Table.TableBody>
							{#each pipeline as row (row.id)}
								<Table.TableRow>
									<Table.TableCell class="font-medium text-foreground">{row.label}</Table.TableCell>
									<Table.TableCell>{row.owner}</Table.TableCell>
									<Table.TableCell>{row.status}</Table.TableCell>
									<Table.TableCell class="text-right font-mono">{row.cost}</Table.TableCell>
								</Table.TableRow>
							{/each}
						</Table.TableBody>
						<Table.TableFooter>
							<Table.TableRow>
								<Table.TableCell colspan={3}>Projected hourly total</Table.TableCell>
								<Table.TableCell class="text-right font-mono">$144</Table.TableCell>
							</Table.TableRow>
						</Table.TableFooter>
					</Table.Table>
				</Card.CardContent>
			</Card.Card>
		</div>

		<div class="grid gap-5 xl:grid-cols-[1fr_1fr_0.85fr]">
			<Card.Card class="bg-card/95" size="sm">
				<Card.CardHeader>
					<Card.CardTitle>Single window</Card.CardTitle>
					<Card.CardDescription>Calendar selection for the release checkpoint.</Card.CardDescription
					>
				</Card.CardHeader>
				<Card.CardContent>
					<Calendar
						type="single"
						bind:value={selectedDate}
						placeholder={new CalendarDate(2026, 6, 1)}
						fixedWeeks
						buttonVariant="outline"
						day={calendarDay}
						class="w-full"
					/>
				</Card.CardContent>
			</Card.Card>

			<Card.Card class="bg-card/95" size="sm">
				<Card.CardHeader>
					<Card.CardTitle>Maintenance range</Card.CardTitle>
					<Card.CardDescription
						>Range calendar for freeze and migration windows.</Card.CardDescription
					>
				</Card.CardHeader>
				<Card.CardContent>
					<RangeCalendar
						bind:value={selectedRange}
						placeholder={new CalendarDate(2026, 6, 1)}
						fixedWeeks
						buttonVariant="outline"
						day={rangeDay}
						class="w-full"
					/>
				</Card.CardContent>
			</Card.Card>

			<Empty.Empty class="min-h-80 bg-black/90">
				<Empty.EmptyMedia variant="icon">
					<CalendarBlank aria-hidden="true" />
				</Empty.EmptyMedia>
				<Empty.EmptyHeader>
					<Empty.EmptyTitle>No incidents queued</Empty.EmptyTitle>
					<Empty.EmptyDescription>On-call can keep the deploy calendar open.</Empty.EmptyDescription
					>
				</Empty.EmptyHeader>
				<Empty.EmptyContent>
					<Button variant="outline" size="sm">
						<CircleNotch data-icon="inline-start" aria-hidden="true" />
						Refresh
					</Button>
				</Empty.EmptyContent>
			</Empty.Empty>
		</div>
	</div>
</section>
