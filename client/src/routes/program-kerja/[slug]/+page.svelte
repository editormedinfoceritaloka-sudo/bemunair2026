<script lang="ts">
	import { onMount } from 'svelte';

	import TimelinePelaksanaan from '$lib/components/program-kerja/TimelinePelaksanaan.svelte';
	import Documentation from '$lib/components/program-kerja/Documentation.svelte';

	import type { PageData } from './$types';

	let {
		data
	}: { data: PageData } = $props();

	let pageElement: HTMLElement;
	let titleElement: HTMLHeadingElement;
	let descriptionElement = $state<HTMLParagraphElement>();
	let timelineElement = $state<HTMLDivElement>();
	let documentationElement = $state<HTMLDivElement>();

	const DAY_MS = 24 * 60 * 60 * 1000;

	const shortMonthNames = [
		'Jan',
		'Feb',
		'Mar',
		'Apr',
		'Mei',
		'Jun',
		'Jul',
		'Agu',
		'Sep',
		'Okt',
		'Nov',
		'Des'
	];

	let now = $state(Date.now());
	let mounted = $state(false);

	let introActive = $state(false);
	let introLeaving = $state(false);
	let introStyle = $state('');
	let timelineReady = $state(false);

	function splitTitleLines(name: string): { first: string; second: string } {
		const words = name.trim().split(/\s+/);

		if (words.length <= 2) {
			return { first: name, second: '' };
		}

		const midpoint = Math.ceil(words.length / 2);

		return {
			first: words.slice(0, midpoint).join(' '),
			second: words.slice(midpoint).join(' ')
		};
	}

	function parseLocalDate(value: string): Date | null {
		const [year, month, day] = value.split('-').map(Number);

		if (!year || !month || !day) {
			return null;
		}

		return new Date(year, month - 1, day);
	}

	function formatShortDate(start: Date, end: Date | null): string {
		if (!end || start.getTime() === end.getTime()) {
			return `${start.getDate()} ${shortMonthNames[start.getMonth()]} ${start.getFullYear()}`;
		}

		if (start.getFullYear() !== end.getFullYear()) {
			return `${start.getDate()} ${shortMonthNames[start.getMonth()]} ${start.getFullYear()} – ${end.getDate()} ${shortMonthNames[end.getMonth()]} ${end.getFullYear()}`;
		}

		if (start.getMonth() !== end.getMonth()) {
			return `${start.getDate()} ${shortMonthNames[start.getMonth()]} – ${end.getDate()} ${shortMonthNames[end.getMonth()]} ${end.getFullYear()}`;
		}

		return `${start.getDate()}–${end.getDate()} ${shortMonthNames[start.getMonth()]} ${start.getFullYear()}`;
	}

	function pad(value: number): string {
		return String(value).padStart(2, '0');
	}

	const program = $derived(data.program);
	const titleLines = $derived(splitTitleLines(program.name));
	const documentationImages = $derived(
		[...(program.documentations ?? [])]
			.sort((a, b) => Number(b.is_cover) - Number(a.is_cover) || a.display_order - b.display_order)
			.map((documentation) => documentation.media?.url)
			.filter((url): url is string => Boolean(url))
	);
	const descriptionParts = $derived.by(() => {
		const description = program.description ?? '';
		const name = program.name.trim();

		if (!name) {
			return [{ text: description, highlighted: false }];
		}

		const lowerDescription = description.toLocaleLowerCase('id-ID');
		const lowerName = name.toLocaleLowerCase('id-ID');
		const parts: { text: string; highlighted: boolean }[] = [];
		let currentIndex = 0;

		while (currentIndex < description.length) {
			const matchIndex = lowerDescription.indexOf(lowerName, currentIndex);

			if (matchIndex === -1) {
				parts.push({ text: description.slice(currentIndex), highlighted: false });
				break;
			}

			if (matchIndex > currentIndex) {
				parts.push({ text: description.slice(currentIndex, matchIndex), highlighted: false });
			}

			parts.push({
				text: description.slice(matchIndex, matchIndex + name.length),
				highlighted: true
			});

			currentIndex = matchIndex + name.length;
		}

		return parts;
	});
	const startDate = $derived(program.start_date?.slice(0, 10) ?? '');
	const endDate = $derived(program.end_date?.slice(0, 10) ?? '');

	const schedule = $derived.by(() => {
		const start = startDate ? parseLocalDate(startDate) : null;

		if (!start) {
			return null;
		}

		const end = (endDate ? parseLocalDate(endDate) : null) ?? start;
		const startMs = start.getTime();
		const endMs = end.getTime() + DAY_MS;
		const totalDays = Math.max(1, Math.round((endMs - startMs) / DAY_MS));

		return {
			startMs,
			endMs,
			totalDays,
			label: formatShortDate(start, endDate ? end : null)
		};
	});

	const phase = $derived.by(() => {
		if (!schedule) {
			return null;
		}

		if (now < schedule.startMs) {
			return 'upcoming' as const;
		}

		if (now < schedule.endMs) {
			return 'ongoing' as const;
		}

		return 'done' as const;
	});

	const statusLabel = $derived(
		phase === 'upcoming' ? 'Akan Datang' : phase === 'ongoing' ? 'Berlangsung' : 'Selesai'
	);

	const countdown = $derived.by(() => {
		if (!schedule || phase === 'done' || !phase) {
			return null;
		}

		const target = phase === 'upcoming' ? schedule.startMs : schedule.endMs;
		const remaining = Math.max(0, target - now);
		const totalSeconds = Math.floor(remaining / 1000);

		return [
			{ value: pad(Math.floor(totalSeconds / 86400)), label: 'HARI' },
			{ value: pad(Math.floor((totalSeconds % 86400) / 3600)), label: 'JAM' },
			{ value: pad(Math.floor((totalSeconds % 3600) / 60)), label: 'MENIT' },
			{ value: pad(totalSeconds % 60), label: 'DETIK' }
		];
	});

	const progress = $derived.by(() => {
		if (!schedule || phase !== 'ongoing') {
			return null;
		}

		const elapsed = now - schedule.startMs;
		const day = Math.min(schedule.totalDays, Math.floor(elapsed / DAY_MS) + 1);
		const percent = Math.min(100, Math.max(0, Math.round((elapsed / (schedule.endMs - schedule.startMs)) * 100)));

		return { day, percent };
	});

	function portal(node: HTMLElement) {
		document.body.appendChild(node);

		return {
			destroy() {
				node.remove();
			}
		};
	}

	function playIntroSound(): void {
		try {
			const AudioContextClass =
				window.AudioContext ??
				(window as unknown as { webkitAudioContext: typeof AudioContext }).webkitAudioContext;

			if (!AudioContextClass) {
				return;
			}

			const context = new AudioContextClass();

			void context.resume().then(() => {
				if (context.state !== 'running') {
					void context.close();
					return;
				}

				const master = context.createGain();
				master.gain.value = 0.5;
				master.connect(context.destination);

				const start = context.currentTime + 0.02;

				function click(at: number): void {
					const length = Math.floor(context.sampleRate * 0.06);
					const buffer = context.createBuffer(1, length, context.sampleRate);
					const channel = buffer.getChannelData(0);

					for (let index = 0; index < length; index += 1) {
						channel[index] = (Math.random() * 2 - 1) * (1 - index / length) ** 3;
					}

					const noise = context.createBufferSource();
					const filter = context.createBiquadFilter();
					const gain = context.createGain();

					noise.buffer = buffer;
					filter.type = 'highpass';
					filter.frequency.value = 1800;
					gain.gain.value = 0.6;
					noise.connect(filter).connect(gain).connect(master);
					noise.start(at);

					const thump = context.createOscillator();
					const thumpGain = context.createGain();

					thump.frequency.setValueAtTime(160, at);
					thump.frequency.exponentialRampToValueAtTime(60, at + 0.12);
					thumpGain.gain.setValueAtTime(0.5, at);
					thumpGain.gain.exponentialRampToValueAtTime(0.0001, at + 0.14);
					thump.connect(thumpGain).connect(master);
					thump.start(at);
					thump.stop(at + 0.16);
				}

				click(start + 0.2);
				click(start + 0.34);

				const swellLength = Math.floor(context.sampleRate * 1.9);
				const swellBuffer = context.createBuffer(1, swellLength, context.sampleRate);
				const swellChannel = swellBuffer.getChannelData(0);

				for (let index = 0; index < swellLength; index += 1) {
					swellChannel[index] = Math.random() * 2 - 1;
				}

				const swell = context.createBufferSource();
				const swellFilter = context.createBiquadFilter();
				const swellGain = context.createGain();

				swell.buffer = swellBuffer;
				swellFilter.type = 'bandpass';
				swellFilter.Q.value = 1.2;
				swellFilter.frequency.setValueAtTime(220, start + 0.3);
				swellFilter.frequency.exponentialRampToValueAtTime(2400, start + 1.95);
				swellGain.gain.setValueAtTime(0.0001, start + 0.3);
				swellGain.gain.exponentialRampToValueAtTime(0.22, start + 1.85);
				swellGain.gain.exponentialRampToValueAtTime(0.0001, start + 2.1);
				swell.connect(swellFilter).connect(swellGain).connect(master);
				swell.start(start + 0.3);

				[196, 294, 392].forEach((frequency) => {
					const pad = context.createOscillator();
					const padGain = context.createGain();

					pad.type = 'sine';
					pad.frequency.value = frequency;
					padGain.gain.setValueAtTime(0.0001, start + 0.4);
					padGain.gain.exponentialRampToValueAtTime(0.05, start + 1.6);
					padGain.gain.exponentialRampToValueAtTime(0.0001, start + 2.4);
					pad.connect(padGain).connect(master);
					pad.start(start + 0.4);
					pad.stop(start + 2.5);
				});

				[784, 988, 1175, 1568].forEach((frequency, index) => {
					const at = start + 1.95 + index * 0.07;
					const tone = context.createOscillator();
					const gain = context.createGain();

					tone.type = 'triangle';
					tone.frequency.value = frequency;
					gain.gain.setValueAtTime(0.0001, at);
					gain.gain.exponentialRampToValueAtTime(0.06, at + 0.02);
					gain.gain.exponentialRampToValueAtTime(0.0001, at + 0.5);
					tone.connect(gain).connect(master);
					tone.start(at);
					tone.stop(at + 0.55);
				});

				setTimeout(() => {
					const shine = new Audio('/audio/shine.mp3');
					shine.volume = 0.55;
					void shine.play().catch(() => {});
				}, 1900);

				setTimeout(() => {
					void context.close();
				}, 3500);
			});
		} catch {
			// browser menolak audio, animasi tetap jalan
		}
	}

	function startIntro(): (() => void) | undefined {
		if (!titleElement) {
			return;
		}

		const rect = titleElement.getBoundingClientRect();
		const width = window.innerWidth;
		const height = window.innerHeight;

		const sx = rect.left + rect.width / 2;
		const sy = Math.min(Math.max(rect.top + rect.height / 2, 80), height - 80);
		const lampY = -40;
		const leftLength = Math.hypot(sx, sy - lampY);
		const rightLength = Math.hypot(width - sx, sy - lampY);
		const leftAngle = (Math.atan2(sx, sy - lampY) * 180) / Math.PI;
		const rightAngle = (Math.atan2(width - sx, sy - lampY) * 180) / Math.PI;
		const spot = Math.max(110, Math.min(170, rect.height * 0.9));

		introStyle = [
			`--sx:${sx}px`,
			`--sy:${sy}px`,
			`--spot:${spot}px`,
			`--len-l:${leftLength + 60}px`,
			`--len-r:${rightLength + 60}px`,
			`--angle-l:${-leftAngle}deg`,
			`--angle-r:${rightAngle}deg`
		].join(';');

		introActive = true;
		playIntroSound();

		let finished = false;

		const finish = (fast: boolean) => {
			if (finished) {
				return;
			}

			finished = true;
			removeListeners();

			if (fast) {
				introLeaving = true;
				setTimeout(() => {
					introActive = false;
					introLeaving = false;
					timelineReady = true;
				}, 320);
			} else {
				introActive = false;
				timelineReady = true;
			}
		};

		const startedAt = performance.now();
		const skip = () => {
			if (performance.now() - startedAt > 500) {
				finish(true);
			}
		};
		const events = ['pointerdown', 'keydown'] as const;

		function removeListeners(): void {
			events.forEach((name) => window.removeEventListener(name, skip));
		}

		events.forEach((name) => window.addEventListener(name, skip, { passive: true }));

		const timer = setTimeout(() => finish(false), 2700);

		return () => {
			clearTimeout(timer);
			removeListeners();
		};
	}

	onMount(() => {
		mounted = true;
		now = Date.now();

		const clock = setInterval(() => {
			now = Date.now();
		}, 1000);

		const reduceMotion = window.matchMedia(
			'(prefers-reduced-motion: reduce)'
		).matches;

		if (reduceMotion) {
			timelineReady = true;
			return () => clearInterval(clock);
		}

		const stopIntro = startIntro();
		const introPlaying = introActive;

		if (!introPlaying) {
			timelineReady = true;
		}

		let cleanup: (() => void) | undefined;

		import('gsap').then(({ gsap }) => {
			const context = gsap.context(() => {
				const timeline = gsap.timeline({
					defaults: { ease: 'power3.out' },
					delay: introPlaying ? 0.5 : 0
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
						{ opacity: 0, y: 24, duration: 0.7 },
						'-=0.45'
					);
				}

				if (timelineElement) {
					timeline.from(
						timelineElement,
						{ opacity: 0, y: 35, duration: 0.8 },
						'-=0.35'
					);
				}

				if (documentationElement) {
					timeline.from(
						documentationElement,
						{ opacity: 0, y: 35, duration: 0.8 },
						'-=0.35'
					);
				}
			}, pageElement);

			cleanup = () => context.revert();
		});

		return () => {
			clearInterval(clock);
			stopIntro?.();
			cleanup?.();
		};
	});
</script>

{#snippet sparkle(className: string)}
	<svg viewBox="0 0 24 24" aria-hidden="true" class={className}>
		<path
			fill="currentColor"
			d="M12 0C13 7 17 11 24 12C17 13 13 17 12 24C11 17 7 13 0 12C7 11 11 7 12 0Z"
		/>
	</svg>
{/snippet}

{#snippet ornamentDots()}
	<span class="flex items-center gap-1">
		<i class="size-1 rounded-full bg-[#164f88]/60"></i>
		<i class="size-1.5 rounded-full bg-[#164f88]"></i>
		{@render sparkle('size-3 text-[#164f88] sm:size-3.5')}
		<i class="size-1.5 rounded-full bg-[#164f88]"></i>
		<i class="size-1 rounded-full bg-[#164f88]/60"></i>
	</span>
{/snippet}

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
		relative
		min-h-screen w-full overflow-hidden
		bg-linear-to-b
		from-blue-50 via-white to-blue-50
	"
>
	<div
		aria-hidden="true"
		class="
			pointer-events-none
			absolute inset-x-0 top-0
			h-[1100px] max-sm:h-[900px]
		"
	>
		<div
			class="
				absolute top-[-200px] left-1/2
				size-[900px] -translate-x-1/2
				rounded-full
				bg-[radial-gradient(circle,rgba(147,197,253,0.45),rgba(147,197,253,0)_65%)]
				max-sm:size-[600px]
			"
		></div>
	</div>

	<section
		class="
			relative
			mx-auto flex min-h-screen
			w-full max-w-7xl
			flex-col items-center
			px-6 pt-20 pb-28
			text-center
			sm:px-8
			md:pt-24 md:pb-32
			lg:px-12
		"
	>
		<div class="mb-6 flex flex-col items-center sm:mb-8">
			<div class="relative flex size-24 items-center justify-center sm:size-32">
				<span
					aria-hidden="true"
					class="
						absolute inset-0 rounded-full
						bg-[radial-gradient(circle,rgba(251,146,60,0.45),rgba(251,146,60,0)_68%)]
					"
				></span>

				<span
					class="
						relative flex size-16 items-center justify-center
						rounded-full border-[3px] border-[#164f88]
						bg-white shadow-[0_5px_0_#0f3a6b]
						sm:size-20
					"
				>
					<img
						src="/logo/logo-kabinet.webp"
						alt="Logo Kabinet Cerita Loka"
						class="size-11 object-contain sm:size-14"
					/>
				</span>
			</div>

			{#if program.ministry_name}
				<div class="mt-3 flex items-center gap-2 sm:gap-3">
					{@render sparkle('size-3 shrink-0 text-[#164f88] sm:size-3.5')}

					<span
						class="
							text-[11px] font-black uppercase
							tracking-[0.22em] text-[#164f88]
							sm:text-sm
						"
					>
						{program.ministry_name}
					</span>

					{@render sparkle('size-3 shrink-0 text-[#164f88] sm:size-3.5')}
				</div>
			{/if}
		</div>

		<h1
			bind:this={titleElement}
			class="
				max-w-6xl
				text-[clamp(2rem,5vw,5rem)]
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
					mt-10 max-w-4xl
					text-base font-semibold
					leading-relaxed
					text-blue-950/70
					sm:text-lg
				"
			>
				{#each descriptionParts as part, index (`${index}-${part.text}`)}
					<span
						class={part.highlighted
							? 'group relative isolate inline-block px-1 font-black text-blue-950'
							: ''}
					>
						{#if part.highlighted}
							<span
								aria-hidden="true"
								class="pointer-events-none absolute inset-x-0 bottom-[0.12em] z-0 h-[0.42em] origin-left scale-x-0 rounded-sm bg-gradient-to-r from-[#fde047]/0 via-[#fde047]/90 to-[#fde047]/0 transition-transform duration-300 ease-out will-change-[transform] group-hover:scale-x-100 motion-reduce:transition-none"
							></span>
						{/if}
						<span class={part.highlighted ? 'relative z-10' : ''}>
							{part.text}
						</span>
					</span>
				{/each}
			</p>
		{/if}

		{#if schedule}
			<div
				class="
					relative mt-10 w-full max-w-3xl
					rounded-[22px]
					bg-linear-to-b from-white via-blue-50 to-blue-200/70
					p-2
					shadow-[0_10px_30px_rgba(22,79,136,0.18)]
					sm:p-2.5
				"
			>
				<div class="relative rounded-[16px] border-2 border-[#164f88]/70 p-1">
					<div
						class="
							relative rounded-[12px]
							border border-[#164f88]/50
							px-2 pt-6 pb-5
							sm:px-6 sm:pt-8 sm:pb-7
						"
					>
						{@render sparkle('absolute top-1.5 left-1.5 size-3 text-[#164f88] sm:size-3.5')}
						{@render sparkle('absolute top-1.5 right-1.5 size-3 text-[#164f88] sm:size-3.5')}
						{@render sparkle('absolute bottom-1.5 left-1.5 size-3 text-[#164f88] sm:size-3.5')}
						{@render sparkle('absolute right-1.5 bottom-1.5 size-3 text-[#164f88] sm:size-3.5')}

						<span class="absolute -top-2 left-1/2 -translate-x-1/2 bg-white px-2">
							{@render ornamentDots()}
						</span>

						<span class="absolute -bottom-2 left-1/2 -translate-x-1/2 rounded bg-blue-200 px-2">
							{@render ornamentDots()}
						</span>

						<div class="grid grid-cols-3 divide-x-2 divide-dotted divide-[#164f88]/30">
							{#each [
								{ label: 'Tanggal', value: schedule.label },
								{ label: 'Durasi', value: `${schedule.totalDays} Hari` },
								{ label: 'Status', value: statusLabel }
							] as item (item.label)}
								<div class="flex flex-col items-center px-1 sm:px-4">
									<span
										class="
											text-[9px] font-black uppercase
											tracking-[0.2em] text-[#164f88]/60
											sm:text-xs
										"
									>
										{item.label}
									</span>

									<span
										class="
											mt-1.5 text-[13px] leading-tight
											font-black tracking-[-0.03em]
											text-[#0f3a6b]
											sm:mt-2 sm:text-2xl
										"
									>
										{item.value}
									</span>
								</div>
							{/each}
						</div>
					</div>
				</div>
			</div>
		{/if}

		{#if mounted && countdown}
			<div class="mt-6 w-full max-w-3xl">
				<p
					class="
						text-[11px] font-black uppercase
						tracking-[0.22em] text-[#164f88]
						sm:text-sm
					"
				>
					✦ {phase === 'upcoming' ? 'Dimulai dalam' : 'Berakhir dalam'} ✦
				</p>

				<div class="mt-3 grid grid-cols-4 gap-2 sm:gap-4">
					{#each countdown as unit (unit.label)}
						<div
							class="
								countdown-tile
								relative overflow-hidden
								rounded-[14px] border-[3px] border-blue-950
								px-1 py-3
								shadow-[0_6px_0_rgba(30,64,175,0.55)]
								sm:rounded-[18px] sm:border-[5px] sm:py-5
							"
						>
							<span
								class="
									block text-[24px] leading-none
									font-black tabular-nums
									text-white
									[-webkit-text-stroke:2px_#ea580c]
									[paint-order:stroke_fill]
									sm:text-6xl
								"
							>
								{unit.value}
							</span>

							<span
								class="
									mt-2 block text-[8px] font-black
									tracking-[0.18em] text-blue-100
									sm:text-xs sm:tracking-[0.2em]
								"
							>
								{unit.label}
							</span>
						</div>
					{/each}
				</div>

				{#if progress && schedule}
					<div class="mx-auto mt-4 max-w-xl">
						<div
							class="
								flex justify-between
								text-[10px] font-black uppercase
								tracking-wider text-[#164f88]/70
								sm:text-xs
							"
						>
							<span>Hari ke-{progress.day} dari {schedule.totalDays}</span>
							<span>{progress.percent}%</span>
						</div>

						<div class="mt-1.5 h-3 overflow-hidden rounded-full border-2 border-blue-950 bg-white">
							<div
								class="h-full rounded-full bg-linear-to-r from-orange-300 to-orange-500 transition-[width] duration-700"
								style={`width:${progress.percent}%`}
							></div>
						</div>
					</div>
				{/if}
			</div>
		{/if}

		{#if startDate && endDate}
			<div
				bind:this={timelineElement}
				class="mt-12 w-full"
			>
				<TimelinePelaksanaan
					{startDate}
					{endDate}
					paused={!timelineReady}
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

	{#if introActive}
		<div
			use:portal
			aria-hidden="true"
			class="spotlight-intro"
			class:spotlight-leaving={introLeaving}
			style={introStyle}
		>
			<div class="spotlight-dark"></div>
			<div class="spotlight-beam spotlight-beam-left"></div>
			<div class="spotlight-beam spotlight-beam-right"></div>
			<div class="spotlight-lamp spotlight-lamp-left"></div>
			<div class="spotlight-lamp spotlight-lamp-right"></div>
		</div>
	{/if}
</main>

<style>
	.countdown-tile {
		background-color: #102f5a;
		background-image:
			linear-gradient(rgba(191, 219, 254, 0.35) 1px, transparent 1px),
			linear-gradient(90deg, rgba(191, 219, 254, 0.35) 1px, transparent 1px);
		background-size: 18px 18px;
	}

	@property --spot-r {
		syntax: '<length>';
		inherits: false;
		initial-value: 0px;
	}

	.spotlight-intro {
		position: fixed;
		inset: 0;
		z-index: 9999;
		overflow: hidden;
		pointer-events: none;
		transition: opacity 0.3s ease;
	}

	.spotlight-leaving {
		opacity: 0;
	}

	.spotlight-dark {
		position: absolute;
		inset: 0;
		background: radial-gradient(ellipse at 50% 30%, #0b2545, #040d1c 70%);
		-webkit-mask-image: radial-gradient(
			ellipse calc(var(--spot-r) * 1.9) var(--spot-r) at var(--sx) var(--sy),
			transparent 60%,
			#000 100%
		);
		mask-image: radial-gradient(
			ellipse calc(var(--spot-r) * 1.9) var(--spot-r) at var(--sx) var(--sy),
			transparent 60%,
			#000 100%
		);
		animation:
			spotlight-open 2.6s linear forwards,
			spotlight-fade 0.5s ease 2.2s forwards;
	}

	.spotlight-beam {
		position: absolute;
		top: -40px;
		width: 240px;
		transform-origin: 50% 0;
		background: linear-gradient(
			to bottom,
			rgba(255, 244, 214, 0.85),
			rgba(255, 236, 190, 0.45) 75%,
			rgba(255, 236, 190, 0)
		);
		clip-path: polygon(47% 0, 53% 0, 100% 100%, 0 100%);
		filter: blur(4px);
		mix-blend-mode: screen;
		opacity: 0;
	}

	.spotlight-beam-left {
		left: -120px;
		height: var(--len-l);
		animation:
			spotlight-beam-left 1s cubic-bezier(0.5, 0, 0.2, 1) 0.2s forwards,
			spotlight-fade 0.6s ease 2.1s forwards;
	}

	.spotlight-beam-right {
		right: -120px;
		height: var(--len-r);
		animation:
			spotlight-beam-right 1s cubic-bezier(0.5, 0, 0.2, 1) 0.2s forwards,
			spotlight-fade 0.6s ease 2.1s forwards;
	}

	.spotlight-lamp {
		position: absolute;
		top: -30px;
		width: 70px;
		height: 70px;
		border-radius: 999px;
		background: radial-gradient(circle, #fff7e0, rgba(255, 220, 150, 0) 70%);
		opacity: 0;
		animation:
			spotlight-show 0.3s ease 0.2s forwards,
			spotlight-fade 0.6s ease 2.1s forwards;
	}

	.spotlight-lamp-left {
		left: -35px;
	}

	.spotlight-lamp-right {
		right: -35px;
	}

	@media (max-width: 640px) {
		.spotlight-beam {
			width: 150px;
		}

		.spotlight-beam-left {
			left: -75px;
		}

		.spotlight-beam-right {
			right: -75px;
		}
	}

	@keyframes spotlight-open {
		0%,
		38% {
			--spot-r: 0px;
		}

		50% {
			--spot-r: var(--spot);
		}

		72% {
			--spot-r: calc(var(--spot) + 20px);
		}

		100% {
			--spot-r: 1600px;
		}
	}

	@keyframes spotlight-beam-left {
		0% {
			opacity: 0;
			transform: rotate(-8deg);
		}

		15% {
			opacity: 1;
		}

		35% {
			opacity: 0.4;
		}

		45%,
		100% {
			opacity: 1;
		}

		100% {
			transform: rotate(var(--angle-l));
		}
	}

	@keyframes spotlight-beam-right {
		0% {
			opacity: 0;
			transform: rotate(8deg);
		}

		15% {
			opacity: 1;
		}

		35% {
			opacity: 0.4;
		}

		45%,
		100% {
			opacity: 1;
		}

		100% {
			transform: rotate(var(--angle-r));
		}
	}

	@keyframes spotlight-show {
		to {
			opacity: 1;
		}
	}

	@keyframes spotlight-fade {
		to {
			opacity: 0;
		}
	}
</style>