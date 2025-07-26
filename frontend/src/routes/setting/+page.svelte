<script>
	import * as Select from '$lib/components/ui/select/index.js';
	import Button from '@/components/ui/button/button.svelte';
	import { toast } from 'svelte-sonner';

	let settings = $state({
		checkTime: {
			hour: '',
			minute: '',
			second: ''
		}
	});

	const times = {
		hour: [...new Array(24)].map((_, i) => (i + 1).toString()),
		minute: [...new Array(61)].map((_, i) => i.toString())
	};

	const triggeredContentHour = $derived(
		times.hour.find((f) => f === settings.checkTime.hour) ?? 'Hour'
	);
	const triggeredContentMinute = $derived(
		times.minute.find((f) => f === settings.checkTime.minute) ?? 'Minute'
	);
	const triggeredContentSecond = $derived(
		times.minute.find((f) => f === settings.checkTime.second) ?? 'Second'
	);

	async function setCheckTime() {
		if (!settings.checkTime.hour && !settings.checkTime.minute && !settings.checkTime.second)
			toast.warning('Check time cant be empty');
	}
</script>

<div class="">
	<h1 class="mb-2.5 font-semibold">Check Time</h1>
	<form class="flex w-[50%] items-end gap-5" onsubmit={setCheckTime}>
		<div>
			<h1 class="mb-1.5 text-sm">Hour</h1>
			<Select.Root type="single" bind:value={settings.checkTime.hour} required>
				<Select.Trigger class="w-[180px]">{triggeredContentHour}</Select.Trigger>
				<Select.Content class="h-[50vh]">
					<Select.Label>Hour</Select.Label>
					{#each times.hour as hour}
						<Select.Item value={hour} label={hour}>{hour}</Select.Item>
					{/each}
				</Select.Content>
			</Select.Root>
		</div>

		<div>
			<h1 class="mb-1.5 text-sm">Minute</h1>
			<Select.Root type="single" bind:value={settings.checkTime.minute} required>
				<Select.Trigger class="w-[180px]">{triggeredContentMinute}</Select.Trigger>
				<Select.Content class="h-[50vh]">
					<Select.Label>Minute</Select.Label>
					{#each times.minute as minute}
						<Select.Item value={minute} label={minute}>{minute}</Select.Item>
					{/each}
				</Select.Content>
			</Select.Root>
		</div>

		<div>
			<h1 class="mb-1.5 text-sm">Second</h1>
			<Select.Root type="single" bind:value={settings.checkTime.second} required>
				<Select.Trigger class="w-[180px]">{triggeredContentSecond}</Select.Trigger>
				<Select.Content class="h-[50vh]">
					<Select.Label>Second</Select.Label>
					{#each times.minute as minute}
						<Select.Item value={minute} label={minute}>{minute}</Select.Item>
					{/each}
				</Select.Content>
			</Select.Root>
		</div>
		<Button class="text-sm" type="submit">Set check time</Button>
	</form>
</div>
