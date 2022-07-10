<script>
	import Task from './task.svelte';
	export let name;
	let tasks = ['Finish journal', 'Finish todo list', 'Finish todo list', 'Finish todo list'];
</script>

<div class="rounded bg-grey-light  flex-no-shrink w-64 p-2 mr-3">
	<div class="flex justify-between py-1">
		<h3 class="text-sm">{name}</h3>
		<svg
			class="h-4 fill-current text-grey-dark cursor-pointer"
			xmlns="http://www.w3.org/2000/svg"
			viewBox="0 0 24 24"
			><path
				d="M5 10a1.999 1.999 0 1 0 0 4 1.999 1.999 0 1 0 0-4zm7 0a1.999 1.999 0 1 0 0 4 1.999 1.999 0 1 0 0-4zm7 0a1.999 1.999 0 1 0 0 4 1.999 1.999 0 1 0 0-4z"
			/></svg
		>
	</div>
	<div class="text-sm mt-2">
		{#each tasks as task}
			<Task {task} />
		{/each}

		<p class="mt-3 text-grey-dark">Add a card...</p>
	</div>
</div>
