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
	<div
		style="min-height: 93vh;"
		class="md:ml-64 xl:ml-80 container mx-auto py-2 sm:py-10 md:w-4/5 w-11/12 px-6 overflow--y-scroll "
	>
		<div class="">
			<div class="sm:px-6 2xl:container">
				<div class="grid gap-6 xl:grid-cols-4 md:grid-cols-2">
					<div class="flex dashboard-bg justify-between p-4 items-center rounded-lg">
						<div class="text-gold bg-white rounded-full p-1">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="icon icon-tabler icon-tabler-rocket"
								width="34"
								height="34"
								viewBox="0 0 24 24"
								stroke-width="2"
								stroke="currentColor"
								fill="none"
								stroke-linecap="round"
								stroke-linejoin="round"
							>
								<path stroke="none" d="M0 0h24v24H0z" fill="none" />
								<path
									d="M4 13a8 8 0 0 1 7 7a6 6 0 0 0 3 -5a9 9 0 0 0 6 -8a3 3 0 0 0 -3 -3a9 9 0 0 0 -8 6a6 6 0 0 0 -5 3"
								/>
								<path d="M7 14a6 6 0 0 0 -3 6a6 6 0 0 0 6 -3" />
								<circle cx="15" cy="9" r="1" />
							</svg>
						</div>
						<div class="flex flex-col gap-2 text-gray-200">
							<div class="text-xl font-bold">
								<span class="">122</span>
							</div>
							<div>Projects</div>
						</div>
					</div>

					<div class="flex justify-between p-4  dashboard-bg items-center rounded-lg">
						<div class="text-gold bg-white rounded-full p-1">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="icon icon-tabler icon-tabler-rocket"
								width="34"
								height="34"
								viewBox="0 0 24 24"
								stroke-width="2"
								stroke="currentColor"
								fill="none"
								stroke-linecap="round"
								stroke-linejoin="round"
							>
								<path stroke="none" d="M0 0h24v24H0z" fill="none" />
								<path
									d="M4 13a8 8 0 0 1 7 7a6 6 0 0 0 3 -5a9 9 0 0 0 6 -8a3 3 0 0 0 -3 -3a9 9 0 0 0 -8 6a6 6 0 0 0 -5 3"
								/>
								<path d="M7 14a6 6 0 0 0 -3 6a6 6 0 0 0 6 -3" />
								<circle cx="15" cy="9" r="1" />
							</svg>
						</div>
						<div class="flex flex-col gap-2 text-gray-200">
							<div class="text-xl font-bold">122</div>
							<div>Clients</div>
						</div>
					</div>

					<div class="flex  dashboard-bg justify-between p-4 items-center rounded-lg">
						<div class="text-gold bg-white rounded-full p-1">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="icon icon-tabler icon-tabler-rocket"
								width="34"
								height="34"
								viewBox="0 0 24 24"
								stroke-width="2"
								stroke="currentColor"
								fill="none"
								stroke-linecap="round"
								stroke-linejoin="round"
							>
								<path stroke="none" d="M0 0h24v24H0z" fill="none" />
								<path
									d="M4 13a8 8 0 0 1 7 7a6 6 0 0 0 3 -5a9 9 0 0 0 6 -8a3 3 0 0 0 -3 -3a9 9 0 0 0 -8 6a6 6 0 0 0 -5 3"
								/>
								<path d="M7 14a6 6 0 0 0 -3 6a6 6 0 0 0 6 -3" />
								<circle cx="15" cy="9" r="1" />
							</svg>
						</div>
						<div class="flex flex-col gap-2 text-gray-200">
							<div class="text-xl font-bold">122</div>
							<div>Tasks</div>
						</div>
					</div>

					<div class="flex dashboard-bg  justify-between p-4 items-center rounded-lg">
						<div class="text-gold bg-white rounded-full p-1">
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="icon icon-tabler icon-tabler-rocket"
								width="34"
								height="34"
								viewBox="0 0 24 24"
								stroke-width="2"
								stroke="currentColor"
								fill="none"
								stroke-linecap="round"
								stroke-linejoin="round"
							>
								<path stroke="none" d="M0 0h24v24H0z" fill="none" />
								<path
									d="M4 13a8 8 0 0 1 7 7a6 6 0 0 0 3 -5a9 9 0 0 0 6 -8a3 3 0 0 0 -3 -3a9 9 0 0 0 -8 6a6 6 0 0 0 -5 3"
								/>
								<path d="M7 14a6 6 0 0 0 -3 6a6 6 0 0 0 6 -3" />
								<circle cx="15" cy="9" r="1" />
							</svg>
						</div>
						<div class="flex flex-col gap-2 text-gray-200">
							<div class="text-xl font-bold">122</div>
							<div>Employees</div>
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
	<title>Dashboard</title>
</svelte:head>
