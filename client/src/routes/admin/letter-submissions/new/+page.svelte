<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Label } from '$lib/components/ui/label';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import type { LetterTemplate } from '$lib/types';
  import { Download, FileText } from '@lucide/svelte';
  import { resolve } from '$app/paths';

  let { data, form } = $props<{ data: { templates: LetterTemplate[]; user: { ministry?: string; role: string } }; form?: { error?: string } }>();
  let type = $state('');
  let subject = $state('');
  let body = $state('');
</script>

<PageHeader title="Pengajuan surat baru" description="Unduh template PDF yang diperlukan, lalu lengkapi pengajuan surat." />
{#if form?.error}<Alert variant="destructive" class="mb-4"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}
<div class="grid gap-6 lg:grid-cols-[320px_1fr]">
  <section class="rounded-xl border bg-card p-5">
    <div class="flex items-center gap-3"><div class="grid size-10 place-items-center rounded-xl bg-red-50 text-red-600"><FileText class="size-5" /></div><div><h2 class="font-bold text-blue-900">Template tersedia</h2><p class="text-xs text-muted-foreground">Gunakan template sesuai jenis surat.</p></div></div>
    <div class="mt-5 space-y-3">{#each data.templates as template (template.id)}<div class="rounded-lg border p-3"><p class="font-semibold text-blue-900">{template.name}</p><p class="mt-1 text-xs text-muted-foreground">{template.type}</p>{#if template.download_url || template.file?.url}<a href={template.download_url || template.file?.url} target="_blank" rel="external noreferrer" class="mt-3 inline-flex items-center gap-2 text-sm font-semibold text-blue-700 hover:underline"><Download class="size-4" />Download PDF</a>{:else}<p class="mt-3 text-xs text-amber-700">File PDF belum tersedia.</p>{/if}</div>{:else}<p class="text-sm text-muted-foreground">Belum ada template PDF.</p>{/each}</div>
    {#if data.user.role === 'ADMIN_MEDINFO'}<Button href={resolve('/admin/letter-templates')} variant="outline" class="mt-5 w-full">Kelola template</Button>{/if}
  </section>
  <form method="POST" class="space-y-5 rounded-xl border bg-card p-5 md:p-8">
    <div class="grid gap-4 sm:grid-cols-2"><div class="space-y-2"><Label for="ministry_view">Kementerian</Label><Input id="ministry_view" value={data.user.ministry || 'Belum ditentukan'} disabled /><input type="hidden" name="ministry" value={data.user.ministry || ''} /></div><div class="space-y-2"><Label for="letter_type">Jenis surat</Label><Input id="letter_type" name="letter_type" bind:value={type} required /></div><div class="space-y-2 sm:col-span-2"><Label for="subject">Perihal</Label><Input id="subject" name="subject" bind:value={subject} required /></div><div class="space-y-2 sm:col-span-2"><Label for="body">Isi/keterangan surat</Label><Textarea id="body" name="body" bind:value={body} rows={14} required /></div><div class="space-y-2"><Label for="deadline">Deadline (WIB)</Label><Input id="deadline" name="deadline" type="datetime-local" required /></div></div>
    <div class="flex justify-end gap-2"><Button href={resolve('/admin/letter-submissions')} variant="outline">Batal</Button><Button type="submit" class="bg-blue-500">Kirim pengajuan</Button></div>
  </form>
</div>
