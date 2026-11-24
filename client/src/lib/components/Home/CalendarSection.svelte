<script lang="ts">
  import {
    onDestroy,
    onMount,
    tick,
    untrack
  } from 'svelte';
  import { resolve } from '$app/paths';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import CalendarGrid from './calender/CalendarGrid.svelte';
  import CalendarNavigation from './calender/CalendarNavigation.svelte';

  import {
    addDays,
    addMonths,
    buildProgramWorkHref,
    dateKey,
    formatEventRange,
    monthLabel,
    parseDate,
    startOfMonth
  } from './calender/calendar-utils';

  import type { CalendarEvent } from './calender/types';

  import type {
    Cabinet,
    OrganizationUnit,
    WorkProgram
  } from '$lib/types';

  function eventTargetKey(
    event: CalendarEvent
  ): string {
    return event.endDate ?? event.startDate;
  }

  function buildDateKeys(
    startKey: string,
    endKey: string
  ): string[] {
    const keys: string[] = [];
    let cursor = parseDate(startKey);
    const last = parseDate(endKey);

    while (
      cursor <= last &&
      keys.length < 62
    ) {
      keys.push(dateKey(cursor));
      cursor = addDays(cursor, 1);
    }

    return keys;
  }

  function programToEvent(
    program: WorkProgram,
    unit: OrganizationUnit,
    fallbackLogo: string
  ): CalendarEvent | null {
    if (
      !program.start_date ||
      program.is_published === false
    ) {
      return null;
    }

    const ministryName =
      unit.short_name ??
      unit.name;

    return {
      id: `program-${program.id}`,
      title: program.name,
      startDate:
        program.start_date.slice(
          0,
          10
        ),
      endDate:
        program.end_date?.slice(
          0,
          10
        ),
      ministryName,
      ministrySlug: unit.slug,
      programSlug: program.slug,
      logo:
        unit.logo?.url ??
        unit.logo?.thumbnail_url ??
        fallbackLogo,
      logoAlt:
        unit.logo?.alt_text ??
        `Logo ${unit.name}`
    };
  }

  function addUnitPrograms(
    result: CalendarEvent[],
    seen: Set<string>,
    unit: OrganizationUnit,
    fallbackLogo: string
  ): void {
    if (
      unit.is_active === false ||
      unit.is_published === false
    ) {
      return;
    }

    for (
      const program of
      unit.programs ?? []
    ) {
      const event = programToEvent(
        program,
        unit,
        fallbackLogo
      );

      if (
        !event ||
        seen.has(event.id)
      ) {
        continue;
      }

      seen.add(event.id);
      result.push(event);
    }

    for (
      const child of
      unit.children ?? []
    ) {
      addUnitPrograms(
        result,
        seen,
        child,
        fallbackLogo
      );
    }
  }

  function cabinetToEvents(
    cabinet: Cabinet,
    fallbackLogo: string
  ): CalendarEvent[] {
    const result: CalendarEvent[] = [];
    const seen = new Set<string>();

    for (
      const unit of
      cabinet.kemenkoan ?? []
    ) {
      addUnitPrograms(
        result,
        seen,
        unit,
        fallbackLogo
      );
    }

    return result.sort(
      (first, second) => {
        const dateComparison =
          first.startDate.localeCompare(
            second.startDate
          );

        if (
          dateComparison !== 0
        ) {
          return dateComparison;
        }

        return first.title.localeCompare(
          second.title,
          'id-ID'
        );
      }
    );
  }

  let {
    events = [],
    logo = '/logo/logo-kabinet.webp',
    initialMonth,
    cabinet = null
  }: {
    events?: CalendarEvent[];
    logo?: string;
    initialMonth?: string;
    cabinet?: Cabinet | null;
  } = $props();

  const resolvedEvents =
    $derived.by<CalendarEvent[]>(() => {
      if (cabinet) {
        return cabinetToEvents(
          cabinet,
          logo
        );
      }

      return [...events].sort(
        (first, second) =>
          first.startDate.localeCompare(
            second.startDate
          )
      );
    });

  let currentMonth = $state(untrack(() =>
    startOfMonth(
      initialMonth
        ? parseDate(initialMonth)
        : new Date()
    )
  ));

  const monthKey = $derived(
    dateKey(currentMonth).slice(0, 7)
  );

  const monthAgenda = $derived.by(() =>
    resolvedEvents
      .filter((event) => {
        const start = event.startDate.slice(0, 7);
        const end = eventTargetKey(event).slice(0, 7);

        return (
          start <= monthKey && end >= monthKey
        );
      })
      .sort((first, second) => {
        const dateComparison =
          eventTargetKey(first).localeCompare(
            eventTargetKey(second)
          );

        return (
          dateComparison ||
          first.title.localeCompare(
            second.title,
            'id-ID'
          )
        );
      })
  );

  let flashDate = $state<string | null>(null);
  let selectedDate = $state<string | null>(null);
  let targetEvent = $state<CalendarEvent | null>(
    null
  );
  let animatingEventId = $state<string | null>(
    null
  );
  let barState = $state<{
    x: number;
    y: number;
    width: number;
    height: number;
    points: { x: number; y: number }[];
    event: CalendarEvent;
  } | null>(null);
  let barElement = $state<
    SVGPolylineElement | undefined
  >(undefined);

  let barTween: gsap.core.Animation | undefined;
  let flashTimer:
    | ReturnType<typeof setTimeout>
    | undefined;

  let sectionElement = $state<
    HTMLElement | undefined
  >(undefined);

  let headingElement = $state<
    HTMLHeadingElement | undefined
  >(undefined);

  let navigationElement = $state<
    HTMLDivElement | undefined
  >(undefined);

  let gridElement = $state<
    HTMLDivElement | undefined
  >(undefined);

  function changeMonth(
    amount: number
  ): void {
    const grid = gridElement;

    if (!grid) {
      currentMonth = addMonths(
        currentMonth,
        amount
      );

      return;
    }

    const reduceMotion =
      window.matchMedia(
        '(prefers-reduced-motion: reduce)'
      ).matches;

    if (reduceMotion) {
      currentMonth = addMonths(
        currentMonth,
        amount
      );

      return;
    }

    gsap.killTweensOf(grid);

    gsap.to(grid, {
      x: amount > 0 ? -30 : 30,
      opacity: 0,
      duration: 0.18,
      ease: 'power2.in',
      onComplete: () => {
        currentMonth = addMonths(
          currentMonth,
          amount
        );

        gsap.fromTo(
          grid,
          {
            x: amount > 0 ? 30 : -30,
            opacity: 0
          },
          {
            x: 0,
            opacity: 1,
            duration: 0.3,
            ease: 'power2.out',
            clearProps:
              'transform,opacity'
          }
        );
      }
    });
  }

  function showPreviousMonth(): void {
    changeMonth(-1);
  }

  function showNextMonth(): void {
    changeMonth(1);
  }

  function selectDate(date: Date): void {
    selectedDate = dateKey(date);
    targetEvent = null;
    flashDate = null;

    if (
      dateKey(currentMonth).slice(0, 7) !==
      selectedDate.slice(0, 7)
    ) {
      currentMonth = startOfMonth(date);
    }
  }

  async function showEventOnCalendar(
    event: CalendarEvent
  ): Promise<void> {
    const targetKey = eventTargetKey(event);
    const targetDate = parseDate(targetKey);
    const reduceMotion =
      window.matchMedia(
        '(prefers-reduced-motion: reduce)'
      ).matches;

    barTween?.kill();
    clearTimeout(flashTimer);
    flashDate = null;
    targetEvent = null;
    barState = null;
    animatingEventId = event.id;

    const section = sectionElement;

    if (!section) {
      animatingEventId = null;

      return;
    }

    const source =
      section.querySelector<HTMLElement>(
        `[data-event-start="${event.id}"]`
      ) ??
      section.querySelector<HTMLElement>(
        `[data-agenda-id="${event.id}"] button`
      );

    if (!source) {
      animatingEventId = null;

      return;
    }

    const sectionRect =
      section.getBoundingClientRect();
    const sourceRect =
      source.getBoundingClientRect();
    const origin = {
      x:
        sourceRect.left +
        sourceRect.width / 2 -
        sectionRect.left,
      y:
        sourceRect.top +
        sourceRect.height / 2 -
        sectionRect.top
    };

    if (
      dateKey(currentMonth).slice(0, 7) !==
      targetKey.slice(0, 7)
    ) {
      currentMonth = startOfMonth(targetDate);
    }

    if (reduceMotion) {
      flashDate = targetKey;
      targetEvent = event;
      flashTimer = setTimeout(() => {
        flashDate = null;
        targetEvent = null;
        animatingEventId = null;
      }, 1400);

      return;
    }

    await tick();

    const cell =
      section.querySelector<HTMLElement>(
        `[data-date="${targetKey}"]`
      );

    if (!cell) {
      animatingEventId = null;

      return;
    }

    const cellRect =
      cell.getBoundingClientRect();
    const x2 =
      cellRect.left +
      cellRect.width / 2 -
      sectionRect.left;
    const y2 =
      cellRect.top +
      cellRect.height / 2 -
      sectionRect.top;
    const startWidth = sourceRect.width;
    const points = [
      { x: origin.x, y: origin.y }
    ];

    for (const key of buildDateKeys(
      event.startDate,
      targetKey
    )) {
      const dayCell =
        section.querySelector<HTMLElement>(
          `[data-date="${key}"]`
        );

      if (!dayCell) {
        continue;
      }

      const dayRect =
        dayCell.getBoundingClientRect();

      points.push({
        x:
          dayRect.left +
          dayRect.width / 2 -
          sectionRect.left,
        y:
          dayRect.top +
          dayRect.height / 2 -
          sectionRect.top
      });
    }

    if (points.length < 2) {
      points.push({ x: x2, y: y2 });
    }

    barState = {
      x: points[0].x,
      y: points[0].y - sourceRect.height / 2,
      width: startWidth,
      height: sourceRect.height,
      points,
      event
    };

    const totalLength = points.reduce(
      (total, point, index) => {
        if (index === 0) {
          return total;
        }

        const previous = points[index - 1];

        return (
          total +
          Math.hypot(
            point.x - previous.x,
            point.y - previous.y
          )
        );
      },
      0
    );

    await tick();

    const bar = barElement;

    if (!bar) {
      animatingEventId = null;

      return;
    }

    bar.setAttribute(
      'stroke-dasharray',
      `${totalLength}`
    );
    bar.setAttribute(
      'stroke-dashoffset',
      `${totalLength}`
    );
    gsap.set(bar, { opacity: 1 });

    barTween = gsap.to(bar, {
      attr: { 'stroke-dashoffset': 0 },
      duration: 0.8,
      ease: 'power2.inOut',
      onComplete: () => {
        flashDate = targetKey;
        targetEvent = event;

        flashTimer = setTimeout(() => {
          flashDate = null;
          targetEvent = null;

          barTween = gsap.to(bar, {
            attr: {
              'stroke-dashoffset': totalLength
            },
            duration: 0.4,
            ease: 'power2.in',
            onComplete: () => {
              barState = null;
              animatingEventId = null;
            }
          });
        }, 1400);
      }
    });
  }

  onDestroy(() => {
    barTween?.kill();
    clearTimeout(flashTimer);
  });

  onMount(() => {
    gsap.registerPlugin(
      ScrollTrigger
    );

    const section =
      sectionElement;

    const heading =
      headingElement;

    const navigation =
      navigationElement;

    const grid =
      gridElement;

    if (
      !section ||
      !heading ||
      !navigation ||
      !grid
    ) {
      return;
    }

    const reduceMotion =
      window.matchMedia(
        '(prefers-reduced-motion: reduce)'
      ).matches;

    if (reduceMotion) {
      return;
    }

    const context = gsap.context(
      () => {
        gsap.set(heading, {
          y: 55,
          opacity: 0
        });

        gsap.set(navigation, {
          y: 35,
          opacity: 0
        });

        gsap.set(grid, {
          y: 60,
          opacity: 0
        });

        const timeline =
          gsap.timeline({
            defaults: {
              ease: 'power3.out'
            },
            scrollTrigger: {
              trigger: section,
              start: 'top 75%',
              once: true
            }
          });

        timeline
          .to(
            heading,
            {
              y: 0,
              opacity: 1,
              duration: 0.8
            },
            0
          )
          .to(
            navigation,
            {
              y: 0,
              opacity: 1,
              duration: 0.65
            },
            0.2
          )
          .to(
            grid,
            {
              y: 0,
              opacity: 1,
              duration: 0.85
            },
            0.35
          )
          .set(
            [
              heading,
              navigation,
              grid
            ],
            {
              clearProps:
                'transform,opacity'
            }
          );
      },
      section
    );

    return () => {
      context.revert();

      gsap.killTweensOf(grid);
    };
  });
</script>

<section
  bind:this={sectionElement}
  id="calendar"
  aria-labelledby="calendar-heading"
  class="
    relative min-h-screen
    overflow-hidden
    bg-linear-to-b
    from-[#8fb2d8]
    via-[#dce8f5]
    to-[#8fb2d8]
    px-4 pb-24 pt-24
    sm:px-7 sm:pt-28
    lg:px-12 lg:pb-32
    lg:pt-32
  "
>
  <div
    aria-hidden="true"
    class="
      pointer-events-none
      absolute inset-x-0 top-0
      h-48
      bg-linear-to-b
      from-white/25
      to-transparent
    "
  ></div>

  <div
    class="
      relative z-10
      mx-auto w-full
      max-w-7xl
    "
  >
    <header class="text-center">
      <h2
        bind:this={headingElement}
        id="calendar-heading"
        class="
          text-4xl leading-[0.94]
          font-black
          tracking-[-0.065em]
          text-[#164f88]
          sm:text-6xl
          md:text-7xl
          lg:text-8xl
        "
      >
        Cerita Loka’s Calendar
      </h2>
    </header>

    <div
      bind:this={navigationElement}
      class="
        mx-auto w-full
        max-w-2xl
      "
    >
      <CalendarNavigation
        {currentMonth}
        onPrevious={showPreviousMonth}
        onNext={showNextMonth}
      />
    </div>

    <div
      bind:this={gridElement}
    >
      <CalendarGrid
        {currentMonth}
        events={resolvedEvents}
        {flashDate}
        {selectedDate}
        {targetEvent}
        onDateSelect={selectDate}
        onEventSelect={showEventOnCalendar}
      />

      {#if selectedDate}
        <div class="mx-auto mt-10 w-full max-w-4xl">
        <div class="flex flex-wrap items-baseline justify-between gap-2">
          <h3 class="text-xl font-black tracking-tight text-[#164f88] sm:text-2xl">
            Agenda bulan ini
          </h3>
          <span class="text-sm font-bold text-blue-900/60 capitalize">
            {monthLabel(currentMonth)} {currentMonth.getFullYear()}
          </span>
        </div>

        {#if monthAgenda.length}
          <ul class="mt-5 grid gap-3">
            {#each monthAgenda as event (event.id)}
              <li
                data-agenda-id={event.id}
                class="
                  flex flex-col gap-3 rounded-2xl
                  border border-blue-900/10
                  bg-white p-4 shadow-sm
                  sm:flex-row sm:items-center
                  sm:justify-between
                "
              >
                <div class="flex min-w-0 items-center gap-3">
                  <span class="flex size-9 shrink-0 items-center justify-center rounded-xl bg-orange-100 text-sm font-black text-orange-600">
                    {parseDate(eventTargetKey(event)).getDate()}
                  </span>

                  <div class="min-w-0">
                    <a
                      href={resolve(buildProgramWorkHref(event))}
                      class="block truncate font-extrabold text-blue-900 no-underline transition hover:text-orange-600"
                    >
                      {event.title}
                    </a>
                    <p class="mt-0.5 text-xs font-semibold text-blue-900/55">
                      {formatEventRange(event)} · {event.ministryName}
                    </p>
                  </div>
                </div>

                <button
                  type="button"
                  onclick={() => showEventOnCalendar(event)}
                  disabled={animatingEventId === event.id}
                  class="
                    shrink-0 rounded-full
                    bg-orange-500 px-4 py-2
                    text-xs font-black text-white
                    shadow-sm transition
                    hover:bg-orange-600
                    focus-visible:outline-2
                    focus-visible:outline-offset-2
                    focus-visible:outline-orange-500
                    disabled:cursor-wait disabled:opacity-70
                  "
                >
                  {animatingEventId === event.id ? 'Menunjukkan…' : 'Lihat di kalender'}
                </button>
              </li>
            {/each}
          </ul>
        {:else}
          <p class="mt-5 rounded-2xl border border-dashed border-blue-900/20 bg-white/60 p-5 text-center text-sm font-semibold text-blue-900/55">
            Belum ada agenda untuk bulan ini.
          </p>
        {/if}
        </div>
      {/if}
    </div>
  </div>

  {#if barState}
    <svg
      aria-hidden="true"
      class="pointer-events-none absolute inset-0 z-30 h-full w-full overflow-visible"
    >
      <polyline
        bind:this={barElement}
        points={barState.points
          .map((point) => `${point.x},${point.y}`)
          .join(' ')}
        fill="none"
        stroke="#f97316"
        stroke-width={Math.min(
          barState.height,
          14
        )}
        stroke-linecap="round"
        stroke-linejoin="round"
        opacity="0"
      />
    </svg>

    <div
      aria-hidden="true"
      class="
        pointer-events-none absolute z-40
        flex items-center gap-2 rounded-full
        bg-orange-500 px-3 py-1
        text-white shadow-lg
      "
      style:left={`${barState.x}px`}
      style:top={`${barState.y}px`}
      style:max-width={`${barState.width}px`}
    >
      <img
        src={barState.event.logo}
        alt=""
        class="size-4 shrink-0 rounded-full bg-white/90 p-0.5 object-contain"
      />
      <span class="truncate text-xs font-black">
        {barState.event.title}
      </span>
    </div>
  {/if}
</section>
