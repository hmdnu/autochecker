<script lang="ts">
	import * as Table from '$lib/components/ui/table/index.js';
	import * as Select from '$lib/components/ui/select/index.js';
	import { HandleReadLog } from '@/wailsjs/go/main/App';
	import type { service } from '@/wailsjs/go/models';
	import { onMount } from 'svelte';
	import { toast } from 'svelte-sonner';

	let logs = $state<Array<service.Response> | undefined>();

	onMount(async () => {
		const toastId = toast.loading('Retrieving logs...');
		try {
			logs = (await HandleReadLog()).reverse();
			toast.success('', {
				id: toastId,
				duration: 0
			});
		} catch (err) {
			toast.error('Error retrieving logs', {
				description: `${err}`
			});
		}
	});

	const filter = [
		{
			value: 'all',
			label: 'All'
		},
		{
			value: 'today',
			label: 'Today'
		}
	];

	let value = $state('');
	const triggerContent = $derived(filter.find((f) => f.value === value)?.label ?? 'Filter');

	function filterLogs(value: string) {
		switch (value) {
			case 'today':
				return logs?.filter((log) => {
					const today = new Date();
					const date = new Date(log.DateCheck.replace('WIB', 'GMT+0700'));
					return (
						date.getFullYear() === today.getFullYear() &&
						date.getMonth() === today.getMonth() &&
						date.getDate() === today.getDate()
					);
				});
			default:
				return logs;
		}
	}

	const filteredLogs = $derived.by(() => {
		return filterLogs(value);
	});
</script>

<Select.Root type="single" bind:value>
	<Select.Trigger class="w-[180px]">{triggerContent}</Select.Trigger>
	<Select.Content>
		<Select.Label>Filter</Select.Label>
		{#each filter as f}
			<Select.Item value={f.value} label={f.label}>{f.label}</Select.Item>
		{/each}
	</Select.Content>
</Select.Root>

<Table.Root class="mt-5 font-medium">
	<Table.Header>
		<Table.Row>
			<Table.Head>Account</Table.Head>
			<Table.Head>Game</Table.Head>
			<Table.Head>Status</Table.Head>
			<Table.Head>Date/Time</Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		{#if filteredLogs}
			{#each filteredLogs as log}
				{#each log.Game as game}
					<Table.Row>
						<Table.Cell>{log.AccountName}</Table.Cell>
						<Table.Cell>{game.name}</Table.Cell>
						<Table.Cell>{game.message}</Table.Cell>
						<Table.Cell>{log.DateCheck}</Table.Cell>
					</Table.Row>
				{/each}
			{/each}
		{/if}
	</Table.Body>
</Table.Root>
