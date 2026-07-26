<script lang="ts">
  import { enhance } from '$app/forms';
  import { Button } from '$lib/components/ui/button';
  import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '$lib/components/ui/card';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { Eye, EyeOff, LoaderCircle, LockKeyhole } from '@lucide/svelte';
  let { form } = $props();
  let showPassword = $state(false); let pending = $state(false);
</script>
<svelte:head><title>Login Admin · BEM UNAIR</title></svelte:head>
<div class="grid min-h-screen place-items-center bg-background px-4 py-12">
  <div class="w-full max-w-md">
    <Card class="border-border bg-card shadow-lg shadow-black-900/5">
      <CardHeader class="space-y-4 pb-4">
        <div class="flex items-center justify-center gap-4" aria-label="Universitas Airlangga dan BEM UNAIR">
          <div class="grid size-20 place-items-center rounded-2xl border border-blue-100 bg-white p-2 shadow-sm">
            <img src="/brand/unair-logo.png" alt="Logo Universitas Airlangga" class="size-full object-contain" />
          </div>
          <div class="h-12 w-px bg-border"></div>
          <div class="grid size-20 place-items-center rounded-2xl border border-blue-100 bg-white p-2 shadow-sm">
            <img src="/brand/bem-unair-2026-logo.png" alt="Logo BEM UNAIR 2026 Kabinet Cerita Loka" class="size-full object-contain" />
          </div>
        </div>
        <div class="text-center"><CardTitle class="text-2xl text-foreground">Selamat datang kembali</CardTitle><CardDescription>Masuk untuk mengelola operasional BEM UNAIR.</CardDescription></div>
      </CardHeader>
      <CardContent>
        {#if form?.error}<Alert variant="destructive" class="mb-4"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}
        <form method="POST" use:enhance={() => { pending=true; return async ({ update }) => { await update(); pending=false; }; }} class="space-y-4">
          <div class="space-y-2"><Label for="email">Email</Label><Input id="email" name="email" type="email" autocomplete="email" required value={form?.email || ''} placeholder="admin@bem.unair.ac.id" /></div>
          <div class="space-y-2"><Label for="password">Password</Label><div class="relative"><Input id="password" name="password" type={showPassword?'text':'password'} autocomplete="current-password" required class="pr-10" /><button type="button" onclick={() => showPassword=!showPassword} class="absolute right-2 top-1/2 -translate-y-1/2 rounded-md p-1 text-muted-foreground transition-colors hover:bg-muted hover:text-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring" aria-label="Tampilkan password">{#if showPassword}<EyeOff class="size-4" />{:else}<Eye class="size-4" />{/if}</button></div></div>
          <Button type="submit" class="h-10 w-full bg-blue-500 shadow-sm hover:bg-blue-600 active:bg-blue-700" disabled={pending}>{#if pending}<LoaderCircle class="animate-spin" />{/if}Masuk ke Dashboard</Button>
        </form>
        <p class="mt-5 text-center text-xs text-muted-foreground">Akses terbatas untuk pengelola resmi BEM UNAIR.</p>
      </CardContent>
    </Card>
  </div>
</div>
