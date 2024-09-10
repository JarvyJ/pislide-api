<script lang="ts">
	import type { BaseSlideShowOutput } from '$lib/api-client';
	type SlideshowSettings = {
		foldername: string;
		settings: BaseSlideShowOutput;
		newSlideshow?: boolean;
		modal?: boolean;
	};

	let {
		foldername,
		settings = $bindable(),
		newSlideshow = false,
		modal = $bindable()
	}: SlideshowSettings = $props();

	import { goto } from '$app/navigation';

	import ExpandSettings from '$lib/icons/chevron-down.svelte';
	import { SlideshowsService } from '$lib/api-client';

	let settingsExpanded = $state(false);
	if (newSlideshow) {
		settingsExpanded = true;
	}

	async function setActive() {
		await SlideshowsService.setActiveSlideshow({
			requestBody: { active_slideshow: foldername }
		});
		await goto('/', { invalidateAll: true });
	}

	async function updateSlideshow() {
		if (newSlideshow) {
			await SlideshowsService.createSlideshow({
				foldername: foldername,
				requestBody: settings
			});
			modal = false;
		} else {
			await SlideshowsService.updateSpecificSlideshow({
				foldername: foldername,
				requestBody: settings
			});
		}
		await goto('/', { invalidateAll: true });
	}
</script>

<div class="card">
	<header class="card-header">
		<p
			class="card-header-title"
			class:is-clickable={!newSlideshow}
			onclick={() => (settingsExpanded = !settingsExpanded)}
		>
			{foldername}
		</p>
		<button
			class="card-header-icon"
			aria-label="more options"
			onclick={() => (settingsExpanded = !settingsExpanded)}
		>
			<span class="icon">
				<ExpandSettings />
			</span>
		</button>
	</header>

	{#if settingsExpanded}
		<div class="card-content">
			<div class="content">
				{#if newSlideshow}
					<div class="field">
						<label for="Foldername" class="label">Foldername</label>
						<div class="control">
							<input class="input" type="text" id="Foldername" bind:value={foldername} />
						</div>
					</div>
				{/if}

				<div class="field">
					<label for="slide-duration" class="label">Slide Duration (seconds)</label>
					<div class="control">
						<input
							class="input"
							type="number"
							placeholder="4.0"
							min="0.1"
							step="0.1"
							required
							id="slide-duration"
							bind:value={settings.duration}
						/>
					</div>
				</div>

				<div class="field">
					<label for="transition-duration" class="label">Transition Duration (seconds)</label>
					<div class="control">
						<input
							class="input"
							type="number"
							placeholder="4.0"
							min="0"
							step="0.1"
							required
							id="transition-duration"
							bind:value={settings.transition_duration}
						/>
					</div>
				</div>

				<div class="field">
					<label for="sort" class="label">Sort</label>
					<div class="control">
						<div class="select">
							<select id="sort" bind:value={settings.sort}>
								<option value="filename">filename</option>
								<option value="natural">natural</option>
								<option value="random">random</option>
							</select>
						</div>
					</div>
				</div>

				<div class="field">
					<label for="sort" class="label">Text to display</label>
					<div class="control">
						<div class="select">
							<select id="sort" bind:value={settings.display}>
								<option value="filename">filename</option>
								<option value="caption">caption</option>
								<option value="none">none</option>
							</select>
						</div>
					</div>
				</div>
			</div>
		</div>

		<footer class="card-footer">
			{#if newSlideshow}
				<a class="card-footer-item" onclick={() => (modal = false)}>Cancel</a>
			{:else}
				<a class="card-footer-item" onclick={setActive}>Set Active</a>
			{/if}
			<a class="card-footer-item" onclick={updateSlideshow}>Save</a>
			<!-- <a class="card-footer-item">Delete</a> -->
		</footer>
	{/if}
</div>

<style>
	input:invalid {
		--bulma-input-h: var(--bulma-danger-h);
		--bulma-input-s: var(--bulma-danger-s);
		--bulma-input-l: var(--bulma-danger-l);
		--bulma-input-focus-h: var(--bulma-danger-h);
		--bulma-input-focus-s: var(--bulma-danger-s);
		--bulma-input-focus-l: var(--bulma-danger-l);
		--bulma-input-border-l: var(--bulma-danger-l);
	}

	.disabled {
		color: var(--bulma-text-weak);
		cursor: not-allowed;
	}
</style>
