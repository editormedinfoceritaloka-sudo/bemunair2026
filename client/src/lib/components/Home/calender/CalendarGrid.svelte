<script lang="ts">
  import { tick } from 'svelte';
  import CalendarRangeBadge from './CalendarRangeBadge.svelte';

  import {
    buildCalendarWeeks,
    dateKey,
    getEventsForMonth,
    isCurrentMonth,
    isToday,
    parseDate
  } from './calendar-utils';

  import type { CalendarEvent } from './types';

  let {
    currentMonth,
    events,
    onShowDate
  }: {
    currentMonth: Date;
    events: CalendarEvent[];
    onShowDate?: (date: Date) => void;
  } = $props();

  let focusedDate = $state<string | null>(null);
  let focusTimer: ReturnType<typeof setTimeout> | undefined;

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

  type WeekSegment = {
    event: CalendarEvent;
    colStart: number;
    span: number;
    lane: number;
    continuesBefore: boolean;
    continuesAfter: boolean;
  };

  const barTop = 48;
  const barGap = 38;

  function eventEndDate(
    event: CalendarEvent
  ): string {
    return event.endDate &&
      event.endDate >= event.startDate
      ? event.endDate
      : event.startDate;
  }

  function buildWeekSegments(
    week: Date[]
  ): {
    segments: WeekSegment[];
    laneCount: number;
  } {
    const weekKeys = week.map(dateKey);
    const weekStart = weekKeys[0];
    const weekEnd = weekKeys[6];

    const weekEvents = events
      .filter(
        (event) =>
          event.startDate <= weekEnd &&
          eventEndDate(event) >= weekStart
      )
      .sort((first, second) => {
        const startComparison =
          first.startDate.localeCompare(
            second.startDate
          );

        if (startComparison !== 0) {
          return startComparison;
        }

        return eventEndDate(second).localeCompare(
          eventEndDate(first)
        );
      });

    const laneEnds: number[] = [];
    const segments: WeekSegment[] = [];

    for (const event of weekEvents) {
      const endDate = eventEndDate(event);

      const colStart =
        event.startDate < weekStart
          ? 0
          : weekKeys.indexOf(event.startDate);

      const colEnd =
        endDate > weekEnd
          ? 6
          : weekKeys.indexOf(endDate);

      let lane = laneEnds.findIndex(
        (laneEnd) => laneEnd < colStart
      );

      if (lane === -1) {
        lane = laneEnds.length;
        laneEnds.push(colEnd);
      } else {
        laneEnds[lane] = colEnd;
      }

      segments.push({
        event,
        colStart,
        span: colEnd - colStart + 1,
        lane,
        continuesBefore:
          event.startDate < weekStart,
        continuesAfter:
          endDate > weekEnd
      });
    }

    return {
      segments,
      laneCount: laneEnds.length
    };
  }

  function segmentStyle(
    segment: WeekSegment
  ): string {
    return `
      left: calc(${segment.colStart} * 100% / 7 + 6px);
      width: calc(${segment.span} * 100% / 7 - 12px);
      top: ${barTop + segment.lane * barGap}px;
    `;
  }

  function rowMinHeight(
    laneCount: number
  ): number {
    return Math.max(
      176,
      barTop + laneCount * barGap + 12
    );
  }

  async function showEndDate(
    event: CalendarEvent
  ): Promise<void> {
    const endKey = eventEndDate(event);
    const endDate = parseDate(endKey);

    clearTimeout(focusTimer);
    focusedDate = null;

    await tick();

    onShowDate?.(endDate);
    focusedDate = endKey;

    setTimeout(() => {
      document
        .querySelector(`[data-date="${endKey}"]`)
        ?.scrollIntoView({
          behavior: 'smooth',
          block: 'center'
        });
    }, 550);

    focusTimer = setTimeout(() => {
      focusedDate = null;
    }, 3600);
  }

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
    mt-6 overflow-hidden
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
      class="min-w-195"
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
        {#each weekDayLabels as label, index (index)}
          <div
            role="columnheader"
            class={`
              border-r
              border-white/15
              px-2.5 py-3
              text-center
              text-[11px] font-black
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
        {@const weekLayout =
          buildWeekSegments(week)}

        <div
          role="row"
          style={`min-height: ${rowMinHeight(weekLayout.laneCount)}px;`}
          class="
            relative grid
            grid-cols-7
            border-b
            border-blue-950/10
            last:border-b-0
          "
        >
          {#each week as day, dayIndex (dateKey(day))}
            {@const currentMonthDay =
              isCurrentMonth(
                day,
                currentMonth
              )}

            <div
              role="gridcell"
              aria-label={dayAriaLabel(day)}
              data-date={dateKey(day)}
              class={`
                min-w-0
                border-r
                border-blue-950/10
                p-2
                last:border-r-0

                ${
                  currentMonthDay
                    ? dayIndex >= 5
                      ? 'bg-blue-50'
                      : 'bg-white'
                    : 'bg-slate-100'
                }

                ${
                  focusedDate === dateKey(day)
                    ? 'calendar-end-focus'
                    : ''
                }
              `}
            >
              <div
                class="
                  flex items-center
                  justify-between gap-1.5
                "
              >
                <span
                  class={`
                    flex size-7
                    items-center
                    justify-center
                    rounded-xl
                    text-[11px] font-black

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
            </div>
          {/each}

          {#each weekLayout.segments as segment (segment.event.id)}
            <CalendarRangeBadge
              event={segment.event}
              logo={segment.event.logo}
              positionStyle={segmentStyle(segment)}
              continuesBefore={segment.continuesBefore}
              continuesAfter={segment.continuesAfter}
              onShowEnd={() => showEndDate(segment.event)}
            />
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
        px-5 py-6
        text-center
      "
    >
      <p
        class="
          text-xs font-bold
          text-[#164f88]/65
        "
      >
        Belum ada program kerja yang dimulai pada bulan ini.
      </p>
    </div>
  {/if}
</div>

<style>
  :global(.calendar-end-focus) {
    animation: calendar-end-focus 0.9s ease-in-out 4;
  }

  @keyframes -global-calendar-end-focus {
    0%,
    100% {
      background-color: #ffffff;
      box-shadow: inset 0 0 0 0 rgba(244, 123, 32, 0);
    }

    50% {
      background-color: #fef0e3;
      box-shadow:
        inset 0 0 0 3px #f47b20,
        0 0 18px rgba(244, 123, 32, 0.55);
    }
  }

  @media (prefers-reduced-motion: reduce) {
    :global(.calendar-end-focus) {
      animation: none;
      box-shadow: inset 0 0 0 3px #f47b20;
    }
  }
</style>