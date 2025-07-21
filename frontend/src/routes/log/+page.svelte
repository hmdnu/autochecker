<script lang="ts">
	import * as Table from '$lib/components/ui/table/index.js';
	import { HandleReadLog } from '@/wailsjs/go/main/App';
	import { onMount } from 'svelte';

	onMount(async () => {
		try {
			const logs = await HandleReadLog();
			console.log(logs);
		} catch (err) {
			console.error('error', err);
		}
	});
</script>

<Table.Root class="font-medium">
	<Table.Header>
		<Table.Row>
			<Table.Head>Account</Table.Head>
			<Table.Head>Game</Table.Head>
			<Table.Head>Status</Table.Head>
			<Table.Head>Date/Time</Table.Head>
		</Table.Row>
	</Table.Header>
	<Table.Body>
		<Table.Row>
			<Table.Cell>Nerots</Table.Cell>
			<Table.Cell>Zenless Zone Zero</Table.Cell>
			<Table.Cell>Done</Table.Cell>
			<Table.Cell>12 Jan 2025 / 00:00</Table.Cell>
		</Table.Row>
	</Table.Body>
</Table.Root>
