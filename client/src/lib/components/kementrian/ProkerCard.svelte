<script lang="ts">
  let {
    title,
    description,
    slug
  }: {
    title: string;
    description: string;
    slug: string;
  } = $props();
</script>

<a
  href={`/program-kerja/${slug}`}
  data-program-card
  class="
    group relative mx-auto
    w-full
    overflow-hidden rounded-[22px]
    bg-black
    shadow-[0_12px_26px_rgba(5,34,70,0.18)]
    transition-shadow duration-300
    hover:shadow-[0_20px_38px_rgba(5,34,70,0.38)]
  "
>
  <div
    class="
      relative z-10
      flex h-full w-full flex-col
      justify-center items-center
      overflow-hidden rounded-[22px]
      bg-gradient-to-br
      from-blue-100
      via-blue-300
      to-blue-500
      px-5 py-6
      text-left
      [clip-path:polygon(0_0,86%_0,100%_14%,100%_100%,0_100%)]
    "
  >
    <div
      aria-hidden="true"
      class="
        pointer-events-none absolute
        right-0 top-0
        size-11
        bg-black
        [clip-path:polygon(0_0,100%_0,100%_100%)]
      "
    ></div>

    <h3
      class="
        shrink-0
        text-md leading-none font-black
        tracking-[-0.055em]
        text-blue-700
        [-webkit-text-stroke:1.5px_white]
        [paint-order:stroke_fill]
        drop-shadow-[0_2px_0_rgba(5,34,70,0.55)]
        sm:text-lg
        lg:text-xl
      "
    >
      {title}
    </h3>

    <p
      class="
        mt-3 overflow-hidden
        text-[11px] leading-[1.5]
        font-medium
        text-slate-950
        sm:text-xs
      "
    >
      {description}
    </p>
  </div>
</a>