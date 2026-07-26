<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte'; import StatusBadge from '$lib/admin/StatusBadge.svelte'; import EmptyState from '$lib/admin/EmptyState.svelte'; import AdminCalendar from '$lib/admin/AdminCalendar.svelte';
  import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'; import { Button } from '$lib/components/ui/button'; import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { Users, FileImage, Mail, Newspaper, ArrowRight, Clock3, UserCheck } from '@lucide/svelte';
  let { data } = $props();
  const activeStatuses = ['SUBMITTED', 'PENDING_REVIEW', 'REVISION_REQUIRED', 'REVISION_SUBMITTED', 'APPROVED', 'SCHEDULED', 'PENDING', 'IN_REVIEW'];
  const open = (rows: {status:string}[]) => rows.filter((r) => activeStatuses.includes(r.status)).length;
  const current = $derived(data.queue.find((q:any)=>q.is_current));
  const metrics = $derived(data.user.role === "ADMIN_MEDINFO" ? [
    { label:'Pengguna aktif', value:data.users.length, icon:Users, href:'/admin/users' },
    { label:'Konten terbuka', value:open(data.content), icon:FileImage, href:'/admin/content-submissions' },
    { label:'Surat terbuka', value:open(data.letters), icon:Mail, href:'/admin/letter-submissions' },
    { label:'Artikel terbit', value:data.articles.filter((a:any)=>a.status==='PUBLISHED').length, icon:Newspaper, href:'/admin/articles' }
  ] : [
    { label:'Pengajuan media', value:data.content.length, icon:FileImage, href:'/admin/content-submissions' },
    { label:'Media diproses', value:open(data.content), icon:Clock3, href:'/admin/content-submissions' },
    { label:'Pengajuan surat', value:data.letters.length, icon:Mail, href:'/admin/letter-submissions' },
    { label:'Surat diproses', value:open(data.letters), icon:Clock3, href:'/admin/letter-submissions' }
  ]);
  const deadlines = $derived([...data.content.map((x:any)=>({...x,kind:'Konten',name:x.title})), ...data.letters.map((x:any)=>({...x,kind:'Surat',name:x.subject}))].filter((x:any)=>x.deadline && activeStatuses.includes(x.status)).sort((a:any,b:any)=>+new Date(a.deadline)-+new Date(b.deadline)).slice(0,5));
</script>
<svelte:head><title>Dashboard Admin · BEM UNAIR</title></svelte:head>
<PageHeader title={data.user.role === "ADMIN_MEDINFO" ? "Ringkasan operasional" : "Progress pengajuan"} description={data.user.role === "ADMIN_MEDINFO" ? "Pantau pekerjaan, deadline, dan publikasi dalam satu workspace." : "Pantau timeline pengajuan kementerian dan buat request baru."}>{#if data.user.role === "ADMIN_MEDINFO"}<Button href="/admin/articles/new">Tulis artikel</Button>{:else}<Button href="/admin/content-submissions/new/select" class="bg-blue-500">Buat pengajuan media</Button>{/if}</PageHeader>
{#if data.partialFailure}<Alert class="mb-5"><AlertDescription>Sebagian data belum dapat dimuat. Modul yang tersedia tetap ditampilkan.</AlertDescription></Alert>{/if}
<div class="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">{#each metrics as item}<a href={item.href} class="group no-underline"><Card class="h-full border-border/80 bg-card shadow-sm transition hover:-translate-y-0.5 hover:border-primary/30 hover:shadow-md"><CardContent class="flex items-start justify-between p-5"><div><p class="text-sm font-medium text-muted-foreground">{item.label}</p><p class="mt-2 text-3xl font-bold text-foreground">{item.value}</p></div><div class="grid size-10 place-items-center rounded-xl bg-primary/10 text-primary transition-colors group-hover:bg-primary group-hover:text-primary-foreground"><item.icon class="size-5" /></div></CardContent></Card></a>{/each}</div>
<AdminCalendar content={data.content} letters={data.letters} articles={data.articles} />
<div class="mt-6 grid gap-6 xl:grid-cols-[1.4fr_.6fr]">
  <Card class="border-border/80 shadow-sm"><CardHeader><CardTitle class="flex items-center gap-2 text-base"><span class="grid size-7 place-items-center rounded-lg bg-warning/10 text-warning"><Clock3 class="size-4"/></span>Deadline terdekat</CardTitle></CardHeader><CardContent>{#if deadlines.length}<div class="divide-y divide-border">{#each deadlines as row}<a href={row.kind==='Konten'?`/admin/content-submissions/${row.id}`:`/admin/letter-submissions/${row.id}`} class="flex items-center justify-between gap-4 rounded-lg px-2 py-3 text-inherit no-underline transition-colors hover:bg-muted"><div class="min-w-0"><p class="truncate font-medium">{row.name}</p><p class="text-xs text-muted-foreground">{row.kind} · {row.ministry} · {new Intl.DateTimeFormat('id-ID',{dateStyle:'medium'}).format(new Date(row.deadline))}</p></div><StatusBadge status={row.status}/></a>{/each}</div>{:else}<EmptyState title="Tidak ada deadline aktif" description="Pengajuan baru akan muncul di sini." />{/if}</CardContent></Card>
  {#if data.user.role === "ADMIN_MEDINFO"}
    <Card class="border-border/80 shadow-sm"><CardHeader><CardTitle class="flex items-center gap-2 text-base"><span class="grid size-7 place-items-center rounded-lg bg-success/10 text-success"><UserCheck class="size-4"/></span>PJ berjalan</CardTitle></CardHeader><CardContent>{#if current}<div class="rounded-xl border border-border bg-muted/40 p-4"><p class="text-lg font-bold text-foreground">{current.user?.name || `User #${current.user_id}`}</p><p class="mt-1 text-sm text-muted-foreground">Posisi {current.position} dari {data.queue.length} anggota antrean.</p></div><Button href="/admin/medinfo-queue" variant="outline" class="mt-4 w-full">Kelola antrean <ArrowRight/></Button>{:else}<EmptyState title="Antrean belum aktif" description="Tambahkan PJ Medinfo untuk assignment otomatis." />{/if}</CardContent></Card>
  {:else}
    <Card class="border-border/80 shadow-sm">
      <CardHeader><CardTitle class="flex items-center gap-2 text-base"><span class="grid size-7 place-items-center rounded-lg bg-primary/10 text-primary"><FileImage class="size-4"/></span>Buat pengajuan</CardTitle></CardHeader>
      <CardContent><p class="text-sm leading-6 text-muted-foreground">Ajukan kebutuhan publikasi media atau surat, lalu pantau progresnya melalui timeline.</p><div class="mt-4 grid gap-2"><Button href="/admin/content-submissions/new/select">Pengajuan media <ArrowRight/></Button><Button href="/admin/letter-submissions/new" variant="outline">Pengajuan surat <ArrowRight/></Button></div></CardContent>
    </Card>
  {/if}
</div>
