# sample-docker-build

This repository is a build-variation testbed for GitHub Actions.

## Workflows

1. `.github/workflows/docker-build.yml`
	- Baseline manual Docker build with selectable Dockerfile and context.
2. `.github/workflows/variation-core.yml`
	- Hybrid core matrix for raw Docker CLI and `docker/build-push-action` local-load builds.
3. `.github/workflows/variation-docker-advanced.yml`
	- Advanced Docker variants: base image families, multi-stage targets, multi-platform OCI export, image artifact roundtrip.
4. `.github/workflows/variation-nondocker.yml`
	- Non-Docker builds for Node, Python (PyPI-style package build), Ruby, and Go.
5. `.github/workflows/variation-pipeline-stages.yml`
	- Multi-stage pipeline orchestration: lint -> build -> smoke -> artifact verify.

## Docker fixture variants

Located in `docker/`:

1. `Dockerfile.raw`
2. `Dockerfile.build-action`
3. `Dockerfile.custom-frontend`
4. `Dockerfile.cd-pattern`
5. `Dockerfile.chained`
6. `Dockerfile.split`
7. `Dockerfile.alpine`
8. `Dockerfile.debian`
9. `Dockerfile.multistage`
10. `Dockerfile.distroless`
11. `Dockerfile.nonroot`
12. `Dockerfile.app-context`

## Non-Docker fixture variants

1. `samples/python` for package build (`python -m build`).
2. `samples/ruby` for bundle + rake build.
3. `samples/go` for `go test` + `go build`.

## Run strategy

All workflows use `workflow_dispatch` (manual trigger), local-only image build/test (no registry push), and green-only reference variants.