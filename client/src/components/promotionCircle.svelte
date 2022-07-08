<div class="grid">
	<div class="inline-flex items-center justify-center rounded-full">
		<!-- Building a Progress Ring: https://css-tricks.com/building-progress-ring-quickly/ -->
		<svg class="w-80 h-80 -rotate-90">
			<circle
				class="text-gray-200"
				stroke-width="25"
				stroke="currentColor"
				fill="transparent"
				r="120"
				cx="160"
				cy="160"
			/>
			<circle
				class="text-blue-800"
				stroke-width="25"
				stroke-dasharray="300, 800"
				stroke-dashoffset="1"
				stroke-linecap="round"
				stroke="currentColor"
				fill="transparent"
				r="120"
				cx="160"
				cy="160"
			/>
		</svg>
		<span class="absolute text-5xl text-blue-800 font-bold">45</span>
	</div>
	<p class="text-center pb-5 font-bold text-xl">Promotion Satus</p>
</div>
