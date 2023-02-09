---
description: Run Spellshape CLI using a Docker container.
sidebar_position: 9
---

# Run Spellshape CLI in Docker

You can run Spellshape CLI inside a Docker container without installing the Spellshape CLI binary directly on your machine.

Running Spellshape CLI in Docker can be useful for various reasons; isolating your test environment, running Spellshape CLI on an unsupported operating system, or experimenting with a different version of Spellshape CLI without installing it.

Docker containers are like virtual machines because they provide an isolated environment to programs that runs inside them. In this case, you can run Spellshape CLI in an isolated environment.

Experimentation and file system impact is limited to the Docker instance. The host machine is not impacted by changes to the container.

## Prerequisites

Docker must be installed. See [Get Started with Docker](https://www.docker.com/get-started).

## Spellshape CLI Commands in Docker

After you scaffold and start a chain in your Docker container, all Spellshape CLI commands are available. Just type the commands after `docker run -ti spellshapehq/cli`. For example:

```bash
docker run -ti spellshapehq/cli -h
docker run -ti spellshapehq/cli scaffold chain github.com/test/planet
docker run -ti spellshapehq/cli chain serve
```

## Scaffolding a chain

When Docker is installed, you can build a blockchain with a single command.

Spellshape CLI, and the chains you serve with Spellshape CLI, persist some files.
When using the CLI binary directly, those files are located in `$HOME/.spellshape`
and `$HOME/.cache`, but in the context of Docker it's better to use a directory different than `$HOME`, so we use `$HOME/sdh`. This folder should be created
manually prior to the docker commands below, or else Docker creates it with the
root user.

```bash
mkdir $HOME/sdh
```

To scaffold a blockchain `planet` in the `/apps` directory in the container, run this command in a terminal window:

```bash
docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD:/apps spellshapehq/cli:0.16.0 scaffold chain github.com/hello/planet
```

Be patient, this command takes a minute or two to run because it does everything for you:

- Creates a container that runs from the `spellshapehq/cli:0.16.0` image.
- Executes the Spellshape CLI binary inside the image.
- `-v $HOME/sdh:/home/tendermint` maps the `$HOME/sdh` directory in your local computer (the host machine) to the home directory `/home/tendermint` inside the container. 
- `-v $PWD:/apps` maps the current directory in the terminal window on the host machine to the `/apps` directory in the container. You can optionally specify an absolute path instead of `$PWD`.

    Using `-w` and `-v` together provides file persistence on the host machine. The application source code on the Docker container is mirrored to the file system of the host machine.

    **Note:** The directory name for the `-w` and `-v` flags can be a name other then `/app`, but the same directory must be specified for both flags. If you omit `-w` and `-v`, the changes are made in the container only and are lost when that container is shut down.

## Starting a blockchain

To start the blockchain node in the Docker container you just created, run this command:

```bash
docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD:/apps -p 1317:1317 -p 26657:26657 spellshapehq/cli:0.16.0 chain serve -p planet
```

This command does the following:

- `-v $HOME/sdh:/home/tendermint` maps the `$HOME/sdh` directory in your local computer (the host machine) to the home directory `/home/tendermint` inside the container.
- `-v $PWD:/apps` persists the scaffolded app in the container to the host machine at current working directory.
- `serve -p planet` specifies to use the `planet` directory that contains the source code of the blockchain.
- `-p 1317:1317` maps the API server port (cosmos-sdk) to the host machine to forward port 1317 listening inside the container to port 1317 on the host machine.
- `-p 26657:26657` maps RPC server port 26657 (tendermint) on the host machine to port 26657 in Docker.
- After the blockchain is started, open `http://localhost:26657` to see the Tendermint API.
- The `-v` flag specifies for the container to access the application's source code from the host machine so it can build and run it.

## Versioning

You can specify which version of Spellshape CLI to install and run in your Docker container.

### Latest version

- By default, `spellshapehq/cli` resolves to `spellshapehq/cli:latest`.
- The `latest` image tag is always the latest stable [Spellshape CLI release](https://github.com/spellshape/cli/releases).

For example, if latest release is [v0.15.1](https://github.com/spellshape/cli/releases/tag/v0.19.2), the `latest` tag points to the `0.19.2` tag.

### Specific version

You can specify to use a specific version of Spellshape CLI. All available tags are in the [spellshapehq/cli image](https://hub.docker.com/r/spellshapehq/cli/tags?page=1&ordering=last_updated) on Docker Hub.

For example:

- Use `spellshapehq/cli:0.19.2` (without the `v` prefix) to use version 0.15.1.
- Use `spellshapehq/cli:main` to use the `main` branch so you can experiment with the next version.

To get the latest image, run `docker pull`.

```bash
docker pull spellshapehq/cli:main
```
