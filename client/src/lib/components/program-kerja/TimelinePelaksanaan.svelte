<script lang="ts">
  let {
    startDate,
    endDate,
    paused = false
  }: {
    startDate: string;
    endDate?: string | null;
    paused?: boolean;
  } = $props();

  const monthNames = [
    'JANUARI',
    'FEBRUARI',
    'MARET',
    'APRIL',
    'MEI',
    'JUNI',
    'JULI',
    'AGUSTUS',
    'SEPTEMBER',
    'OKTOBER',
    'NOVEMBER',
    'DESEMBER'
  ];

  function parseDate(value: string): Date | null {
    const [year, month, day] = value
      .split('-')
      .map(Number);

    if (
      !Number.isInteger(year) ||
      !Number.isInteger(month) ||
      !Number.isInteger(day)
    ) {
      return null;
    }

    const date = new Date(year, month - 1, day);

    if (
      date.getFullYear() !== year ||
      date.getMonth() !== month - 1 ||
      date.getDate() !== day
    ) {
      return null;
    }

    return date;
  }

  function isSameDate(
    firstDate: Date,
    secondDate: Date
  ): boolean {
    return (
      firstDate.getFullYear() ===
        secondDate.getFullYear() &&
      firstDate.getMonth() ===
        secondDate.getMonth() &&
      firstDate.getDate() ===
        secondDate.getDate()
    );
  }

  const dateLabel = $derived.by(() => {
    const start = parseDate(startDate);

    const end = endDate
      ? parseDate(endDate)
      : null;

    if (!start) {
      return startDate;
    }

    if (!end || isSameDate(start, end)) {
      return `${start.getDate()} ${
        monthNames[start.getMonth()]
      } ${start.getFullYear()}`;
    }

    const sameMonth =
      start.getMonth() === end.getMonth();

    const sameYear =
      start.getFullYear() === end.getFullYear();

    if (sameMonth && sameYear) {
      return `${start.getDate()}–${end.getDate()} ${
        monthNames[start.getMonth()]
      } ${start.getFullYear()}`;
    }

    if (sameYear) {
      return `${start.getDate()} ${
        monthNames[start.getMonth()]
      }–${end.getDate()} ${
        monthNames[end.getMonth()]
      } ${start.getFullYear()}`;
    }

    return `${start.getDate()} ${
      monthNames[start.getMonth()]
    } ${start.getFullYear()}–${end.getDate()} ${
      monthNames[end.getMonth()]
    } ${end.getFullYear()}`;
  });

  const timelineLetters = 'TIMELINE'.split('');
  const pelaksanaanLetters = 'PELAKSANAAN'.split('');

  let root = $state<HTMLDivElement | null>(null);
  let animReady = $state(false);
  let playing = $state(false);

  function createNoise(context: AudioContext, duration: number): AudioBuffer {
    const buffer = context.createBuffer(
      1,
      Math.floor(context.sampleRate * duration),
      context.sampleRate
    );
    const data = buffer.getChannelData(0);

    for (let index = 0; index < data.length; index += 1) {
      data[index] = Math.random() * 2 - 1;
    }

    return buffer;
  }

  function playNoise(
    context: AudioContext,
    output: AudioNode,
    {
      at,
      duration,
      type,
      frequency,
      endFrequency,
      volume
    }: {
      at: number;
      duration: number;
      type: BiquadFilterType;
      frequency: number;
      endFrequency?: number;
      volume: number;
    }
  ): void {
    const source = context.createBufferSource();
    const filter = context.createBiquadFilter();
    const gain = context.createGain();

    source.buffer = createNoise(context, duration);
    filter.type = type;
    filter.frequency.setValueAtTime(frequency, at);

    if (endFrequency) {
      filter.frequency.exponentialRampToValueAtTime(endFrequency, at + duration);
    }

    gain.gain.setValueAtTime(0.0001, at);
    gain.gain.exponentialRampToValueAtTime(volume, at + Math.min(0.04, duration / 3));
    gain.gain.exponentialRampToValueAtTime(0.0001, at + duration);

    source.connect(filter).connect(gain).connect(output);
    source.start(at);
    source.stop(at + duration + 0.05);
  }

  function playTone(
    context: AudioContext,
    output: AudioNode,
    {
      at,
      duration,
      frequency,
      endFrequency,
      volume,
      type = 'sine'
    }: {
      at: number;
      duration: number;
      frequency: number;
      endFrequency?: number;
      volume: number;
      type?: OscillatorType;
    }
  ): void {
    const oscillator = context.createOscillator();
    const gain = context.createGain();

    oscillator.type = type;
    oscillator.frequency.setValueAtTime(frequency, at);

    if (endFrequency) {
      oscillator.frequency.exponentialRampToValueAtTime(endFrequency, at + duration);
    }

    gain.gain.setValueAtTime(0.0001, at);
    gain.gain.exponentialRampToValueAtTime(volume, at + 0.012);
    gain.gain.exponentialRampToValueAtTime(0.0001, at + duration);

    oscillator.connect(gain).connect(output);
    oscillator.start(at);
    oscillator.stop(at + duration + 0.05);
  }

  async function playBacksound(): Promise<void> {
    try {
      const AudioContextClass =
        window.AudioContext ??
        (window as unknown as { webkitAudioContext: typeof AudioContext })
          .webkitAudioContext;

      if (!AudioContextClass) {
        return;
      }

      const context = new AudioContextClass();

      if (context.state === 'suspended') {
        await context.resume();
      }

      if (context.state !== 'running') {
        await context.close();
        return;
      }

      const master = context.createGain();
      master.gain.value = 0.55;
      master.connect(context.destination);

      const now = context.currentTime + 0.02;
      const impact = now + 0.63;

      playNoise(context, master, {
        at: now,
        duration: 0.6,
        type: 'bandpass',
        frequency: 400,
        endFrequency: 1800,
        volume: 0.12
      });

      playTone(context, master, {
        at: impact,
        duration: 0.45,
        frequency: 140,
        endFrequency: 42,
        volume: 0.9
      });

      playNoise(context, master, {
        at: impact,
        duration: 0.22,
        type: 'lowpass',
        frequency: 900,
        endFrequency: 120,
        volume: 0.5
      });

      [0.04, 0.09, 0.15, 0.2, 0.28, 0.36].forEach((offset, index) => {
        playNoise(context, master, {
          at: impact + offset,
          duration: 0.05 + (index % 2) * 0.03,
          type: 'highpass',
          frequency: 1800 + index * 250,
          volume: 0.22 - index * 0.025
        });
      });

      const notes = [523.25, 587.33, 659.25, 783.99, 880, 1046.5];

      timelineLetters.forEach((_, index) => {
        playTone(context, master, {
          at: now + 0.95 + index * 0.045,
          duration: 0.12,
          frequency: notes[index % notes.length],
          volume: 0.05,
          type: 'triangle'
        });
      });

      pelaksanaanLetters.forEach((_, index) => {
        playTone(context, master, {
          at: now + 1.3 + index * 0.04,
          duration: 0.12,
          frequency: notes[index % notes.length] * 1.5,
          volume: 0.045,
          type: 'triangle'
        });
      });

      setTimeout(() => {
        void context.close();
      }, 3500);
    } catch {
      // browser menolak audio, animasi tetap jalan
    }
  }

  $effect(() => {
    if (!root) {
      return;
    }

    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (reduceMotion || !('IntersectionObserver' in window)) {
      return;
    }

    animReady = true;
  });

  $effect(() => {
    if (!root || !animReady || paused || playing) {
      return;
    }

    const observer = new IntersectionObserver(
      (entries) => {
        if (entries.some((entry) => entry.isIntersecting)) {
          playing = true;
          void playBacksound();
          observer.disconnect();
        }
      },
      { threshold: 0.55 }
    );

    observer.observe(root);

    return () => observer.disconnect();
  });
</script>


<div
  bind:this={root}
  class:tl-anim={animReady}
  class:tl-play={playing}
  class="
    relative mx-auto w-full
    max-w-[900px]
    px-[7cqw] pt-[12cqw] pb-[2cqw]
    [container-type:inline-size]
    max-sm:px-[3cqw] max-sm:pt-[17cqw]
  "
>
  <div class="tl-card relative z-10">
    <div
      aria-hidden="true"
      class="
        absolute inset-0
        translate-x-[1.2%]
        translate-y-[4%]
        rounded-[18px]
        bg-blue-950/30
      "
    ></div>

    <div
      class="
        timeline-grid
        relative
        overflow-hidden rounded-[18px]
        border-[6px] border-blue-950
        bg-blue-950
        px-[6cqw] pt-[7cqw] pb-[5cqw]
        max-sm:border-4
        max-sm:px-[5cqw] max-sm:pt-[10cqw] max-sm:pb-[6cqw]
        shadow-[0_10px_0_rgba(30,64,175,0.65)]
      "
    >
      <div
        class="
          absolute inset-[14px]
          border-2 border-blue-200/70
          max-sm:inset-[8px]
        "
      ></div>

      <div
        class="
          relative z-20 flex
          min-h-[15cqw]
          items-center justify-center
          rounded-[14px]
          bg-white px-[4cqw] py-[3cqw]
          max-sm:min-h-[22cqw]
          max-sm:rounded-[10px]
          max-sm:py-[4cqw]
          shadow-[0_12px_14px_rgba(15,23,42,0.28)]
        "
      >
        <div
          class="
            flex w-full flex-col
            items-center justify-center
          "
        >
          <h2
            class="
              timeline-date
              max-w-full text-center
              text-balance
              text-[5cqw]
              max-sm:text-[6.4cqw]
              leading-[1.05] font-black
              tracking-[-0.06em]
              text-blue-900
            "
          >
            {dateLabel}
          </h2>

          <div
            class="
              mt-[2cqw] h-[5px]
              w-[46%] rounded-full
              max-sm:h-[3px]
              bg-orange-400
              shadow-[0_2px_4px_rgba(249,115,22,0.35)]
            "
          ></div>
        </div>
      </div>
    </div>

    <div
      class="
        absolute left-[3cqw] bottom-[calc(100%-4.2cqw)]
        z-40 flex flex-col items-start
        origin-bottom-left
        -rotate-[3deg]
        max-sm:left-[2cqw] max-sm:bottom-[calc(100%-6cqw)]
      "
    >
      <div
        class="
          tl-label tl-label-first
          relative z-10 ml-[1.6cqw] flex w-fit
          items-center
          rounded-t-[12px]
          border-[3px] border-b-0 border-orange-700
          bg-gradient-to-b
          from-orange-300
          via-orange-400
          to-orange-500
          px-[2.6cqw] pt-[1.2cqw] pb-[1cqw]
          max-sm:rounded-t-[8px] max-sm:border-2 max-sm:border-b-0
          max-sm:px-[3.2cqw] max-sm:pt-[1.6cqw] max-sm:pb-[1.3cqw]
        "
      >
        <span
          aria-label="TIMELINE"
          class="
            whitespace-nowrap
            text-[3.2cqw]
            max-sm:text-[4.6cqw]
            leading-none font-black
            tracking-[-0.04em]
            text-white
            [-webkit-text-stroke:2px_#c2410c]
            [paint-order:stroke_fill]
          "
        >
          {#each timelineLetters as letter, index (index)}
          <span class="tl-letter tl-letter-first" style={`--i:${index}`}>{letter}</span>
        {/each}
        </span>
      </div>

      <div
        class="
          tl-label tl-label-second
          flex w-fit
          items-center
          rounded-[12px] rounded-tl-[4px]
          border-[3px] border-blue-950
          bg-gradient-to-b
          from-[#2563a8]
          via-[#164f88]
          to-[#0f3a6b]
          px-[3cqw] py-[1.6cqw]
          max-sm:rounded-[8px] max-sm:rounded-tl-[3px] max-sm:border-2
          max-sm:px-[3.6cqw] max-sm:py-[2.2cqw]
          shadow-[0_6px_0_#0b2545,0_12px_16px_rgba(15,23,42,0.3)]
          max-sm:shadow-[0_4px_0_#0b2545,0_8px_10px_rgba(15,23,42,0.3)]
        "
      >
        <span
          aria-label="PELAKSANAAN"
          class="
            whitespace-nowrap
            text-[4.4cqw]
            max-sm:text-[6.2cqw]
            leading-none font-black
            tracking-[-0.05em]
            text-white
            [-webkit-text-stroke:2px_#ea580c]
            [paint-order:stroke_fill]
            drop-shadow-[0_3px_0_rgba(255,255,255,0.35)]
          "
        >
          {#each pelaksanaanLetters as letter, index (index)}
          <span class="tl-letter tl-letter-second" style={`--i:${index}`}>{letter}</span>
        {/each}
        </span>
      </div>
    </div>
  </div>

  <div
    aria-hidden="true"
    class="
      tl-ground
      relative z-30
      -mx-[7cqw] -mt-[1.6cqw]
      h-[13cqw]
      max-sm:-mx-[3cqw] max-sm:-mt-[2.2cqw] max-sm:h-[18cqw]
    "
  >
    <div
      class="
        tl-soil
        absolute inset-x-0 top-[1.8cqw] bottom-0
        overflow-hidden
        max-sm:top-[2.4cqw]
      "
    >
      <svg
        class="absolute inset-x-0 top-0 h-full w-full"
        viewBox="0 0 180 20"
        preserveAspectRatio="xMidYMin meet"
      >
        <g
          fill="none"
          stroke="rgba(78, 45, 18, 0.6)"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="0.7"
        >
          <path class="tl-crack" style="--d:0s" pathLength="1" d="M90 0 L88 4 L92 7 L89 11 L93 15 L91 19" />
          <path class="tl-crack" style="--d:0.02s" pathLength="1" d="M25 0 L27 4 L23 8 L26 12 L22 17" />
          <path class="tl-crack" style="--d:0.02s" pathLength="1" d="M155 0 L153 4 L157 8 L154 12 L158 17" />
          <path class="tl-crack" style="--d:0.08s" pathLength="1" d="M88 4 L80 6 L74 5 L66 9 L58 8 L50 12 L42 11" />
          <path class="tl-crack" style="--d:0.1s" pathLength="1" d="M92 7 L100 9 L107 7 L115 11 L124 10 L132 14 L140 13" />
          <path class="tl-crack" style="--d:0.12s" pathLength="1" d="M27 4 L34 6 L39 10 L46 9" />
          <path class="tl-crack" style="--d:0.12s" pathLength="1" d="M153 4 L146 6 L141 10 L134 9" />
          <path class="tl-crack" style="--d:0.16s" pathLength="1" d="M23 8 L16 10 L10 9 L4 13" />
          <path class="tl-crack" style="--d:0.16s" pathLength="1" d="M157 8 L164 10 L170 9 L176 13" />
        </g>
        <g
          fill="none"
          stroke="rgba(78, 45, 18, 0.6)"
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="0.45"
        >
          <path class="tl-crack" style="--d:0.3s" pathLength="1" d="M74 5 L70 1" />
          <path class="tl-crack" style="--d:0.34s" pathLength="1" d="M66 9 L64 14 L60 17" />
          <path class="tl-crack" style="--d:0.4s" pathLength="1" d="M50 12 L47 16" />
          <path class="tl-crack" style="--d:0.3s" pathLength="1" d="M107 7 L110 3" />
          <path class="tl-crack" style="--d:0.34s" pathLength="1" d="M115 11 L117 16 L121 18" />
          <path class="tl-crack" style="--d:0.4s" pathLength="1" d="M132 14 L135 18" />
          <path class="tl-crack" style="--d:0.22s" pathLength="1" d="M89 11 L84 15 L79 14 L74 18" />
          <path class="tl-crack" style="--d:0.26s" pathLength="1" d="M93 15 L98 17 L104 16" />
          <path class="tl-crack" style="--d:0.3s" pathLength="1" d="M26 12 L31 15" />
          <path class="tl-crack" style="--d:0.3s" pathLength="1" d="M154 12 L149 15" />
          <path class="tl-crack" style="--d:0.36s" pathLength="1" d="M10 9 L8 4" />
          <path class="tl-crack" style="--d:0.36s" pathLength="1" d="M170 9 L172 4" />
        </g>
      </svg>
    </div>

    <div
      class="tl-grass absolute inset-x-0 top-0 h-[2.6cqw] max-sm:h-[3.4cqw]"
    ></div>

    <span class="tl-dust tl-dust-left" style="--x:-1"></span>
    <span class="tl-dust tl-dust-right" style="--x:1"></span>
  </div>
</div>

<style>
  .timeline-grid {
    background-color: #102f5a;
    background-image:
      linear-gradient(
        rgba(191, 219, 254, 0.58) 1px,
        transparent 1px
      ),
      linear-gradient(
        90deg,
        rgba(191, 219, 254, 0.58) 1px,
        transparent 1px
      );
    background-size:
      clamp(16px, 2.5vw, 26px)
      clamp(16px, 2.5vw, 26px);
  }

  .timeline-date {
    -webkit-text-stroke: 2px white;
    paint-order: stroke fill;
    text-shadow:
      0 5px 0 #dbeafe,
      0 9px 4px rgba(15, 23, 42, 0.38);
  }

  @media (max-width: 640px) {
    .timeline-date {
      -webkit-text-stroke-width: 1px;
      text-shadow:
        0 3px 0 #dbeafe,
        0 5px 3px rgba(15, 23, 42, 0.32);
    }
  }

  .tl-grass {
    background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 60 24' preserveAspectRatio='none'%3E%3Cpath fill='%2315803d' d='M0 24L4 9L7 24L11 3L15 24L19 11L22 24L27 1L31 24L35 10L38 24L43 5L47 24L51 12L54 24L58 6L60 24Z'/%3E%3Cpath fill='%2322c55e' d='M2 24L6 13L9 24L14 8L17 24L24 7L28 24L33 14L36 24L40 9L44 24L49 15L52 24L56 11L60 24Z'/%3E%3C/svg%3E");
    background-repeat: repeat-x;
    background-size: 3cqw 100%;
    opacity: 0.75;
    -webkit-mask-image: linear-gradient(to right, transparent, #000 12%, #000 88%, transparent);
    mask-image: linear-gradient(to right, transparent, #000 12%, #000 88%, transparent);
  }

  .tl-soil {
    background-image: linear-gradient(
      to bottom,
      rgba(74, 155, 82, 0.4) 0%,
      rgba(139, 90, 43, 0.48) 18%,
      rgba(139, 90, 43, 0.34) 50%,
      rgba(139, 90, 43, 0.14) 78%,
      rgba(255, 255, 255, 0) 100%
    );
    -webkit-mask-image: linear-gradient(to right, transparent, #000 12%, #000 88%, transparent);
    mask-image: linear-gradient(to right, transparent, #000 12%, #000 88%, transparent);
  }

  .tl-crack {
    stroke-dasharray: 1;
    stroke-dashoffset: 0;
  }

  .tl-dust {
    position: absolute;
    top: 0;
    width: 9cqw;
    height: 4cqw;
    border-radius: 999px;
    background: radial-gradient(
      ellipse at center,
      rgba(231, 211, 177, 0.95),
      rgba(231, 211, 177, 0) 70%
    );
    opacity: 0;
    pointer-events: none;
  }

  .tl-dust-left {
    left: 2cqw;
  }

  .tl-dust-right {
    right: 2cqw;
  }

  .tl-letter {
    display: inline-block;
  }

  .tl-anim .tl-card {
    opacity: 0;
    transform: translateY(-14cqw);
  }

  .tl-anim .tl-crack {
    stroke-dashoffset: 1;
  }

  .tl-anim .tl-label,
  .tl-anim .tl-letter {
    opacity: 0;
  }

  .tl-anim.tl-play .tl-card {
    transform-origin: 50% 100%;
    animation: tl-drop 0.9s cubic-bezier(0.55, 0, 1, 0.45) both;
  }

  .tl-anim.tl-play .tl-ground {
    animation: tl-shake 0.38s ease-out 0.63s both;
  }

  .tl-anim.tl-play .tl-crack {
    animation: tl-crack 0.6s ease-out calc(0.66s + var(--d)) both;
  }

  .tl-anim.tl-play .tl-dust {
    animation: tl-dust 0.8s ease-out 0.62s both;
  }

  .tl-anim.tl-play .tl-label {
    animation: tl-label 0.6s cubic-bezier(0.22, 1, 0.36, 1) both;
  }

  .tl-anim.tl-play .tl-label-first {
    animation-delay: 0.9s;
  }

  .tl-anim.tl-play .tl-label-second {
    animation-delay: 1.22s;
  }

  .tl-anim.tl-play .tl-letter {
    animation: tl-letter 0.55s cubic-bezier(0.22, 1, 0.36, 1) both;
  }

  .tl-anim.tl-play .tl-letter-first {
    animation-delay: calc(0.95s + var(--i) * 0.045s);
  }

  .tl-anim.tl-play .tl-letter-second {
    animation-delay: calc(1.3s + var(--i) * 0.04s);
  }

  @keyframes tl-drop {
    0% {
      opacity: 0;
      transform: translateY(-14cqw);
    }

    15% {
      opacity: 1;
    }

    70% {
      transform: translateY(0) scale(1, 1);
      animation-timing-function: ease-out;
    }

    80% {
      transform: translateY(0) scale(1.015, 0.96);
    }

    90% {
      transform: translateY(-0.6cqw) scale(1, 1);
    }

    100% {
      opacity: 1;
      transform: translateY(0) scale(1, 1);
    }
  }

  @keyframes tl-shake {
    0%,
    100% {
      transform: translate(0, 0);
    }

    20% {
      transform: translate(-0.3cqw, 0.4cqw);
    }

    45% {
      transform: translate(0.3cqw, -0.15cqw);
    }

    70% {
      transform: translate(-0.15cqw, 0.1cqw);
    }
  }

  @keyframes tl-crack {
    from {
      stroke-dashoffset: 1;
    }

    to {
      stroke-dashoffset: 0;
    }
  }

  @keyframes tl-dust {
    0% {
      opacity: 0;
      transform: translate(0, 0) scale(0.3);
    }

    25% {
      opacity: 1;
    }

    100% {
      opacity: 0;
      transform: translate(calc(var(--x) * 5cqw), -2.5cqw) scale(1.5);
    }
  }

  @keyframes tl-label {
    from {
      opacity: 0;
      transform: translateY(-3cqw);
    }

    to {
      opacity: 1;
      transform: translateY(0);
    }
  }

  @keyframes tl-letter {
    from {
      opacity: 0;
      transform: translateY(-1.1em);
    }

    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
</style>