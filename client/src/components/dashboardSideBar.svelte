<script>
	import { goto } from '$app/navigation';
	let openSideBar = 'flex';
	let closeSideBar = 'none';
	let sideBarMove = 'translateX(-260px)';
	function sidebarHandler() {
		if (openSideBar == 'flex') {
			openSideBar = 'none';
			closeSideBar = 'flex';
			sideBarMove = 'translateX(0px)';
		} else {
			closeSideBar = 'none';
			openSideBar = 'flex';
			sideBarMove = 'translateX(-260px)';
		}
	}
	//Options
	const options = [
		{
			name: 'Dashboard',
			icon: `<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-gauge" width="32" height="32" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <circle cx="12" cy="12" r="9"></circle>
   <circle cx="12" cy="12" r="1"></circle>
   <line x1="13.41" y1="10.59" x2="16" y2="8"></line>
   <path d="M7 12a5 5 0 0 1 5 -5"></path>
</svg>`
		},
		{
			name: 'Employees',
			icon: `<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-user" width="32" height="32" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <circle cx="12" cy="7" r="4"></circle>
   <path d="M6 21v-2a4 4 0 0 1 4 -4h4a4 4 0 0 1 4 4v2"></path>
</svg>`
		},
		{
			name: 'Clients',
			icon: `<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-users" width="32" height="32" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <circle cx="9" cy="7" r="4"></circle>
   <path d="M3 21v-2a4 4 0 0 1 4 -4h4a4 4 0 0 1 4 4v2"></path>
   <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
   <path d="M21 21v-2a4 4 0 0 0 -3 -3.85"></path>
</svg>`
		},
		{
			name: 'Projects',
			icon: `<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-rocket" width="32" height="32" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <path d="M4 13a8 8 0 0 1 7 7a6 6 0 0 0 3 -5a9 9 0 0 0 6 -8a3 3 0 0 0 -3 -3a9 9 0 0 0 -8 6a6 6 0 0 0 -5 3"></path>
   <path d="M7 14a6 6 0 0 0 -3 6a6 6 0 0 0 6 -3"></path>
   <circle cx="15" cy="9" r="1"></circle>
</svg>`
		},
		{
			name: 'Tasks',
			icon: `<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-list" width="32" height="32" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <line x1="9" y1="6" x2="20" y2="6"></line>
   <line x1="9" y1="12" x2="20" y2="12"></line>
   <line x1="9" y1="18" x2="20" y2="18"></line>
   <line x1="5" y1="6" x2="5" y2="6.01"></line>
   <line x1="5" y1="12" x2="5" y2="12.01"></line>
   <line x1="5" y1="18" x2="5" y2="18.01"></line>
</svg>`
		},
		{
			name: 'Teams',
			icon: `<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-gauge" width="32" height="32" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <circle cx="12" cy="12" r="9"></circle>
   <circle cx="12" cy="12" r="1"></circle>
   <line x1="13.41" y1="10.59" x2="16" y2="8"></line>
   <path d="M7 12a5 5 0 0 1 5 -5"></path>
</svg>`
		},
		{
			name: 'Payroll',
			icon: `<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-cash" width="32" height="32" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <rect x="7" y="9" width="14" height="10" rx="2"></rect>
   <circle cx="14" cy="14" r="2"></circle>
   <path d="M17 9v-2a2 2 0 0 0 -2 -2h-10a2 2 0 0 0 -2 2v6a2 2 0 0 0 2 2h2"></path>
</svg>`
		},
		{
			name: 'Settings',
			icon: `<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-settings" width="32" height="32" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <path d="M10.325 4.317c.426 -1.756 2.924 -1.756 3.35 0a1.724 1.724 0 0 0 2.573 1.066c1.543 -.94 3.31 .826 2.37 2.37a1.724 1.724 0 0 0 1.065 2.572c1.756 .426 1.756 2.924 0 3.35a1.724 1.724 0 0 0 -1.066 2.573c.94 1.543 -.826 3.31 -2.37 2.37a1.724 1.724 0 0 0 -2.572 1.065c-.426 1.756 -2.924 1.756 -3.35 0a1.724 1.724 0 0 0 -2.573 -1.066c-1.543 .94 -3.31 -.826 -2.37 -2.37a1.724 1.724 0 0 0 -1.065 -2.572c-1.756 -.426 -1.756 -2.924 0 -3.35a1.724 1.724 0 0 0 1.066 -2.573c-.94 -1.543 .826 -3.31 2.37 -2.37c1 .608 2.296 .07 2.572 -1.065z"></path>
   <circle cx="12" cy="12" r="3"></circle>
</svg>`
		},
		{
			name: 'Log out',
			icon: `<svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-logout" width="32" height="32" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
   <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
   <path d="M14 8v-2a2 2 0 0 0 -2 -2h-7a2 2 0 0 0 -2 2v12a2 2 0 0 0 2 2h7a2 2 0 0 0 2 -2v-2"></path>
   <path d="M7 12h14l-3 -3m0 6l3 -3"></path>
</svg>`
		}
	];

	const changeOption = (option) => {
		if (option == 'Chat') {
			goto('/chat');
		} else if (option == 'Questions') {
			goto('/questions');
		} else {
			goto('/match');
		}
	};
</script>

<!-- Desktop view side bar-->
<div class="w-56 absolute md:relative border md:h-full flex-col hidden md:flex">
	<div>
		<ul>
			{#each options as option}
				<li
					class="flex w-full  nav-gray-text hover:text-blue-600 hover:bg-white cursor-pointer p-3 pl-10 rounded-r-full mb-1 "
				>
					<button
						on:click|preventDefault={() => changeOption(option)}
						class="flex items-center focus:outline-none focus:ring-2 focus:ring-white"
					>
						{@html option.icon}

						<span class="text-sm ml-2">{option.name} </span>
					</button>
				</li>
			{/each}
		</ul>
	</div>
</div>

<!-- Mobile view side bar -->
<div
	style="min-height:120vh;  transform: {sideBarMove};"
	class="w-64 z-40 absolute   shadow md:h-full flex-col justify-between md:hidden transition duration-150 ease-in-out"
	id="mobile-nav"
>
	<button
		style="display: {openSideBar};"
		aria-label="toggle sidebar"
		id="openSideBar"
		class="h-10 w-10   absolute right-0 mt-16 -mr-10 flex items-center shadow rounded-tr rounded-br justify-center cursor-pointer focus:outline-none focus:ring-2 focus:ring-offset-2 rounded focus:ring-gray-800"
		on:click={sidebarHandler}
	>
		<img
			src="https://tuk-cdn.s3.amazonaws.com/can-uploader/light_with_icons_at_bottom-svg7.svg"
			alt="toggler"
		/>
	</button>
	<button
		style="display: {closeSideBar};"
		aria-label="Close sidebar"
		id="closeSideBar"
		class="h-10 w-10 nav-blue-bg absolute right-0 mt-16 -mr-10 flex items-center shadow rounded-tr rounded-br justify-center cursor-pointer text-white"
		on:click={sidebarHandler}
	>
		<img
			src="https://tuk-cdn.s3.amazonaws.com/can-uploader/light_with_icons_at_bottom-svg8.svg"
			alt="cross"
		/>
	</button>
	<div class="px-8">
		<ul class="mt-12">
			{#each options as option}
				<li
					class="flex w-full justify-between  nav-gray-text hover:text-yellow-400 cursor-pointer items-center mb-6"
				>
					<button
						on:click|preventDefault={() => changeOption(option)}
						class="flex items-center focus:outline-none focus:ring-2 focus:ring-white"
					>
						{@html option.icon}
						<span class="text-sm ml-2">{option.name}</span>
					</button>
				</li>
			{/each}
		</ul>
	</div>
</div>
