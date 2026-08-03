import type {
  CalendarEvent,
  EventSegment
} from './types';

export function parseDate(value: string): Date {
  const [year, month, day] = value.split('-').map(Number);

  return new Date(
    year,
    month - 1,
    day,
    12,
    0,
    0,
    0
  );
}

export function startOfMonth(date: Date): Date {
  return new Date(
    date.getFullYear(),
    date.getMonth(),
    1,
    12,
    0,
    0,
    0
  );
}

export function endOfMonth(date: Date): Date {
  return new Date(
    date.getFullYear(),
    date.getMonth() + 1,
    0,
    12,
    0,
    0,
    0
  );
}

export function addMonths(
  date: Date,
  amount: number
): Date {
  return new Date(
    date.getFullYear(),
    date.getMonth() + amount,
    1,
    12,
    0,
    0,
    0
  );
}

export function addDays(
  date: Date,
  amount: number
): Date {
  const result = new Date(date);

  result.setDate(result.getDate() + amount);

  return result;
}

export function startOfWeek(date: Date): Date {
  const currentDay = date.getDay();
  const offset = currentDay === 0 ? -6 : 1 - currentDay;

  return addDays(date, offset);
}

export function endOfWeek(date: Date): Date {
  return addDays(startOfWeek(date), 6);
}

export function dateKey(date: Date): string {
  const year = date.getFullYear();
  const month = String(
    date.getMonth() + 1
  ).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');

  return `${year}-${month}-${day}`;
}

export function differenceInDays(
  start: Date,
  end: Date
): number {
  const millisecondsPerDay = 1000 * 60 * 60 * 24;

  return Math.round(
    (end.getTime() - start.getTime()) /
      millisecondsPerDay
  );
}

export function buildCalendarWeeks(
  month: Date
): Date[][] {
  const firstVisibleDate = startOfWeek(
    startOfMonth(month)
  );

  const lastVisibleDate = endOfWeek(
    endOfMonth(month)
  );

  const days: Date[] = [];
  const weeks: Date[][] = [];

  let cursor = firstVisibleDate;

  while (cursor <= lastVisibleDate) {
    days.push(cursor);
    cursor = addDays(cursor, 1);
  }

  for (
    let index = 0;
    index < days.length;
    index += 7
  ) {
    weeks.push(days.slice(index, index + 7));
  }

  return weeks;
}

export function isRangeEvent(
  event: CalendarEvent
): boolean {
  return Boolean(
    event.endDate &&
      event.endDate !== event.startDate
  );
}

export function getSingleDayEvents(
  day: Date,
  events: CalendarEvent[]
): CalendarEvent[] {
  const selectedDate = dateKey(day);

  return events.filter(
    (event) =>
      !isRangeEvent(event) &&
      event.startDate === selectedDate
  );
}

export function buildWeekSegments(
  week: Date[],
  events: CalendarEvent[]
): EventSegment[] {
  const weekStart = week[0];
  const weekEnd = week[6];

  const rawSegments: Array<{
    event: CalendarEvent;
    startColumn: number;
    endColumn: number;
    span: number;
    continuesBefore: boolean;
    continuesAfter: boolean;
  }> = [];

  for (const event of events) {
    if (!isRangeEvent(event)) {
      continue;
    }

    const eventStart = parseDate(event.startDate);
    const eventEnd = parseDate(
      event.endDate ?? event.startDate
    );

    if (
      eventStart > weekEnd ||
      eventEnd < weekStart
    ) {
      continue;
    }

    const clippedStart =
      eventStart < weekStart
        ? weekStart
        : eventStart;

    const clippedEnd =
      eventEnd > weekEnd
        ? weekEnd
        : eventEnd;

    const startColumn = differenceInDays(
      weekStart,
      clippedStart
    );

    const endColumn = differenceInDays(
      weekStart,
      clippedEnd
    );

    rawSegments.push({
      event,
      startColumn,
      endColumn,
      span: endColumn - startColumn + 1,
      continuesBefore: eventStart < weekStart,
      continuesAfter: eventEnd > weekEnd
    });
  }

  rawSegments.sort((first, second) => {
    if (
      first.startColumn !== second.startColumn
    ) {
      return first.startColumn - second.startColumn;
    }

    return second.span - first.span;
  });

  const lanes: Array<
    Array<{
      start: number;
      end: number;
    }>
  > = [];

  return rawSegments.map((segment) => {
    let lane = 0;

    while (
      lanes[lane]?.some(
        (occupied) =>
          segment.startColumn <= occupied.end &&
          segment.endColumn >= occupied.start
      )
    ) {
      lane += 1;
    }

    if (!lanes[lane]) {
      lanes[lane] = [];
    }

    lanes[lane].push({
      start: segment.startColumn,
      end: segment.endColumn
    });

    return {
      event: segment.event,
      startColumn: segment.startColumn,
      span: segment.span,
      lane,
      continuesBefore:
        segment.continuesBefore,
      continuesAfter:
        segment.continuesAfter
    };
  });
}

export function isCurrentMonth(
  date: Date,
  currentMonth: Date
): boolean {
  return (
    date.getMonth() ===
      currentMonth.getMonth() &&
    date.getFullYear() ===
      currentMonth.getFullYear()
  );
}

export function isToday(date: Date): boolean {
  return dateKey(date) === dateKey(new Date());
}

export function monthLabel(date: Date): string {
  return new Intl.DateTimeFormat('en-US', {
    month: 'long'
  }).format(date);
}

export function buildProgramWorkHref(
  event: CalendarEvent
): string {
  const ministrySlug = encodeURIComponent(
    event.ministrySlug
  );

  const programSlug = encodeURIComponent(
    event.programSlug
  );

  return `/kementrian/${ministrySlug}/program-kerja/${programSlug}`;
}
