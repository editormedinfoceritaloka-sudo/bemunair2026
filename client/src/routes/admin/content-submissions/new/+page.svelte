<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte'; import { Button } from '$lib/components/ui/button'; import { Input } from '$lib/components/ui/input'; import { Label } from '$lib/components/ui/label'; import { Textarea } from '$lib/components/ui/textarea'; import { Alert, AlertDescription } from '$lib/components/ui/alert';
  let { form } = $props(); let type = $state('FEEDS_REELS');
</script>
<PageHeader title="Pengajuan konten baru" description="Kirim brief lengkap untuk tim Media dan Informasi." />
{#if form?.error}<Alert variant="destructive" class="mb-4"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}
<form method="POST" enctype="multipart/form-data" class="max-w-4xl space-y-6 rounded-xl border bg-card p-5 md:p-8">
  <section><h2 class="font-serif text-xl font-semibold">Informasi utama</h2><p class="mb-4 text-sm text-muted-foreground">Identitas brief dan jenis output publikasi.</p><div class="grid gap-4 sm:grid-cols-2">
    <div class="space-y-2"><Label for="ministry">Kementerian</Label><Input id="ministry" name="ministry" required/></div>
    <div class="space-y-2"><Label for="submission_type">Jenis konten</Label><select id="submission_type" name="submission_type" bind:value={type} class="h-9 w-full rounded-lg border bg-card px-3 text-sm"><option>FEEDS_REELS</option><option>INSTASTORY</option><option>ARTIKEL</option></select></div>
    <div class="space-y-2 sm:col-span-2"><Label for="title">Judul</Label><Input id="title" name="title" required/></div>
    <div class="space-y-2 sm:col-span-2"><Label for="caption">Caption / isi utama</Label><Textarea id="caption" name="caption" rows={6} required/></div>
    <div class="space-y-2 sm:col-span-2"><Label for="brief_link">Tautan brief</Label><Input id="brief_link" name="brief_link" type="url" placeholder="https://..." required/></div>
  </div></section>
  {#if type === 'ARTIKEL'}<section class="grid gap-4 border-t pt-6 sm:grid-cols-2"><div class="space-y-2 sm:col-span-2"><Label for="article_drive_link">Tautan naskah artikel</Label><Input id="article_drive_link" name="article_drive_link" type="url" required/></div></section>{:else}<section class="grid gap-4 border-t pt-6 sm:grid-cols-2"><div class="space-y-2"><Label for="publish_date">Tanggal publikasi</Label><Input id="publish_date" name="publish_date" type="date" required/></div><div class="space-y-2"><Label for="publish_time">Jam publikasi</Label><Input id="publish_time" name="publish_time" type="time" required/></div><div class="space-y-2"><Label for="design_drive_link">Tautan desain Drive</Label><Input id="design_drive_link" name="design_drive_link" type="url" required/></div><div class="space-y-2"><Label for="canva_link">Tautan Canva</Label><Input id="canva_link" name="canva_link" type="url" required/></div><div class="space-y-2 sm:col-span-2"><Label for="add_song">Referensi lagu <span class="text-black-200">(opsional)</span></Label><Input id="add_song" name="add_song"/></div></section>{/if}
  <section class="grid gap-4 border-t pt-6"><div class="space-y-2"><Label for="additional_notes">Catatan tambahan</Label><Textarea id="additional_notes" name="additional_notes" rows={4}/></div></section>
  <div class="flex justify-end gap-2"><Button href="/admin/content-submissions" variant="outline">Batal</Button><Button type="submit" class="bg-blue-500">Kirim pengajuan</Button></div>
</form>
