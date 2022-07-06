<!-- Renders the dashboard -->
<script>
	import Nav from '../components/nav.svelte';
	import DashboardSideBar from '../components/dashboardSideBar.svelte';
	import { onMount } from 'svelte/internal';
	import { goto } from '$app/navigation';

	//Checks if newuser of logged in
	//If new user then send the user to fill out the question
	//if login and not a new user then they can access the dashboard
	// onMount(() => {
	// 	if (!$CurrentUser.loggedIn) {
	// 		goto('/login');
	// 	}
	// 	if ($CurrentUser.newUser) {
	// 		goto('/questions');
	// 	}
	// });
</script>

<div class="sticky top-0 w-full z-40">
	<Nav />
</div>
<div class="flex flex-no-wrap">
	<div class="fixed z-30">
		<DashboardSideBar />
	</div>

	<div class="md:ml-64 xl:ml-80 container mx-auto jmd:w-4/5 w-11/12 px-6 overflow--y-scroll ">
		<div>
			<h1 class="text-xl font-bold text-center">Overview</h1>
		</div>

		<div class="mt-5">
			<div class="sm:px-6 2xl:container">
				<div class="grid gap-6 xl:grid-cols-3 md:grid-cols-2">
					<div class="flex justify-between p-4 items-center rounded-lg border">
						<div class="card-body w-full">
							<div>
								<p class="">Number Hours Worked</p>
								<div class="flex mt-2 justify-between">
									<h3>36</h3>
									<div class="flex gap-2">
										<svg
											xmlns="http://www.w3.org/2000/svg"
											xmlns:xlink="http://www.w3.org/1999/xlink"
											viewBox="0 0 122.88 68.04"
											width="34"
											height="17"
											stroke="currentColor"
											fill="currentColor"
											class="text-green-600"
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

					<div class="flex justify-between p-4 items-center rounded-lg border">
						<div class="card-body w-full">
							<div>
								<p class="">Tasks Completed</p>
								<div class="flex mt-2 justify-between">
									<h3>36</h3>
									<div class="flex gap-2">
										<svg
											xmlns="http://www.w3.org/2000/svg"
											xmlns:xlink="http://www.w3.org/1999/xlink"
											viewBox="0 0 122.88 68.04"
											width="34"
											height="17"
											stroke="currentColor"
											fill="currentColor"
											class="text-red-600"
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
					<div class="flex justify-between p-4 items-center rounded-lg border">
						<div class="card-body w-full">
							<div>
								<p class="">Rank</p>
								<div class="flex mt-2 justify-between">
									<h3>7th</h3>
									<div class="flex gap-2">
										<svg
											xmlns="http://www.w3.org/2000/svg"
											xmlns:xlink="http://www.w3.org/1999/xlink"
											viewBox="0 0 122.88 68.04"
											width="34"
											height="17"
											stroke="currentColor"
											fill="currentColor"
											class="text-red-600"
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
				<!-- 
				<div class="grid gap-6 xl:grid-cols-3">
					<div>
						<CarouselGroupCard name={'Popular Groups'} />
					</div>
					<div>
						<CarouselGroupCard name={'Suggested Groups'} />
					</div>
					<div>
						<CarouselGroupCard name={'Some other listing'} />
					</div>
				</div> -->
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
