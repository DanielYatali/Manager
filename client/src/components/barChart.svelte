<script>
	// chart 2
	import { Chart } from 'chart.js/auto/auto.js';
	import { onMount } from 'svelte';
	let ctx2;

	onMount(() => {
		ctx2 = ctx2.getContext('2d');
		var gradientStroke1 = ctx2.createLinearGradient(5, 0, 5, 100);

		gradientStroke1.addColorStop(1, 'rgba(26, 115, 232, 0.02)');
		gradientStroke1.addColorStop(0, 'rgba(26, 115, 232, 0.18)'); //purple colors
		var gradientStroke2 = ctx2.createLinearGradient(100, 0, 50, 150);
		gradientStroke2.addColorStop(1, 'rgba(0, 208, 255, 0.03)');
		gradientStroke2.addColorStop(0, 'rgba(0, 208, 255, 0.19)'); //purple colors

		new Chart(ctx2, {
			type: 'bar',
			data: {
				labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri'],
				datasets: [
					{
						label: 'Last Week',
						tension: 0.4,
						pointRadius: 0,
						borderColor: '#1F3BB3',
						borderWidth: 3,
						backgroundColor: gradientStroke1,
						fill: true,
						data: [5, 4, 6, 8, 2],
						maxBarThickness: 6
					},
					{
						label: 'This Week',
						tension: 0.4,
						pointRadius: 0,
						borderColor: '#52CDFF',
						borderWidth: 3,
						backgroundColor: gradientStroke2,
						fill: true,
						data: [4, 8, 6, 3, 5],
						maxBarThickness: 6
					}
				]
			},
			options: {
				responsive: true,
				maintainAspectRatio: false,
				plugins: {
					legend: {
						display: false
					}
				},
				interaction: {
					intersect: false,
					mode: 'index'
				},
				scales: {
					y: {
						grid: {
							drawBorder: false,
							display: true,
							drawOnChartArea: true,
							drawTicks: false,
							borderDash: [5, 5]
						},
						ticks: {
							display: true,
							padding: 10,
							color: '#b2b9bf',
							font: {
								size: 11,
								family: 'Open Sans',
								style: 'normal',
								lineHeight: 2
							}
						}
					},
					x: {
						grid: {
							drawBorder: false,
							display: false,
							drawOnChartArea: false,
							drawTicks: false,
							borderDash: [5, 5]
						},
						ticks: {
							display: true,
							color: '#b2b9bf',
							padding: 20,
							font: {
								size: 11,
								family: 'Open Sans',
								style: 'normal',
								lineHeight: 2
							}
						}
					}
				}
			}
		});
	});
</script>

<div class="w-full max-w-full px-6 lg:flex-none mt-6 container h-full">
	<div
		class="border-black/12.5 shadow-soft-xl relative flex min-w-0 flex-col break-words rounded-2xl border-0 border-solid bg-white bg-clip-border"
	>
		<div class="flex justify-between">
			<h4 class="font-bold">Performance Chart</h4>
			<div class="gap-10 text-xs flex light-gray-text">
				<div class=" text-center flex gap-2">
					<div class="blue-bg w-2 h-2 rounded-full mt-1" />
					<span>This Week</span>
				</div>

				<div class="flex gap-2">
					<div class="light-blue-bg w-2 h-2 rounded-full mt-1" />
					<span> Last Week </span>
				</div>
			</div>
		</div>
		<h5 class="nav-gray-text text-sm mt-2">Lorem Ipsum is simply dummy text of the printing</h5>
		<br />
		<div class=" h-full w-full">
			<div>
				<canvas bind:this={ctx2} id="chart-line" height="250" />
			</div>
		</div>
	</div>
</div>
