<script lang="ts">
  import { buildProgramWorkHref } from './calendar-utils';
  import type { CalendarEvent } from './types';

  let {
    event,
    logo
  }: {
    event: CalendarEvent;
    logo: string;
  } = $props();
</script>

<a
  href={buildProgramWorkHref(event)}
  aria-label={`Lihat program kerja ${event.title}`}
  title={event.title}
  class="
    flex min-w-0 items-center gap-1.5
    rounded-full
    bg-gradient-to-r
    from-orange-500 to-orange-400
    px-2 py-1
    text-[9px] leading-none
    font-extrabold text-white
    shadow-sm
    transition duration-200
    hover:-translate-y-px
    hover:from-orange-600
    hover:to-orange-500
  "
>
  <span
    class="
      flex size-4 shrink-0
      items-center justify-center
      rounded-full bg-white
    "
  >
    <img
      src={logo}
      alt=""
      draggable="false"
      class="size-3.5 object-contain"
    />
  </span>

  <span class="truncate text-white">
    {event.title}
  </span>
</a>
