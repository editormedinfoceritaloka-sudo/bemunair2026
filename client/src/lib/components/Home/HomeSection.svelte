<script lang="ts">
  import { onMount } from 'svelte';
  import { gsap } from 'gsap';

  let heroSection: HTMLElement;
  let leftSide: HTMLDivElement;
  let rightSide: HTMLDivElement;
  let bottomSide: HTMLDivElement;
  let centerText: HTMLDivElement;

  onMount(() => {
    const context = gsap.context(() => {
      gsap.set(leftSide, {
        xPercent: -110
      });

      gsap.set(rightSide, {
        xPercent: 110
      });

      gsap.set(bottomSide, {
        yPercent: 100,
        opacity: 0
      });

      gsap.set(centerText, {
        opacity: 0
      });

      const timeline = gsap.timeline({
        defaults: {
          ease: 'power3.out'
        }
      });

      timeline
        .to(
          leftSide,
          {
            xPercent: 0,
            duration: 1.2
          },
          0
        )
        .to(
          rightSide,
          {
            xPercent: 0,
            duration: 1.2
          },
          0
        )
        .call(
          () => {
            window.dispatchEvent(
              new CustomEvent('hero-bottom-enter')
            );
          },
          [],
          0.85
        )
        .to(
          bottomSide,
          {
            yPercent: 0,
            opacity: 1,
            duration: 0.9
          },
          0.85
        )
        .to(
          centerText,
          {
            opacity: 1,
            duration: 0.8,
            ease: 'power2.out'
          },
          1.55
        );
    }, heroSection);

    return () => {
      context.revert();
    };
  });
</script>

<section
  bind:this={heroSection}
  id="home"
  class="relative h-[80vh] overflow-hidden bg-linear-to-b from-blue-200 via-blue-50 to-white sm:h-[100vh] md:h-screen"
>
  <div class="absolute inset-0 flex flex-col">
    <div
      class="grid min-h-0 flex-1 grid-cols-[0.7fr_1.6fr_0.7fr] items-start md:grid-cols-3"
    >
      <div
        bind:this={leftSide}
        class="flex h-full min-w-0 items-start justify-center will-change-transform"
      >
        <img
          src="/landing/home/left.png"
          alt="Ilustrasi landing bagian kiri"
          class="h-full w-full -rotate-4 object-contain object-top md:rotate-0"
          draggable="false"
        />
      </div>

      <div
        bind:this={centerText}
        class="
          flex h-full min-w-0 flex-col
          items-center justify-center gap-2 px-1
          md:justify-start md:gap-4 md:px-0 md:py-24
        "
      >
        <h1
          class="
            whitespace-nowrap
            bg-linear-to-b
            from-blue-900
            via-blue-700
            to-blue-400
            bg-clip-text
            text-[clamp(2.25rem,10vw,4rem)]
            leading-none
            font-black
            tracking-[-0.06em]
            text-transparent
            drop-shadow-[0_8px_6px_rgba(15,23,42,0.28)]
            md:text-[clamp(4rem,13vw,12rem)]
          "
        >
          BEM UNAIR
        </h1>

        <p
          class="
            whitespace-nowrap
            text-center
            text-[clamp(1rem,4vw,1.5rem)]
            leading-none
            font-black
            tracking-[-0.04em]
            text-orange-600
            drop-shadow-[0_4px_3px_rgba(120,53,15,0.35)]
            md:text-[clamp(1.5rem,3.2vw,3.5rem)]
          "
        >
          Kabinet Cerita Loka
        </p>

        <p
          class="
            max-w-md text-center
            text-[10px] leading-snug
            md:max-w-none md:text-lg md:leading-normal
          "
        >
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed pulvinar
          facilisis ornare. Maecenas ut dignissim orci. Mauris ac accumsan
          tellus.
        </p>
      </div>

      <div
        bind:this={rightSide}
        class="flex h-full min-w-0 items-start justify-center will-change-transform"
      >
        <img
          src="/landing/home/right.png"
          alt="Ilustrasi landing bagian kanan"
          class="h-full w-full object-contain object-top"
          draggable="false"
        />
      </div>
    </div>

    <div
      bind:this={bottomSide}
      class="relative z-10 w-full shrink-0 will-change-transform"
    >
      <img
        src="/landing/home/bottom.png"
        alt="Ilustrasi landing bagian bawah"
        class="block h-auto w-full object-cover object-bottom"
        draggable="false"
      />
    </div>
  </div>
</section>