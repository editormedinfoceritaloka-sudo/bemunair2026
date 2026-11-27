<script lang="ts">
  let {
    name = '',
    position = 'Menteri Koordinator',
    image = null
  }: {
    name?: string;
    position?: string;
    image?: string | null;
  } = $props();
</script>

<article
  class="
    group relative
    h-[290px] w-[310px]
    overflow-hidden rounded-[20px]
    border-[3px] border-white
    bg-white
    shadow-[0_14px_35px_rgba(15,50,92,0.25)]
    transition-transform duration-300
    hover:-translate-y-2
  "
>
  <div
    class="
      absolute inset-[8px]
      overflow-hidden rounded-[14px]
      bg-linear-to-b
      from-[#eaf4ff]
      via-[#c6ddf1]
      to-[#82b3d9]
    "
  >
    <div
      class="
        pointer-events-none absolute inset-x-0 top-0
        h-1/2
        bg-linear-to-b from-white/55 to-transparent
      "
    ></div>

    {#if image}
      <img
        src={image}
        alt={name}
        class="
          absolute bottom-0 left-1/2
          h-[94%] w-auto max-w-none
          -translate-x-1/2
          object-contain object-bottom
          transition-transform duration-500
          group-hover:scale-[1.035]
        "
      />
    {:else}
      <span aria-hidden="true" class="absolute inset-0 grid place-items-center text-6xl font-bold text-blue-700">{name.slice(0, 1)}</span>
    {/if}
  </div>

  <div
    class="
      absolute bottom-0 left-0 z-20
      min-w-[155px] max-w-[235px]
      rounded-tr-[20px]
      bg-white
      px-4 pb-3 pt-2.5
    "
  >
    <h3
      class="
        text-[21px] leading-[0.9]
        font-black
        tracking-[-0.045em]
        text-[#174b83]
      "
    >
      {name}
    </h3>

    <p
      class="
        mt-1.5
        text-[10px] leading-tight
        font-bold uppercase
        tracking-[0.12em]
        text-[#5480aa]
      "
    >
      {position}
    </p>
  </div>
</article>