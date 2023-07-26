<script>
	import "../app.postcss";
</script>

<div class=" flex flex-col w-screen bg-gradient-to-tr from-emerald-300 to-teal-500">
	<div class="text-3xl font-bold m-5">RocketBase</div>
	<div
		class=" flex flex-row p-4 space-x-4 h-screen "
	>
		<slot />
	</div>
</div>
