<script lang="ts">
  let {
    title,
    slug,
    logo
  }: {
    title: string;
    slug: string;
    logo: string;
  } = $props();
</script>

<a
  href={`/kementrian/${slug}`}
  aria-label={`Lihat halaman ${title}`}
  class="
    group relative block
    aspect-[3/4] w-full
    overflow-hidden rounded-[18px]
    border-[3px] border-white/90
    bg-gradient-to-b
    from-[#eef7ff]
    via-[#d7e9f8]
    to-[#9cc4e4]
    p-2
    shadow-[0_12px_30px_rgba(2,30,70,0.35)]
    transition-all duration-300
    hover:-translate-y-2
    hover:shadow-[0_20px_40px_rgba(2,30,70,0.5)]
    focus-visible:outline-4
    focus-visible:outline-offset-4
    focus-visible:outline-white
  "
>
  <div
    class="
      pointer-events-none absolute inset-[7px]
      rounded-[12px]
      border border-blue-700/30
    "
  ></div>

  <div
    class="
      pointer-events-none absolute
      left-4 right-4 top-3
      h-px bg-blue-700/35
    "
  ></div>

  <div
    class="
      pointer-events-none absolute
      bottom-3 left-4 right-4
      h-px bg-blue-700/35
    "
  ></div>

  <span
    class="
      pointer-events-none absolute left-3 top-3
      h-4 w-4
      border-l-2 border-t-2 border-blue-700/45
    "
  ></span>

  <span
    class="
      pointer-events-none absolute right-3 top-3
      h-4 w-4
      border-r-2 border-t-2 border-blue-700/45
    "
  ></span>

  <span
    class="
      pointer-events-none absolute bottom-3 left-3
      h-4 w-4
      border-b-2 border-l-2 border-blue-700/45
    "
  ></span>

  <span
    class="
      pointer-events-none absolute bottom-3 right-3
      h-4 w-4
      border-b-2 border-r-2 border-blue-700/45
    "
  ></span>

  <div
    class="
      relative z-10 flex h-full
      flex-col items-center justify-center
      px-3 py-6 text-center
    "
  >
    <img
      src={logo}
      alt={`Logo ${title}`}
      class="
        h-20 w-20 object-contain
        drop-shadow-[0_6px_10px_rgba(15,55,100,0.2)]
        transition-transform duration-300
        group-hover:scale-110
        sm:h-24 sm:w-24
        lg:h-28 lg:w-28
      "
    />

    <h3
      class="
        mt-4 max-w-full
        text-xs leading-tight font-black
        uppercase tracking-[-0.025em]
        text-[#174f88]
        transition-colors duration-300
        group-hover:text-[#0b315b]
        sm:text-sm
      "
    >
      {title}
    </h3>

    <span
      class="
        mt-2 text-[9px] font-extrabold
        uppercase tracking-[0.16em]
        text-blue-700/70
        opacity-0 transition-opacity duration-300
        group-hover:opacity-100
      "
    >
      Lihat Kementerian
    </span>
  </div>
</a>