import { SlideshowsService } from '$lib/api-client';

export const ssr = false;
export const prerender = true;

export async function load({ params }) {
	const slideshows = await SlideshowsService.getAllSlideshows();

	return slideshows;
}
