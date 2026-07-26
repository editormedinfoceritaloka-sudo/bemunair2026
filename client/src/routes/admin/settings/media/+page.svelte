<script lang="ts">
  import PageHeader from '$lib/admin/PageHeader.svelte';
  import { Button } from '$lib/components/ui/button';
  import { Input } from '$lib/components/ui/input';
  import { Label } from '$lib/components/ui/label';
  import { Textarea } from '$lib/components/ui/textarea';
  import { Alert, AlertDescription } from '$lib/components/ui/alert';
  import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
  import { enhance } from '$app/forms';
  import { invalidateAll } from '$app/navigation';
  import { toast } from 'svelte-sonner';
  let { data, form } = $props();
  const value = (row: any, snake: string, pascal: string, fallback: string | number = '') => row[snake] ?? row[pascal] ?? fallback;
  const done = () => () => async ({ update }: any) => {
    await update(); await invalidateAll(); toast.success('Pengaturan pengajuan disimpan');
  };
</script>

<PageHeader title="Pengaturan pengajuan media" description="Kelola SOP, template, PIC, ketentuan, dan aturan jadwal tanpa mengubah source code." />
{#if form?.error}<Alert variant="destructive" class="mb-4"><AlertDescription>{form.error}</AlertDescription></Alert>{/if}

<div class="grid gap-6 xl:grid-cols-2">
  {#each [['CONTENT', data.content, 'Konten'], ['ARTICLE', data.article, 'Artikel']] as [service, row, label]}
    <Card>
      <CardHeader><CardTitle>{label}</CardTitle></CardHeader>
      <CardContent>
        <form method="POST" action="?/save" use:enhance={done()} class="grid gap-4">
          <input type="hidden" name="service_type" value={service} />
          <div class="space-y-2"><Label>SOP URL</Label><Input name="sop_url" type="url" value={value(row, 'sop_url', 'SOPURL')} /></div>
          <div class="space-y-2"><Label>Template kementerian</Label><Input name="ministry_template_url" type="url" value={value(row, 'ministry_template_url', 'MinistryTemplateURL')} /></div>
          <div class="space-y-2"><Label>{service === 'CONTENT' ? 'Template brief' : 'Template caption'}</Label><Input name={service === 'CONTENT' ? 'brief_template_url' : 'caption_template_url'} type="url" value={service === 'CONTENT' ? value(row, 'brief_template_url', 'BriefTemplateURL') : value(row, 'caption_template_url', 'CaptionTemplateURL')} /></div>
          <div class="grid gap-4 sm:grid-cols-2"><div class="space-y-2"><Label>Nama PIC</Label><Input name="pic_name" value={value(row, 'pic_name', 'PICName')} /></div><div class="space-y-2"><Label>WhatsApp PIC</Label><Input name="pic_whatsapp" value={value(row, 'pic_whatsapp', 'PICWhatsApp')} /></div></div>
          <div class="space-y-2"><Label>Ketentuan <span class="text-muted-foreground">(satu per baris)</span></Label><Textarea name="terms" rows={7} value={((row as any).terms || []).join('\n')} required /></div>
          <div class="grid gap-4 sm:grid-cols-2"><div class="space-y-2"><Label>Minimal lead day</Label><Input name="minimum_lead_days" type="number" min="0" value={value(row, 'minimum_lead_days', 'MinimumLeadDays')} required /></div><div class="space-y-2"><Label>Interval slot (menit)</Label><Input name="slot_interval_minutes" type="number" min="15" value={value(row, 'slot_interval_minutes', 'SlotIntervalMinutes', 30)} required /></div><div class="space-y-2"><Label>Jam mulai</Label><Input name="publish_time_start" type="time" value={value(row, 'publish_time_start', 'PublishTimeStart', '08:00')} required /></div><div class="space-y-2"><Label>Jam selesai</Label><Input name="publish_time_end" type="time" value={value(row, 'publish_time_end', 'PublishTimeEnd', '17:00')} required /></div><div class="space-y-2 sm:col-span-2"><Label>Kapasitas harian <span class="text-muted-foreground">(opsional)</span></Label><Input name="daily_capacity" type="number" min="1" value={value(row, 'daily_capacity', 'DailyCapacity')} /></div></div>
          <Button type="submit" class="bg-blue-500">Simpan pengaturan {label}</Button>
        </form>
      </CardContent>
    </Card>
  {/each}
</div>
