# docker-build-workflow [![docker](https://github.com/int128/docker-build-workflow/actions/workflows/docker.yaml/badge.svg)](https://github.com/int128/docker-build-workflow/actions/workflows/docker.yaml)

This is an example of workflow to build a Docker image in GitHub Actions.

See [`docker.yaml`](.github/workflows/docker.yaml).

## Features

- Build and push into GitHub Container Registry
- [Build a multi-architecture image in parallel](https://github.com/int128/docker-manifest-create-action) (e.g., amd64 or arm64)
- BuildKit cache support
- [Effective cache strategy](https://github.com/int128/docker-build-cache-config-action)
