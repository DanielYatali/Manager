<script>
	// chart 2
	import { Chart } from 'chart.js/auto/auto.js';
	import { onMount } from 'svelte';
	let ctx2;

	onMount(() => {
		let ctx2 = document.getElementById('chart-line').getContext('2d');
		var gradientStroke1 = ctx2.createLinearGradient(0, 230, 0, 50);

		gradientStroke1.addColorStop(1, 'rgba(203,12,159,0.2)');
		gradientStroke1.addColorStop(0.2, 'rgba(72,72,176,0.0)');
		gradientStroke1.addColorStop(0, 'rgba(203,12,159,0)'); //purple colors

		var gradientStroke2 = ctx2.createLinearGradient(0, 230, 0, 50);

		gradientStroke2.addColorStop(1, 'rgba(20,23,39,0.2)');
		gradientStroke2.addColorStop(0.2, 'rgba(72,72,176,0.0)');
		gradientStroke2.addColorStop(0, 'rgba(20,23,39,0)'); //purple colors

		new Chart(ctx2, {
			type: 'line',
			data: {
				labels: ['Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
				datasets: [
					{
						label: 'Last Month',
						tension: 0.4,
						borderWidth: 0,
						pointRadius: 0,
						borderColor: '#cb0c9f',
						borderWidth: 3,
						backgroundColor: gradientStroke1,
						fill: true,
						data: [50, 40, 300, 220, 500, 250, 400, 230, 500],
						maxBarThickness: 6
					},
					{
						label: 'This Month',
						tension: 0.4,
						borderWidth: 0,
						pointRadius: 0,
						borderColor: '#3A416F',
						borderWidth: 3,
						backgroundColor: gradientStroke2,
						fill: true,
						data: [30, 90, 40, 140, 290, 290, 340, 230, 400],
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

<div class="w-full max-w-full px-3 mt-0 lg:flex-none">
	<div
		class="border-black/12.5 shadow-soft-xl relative z-20 flex min-w-0 flex-col break-words rounded-2xl border-0 border-solid bg-white bg-clip-border"
	>
		<div class="border-black/12.5 mb-0 rounded-t-2xl border-b-0 border-solid bg-white p-6 pb-0">
			<h6>Monthly Performance Chart</h6>
			<p class="leading-normal text-size-sm">
				<i class="fa fa-arrow-up text-lime-500" />
				<span class="font-semibold">4% more</span> in 2021
			</p>
		</div>
		<div class="flex-auto p-4">
			<div>
				<canvas bind:this={ctx2} id="chart-line" height="300" />
			</div>
		</div>
	</div>
</div>
