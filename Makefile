.PHONY: test test-coverage

test:
	cd server && env GOCACHE=/tmp/go-build-cache GOMODCACHE=/tmp/go-mod-cache go test ./... -v
	cd wa-engine && env PNPM_HOME=/tmp/pnpm-home PNPM_STORE_DIR=/tmp/pnpm-store pnpm test

test-coverage:
	cd server && env GOCACHE=/tmp/go-build-cache GOMODCACHE=/tmp/go-mod-cache go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out
	cd wa-engine && env PNPM_HOME=/tmp/pnpm-home PNPM_STORE_DIR=/tmp/pnpm-store pnpm run test:coverage
