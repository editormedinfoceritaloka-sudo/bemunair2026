<script lang="ts">
  import AtSign from '@lucide/svelte/icons/at-sign';
  import Camera from '@lucide/svelte/icons/camera';
  import MessageCircle from '@lucide/svelte/icons/message-circle';
  import Music2 from '@lucide/svelte/icons/music-2';
  import Play from '@lucide/svelte/icons/play';

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
      label: 'TikTok',
      href: '#',
      icon: Music2
    },
    {
      label: 'Instagram',
      href: '#',
      icon: Camera
    },
    {
      label: 'LINE',
      href: '#',
      icon: MessageCircle
    },
    {
      label: 'YouTube',
      href: '#',
      icon: Play
    },
    {
      label: 'X',
      href: '#',
      icon: AtSign
    }
  ] as const;
</script>

<footer class="bg-linear-to-b from-white via-blue-300 to-blue-900 text-white">
  <div class="mx-auto w-full max-w-7xl px-5 pb-8 pt-24 sm:px-8 lg:px-10">
    <div
      class="grid gap-12 md:grid-cols-2 lg:grid-cols-[1.5fr_0.7fr_0.8fr] lg:items-center"
    >
      <div class="flex items-center gap-5">
        <img
          src="/logo/logo-kabinet.png"
          alt="Logo Kabinet Cerita Loka"
          class="size-28 shrink-0 object-contain sm:size-32"
        />

        <h2 class="max-w-xs text-3xl font-black leading-tight sm:text-4xl">
          Cerita Loka
          <br />
          {year}
        </h2>
      </div>

      <div>
        <h3 class="mb-3 text-sm font-extrabold">
          Links
        </h3>

        <nav aria-label="Navigasi footer">
          <ul class="space-y-1.5 text-sm">
            {#each navigation as item (item.href)}
              <li>
                <a
                  href={item.href}
                  class="transition hover:text-blue-100 hover:underline"
                >
                  {item.label}
                </a>
              </li>
            {/each}
          </ul>
        </nav>
      </div>

      <div class="space-y-5 text-sm">
        <div>
          <h3 class="mb-2 font-extrabold">
            Address
          </h3>

          <p class="max-w-xs leading-relaxed text-white/90">
            Kampus C Universitas Airlangga,
            <br />
            Mulyorejo, Surabaya.
          </p>
        </div>

        <div>
          <h3 class="mb-1 font-extrabold">
            Contact us
          </h3>

          <p class="text-white/90">
            Email:
            <br />

            <a
              href="mailto:bemunair@unair.ac.id"
              class="transition hover:text-blue-100 hover:underline"
            >
              bemunair@unair.ac.id
            </a>
          </p>
        </div>
      </div>
    </div>

    <div class="mt-10 border-t-2 border-white/90 pt-4">
      <div
        class="flex flex-col gap-5 text-xs sm:flex-row sm:items-center sm:justify-between"
      >
        <p class="text-center text-white/90 sm:text-left">
          © BEM UNAIR {year} | All Rights Reserved
        </p>

        <p class="text-center text-sm">
          <span class="font-extrabold">Merajut</span>
          Cerita Tumbuh
          <span class="font-extrabold">Bersama</span>
        </p>

        <div class="flex items-center justify-center gap-2 sm:justify-end">
          {#each socials as social (social.label)}
            {@const Icon = social.icon}

            <a
              href={social.href}
              aria-label={social.label}
              title={social.label}
              class="grid size-8 place-items-center rounded-full bg-blue-100/80 text-blue-700 transition duration-200 hover:-translate-y-1 hover:bg-white"
            >
              <Icon size={17} strokeWidth={2.4} />
            </a>
          {/each}
        </div>
      </div>
    </div>
  </div>
</footer>