<script lang="ts">
  import { resolve } from '$app/paths';
  import { untrack } from 'svelte';
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import CalendarGrid from './calender/CalendarGrid.svelte';
  import CalendarNavigation from './calender/CalendarNavigation.svelte';

  import {
    addMonths,
    buildProgramWorkHref,
    dateKey,
    formatEventRange,
    getEventsForDay,
    getEventsForMonth,
    parseDate,
    startOfMonth
  } from './calender/calendar-utils';

  import type { CalendarEvent } from './calender/types';

  import type {
    Cabinet,
    OrganizationUnit,
    WorkProgram
  } from '$lib/types';

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

  let selectedDate = $state<Date | null>(null);
  const dateFormatter = new Intl.DateTimeFormat('id-ID', { dateStyle: 'full' });
  const agendaEvents = $derived(
    selectedDate
      ? getEventsForDay(selectedDate, resolvedEvents)
      : getEventsForMonth(currentMonth, resolvedEvents)
  );

  function selectDate(day: Date): void {
    selectedDate = selectedDate && dateKey(selectedDate) === dateKey(day) ? null : day;
  }

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
    selectedDate = null;
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
        selectedDate={selectedDate ? dateKey(selectedDate) : null}
        onSelectDate={selectDate}
      />
    </div>

    <div class="mx-auto mt-6 w-full max-w-4xl">
      <div class="flex flex-wrap items-center justify-between gap-2">
        <h3 class="text-lg font-black text-blue-900" aria-live="polite">
          {selectedDate ? `Program mulai ${dateFormatter.format(selectedDate)}` : 'Agenda bulan ini'}
        </h3>
        {#if selectedDate}
          <button
            type="button"
            class="rounded-md px-2 py-1 text-sm font-semibold text-blue-800 underline hover:bg-blue-50"
            onclick={() => selectedDate = null}
          >
            Semua agenda bulan ini
          </button>
        {/if}
      </div>

      {#if agendaEvents.length}
        <ul class="mt-4 grid gap-3">
          {#each agendaEvents as event (event.id)}
            <li>
              <a
                href={resolve(buildProgramWorkHref(event))}
                class="
                  flex flex-col gap-1
                  rounded-xl border
                  border-blue-900/10
                  bg-white p-4
                  no-underline shadow-sm
                  transition
                  hover:border-orange-400
                  hover:shadow-md
                  sm:flex-row
                  sm:items-center
                  sm:justify-between
                "
              >
                <span class="font-semibold text-blue-900">
                  {event.title}
                </span>

                <span class="text-sm text-muted-foreground">
                  {formatEventRange(event)} · {event.ministryName}
                </span>
              </a>
            </li>
          {/each}
        </ul>
      {:else}
        <p class="mt-3 text-sm text-muted-foreground">Belum ada agenda untuk bulan ini.</p>
      {/if}
    </div>
  </div>
</section>
