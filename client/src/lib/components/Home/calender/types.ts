export type CalendarEvent = {
  id: string;
  title: string;
  startDate: string;
  endDate?: string;
  ministrySlug: string;
  programSlug: string;
};

export type EventSegment = {
  event: CalendarEvent;
  startColumn: number;
  span: number;
  lane: number;
  continuesBefore: boolean;
  continuesAfter: boolean;
};

export type WeekView = {
  days: Date[];
  segments: EventSegment[];
  height: number;
  rangeTop: number;
};
