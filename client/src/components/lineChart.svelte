<script>
	import Chart from 'chart.js/auto';
	import { onMount } from 'svelte';

	/**
	 * @type {HTMLCanvasElement}
	 */
	let myChart;
	Chart.defaults.font.family = 'Manrope';
	Chart.defaults.color = '#706f73';

	onMount(() => {
		let graphGradient = myChart.getContext('2d');
		let graphGradient2 = myChart.getContext('2d');
		let GradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
		GradientBg.addColorStop(0, 'rgba(26, 115, 232, 0.18)');
		GradientBg.addColorStop(1, 'rgba(26, 115, 232, 0.02)');
		var GradientBg2 = graphGradient2.createLinearGradient(100, 0, 50, 150);
		GradientBg2.addColorStop(0, 'rgba(0, 208, 255, 0.19)');
		GradientBg2.addColorStop(1, 'rgba(0, 208, 255, 0.03)');
		let lineChart = new Chart(myChart, {
			type: 'line',
			data: {
				labels: ['MON', 'TUE', 'WED', 'THU', 'FRI'],
				datasets: [
					{
						label: 'This Week',
						data: [5, 5, 3, 6, 7],
						backgroundColor: GradientBg,
						borderColor: '#1F3BB3',
						borderWidth: 2,
						fill: true,
						pointBorderWidth: 1,
						pointRadius: 4,
						pointHoverRadius: 2,
						pointBackgroundColor: '#1F3BB3',
						pointBorderColor: '#fff'
					},
					{
						label: 'Last Week',
						data: [3, 7, 4, 3, 6],
						backgroundColor: GradientBg2,
						borderColor: ['#52CDFF'],
						borderWidth: 2,
						fill: true,
						pointBorderWidth: 1,
						pointRadius: 4,
						pointHoverRadius: 2,
						pointBackgroundColor: '#52CDFF',
						pointBorderColor: '#fff'
					}
				]
			},
			options: {
				maintainAspectRatio: false,
				responsive: true,
				elements: {
					line: {
						tension: 0.4
					}
				},
				plugins: {
					legend: {
						display: false
					}
				},
				scales: {
					x: {
						grid: {
							drawBorder: false,
							display: false,
							lineWidth: 10
						}
					}
				}
			}
		});
	});
</script>

<div class="container p-6">
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
	<h5 class="nav-gray-text text-sm">Lorem Ipsum is simply dummy text of the printing</h5>
	<div class="flex flex-col items-center">
		<div class="container">
			<canvas
				style="display: block; width: 988px; height: 150px;"
				width="1976"
				height="300"
				bind:this={myChart}
				id="myChart"
			/>
		</div>
	</div>
</div>
