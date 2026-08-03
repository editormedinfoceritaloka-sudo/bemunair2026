<script lang="ts">
  let {
    startDate,
    endDate
  }: {
    startDate: string;
    endDate?: string | null;
  } = $props();

  const monthNames = [
    'JANUARI',
    'FEBRUARI',
    'MARET',
    'APRIL',
    'MEI',
    'JUNI',
    'JULI',
    'AGUSTUS',
    'SEPTEMBER',
    'OKTOBER',
    'NOVEMBER',
    'DESEMBER'
  ];

  function parseDate(value: string): Date | null {
    const [year, month, day] = value
      .split('-')
      .map(Number);

    if (
      !Number.isInteger(year) ||
      !Number.isInteger(month) ||
      !Number.isInteger(day)
    ) {
      return null;
    }

    const date = new Date(year, month - 1, day);

    if (
      date.getFullYear() !== year ||
      date.getMonth() !== month - 1 ||
      date.getDate() !== day
    ) {
      return null;
    }

    return date;
  }

  function isSameDate(
    firstDate: Date,
    secondDate: Date
  ): boolean {
    return (
      firstDate.getFullYear() ===
        secondDate.getFullYear() &&
      firstDate.getMonth() ===
        secondDate.getMonth() &&
      firstDate.getDate() ===
        secondDate.getDate()
    );
  }

  const dateLabel = $derived.by(() => {
    const start = parseDate(startDate);

    const end = endDate
      ? parseDate(endDate)
      : null;

    if (!start) {
      return startDate;
    }

    if (!end || isSameDate(start, end)) {
      return `${start.getDate()} ${
        monthNames[start.getMonth()]
      } ${start.getFullYear()}`;
    }

    const sameMonth =
      start.getMonth() === end.getMonth();

    const sameYear =
      start.getFullYear() === end.getFullYear();

    if (sameMonth && sameYear) {
      return `${start.getDate()}–${end.getDate()} ${
        monthNames[start.getMonth()]
      } ${start.getFullYear()}`;
    }

    if (sameYear) {
      return `${start.getDate()} ${
        monthNames[start.getMonth()]
      }–${end.getDate()} ${
        monthNames[end.getMonth()]
      } ${start.getFullYear()}`;
    }

    return `${start.getDate()} ${
      monthNames[start.getMonth()]
    } ${start.getFullYear()}–${end.getDate()} ${
      monthNames[end.getMonth()]
    } ${end.getFullYear()}`;
  });
</script>

<div
  class="
    relative mx-auto
    aspect-[1054/544] w-full
    max-w-[900px]
    overflow-hidden
  "
>
  <div
    aria-hidden="true"
    class="
      absolute left-[10%] top-[27%]
      h-[55%] w-[80%]
      translate-x-[1.2%]
      translate-y-[1.5%]
      rounded-[18px]
      bg-blue-950/30
    "
  ></div>

  <div
    class="
      timeline-grid
      absolute left-[10%] top-[26%]
      h-[55%] w-[80%]
      overflow-hidden rounded-[18px]
      border-[6px] border-blue-950
      bg-blue-950
      shadow-[0_10px_0_rgba(30,64,175,0.65)]
    "
  >
    <div
      class="
        absolute inset-[14px]
        border-2 border-blue-200/70
      "
    ></div>
  </div>

  <div
    class="
      absolute left-[15.5%] top-[37%]
      z-20 flex h-[34%] w-[68%]
      items-center justify-center
      rounded-[14px]
      bg-white px-5
      shadow-[0_12px_14px_rgba(15,23,42,0.28)]
    "
  >
    <div
      class="
        flex w-full flex-col
        items-center justify-center
      "
    >
      <h2
        class="
          timeline-date
          max-w-full text-center
          text-[clamp(1.25rem,5vw,4.4rem)]
          leading-none font-black
          tracking-[-0.06em]
          text-blue-900
        "
      >
        {dateLabel}
      </h2>

      <div
        class="
          mt-[4%] h-[5px]
          w-[62%] rounded-full
          bg-orange-400
          shadow-[0_2px_4px_rgba(249,115,22,0.35)]
        "
      ></div>
    </div>
  </div>

  <div
    class="
      absolute left-[3.8%] top-[4.5%]
      z-40 w-[31%]
      -rotate-[11deg]
    "
  >
    <div
      class="
        relative flex min-h-[70px]
        items-center
        rounded-[12px]
        border-[5px] border-orange-700
        bg-gradient-to-b
        from-orange-300
        via-orange-400
        to-orange-300
        px-[7%] py-[4%]
        shadow-[0_8px_0_#c2410c,0_14px_18px_rgba(15,23,42,0.28)]
        [clip-path:polygon(0_10%,82%_0,100%_30%,84%_100%,4%_100%)]
      "
    >
      <span
        class="
          text-[clamp(0.9rem,3.7vw,3rem)]
          leading-none font-black
          tracking-[-0.055em]
          text-white
          [-webkit-text-stroke:2px_#2563eb]
          [paint-order:stroke_fill]
          drop-shadow-[0_3px_0_rgba(255,255,255,0.7)]
        "
      >
        TIMELINE
      </span>
    </div>
  </div>

  <div
    class="
      absolute left-[3%] top-[20%]
      z-30 w-[47%]
      -rotate-[11deg]
    "
  >
    <div
      class="
        flex min-h-[90px]
        items-center
        rounded-[10px]
        border border-slate-300
        bg-gradient-to-b
        from-slate-100
        via-white
        to-slate-300
        px-[5%] pb-[4%] pt-[6%]
        shadow-[0_8px_0_#1e4f91,0_14px_18px_rgba(15,23,42,0.28)]
        [clip-path:polygon(0_18%,94%_0,100%_18%,97%_75%,7%_100%,0_82%)]
      "
    >
      <span
        class="
          text-[clamp(0.9rem,3.8vw,3.2rem)]
          leading-[0.9] font-black
          tracking-[-0.055em]
          text-white
          [-webkit-text-stroke:2px_#ea580c]
          [paint-order:stroke_fill]
          drop-shadow-[0_3px_0_rgba(255,255,255,0.7)]
        "
      >
        PELAKSANAAN
      </span>
    </div>
  </div>
</div>

<style>
  .timeline-grid {
    background-color: #102f5a;
    background-image:
      linear-gradient(
        rgba(191, 219, 254, 0.58) 1px,
        transparent 1px
      ),
      linear-gradient(
        90deg,
        rgba(191, 219, 254, 0.58) 1px,
        transparent 1px
      );
    background-size:
      clamp(16px, 2.5vw, 26px)
      clamp(16px, 2.5vw, 26px);
  }

  .timeline-date {
    -webkit-text-stroke: 2px white;
    paint-order: stroke fill;
    text-shadow:
      0 5px 0 #dbeafe,
      0 9px 4px rgba(15, 23, 42, 0.38);
  }

  @media (max-width: 640px) {
    .timeline-date {
      -webkit-text-stroke-width: 1px;
      text-shadow:
        0 3px 0 #dbeafe,
        0 5px 3px rgba(15, 23, 42, 0.32);
    }
  }
</style>