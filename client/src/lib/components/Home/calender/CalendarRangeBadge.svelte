<script lang="ts">
  import { resolve } from '$app/paths';
  import {
    buildProgramWorkHref,
    formatEventRange,
    parseDate
  } from './calendar-utils';
  import type { CalendarEvent } from './types';

  let {
    event,
    logo,
    positionStyle,
    continuesBefore = false,
    continuesAfter = false,
    onShowEnd
  }: {
    event: CalendarEvent;
    logo: string;
    positionStyle: string;
    continuesBefore?: boolean;
    continuesAfter?: boolean;
    onShowEnd?: () => void;
  } = $props();

  const dayCount = $derived.by(() => {
    if (!event.endDate) {
      return 1;
    }

    const difference =
      parseDate(event.endDate).getTime() -
      parseDate(event.startDate).getTime();

    return Math.max(
      1,
      Math.round(difference / 86400000) + 1
    );
  });
</script>

<div
  style={positionStyle}
  class="
    group/range absolute z-20
    h-8
    hover:z-40
    focus-within:z-40
  "
>
  <a
    href={resolve(buildProgramWorkHref(event))}
    aria-label={`Lihat program kerja ${event.title}`}
    class={`
      group relative
      flex h-8 w-full items-center
      overflow-hidden
      border border-blue-950/15
      bg-linear-to-r
      from-[#0d4d88]
      to-[#277bb8]
      text-white
      shadow-sm
      shadow-blue-950/15
      transition duration-200
      hover:-translate-y-px
      hover:from-[#093f73]
      hover:to-[#1e679d]
      hover:shadow-md
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
          ? 'rounded-r-lg pr-10'
          : 'rounded-r-full pr-3'
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
  </a>

  {#if continuesAfter}
    <button
      type="button"
      aria-label={`Lihat tanggal selesai ${event.title}`}
      onclick={onShowEnd}
      class="
        absolute top-0 right-0
        flex h-8 w-9
        items-center justify-center
        rounded-r-lg
        text-white/75
        transition duration-200
        hover:bg-white/15
        hover:text-white
        focus-visible:outline-2
        focus-visible:outline-offset-1
        focus-visible:outline-orange-500
      "
    >
      <svg
        viewBox="0 0 24 24"
        fill="none"
        aria-hidden="true"
        class="
          size-3.5
          shrink-0
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
    </button>
  {/if}

  <div
    role="tooltip"
    class="
      pointer-events-none
      invisible absolute
      bottom-full left-0
      mb-2 w-max max-w-64
      translate-y-1
      rounded-xl
      border border-blue-900/10
      bg-white px-3 py-2
      text-left
      opacity-0
      shadow-lg
      shadow-blue-950/15
      transition duration-150
      group-hover/range:visible
      group-hover/range:translate-y-0
      group-hover/range:opacity-100
      group-focus-within/range:visible
      group-focus-within/range:translate-y-0
      group-focus-within/range:opacity-100
    "
  >
    <p
      class="
        text-[11px] leading-tight
        font-extrabold
        text-[#164f88]
      "
    >
      {event.title}
    </p>

    <p
      class="
        mt-1 text-[10px]
        font-bold
        text-orange-600
      "
    >
      {formatEventRange(event)}
      {#if dayCount > 1}
        · {dayCount} hari
      {/if}
    </p>

    <p
      class="
        mt-0.5 text-[10px]
        font-medium
        text-[#164f88]/70
      "
    >
      {event.ministryName}
    </p>
  </div>
</div>
