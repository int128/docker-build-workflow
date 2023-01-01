# docker-build-workflow [![test](https://github.com/int128/docker-build-workflow/actions/workflows/test.yaml/badge.svg)](https://github.com/int128/docker-build-workflow/actions/workflows/test.yaml)

This is a set of reusable workflows to build a container image in GitHub Actions.

## Features

- Build and push an image into GitHub Container Registry
- [Build a multi-architecture image in parallel](https://github.com/int128/docker-manifest-create-action) (e.g., amd64 or arm64)
- BuildKit cache support with [effective cache strategy](https://github.com/int128/docker-build-cache-config-action)

## Example

See [test workflow](.github/workflows/test.yaml).
