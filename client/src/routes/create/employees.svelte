<!-- Homepage of todo app -->
<script>
	// @ts-ignore
	import { endpoints } from '$lib/endpoints';
	import JobTitles from '../../components/jobTitles.svelte';
	import { onMount } from 'svelte';

	let response = {};
	let employees = [];
	let firstName = '';
	let lastName = '';
	let jobTitle = '';
	let age = '';
	let dob = '';
	let email = '';

	//When component mounted fetch all todos
	// onMount(() => {
	//   getTodos();
	// });

	//Fetches all todos from WeOS server
	const getEmployees = async () => {
		try {
			const rawResponse = await fetch(endpoints.server + '/employees', {
				method: 'GET',
				headers: {
					Accept: 'application/json',
					'Content-Type': 'application/json'
				}
			});
			response = await rawResponse.json();
			employees = response.items;
			console.log(employees);
		} catch (error) {
			console.error(error);
		}
	};

	//Creates a todo
	const createEmployee = async () => {
		//Prevent user from making blank todo
		if (firstName != '' && lastName != '') {
			let employee = {
				firstName: firstName,
				lastName: lastName,
				email: email,
				age: age,
				dob: dob,
				jobTitle: jobTitle
			};
			try {
				const rawResponse = await fetch(endpoints.server + '/employees', {
					method: 'POST',
					headers: {
						Accept: 'application/json',
						'Content-Type': 'application/json'
					},
					body: JSON.stringify(employee)
				});
				response = await rawResponse.json();
				console.log(response);
				// description = "";
			} catch (error) {
				console.error(error);
			}
		}
	};

	// //Removes a todo
	// const removeTodo = async (id) => {
	//   try {
	//     const rawResponse = await fetch(endpoints.server + "/todos/" + id, {
	//       method: "DELETE",
	//       headers: {
	//         Accept: "application/json",
	//         "Content-Type": "application/json",
	//       },
	//     });
	//     response = await rawResponse.json();
	//     // console.log(response);
	//     //Fetches updates todo list
	//     getTodos();
	//   } catch (error) {
	//     console.error(error);
	//   }
	// };

	// //Edits todo by sending updated values to WeOS server
	// const editTodo = async (updateTodo) => {
	//   try {
	//     const rawResponse = await fetch(
	//       endpoints.server + "/todos/" + updateTodo.id,
	//       {
	//         method: "PUT",
	//         headers: {
	//           Accept: "application/json",
	//           "Content-Type": "application/json",
	//         },
	//         body: JSON.stringify(updateTodo),
	//       }
	//     );
	//     let response = await rawResponse.json();
	//     // console.log(response);
	//     //Fetching updated todo list
	//     getTodos();
	//   } catch (error) {
	//     console.error(error);
	//   }
	// };
</script>

<div class="h-100 w-full flex items-center justify-center bg-teal-lightest">
	<div class="bg-red rounded shadow p-6 m-4 w-full lg:w-3/4 bg-gray-200">
		<h1 class="text-3xl text-center text-gray-700 pb-4">Employee List</h1>
		<div class="mb-4">
			<div class="w-full flex flex-col gap-2">
				<div class="w-full flex  gap-2 flex-col sm:flex-row items-center">
					<!-- binds the title variable to the users input -->
					<input
						bind:value={firstName}
						type="text"
						class="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-gray-600"
						placeholder="First Name"
					/>
					<!-- binds the description variable to the users input -->
					<input
						bind:value={lastName}
						type="text"
						class="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-gray-600"
						placeholder="Last Name"
					/>
				</div>
				<div class="w-full flex flex-col gap-2 sm:flex-row items-center">
					<!-- binds the title variable to the users input -->
					<input
						bind:value={age}
						type="number"
						class="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-gray-600"
						placeholder="age"
					/>
					<!-- binds the description variable to the users input -->
					<input
						bind:value={dob}
						type="date"
						class="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-gray-600"
						placeholder="dob"
					/>
				</div>
				<div class="w-full flex flex-col  gap-2 sm:flex-row items-center">
					<div>
						<JobTitles />
					</div>
					<div>
						<input
							bind:value={email}
							type="email"
							class="shadow appearance-none border rounded w-full py-2 px-3 mr-4 text-gray-600"
							placeholder="Email"
						/>
					</div>
				</div>

				<button
					on:click|preventDefault={createEmployee}
					class="flex-no-shrink p-2 rounded-lg text-white bg-gray-500 hover:text-red hover:bg-gray-700"
					>Add</button
				>
			</div>

			<div class="flex flex-col-reverse">
				{#each employees as employee, Index (employee.id)}
					<div class="shadow m-2 bg-gray-100 rounded-lg">
						<!-- <TodoInput {todo} {removeTodo} {editTodo} /> -->
					</div>
				{/each}
			</div>
		</div>
	</div>
</div>

<svelte:head>
	<style>
		body {
			--tw-bg-opacity: 1;
			background-color: rgb(243 244 246 / var(--tw-bg-opacity));
		}
	</style>
</svelte:head>
