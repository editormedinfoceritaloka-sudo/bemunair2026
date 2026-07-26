<script lang="ts">
  import type { SubmissionHistory } from '$lib/types';
  import StatusBadge from '$lib/admin/StatusBadge.svelte';
  import { CheckCircle2, UserCheck } from '@lucide/svelte';
  let { history }: { history: SubmissionHistory[] } = $props();
</script>

<section class="rounded-xl border bg-card p-6">
  <h2 class="font-serif text-xl font-semibold">Progress & timeline</h2>
  <p class="mt-1 text-sm text-muted-foreground">Riwayat status dan penetapan PJ secara kronologis.</p>
  <ol class="mt-5 space-y-0">
    {#each history as item, index}
      {@const assignment = item.event_type === 'PJ_ASSIGNED' || item.event_type === 'PJ_REASSIGNED'}
      <li class="relative grid grid-cols-[28px_1fr] gap-3 pb-6 last:pb-0">
        {#if index < history.length - 1}<span class="absolute left-[13px] top-6 h-full w-px bg-border"></span>{/if}
        <span class="relative z-10 grid size-7 place-items-center rounded-full bg-blue-50 text-blue-600">
          {#if assignment}<UserCheck class="size-4" />{:else}<CheckCircle2 class="size-4" />{/if}
        </span>
        <div class="min-w-0 pt-0.5">
          <div class="flex flex-wrap items-center gap-2">
            {#if assignment}
              <span class="rounded-full bg-blue-50 px-2.5 py-1 text-xs font-semibold text-blue-700">
                {item.event_type === 'PJ_REASSIGNED' ? 'PJ diganti' : 'PJ ditetapkan'}
              </span>
            {:else if item.to_status}
              <StatusBadge status={item.to_status} />
            {/if}
            <time class="text-xs text-muted-foreground">{new Intl.DateTimeFormat('id-ID', { dateStyle: 'medium', timeStyle: 'short' }).format(new Date(item.created_at))}</time>
          </div>
          {#if assignment}
            <p class="mt-1 text-sm">
              {#if item.from_pj}<span class="text-muted-foreground">{item.from_pj.name} → </span>{/if}
              <strong>{item.to_pj?.name || 'PJ Medinfo'}</strong>
            </p>
          {:else}
            <p class="mt-1 text-sm">{item.actor?.name || 'Sistem'}</p>
          {/if}
          {#if item.note}<p class="mt-1 whitespace-pre-wrap rounded-lg bg-muted/50 p-3 text-sm text-muted-foreground">{item.note}</p>{/if}
        </div>
      </li>
    {/each}
  </ol>
</section>
