<script lang="ts">
  import { resolve } from '$app/paths';
  import { buildProgramWorkHref, formatEventRange } from './calendar-utils';
  import type { CalendarEvent } from './types';

  let {
    event
  }: {
    event: CalendarEvent;
  } = $props();
</script>

<a
  href={resolve(buildProgramWorkHref(event))}
  aria-label={`Lihat program kerja ${event.title}, ${formatEventRange(event)}, dari ${event.ministryName}`}
  title={`${event.title} — ${formatEventRange(event)} — ${event.ministryName}`}
  class="
    group flex min-h-8 min-w-7 max-w-full
    items-center justify-center gap-1.5 rounded-lg
    border border-blue-900/10
    bg-white p-0.5
    shadow-sm
    transition duration-200
    hover:-translate-y-px
    hover:border-orange-400
    hover:shadow-md
    focus-visible:outline-2
    focus-visible:outline-offset-1
    focus-visible:outline-orange-500
    sm:min-w-0 sm:justify-start sm:px-1.5
  "
>
  <span
    class="
      flex size-5 shrink-0
      items-center justify-center
      rounded-lg bg-blue-50
      p-0.5
    "
  >
    <img
      src={event.logo}
      alt={event.logoAlt}
      draggable="false"
      loading="lazy"
      class="
        h-full w-full
        object-contain
      "
    />
  </span>

  <span
    class="
      hidden min-w-0 truncate sm:block
      text-[10px] leading-tight
      font-extrabold
      text-[#164f88]
      transition-colors
      group-hover:text-orange-600
      sm:text-[11px]
    "
  >
    {event.title}
  </span>
</a>
