<script lang="ts">
	import { page } from '$app/state';
	import { Toaster } from '@/components/ui/sonner';
	import '../app.css';

	let { children } = $props();

	const navs = [
		{
			name: 'Account',
			link: '/'
		},
		{
			name: 'Log',
			link: '/log'
		},
		{
			name: 'Setting',
			link: '/setting'
		}
	];
</script>

<nav class="flex items-center justify-center gap-5 p-5">
	{#each navs as nav}
		<a
			href={nav.link}
			class={`font-semibold hover:underline ${page.url.pathname === nav.link ? 'text-blue-500' : ''}`}
			>{nav.name}</a
		>
	{/each}
</nav>

<main class="p-5">
	<Toaster richColors closeButton />

	{@render children()}
</main>
