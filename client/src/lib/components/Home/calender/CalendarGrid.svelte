<script lang="ts">
  import CalendarDayBadge from './CalendarDayBadge.svelte';
  import CalendarRangeBadge from './CalendarRangeBadge.svelte';

  import {
    buildCalendarWeeks,
    buildWeekSegments,
    dateKey,
    getSingleDayEvents,
    isCurrentMonth,
    isToday
  } from './calendar-utils';

  import type {
    CalendarEvent,
    WeekView
  } from './types';

  let {
    currentMonth,
    events,
    logo
  }: {
    currentMonth: Date;
    events: CalendarEvent[];
    logo: string;
  } = $props();

  const weekDayLabels = [
    'Mon',
    'Tue',
    'Wed',
    'Thu',
    'Fri',
    'Sat',
    'Sun'
  ];

  const weekViews = $derived.by((): WeekView[] => {
    return buildCalendarWeeks(currentMonth).map(
      (days) => {
        const segments = buildWeekSegments(
          days,
          events
        );

        const laneCount =
          segments.length > 0
            ? Math.max(
                ...segments.map(
                  (segment) => segment.lane
                )
              ) + 1
            : 0;

        const maximumSingleEvents = Math.max(
          ...days.map(
            (day) =>
              getSingleDayEvents(day, events).length
          ),
          0
        );

        const rangeTop =
          48 + maximumSingleEvents * 28;

        const contentHeight =
          rangeTop +
          Math.max(laneCount, 1) * 31 +
          12;

        return {
          days,
          segments,
          rangeTop,
          height: Math.max(116, contentHeight)
        };
      }
    );
  });
</script>

<div
  class="
    mt-10 overflow-hidden
    rounded-[28px]
    border border-white/70
    bg-white/25
    shadow-sm
    backdrop-blur-md
  "
>
  <div class="overflow-x-auto">
    <div class="min-w-[900px]">
      <div
        class="
          grid grid-cols-7
          border-b border-blue-900/15
          bg-gradient-to-r
          from-[#327fbd]
          via-[#5ba0d5]
          to-[#327fbd]
        "
      >
        {#each weekDayLabels as label}
          <div
            class="
              border-r border-white/30
              px-2 py-3 text-center
              text-xs font-black uppercase
              tracking-[0.12em]
              text-white
              last:border-r-0
              sm:text-sm
            "
          >
            {label}
          </div>
        {/each}
      </div>

      {#each weekViews as week (dateKey(week.days[0]))}
        <div
          class="
            relative grid grid-cols-7
            border-b border-blue-900/15
            last:border-b-0
          "
          style={`height: ${week.height}px;`}
        >
          {#each week.days as day (dateKey(day))}
            {@const singleDayEvents =
              getSingleDayEvents(day, events)}

            <div
              class={`
                relative
                border-r border-blue-800/20
                px-2 py-2
                last:border-r-0
                ${
                  isCurrentMonth(day, currentMonth)
                    ? 'bg-white/25'
                    : 'bg-blue-300/15'
                }
              `}
            >
              <span
                class={`
                  flex size-7
                  items-center justify-center
                  rounded-full
                  text-xs font-extrabold
                  ${
                    isToday(day)
                      ? 'bg-orange-500 text-white shadow-sm'
                      : isCurrentMonth(
                            day,
                            currentMonth
                          )
                        ? 'text-blue-950'
                        : 'text-blue-800/35'
                  }
                `}
              >
                {day.getDate()}
              </span>

              {#if singleDayEvents.length > 0}
                <div class="mt-2 flex flex-col gap-1">
                  {#each singleDayEvents as event (event.id)}
                    <CalendarDayBadge
                      {event}
                      {logo}
                    />
                  {/each}
                </div>
              {/if}
            </div>
          {/each}

          {#each week.segments as segment (`${dateKey(week.days[0])}-${segment.event.id}`)}
            <CalendarRangeBadge
              event={segment.event}
              {logo}
              continuesBefore={segment.continuesBefore}
              continuesAfter={segment.continuesAfter}
              positionStyle={`
                left: calc(${(segment.startColumn / 7) * 100}% + 5px);
                width: calc(${(segment.span / 7) * 100}% - 10px);
                top: ${week.rangeTop + segment.lane * 31}px;
              `}
            />
          {/each}
        </div>
      {/each}
    </div>
  </div>
</div>
