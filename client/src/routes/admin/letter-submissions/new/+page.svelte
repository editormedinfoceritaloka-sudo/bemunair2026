<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte'; import { Button } from '$lib/components/ui/button'; import { Input } from '$lib/components/ui/input'; import { Textarea } from '$lib/components/ui/textarea'; import { Label } from '$lib/components/ui/label'; import { Alert, AlertDescription } from '$lib/components/ui/alert';
  let {data,form}=$props(); let type=$state(''); let subject=$state(''); let body=$state(''); let selected=$state('');
  function choose(){const item=data.templates.find((t)=>String(t.id)===selected);if(item){type=item.type;subject=item.subject;body=item.body}}
</script>
<PageHeader title="Pengajuan surat baru" description="Gunakan template atau susun permintaan surat dari awal."/>
{#if form?.error}<Alert variant="destructive" class="mb-4"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}
<form method="POST" class="max-w-4xl space-y-6 rounded-xl border bg-card p-5 md:p-8">
  <div class="rounded-lg border border-dashed bg-background p-4"><Label for="template">Mulai dari template <span class="text-black-200">(opsional)</span></Label><select id="template" bind:value={selected} onchange={choose} class="mt-2 h-9 w-full rounded-lg border bg-card px-3 text-sm"><option value="">Tanpa template</option>{#each data.templates as item}<option value={item.id}>{item.name} · {item.type}</option>{/each}</select></div>
  <div class="grid gap-4 sm:grid-cols-2"><div class="space-y-2"><Label for="ministry">Kementerian</Label><Input id="ministry" name="ministry" required/></div><div class="space-y-2"><Label for="letter_type">Jenis surat</Label><Input id="letter_type" name="letter_type" bind:value={type} required/></div><div class="space-y-2 sm:col-span-2"><Label for="subject">Perihal</Label><Input id="subject" name="subject" bind:value={subject} required/></div><div class="space-y-2 sm:col-span-2"><Label for="body">Isi surat</Label><Textarea id="body" name="body" bind:value={body} rows={14} required/></div><div class="space-y-2"><Label for="deadline">Deadline (WIB)</Label><Input id="deadline" name="deadline" type="datetime-local" required/></div></div>
  <div class="flex justify-end gap-2"><Button href="/admin/letter-submissions" variant="outline">Batal</Button><Button type="submit" class="bg-blue-500">Kirim pengajuan</Button></div>
</form>
