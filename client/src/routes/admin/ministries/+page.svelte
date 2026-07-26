<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import { Badge } from '$lib/components/ui/badge';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import * as Table from '$lib/components/ui/table';
  import * as Dialog from '$lib/components/ui/dialog';
  import { Plus, Power } from '@lucide/svelte';
  import { enhance } from '$app/forms';
  import { invalidateAll } from '$app/navigation';
  import { toast } from 'svelte-sonner';
  let { data, form } = $props();
  let open = $state(false);
  const done = () => () => async ({ update }: any) => {
    await update(); await invalidateAll(); open = false; toast.success('Data kementerian diperbarui');
  };
</script>

<PageHeader title="Master kementerian" description="Standarkan identitas kementerian untuk akun dan snapshot setiap pengajuan.">
  <Button onclick={() => open = true} class="bg-blue-500"><Plus />Tambah kementerian</Button>
</PageHeader>
{#if form?.error}<Alert variant="destructive" class="mb-4"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}

<div class="overflow-hidden rounded-xl border bg-card">
  <Table.Root>
    <Table.Header><Table.Row><Table.Head>Kode</Table.Head><Table.Head>Nama</Table.Head><Table.Head>Status</Table.Head><Table.Head class="text-right">Aksi</Table.Head></Table.Row></Table.Header>
    <Table.Body>
      {#each data.ministries as item}
        <Table.Row>
          <Table.Cell class="font-mono text-sm">{item.code}</Table.Cell>
          <Table.Cell class="font-medium">{item.name}</Table.Cell>
          <Table.Cell><Badge variant="outline" class={item.is_active ? 'border-green-200 bg-green-50 text-green-700' : 'bg-muted text-muted-foreground'}>{item.is_active ? 'Aktif' : 'Nonaktif'}</Badge></Table.Cell>
          <Table.Cell>
            <form method="POST" action="?/toggle" use:enhance={done()} class="flex justify-end">
              <input type="hidden" name="id" value={item.id} /><input type="hidden" name="active" value={String(item.is_active)} />
              <Button type="submit" variant="ghost" size="sm"><Power />{item.is_active ? 'Nonaktifkan' : 'Aktifkan'}</Button>
            </form>
          </Table.Cell>
        </Table.Row>
      {/each}
    </Table.Body>
  </Table.Root>
</div>

<Dialog.Root bind:open>
  <Dialog.Content>
    <Dialog.Header><Dialog.Title>Tambah kementerian</Dialog.Title><Dialog.Description>Kode digunakan sebagai identitas stabil pada integrasi API.</Dialog.Description></Dialog.Header>
    <form method="POST" action="?/create" use:enhance={done()} class="space-y-4">
      <div class="space-y-2"><Label for="code">Kode</Label><Input id="code" name="code" placeholder="CONTOH_KODE" required /></div>
      <div class="space-y-2"><Label for="name">Nama kementerian</Label><Input id="name" name="name" required /></div>
      <Dialog.Footer><Button type="button" variant="outline" onclick={() => open = false}>Batal</Button><Button type="submit" class="bg-blue-500">Simpan</Button></Dialog.Footer>
    </form>
  </Dialog.Content>
</Dialog.Root>
