<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import EmptyState from '$lib/admin/EmptyState.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Badge } from '$lib/components/ui/badge';
  import * as Dialog from '$lib/components/ui/dialog';
  import { Label } from '$lib/components/ui/label';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { Plus, Trash2, UserCheck, BriefcaseBusiness } from '@lucide/svelte';
  import { enhance } from '$app/forms';
  import { invalidateAll } from '$app/navigation';
  import { toast } from 'svelte-sonner';

  let { data, form } = $props();
  let open = $state(false);
  let userId = $state('');
  let roster = $derived([...data.queue].sort((a, b) => a.position - b.position));
  let available = $derived(data.users.filter((user) =>
    user.role === 'ADMIN_MEDINFO' && user.ministry === 'MEDINFO' &&
    !data.queue.some((item) => item.user_id === user.id)));
  const done = (message: string) => () => async ({ update }: any) => {
    await update(); await invalidateAll(); toast.success(message); open = false;
  };
</script>

<PageHeader title="Roster PJ Medinfo" description="Kelola petugas yang dapat dipilih untuk menangani satu task aktif.">
  <Button onclick={() => { userId = ''; open = true; }} class="bg-blue-500"><Plus />Tambah PJ</Button>
</PageHeader>

{#if form?.error}<Alert variant="destructive" class="mb-4"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}

{#if roster.length}
  <div class="grid max-w-5xl gap-4 md:grid-cols-2">
    {#each roster as item}
      <article class="rounded-xl border bg-card p-5 shadow-sm">
        <div class="flex items-start gap-4">
          <div class="grid size-11 place-items-center rounded-xl bg-blue-50 text-blue-700">
            {#if item.is_busy}<BriefcaseBusiness class="size-5" />{:else}<UserCheck class="size-5" />{/if}
          </div>
          <div class="min-w-0 flex-1">
            <div class="flex flex-wrap items-center gap-2">
              <h2 class="font-semibold">{item.user?.name || `User #${item.user_id}`}</h2>
              <Badge class={item.is_busy ? 'bg-orange-500' : 'bg-green-700'}>{item.is_busy ? 'Sedang menangani' : 'Tersedia'}</Badge>
            </div>
            <p class="mt-1 text-xs text-muted-foreground">{item.user?.email || '—'}</p>
            {#if item.is_busy}
              <a href={item.active_task_type === 'LETTER' ? `/admin/letter-submissions/${item.active_task_id}` : `/admin/content-submissions/${item.active_task_id}`} class="mt-3 block rounded-lg bg-blue-50 p-3 text-sm text-blue-800 hover:underline">
                {item.active_request_code || 'Task aktif'} · {item.active_task_title}
              </a>
            {:else}
              <p class="mt-3 rounded-lg bg-green-50 p-3 text-sm text-green-800">Siap menerima assignment baru.</p>
            {/if}
          </div>
          <form method="POST" action="?/delete" use:enhance={done('PJ dihapus dari roster')} onsubmit={(event) => { if (!confirm('Hapus PJ dari roster?')) event.preventDefault(); }}>
            <input type="hidden" name="id" value={item.id} />
            <Button type="submit" variant="ghost" size="icon" class="text-red-700" disabled={item.is_busy} aria-label="Hapus dari roster"><Trash2 /></Button>
          </form>
        </div>
      </article>
    {/each}
  </div>
{:else}
  <EmptyState title="Roster PJ masih kosong" description="Tambahkan akun ADMIN_MEDINFO dari kementerian MEDINFO." />
{/if}

<Dialog.Root bind:open>
  <Dialog.Content>
    <Dialog.Header><Dialog.Title>Tambah PJ Medinfo</Dialog.Title><Dialog.Description>Hanya akun ADMIN_MEDINFO kementerian MEDINFO yang dapat dimasukkan.</Dialog.Description></Dialog.Header>
    {#if available.length}
      <form method="POST" action="?/create" use:enhance={done('PJ ditambahkan')} class="space-y-4">
        <div class="space-y-2">
          <Label for="user_id">Pengguna</Label>
          <select id="user_id" name="user_id" bind:value={userId} required class="h-10 w-full rounded-lg border bg-card px-3 text-sm">
            <option value="" disabled>Pilih pengguna</option>
            {#each available as user}<option value={user.id}>{user.name} · {user.email}</option>{/each}
          </select>
          <p class="text-xs text-muted-foreground">Akun baru yang dibuat setelah proses seeding akan otomatis muncul sebagai pilihan selama memenuhi role dan kementerian.</p>
        </div>
        <Dialog.Footer><Button type="button" variant="outline" onclick={() => open = false}>Batal</Button><Button type="submit" class="bg-blue-500">Tambahkan</Button></Dialog.Footer>
      </form>
    {:else}
      <div class="space-y-4">
        <Alert><UserCheck class="size-4" /><AlertDescription>Semua akun ADMIN_MEDINFO dari kementerian MEDINFO saat ini sudah masuk roster. Anda tetap dapat menambahkan PJ baru setelah membuat akun yang memenuhi kriteria tersebut.</AlertDescription></Alert>
        <Dialog.Footer><Button type="button" variant="outline" onclick={() => open = false}>Tutup</Button><Button href="/admin/users" class="bg-blue-500">Kelola Pengguna</Button></Dialog.Footer>
      </div>
    {/if}
  </Dialog.Content>
</Dialog.Root>
