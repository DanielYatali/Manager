<script>
	import { endpoints } from '$lib/endpoints';
	let jobTitle = '';
	let jobDescription = '';
	//Fetches all jobTitles
	const createJob = async () => {
		if (jobTitle != '' && jobDescription != '') {
			try {
				const rawResponse = await fetch(endpoints.server + '/jobs', {
					method: 'POST',
					headers: {
						Accept: 'application/json',
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						title: jobTitle,
						description: jobDescription
					})
				});
				let response = await rawResponse.json();
				console.log(response);
			} catch (error) {
				console.error(error);
			}
		}
	};
</script>

<div>
	<div class="flex flex-col sm:flex-row gap-2 m-4">
		<input
			bind:value={jobTitle}
			class="shadow border rounded w-full py-2 px-3 mr-4 text-gray-600"
			placeholder="Job Title"
		/>
		<input
			bind:value={jobDescription}
			class="shadow border rounded w-full py-2 px-3 mr-4 text-gray-600"
			placeholder="Job Description"
		/>
		<button
			on:click|preventDefault={createJob}
			class="flex-no-shrink p-2 rounded-lg text-white bg-gray-500 hover:text-red hover:bg-gray-700"
			>create</button
		>
	</div>
</div>
