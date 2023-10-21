import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	server: {
    fs: {
      // To allow serving files from the frontend project root.
      //
      // allow: ['.'],
    },
  },
	plugins: [sveltekit()],
	resolve: {
    alias: {
      '~': './src', 
    },
  },
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
