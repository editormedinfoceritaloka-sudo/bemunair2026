<script lang="ts">
  const year = 2026;

  const navigation = [
    {
      label: 'Home',
      href: '/#home'
    },
    {
      label: 'Speech',
      href: '/#speech'
    },
    {
      label: 'Cabinet',
      href: '/#cabinet'
    },
    {
      label: 'Calendar',
      href: '/#calendar'
    },
    {
      label: 'News',
      href: '/#news'
    }
  ] as const;

  const socials = [
    {
      label: 'Instagram',
      username: '@bem_unair',
      href: 'https://www.instagram.com/bem_unair/',
      icon: 'instagram'
    },
    {
      label: 'TikTok',
      username: '@bem_unair',
      href: 'https://www.tiktok.com/@bem_unair',
      icon: 'tiktok'
    },
    {
      label: 'YouTube',
      username: 'BEMUniversitasAirlangga',
      href: 'https://www.youtube.com/@BEMUniversitasAirlangga',
      icon: 'youtube'
    },
    {
      label: 'X',
      username: '@bemunair_id',
      href: 'https://x.com/bemunair_id',
      icon: 'x'
    }
  ] as const;
</script>

<footer
  class="relative overflow-hidden bg-linear-to-b from-blue-800 via-blue-900 to-blue-950 text-white"
>
  <div
    class="relative mx-auto w-full max-w-7xl px-5 pb-7 pt-16 sm:px-8 sm:pt-20 lg:px-10 lg:pt-24"
  >
    <div
      class="grid gap-12 md:grid-cols-2 lg:grid-cols-[1.5fr_0.7fr_0.8fr] lg:items-start lg:gap-10"
    >
      <div class="flex items-center gap-4 sm:gap-5">
        <img
          src="/logo/logo-kabinet.webp"
          alt="Logo Kabinet Cerita Loka"
          class="size-24 shrink-0 object-contain sm:size-28 lg:size-32"
          loading="lazy"
          decoding="async"
        />

        <div class="flex flex-col">
          <h2
            class="text-3xl leading-tight font-black tracking-tight sm:text-4xl"
          >
            Cerita Loka
          </h2>

          <span
            class="mt-1 text-2xl font-bold tracking-tight text-blue-100 sm:text-3xl"
          >
            {year}
          </span>
        </div>
      </div>

      <div>
        <h3
          class="mb-4 text-sm font-extrabold tracking-wide text-white"
        >
          Links
        </h3>

        <nav aria-label="Navigasi footer">
          <ul class="space-y-2.5">
            {#each navigation as item (item.href)}
              <li>
                <a
                  href={item.href}
                  class="inline-flex text-sm text-blue-100 transition-colors duration-200 hover:text-white hover:underline hover:underline-offset-4 focus-visible:rounded-sm focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-white"
                >
                  {item.label}
                </a>
              </li>
            {/each}
          </ul>
        </nav>
      </div>

      <div class="space-y-6 text-sm">
        <div>
          <h3
            class="mb-3 font-extrabold tracking-wide text-white"
          >
            Address
          </h3>

          <p class="max-w-xs leading-relaxed text-blue-100">
            Kampus C Universitas Airlangga,
            <br />
            Mulyorejo, Surabaya.
          </p>
        </div>

        <div>
          <h3
            class="mb-3 font-extrabold tracking-wide text-white"
          >
            Contact us
          </h3>

          <div class="space-y-1">
            <p class="text-blue-100">
              Email:
            </p>

            <a
              href="mailto:bemunair@unair.ac.id"
              class="inline-block text-blue-100 transition-colors duration-200 hover:text-white hover:underline hover:underline-offset-4 focus-visible:rounded-sm focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-white"
            >
              bemunair@unair.ac.id
            </a>
          </div>
        </div>
      </div>
    </div>

    <div
      class="mt-12 border-t border-white/25 pt-6 lg:mt-16"
    >
      <div
        class="flex flex-col items-center gap-6 lg:flex-row lg:justify-between lg:gap-8"
      >
        <p
          class="order-3 text-center text-xs leading-relaxed text-blue-100 lg:order-1 lg:text-left"
        >
          © BEM UNAIR {year} | All Rights Reserved
        </p>

        <p
          class="order-1 text-center text-sm leading-relaxed text-white lg:order-2"
        >
          <span class="font-extrabold">
            Merajut
          </span>

          Cerita Tumbuh

          <span class="font-extrabold">
            Bersama
          </span>
        </p>

        <div
          class="order-2 flex flex-wrap items-center justify-center gap-3 lg:order-3 lg:justify-end"
        >
          {#each socials as social (social.icon)}
            <a
              href={social.href}
              target="_blank"
              rel="noopener noreferrer"
              aria-label={`${social.label} ${social.username}`}
              title={`${social.label} — ${social.username}`}
              class="group relative flex size-10 shrink-0 items-center justify-center rounded-full border border-white/30 bg-white shadow-sm transition-all duration-300 ease-out hover:-translate-y-1 hover:scale-110 hover:shadow-xl focus-visible:outline-2 focus-visible:outline-offset-4 focus-visible:outline-white active:scale-95"
            >
              {#if social.icon === 'instagram'}
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="none"
                  class="size-[21px] transition-transform duration-300 group-hover:scale-110"
                  aria-hidden="true"
                >
                  <defs>
                    <linearGradient
                      id="footer-instagram-gradient"
                      x1="0%"
                      y1="100%"
                      x2="100%"
                      y2="0%"
                    >
                      <stop offset="0%" stop-color="#FEDA75" />
                      <stop offset="25%" stop-color="#FA7E1E" />
                      <stop offset="50%" stop-color="#D62976" />
                      <stop offset="75%" stop-color="#962FBF" />
                      <stop offset="100%" stop-color="#4F5BD5" />
                    </linearGradient>
                  </defs>

                  <rect
                    x="2"
                    y="2"
                    width="20"
                    height="20"
                    rx="5.5"
                    stroke="url(#footer-instagram-gradient)"
                    stroke-width="2.5"
                  />

                  <circle
                    cx="12"
                    cy="12"
                    r="4.2"
                    stroke="url(#footer-instagram-gradient)"
                    stroke-width="2.5"
                  />

                  <circle
                    cx="17.6"
                    cy="6.5"
                    r="1.3"
                    fill="#962FBF"
                  />
                </svg>

              {:else if social.icon === 'tiktok'}
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  class="size-[21px] transition-transform duration-300 group-hover:scale-110"
                  aria-hidden="true"
                >
                  <path
                    fill="#25F4EE"
                    d="M16.7 2h-4.2v13.2a3.3 3.3 0 1 1-3.1-3.3c.4 0 .8.1 1.2.2V7.9c-.4-.1-.8-.1-1.2-.1a7.4 7.4 0 1 0 7.4 7.4V8.6a10.5 10.5 0 0 0 5.2 1.4V5.8C19.2 5.6 17.3 4.2 16.7 2Z"
                    transform="translate(-0.6, 0.3)"
                  />

                  <path
                    fill="#FE2C55"
                    d="M16.7 2h-4.2v13.2a3.3 3.3 0 1 1-3.1-3.3c.4 0 .8.1 1.2.2V7.9c-.4-.1-.8-.1-1.2-.1a7.4 7.4 0 1 0 7.4 7.4V8.6a10.5 10.5 0 0 0 5.2 1.4V5.8C19.2 5.6 17.3 4.2 16.7 2Z"
                    transform="translate(0.5, -0.2)"
                  />

                  <path
                    fill="#111111"
                    d="M16.7 2h-4.2v13.2a3.3 3.3 0 1 1-3.1-3.3c.4 0 .8.1 1.2.2V7.9c-.4-.1-.8-.1-1.2-.1a7.4 7.4 0 1 0 7.4 7.4V8.6a10.5 10.5 0 0 0 5.2 1.4V5.8C19.2 5.6 17.3 4.2 16.7 2Z"
                  />
                </svg>
              {:else if social.icon === 'youtube'}
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="#FF0033"
                  class="size-[23px] transition-transform duration-300 group-hover:scale-110"
                  aria-hidden="true"
                >
                  <path
                    d="M21.582 6.186a2.506 2.506 0 0 0-1.765-1.774C18.263 4 12 4 12 4s-6.263 0-7.817.412a2.506 2.506 0 0 0-1.765 1.774C2 7.752 2 12 2 12s0 4.248.418 5.814a2.506 2.506 0 0 0 1.765 1.774C5.737 20 12 20 12 20s6.263 0 7.817-.412a2.506 2.506 0 0 0 1.765-1.774C22 16.248 22 12 22 12s0-4.248-.418-5.814zM10 15.5v-7l6 3.5-6 3.5z"
                    fill-rule="evenodd"
                  />
                </svg>

              {:else if social.icon === 'x'}
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  fill="#000000"
                  class="size-[19px] transition-transform duration-300 group-hover:scale-110"
                  aria-hidden="true"
                >
                  <path
                    d="M18.901 1.153h3.68l-8.04 9.19L24 22.846h-7.406l-5.8-7.584-6.64 7.584H.47l8.6-9.83L0 1.154h7.594l5.243 6.932 6.064-6.933ZM17.61 20.644h2.039L6.482 3.24H4.295L17.61 20.644Z"
                  />
                </svg>
              {/if}
            </a>
          {/each}
        </div>
      </div>
    </div>
  </div>
</footer>