<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte'; import StatusBadge from '$lib/admin/StatusBadge.svelte'; import EmptyState from '$lib/admin/EmptyState.svelte'; import AdminCalendar from '$lib/admin/AdminCalendar.svelte';
  import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card'; import { Button } from '$lib/components/ui/button'; import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { Users, FileImage, Mail, Newspaper, ArrowRight, Clock3, UserCheck } from '@lucide/svelte';
  let { data } = $props();
  const open = (rows: {status:string}[]) => rows.filter((r) => r.status === 'PENDING' || r.status === 'IN_REVIEW').length;
  const current = $derived(data.queue.find((q:any)=>q.is_current));
  const metrics = $derived([
    { label:'Pengguna aktif', value:data.users.length, icon:Users, href:'/admin/users' },
    { label:'Konten terbuka', value:open(data.content), icon:FileImage, href:'/admin/content-submissions' },
    { label:'Surat terbuka', value:open(data.letters), icon:Mail, href:'/admin/letter-submissions' },
    { label:'Artikel terbit', value:data.articles.filter((a:any)=>a.status==='PUBLISHED').length, icon:Newspaper, href:'/admin/articles' }
  ]);
  const deadlines = $derived([...data.content.map((x:any)=>({...x,kind:'Konten',name:x.title})), ...data.letters.map((x:any)=>({...x,kind:'Surat',name:x.subject}))].filter((x:any)=>x.deadline && ['PENDING','IN_REVIEW'].includes(x.status)).sort((a:any,b:any)=>+new Date(a.deadline)-+new Date(b.deadline)).slice(0,5));
</script>
<svelte:head><title>Dashboard Admin · BEM UNAIR</title></svelte:head>
<PageHeader title="Ringkasan operasional" description="Pantau pekerjaan, deadline, dan publikasi dalam satu workspace."><Button href="/admin/articles/new">Tulis artikel</Button></PageHeader>
{#if data.partialFailure}<Alert class="mb-5"><AlertDescription>Sebagian data belum dapat dimuat. Modul yang tersedia tetap ditampilkan.</AlertDescription></Alert>{/if}
<div class="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">{#each metrics as item}<a href={item.href} class="group no-underline"><Card class="h-full border-border/80 bg-card shadow-sm transition hover:-translate-y-0.5 hover:border-primary/30 hover:shadow-md"><CardContent class="flex items-start justify-between p-5"><div><p class="text-sm font-medium text-muted-foreground">{item.label}</p><p class="mt-2 text-3xl font-bold text-foreground">{item.value}</p></div><div class="grid size-10 place-items-center rounded-xl bg-primary/10 text-primary transition-colors group-hover:bg-primary group-hover:text-primary-foreground"><item.icon class="size-5" /></div></CardContent></Card></a>{/each}</div>
<AdminCalendar content={data.content} letters={data.letters} articles={data.articles} />
<div class="mt-6 grid gap-6 xl:grid-cols-[1.4fr_.6fr]">
  <Card class="border-border/80 shadow-sm"><CardHeader><CardTitle class="flex items-center gap-2 text-base"><span class="grid size-7 place-items-center rounded-lg bg-warning/10 text-warning"><Clock3 class="size-4"/></span>Deadline terdekat</CardTitle></CardHeader><CardContent>{#if deadlines.length}<div class="divide-y divide-border">{#each deadlines as row}<a href={row.kind==='Konten'?`/admin/content-submissions/${row.id}`:`/admin/letter-submissions/${row.id}`} class="flex items-center justify-between gap-4 rounded-lg px-2 py-3 text-inherit no-underline transition-colors hover:bg-muted"><div class="min-w-0"><p class="truncate font-medium">{row.name}</p><p class="text-xs text-muted-foreground">{row.kind} · {row.ministry} · {new Intl.DateTimeFormat('id-ID',{dateStyle:'medium'}).format(new Date(row.deadline))}</p></div><StatusBadge status={row.status}/></a>{/each}</div>{:else}<EmptyState title="Tidak ada deadline aktif" description="Pengajuan baru akan muncul di sini." />{/if}</CardContent></Card>
  <Card class="border-border/80 shadow-sm"><CardHeader><CardTitle class="flex items-center gap-2 text-base"><span class="grid size-7 place-items-center rounded-lg bg-success/10 text-success"><UserCheck class="size-4"/></span>PJ berjalan</CardTitle></CardHeader><CardContent>{#if current}<div class="rounded-xl border border-border bg-muted/40 p-4"><p class="text-lg font-bold text-foreground">{current.user?.name || `User #${current.user_id}`}</p><p class="mt-1 text-sm text-muted-foreground">Posisi {current.position} dari {data.queue.length} anggota antrean.</p></div><Button href="/admin/medinfo-queue" variant="outline" class="mt-4 w-full">Kelola antrean <ArrowRight/></Button>{:else}<EmptyState title="Antrean belum aktif" description="Tambahkan PJ Medinfo untuk assignment otomatis." />{/if}</CardContent></Card>
</div>
