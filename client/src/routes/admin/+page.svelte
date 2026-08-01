<script lang="ts">
  import AdminCalendar from '$lib/admin/AdminCalendar.svelte';
  import EmptyState from '$lib/admin/EmptyState.svelte';
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import StatusBadge from '$lib/admin/StatusBadge.svelte';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { resolve } from '$app/paths';
  import { Button } from '$lib/components/ui/button';
  import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
  import { ArrowRight, Clock3, FileImage, Files, Newspaper, UserCheck, Users } from '@lucide/svelte';

  let { data } = $props();
  const activeStatuses = ['SUBMITTED', 'PENDING_REVIEW', 'REVISION_REQUIRED', 'REVISION_SUBMITTED', 'APPROVED', 'SCHEDULED', 'PENDING', 'IN_REVIEW'];
  const open = (rows: { status: string }[]) => rows.filter((row) => activeStatuses.includes(row.status)).length;
  const current = $derived(data.queue.find((queueItem: { is_current: boolean }) => queueItem.is_current));
  const metrics = $derived(data.user.role === 'ADMIN_MEDINFO'
    ? [
        { label: 'Pengguna aktif', value: data.users.length, icon: Users, href: '/admin/users' },
        { label: 'Konten terbuka', value: open(data.content), icon: FileImage, href: '/admin/content-submissions' },
        { label: 'Artikel terbit', value: data.articles.filter((article: { status: string }) => article.status === 'PUBLISHED').length, icon: Newspaper, href: '/admin/articles' }
      ]
    : [
        { label: 'Pengajuan media', value: data.content.length, icon: FileImage, href: '/admin/content-submissions' },
        { label: 'Media diproses', value: open(data.content), icon: Clock3, href: '/admin/content-submissions' },
        { label: 'Template surat', value: data.templates.length, icon: Files, href: '/admin/letter-submissions' }
      ]);
  const deadlines = $derived(data.content
    .map((item: { deadline?: string; title: string; ministry?: string; status: string; id: number }) => ({ ...item, kind: 'Konten', name: item.title }))
    .filter((item: { deadline?: string; status: string }) => item.deadline && activeStatuses.includes(item.status))
    .sort((first: { deadline?: string }, second: { deadline?: string }) => +new Date(first.deadline || '') - +new Date(second.deadline || ''))
    .slice(0, 5));
</script>

<svelte:head><title>Dashboard Admin - BEM UNAIR</title></svelte:head>

<PageHeader title={data.user.role === 'ADMIN_MEDINFO' ? 'Ringkasan operasional' : 'Workspace kementerian'} description={data.user.role === 'ADMIN_MEDINFO' ? 'Pantau pekerjaan, deadline, dan publikasi dalam satu workspace.' : 'Kelola pengajuan media dan akses template surat PDF.'}>
  {#if data.user.role === 'ADMIN_MEDINFO'}
    <Button href="/admin/articles/new">Tulis artikel</Button>
  {:else}
    <Button href="/admin/content-submissions/new/content/form" class="bg-blue-500">Buat pengajuan media</Button>
  {/if}
</PageHeader>

{#if data.partialFailure}
  <Alert class="mb-5"><AlertDescription>Sebagian data belum dapat dimuat. Modul yang tersedia tetap ditampilkan.</AlertDescription></Alert>
{/if}

<div class="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
  {#each metrics as item (item.label)}
    <a href={resolve(item.href as '/admin')} class="group no-underline">
      <Card class="h-full border-border/80 bg-card shadow-sm transition hover:-translate-y-0.5 hover:border-primary/30 hover:shadow-md">
        <CardContent class="flex items-start justify-between p-5">
          <div><p class="text-sm font-medium text-muted-foreground">{item.label}</p><p class="mt-2 text-3xl font-bold text-foreground">{item.value}</p></div>
          <div class="grid size-10 place-items-center rounded-xl bg-primary/10 text-primary transition-colors group-hover:bg-primary group-hover:text-primary-foreground"><item.icon class="size-5" /></div>
        </CardContent>
      </Card>
    </a>
  {/each}
</div>

<AdminCalendar content={data.content} letters={[]} articles={data.articles} />

<div class="mt-6 grid gap-6 xl:grid-cols-[1.4fr_.6fr]">
  <Card class="border-border/80 shadow-sm">
    <CardHeader><CardTitle class="flex items-center gap-2 text-base"><span class="grid size-7 place-items-center rounded-lg bg-warning/10 text-warning"><Clock3 class="size-4" /></span>Deadline terdekat</CardTitle></CardHeader>
    <CardContent>
      {#if deadlines.length}
        <div class="divide-y divide-border">
          {#each deadlines as row (row.id)}
            <a href={resolve((row.kind === 'Konten' ? `/admin/content-submissions/${row.id}` : '/admin') as '/admin')} class="flex items-center justify-between gap-4 rounded-lg px-2 py-3 text-inherit no-underline transition-colors hover:bg-muted">
              <div class="min-w-0"><p class="truncate font-medium">{row.name}</p><p class="text-xs text-muted-foreground">{row.kind} - {row.ministry || 'Kementerian'} - {new Intl.DateTimeFormat('id-ID', { dateStyle: 'medium' }).format(new Date(row.deadline || ''))}</p></div>
              <StatusBadge status={row.status} />
            </a>
          {/each}
        </div>
      {:else}
        <EmptyState title="Tidak ada deadline aktif" description="Pengajuan media baru akan muncul di sini." />
      {/if}
    </CardContent>
  </Card>

  {#if data.user.role === 'ADMIN_MEDINFO'}
    <Card class="border-border/80 shadow-sm">
      <CardHeader><CardTitle class="flex items-center gap-2 text-base"><span class="grid size-7 place-items-center rounded-lg bg-success/10 text-success"><UserCheck class="size-4" /></span>PJ berjalan</CardTitle></CardHeader>
      <CardContent>
        {#if current}
          <div class="rounded-xl border border-border bg-muted/40 p-4"><p class="text-lg font-bold text-foreground">{current.user?.name || `User #${current.user_id}`}</p><p class="mt-1 text-sm text-muted-foreground">Posisi {current.position} dari {data.queue.length} anggota antrean.</p></div>
          <Button href="/admin/medinfo-queue" variant="outline" class="mt-4 w-full">Kelola antrean <ArrowRight /></Button>
        {:else}
          <EmptyState title="Antrean belum aktif" description="Tambahkan PJ Medinfo untuk assignment otomatis." />
        {/if}
      </CardContent>
    </Card>
  {:else}
    <Card class="border-border/80 shadow-sm">
      <CardHeader><CardTitle class="flex items-center gap-2 text-base"><span class="grid size-7 place-items-center rounded-lg bg-primary/10 text-primary"><Files class="size-4" /></span>Template surat</CardTitle></CardHeader>
      <CardContent>
        <p class="text-sm leading-6 text-muted-foreground">Download template surat PDF yang tersedia sebelum menyiapkan kebutuhan administrasi kementerian.</p>
        <Button href="/admin/letter-submissions" variant="outline" class="mt-4 w-full">Lihat semua template <ArrowRight /></Button>
      </CardContent>
    </Card>
  {/if}
</div>
