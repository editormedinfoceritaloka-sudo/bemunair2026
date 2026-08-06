import type { CalendarEvent } from './types';

export function parseDate(value: string): Date {
  const normalizedValue = value.slice(0, 10);

  const [year, month, day] =
    normalizedValue.split('-').map(Number);

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

export function startOfMonth(
  date: Date
): Date {
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

export function endOfMonth(
  date: Date
): Date {
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

  result.setDate(
    result.getDate() + amount
  );

  return result;
}

export function startOfWeek(
  date: Date
): Date {
  const currentDay = date.getDay();

  const offset =
    currentDay === 0
      ? -6
      : 1 - currentDay;

  return addDays(date, offset);
}

export function endOfWeek(
  date: Date
): Date {
  return addDays(
    startOfWeek(date),
    6
  );
}

export function dateKey(
  date: Date
): string {
  const year = date.getFullYear();

  const month = String(
    date.getMonth() + 1
  ).padStart(2, '0');

  const day = String(
    date.getDate()
  ).padStart(2, '0');

  return `${year}-${month}-${day}`;
}

export function buildCalendarWeeks(
  month: Date
): Date[][] {
  const firstVisibleDate =
    startOfWeek(
      startOfMonth(month)
    );

  const lastVisibleDate =
    endOfWeek(
      endOfMonth(month)
    );

  const days: Date[] = [];
  const weeks: Date[][] = [];

  let cursor = firstVisibleDate;

  while (cursor <= lastVisibleDate) {
    days.push(cursor);

    cursor = addDays(
      cursor,
      1
    );
  }

  for (
    let index = 0;
    index < days.length;
    index += 7
  ) {
    weeks.push(
      days.slice(
        index,
        index + 7
      )
    );
  }

  return weeks;
}

export function getEventsForDay(
  day: Date,
  events: CalendarEvent[]
): CalendarEvent[] {
  const selectedDate = dateKey(day);

  return events
    .filter(
      (event) =>
        event.startDate ===
        selectedDate
    )
    .sort((first, second) =>
      first.title.localeCompare(
        second.title,
        'id-ID'
      )
    );
}

export function getEventsForMonth(
  month: Date,
  events: CalendarEvent[]
): CalendarEvent[] {
  return events.filter((event) => {
    const eventDate =
      parseDate(event.startDate);

    return (
      eventDate.getMonth() ===
        month.getMonth() &&
      eventDate.getFullYear() ===
        month.getFullYear()
    );
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

export function isToday(
  date: Date
): boolean {
  return (
    dateKey(date) ===
    dateKey(new Date())
  );
}

export function monthLabel(
  date: Date
): string {
  return new Intl.DateTimeFormat(
    'id-ID',
    {
      month: 'long'
    }
  ).format(date);
}

export function buildProgramWorkHref(
  event: CalendarEvent
): string {
  const programSlug =
    encodeURIComponent(
      event.programSlug
    );

  return `/program-kerja/${programSlug}`;
}