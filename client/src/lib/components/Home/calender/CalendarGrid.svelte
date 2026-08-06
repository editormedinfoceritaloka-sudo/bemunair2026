<script lang="ts">
  import CalendarDayBadge from './CalendarDayBadge.svelte';

  import {
    buildCalendarWeeks,
    dateKey,
    getEventsForDay,
    getEventsForMonth,
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

  const weekDayLabels = [
    'Senin',
    'Selasa',
    'Rabu',
    'Kamis',
    'Jumat',
    'Sabtu',
    'Minggu'
  ];

  const weeks = $derived(
    buildCalendarWeeks(
      currentMonth
    )
  );

  const monthEvents = $derived(
    getEventsForMonth(
      currentMonth,
      events
    )
  );

  function dayAriaLabel(
    date: Date
  ): string {
    return new Intl.DateTimeFormat(
      'id-ID',
      {
        weekday: 'long',
        day: 'numeric',
        month: 'long',
        year: 'numeric'
      }
    ).format(date);
  }
</script>

<div
  class="
    mt-8 overflow-hidden
    rounded-3xl
    border border-blue-900/15
    bg-white
    shadow-xl
    shadow-blue-950/10
  "
>
  <div
    class="
      overflow-x-auto
      overscroll-x-contain
    "
  >
    <div
      role="grid"
      aria-label="Kalender program kerja Kabinet Cerita Loka"
      class="min-w-245"
    >
      <div
        role="row"
        class="
          grid grid-cols-7
          border-b
          border-blue-950/20
          bg-[#164f88]
        "
      >
        {#each weekDayLabels as label, index}
          <div
            role="columnheader"
            class={`
              border-r
              border-white/15
              px-3 py-4
              text-center
              text-xs font-black
              uppercase
              tracking-[0.1em]
              text-white
              last:border-r-0

              ${
                index >= 5
                  ? 'bg-[#0f4578]'
                  : ''
              }
            `}
          >
            {label}
          </div>
        {/each}
      </div>

      {#each weeks as week (dateKey(week[0]))}
        <div
          role="row"
          class="
            grid min-h-54
            grid-cols-7
            border-b
            border-blue-950/10
            last:border-b-0
          "
        >
          {#each week as day, dayIndex (dateKey(day))}
            {@const dayEvents =
              getEventsForDay(
                day,
                events
              )}

            {@const currentMonthDay =
              isCurrentMonth(
                day,
                currentMonth
              )}

            <div
              role="gridcell"
              aria-label={dayAriaLabel(day)}
              class={`
                min-w-0
                border-r
                border-blue-950/10
                p-2.5
                last:border-r-0

                ${
                  currentMonthDay
                    ? dayIndex >= 5
                      ? 'bg-blue-50'
                      : 'bg-white'
                    : 'bg-slate-100'
                }
              `}
            >
              <div
                class="
                  flex items-center
                  justify-between gap-2
                "
              >
                <span
                  class={`
                    flex size-8
                    items-center
                    justify-center
                    rounded-xl
                    text-xs font-black

                    ${
                      isToday(day)
                        ? 'bg-orange-500 text-white shadow-md'
                        : currentMonthDay
                          ? 'bg-blue-100 text-[#164f88]'
                          : 'text-blue-950/25'
                    }
                  `}
                >
                  {day.getDate()}
                </span>
              </div>

              {#if dayEvents.length > 0}
                <div
                  class="
                    mt-2.5 flex
                    max-h-40
                    flex-col gap-1.5
                    overflow-y-auto
                    pr-1
                    [scrollbar-color:#164f88_transparent]
                    [scrollbar-width:thin]
                  "
                >
                  {#each dayEvents as event (event.id)}
                    <CalendarDayBadge
                      {event}
                    />
                  {/each}
                </div>
              {/if}
            </div>
          {/each}
        </div>
      {/each}
    </div>
  </div>

  {#if monthEvents.length === 0}
    <div
      class="
        border-t
        border-blue-950/10
        bg-blue-50
        px-6 py-8
        text-center
      "
    >
      <p
        class="
          text-sm font-bold
          text-[#164f88]/65
        "
      >
        Belum ada program kerja yang dimulai pada bulan ini.
      </p>
    </div>
  {/if}
</div>