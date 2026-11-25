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

  function pointAtProgress(
    points: { x: number; y: number }[],
    totalLength: number,
    progress: number
  ): { x: number; y: number; angle: number } {
    let remaining =
      totalLength *
      Math.min(Math.max(progress, 0), 1);

    for (
      let index = 1;
      index < points.length;
      index += 1
    ) {
      const from = points[index - 1];
      const to = points[index];
      const segmentLength = Math.hypot(
        to.x - from.x,
        to.y - from.y
      );

      if (
        remaining <= segmentLength ||
        index === points.length - 1
      ) {
        const ratio =
          segmentLength === 0
            ? 0
            : Math.min(
                Math.max(
                  remaining / segmentLength,
                  0
                ),
                1
              );

        return {
          x:
            from.x +
            (to.x - from.x) * ratio,
          y:
            from.y +
            (to.y - from.y) * ratio,
          angle:
            (Math.atan2(
              to.y - from.y,
              to.x - from.x
            ) *
              180) /
            Math.PI
        };
      }

      remaining -= segmentLength;
    }

    const last =
      points[points.length - 1];

    return {
      x: last.x,
      y: last.y,
      angle: 0
    };
  }

  function buildRoundedPath(
    points: { x: number; y: number }[]
  ): string {
    if (points.length === 0) {
      return '';
    }

    if (points.length === 1) {
      return `M ${points[0].x} ${points[0].y}`;
    }

    const radius = 22;
    let path = `M ${points[0].x} ${points[0].y}`;

    for (
      let index = 1;
      index < points.length - 1;
      index += 1
    ) {
      const previous = points[index - 1];
      const current = points[index];
      const next = points[index + 1];

      const inLength = Math.hypot(
        current.x - previous.x,
        current.y - previous.y
      );
      const outLength = Math.hypot(
        next.x - current.x,
        next.y - current.y
      );

      if (inLength === 0 || outLength === 0) {
        continue;
      }

      const corner = Math.min(
        radius,
        inLength / 2,
        outLength / 2
      );

      const startX =
        current.x +
        ((previous.x - current.x) * corner) / inLength;
      const startY =
        current.y +
        ((previous.y - current.y) * corner) / inLength;
      const endX =
        current.x +
        ((next.x - current.x) * corner) / outLength;
      const endY =
        current.y +
        ((next.y - current.y) * corner) / outLength;

      path += ` L ${startX} ${startY} Q ${current.x} ${current.y} ${endX} ${endY}`;
    }

    const last = points[points.length - 1];

    path += ` L ${last.x} ${last.y}`;

    return path;
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
    endX: number;
    endY: number;
    width: number;
    height: number;
    points: { x: number; y: number }[];
    path: string;
    event: CalendarEvent;
  } | null>(null);
  let barElement = $state<
    SVGPathElement | undefined
  >(undefined);
  let sparkleElement = $state<
    SVGPathElement | undefined
  >(undefined);
  let rocketElement = $state<
    SVGSVGElement | undefined
  >(undefined);
  let smokeElement = $state<
    HTMLDivElement | undefined
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

    if (barElement) {
      gsap.killTweensOf(barElement);
    }

    if (rocketElement) {
      gsap.killTweensOf(rocketElement);
    }

    if (smokeElement) {
      gsap.killTweensOf(smokeElement);
    }

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

    const badgeSource =
      section.querySelector<HTMLElement>(
        `[data-event-start="${event.id}"]`
      );

    const source =
      (badgeSource &&
      badgeSource.getBoundingClientRect().width > 0
        ? badgeSource
        : null) ??
      section.querySelector<HTMLElement>(
        `[data-date="${event.startDate}"] [data-day-list]`
      ) ??
      section.querySelector<HTMLElement>(
        `[data-date="${event.startDate}"]`
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

      const nextPoint = {
        x:
          dayRect.left +
          dayRect.width / 2 -
          sectionRect.left,
        y:
          dayRect.top +
          dayRect.height / 2 -
          sectionRect.top
      };

      const previousPoint =
        points[points.length - 1];

      if (
        points.length > 1 &&
        nextPoint.y > previousPoint.y + 5 &&
        nextPoint.x < previousPoint.x
      ) {
        const middleY =
          (previousPoint.y + nextPoint.y) / 2;

        points.push(
          { x: previousPoint.x, y: middleY },
          { x: nextPoint.x, y: middleY }
        );
      }

      points.push(nextPoint);
    }

    if (points.length < 2) {
      points.push({ x: x2, y: y2 });
    }

    const endPoint =
      points[points.length - 1];

    barState = {
      x: points[0].x,
      y: points[0].y - sourceRect.height / 2,
      endX: endPoint.x,
      endY: endPoint.y,
      width: startWidth,
      height: sourceRect.height,
      points,
      path: buildRoundedPath(points),
      event
    };

    const pointsLength = points.reduce(
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

    const pathLength = bar.getTotalLength();
    const totalLength =
      pathLength > 0 ? pathLength : pointsLength;
    const sparkle = sparkleElement;

    if (sparkle) {
      gsap.killTweensOf(sparkle);
      gsap.set(sparkle, { opacity: 0 });
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

    const rocket = rocketElement;
    const smoke = smokeElement;

    const placeRocket = (
      progress: number
    ) => {
      if (!rocket) {
        return;
      }

      if (pathLength <= 0) {
        const point = pointAtProgress(
          points,
          totalLength,
          progress
        );

        gsap.set(rocket, {
          left: point.x,
          top: point.y,
          rotation: point.angle
        });

        return;
      }

      const distance =
        totalLength *
        Math.min(Math.max(progress, 0), 1);
      const point = bar.getPointAtLength(distance);
      const behind = bar.getPointAtLength(
        Math.max(distance - 2, 0)
      );
      const ahead = bar.getPointAtLength(
        Math.min(distance + 2, totalLength)
      );

      gsap.set(rocket, {
        left: point.x,
        top: point.y,
        rotation:
          (Math.atan2(
            ahead.y - behind.y,
            ahead.x - behind.x
          ) *
            180) /
          Math.PI
      });
    };

    const playSmoke = () => {
      if (!smoke) {
        return;
      }

      const puffs = Array.from(
        smoke.querySelectorAll<HTMLElement>(
          '[data-smoke-puff]'
        )
      );

      gsap.set(smoke, { autoAlpha: 1 });
      gsap.fromTo(
        puffs,
        {
          x: 0,
          y: 0,
          scale: 0.35,
          opacity: 0.9
        },
        {
          x: (index: number) =>
            [-18, 12, 26, -8, 4][
              index % 5
            ],
          y: (index: number) =>
            [-6, -18, 8, 14, -12][
              index % 5
            ],
          scale: 2.4,
          opacity: 0,
          duration: 0.75,
          stagger: 0.05,
          ease: 'power2.out'
        }
      );
    };

    if (rocket) {
      gsap.set(rocket, {
        autoAlpha: 1,
        scale: 1
      });
      placeRocket(0);
    }

    if (smoke) {
      gsap.set(smoke, { autoAlpha: 0 });
    }

    const flight = { progress: 0 };

    barTween = gsap
      .timeline({
        onComplete: () => {
          flashDate = targetKey;
          targetEvent = event;

          if (sparkle) {
            gsap.to(sparkle, {
              opacity: 0.9,
              duration: 0.3,
              ease: 'power1.out'
            });
          }

          flashTimer = setTimeout(() => {
            flashDate = null;
            targetEvent = null;

            if (sparkle) {
              gsap.to(sparkle, {
                opacity: 0,
                duration: 0.2
              });
            }

            barTween = gsap
              .timeline({
                onComplete: () => {
                  if (smoke) {
                    gsap.set(smoke, {
                      autoAlpha: 0
                    });
                  }

                  barState = null;
                  animatingEventId = null;
                }
              })
              .to(bar, {
                attr: {
                  'stroke-dashoffset': totalLength
                },
                duration: 0.4,
                ease: 'power2.in'
              })
              .call(() => {
                playSmoke();
              })
              .to(rocket ?? {}, {
                autoAlpha: 0,
                scale: 0.6,
                duration: 0.25,
                ease: 'power2.in'
              })
              .to({}, {
                duration: 0.75
              });
          }, 1400);
        }
      })
      .to(
        bar,
        {
          attr: {
            'stroke-dashoffset': 0
          },
          duration: 1.1,
          ease: 'power2.inOut'
        },
        0
      )
      .to(
        flight,
        {
          progress: 1,
          duration: 1.1,
          ease: 'power2.inOut',
          onUpdate: () =>
            placeRocket(flight.progress)
        },
        0
      );
  }

  async function showEventGlowOnly(event: CalendarEvent): Promise<void> {
    const targetKey = eventTargetKey(event);
    const targetDate = parseDate(targetKey);

    barTween?.kill();

    if (barElement) {
      gsap.killTweensOf(barElement);
    }

    if (rocketElement) {
      gsap.killTweensOf(rocketElement);
      gsap.set(rocketElement, { autoAlpha: 0 });
    }

    if (smokeElement) {
      gsap.killTweensOf(smokeElement);
      gsap.set(smokeElement, { autoAlpha: 0 });
    }

    clearTimeout(flashTimer);
    flashDate = null;
    targetEvent = null;
    barState = null;
    animatingEventId = event.id;

    if (
      dateKey(currentMonth).slice(0, 7) !==
      targetKey.slice(0, 7)
    ) {
      currentMonth = startOfMonth(targetDate);
    }

    await tick();

    const cell =
      sectionElement?.querySelector<HTMLElement>(
        `[data-date="${targetKey}"]`
      );

    cell?.scrollIntoView({
      behavior: 'smooth',
      block: 'center'
    });

    flashTimer = setTimeout(() => {
      flashDate = targetKey;
      flashTimer = setTimeout(() => {
        flashDate = null;
        animatingEventId = null;
      }, 1400);
    }, cell ? 650 : 0);
  }

  onDestroy(() => {
    barTween?.kill();

    if (barElement) {
      gsap.killTweensOf(barElement);
    }

    if (rocketElement) {
      gsap.killTweensOf(rocketElement);
    }

    if (smokeElement) {
      gsap.killTweensOf(smokeElement);
    }

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
                  onclick={() => showEventGlowOnly(event)}
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
      <defs>
        <linearGradient
          id="calendar-trail-gradient"
          gradientUnits="userSpaceOnUse"
          x1={barState.points[0].x}
          y1={barState.points[0].y}
          x2={barState.endX}
          y2={barState.endY}
        >
          <stop
            offset="0"
            stop-color="#fed7aa"
          />
          <stop
            offset="0.45"
            stop-color="#fb923c"
          />
          <stop
            offset="1"
            stop-color="#ea580c"
          />
        </linearGradient>
      </defs>

      <path
        bind:this={barElement}
        d={barState.path}
        fill="none"
        stroke="url(#calendar-trail-gradient)"
        class="
          [stroke-width:4px]
          drop-shadow-[0_0_6px_rgba(249,115,22,0.55)]
          sm:[stroke-width:8px]
        "
        stroke-linecap="round"
        stroke-linejoin="round"
        style:opacity="0"
      />

      <path
        bind:this={sparkleElement}
        d={barState.path}
        fill="none"
        stroke="#fff7ed"
        class="
          calendar-trail-sparkle
          [stroke-width:1.5px]
          sm:[stroke-width:2.5px]
        "
        stroke-linecap="round"
        stroke-dasharray="1 12"
        style:opacity="0"
      />
    </svg>

    <svg
      bind:this={rocketElement}
      aria-hidden="true"
      viewBox="0 0 72 72"
      style:opacity="0"
      class="
        pointer-events-none absolute z-50
        size-10
        drop-shadow-[0_6px_14px_rgba(249,115,22,0.45)]
        sm:size-14
      "
      style:translate="-50% -50%"
    >
      <defs>
        <linearGradient
          id="rocket-body-gradient"
          x1="0"
          y1="0"
          x2="0"
          y2="1"
        >
          <stop
            offset="0"
            stop-color="#ffffff"
          />
          <stop
            offset="1"
            stop-color="#e0ecfb"
          />
        </linearGradient>
        <linearGradient
          id="rocket-nose-gradient"
          x1="0"
          y1="0"
          x2="1"
          y2="1"
        >
          <stop
            offset="0"
            stop-color="#fdba74"
          />
          <stop
            offset="1"
            stop-color="#ea580c"
          />
        </linearGradient>
        <linearGradient
          id="rocket-window-gradient"
          x1="0"
          y1="0"
          x2="1"
          y2="1"
        >
          <stop
            offset="0"
            stop-color="#7dd3fc"
          />
          <stop
            offset="1"
            stop-color="#164f88"
          />
        </linearGradient>
        <linearGradient
          id="rocket-flame-gradient"
          x1="1"
          y1="0"
          x2="0"
          y2="0"
        >
          <stop
            offset="0"
            stop-color="#fef3c7"
          />
          <stop
            offset="0.45"
            stop-color="#fb923c"
          />
          <stop
            offset="1"
            stop-color="#f97316"
            stop-opacity="0"
          />
        </linearGradient>
      </defs>

      <g class="rocket-flame">
        <path
          d="M20 28 C11 29 5 33 1 36 C5 39 11 43 20 44 Z"
          fill="url(#rocket-flame-gradient)"
        />
        <path
          d="M20 32 C15 33 11 35 8 36 C11 37 15 39 20 40 Z"
          fill="#fffbeb"
        />
      </g>

      <path
        d="M27 26 L17 15 C24 14 31 17 36 23 Z"
        fill="#c2410c"
      />
      <path
        d="M27 46 L17 57 C24 58 31 55 36 49 Z"
        fill="#c2410c"
      />
      <path
        d="M19 30 C19 27 21 25 24 25 L50 25 C59 25 66 30 69 36 C66 42 59 47 50 47 L24 47 C21 47 19 45 19 42 Z"
        fill="url(#rocket-body-gradient)"
        stroke="#164f88"
        stroke-width="2.5"
        stroke-linejoin="round"
      />
      <path
        d="M54 25.6 C61 27 66 31 69 36 C66 41 61 45 54 46.4 C57 43 58 39.5 58 36 C58 32.5 57 29 54 25.6 Z"
        fill="url(#rocket-nose-gradient)"
      />
      <path
        d="M28 26.5 L28 45.5"
        stroke="#fb923c"
        stroke-width="2.5"
        stroke-linecap="round"
      />
      <path
        d="M33 29.5 L48 29.5"
        stroke="#ffffff"
        stroke-width="2"
        stroke-linecap="round"
        opacity="0.9"
      />
      <circle
        cx="43"
        cy="36"
        r="6.5"
        fill="url(#rocket-window-gradient)"
        stroke="#164f88"
        stroke-width="2.5"
      />
      <circle
        cx="41"
        cy="34"
        r="1.8"
        fill="#ffffff"
        opacity="0.95"
      />
    </svg>

    <div
      bind:this={smokeElement}
      aria-hidden="true"
      style:opacity="0"
      class="pointer-events-none absolute z-40 size-20"
      style:left={`${barState.endX}px`}
      style:top={`${barState.endY}px`}
      style:translate="-50% -50%"
    >
      <span
        data-smoke-puff
        class="absolute left-1/2 top-1/2 size-3 rounded-full bg-white/85 blur-[2px]"
        style:translate="-50% -50%"
      ></span>
      <span
        data-smoke-puff
        class="absolute left-1/2 top-1/2 size-4 rounded-full bg-sky-100/75 blur-[3px]"
        style:translate="-50% -50%"
      ></span>
      <span
        data-smoke-puff
        class="absolute left-1/2 top-1/2 size-3 rounded-full bg-white/75 blur-[2px]"
        style:translate="-50% -50%"
      ></span>
      <span
        data-smoke-puff
        class="absolute left-1/2 top-1/2 size-5 rounded-full bg-blue-50/70 blur-[4px]"
        style:translate="-50% -50%"
      ></span>
      <span
        data-smoke-puff
        class="absolute left-1/2 top-1/2 size-3 rounded-full bg-white/80 blur-[2px]"
        style:translate="-50% -50%"
      ></span>
    </div>
  {/if}
</section>

<style>
  .rocket-flame {
    transform-origin: right center;
    animation: rocket-flame 0.16s ease-in-out
      infinite alternate;
  }

  @keyframes rocket-flame {
    from {
      opacity: 0.75;
      transform: scaleX(0.72) scaleY(0.85);
    }

    to {
      opacity: 1;
      transform: scaleX(1) scaleY(1.15);
    }
  }

  .calendar-trail-sparkle {
    animation: calendar-trail-sparkle 0.9s linear
      infinite;
  }

  @keyframes calendar-trail-sparkle {
    to {
      stroke-dashoffset: -26;
    }
  }

  @media (prefers-reduced-motion: reduce) {
    .rocket-flame,
    .calendar-trail-sparkle {
      animation: none;
    }
  }
</style>