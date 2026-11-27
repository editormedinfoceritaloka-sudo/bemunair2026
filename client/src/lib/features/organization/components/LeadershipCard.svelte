<script lang="ts">
  import { memberCardRole } from '../member-card-role';
  import { resolveOrganizationPortrait } from '../portraits';
  import type { OrganizationMember } from '$lib/types';
  let { member }: { member: OrganizationMember } = $props();
  const portrait = $derived(resolveOrganizationPortrait(member));

</script>

<article class="overflow-hidden rounded-2xl border border-blue-100 bg-white shadow-[0_14px_40px_rgba(16,56,100,0.1)]">
  <div class="aspect-[4/3] bg-gradient-to-br from-blue-50 to-blue-200">
    {#if portrait}<img src={portrait} alt={member.name} class="size-full object-cover" />{:else}<div class="grid size-full place-items-center text-6xl font-black text-blue-500">{member.name.slice(0, 1)}</div>{/if}
  </div>
  <div class="p-5">
    <p class="text-xs font-bold uppercase tracking-[0.18em] text-orange-600">{memberCardRole(member)}</p>
    <h3 class="mt-2 text-xl font-black text-blue-900">{member.name}</h3>
    {#if member.biography}<p class="mt-3 text-sm leading-6 text-slate-600">{member.biography}</p>{/if}
  </div>
</article>
