<script>
	import { endpoints } from '$lib/endpoints';
	import { onMount } from 'svelte';
	// let jobTitle = "";
	// let jobDescription = "";
	let jobs = [];
	//Fetches all jobTitles

	onMount(() => {
		getJobs();
	});

	const getJobs = async () => {
		try {
			const rawResponse = await fetch(endpoints.server + '/jobs', {
				method: 'GET',
				headers: {
					Accept: 'application/json',
					'Content-Type': 'application/json'
				}
			});
			let response = await rawResponse.json();
			jobs = response.items;
			console.log(jobs);
		} catch (error) {
			console.error(error);
		}
	};
</script>

<div>
	<select name="jobs" id="jobs">
		{#each jobs as job}
			<option value={job.title}>{job.title}</option>
		{/each}
	</select>
</div>
