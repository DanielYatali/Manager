<script>
	import SideBar from '../components/sideBar.svelte';

	import Nav from '../components/nav.svelte';

	import '../app.css';
</script>

<main>
	<div class="sticky top-0 w-full z-40">
		<Nav />
	</div>
	<div class="flex flex-no-wrap">
		<div class="fixed z-30">
			<SideBar />
		</div>
		<div class="md:ml-56 xl:ml-80 container mx-auto jmd:w-4/5 w-11/12 px-6 overflow--y-scroll ">
			<slot />
		</div>
	</div>
</main>

<svelte:head>
	<style>
		body {
			background-color: #f4f5f7;
		}
	</style>
</svelte:head>
