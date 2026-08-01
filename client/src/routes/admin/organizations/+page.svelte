<script lang="ts">
  import { managedOrganizationUnitTypes } from '$lib/constants/cabinet';
  import { enhance } from '$app/forms';
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import EmptyState from '$lib/admin/EmptyState.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Label } from '$lib/components/ui/label';
  import type { Cabinet, OrganizationUnit } from '$lib/types';
  import { useFormSubmit } from '$lib/hooks/use-form-submit.svelte';

  let { data, form } = $props<{ data: { cabinets: Cabinet[]; units: OrganizationUnit[]; cabinetId: number | null }; form?: { error?: string } }>();
  const done = () => useFormSubmit('Kementerian tersimpan');
  const parents = $derived(data.units.filter((unit: OrganizationUnit) => unit.unit_type === 'KEMENKOAN'));
  const childrenFor = (parentId: number) => data.units.filter((unit: OrganizationUnit) => unit.parent_id === parentId);
</script>

<PageHeader title="Struktur Organisasi" description="Kelola kementerian pada struktur Kemenkoan yang sudah tersedia." />
{#if form?.error}<p class="mb-5 rounded-xl bg-red-50 p-4 text-sm text-red-700">{form.error}</p>{/if}
<div class="grid gap-6 lg:grid-cols-[1fr_360px]">
  <section class="space-y-4">{#each parents as parent (parent.id)}<article class="rounded-2xl border bg-card p-5"><div class="flex items-center justify-between"><div><p class="text-xs font-bold uppercase tracking-wider text-orange-600">Kemenkoan</p><h2 class="mt-1 text-xl font-black text-blue-900">{parent.name}</h2></div><span class="rounded-full bg-blue-50 px-3 py-1 text-xs font-bold text-blue-700">{childrenFor(parent.id).length} kementerian</span></div><div class="mt-4 grid gap-2 sm:grid-cols-2">{#each childrenFor(parent.id) as child (child.id)}<div class="rounded-xl border border-blue-100 bg-blue-50/50 p-3 text-sm font-semibold text-blue-900">{child.name}</div>{/each}</div></article>{:else}<EmptyState title="Belum ada struktur" description="Kemenkoan akan ditampilkan setelah tersedia pada seed atau API administrasi." />{/each}</section>
  <section class="rounded-2xl border bg-card p-5"><h2 class="text-xl font-black text-blue-900">Tambah kementerian</h2><form method="POST" action="?/create" use:enhance={done()} class="mt-5 space-y-4"><input type="hidden" name="cabinet_term_id" value={data.cabinetId || ""}/><div class="space-y-2"><Label for="unit_type">Tipe</Label><select id="unit_type" name="unit_type" class="h-10 w-full rounded-lg border bg-background px-3 text-sm">{#each managedOrganizationUnitTypes as type (type.value)}<option value={type.value}>{type.label}</option>{/each}</select></div><div class="space-y-2"><Label for="parent_id">Kemenkoan</Label><select id="parent_id" name="parent_id" required class="h-10 w-full rounded-lg border bg-background px-3 text-sm"><option value="">Pilih Kemenkoan</option>{#each parents as parent (parent.id)}<option value={parent.id}>{parent.name}</option>{/each}</select></div><div class="space-y-2"><Label for="code">Kode</Label><Input id="code" name="code" required /></div><div class="space-y-2"><Label for="name">Nama</Label><Input id="name" name="name" required /></div><div class="space-y-2"><Label for="slug">Slug</Label><Input id="slug" name="slug" /></div><div class="space-y-2"><Label for="short_name">Nama singkat</Label><Input id="short_name" name="short_name" /></div><div class="space-y-2"><Label for="description">Deskripsi</Label><Textarea id="description" name="description" rows={4} /></div><label class="flex items-center gap-2 text-sm"><input type="checkbox" name="is_active" checked /> Aktif</label><label class="flex items-center gap-2 text-sm"><input type="checkbox" name="is_published" /> Published</label><Button type="submit" class="w-full">Simpan kementerian</Button></form></section>
</div>
