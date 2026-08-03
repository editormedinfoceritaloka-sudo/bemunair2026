<script lang="ts">
  import CalendarGrid from './calender/CalendarGrid.svelte';
  import CalendarNavigation from './calender/CalendarNavigation.svelte';

  import {
    addMonths,
    parseDate,
    startOfMonth
  } from './calender/calendar-utils';

  import type { CalendarEvent } from './calender/types';

  const logo = '/logo/logo-kabinet.png';

  const events: CalendarEvent[] = [
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

  let currentMonth = $state(
    startOfMonth(parseDate('2026-06-01'))
  );

  function showPreviousMonth(): void {
    currentMonth = addMonths(currentMonth, -1);
  }

  function showNextMonth(): void {
    currentMonth = addMonths(currentMonth, 1);
  }
</script>

<section
  id="calendar"
  aria-labelledby="calendar-heading"
  class="
    relative py-24 overflow-hidden
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
    class="
      pointer-events-none
      absolute left-1/2 top-16
      h-[480px] w-[820px]
      -translate-x-1/2
      rounded-full
      bg-white/30
      blur-[120px]
    "
  ></div>

  <div class="relative mx-auto w-full max-w-7xl">
    <header class="text-center">
      <h2
        id="calendar-heading"
        class="
          text-5xl leading-[0.94]
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

    <CalendarNavigation
      {currentMonth}
      onPrevious={showPreviousMonth}
      onNext={showNextMonth}
    />

    <CalendarGrid
      {currentMonth}
      {events}
      {logo}
    />
  </div>
</section>