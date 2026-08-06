<script lang="ts">
  import { buildProgramWorkHref } from './calendar-utils';
  import type { CalendarEvent } from './types';

  let {
    event
  }: {
    event: CalendarEvent;
  } = $props();
</script>

<a
  href={buildProgramWorkHref(event)}
  aria-label={`Lihat program kerja ${event.title} dari ${event.ministryName}`}
  title={`${event.title} — ${event.ministryName}`}
  class="
    group flex min-h-9
    min-w-0 items-center
    gap-2 rounded-xl
    border border-blue-900/10
    bg-white px-2 py-1.5
    shadow-sm
    transition duration-200
    hover:-translate-y-px
    hover:border-orange-400
    hover:shadow-md
    focus-visible:outline-2
    focus-visible:outline-offset-1
    focus-visible:outline-orange-500
  "
>
  <span
    class="
      flex size-6 shrink-0
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
      min-w-0 truncate
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