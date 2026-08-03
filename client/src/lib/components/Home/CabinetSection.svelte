<script lang="ts">
  import LogoCarrousel from './cabinet/LogoCarrousel.svelte';
  import MenkoSlider from './cabinet/MenkoSlider.svelte';
</script>

<section
  id="cabinet"
  class="
    min-h-screen w-full overflow-hidden
    bg-linear-to-b from-blue-100 via-blue-50 to-white
    px-4 py-20
    sm:px-6 sm:py-24
    md:py-28
  "
>
  <div class="mx-auto flex w-full max-w-7xl flex-col items-center">
    <div
      class="
        flex flex-col items-center
        text-center font-black uppercase
        leading-none tracking-[-0.05em]
      "
    >
      <h2
        class="
          whitespace-nowrap
          text-[clamp(2.3rem,8vw,7rem)]
          text-orange-600
          [-webkit-text-stroke:1.5px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_5px_3px_rgba(15,23,42,0.3)]
          sm:[-webkit-text-stroke:2px_white]
          md:[-webkit-text-stroke:3px_white]
        "
      >
        Kementerian
      </h2>

      <p
        class="
          mt-3 whitespace-nowrap
          text-[clamp(1.1rem,3vw,2.5rem)]
          text-blue-700
          [-webkit-text-stroke:1px_white]
          [paint-order:stroke_fill]
          drop-shadow-[0_3px_2px_rgba(15,23,42,0.3)]
          sm:mt-4
          sm:[-webkit-text-stroke:1.5px_white]
          md:[-webkit-text-stroke:2px_white]
        "
      >
        Cerita Loka
      </p>
    </div>

    <div class="mt-6 w-full sm:mt-8 md:mt-10">
      <LogoCarrousel />
    </div>

    <div class="mt-10 w-full sm:mt-14 md:mt-16">
      <MenkoSlider autoplayDelay={3000} />
    </div>
  </div>
</section>