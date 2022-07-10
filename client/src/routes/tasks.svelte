<script>
	import TaskTable from '../components/taskTable.svelte';
</script>

<div class="bg-blue w-full h-screen font-sans">
	<div class="flex px-4 pb-8 items-start overflow-x-scroll lg:overflow-hidden">
		<TaskTable name="Todo" />
		<TaskTable name="In Progress" />
		<TaskTable name="Paused" />
		<TaskTable name="Done" />
	</div>
</div>
<svelte:head>
	<style>
		body {
			background-color: #f4f5f7;
		}
	</style>

	<title>Tasks</title>
</svelte:head>
