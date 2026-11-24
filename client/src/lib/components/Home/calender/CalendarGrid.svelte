<script lang="ts">
  import { resolve } from '$app/paths';
  import CalendarDayBadge from './CalendarDayBadge.svelte';
  import {
    buildCalendarWeeks,
    buildProgramWorkHref,
    dateKey,
    getEventsForDay,
    isCurrentMonth,
    isToday
  } from './calendar-utils';
  import type { CalendarEvent } from './types';

  let {
    currentMonth,
    events,
    flashDate = null,
    selectedDate = null,
    targetEvent = null,
    onDateSelect,
    onEventSelect
  }: {
    currentMonth: Date;
    events: CalendarEvent[];
    flashDate?: string | null;
    selectedDate?: string | null;
    targetEvent?: CalendarEvent | null;
    onDateSelect?: (date: Date) => void;
    onEventSelect?: (event: CalendarEvent) => void;
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
          {@const isTarget = targetEvent ? (targetEvent.endDate ?? targetEvent.startDate) === key : false}
          <div
            role="gridcell"
            aria-label={`${label}${startingEvents.length ? `, ${startingEvents.map((event) => event.title).join(', ')}` : ''}`}
            data-date={key}
            class:bg-slate-100={!inMonth}
            class:bg-blue-50={inMonth && dayIndex >= 5}
            class:calendar-cell-flash={flashDate === key}
            class:z-10={flashDate === key}
            class="
              relative flex aspect-square min-w-0
              flex-col items-start justify-start
              gap-1 overflow-hidden
              border-r border-blue-950/10 p-1
              last:border-r-0 sm:p-2
            "
          >
            {#if startingEvents.length && onDateSelect}
              <button
                type="button"
                data-date-chip={key}
                onclick={() => onDateSelect(day)}
                aria-label={`Lihat agenda ${label}`}
                title="Lihat agenda bulan ini"
                class={`flex size-7 shrink-0 items-center justify-center rounded-full text-[11px] font-black transition hover:scale-105 focus-visible:outline-2 focus-visible:outline-offset-1 focus-visible:outline-orange-500 ${selectedDate === key ? 'ring-2 ring-orange-400 ring-offset-1' : ''} ${today ? 'bg-orange-500 text-white shadow-md' : inMonth ? 'bg-blue-100 text-[#164f88]' : 'text-blue-950/25'}`}
              >
                {day.getDate()}
              </button>
            {:else}
              <span
                aria-current={today ? 'date' : undefined}
                class={`flex size-7 shrink-0 items-center justify-center rounded-full text-[11px] font-black ${today ? 'bg-orange-500 text-white shadow-md' : inMonth ? 'bg-blue-100 text-[#164f88]' : 'text-blue-950/25'}`}
              >
                {day.getDate()}
              </span>
            {/if}

            {#if startingEvents.length}
              <div class="flex w-full min-w-0 flex-col gap-1 overflow-hidden">
                {#each startingEvents.slice(0, 2) as event (event.id)}
                  <CalendarDayBadge {event} onSelect={onEventSelect} />
                {/each}
              </div>
            {/if}

            {#if isTarget && targetEvent}
              <a
                href={resolve(buildProgramWorkHref(targetEvent))}
                aria-label={`Buka detail program kerja ${targetEvent.title}`}
                title={`Buka detail ${targetEvent.title}`}
                class="
                  absolute inset-0 z-20
                  flex items-end justify-center
                  bg-orange-500/10
                  outline-2 -outline-offset-2 outline-orange-500
                "
              >
                <span class="mb-1 rounded-full bg-orange-500 px-2 py-0.5 text-[8px] font-black text-white shadow-sm sm:text-[9px]">
                  Lihat detail
                </span>
              </a>
            {/if}
          </div>
        {/each}
      </div>
    {/each}
  </div>
</div>

<style>
  .calendar-cell-flash {
    animation: calendar-cell-flash 0.85s ease-out;
  }

  @keyframes calendar-cell-flash {
    0%,
    100% {
      box-shadow: 0 0 0 0 rgba(249, 115, 22, 0);
      background-color: transparent;
    }

    35% {
      box-shadow:
        0 0 0 6px rgba(249, 115, 22, 0.9),
        0 0 28px 8px rgba(249, 115, 22, 0.6);
      background-color: #ffedd5;
    }
  }
</style>
