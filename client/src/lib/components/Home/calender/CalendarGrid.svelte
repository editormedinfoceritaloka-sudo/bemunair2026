<script lang="ts">
  import { resolve } from '$app/paths';
  import { fade, fly } from 'svelte/transition';
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

  let openDay = $state<{ label: string; events: CalendarEvent[] } | null>(null);

  function openDayList(label: string, dayEvents: CalendarEvent[]): void {
    openDay = { label, events: dayEvents };
  }

  function closeDayList(): void {
    openDay = null;
  }

  function chooseFromList(event: CalendarEvent): void {
    openDay = null;

    setTimeout(() => {
      onEventSelect?.(event);
    }, 280);
  }
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
              max-sm:aspect-auto max-sm:min-h-20
              max-sm:items-center max-sm:pt-1.5
            "
          >
            {#if startingEvents.length && onDateSelect}
              <button
                type="button"
                data-date-chip={key}
                onclick={() => onDateSelect(day)}
                aria-label={`Lihat agenda ${label}`}
                title="Lihat agenda bulan ini"
                class={`flex size-7 shrink-0 items-center justify-center rounded-full text-[11px] font-black transition hover:scale-105 max-sm:hover:scale-100 focus-visible:outline-2 focus-visible:outline-offset-1 focus-visible:outline-orange-500 max-sm:size-8 max-sm:min-h-8 max-sm:min-w-8 max-sm:p-0 max-sm:[font-family:inherit] max-sm:text-[12px]! max-sm:leading-none max-sm:tabular-nums ${selectedDate === key ? 'ring-2 ring-orange-400 ring-offset-1 max-sm:ring-0 max-sm:outline-2 max-sm:outline-orange-400 max-sm:outline-offset-1' : ''} ${today ? 'bg-orange-500 text-white shadow-md' : inMonth ? 'bg-blue-100 text-[#164f88]' : 'text-blue-950/25 max-sm:text-[#164f88]'}`}
              >
                {day.getDate()}
              </button>
            {:else}
              <span
                aria-current={today ? 'date' : undefined}
                class={`flex size-7 shrink-0 items-center justify-center rounded-full text-[11px] font-black max-sm:size-8 max-sm:min-h-8 max-sm:min-w-8 max-sm:p-0 max-sm:[font-family:inherit] max-sm:text-[12px]! max-sm:leading-none max-sm:tabular-nums ${today ? 'bg-orange-500 text-white shadow-md' : inMonth ? 'bg-blue-100 text-[#164f88]' : 'text-blue-950/25 max-sm:text-[#164f88]'}`}
              >
                {day.getDate()}
              </span>
            {/if}

            {#if startingEvents.length}
              <div class={`flex w-full min-w-0 flex-col gap-1 overflow-hidden max-sm:items-center max-sm:overflow-visible ${startingEvents.length > 1 ? 'max-sm:hidden' : ''}`}>
                {#each startingEvents.slice(0, 2) as event (event.id)}
                  <CalendarDayBadge {event} large={startingEvents.length === 1} onSelect={onEventSelect} />
                {/each}
              </div>

              {#if startingEvents.length > 1}
                <button
                  type="button"
                  data-day-list
                  onclick={() => openDayList(label, startingEvents)}
                  aria-label={`Lihat ${startingEvents.length} program kerja pada ${label}`}
                  class="relative z-10 flex min-h-7 min-w-7 items-center justify-center rounded-lg border-2 border-orange-500 bg-white p-1 shadow-[0_2px_8px_rgba(249,115,22,0.35)] sm:hidden"
                >
                  <span class="flex size-6 shrink-0 items-center justify-center rounded-lg bg-orange-100 p-0.5">
                    <img
                      src={startingEvents[0].logo}
                      alt=""
                      draggable="false"
                      loading="lazy"
                      class="h-full w-full object-contain"
                    />
                  </span>

                  <span
                    aria-hidden="true"
                    class="absolute -top-1.5 -right-1.5 flex h-4 min-w-4 items-center justify-center rounded-full border-2 border-white bg-orange-500 px-0.5 text-[8px]! leading-none font-black text-white shadow-sm"
                  >
                    +{startingEvents.length - 1}
                  </span>
                </button>
              {/if}
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
                <span class="mb-1 rounded-full bg-orange-500 px-2 py-0.5 text-[8px] font-black text-white shadow-sm sm:mb-3 sm:px-4 sm:py-1.5 sm:text-xs md:text-sm max-sm:mb-1.5 max-sm:text-[8px]! max-sm:whitespace-nowrap max-sm:px-1.5 max-sm:py-[3px] max-sm:leading-none">
                  <span class="sm:hidden">Detail</span>
                  <span class="hidden sm:inline">Lihat detail</span>
                </span>
              </a>
            {/if}
          </div>
        {/each}
      </div>
    {/each}
  </div>
</div>

{#if openDay}
  <div
    class="fixed inset-0 z-[90] flex items-end justify-center p-4 sm:hidden"
    transition:fade={{ duration: 180 }}
  >
    <button
      type="button"
      aria-label="Tutup daftar program kerja"
      onclick={closeDayList}
      class="absolute inset-0 bg-blue-950/45 backdrop-blur-[2px]"
    ></button>

    <div
      role="dialog"
      aria-modal="true"
      aria-label={`Program kerja ${openDay.label}`}
      class="relative w-full max-w-sm rounded-2xl bg-white p-4 shadow-2xl"
      transition:fly={{ y: 40, duration: 240 }}
    >
      <div class="flex items-start justify-between gap-3">
        <div class="min-w-0">
          <p class="text-[10px] font-black uppercase tracking-[0.18em] text-orange-500">
            {openDay.events.length} program kerja
          </p>
          <p class="mt-0.5 text-sm font-black text-[#164f88]">
            {openDay.label}
          </p>
        </div>

        <button
          type="button"
          onclick={closeDayList}
          aria-label="Tutup"
          class="flex size-8 shrink-0 items-center justify-center rounded-full bg-blue-50 text-lg leading-none font-black text-[#164f88]"
        >
          ×
        </button>
      </div>

      <ul class="mt-3 grid max-h-[50vh] gap-2 overflow-y-auto">
        {#each openDay.events as event (event.id)}
          <li>
            <button
              type="button"
              onclick={() => chooseFromList(event)}
              class="flex w-full items-center gap-3 rounded-xl border border-blue-900/10 bg-blue-50/60 p-2.5 text-left transition active:scale-[0.98] active:bg-orange-50"
            >
              <span class="flex size-9 shrink-0 items-center justify-center rounded-lg border-2 border-orange-500 bg-orange-100 p-1">
                <img
                  src={event.logo}
                  alt=""
                  draggable="false"
                  loading="lazy"
                  class="h-full w-full object-contain"
                />
              </span>

              <span class="min-w-0 flex-1">
                <span class="block truncate text-[13px] font-extrabold text-[#164f88]">
                  {event.title}
                </span>
                <span class="mt-0.5 block truncate text-[11px] font-semibold text-blue-900/55">
                  {event.ministryName}
                </span>
              </span>

              <span aria-hidden="true" class="text-lg font-black text-orange-500">›</span>
            </button>
          </li>
        {/each}
      </ul>
    </div>
  </div>
{/if}

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