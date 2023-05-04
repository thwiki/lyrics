import { sveltekit } from '@sveltejs/kit/vite';
import svelteSVG from 'vite-plugin-svelte-svg';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	plugins: [
		svelteSVG({
			svgoConfig: {
				plugins: [
					{
						name: 'removeViewBox',
						active: false
					},
					'removeDimensions'
				]
			},
			requireSuffix: true
		}),
		sveltekit()
	],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
