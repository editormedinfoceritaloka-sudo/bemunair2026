<script lang="ts">
  import { page } from '$app/state';
  import * as Sidebar from '$lib/components/ui/sidebar';
  import { Avatar, AvatarFallback } from '$lib/components/ui/avatar';
  import type { User } from '$lib/types';
  import { LayoutDashboard, Users, FileImage, Mail, ListOrdered, Files, Newspaper, Activity, LogOut, ShieldCheck, Building2, Settings2 } from '@lucide/svelte';

  let { user }: { user: User } = $props();
  const sidebar = Sidebar.useSidebar();

  const navigation = [
    {
      label: 'Workspace',
      items: [
        { href: '/admin', label: 'Ringkasan', icon: LayoutDashboard, medinfoOnly: false },
        { href: '/admin/users', label: 'Pengguna', icon: Users, medinfoOnly: true },
        { href: '/admin/ministries', label: 'Kementerian', icon: Building2, medinfoOnly: true }
      ]
    },
    {
      label: 'Layanan',
      items: [
        { href: '/admin/content-submissions', label: 'Pengajuan Media', icon: FileImage, medinfoOnly: false },
        { href: '/admin/letter-submissions', label: 'Pengajuan Surat', icon: Mail, medinfoOnly: false },
        { href: '/admin/medinfo-queue', label: 'Antrean PJ', icon: ListOrdered, medinfoOnly: true },
        { href: '/admin/letter-templates', label: 'Template Surat', icon: Files, medinfoOnly: true }
      ]
    },
    {
      label: 'Publikasi & sistem',
      items: [
        { href: '/admin/articles', label: 'Artikel', icon: Newspaper, medinfoOnly: true },
        { href: '/admin/system', label: 'Status Sistem', icon: Activity, medinfoOnly: true },
        { href: '/admin/settings/media', label: 'Pengaturan Media', icon: Settings2, medinfoOnly: true }
      ]
    }
  ];

  const active = (href: string) => href === '/admin' ? page.url.pathname === href : page.url.pathname.startsWith(href);
  const initials = $derived(user.name.split(' ').filter(Boolean).slice(0, 2).map((part) => part[0]).join('').toUpperCase());
  const closeMobile = () => {
    if (sidebar.isMobile) sidebar.setOpenMobile(false);
  };
</script>

<Sidebar.Root collapsible="icon" class="overflow-x-hidden border-r border-sidebar-border bg-sidebar">
  <Sidebar.Header class="overflow-hidden border-b border-sidebar-border px-3 py-4 group-data-[collapsible=icon]:p-2">
    <a href="/admin" onclick={closeMobile} class="flex h-11 min-w-0 items-center gap-3 overflow-hidden rounded-xl px-1 no-underline group-data-[collapsible=icon]:h-8 group-data-[collapsible=icon]:justify-center group-data-[collapsible=icon]:p-0">
      <div class="relative grid size-10 shrink-0 place-items-center rounded-xl border border-blue-100 bg-white p-1 shadow-sm group-data-[collapsible=icon]:size-8 group-data-[collapsible=icon]:rounded-lg">
        <img src="/brand/bem-unair-2026-logo.png" alt="Logo BEM UNAIR 2026 Kabinet Cerita Loka" class="size-full object-contain" />
        <span class="absolute -bottom-0.5 -right-0.5 size-3 rounded-full border-2 border-sidebar bg-orange-500 group-data-[collapsible=icon]:hidden"></span>
      </div>
      <div class="min-w-0 group-data-[collapsible=icon]:hidden">
        <p class="truncate text-sm font-extrabold tracking-wide text-blue-900">BEM UNAIR</p>
        <p class="truncate text-[11px] font-medium text-black-300">Admin Workspace</p>
      </div>
    </a>
  </Sidebar.Header>

  <Sidebar.Content class="overflow-x-hidden px-2 py-3 group-data-[collapsible=icon]:px-2 group-data-[collapsible=icon]:py-2">
    {#each navigation as section}
      <Sidebar.Group class="py-1.5 group-data-[collapsible=icon]:p-0 group-data-[collapsible=icon]:py-0.5">
        <Sidebar.GroupLabel class="px-2 text-[10px] font-bold uppercase tracking-[0.14em] text-black-200 group-data-[collapsible=icon]:hidden">{section.label}</Sidebar.GroupLabel>
        <Sidebar.GroupContent>
          <Sidebar.Menu class="gap-1 group-data-[collapsible=icon]:items-center">
            {#each section.items.filter((item) => !item.medinfoOnly || user.role === 'ADMIN_MEDINFO') as item}
              <Sidebar.MenuItem class="group-data-[collapsible=icon]:flex group-data-[collapsible=icon]:justify-center">
                <Sidebar.MenuButton
                  isActive={active(item.href)}
                  tooltipContent={item.label}
                  class="h-10 rounded-lg px-3 text-black-300 hover:bg-blue-50 hover:text-blue-700 data-active:bg-blue-50 data-active:font-semibold data-active:text-blue-700 data-active:shadow-[inset_3px_0_0_var(--orange-500)] group-data-[collapsible=icon]:mx-auto group-data-[collapsible=icon]:size-8! group-data-[collapsible=icon]:justify-center group-data-[collapsible=icon]:p-0!"
                >
                  {#snippet child({ props })}
                    <a {...props} href={item.href} onclick={closeMobile}>
                      <item.icon class="text-current" />
                      <span class="group-data-[collapsible=icon]:hidden">{item.label}</span>
                    </a>
                  {/snippet}
                </Sidebar.MenuButton>
              </Sidebar.MenuItem>
            {/each}
          </Sidebar.Menu>
        </Sidebar.GroupContent>
      </Sidebar.Group>
    {/each}
  </Sidebar.Content>

  <Sidebar.Footer class="overflow-hidden border-t border-sidebar-border p-3 group-data-[collapsible=icon]:p-2">
    <div class="w-full overflow-hidden rounded-xl border border-white-600 bg-card-200 p-2 group-data-[collapsible=icon]:grid group-data-[collapsible=icon]:place-items-center group-data-[collapsible=icon]:border-transparent group-data-[collapsible=icon]:bg-transparent group-data-[collapsible=icon]:p-0">
      <div class="flex min-w-0 items-center gap-2.5 overflow-hidden group-data-[collapsible=icon]:justify-center">
        <Avatar class="size-9 shrink-0 border border-blue-100 group-data-[collapsible=icon]:size-8" title={`${user.name} · ${user.role === 'ADMIN_MEDINFO' ? 'Admin Medinfo' : 'Admin Kementerian'}`}>
          <AvatarFallback class="bg-blue-50 text-xs font-bold text-blue-700">{initials}</AvatarFallback>
        </Avatar>
        <div class="min-w-0 flex-1 group-data-[collapsible=icon]:hidden">
          <p class="truncate text-sm font-semibold text-black-500">{user.name}</p>
          <div class="mt-0.5 flex items-center gap-1 text-[11px] text-black-300"><ShieldCheck class="size-3" />{user.role === 'ADMIN_MEDINFO' ? 'Admin Medinfo' : 'Admin Kementerian'}</div>
        </div>
        <form method="POST" action="/admin/logout" class="shrink-0 group-data-[collapsible=icon]:hidden">
          <button class="rounded-lg p-2 text-black-200 transition-colors hover:bg-red-50 hover:text-red-600 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-sidebar-ring" aria-label="Keluar" title="Keluar">
            <LogOut class="size-4" />
          </button>
        </form>
      </div>
    </div>
  </Sidebar.Footer>
  <Sidebar.Rail />
</Sidebar.Root>
