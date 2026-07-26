<script lang="ts">
  import StatusBadge from '$lib/admin/StatusBadge.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
  import * as Dialog from '$lib/components/ui/dialog';
  import type { Article, ContentSubmission, LetterSubmission } from '$lib/types';
  import {
    ArrowRight,
    CalendarDays,
    ChevronLeft,
    ChevronRight,
    Clock3,
    FileImage,
    Mail,
    Newspaper
  } from '@lucide/svelte';

  type EventKind = 'CONTENT' | 'LETTER' | 'ARTICLE';

  interface CalendarEvent {
    id: string;
    kind: EventKind;
    title: string;
    date: string;
    href: string;
    status: string;
    ministry?: string;
    description: string;
  }

  interface CalendarDay {
    key: string;
    day: number;
    inMonth: boolean;
    isToday: boolean;
  }

  let {
    content = [],
    letters = [],
    articles = []
  }: {
    content?: ContentSubmission[];
    letters?: LetterSubmission[];
    articles?: Article[];
  } = $props();

  const now = new Date();
  let viewYear = $state(now.getFullYear());
  let viewMonth = $state(now.getMonth());
  let detailOpen = $state(false);
  let selectedEvents = $state<CalendarEvent[]>([]);
  let selectedDate = $state('');

  const weekdays = ['Sen', 'Sel', 'Rab', 'Kam', 'Jum', 'Sab', 'Min'];

  function pad(value: number) {
    return String(value).padStart(2, '0');
  }

  function localKey(year: number, month: number, day: number) {
    return `${year}-${pad(month + 1)}-${pad(day)}`;
  }

  function jakartaDateKey(value: string) {
    const date = new Date(value);
    if (Number.isNaN(date.getTime())) return '';
    const parts = new Intl.DateTimeFormat('en-US', {
      timeZone: 'Asia/Jakarta',
      year: 'numeric',
      month: '2-digit',
      day: '2-digit'
    }).formatToParts(date);
    const part = (type: Intl.DateTimeFormatPartTypes) =>
      parts.find((item) => item.type === type)?.value || '';
    return `${part('year')}-${part('month')}-${part('day')}`;
  }

  function buildEvents() {
    const contentEvents: CalendarEvent[] = content
      .filter((item) => item.deadline || item.publish_date)
      .map((item) => ({
        id: `content-${item.id}`,
        kind: 'CONTENT',
        title: item.title,
        date: item.deadline || item.publish_date || '',
        href: `/admin/content-submissions/${item.id}`,
        status: item.status,
        ministry: item.ministry,
        description: `${item.submission_type.replaceAll('_', ' ')} · ${item.assigned_pj?.name || 'PJ belum ditetapkan'}`
      }));

    const letterEvents: CalendarEvent[] = letters
      .filter((item) => item.deadline)
      .map((item) => ({
        id: `letter-${item.id}`,
        kind: 'LETTER',
        title: item.subject,
        date: item.deadline,
        href: `/admin/letter-submissions/${item.id}`,
        status: item.status,
        ministry: item.ministry,
        description: `${item.letter_type} · ${item.assigned_pj?.name || 'PJ belum ditetapkan'}`
      }));

    const articleEvents: CalendarEvent[] = articles
      .filter((item) => item.published_at)
      .map((item) => ({
        id: `article-${item.id}`,
        kind: 'ARTICLE',
        title: item.title,
        date: item.published_at || '',
        href: `/admin/articles/${item.id}/preview`,
        status: item.status,
        description: `Dipublikasikan oleh ${item.author?.name || 'Admin BEM UNAIR'}`
      }));

    return [...contentEvents, ...letterEvents, ...articleEvents].sort(
      (left, right) => new Date(left.date).getTime() - new Date(right.date).getTime()
    );
  }

  function groupEvents(rows: CalendarEvent[]) {
    const grouped: Record<string, CalendarEvent[]> = {};
    for (const event of rows) {
      const key = jakartaDateKey(event.date);
      if (!key) continue;
      (grouped[key] ||= []).push(event);
    }
    return grouped;
  }

  function buildDays(year: number, month: number) {
    const first = new Date(year, month, 1);
    const mondayOffset = (first.getDay() + 6) % 7;
    const gridStart = new Date(year, month, 1 - mondayOffset);
    const todayKey = localKey(now.getFullYear(), now.getMonth(), now.getDate());

    return Array.from({ length: 42 }, (_, index): CalendarDay => {
      const date = new Date(gridStart.getFullYear(), gridStart.getMonth(), gridStart.getDate() + index);
      const key = localKey(date.getFullYear(), date.getMonth(), date.getDate());
      return {
        key,
        day: date.getDate(),
        inMonth: date.getMonth() === month,
        isToday: key === todayKey
      };
    });
  }

  const events = $derived(buildEvents());
  const eventsByDate = $derived(groupEvents(events));
  const calendarDays = $derived(buildDays(viewYear, viewMonth));
  const monthTitle = $derived(
    new Intl.DateTimeFormat('id-ID', { month: 'long', year: 'numeric' }).format(
      new Date(viewYear, viewMonth, 1)
    )
  );
  const monthEventCount = $derived(
    events.filter((event) => jakartaDateKey(event.date).startsWith(`${viewYear}-${pad(viewMonth + 1)}`)).length
  );

  function changeMonth(offset: number) {
    const next = new Date(viewYear, viewMonth + offset, 1);
    viewYear = next.getFullYear();
    viewMonth = next.getMonth();
  }

  function goToToday() {
    viewYear = now.getFullYear();
    viewMonth = now.getMonth();
  }

  function openDetails(rows: CalendarEvent[], key: string) {
    if (!rows.length) return;
    selectedEvents = rows;
    selectedDate = key;
    detailOpen = true;
  }

  function formattedSelectedDate() {
    if (!selectedDate) return '';
    const [year, month, day] = selectedDate.split('-').map(Number);
    return new Intl.DateTimeFormat('id-ID', { dateStyle: 'full' }).format(
      new Date(year, month - 1, day)
    );
  }

  function formattedEventTime(value: string) {
    return new Intl.DateTimeFormat('id-ID', {
      timeZone: 'Asia/Jakarta',
      hour: '2-digit',
      minute: '2-digit',
      hourCycle: 'h23'
    }).format(new Date(value));
  }

  function kindLabel(kind: EventKind) {
    if (kind === 'CONTENT') return 'Konten';
    if (kind === 'LETTER') return 'Surat';
    return 'Artikel';
  }

  function kindClasses(kind: EventKind) {
    if (kind === 'CONTENT') return 'border-blue-100 bg-blue-50 text-blue-700 hover:bg-blue-100';
    if (kind === 'LETTER') return 'border-orange-100 bg-orange-50 text-orange-700 hover:bg-orange-100';
    return 'border-green-100 bg-green-50 text-green-800 hover:bg-green-100';
  }
</script>

<Card class="mt-6 overflow-hidden border-border/80 shadow-sm">
  <CardHeader class="gap-4 border-b border-border/80 pb-4 lg:flex-row lg:items-center lg:justify-between">
    <div>
      <CardTitle class="flex items-center gap-2 text-base">
        <span class="grid size-8 place-items-center rounded-lg bg-blue-50 text-blue-700">
          <CalendarDays class="size-4" />
        </span>
        Kalender operasional
      </CardTitle>
      <p class="mt-1 text-sm text-muted-foreground">
        {monthEventCount} agenda pada {monthTitle}
      </p>
    </div>
    <div class="flex flex-wrap items-center gap-2">
      <Button type="button" variant="outline" size="sm" onclick={goToToday}>Hari ini</Button>
      <div class="flex items-center rounded-lg border border-border bg-card">
        <Button
          type="button"
          variant="ghost"
          size="icon-sm"
          onclick={() => changeMonth(-1)}
          aria-label="Bulan sebelumnya"
        >
          <ChevronLeft />
        </Button>
        <p class="min-w-36 px-2 text-center text-sm font-semibold capitalize">{monthTitle}</p>
        <Button
          type="button"
          variant="ghost"
          size="icon-sm"
          onclick={() => changeMonth(1)}
          aria-label="Bulan berikutnya"
        >
          <ChevronRight />
        </Button>
      </div>
    </div>
  </CardHeader>

  <CardContent class="p-0">
    <div class="grid grid-cols-7 border-b border-border bg-white-500/70">
      {#each weekdays as weekday}
        <div class="px-1 py-2 text-center text-[10px] font-bold uppercase tracking-wider text-black-300 sm:text-xs">
          {weekday}
        </div>
      {/each}
    </div>
    <div class="grid grid-cols-7 bg-border/80">
      {#each calendarDays as day}
        {@const dayEvents = eventsByDate[day.key] || []}
        <div
          class:opacity-45={!day.inMonth}
          class="min-h-20 min-w-0 bg-card p-1.5 sm:min-h-28 sm:p-2"
        >
          <button
            type="button"
            class:cursor-default={!dayEvents.length}
            class="mb-1 grid size-7 place-items-center rounded-full text-xs font-semibold transition-colors hover:bg-muted sm:size-8 sm:text-sm"
            class:bg-blue-600={day.isToday}
            class:text-white-50={day.isToday}
            aria-label={dayEvents.length ? `Lihat ${dayEvents.length} agenda tanggal ${day.day}` : `Tanggal ${day.day}`}
            onclick={() => openDetails(dayEvents, day.key)}
          >
            {day.day}
          </button>

          <div class="space-y-1">
            {#each dayEvents.slice(0, 2) as event}
              <button
                type="button"
                class={`flex h-5 w-full min-w-0 items-center justify-center rounded-md border px-1 text-left text-[10px] font-medium transition-colors sm:h-auto sm:justify-start sm:py-1 ${kindClasses(event.kind)}`}
                aria-label={`${kindLabel(event.kind)}: ${event.title}`}
                title={event.title}
                onclick={() => openDetails([event], day.key)}
              >
                <span class="size-1.5 shrink-0 rounded-full bg-current sm:mr-1.5"></span>
                <span class="hidden truncate sm:block">{event.title}</span>
              </button>
            {/each}
            {#if dayEvents.length > 2}
              <button
                type="button"
                class="block w-full truncate px-1 text-center text-[10px] font-semibold text-blue-700 hover:underline sm:text-left"
                onclick={() => openDetails(dayEvents, day.key)}
              >
                +{dayEvents.length - 2} agenda
              </button>
            {/if}
          </div>
        </div>
      {/each}
    </div>
    <div class="flex flex-wrap gap-x-5 gap-y-2 border-t border-border px-4 py-3 text-xs text-muted-foreground">
      <span class="flex items-center gap-1.5"><span class="size-2 rounded-full bg-blue-500"></span>Deadline konten</span>
      <span class="flex items-center gap-1.5"><span class="size-2 rounded-full bg-orange-500"></span>Deadline surat</span>
      <span class="flex items-center gap-1.5"><span class="size-2 rounded-full bg-green-600"></span>Artikel terbit</span>
    </div>
  </CardContent>
</Card>

<Dialog.Root bind:open={detailOpen}>
  <Dialog.Content class="max-h-[85vh] overflow-y-auto sm:max-w-xl">
    <Dialog.Header>
      <Dialog.Title>
        {selectedEvents.length === 1 ? 'Detail agenda' : `${selectedEvents.length} agenda`}
      </Dialog.Title>
      <Dialog.Description class="capitalize">{formattedSelectedDate()}</Dialog.Description>
    </Dialog.Header>

    <div class="space-y-3">
      {#each selectedEvents as event}
        {@const EventIcon = event.kind === 'CONTENT' ? FileImage : event.kind === 'LETTER' ? Mail : Newspaper}
        <article class="rounded-xl border border-border bg-card p-4">
          <div class="flex items-start gap-3">
            <span class={`grid size-10 shrink-0 place-items-center rounded-xl border ${kindClasses(event.kind)}`}>
              <EventIcon class="size-5" />
            </span>
            <div class="min-w-0 flex-1">
              <div class="flex flex-wrap items-center gap-2">
                <p class="text-xs font-bold uppercase tracking-wider text-muted-foreground">{kindLabel(event.kind)}</p>
                <StatusBadge status={event.status} />
              </div>
              <h3 class="mt-2 font-serif text-lg font-semibold leading-snug">{event.title}</h3>
              <p class="mt-1 text-sm text-muted-foreground">{event.description}</p>
              {#if event.ministry}<p class="mt-1 text-xs font-medium text-black-300">{event.ministry}</p>{/if}
              <p class="mt-3 flex items-center gap-1.5 text-xs text-black-300">
                <Clock3 class="size-3.5" /> {formattedEventTime(event.date)} WIB
              </p>
            </div>
          </div>
          <Button href={event.href} variant="outline" size="sm" class="mt-4 w-full">
            Buka detail <ArrowRight />
          </Button>
        </article>
      {/each}
    </div>
  </Dialog.Content>
</Dialog.Root>
