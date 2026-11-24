<script lang="ts">
  import { resolve } from '$app/paths';
  import { buildProgramWorkHref, formatEventRange } from './calendar-utils';
  import type { CalendarEvent } from './types';

  let {
    event,
    onSelect
  }: {
    event: CalendarEvent;
    onSelect?: (event: CalendarEvent) => void;
  } = $props();

  const label = $derived(`Lihat program kerja ${event.title}, ${formatEventRange(event)}, dari ${event.ministryName}`);
  const title = $derived(`${event.title} — ${formatEventRange(event)} — ${event.ministryName}`);
  const className = `
    group flex min-h-7 min-w-7 max-w-full
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
  `;
</script>

{#if onSelect}
  <button
    type="button"
    data-event-start={event.id}
    onclick={() => onSelect(event)}
    aria-label={label}
    {title}
    class={className}
  >
    <span class="flex size-5 shrink-0 items-center justify-center rounded-lg bg-blue-50 p-0.5">
      <img
        src={event.logo}
        alt={event.logoAlt}
        draggable="false"
        loading="lazy"
        class="h-full w-full object-contain"
      />
    </span>
    <span class="hidden min-w-0 truncate text-[10px] leading-tight font-extrabold text-[#164f88] transition-colors group-hover:text-orange-600 sm:block sm:text-[11px]">
      {event.title}
    </span>
  </button>
{:else}
  <a
    href={resolve(buildProgramWorkHref(event))}
    aria-label={label}
    {title}
    class={className}
  >
    <span class="flex size-5 shrink-0 items-center justify-center rounded-lg bg-blue-50 p-0.5">
      <img
        src={event.logo}
        alt={event.logoAlt}
        draggable="false"
        loading="lazy"
        class="h-full w-full object-contain"
      />
    </span>
    <span class="hidden min-w-0 truncate text-[10px] leading-tight font-extrabold text-[#164f88] transition-colors group-hover:text-orange-600 sm:block sm:text-[11px]">
      {event.title}
    </span>
  </a>
{/if}
