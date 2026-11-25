<script lang="ts">
  import { resolve } from '$app/paths';
  import { buildProgramWorkHref, formatEventRange } from './calendar-utils';
  import type { CalendarEvent } from './types';

  let {
    event,
    large = false,
    onSelect
  }: {
    event: CalendarEvent;
    large?: boolean;
    onSelect?: (event: CalendarEvent) => void;
  } = $props();

  const label = $derived(`Lihat program kerja ${event.title}, ${formatEventRange(event)}, dari ${event.ministryName}`);
  const title = $derived(`${event.title} — ${formatEventRange(event)} — ${event.ministryName}`);
  const iconClass = $derived(`flex ${large ? 'size-8 sm:size-5' : 'size-6 sm:size-5'} shrink-0 items-center justify-center rounded-lg bg-orange-100 p-0.5 sm:bg-blue-50`);
  const className = $derived(`
    group relative z-10 flex min-h-7 min-w-7 max-w-full
    items-center justify-center gap-1.5 rounded-lg
    border-2 border-orange-500
    bg-white p-1
    shadow-[0_2px_8px_rgba(249,115,22,0.35)]
    transition duration-200
    hover:-translate-y-px
    hover:border-orange-600
    hover:shadow-md
    focus-visible:outline-2
    focus-visible:outline-offset-1
    focus-visible:outline-orange-500
    sm:border sm:border-blue-900/10 sm:bg-white sm:p-0.5
    sm:shadow-sm sm:min-w-0 sm:justify-start sm:px-1.5
    ${large ? 'min-h-9 min-w-9 p-1.5 sm:min-h-7 sm:min-w-7 sm:p-0.5' : ''}
  `);
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
    <span class={iconClass}>
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
    <span class={iconClass}>
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
