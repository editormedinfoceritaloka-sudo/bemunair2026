<script lang="ts">
  import { enhance } from '$app/forms';
  import type { SubmitFunction } from '@sveltejs/kit';
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import EmptyState from '$lib/admin/EmptyState.svelte';
  import { uploadImageFile, type UploadedImage } from '$lib/admin/upload-image';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { Badge } from '$lib/components/ui/badge';
  import { Button } from '$lib/components/ui/button';
  import * as Dialog from '$lib/components/ui/dialog';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import { Textarea } from '$lib/components/ui/textarea';
  import type { OrganizationMember, OrganizationUnit, OrganizationUnitType } from '$lib/types';
  import {
    Building2,
    CheckCircle2,
    Edit3,
    ImagePlus,
    LoaderCircle,
    Plus,
    Search,
    ShieldCheck,
    Trash2,
    UploadCloud,
    UserRoundCog,
    Users
  } from '@lucide/svelte';
  import { toast } from 'svelte-sonner';

  let { data, form } = $props();

  let unitOpen = $state(false);
  let editing = $state<OrganizationUnit | null>(null);
  let unitType = $state<OrganizationUnitType>('KEMENTERIAN');
  let logoInput = $state<HTMLInputElement>();
  let coverInput = $state<HTMLInputElement>();
  let logoUpload = $state<UploadedImage | null>(null);
  let coverUpload = $state<UploadedImage | null>(null);
  let logoPreview = $state('');
  let coverPreview = $state('');
  let logoMediaID = $state('');
  let coverMediaID = $state('');
  let uploading = $state<'logo' | 'cover' | 'photo' | null>(null);

  let deleteOpen = $state(false);
  let deleteTarget = $state<OrganizationUnit | null>(null);

  let membersOpen = $state(false);
  let memberUnit = $state<OrganizationUnit | null>(null);
  let memberEditing = $state<OrganizationMember | null>(null);
  let photoInput = $state<HTMLInputElement>();
  let photoUpload = $state<UploadedImage | null>(null);
  let photoPreview = $state('');
  let photoMediaID = $state('');
  let submitting = $state(false);

  const normalizedQuery = $derived(data.filters.q.trim().toLowerCase());
  const units = $derived(data.units as OrganizationUnit[]);
  const parents = $derived(units.filter((unit) => unit.unit_type !== 'KEMENTERIAN'));
  const ministries = $derived(units.filter((unit) => unit.unit_type === 'KEMENTERIAN'));

  const childrenFor = (parentID: number) =>
    ministries.filter((unit) => unit.parent_id === parentID);

  const unitTypeLabel = (type: OrganizationUnitType) => {
    if (type === 'MENKO' || type === 'KEMENKOAN') return 'Kemenkoan';
    if (type === 'BPII') return 'BPII';
    return 'Kementerian';
  };

  const matches = (unit: OrganizationUnit) => {
    const haystack = `${unit.code} ${unit.name} ${unit.short_name || ''} ${unit.slug}`.toLowerCase();
    const queryMatches = !normalizedQuery || haystack.includes(normalizedQuery);
    const typeMatches = data.filters.type === 'ALL' || unit.unit_type === data.filters.type;
    const statusMatches =
      data.filters.status === 'ALL'
      || (data.filters.status === 'ACTIVE' && unit.is_active)
      || (data.filters.status === 'INACTIVE' && !unit.is_active)
      || (data.filters.status === 'PUBLISHED' && unit.is_published)
      || (data.filters.status === 'DRAFT' && !unit.is_published);
    return queryMatches && typeMatches && statusMatches;
  };

  const visibleParents = $derived(
    parents.filter((parent) => matches(parent) || childrenFor(parent.id).some(matches))
  );
  const visibleOrphans = $derived(
    ministries.filter((unit) => !parents.some((parent) => parent.id === unit.parent_id) && matches(unit))
  );

  const submit = (onSuccess?: () => void): SubmitFunction => () => {
    submitting = true;
    return async ({ result, update }) => {
      await update({ reset: false });
      submitting = false;
      if (result.type === 'success') {
        onSuccess?.();
        toast.success('Perubahan berhasil disimpan');
      }
    };
  };

  function openUnit(unit: OrganizationUnit | null, type: OrganizationUnitType = 'KEMENTERIAN') {
    editing = unit;
    unitType = unit?.unit_type ?? type;
    logoUpload = null;
    coverUpload = null;
    logoPreview = unit?.logo?.url ?? '';
    coverPreview = unit?.cover?.url ?? '';
    logoMediaID = unit?.logo?.id ? String(unit.logo.id) : '';
    coverMediaID = unit?.cover?.id ? String(unit.cover.id) : '';
    unitOpen = true;
  }

  function requestDelete(unit: OrganizationUnit) {
    deleteTarget = unit;
    deleteOpen = true;
  }

  function openMembers(unit: OrganizationUnit) {
    memberUnit = unit;
    memberEditing = null;
    resetMemberMedia(null);
    membersOpen = true;
  }

  function editMember(member: OrganizationMember | null) {
    memberEditing = member;
    resetMemberMedia(member);
  }

  function resetMemberMedia(member: OrganizationMember | null) {
    photoUpload = null;
    photoPreview = member?.photo?.url ?? '';
    photoMediaID = member?.photo?.id ? String(member.photo.id) : '';
  }

  async function uploadAsset(event: Event, kind: 'logo' | 'cover' | 'photo') {
    const input = event.currentTarget as HTMLInputElement;
    const file = input.files?.[0];
    if (!file) return;
    uploading = kind;
    try {
      const value = await uploadImageFile(file, kind === 'cover' ? 'cover' : 'profile');
      if (kind === 'logo') {
        logoUpload = value;
        logoPreview = value.thumbnail_url || value.url;
        logoMediaID = '';
      } else if (kind === 'cover') {
        coverUpload = value;
        coverPreview = value.thumbnail_url || value.url;
        coverMediaID = '';
      } else {
        photoUpload = value;
        photoPreview = value.thumbnail_url || value.url;
        photoMediaID = '';
      }
      toast.success('Gambar berhasil diunggah');
    } catch (error) {
      toast.error(error instanceof Error ? error.message : 'Gagal mengunggah gambar');
    } finally {
      uploading = null;
      input.value = '';
    }
  }

  function removeMedia(kind: 'logo' | 'cover' | 'photo') {
    if (kind === 'logo') {
      logoUpload = null;
      logoPreview = '';
      logoMediaID = '';
    } else if (kind === 'cover') {
      coverUpload = null;
      coverPreview = '';
      coverMediaID = '';
    } else {
      photoUpload = null;
      photoPreview = '';
      photoMediaID = '';
    }
  }
</script>

<svelte:head><title>Kementerian & Kemenkoan - Admin BEM UNAIR</title></svelte:head>

<PageHeader
  title="Kementerian & Kemenkoan"
  description="Kelola MENKO, BPII, kementerian, identitas publik, media, dan tim setelah proses seeding."
>
  <Button type="button" onclick={() => openUnit(null)} class="bg-blue-500">
    <Plus /> Tambah unit
  </Button>
</PageHeader>

{#if form?.error}
  <Alert variant="destructive" class="mb-5">
    <AlertDescription>{form.error}</AlertDescription>
  </Alert>
{/if}

<div class="mb-6 grid gap-4 rounded-xl border bg-card p-4 lg:grid-cols-[220px_1fr_170px_170px_auto]">
  <form method="GET" class="contents">
    <div class="space-y-1.5">
      <Label for="cabinet">Kabinet</Label>
      <select id="cabinet" name="cabinet" class="h-9 w-full rounded-lg border bg-background px-3 text-sm" value={data.cabinetId ?? ''}>
        {#each data.cabinets as cabinet (cabinet.id)}
          <option value={cabinet.id}>{cabinet.name}</option>
        {/each}
      </select>
    </div>
    <div class="space-y-1.5">
      <Label for="q">Cari unit</Label>
      <div class="relative">
        <Search class="pointer-events-none absolute left-2.5 top-2.5 size-4 text-muted-foreground" aria-hidden="true" />
        <Input id="q" name="q" value={data.filters.q} placeholder="Kode, nama, atau slug…" class="pl-8" />
      </div>
    </div>
    <div class="space-y-1.5">
      <Label for="type">Tipe</Label>
      <select id="type" name="type" class="h-9 w-full rounded-lg border bg-background px-3 text-sm" value={data.filters.type}>
        <option value="ALL">Semua tipe</option>
        <option value="MENKO">Kemenkoan</option>
        <option value="BPII">BPII</option>
        <option value="KEMENTERIAN">Kementerian</option>
      </select>
    </div>
    <div class="space-y-1.5">
      <Label for="status">Status</Label>
      <select id="status" name="status" class="h-9 w-full rounded-lg border bg-background px-3 text-sm" value={data.filters.status}>
        <option value="ALL">Semua status</option>
        <option value="ACTIVE">Aktif</option>
        <option value="INACTIVE">Nonaktif</option>
        <option value="PUBLISHED">Published</option>
        <option value="DRAFT">Draft</option>
      </select>
    </div>
    <div class="flex items-end">
      <Button type="submit" variant="outline" class="w-full">Terapkan</Button>
    </div>
  </form>
</div>

<div class="mb-5 grid gap-3 sm:grid-cols-3">
  <div class="rounded-xl border bg-card p-4">
    <p class="text-sm text-muted-foreground">Unit induk</p>
    <p class="mt-1 text-2xl font-bold tabular-nums">{parents.length}</p>
  </div>
  <div class="rounded-xl border bg-card p-4">
    <p class="text-sm text-muted-foreground">Kementerian</p>
    <p class="mt-1 text-2xl font-bold tabular-nums">{ministries.length}</p>
  </div>
  <div class="rounded-xl border bg-card p-4">
    <p class="text-sm text-muted-foreground">Unit published</p>
    <p class="mt-1 text-2xl font-bold tabular-nums">{units.filter((unit) => unit.is_published).length}</p>
  </div>
</div>

{#if visibleParents.length || visibleOrphans.length}
  <div class="space-y-4">
    {#each visibleParents as parent (parent.id)}
      {@const children = childrenFor(parent.id).filter(matches)}
      <section class="overflow-hidden rounded-2xl border bg-card shadow-sm">
        <header class="flex flex-col gap-4 border-b bg-blue-50/60 p-5 sm:flex-row sm:items-center">
          <div class="flex min-w-0 flex-1 items-center gap-3">
            <div class="grid size-12 shrink-0 place-items-center overflow-hidden rounded-xl border border-blue-100 bg-white">
              {#if parent.logo?.url}
                <img src={parent.logo.url} alt={parent.logo.alt_text || `Logo ${parent.name}`} width="48" height="48" class="size-full object-contain p-1" />
              {:else}
                <Building2 class="size-5 text-blue-700" aria-hidden="true" />
              {/if}
            </div>
            <div class="min-w-0">
              <div class="flex flex-wrap items-center gap-2">
                <p class="text-xs font-bold uppercase tracking-wider text-orange-600">{unitTypeLabel(parent.unit_type)} · {parent.code}</p>
                <Badge variant="outline" class={parent.is_active ? 'border-green-200 bg-green-50 text-green-700' : 'bg-muted text-muted-foreground'}>
                  {parent.is_active ? 'Aktif' : 'Nonaktif'}
                </Badge>
                <Badge variant="outline">{parent.is_published ? 'Published' : 'Draft'}</Badge>
              </div>
              <h2 class="mt-1 truncate text-lg font-bold text-blue-900">{parent.name}</h2>
              <p class="mt-1 text-sm text-muted-foreground">{childrenFor(parent.id).length} kementerian · urutan {parent.display_order}</p>
            </div>
          </div>
          <div class="flex flex-wrap gap-2">
            <Button type="button" size="sm" variant="outline" onclick={() => openMembers(parent)}>
              <Users /> Tim
            </Button>
            <Button type="button" size="sm" variant="outline" onclick={() => openUnit(parent)}>
              <Edit3 /> Kelola
            </Button>
            <Button type="button" size="sm" variant="ghost" class="text-red-700" onclick={() => requestDelete(parent)}>
              <Trash2 /> Hapus
            </Button>
          </div>
        </header>

        <div class="grid gap-3 p-4 lg:grid-cols-2">
          {#each children as child (child.id)}
            <article class="flex min-w-0 flex-col gap-4 rounded-xl border border-border/80 bg-background p-4 sm:flex-row sm:items-center">
              <div class="flex min-w-0 flex-1 items-center gap-3">
                <div class="grid size-10 shrink-0 place-items-center overflow-hidden rounded-lg border bg-white">
                  {#if child.logo?.url}
                    <img src={child.logo.url} alt={child.logo.alt_text || `Logo ${child.name}`} width="40" height="40" class="size-full object-contain p-1" loading="lazy" />
                  {:else}
                    <ShieldCheck class="size-4 text-blue-700" aria-hidden="true" />
                  {/if}
                </div>
                <div class="min-w-0">
                  <p class="truncate font-semibold text-foreground">{child.name}</p>
                  <p class="mt-0.5 truncate text-xs text-muted-foreground">{child.code} · {child.short_name || child.slug}</p>
                  <div class="mt-2 flex flex-wrap gap-1.5">
                    <Badge variant="outline" class={child.is_active ? 'border-green-200 bg-green-50 text-green-700' : 'bg-muted text-muted-foreground'}>
                      {child.is_active ? 'Aktif' : 'Nonaktif'}
                    </Badge>
                    <Badge variant="outline">{child.is_published ? 'Published' : 'Draft'}</Badge>
                    <Badge variant="secondary">{child.members?.length || 0} anggota</Badge>
                  </div>
                </div>
              </div>
              <div class="flex shrink-0 gap-1 self-end sm:self-auto">
                <Button type="button" variant="ghost" size="icon" onclick={() => openMembers(child)} aria-label={`Kelola tim ${child.name}`}>
                  <UserRoundCog />
                </Button>
                <Button type="button" variant="ghost" size="icon" onclick={() => openUnit(child)} aria-label={`Edit ${child.name}`}>
                  <Edit3 />
                </Button>
                <Button type="button" variant="ghost" size="icon" class="text-red-700" onclick={() => requestDelete(child)} aria-label={`Hapus ${child.name}`}>
                  <Trash2 />
                </Button>
              </div>
            </article>
          {:else}
            <div class="rounded-xl border border-dashed p-5 text-sm text-muted-foreground lg:col-span-2">
              Tidak ada kementerian yang cocok dengan filter pada unit induk ini.
            </div>
          {/each}
        </div>
      </section>
    {/each}

    {#if visibleOrphans.length}
      <section class="rounded-2xl border border-orange-200 bg-orange-50/40 p-5">
        <h2 class="font-semibold text-orange-900">Kementerian tanpa parent valid</h2>
        <p class="mt-1 text-sm text-orange-800">Perbaiki parent agar struktur publik dapat ditampilkan dengan benar.</p>
        <div class="mt-4 grid gap-3 lg:grid-cols-2">
          {#each visibleOrphans as unit (unit.id)}
            <button type="button" class="flex items-center justify-between rounded-xl border bg-card p-4 text-left hover:border-orange-300" onclick={() => openUnit(unit)}>
              <span><strong>{unit.name}</strong><small class="mt-1 block text-muted-foreground">{unit.code}</small></span>
              <Edit3 class="size-4" />
            </button>
          {/each}
        </div>
      </section>
    {/if}
  </div>
{:else}
  <EmptyState title="Unit organisasi tidak ditemukan" description="Ubah filter atau tambahkan Kemenkoan dan Kementerian baru." />
{/if}

<Dialog.Root bind:open={unitOpen}>
  <Dialog.Content class="max-h-[92vh] overflow-y-auto sm:max-w-3xl">
    <Dialog.Header>
      <Dialog.Title>{editing ? `Edit ${editing.name}` : 'Tambah unit organisasi'}</Dialog.Title>
      <Dialog.Description>Data ini dipakai oleh halaman publik, akun admin, dan modul program kerja.</Dialog.Description>
    </Dialog.Header>

    {#key editing?.id ?? 'new'}
      <form method="POST" action="?/saveUnit" use:enhance={submit(() => unitOpen = false)} class="space-y-6">
        <input type="hidden" name="id" value={editing?.id ?? ''} />
        <input type="hidden" name="cabinet_term_id" value={data.cabinetId ?? ''} />
        <input type="hidden" name="logo_media_id" value={logoMediaID} />
        <input type="hidden" name="cover_media_id" value={coverMediaID} />
        <input type="hidden" name="logo_upload" value={logoUpload ? JSON.stringify(logoUpload) : ''} />
        <input type="hidden" name="cover_upload" value={coverUpload ? JSON.stringify(coverUpload) : ''} />

        <section class="space-y-4">
          <div>
            <h3 class="font-semibold">Identitas dan struktur</h3>
            <p class="text-sm text-muted-foreground">Kode harus stabil karena digunakan oleh integrasi dan akun.</p>
          </div>
          <div class="grid gap-4 sm:grid-cols-2">
            <div class="space-y-2">
              <Label for="unit_type">Tipe unit</Label>
              <select id="unit_type" name="unit_type" bind:value={unitType} class="h-9 w-full rounded-lg border bg-background px-3 text-sm">
                <option value="MENKO">Kemenkoan</option>
                <option value="BPII">BPII</option>
                <option value="KEMENTERIAN">Kementerian</option>
              </select>
            </div>
            {#if unitType === 'KEMENTERIAN'}
              <div class="space-y-2">
                <Label for="parent_id">Unit induk</Label>
                <select id="parent_id" name="parent_id" required class="h-9 w-full rounded-lg border bg-background px-3 text-sm" value={editing?.parent_id ?? ''}>
                  <option value="">Pilih unit induk</option>
                  {#each parents.filter((parent) => parent.id !== editing?.id) as parent (parent.id)}
                    <option value={parent.id}>{parent.name}</option>
                  {/each}
                </select>
              </div>
            {/if}
            <div class="space-y-2">
              <Label for="code">Kode</Label>
              <Input id="code" name="code" value={editing?.code ?? ''} maxlength={50} autocomplete="off" required />
            </div>
            <div class="space-y-2">
              <Label for="name">Nama lengkap</Label>
              <Input id="name" name="name" value={editing?.name ?? ''} maxlength={120} autocomplete="off" required />
            </div>
            <div class="space-y-2">
              <Label for="short_name">Nama singkat</Label>
              <Input id="short_name" name="short_name" value={editing?.short_name ?? ''} maxlength={80} autocomplete="off" />
            </div>
            <div class="space-y-2">
              <Label for="slug">Slug halaman publik</Label>
              <Input id="slug" name="slug" value={editing?.slug ?? ''} placeholder="Otomatis dari nama jika kosong" autocomplete="off" />
            </div>
            <div class="space-y-2">
              <Label for="display_order">Urutan tampilan</Label>
              <Input id="display_order" name="display_order" type="number" min="0" value={editing?.display_order ?? 0} inputmode="numeric" />
            </div>
          </div>
        </section>

        <section class="space-y-4 border-t pt-5">
          <div>
            <h3 class="font-semibold">Konten halaman publik</h3>
            <p class="text-sm text-muted-foreground">Lengkapi narasi yang tampil pada profil organisasi.</p>
          </div>
          <div class="space-y-2">
            <Label for="description">Deskripsi</Label>
            <Textarea id="description" name="description" rows={4} value={editing?.description ?? ''} />
          </div>
          <div class="grid gap-4 sm:grid-cols-2">
            <div class="space-y-2">
              <Label for="vision">Visi</Label>
              <Textarea id="vision" name="vision" rows={5} value={editing?.vision ?? ''} />
            </div>
            <div class="space-y-2">
              <Label for="mission">Misi</Label>
              <Textarea id="mission" name="mission" rows={5} value={editing?.mission ?? ''} />
            </div>
          </div>
        </section>

        <section class="space-y-4 border-t pt-5">
          <div>
            <h3 class="font-semibold">Logo dan cover</h3>
            <p class="text-sm text-muted-foreground">Logo disarankan rasio 1:1 dan cover rasio 16:9, maksimum 10 MB.</p>
          </div>
          <div class="grid gap-4 sm:grid-cols-2">
            <div class="space-y-2">
              <Label for="logo-upload">Logo organisasi</Label>
              <input bind:this={logoInput} id="logo-upload" type="file" accept="image/jpeg,image/png,image/webp,image/gif,image/avif" class="sr-only" onchange={(event) => uploadAsset(event, 'logo')} />
              {#if logoPreview}
                <div class="flex items-center gap-3 rounded-xl border p-3">
                  <img src={logoPreview} alt="Preview logo organisasi" width="64" height="64" class="size-16 rounded-lg object-contain" />
                  <div class="flex flex-1 justify-end gap-1">
                    <Button type="button" variant="outline" size="sm" onclick={() => logoInput?.click()} disabled={uploading !== null}><ImagePlus /> Ganti</Button>
                    <Button type="button" variant="ghost" size="icon" class="text-red-700" onclick={() => removeMedia('logo')} aria-label="Hapus logo"><Trash2 /></Button>
                  </div>
                </div>
              {:else}
                <button type="button" class="flex h-28 w-full flex-col items-center justify-center gap-2 rounded-xl border border-dashed bg-muted/30 hover:border-blue-300 hover:bg-blue-50" onclick={() => logoInput?.click()} disabled={uploading !== null}>
                  {#if uploading === 'logo'}<LoaderCircle class="size-5 animate-spin" /> Mengunggah…
                  {:else}<UploadCloud class="size-5" /> Unggah logo{/if}
                </button>
              {/if}
            </div>
            <div class="space-y-2">
              <Label for="cover-upload">Cover organisasi</Label>
              <input bind:this={coverInput} id="cover-upload" type="file" accept="image/jpeg,image/png,image/webp,image/gif,image/avif" class="sr-only" onchange={(event) => uploadAsset(event, 'cover')} />
              {#if coverPreview}
                <div class="space-y-2 rounded-xl border p-3">
                  <img src={coverPreview} alt="Preview cover organisasi" width="320" height="180" class="aspect-video w-full rounded-lg object-cover" />
                  <div class="flex justify-end gap-1">
                    <Button type="button" variant="outline" size="sm" onclick={() => coverInput?.click()} disabled={uploading !== null}><ImagePlus /> Ganti</Button>
                    <Button type="button" variant="ghost" size="icon" class="text-red-700" onclick={() => removeMedia('cover')} aria-label="Hapus cover"><Trash2 /></Button>
                  </div>
                </div>
              {:else}
                <button type="button" class="flex h-28 w-full flex-col items-center justify-center gap-2 rounded-xl border border-dashed bg-muted/30 hover:border-blue-300 hover:bg-blue-50" onclick={() => coverInput?.click()} disabled={uploading !== null}>
                  {#if uploading === 'cover'}<LoaderCircle class="size-5 animate-spin" /> Mengunggah…
                  {:else}<UploadCloud class="size-5" /> Unggah cover{/if}
                </button>
              {/if}
            </div>
          </div>
        </section>

        <section class="grid gap-3 border-t pt-5 sm:grid-cols-2">
          <label class="flex cursor-pointer items-start gap-3 rounded-xl border p-4">
            <input type="checkbox" name="is_active" checked={editing?.is_active ?? true} class="mt-1 size-4" />
            <span><strong class="block text-sm">Unit aktif</strong><small class="text-muted-foreground">Dapat dipakai oleh akun dan modul internal.</small></span>
          </label>
          <label class="flex cursor-pointer items-start gap-3 rounded-xl border p-4">
            <input type="checkbox" name="is_published" checked={editing?.is_published ?? false} class="mt-1 size-4" />
            <span><strong class="block text-sm">Tampilkan ke publik</strong><small class="text-muted-foreground">Profil tampil pada website utama.</small></span>
          </label>
        </section>

        <Dialog.Footer>
          <Button type="button" variant="outline" onclick={() => unitOpen = false}>Batal</Button>
          <Button type="submit" class="bg-blue-500" disabled={submitting || uploading !== null}>
            {#if submitting}<LoaderCircle class="animate-spin" /> Menyimpan…
            {:else}<CheckCircle2 /> Simpan unit{/if}
          </Button>
        </Dialog.Footer>
      </form>
    {/key}
  </Dialog.Content>
</Dialog.Root>

<Dialog.Root bind:open={membersOpen}>
  <Dialog.Content class="max-h-[92vh] overflow-y-auto sm:max-w-4xl">
    <Dialog.Header>
      <Dialog.Title>Tim {memberUnit?.name}</Dialog.Title>
      <Dialog.Description>Kelola pimpinan dan anggota yang tampil pada profil organisasi.</Dialog.Description>
    </Dialog.Header>

    {#if memberUnit}
      <div class="grid gap-6 lg:grid-cols-[1fr_360px]">
        <section class="space-y-3">
          <div class="flex items-center justify-between">
            <h3 class="font-semibold">Daftar anggota</h3>
            <Button type="button" size="sm" variant="outline" onclick={() => editMember(null)}><Plus /> Tambah</Button>
          </div>
          {#each memberUnit.members || [] as member (member.id)}
            <button type="button" class="flex w-full items-center gap-3 rounded-xl border p-3 text-left hover:border-blue-300 hover:bg-blue-50/40" onclick={() => editMember(member)}>
              <div class="grid size-11 shrink-0 place-items-center overflow-hidden rounded-full border bg-muted">
                {#if member.photo?.url}
                  <img src={member.photo.url} alt={member.photo.alt_text || `Foto ${member.name}`} width="44" height="44" class="size-full object-cover" loading="lazy" />
                {:else}
                  <Users class="size-4 text-muted-foreground" aria-hidden="true" />
                {/if}
              </div>
              <span class="min-w-0 flex-1">
                <strong class="block truncate text-sm">{member.name}</strong>
                <small class="block truncate text-muted-foreground">{member.position}</small>
              </span>
              <Badge variant="outline" class={member.is_active ? 'border-green-200 text-green-700' : ''}>{member.is_active ? 'Aktif' : 'Nonaktif'}</Badge>
            </button>
          {:else}
            <div class="rounded-xl border border-dashed p-5 text-center text-sm text-muted-foreground">Belum ada anggota.</div>
          {/each}
        </section>

        <section class="rounded-xl border bg-muted/20 p-4">
          {#key memberEditing?.id ?? 'new-member'}
            <form method="POST" action="?/saveMember" use:enhance={submit(() => { memberEditing = null; membersOpen = false; })} class="space-y-4">
              <input type="hidden" name="member_id" value={memberEditing?.id ?? ''} />
              <input type="hidden" name="ministry_id" value={memberUnit.id} />
              <input type="hidden" name="photo_media_id" value={photoMediaID} />
              <input type="hidden" name="photo_upload" value={photoUpload ? JSON.stringify(photoUpload) : ''} />
              <div>
                <h3 class="font-semibold">{memberEditing ? 'Edit anggota' : 'Anggota baru'}</h3>
                <p class="text-xs text-muted-foreground">Nonaktifkan anggota untuk mempertahankan histori profil.</p>
              </div>
              <div class="space-y-2">
                <Label for="member_name">Nama</Label>
                <Input id="member_name" name="member_name" value={memberEditing?.name ?? ''} required />
              </div>
              <div class="space-y-2">
                <Label for="position">Jabatan</Label>
                <Input id="position" name="position" value={memberEditing?.position ?? ''} placeholder="Menteri atau Direktur Jenderal" required />
              </div>
              <div class="space-y-2">
                <Label for="position_type">Tipe posisi</Label>
                <select id="position_type" name="position_type" class="h-9 w-full rounded-lg border bg-background px-3 text-sm" value={memberEditing?.position_type ?? (memberUnit.unit_type === 'KEMENTERIAN' ? 'MENTERI' : 'MENKO')}>
                  <option value="MENKO">Menteri Koordinator</option>
                  <option value="MENTERI">Menteri</option>
                  <option value="DIRJEN">Direktur Jenderal</option>
                  <option value="KABIRO">Kepala Biro</option>
                  <option value="KEPALA">Kepala</option>
                  <option value="DEPUTI">Deputi</option>
                  <option value="MINISTER">Menteri (legacy)</option>
                  <option value="DIRECTOR_GENERAL">Direktur Jenderal (legacy)</option>
                </select>
              </div>
              <div class="space-y-2">
                <Label for="biography">Biografi</Label>
                <Textarea id="biography" name="biography" rows={3} value={memberEditing?.biography ?? ''} />
              </div>
              <div class="space-y-2">
                <Label for="quote">Kutipan</Label>
                <Textarea id="quote" name="quote" rows={2} value={memberEditing?.quote ?? ''} />
              </div>
              <div class="space-y-2">
                <Label for="member_display_order">Urutan</Label>
                <Input id="member_display_order" name="member_display_order" type="number" min="0" inputmode="numeric" value={memberEditing?.display_order ?? 0} />
              </div>

              <div class="space-y-2">
                <Label for="photo-upload">Foto profil</Label>
                <input bind:this={photoInput} id="photo-upload" type="file" accept="image/jpeg,image/png,image/webp,image/gif,image/avif" class="sr-only" onchange={(event) => uploadAsset(event, 'photo')} />
                <div class="flex items-center gap-3 rounded-xl border bg-background p-3">
                  {#if photoPreview}
                    <img src={photoPreview} alt="Preview foto anggota" width="56" height="56" class="size-14 rounded-full object-cover" />
                  {:else}
                    <div class="grid size-14 place-items-center rounded-full bg-muted"><Users class="size-5" /></div>
                  {/if}
                  <div class="ml-auto flex gap-1">
                    <Button type="button" size="sm" variant="outline" onclick={() => photoInput?.click()} disabled={uploading !== null}>
                      {#if uploading === 'photo'}<LoaderCircle class="animate-spin" />{:else}<ImagePlus />{/if} Pilih
                    </Button>
                    {#if photoPreview}<Button type="button" size="icon" variant="ghost" class="text-red-700" onclick={() => removeMedia('photo')} aria-label="Hapus foto"><Trash2 /></Button>{/if}
                  </div>
                </div>
              </div>

              <div class="grid gap-2 sm:grid-cols-2">
                <label class="flex items-center gap-2 text-sm"><input type="checkbox" name="is_leader" checked={memberEditing?.is_leader ?? false} /> Pimpinan</label>
                <label class="flex items-center gap-2 text-sm"><input type="checkbox" name="member_is_active" checked={memberEditing?.is_active ?? true} /> Aktif</label>
              </div>
              <Button type="submit" class="w-full bg-blue-500" disabled={submitting || uploading !== null}>
                {#if submitting}<LoaderCircle class="animate-spin" /> Menyimpan…{:else}Simpan anggota{/if}
              </Button>
            </form>
          {/key}
        </section>
      </div>
    {/if}
  </Dialog.Content>
</Dialog.Root>

<Dialog.Root bind:open={deleteOpen}>
  <Dialog.Content class="sm:max-w-md">
    <Dialog.Header>
      <Dialog.Title>Hapus {deleteTarget?.name}?</Dialog.Title>
      <Dialog.Description>
        Unit hanya akan dihapus bila belum memiliki kementerian anak, pengguna, anggota, program kerja, atau pengajuan. Jika masih digunakan, nonaktifkan unit melalui menu edit.
      </Dialog.Description>
    </Dialog.Header>
    {#if deleteTarget}
      <form method="POST" action="?/deleteUnit" use:enhance={submit(() => { deleteOpen = false; deleteTarget = null; })}>
        <input type="hidden" name="id" value={deleteTarget.id} />
        <Dialog.Footer>
          <Button type="button" variant="outline" onclick={() => deleteOpen = false}>Batal</Button>
          <Button type="submit" variant="destructive" disabled={submitting}>
            {#if submitting}<LoaderCircle class="animate-spin" /> Menghapus…{:else}<Trash2 /> Hapus unit{/if}
          </Button>
        </Dialog.Footer>
      </form>
    {/if}
  </Dialog.Content>
</Dialog.Root>
