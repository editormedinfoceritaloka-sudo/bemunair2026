<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';

  const cabinetLogos = [
    { name: 'ADKESMA', src: '/landing/cabinet/LOGO ADKESMA RAW.png' },
    { name: 'HUBLU', src: '/landing/cabinet/logo hublu.png' },
    { name: 'LH', src: '/landing/cabinet/LOGO LH.png' },
    { name: 'MEDINFO', src: '/landing/cabinet/LOGO MEDINFO.png' },
    { name: 'MENKES', src: '/landing/cabinet/LOGO MENKES.png' },
    { name: 'PENGMAS', src: '/landing/cabinet/LOGO PENGMAS.png' },
    { name: 'PENGPROF', src: '/landing/cabinet/LOGO PENGPROF.png' },
    { name: 'PSDM', src: '/landing/cabinet/LOGO PSDM.png' },
    { name: 'RISKEL', src: '/landing/cabinet/LOGO RISKEL.png' },
    { name: 'SENIORA', src: '/landing/cabinet/LOGO SENIORA.png' },
    { name: 'SINEMA', src: '/landing/cabinet/LOGO SINEMA.png' },
    { name: 'SOSPOL', src: '/landing/cabinet/LOGO SOSPOL.png' }
  ] as const;

  let carousel: HTMLDivElement;
  let track: HTMLDivElement;
  let animation: gsap.core.Tween | undefined;

  function pauseCarousel(): void {
    animation?.pause();
  }

  function resumeCarousel(): void {
    animation?.play();
  }

  onMount(() => {
    const reduceMotion = window.matchMedia(
      '(prefers-reduced-motion: reduce)'
    ).matches;

    if (reduceMotion) return;

    const context = gsap.context(() => {
      animation = gsap.to(track, {
        xPercent: -50,
        duration: 28,
        ease: 'none',
        repeat: -1
      });
    }, carousel);

    return () => {
      animation?.kill();
      context.revert();
    };
  });
</script>

<div class="-rotate-2.5 overflow-hidden py-8">
  <div
    bind:this={carousel}
    class="w-full overflow-hidden"
    role="region"
    aria-label="Logo kementerian BEM UNAIR"
    onmouseenter={pauseCarousel}
    onmouseleave={resumeCarousel}
  >
    <div bind:this={track} class="flex w-max will-change-transform">
      {#each [0, 1] as copy}
        <div
          aria-hidden={copy === 1}
          class="flex shrink-0 items-center gap-5 pr-5 sm:gap-8 sm:pr-8 md:gap-10 md:pr-10"
        >
          {#each cabinetLogos as logo (`${copy}-${logo.name}`)}
            <div
              class="
                flex size-24 shrink-0 items-center justify-center
                transition duration-300
                hover:-translate-y-2 hover:scale-105
                sm:size-32 sm:p-4
                md:size-40 md:p-5
              "
            >
              <img
                src={logo.src}
                alt={copy === 0 ? `Logo ${logo.name}` : ''}
                draggable="false"
                class="h-full w-full object-contain"
              />
            </div>
          {/each}
        </div>
      {/each}
    </div>
  </div>
</div>