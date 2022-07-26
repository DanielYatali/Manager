<!-- Renders the dashboard -->
<script>
	import BarChart from '../../components/barChart.svelte';
	import TaskLineChart from '../../components/taskLineChart.svelte';
	import PromotionCircle from '../../components/promotionCircle.svelte';
</script>

<div class="mt-5">
	<div class="sm:px-6 2xl:container">
		<div class="grid gap-6 xl:grid-cols-3 md:grid-cols-2">
			<div class="flex justify-between p-6 items-center rounded-2xl  bg-white">
				<div class="card-body w-full">
					<div>
						<p class="text-lg font-medium text-gray-700">Number of Hours Worked</p>
						<div class="flex mt-2 justify-between">
							<h3 class="text-2xl font-bold">36</h3>
							<div class="flex gap-2 text-green-600">
								<svg
									xmlns="http://www.w3.org/2000/svg"
									xmlns:xlink="http://www.w3.org/1999/xlink"
									viewBox="0 0 122.88 68.04"
									width="34"
									height="17"
									stroke="currentColor"
									fill="currentColor"
								>
									><path
										d="M2.03,56.52c-2.66,2.58-2.72,6.83-0.13,9.49c2.58,2.66,6.83,2.72,9.49,0.13l27.65-26.98l23.12,22.31 c2.67,2.57,6.92,2.49,9.49-0.18l37.77-38.22v19.27c0,3.72,3.01,6.73,6.73,6.73s6.73-3.01,6.73-6.73V6.71h-0.02 c0-1.74-0.67-3.47-2-4.78c-1.41-1.39-3.29-2.03-5.13-1.91H82.4c-3.72,0-6.73,3.01-6.73,6.73c0,3.72,3.01,6.73,6.73,6.73h17.63 L66.7,47.2L43.67,24.97c-2.6-2.5-6.73-2.51-9.33,0.03L2.03,56.52L2.03,56.52z"
									/>
									></svg
								>
								+20%
							</div>
						</div>
					</div>
				</div>
			</div>

			<div class="flex justify-between p-6 items-center rounded-2xl  bg-white">
				<div class="card-body w-full">
					<div>
						<p class="text-lg font-medium text-gray-700">Task Completed</p>
						<div class="flex mt-2 justify-between">
							<h3 class="text-2xl font-bold">36</h3>
							<div class="flex gap-2 text-red-600">
								<svg
									xmlns="http://www.w3.org/2000/svg"
									xmlns:xlink="http://www.w3.org/1999/xlink"
									viewBox="0 0 122.88 68.04"
									width="34"
									height="17"
									stroke="currentColor"
									fill="currentColor"
								>
									><path
										d="M2.03,11.52C-0.63,8.94-0.68,4.69,1.9,2.03c2.58-2.66,6.83-2.72,9.49-0.13l27.65,26.98L62.16,6.57 c2.67-2.57,6.92-2.49,9.49,0.18l37.77,38.22V25.7c0-3.72,3.01-6.73,6.73-6.73s6.73,3.01,6.73,6.73v35.63h-0.02 c0,1.74-0.67,3.47-2,4.78c-1.41,1.39-3.29,2.03-5.13,1.91H82.4c-3.72,0-6.73-3.01-6.73-6.73c0-3.72,3.01-6.73,6.73-6.73h17.63 L66.7,20.84L43.67,43.07c-2.6,2.5-6.73,2.51-9.33-0.03L2.03,11.52L2.03,11.52z"
									/>
									></svg
								>
								-20%
							</div>
						</div>
					</div>
				</div>
			</div>
			<div class="flex justify-between p-6 items-center rounded-2xl  bg-white">
				<div class="card-body w-full">
					<div>
						<p class="text-lg font-medium text-gray-700">Rank</p>
						<div class="flex mt-2 justify-between">
							<h3 class="text-2xl font-bold">7th</h3>
							<div class="flex gap-2 text-red-600">
								<svg
									xmlns="http://www.w3.org/2000/svg"
									xmlns:xlink="http://www.w3.org/1999/xlink"
									viewBox="0 0 122.88 68.04"
									width="34"
									height="17"
									stroke="currentColor"
									fill="currentColor"
								>
									><path
										d="M2.03,11.52C-0.63,8.94-0.68,4.69,1.9,2.03c2.58-2.66,6.83-2.72,9.49-0.13l27.65,26.98L62.16,6.57 c2.67-2.57,6.92-2.49,9.49,0.18l37.77,38.22V25.7c0-3.72,3.01-6.73,6.73-6.73s6.73,3.01,6.73,6.73v35.63h-0.02 c0,1.74-0.67,3.47-2,4.78c-1.41,1.39-3.29,2.03-5.13,1.91H82.4c-3.72,0-6.73-3.01-6.73-6.73c0-3.72,3.01-6.73,6.73-6.73h17.63 L66.7,20.84L43.67,43.07c-2.6,2.5-6.73,2.51-9.33-0.03L2.03,11.52L2.03,11.52z"
									/>
									></svg
								>
								-20%
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
		<br />
		<div class="grid grid-cols-1 gap-6 lg:grid-cols-3">
			<div class="p-2 rounded-2xl  bg-white col-span-2">
				<BarChart />
			</div>
			<div class="grid bg-white rounded-2xl col-span-1 w-full">
				<!-- Circle -->
				<PromotionCircle />
			</div>
		</div>
		<br />
		<div class="grid grid-cols-1">
			<div class="p-6 rounded-2xl  bg-white col-span-1">
				<TaskLineChart />
			</div>
		</div>
	</div>
</div>

<svelte:head>
	<style>
		body {
			background-color: #f4f5f7;
		}
	</style>

	<title>Dashboard</title>
</svelte:head>
