<script lang="ts">
	import type { BaseSlideShowOutput } from '$lib/api-client';

	export let data;

	import SlideshowSettingsCard from '$lib/pages/SlideshowSettingsCard.svelte';

	let openNewSlideshowModal = false;
	let newSlideshowName = '';
	let newSlideshowSettings: BaseSlideShowOutput = {
		display: 'none',
		transition_duration: 0,
		duration: 0,
		sort: 'filename'
	};
</script>

<h2 class="subtitle is-4">Active Slideshow</h2>
{#if data.active_slideshow}
	<SlideshowSettingsCard
		foldername={data.active_slideshow}
		bind:settings={data.slideshows[data.active_slideshow]}
	/>
{:else}
	<p class="block">There is currently no active slideshow set</p>
{/if}

<h2 class="subtitle is-4">Slideshows</h2>

<button class="button is-primary mb-4" on:click={() => (openNewSlideshowModal = true)}
	>New Slideshow</button
>

<div class="modal" class:is-active={openNewSlideshowModal}>
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="modal-background" on:click={() => (openNewSlideshowModal = false)}></div>
	<div class="modal-card">
		<SlideshowSettingsCard
			foldername={newSlideshowName}
			settings={newSlideshowSettings}
			newSlideshow={true}
			bind:modal={openNewSlideshowModal}
		/>
	</div>
</div>

{#each Object.entries(data.slideshows) as [foldername, slideshowMeta], index (foldername)}
	<SlideshowSettingsCard {foldername} bind:settings={slideshowMeta} />
{/each}
