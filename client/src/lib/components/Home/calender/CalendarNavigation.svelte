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
    mt-10 flex items-center
    justify-center gap-5
    sm:mt-12 sm:gap-10
  "
>
  <button
    type="button"
    aria-label="Bulan sebelumnya"
    onclick={onPrevious}
    class="
      group flex size-11
      items-center justify-center
      rounded-2xl
      border border-white/70
      bg-white/30
      text-orange-500
      shadow-sm
      backdrop-blur-md
      transition duration-200
      hover:bg-white/50
      sm:size-12
    "
  >
    <svg
      viewBox="0 0 24 24"
      fill="none"
      aria-hidden="true"
      class="
        size-6 transition-transform
        group-hover:-translate-x-0.5
      "
    >
      <path
        d="M15 5L8 12L15 19"
        stroke="currentColor"
        stroke-width="3"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  </button>

  <div
    class="
      min-w-44 rounded-2xl
      border border-orange-300/70
      bg-gradient-to-b
      from-orange-400 to-orange-300
      px-7 py-3 text-center
      shadow-sm
      sm:min-w-56
    "
  >
    <p
      class="
        text-2xl leading-none
        font-black text-white
        sm:text-3xl
      "
    >
      {monthLabel(currentMonth)}
    </p>

    <p
      class="
        mt-1 text-xs font-extrabold
        tracking-[0.16em] text-white
      "
    >
      {currentMonth.getFullYear()}
    </p>
  </div>

  <button
    type="button"
    aria-label="Bulan berikutnya"
    onclick={onNext}
    class="
      group flex size-11
      items-center justify-center
      rounded-2xl
      border border-white/70
      bg-white/30
      text-orange-500
      shadow-sm
      backdrop-blur-md
      transition duration-200
      hover:bg-white/50
      sm:size-12
    "
  >
    <svg
      viewBox="0 0 24 24"
      fill="none"
      aria-hidden="true"
      class="
        size-6 transition-transform
        group-hover:translate-x-0.5
      "
    >
      <path
        d="M9 5L16 12L9 19"
        stroke="currentColor"
        stroke-width="3"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  </button>
</div>
