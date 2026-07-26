<script lang="ts">
  import MediaSubmissionForm from '$lib/admin/MediaSubmissionForm.svelte';
  let { data, form } = $props();
</script>

<svelte:head><title>Pengajuan Artikel · BEM UNAIR</title></svelte:head>
<MediaSubmissionForm service="ARTICLE" user={data.user} setting={data.setting} {form} />
