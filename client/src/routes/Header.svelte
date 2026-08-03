<script lang="ts">
  import { onMount } from 'svelte';
  import { resolve } from '$app/paths';
  import { page } from '$app/state';
  import { gsap } from 'gsap';
  import { publicNavigation } from '$lib/constants/public-navigation';

  let headerElement: HTMLElement;

  function isActive(hash: string): boolean {
    if (page.url.pathname !== '/') {
      return false;
    }

    if (!page.url.hash) {
      return hash === '#home';
    }

    return page.url.hash === hash;
  }

  onMount(() => {
    gsap.set(headerElement, {
      yPercent: -130,
      opacity: 0
    });

    const showHeader = () => {
      gsap.to(headerElement, {
        yPercent: 0,
        opacity: 1,
        duration: 0.9,
        ease: 'power3.out',
        clearProps: 'transform,opacity'
      });
    };

    window.addEventListener('hero-bottom-enter', showHeader);

    if (page.url.pathname !== '/') {
      showHeader();
    }

    return () => {
      window.removeEventListener('hero-bottom-enter', showHeader);
      gsap.killTweensOf(headerElement);
    };
  });
</script>

<header
  bind:this={headerElement}
  class="fixed inset-x-0 top-0 z-50 px-1 py-1.5 sm:px-4 sm:py-2 md:px-6 md:py-3"
>
  <div
    class="mx-auto flex w-full max-w-7xl items-center justify-between gap-0.5 sm:gap-2 lg:gap-7"
  >
    <a
      href={resolve('/')}
      aria-label="Kembali ke halaman utama"
      class="flex shrink-0 items-center gap-0.5 rounded-md border border-white/70 bg-white/25 px-1 py-1 shadow-sm backdrop-blur-xl transition duration-200 hover:bg-white/40 sm:gap-1.5 sm:rounded-lg sm:px-2 sm:py-1.5 md:rounded-xl md:px-3 md:py-2"
    >
      <div class="flex shrink-0 items-center gap-px sm:gap-0.5 md:gap-1">
        <img
          class="size-4 object-contain sm:size-6 md:size-8"
          src="/logo/logo-bem.png"
          alt="Logo BEM UNAIR"
        />

        <img
          class="size-4 object-contain sm:size-6 md:size-8"
          src="/logo/logo-kabinet.png"
          alt="Logo Kabinet Cerita Loka"
        />
      </div>

      <div class="min-w-0 leading-none">
        <p
          class="whitespace-nowrap text-[5px] font-extrabold text-black-700 sm:text-[7px] md:text-[10px]"
        >
          Kabinet
        </p>

        <p
          class="mt-0.5 whitespace-nowrap text-[5px] font-medium text-black-600 sm:text-[7px] md:mt-1 md:text-[10px]"
        >
          Cerita Loka
        </p>
      </div>
    </a>

    <nav
      aria-label="Navigasi publik"
      class="flex w-fit min-w-0 flex-none items-center justify-center rounded-md border border-white/70 bg-white/25 p-px shadow-sm backdrop-blur-xl sm:rounded-lg sm:p-0.5 md:rounded-xl md:px-4 md:py-1"
    >
      {#each publicNavigation as item (item.hash)}
        <a
          href={item.href}
          aria-current={isActive(item.hash) ? 'location' : undefined}
          class={`
            whitespace-nowrap rounded-sm px-0.5 py-1
            text-[7px] font-semibold transition duration-200
            sm:rounded-md sm:px-1.5 sm:text-[9px]
            md:rounded-lg md:px-3 md:py-2 md:text-xs
            lg:px-5 lg:text-sm
            ${
              isActive(item.hash)
                ? 'bg-white/45 text-blue-900 shadow-sm'
                : 'text-blue-600 hover:bg-white/35 hover:text-blue-800'
            }
          `}
        >
          {item.label}
        </a>
      {/each}
    </nav>

    <div
      class="flex shrink-0 items-center gap-0.5 rounded-md border border-white/70 bg-white/25 px-1 py-1 shadow-sm backdrop-blur-xl transition duration-200 hover:bg-white/40 sm:gap-1 sm:rounded-lg sm:px-2 sm:py-1.5 md:gap-2 md:rounded-xl md:px-4 md:py-2"
    >
      <span
        class="whitespace-nowrap text-right text-[5px] leading-[1.05] text-black-700 sm:text-[7px] md:text-[10px]"
      >
        <span class="font-semibold">Merajut Cerita</span>
        <br />
        <span class="font-medium">Tumbuh Bersama</span>
      </span>

      <span
        aria-hidden="true"
        class="text-sm leading-none font-black text-blue-700 sm:text-xl md:text-3xl"
      >
        #
      </span>
    </div>
  </div>
</header>