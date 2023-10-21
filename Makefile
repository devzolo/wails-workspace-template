install:
	corepack enable && pnpm i

dev:
	cd apps/desktop && wails dev

build:
	cd apps/desktop && wails build

ui:
	cd apps/frontend && npx shadcn-svelte@latest add
