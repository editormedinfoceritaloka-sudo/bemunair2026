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
    absolute z-20
    flex h-7 items-center
    overflow-hidden
    bg-gradient-to-r
    from-[#0d4d88] to-[#1668aa]
    pr-3 text-white
    shadow-sm
    transition duration-200
    hover:z-30
    hover:-translate-y-px
    hover:from-[#093f73]
    hover:to-[#125b97]
    ${
      continuesBefore
        ? 'rounded-l-md'
        : 'rounded-l-full'
    }
    ${
      continuesAfter
        ? 'rounded-r-md'
        : 'rounded-r-full'
    }
  `}
>
  {#if !continuesBefore}
    <span
      class="
        ml-0.5 flex size-6 shrink-0
        items-center justify-center
        rounded-full bg-white
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
      truncate px-2
      text-[10px] font-extrabold
      text-white sm:text-xs
    "
  >
    {event.title}
  </span>
</a>
