<script lang="ts">
	import type { service } from '@/wailsjs/go/models';
	import * as Table from '$lib/components/ui/table/index.js';
	import * as Popover from '$lib/components/ui/popover/index.js';
	import { Button, buttonVariants } from '$lib/components/ui/button/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import {
		HandleAddAccount,
		HandleCheck,
		HandleDeleteAccount,
		HandleReadAccount
	} from '@/wailsjs/go/main/App';
	import { onMount } from 'svelte';
	import { nanoid } from 'nanoid';
	import { toast } from 'svelte-sonner';
	let getUsers = $state([] as service.Account[] | void);
	let user = $state({
		name: '',
		token: '',
		uid: ''
	});
	onMount(async () => {
		await refreshReadAccount();
	});
	async function refreshReadAccount() {
		try {
			getUsers = await HandleReadAccount();
		} catch (err) {
			console.error(err);
			toast.error('Error retrieving accounts', {
				description: `${err}`
			});
		}
	}
	async function manualCheckIn() {
		toast.promise(HandleCheck(), {
			loading: 'Getting your rewards...',
			success: (res) => {
				console.log(res);
				return 'Success';
			},
			error: (err) => {
				console.error(err);
				return `Error checking in: ${err}`;
			}
		});
	}
	async function addAccount(e: SubmitEvent) {
		e.preventDefault();
		const toastId = toast.loading('Adding account');
		try {
			await HandleAddAccount({
				Id: nanoid(),
				Name: user.name,
				Token: user.token,
				Uid: user.uid
			});
			toast.success('Account added', {
				id: toastId
			});

			((user.name = ''), (user.token = ''), (user.uid = ''));
			refreshReadAccount();
		} catch (err) {
			toast.error('Error creating user', {
				description: `${err}`
			});
		}
	}
	async function deleteAccount(id: string) {
		const toastId = toast.loading('Deleting account');
		try {
			await HandleDeleteAccount(id);
			toast.success('Account deleted', {
				id: toastId
			});
			refreshReadAccount();
		} catch (err) {
			toast.error('Error creating user', {
				description: `${err}`
			});
		}
	}
</script>

<div class="mb-5">
	<Button onclick={manualCheckIn} class="mb-5">Manual Check In</Button>
</div>

<div class="mb-10">
	<h1 class="mb-5 font-semibold">Add Account</h1>

	<form class="flex flex-col gap-2.5" onsubmit={addAccount}>
		<Input placeholder="Name (only for display)" bind:value={user.name} required />
		<Input placeholder="Token" name="token" bind:value={user.token} required />
		<Input placeholder="UID" name="uid" bind:value={user.uid} required />
		<Button class="mt-1.5 w-fit" size="sm" type="submit">Add</Button>
	</form>
</div>

<div>
	<h1 class="font-semibold">Accounts</h1>
	<Table.Root class="font-medium">
		<Table.Header>
			<Table.Row>
				<Table.Head>Name</Table.Head>
				<Table.Head>Token</Table.Head>
				<Table.Head>UID</Table.Head>
				<Table.Head>Action</Table.Head>
			</Table.Row>
		</Table.Header>
		<Table.Body>
			{#if (getUsers as service.Account[]).length === 0}
				<p class="mt-2.5">No Account</p>
			{:else}
				{#each getUsers as service.Account[] as user}
					<Table.Row>
						<Table.Cell>{user.Name}</Table.Cell>
						<Table.Cell>
							<Popover.Root>
								<Popover.Trigger class={buttonVariants({ variant: 'outline', size: 'sm' })}>
									View Token
								</Popover.Trigger>
								<Popover.Content class="flex h-max items-center gap-1.5">
									<p class="truncate text-xs" id="user-token">{user.Token}</p>
									<Button
										variant="link"
										size="sm"
										class="px-0 py-0 text-xs"
										onclick={() => {
											const token = document.getElementById('user-token');
											if (token) {
												navigator.clipboard.writeText(token.innerHTML);
												toast.info('Token copied');
											}
										}}>copy</Button
									>
								</Popover.Content>
							</Popover.Root>
						</Table.Cell>
						<Table.Cell>{user.Uid}</Table.Cell>
						<Table.Cell>
							<Button variant="destructive" size="sm" onclick={() => deleteAccount(user.Id)}
								>Delete</Button
							>
						</Table.Cell>
					</Table.Row>
				{/each}
			{/if}
		</Table.Body>
	</Table.Root>
</div>
