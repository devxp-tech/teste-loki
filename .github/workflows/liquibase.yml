name: liquibase
on:
  workflow_dispatch:

  push:
    branches:
      - main
    paths:
      - "liquibase/**"

jobs:
  vars:
    name: Setup Environment 🛠️
    runs-on: ubuntu-latest
    outputs:
      tag: ${{ steps.vars.outputs.tag }}
    steps:
      - id: vars
        # run: echo "::set-output name=tag::$(git show -s --format="%h")"
        run: |
          if [[ ! "$GITHUB_REF" =~ refs/tags ]]; then
            echo "::set-output name=tag::$(echo $GITHUB_SHA | cut -c 1-8)"
          else
            echo "::set-output name=tag::$(echo $GITHUB_REF | cut -d / -f 3)"
          fi
      - run: echo ${{ steps.vars.outputs.tag }}

  test:
    name: Test Docker Compose 🐳
    runs-on: ubuntu-latest
    steps:
      # need checkout before using compose-action
      - uses: actions/checkout@v3
      - uses: cloudposse/github-action-docker-compose-test-run@main
        with:
          file: ./docker-compose.yml
          service: app

  build-and-push:
    name: Build and Push 📦
    runs-on: ubuntu-latest
    needs:
      - vars
      - test
    steps:
      - name: Checkout the repo 🛎️
        uses: actions/checkout@v2

      - name: Set rev.txt 🛠️
        run: git show -s --format="%ai %H %s %aN" HEAD > rev.txt

      - name: Set up QEMU 🛠️
        uses: docker/setup-qemu-action@v2

      - name: Set up Docker Buildx 🛠️
        uses: docker/setup-buildx-action@v2
        with:
          DOCKER_CONFIG: liquibase/Dockerfile

      - name: Login Docker Registry 🪪
        uses: docker/login-action@v2
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ github.token }}

      - name: Build and Push Docker 🚀
        uses: docker/build-push-action@v3
        with:
          push: true
          context: liquibase
          secrets: |
            GIT_AUTH_TOKEN=${{ github.token }}
          tags: ghcr.io/devxp-tech/teste-loki-migration:${{ needs.vars.outputs.tag }},ghcr.io/devxp-tech/teste-loki-migration:latest