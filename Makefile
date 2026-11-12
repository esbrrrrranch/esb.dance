.PHONY: build dev format format-check test

build:
	pnpm build

dev:
	pnpm dev

format:
	pnpm format

format-check:
	pnpm format:check

test:
	pnpm test
