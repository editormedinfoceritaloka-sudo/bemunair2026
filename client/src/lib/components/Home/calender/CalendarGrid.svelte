<script lang="ts">
  import CalendarDayBadge from './CalendarDayBadge.svelte';
  import {
    buildCalendarWeeks,
    dateKey,
    getEventsForDay,
    isCurrentMonth,
    isToday
  } from './calendar-utils';
  import type { CalendarEvent } from './types';

  let {
    currentMonth,
    events
  }: {
    currentMonth: Date;
    events: CalendarEvent[];
  } = $props();

  const weekDayLabels = ['Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu', 'Minggu'];
  const dateFormatter = new Intl.DateTimeFormat('id-ID', { dateStyle: 'full' });
  const weeks = $derived(buildCalendarWeeks(currentMonth));
</script>

<div class="mt-6 overflow-hidden rounded-2xl border border-blue-900/15 bg-white shadow-lg sm:rounded-3xl">
  <div role="grid" aria-label="Kalender program kerja Kabinet Cerita Loka">
    <div role="row" class="grid grid-cols-7 bg-[#164f88]">
      {#each weekDayLabels as label, index (label)}
        <div
          role="columnheader"
          aria-label={label}
          class:bg-[#0f4578]={index >= 5}
          class="border-r border-white/15 px-1 py-3 text-center text-[10px] font-black uppercase text-white last:border-r-0 sm:text-xs"
        >
          <span class="sm:hidden">{label.slice(0, 3)}</span>
          <span class="hidden sm:inline">{label}</span>
        </div>
      {/each}
    </div>

    {#each weeks as week (dateKey(week[0]))}
      <div role="row" class="grid grid-cols-7 border-b border-blue-950/10 last:border-b-0">
        {#each week as day, dayIndex (dateKey(day))}
          {@const key = dateKey(day)}
          {@const inMonth = isCurrentMonth(day, currentMonth)}
          {@const today = isToday(day)}
          {@const startingEvents = inMonth ? getEventsForDay(day, events) : []}
          {@const label = dateFormatter.format(day)}
          <div
            role="gridcell"
            aria-label={label}
            data-date={key}
            class:bg-slate-100={!inMonth}
            class:bg-blue-50={inMonth && dayIndex >= 5}
            class="min-h-14 min-w-0 border-r border-blue-950/10 p-1 last:border-r-0 sm:p-1.5"
          >
            <span
              aria-current={today ? 'date' : undefined}
              class={`mx-auto flex size-6 items-center justify-center rounded-md text-xs font-bold sm:mx-0 sm:text-sm ${today ? 'bg-orange-500 text-white' : inMonth ? 'text-[#164f88]' : 'text-blue-950/35'}`}
            >
              {day.getDate()}
            </span>
            {#if startingEvents.length}
              <div class="mt-1 flex flex-wrap justify-center gap-1 sm:flex-col">
                {#each startingEvents as event (event.id)}
                  <CalendarDayBadge {event} />
                {/each}
              </div>
            {/if}
          </div>
        {/each}
      </div>
    {/each}
  </div>
</div>
