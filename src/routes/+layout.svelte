<script>
  import { page } from "$app/state";
  import "../app.css";

  let { children } = $props();

  const navItems = [
    {
      name: "Accounts",
      url: "/",
    },
    {
      name: "Logs",
      url: "/logs",
    },
    {
      name: "Settings",
      url: "/settings",
    },
  ];
</script>

<nav class="bg-primary-200 w-full flex px-10 justify-between">
  {#each navItems as item}
    <a
      href={item.url}
      class={`relative px-2 py-2.5 text-white transition-colors duration-200
             after:content-[''] after:absolute after:left-0 after:bottom-0
             after:h-[3px] after:w-full after:bg-blue-500
             after:opacity-0 after:transition-opacity after:duration-200
             hover:after:opacity-100 w-fit after:rounded 
             ${page.url.pathname === item.url ? "after:opacity-100" : ""}`}
    >
      {item.name}
    </a>
  {/each}
</nav>

<main class="text-secondary-100 px-10 py-5">
  {@render children()}
</main>
