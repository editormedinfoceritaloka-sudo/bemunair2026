<script lang="ts">
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
    events,
    selectedDate,
    onSelectDate
  }: {
    currentMonth: Date;
    events: CalendarEvent[];
    selectedDate: string | null;
    onSelectDate: (day: Date) => void;
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
          {@const count = inMonth ? getEventsForDay(day, events).length : 0}
          {@const label = dateFormatter.format(day)}
          <div
            role="gridcell"
            aria-label={label}
            data-date={key}
            class:bg-slate-100={!inMonth}
            class:bg-blue-50={inMonth && dayIndex >= 5}
            class="min-w-0 border-r border-blue-950/10 p-1 last:border-r-0 sm:p-1.5"
          >
            {#if count}
              <button
                type="button"
                aria-label={`${label}, ${count} program mulai`}
                aria-pressed={selectedDate === key}
                aria-current={today ? 'date' : undefined}
                onclick={() => onSelectDate(day)}
                class="flex min-h-12 w-full flex-col items-center justify-center gap-0.5 rounded-lg text-[#164f88] hover:bg-blue-100 focus-visible:outline-2 focus-visible:-outline-offset-2 focus-visible:outline-blue-700 aria-pressed:bg-blue-100"
              >
                <span class:bg-orange-500={today} class:text-white={today} class="rounded-md px-1.5 text-xs font-black sm:text-sm">{day.getDate()}</span>
                <span class="rounded-full bg-[#164f88] px-1.5 text-[10px] leading-4 font-bold text-white">{count}</span>
              </button>
            {:else}
              <span
                aria-current={today ? 'date' : undefined}
                class={`flex min-h-12 items-center justify-center text-xs font-semibold sm:text-sm ${inMonth ? 'text-[#164f88]' : 'text-blue-950/35'}`}
              >
                <span class:bg-orange-500={today} class:text-white={today} class="rounded-md px-1.5">{day.getDate()}</span>
              </span>
            {/if}
          </div>
        {/each}
      </div>
    {/each}
  </div>
</div>
