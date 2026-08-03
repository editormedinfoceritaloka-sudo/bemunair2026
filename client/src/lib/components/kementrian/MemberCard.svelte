<script lang="ts">
  let {
    role,
    title,
    image,
    featured = false
  }: {
    role: string;
    title: string;
    image: string;
    featured?: boolean;
  } = $props();
</script>

<article
  data-member-card
  class={`
    group mx-auto w-full
    rounded-xl
    [perspective:1200px]
    ${featured
      ? 'max-w-[310px] md:-translate-y-6'
      : 'max-w-[260px]'}
  `}
>
  <div
    data-member-card-inner
    class="
      relative w-full
      rounded-xl
      [aspect-ratio:0.88]
      [transform-style:preserve-3d]
      will-change-transform
    "
  >
    <div
      class="
        absolute inset-0
        overflow-hidden
        rounded-xl
        bg-white
        p-2.5
        shadow-[8px_10px_0_rgba(4,29,64,0.25)]
        [backface-visibility:hidden]
        transition-shadow duration-300
        group-hover:shadow-[12px_16px_0_rgba(4,29,64,0.22)]
        sm:p-3
      "
    >
      <div
        class="
          relative h-[78%] overflow-hidden
          rounded-xl
          bg-gradient-to-b
          from-blue-50
          via-blue-100
          to-blue-300
        "
      >
        <div
          class="
            pointer-events-none absolute inset-x-0 top-0
            h-1/2
            bg-gradient-to-b from-white/55 to-transparent
          "
        ></div>

        <img
          src={image}
          alt={`${role} ${title}`}
          class="
            absolute bottom-0 left-1/2
            h-[96%] w-auto max-w-none
            -translate-x-1/2
            object-contain object-bottom
            transition-transform duration-500
            group-hover:scale-[1.035]
          "
        />
      </div>

      <div
        class="
          flex h-[22%] flex-col
          items-center justify-center
          px-2 text-center
        "
      >
        <p
          class="
            text-xs leading-none font-black
            uppercase tracking-[0.08em]
            text-blue-800
            sm:text-sm
          "
        >
          {role}
        </p>

        <h3
          class={`
            mt-1 font-black
            leading-[1.02]
            tracking-[-0.035em]
            text-blue-900
            ${featured
              ? 'text-sm sm:text-base'
              : 'text-xs sm:text-sm'}
          `}
        >
          {title}
        </h3>
      </div>
    </div>

    <div
      class="
        absolute inset-0
        flex flex-col items-center justify-center
        overflow-hidden
        border-[10px] border-white
        bg-gradient-to-b
        from-blue-100
        via-blue-300
        to-blue-700
        px-6 text-center
        shadow-[8px_10px_0_rgba(4,29,64,0.25)]
        [backface-visibility:hidden]
        [transform:rotateY(180deg)]
      "
    >
      <div
        class="
          pointer-events-none absolute
          -left-16 -top-16
          size-40 rounded-full
          border border-white/30
        "
      ></div>

      <div
        class="
          pointer-events-none absolute
          -bottom-20 -right-16
          size-48 rounded-full
          border border-white/20
        "
      ></div>

      <div
        class="
          flex size-24 items-center justify-center
          rounded-[24px]
          border-4 border-white
          bg-white/90
          shadow-[0_16px_35px_rgba(4,29,64,0.28)]
          backdrop-blur-xl
          sm:size-28
        "
      >
        <img
          src="/logo/logo-kabinet.png"
          alt="Logo Kabinet Cerita Loka"
          class="h-[76%] w-[76%] object-contain"
        />
      </div>

      <p
        class="
          mt-5 text-xs font-black uppercase
          tracking-[0.18em]
          text-white
          drop-shadow-[0_2px_2px_rgba(0,0,0,0.35)]
        "
      >
        Kabinet Cerita Loka
      </p>

      <p
        class="
          mt-2 text-[10px] font-bold uppercase
          tracking-[0.12em]
          text-blue-50/85
        "
      >
        Universitas Airlangga
      </p>
    </div>
  </div>
</article>