<script lang="ts">
  import { buildProgramWorkHref } from './calendar-utils';
  import type { CalendarEvent } from './types';

  let {
    event,
    logo,
    positionStyle,
    continuesBefore = false,
    continuesAfter = false
  }: {
    event: CalendarEvent;
    logo: string;
    positionStyle: string;
    continuesBefore?: boolean;
    continuesAfter?: boolean;
  } = $props();
</script>

<a
  href={buildProgramWorkHref(event)}
  aria-label={`Lihat program kerja ${event.title}`}
  title={event.title}
  style={positionStyle}
  class={`
    group absolute z-20
    flex h-8 items-center
    overflow-hidden
    border border-blue-950/15
    bg-linear-to-r
    from-[#0d4d88]
    to-[#277bb8]
    pr-3 text-white
    shadow-sm
    shadow-blue-950/15
    transition duration-200
    hover:z-30
    hover:-translate-y-px
    hover:from-[#093f73]
    hover:to-[#1e679d]
    hover:shadow-md
    focus-visible:z-30
    focus-visible:outline-2
    focus-visible:outline-offset-1
    focus-visible:outline-orange-500

    ${
      continuesBefore
        ? 'rounded-l-lg'
        : 'rounded-l-full'
    }

    ${
      continuesAfter
        ? 'rounded-r-lg'
        : 'rounded-r-full'
    }
  `}
>
  {#if continuesBefore}
    <span
      aria-hidden="true"
      class="
        mr-1 h-full w-1.5
        shrink-0
        bg-white/20
      "
    ></span>
  {:else}
    <span
      class="
        ml-0.5 flex size-7
        shrink-0
        items-center justify-center
        rounded-full
        bg-white shadow-sm
        transition-transform
        group-hover:scale-105
      "
    >
      <img
        src={logo}
        alt=""
        draggable="false"
        class="size-5 object-contain"
      />
    </span>
  {/if}

  <span
    class="
      min-w-0 truncate
      px-2
      text-[10px]
      font-extrabold
      text-white
      sm:text-xs
    "
  >
    {event.title}
  </span>

  {#if continuesAfter}
    <svg
      viewBox="0 0 24 24"
      fill="none"
      aria-hidden="true"
      class="
        ml-auto size-3.5
        shrink-0 text-white/75
      "
    >
      <path
        d="M9 6L15 12L9 18"
        stroke="currentColor"
        stroke-width="2.5"
        stroke-linecap="round"
        stroke-linejoin="round"
      />
    </svg>
  {/if}
</a>