<script>
	import TodoItem from '../components/todoItem.svelte';
</script>

<div class="container p-6">
	<div class="flex justify-between">
		<div>
			<p>Todo List</p>
		</div>
		<div>icon</div>
	</div>
	<br />
	<div>
		<TodoItem />
		<TodoItem />
		<TodoItem />
		<!-- <div class="flex flex-col gap-2">
			<div class="flex gap-3">
				<input class="" type="checkbox" />
				<p>Lorem ipsum dolor sit amet consectetur adipisicing elit.</p>
			</div>
			<div class="flex gap-2 ml-6">
				<p>24 June 2019</p>
				<p class="bg-yellow-600 rounded-lg p-1.5 -mt-1.5">Due Tom</p>
			</div>
			<div class="w-full h-0.5 bg-gray-200 mt-3" />
		</div> -->
	</div>
</div>
