<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';
  import { ScrollTrigger } from 'gsap/ScrollTrigger';

  import CalendarGrid from './calender/CalendarGrid.svelte';
  import CalendarNavigation from './calender/CalendarNavigation.svelte';

  import {
    addMonths,
    parseDate,
    startOfMonth
  } from './calender/calendar-utils';

  import type { CalendarEvent } from './calender/types';

  const exampleEvents: CalendarEvent[] = [
    {
      id: 'company-profile',
      title: 'Company Profile',
      startDate: '2026-06-06',
      endDate: '2026-06-20',
      ministrySlug: 'komunikasi-dan-informasi',
      programSlug: 'company-profile'
    },
    {
      id: 'festival-cerita-loka',
      title: 'Festival Cerita Loka',
      startDate: '2026-06-13',
      endDate: '2026-06-27',
      ministrySlug: 'seni-dan-budaya',
      programSlug: 'festival-cerita-loka'
    },
    {
      id: 'cerita-loka-mengabdi',
      title: 'Cerita Loka Mengabdi',
      startDate: '2026-06-08',
      endDate: '2026-06-18',
      ministrySlug: 'pengabdian-masyarakat',
      programSlug: 'cerita-loka-mengabdi'
    },
    {
      id: 'launching-news',
      title: 'Launching News',
      startDate: '2026-06-07',
      ministrySlug: 'komunikasi-dan-informasi',
      programSlug: 'launching-news'
    },
    {
      id: 'forum-aspirasi',
      title: 'Forum Aspirasi',
      startDate: '2026-06-15',
      ministrySlug: 'sosial-dan-politik',
      programSlug: 'forum-aspirasi'
    },
    {
      id: 'anniversary',
      title: 'Cerita Loka Anniversary',
      startDate: '2026-06-24',
      ministrySlug: 'sekretariat-kabinet',
      programSlug: 'cerita-loka-anniversary'
    }
  ];

  let {
    events = exampleEvents,
    logo = '/logo/logo-kabinet.png',
    initialMonth = '2026-06-01'
  }: {
    events?: CalendarEvent[];
    logo?: string;
    initialMonth?: string;
  } = $props();

  let currentMonth = $state(
    startOfMonth(parseDate(initialMonth))
  );

  let sectionElement: HTMLElement;
  let glowElement: HTMLDivElement;
  let headingElement: HTMLHeadingElement;
  let navigationElement: HTMLDivElement;
  let gridElement: HTMLDivElement;

  function showPreviousMonth(): void {
    currentMonth = addMonths(currentMonth, -1);
  }

  function showNextMonth(): void {
    currentMonth = addMonths(currentMonth, 1);
  }

  onMount(() => {
    gsap.registerPlugin(ScrollTrigger);

    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (reduceMotion) {
      gsap.set(
        [
          glowElement,
          headingElement,
          navigationElement,
          gridElement
        ],
        {
          clearProps: 'all'
        }
      );

      return;
    }

    const context = gsap.context(() => {
      gsap.set(glowElement, {
        opacity: 0,
        scale: 0.65
      });

      gsap.set(headingElement, {
        y: 70,
        opacity: 0,
        scale: 0.94,
        filter: 'blur(10px)'
      });

      gsap.set(navigationElement, {
        y: 40,
        opacity: 0
      });

      gsap.set(gridElement, {
        y: 80,
        opacity: 0,
        scale: 0.97,
        transformOrigin: 'center top'
      });

      const timeline = gsap.timeline({
        defaults: {
          ease: 'power3.out'
        },
        scrollTrigger: {
          trigger: sectionElement,
          start: 'top 72%',
          once: true
        }
      });

      timeline
        .to(
          glowElement,
          {
            opacity: 1,
            scale: 1,
            duration: 1.4,
            ease: 'power2.out'
          },
          0
        )
        .to(
          headingElement,
          {
            y: 0,
            opacity: 1,
            scale: 1,
            filter: 'blur(0px)',
            duration: 1
          },
          0.1
        )
        .to(
          navigationElement,
          {
            y: 0,
            opacity: 1,
            duration: 0.75
          },
          0.45
        )
        .to(
          gridElement,
          {
            y: 0,
            opacity: 1,
            scale: 1,
            duration: 1.05
          },
          0.65
        )
        .set(
          [
            glowElement,
            headingElement,
            navigationElement,
            gridElement
          ],
          {
            clearProps: 'willChange,transform,filter,opacity'
          }
        );
    }, sectionElement);

    return () => {
      context.revert();
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
    bg-gradient-to-b
    from-[#8fb2d8]
    via-[#dce8f5]
    to-[#8fb2d8]
    px-4 pb-24 pt-28
    sm:px-7 sm:pt-32
    lg:px-12 lg:pb-32 lg:pt-36
  "
>
  <div
    bind:this={glowElement}
    class="
      pointer-events-none
      absolute left-1/2 top-16
      h-[480px] w-[820px]
      -translate-x-1/2
      rounded-full
      bg-white/30
      blur-[120px]
      will-change-transform
    "
  ></div>

  <div class="relative mx-auto w-full max-w-7xl">
    <header class="text-center">
      <h2
        bind:this={headingElement}
        id="calendar-heading"
        class="
          text-5xl leading-[0.94]
          font-black
          tracking-[-0.065em]
          text-[#164f88]
          will-change-transform
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
      class="will-change-transform"
    >
      <CalendarNavigation
        {currentMonth}
        onPrevious={showPreviousMonth}
        onNext={showNextMonth}
      />
    </div>

    <div
      bind:this={gridElement}
      class="will-change-transform"
    >
      <CalendarGrid
        {currentMonth}
        {events}
        {logo}
      />
    </div>
  </div>
</section>