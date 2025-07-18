<script lang="ts">
	import * as Alert from '$lib/components/ui/alert/index.js';
	import CheckCircle2Icon from '@lucide/svelte/icons/check-circle-2';

	type Props = {
		variant: 'default' | 'destructive';
		isAlert: boolean;
	};

	let { variant, isAlert = $bindable() }: Props = $props();
</script>

<Alert.Root
	class="flex flex-col gap-2 pr-4 sm:flex-row sm:items-start sm:justify-between"
	{variant}
>
	<div class="flex items-start gap-3">
		<CheckCircle2Icon class="mt-0.5 size-4" />
		<div>
			<Alert.Title>Successfully checked in</Alert.Title>
		</div>
	</div>

	<button
		onclick={() => (isAlert = false)}
		class="text-primary text-sm font-medium transition hover:underline"
		aria-label="Close alert"
	>
		Close
	</button>
</Alert.Root>
