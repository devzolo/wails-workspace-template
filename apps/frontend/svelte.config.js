import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/kit/vite';
import path from 'path';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: [vitePreprocess({})],
	kit: {
		adapter: adapter({
      pages: '../desktop/frontend/build',
      assets: '../desktop/frontend/build',
      fallback: 'index.html'
    }),
    alias: {
			'~': path.resolve('./src'),
		},
	}
};

export default config;
