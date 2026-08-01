<script lang="ts">
  import { page } from '$app/state';
  import { resolve } from '$app/paths';
  import type { User } from '$lib/types';
  import {
    BriefcaseBusiness,
    Building2,
    FileImage,
    Files,
    LayoutDashboard,
    ListOrdered,
    LogOut,
    Newspaper,
    ShieldCheck,
    Users
  } from '@lucide/svelte';

  type NavigationItem = {
    href: string;
    label: string;
    icon: typeof LayoutDashboard;
    medinfoOnly?: boolean;
    adminOnly?: boolean;
  };

  let { user }: { user: User } = $props();

  const navigation: { label: string; items: NavigationItem[] }[] = [
    {
      label: 'Workspace',
      items: [
        { href: '/admin', label: 'Ringkasan', icon: LayoutDashboard },
        { href: '/admin/users', label: 'Pengguna', icon: Users, medinfoOnly: true },
        { href: '/admin/ministries', label: 'Kementerian', icon: Building2, medinfoOnly: true },
        { href: '/admin/work-programs', label: 'Program Kerja', icon: BriefcaseBusiness }
      ]
    },
    {
      label: 'Layanan',
      items: [
        { href: '/admin/content-submissions?type=CONTENT', label: 'Pengajuan Media', icon: FileImage },
        { href: '/admin/content-submissions?type=ARTICLE', label: 'Pengajuan Artikel', icon: Newspaper },
        { href: '/admin/letter-submissions', label: 'Template Surat', icon: Files, adminOnly: true },
        { href: '/admin/letter-templates', label: 'Manajemen Template Surat', icon: Files, medinfoOnly: true },
        { href: '/admin/medinfo-queue', label: 'Antrean PJ', icon: ListOrdered, medinfoOnly: true }
      ]
    }
  ];

  const canShow = (item: NavigationItem) =>
    (!item.medinfoOnly || user.role === 'ADMIN_MEDINFO') &&
    (!item.adminOnly || user.role === 'ADMIN');

  const active = (href: string) => {
    const [path = '', query] = href.split('?');

    if (path === '/admin') return page.url.pathname === path;

    return page.url.pathname.startsWith(path) && (!query || page.url.search === '?' + query);
  };
</script>

<aside class="hidden h-svh w-64 shrink-0 flex-col border-r border-sidebar-border bg-sidebar md:flex">
  <div class="border-b border-sidebar-border px-4 py-4">
    <a href={resolve('/admin')} class="flex items-center gap-3 rounded-xl no-underline">
      <div class="relative grid size-10 shrink-0 place-items-center rounded-xl border border-blue-100 bg-white p-1 shadow-sm">
        <img src="/logo/logo-bem.png" alt="Logo BEM UNAIR 2026 Kabinet Cerita Loka" class="size-full object-contain" />
        <span class="absolute -bottom-0.5 -right-0.5 size-3 rounded-full border-2 border-sidebar bg-orange-500"></span>
      </div>
      <div class="min-w-0">
        <p class="truncate text-sm font-extrabold tracking-wide text-blue-900">BEM UNAIR</p>
        <p class="truncate text-[11px] font-medium text-black-300">Admin Workspace</p>
      </div>
    </a>
  </div>

  <nav class="flex-1 overflow-y-auto px-3 py-4" aria-label="Navigasi admin">
    {#each navigation as section (section.label)}
      <section class="mb-5">
        <p class="px-2 text-[10px] font-bold uppercase tracking-[0.14em] text-black-200">{section.label}</p>
        <div class="mt-2 grid gap-1">
          {#each section.items.filter(canShow) as item (item.href)}
            <a
              href={resolve(item.href as '/admin')}
              aria-current={active(item.href) ? 'page' : undefined}
              class={`flex h-10 items-center gap-3 rounded-lg px-3 text-sm no-underline transition-colors ${active(item.href) ? 'bg-blue-50 font-semibold text-blue-700' : 'text-black-300 hover:bg-blue-50 hover:text-blue-700'}`}
            >
              <item.icon class="size-4 shrink-0" />
              <span class="truncate">{item.label}</span>
            </a>
          {/each}
        </div>
      </section>
    {/each}
  </nav>

  <div class="border-t border-sidebar-border p-3">
    <div class="rounded-xl border border-white-600 bg-card-200 p-3">
      <div class="flex min-w-0 items-center gap-2.5">
        <div class="grid size-9 shrink-0 place-items-center rounded-full border border-blue-100 bg-blue-50 text-xs font-bold text-blue-700">
          {user.name.slice(0, 2).toUpperCase()}
        </div>
        <div class="min-w-0 flex-1">
          <p class="truncate text-sm font-semibold text-black-500">{user.name}</p>
          <p class="mt-0.5 flex items-center gap-1 text-[11px] text-black-300">
            <ShieldCheck class="size-3" />
            {user.role === 'ADMIN_MEDINFO' ? 'Admin Medinfo' : 'Admin Kementerian'}
          </p>
        </div>
        <form method="POST" action="/admin/logout">
          <button class="rounded-lg p-2 text-black-200 transition-colors hover:bg-red-50 hover:text-red-600 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-sidebar-ring" aria-label="Keluar" title="Keluar">
            <LogOut class="size-4" />
          </button>
        </form>
      </div>
    </div>
  </div>
</aside>
