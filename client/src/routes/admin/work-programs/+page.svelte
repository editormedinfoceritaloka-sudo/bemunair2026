<script lang="ts">
  import { workProgramStatuses } from '$lib/constants/cabinet';
  import { enhance } from '$app/forms';
  import { resolve } from '$app/paths';
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import EmptyState from '$lib/admin/EmptyState.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Label } from '$lib/components/ui/label';
  import type { OrganizationUnit, WorkProgram } from '$lib/types';
  import { useFormSubmit } from '$lib/hooks/use-form-submit.svelte';

  let { data, form } = $props<{ data: { units: OrganizationUnit[]; programs: WorkProgram[]; unitId: number | null }; form?: { error?: string } }>();
  const done = () => useFormSubmit('Program kerja tersimpan');
</script>

<PageHeader title="Program Kerja" description="Kelola program kerja dan status publikasinya." />
{#if form?.error}<p class="mb-5 rounded-xl bg-red-50 p-4 text-sm text-red-700">{form.error}</p>{/if}
<div class="grid gap-6 lg:grid-cols-[1fr_380px]">
  <section class="space-y-4">{#each data.programs as program (program.id)}<article class="rounded-2xl border bg-card p-5"><div class="flex items-start justify-between gap-4"><div><p class="text-xs font-bold uppercase tracking-wider text-orange-600">{program.status}</p><h2 class="mt-1 text-xl font-black text-blue-900">{program.name}</h2><p class="mt-1 text-sm text-muted-foreground">{program.short_description || 'Belum ada ringkasan.'}</p></div><span class={`rounded-full px-3 py-1 text-xs font-bold ${program.is_published ? 'bg-green-50 text-green-700' : 'bg-slate-100 text-slate-600'}`}>{program.is_published ? 'Published' : 'Draft'}</span></div><div class="mt-4 flex flex-wrap justify-end gap-2"><a href={resolve(`/admin/work-programs/${program.id}`)} class="rounded-lg border px-3 py-2 text-sm font-semibold text-blue-800">Kelola detail</a><form method="POST" action="?/publish" use:enhance={done()}><input type="hidden" name="id" value={program.id}/><input type="hidden" name="published" value={program.is_published ? 'false' : 'true'}/><Button type="submit" variant="outline" size="sm">{program.is_published ? 'Jadikan draft' : 'Publish'}</Button></form></div></article>{:else}<EmptyState title="Belum ada program" description="Buat program kerja pertama dari panel di samping." />{/each}</section>
  <section class="rounded-2xl border bg-card p-5"><h2 class="text-xl font-black text-blue-900">Program baru</h2><form method="POST" action="?/create" use:enhance={done()} class="mt-5 space-y-4"><div class="space-y-2"><Label for="ministry_id">Kementerian</Label><select id="ministry_id" name="ministry_id" class="h-10 w-full rounded-lg border bg-background px-3 text-sm" value={data.unitId || ''}>{#each data.units as unit (unit.id)}<option value={unit.id}>{unit.name}</option>{/each}</select></div><div class="space-y-2"><Label for="name">Nama program</Label><Input id="name" name="name" required /></div><div class="space-y-2"><Label for="slug">Slug</Label><Input id="slug" name="slug" /></div><div class="space-y-2"><Label for="short_description">Ringkasan</Label><Input id="short_description" name="short_description" /></div><div class="space-y-2"><Label for="description">Deskripsi</Label><Textarea id="description" name="description" rows={5} /></div><div class="grid gap-4 sm:grid-cols-2"><div class="space-y-2"><Label for="execution_month">Bulan</Label><Input id="execution_month" name="execution_month" /></div><div class="space-y-2"><Label for="status">Status</Label><select id="status" name="status" class="h-10 w-full rounded-lg border bg-background px-3 text-sm">{#each workProgramStatuses as status (status.value)}<option value={status.value}>{status.label}</option>{/each}</select></div></div><label class="flex items-center gap-2 text-sm"><input type="checkbox" name="is_featured" /> Featured</label><label class="flex items-center gap-2 text-sm"><input type="checkbox" name="is_published" /> Published</label><Button type="submit" class="w-full">Simpan program</Button></form></section>
</div>
