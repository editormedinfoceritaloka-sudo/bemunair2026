<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import EmptyState from '$lib/admin/EmptyState.svelte';
  import StatusBadge from '$lib/admin/StatusBadge.svelte';
  import { formatDate, deadlineTone } from '$lib/admin/format';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import * as Table from '$lib/components/ui/table';
  import { Plus, Search, ArrowUpRight } from '@lucide/svelte';

  let { data } = $props();
  let query = $state('');
  let status = $state('ALL');

  const isArticle = $derived(data.type === 'ARTICLE');
  const title = $derived(isArticle ? 'Pengajuan artikel' : 'Pengajuan media');
  const description = $derived(
    isArticle
      ? 'Kelola permintaan artikel, liputan, dan publikasi website.'
      : 'Kelola permintaan desain, feed, Reels, dan Instastory.'
  );
  const createHref = $derived(
    isArticle
      ? '/admin/content-submissions/new/article/form'
      : '/admin/content-submissions/new/content/form'
  );
  const createLabel = $derived(isArticle ? 'Buat pengajuan artikel' : 'Buat pengajuan media');
  const rowLabel = $derived(isArticle ? 'Artikel' : 'Media');
  const emptyTitle = $derived(isArticle ? 'Belum ada pengajuan artikel' : 'Belum ada pengajuan media');
  const emptyDescription = $derived(
    isArticle
      ? 'Pengajuan artikel yang masuk akan tampil dan dapat diproses dari sini.'
      : 'Pengajuan media yang masuk akan tampil dan dapat diproses dari sini.'
  );
  const filtered = $derived(
    data.submissions.filter(
      (row) =>
        `${row.title} ${row.ministry} ${row.submitter?.name || ''}`
          .toLowerCase()
          .includes(query.toLowerCase()) &&
        (status === 'ALL' || row.status === status)
    )
  );
</script>

<svelte:head><title>{title} - Admin BEM UNAIR</title></svelte:head>

<PageHeader {title} {description}>
  <Button href={createHref} class="bg-blue-500"><Plus />{createLabel}</Button>
</PageHeader>

<div class="mb-5 grid gap-3 md:grid-cols-[1fr_180px]">
  <div class="relative">
    <Search class="absolute left-2.5 top-2 size-4 text-black-200" />
    <Input bind:value={query} placeholder={`Cari judul, kementerian, atau pengaju ${rowLabel.toLowerCase()}...`} class="pl-8" />
  </div>
  <select bind:value={status} class="h-9 rounded-lg border bg-card px-3 text-sm">
    <option value="ALL">Semua status</option>
    <option>SUBMITTED</option>
    <option>PENDING_REVIEW</option>
    <option>REVISION_REQUIRED</option>
    <option>REVISION_SUBMITTED</option>
    <option>APPROVED</option>
    <option>SCHEDULED</option>
    <option>PUBLISHED</option>
    <option>REJECTED</option>
  </select>
</div>

{#if filtered.length}
  <div class="overflow-hidden rounded-xl border bg-card">
    <Table.Root>
      <Table.Header>
        <Table.Row>
          <Table.Head>{rowLabel}</Table.Head>
          <Table.Head>Status</Table.Head>
          <Table.Head>Deadline</Table.Head>
          <Table.Head>PJ</Table.Head>
          <Table.Head class="w-16"></Table.Head>
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {#each filtered as row (row.id)}
          <Table.Row>
            <Table.Cell>
              <p class="font-medium text-foreground">{row.title}</p>
              <p class="text-xs text-muted-foreground">{row.ministry} - {row.submitter?.name || '-'}</p>
            </Table.Cell>
            <Table.Cell><StatusBadge status={row.status} /></Table.Cell>
            <Table.Cell><span class={deadlineTone(row.deadline)}>{formatDate(row.deadline, true)}</span></Table.Cell>
            <Table.Cell>{row.assigned_pj?.name || 'Belum ditetapkan'}</Table.Cell>
            <Table.Cell>
              <Button href={`/admin/content-submissions/${row.id}`} variant="ghost" size="icon" aria-label="Buka detail">
                <ArrowUpRight />
              </Button>
            </Table.Cell>
          </Table.Row>
        {/each}
      </Table.Body>
    </Table.Root>
  </div>
{:else}
  <EmptyState title={emptyTitle} description={emptyDescription} />
{/if}
