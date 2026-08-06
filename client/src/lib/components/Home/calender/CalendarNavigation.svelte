<script lang="ts">
  import { monthLabel } from './calendar-utils';

  let {
    currentMonth,
    onPrevious,
    onNext
  }: {
    currentMonth: Date;
    onPrevious: () => void;
    onNext: () => void;
  } = $props();
</script>

<div
  class="
    mt-8 grid
    grid-cols-[52px_minmax(0,1fr)_52px]
    items-center gap-3
    sm:mt-10
    sm:grid-cols-[60px_minmax(0,1fr)_60px]
    sm:gap-5
  "
>
  <button
    type="button"
    aria-label="Tampilkan bulan sebelumnya"
    onclick={onPrevious}
    class="
      group flex size-13
      items-center justify-center
      rounded-2xl
      bg-[#164f88]
      text-white
      shadow-md
      transition duration-200
      hover:-translate-y-0.5
      hover:bg-[#0d3f70]
      hover:shadow-lg
      focus-visible:outline-2
      focus-visible:outline-offset-3
      focus-visible:outline-orange-500
      active:translate-y-0
      sm:size-15
    "
  >
    <svg
      viewBox="0 0 24 24"
      fill="none"
      aria-hidden="true"
      class="
        size-6
        transition-transform
        group-hover:-translate-x-0.5
      "
    >
      <path
        d="M15 5L8 12L15 19"
        stroke="currentColor"
        stroke-width="2.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  </button>

  <div
    class="
      min-w-0 text-center
    "
  >
    <p
      class="
        truncate text-3xl
        leading-none font-black
        capitalize
        tracking-[-0.045em]
        px-8 py-2.5 rounded-md
        bg-amber-600 text-white
        sm:text-4xl
        md:text-5xl
      "
    >
      {monthLabel(currentMonth)}
    </p>

    <p
      class="
        mt-1.5
        text-xs font-black
        tracking-[0.2em]
        text-orange-600
        sm:text-sm
      "
    >
      {currentMonth.getFullYear()}
    </p>
  </div>

  <button
    type="button"
    aria-label="Tampilkan bulan berikutnya"
    onclick={onNext}
    class="
      group flex size-13
      items-center justify-center
      rounded-2xl
      bg-[#164f88]
      text-white
      shadow-md
      transition duration-200
      hover:-translate-y-0.5
      hover:bg-[#0d3f70]
      hover:shadow-lg
      focus-visible:outline-2
      focus-visible:outline-offset-3
      focus-visible:outline-orange-500
      active:translate-y-0
      sm:size-15
    "
  >
    <svg
      viewBox="0 0 24 24"
      fill="none"
      aria-hidden="true"
      class="
        size-6
        transition-transform
        group-hover:translate-x-0.5
      "
    >
      <path
        d="M9 5L16 12L9 19"
        stroke="currentColor"
        stroke-width="2.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  </button>
</div>