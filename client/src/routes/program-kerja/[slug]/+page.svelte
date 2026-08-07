<script lang="ts">
	import { onMount } from 'svelte';
	import { gsap } from 'gsap';

	import TimelinePelaksanaan from '$lib/components/program-kerja/TimelinePelaksanaan.svelte';
	import Documentation from '$lib/components/program-kerja/Documentation.svelte';

	import type { PageData } from './$types';

	export let data: PageData;

	let pageElement: HTMLElement;
	let titleElement: HTMLHeadingElement;
	let descriptionElement: HTMLParagraphElement;
	let timelineElement: HTMLDivElement;
	let documentationElement: HTMLDivElement;

	const monthNames = [
		'Januari',
		'Februari',
		'Maret',
		'April',
		'Mei',
		'Juni',
		'Juli',
		'Agustus',
		'September',
		'Oktober',
		'November',
		'Desember'
	];

	function formatProgramDate(value?: string | null): string {
		if (!value) {
			return '';
		}

		const match = value.match(/^(\d{4})-(\d{2})-(\d{2})/);

		if (!match) {
			return value;
		}

		const month = Number(match[2]);
		const day = Number(match[3]);

		if (
			Number.isNaN(day) ||
			Number.isNaN(month) ||
			month < 1 ||
			month > 12
		) {
			return value;
		}

		return `${day} ${monthNames[month - 1]}`;
	}

	$: program = data.program;

	$: titleLines = (() => {
		const words = program.name.trim().split(/\s+/);

		if (words.length <= 2) {
			return {
				first: program.name,
				second: ''
			};
		}

		const midpoint = Math.ceil(words.length / 2);

		return {
			first: words.slice(0, midpoint).join(' '),
			second: words.slice(midpoint).join(' ')
		};
	})();

	$: documentationImages = [...(program.documentations ?? [])]
		.sort((a, b) => a.display_order - b.display_order)
		.map((documentation) => documentation.media?.url)
		.filter((url): url is string => Boolean(url));

	$: startDate = formatProgramDate(program.start_date);
	$: endDate = formatProgramDate(program.end_date);

	onMount(() => {
		const reduceMotion = window.matchMedia(
			'(prefers-reduced-motion: reduce)'
		).matches;

		if (reduceMotion) {
			return;
		}

		const context = gsap.context(() => {
			const timeline = gsap.timeline({
				defaults: {
					ease: 'power3.out'
				}
			});

			if (titleElement) {
				timeline.from(titleElement, {
					opacity: 0,
					y: 50,
					duration: 0.9
				});
			}

			if (descriptionElement) {
				timeline.from(
					descriptionElement,
					{
						opacity: 0,
						y: 24,
						duration: 0.7
					},
					'-=0.45'
				);
			}

			if (timelineElement) {
				timeline.from(
					timelineElement,
					{
						opacity: 0,
						y: 35,
						duration: 0.8
					},
					'-=0.35'
				);
			}

			if (documentationElement) {
				timeline.from(
					documentationElement,
					{
						opacity: 0,
						y: 35,
						duration: 0.8
					},
					'-=0.35'
				);
			}
		}, pageElement);

		return () => {
			context.revert();
		};
	});
</script>

<svelte:head>
	<title>{program.name} | BEM UNAIR 2026</title>

	<meta
		name="description"
		content={program.short_description ??
			program.description ??
			`Program kerja ${program.name}`}
	/>
</svelte:head>

<main
	bind:this={pageElement}
	class="
		min-h-screen w-full overflow-hidden
		bg-linear-to-b
		from-blue-50 via-white to-blue-50
	"
>
	<section
		class="
			mx-auto flex min-h-screen
			w-full max-w-7xl
			flex-col items-center
			px-6 py-28
			text-center
			sm:px-8
			md:py-32
			lg:px-12
		"
	>
		<h1
			bind:this={titleElement}
			class="
				max-w-6xl
				text-[clamp(3rem,8vw,7.5rem)]
				font-black uppercase
				leading-[0.85]
				tracking-[-0.06em]
			"
		>
			<span
				class="
					block
					text-blue-400
					[-webkit-text-stroke:2px_white]
					[paint-order:stroke_fill]
					drop-shadow-[0_10px_0_rgba(5,34,70,0.78)]
					sm:[-webkit-text-stroke:3px_white]
				"
			>
				{titleLines.first}
			</span>

			{#if titleLines.second}
				<span
					class="
						relative left-[0.55em]
						block
						text-blue-600
						[-webkit-text-stroke:2px_white]
						[paint-order:stroke_fill]
						drop-shadow-[0_10px_0_rgba(5,34,70,0.78)]
						sm:[-webkit-text-stroke:3px_white]
					"
				>
					{titleLines.second}
				</span>
			{/if}
		</h1>

		{#if program.description}
			<p
				bind:this={descriptionElement}
				class="
					mt-14 max-w-4xl
					text-base font-semibold
					leading-relaxed
					text-blue-950/70
					sm:text-lg
				"
			>
				{program.description}
			</p>
		{/if}

		{#if startDate && endDate}
			<div
				bind:this={timelineElement}
				class="mt-16 w-full"
			>
				<TimelinePelaksanaan
					{startDate}
					{endDate}
				/>
			</div>
		{/if}

		{#if documentationImages.length > 0}
			<div
				bind:this={documentationElement}
				class="mt-20 w-full"
			>
				<Documentation images={documentationImages} />
			</div>
		{/if}
	</section>
</main>