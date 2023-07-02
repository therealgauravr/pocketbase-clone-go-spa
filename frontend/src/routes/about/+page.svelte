<script lang="ts">
	export let data;
</script>

<svelte:head>
	<title>About</title>
</svelte:head>

<h1>About</h1>

<p>
	<strong>server respond</strong>:
	{#if data}
		{data?.message}
	{:else}
		{"No data :("}
	{/if}
</p>

<p>This is about page</p>
