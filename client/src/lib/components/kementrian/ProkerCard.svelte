<script lang="ts">
  import { resolve } from '$app/paths';
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
  href={resolve('/program-kerja/[slug]', { slug })}
  data-program-card
  class="
    group relative mx-auto
    w-full
    rounded-[22px]
    drop-shadow-[0_12px_18px_rgba(5,34,70,0.18)]
    transition-[filter,transform] duration-300
    hover:-translate-y-1
    hover:drop-shadow-[0_18px_26px_rgba(5,34,70,0.32)]
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
      [clip-path:polygon(0_0,calc(100%_-_2.75rem)_0,100%_2.75rem,100%_100%,0_100%)]
    "
  >
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

  <div
    aria-hidden="true"
    class="
      pointer-events-none absolute
      right-0 top-0 z-20
      size-11
      drop-shadow-[-2px_3px_3px_rgba(5,34,70,0.28)]
    "
  >
    <div
      class="
        size-full
        rounded-bl-lg
        bg-gradient-to-br
        from-white
        via-blue-100
        to-blue-300
        [clip-path:polygon(0_0,100%_100%,0_100%)]
      "
    ></div>
  </div>
</a>