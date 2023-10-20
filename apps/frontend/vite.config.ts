import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';
import path from 'path';

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
      '~': path.resolve(__dirname, './src'), 
    },
  },
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
