---
sidebar_position: 1
description: Steps to install Spellshape CLI on your local computer.
---

# Install Spellshape CLI

You can run [Spellshape CLI](https://github.com/spellshape/cli) in a web-based Gitpod IDE or you can install Spellshape CLI on your
local computer.

## Prerequisites

Be sure you have met the prerequisites before you install and use Spellshape CLI.

### Operating systems

Spellshape CLI is supported for the following operating systems:

- GNU/Linux
- macOS
- Windows Subsystem for Linux (WSL)

### Go

Spellshape CLI is written in the Go programming language. To use Spellshape CLI on a local system:

- Install [Go](https://golang.org/doc/install) (**version 1.19** or higher)
- Ensure the Go environment variables are [set properly](https://golang.org/doc/gopath_code#GOPATH) on your system

## Verify your Spellshape CLI version

To verify the version of Spellshape CLI you have installed, run the following command:

```bash
spellshape version
```

## Installing Spellshape CLI

To install the latest version of the `spellshape` binary use the following command.

```bash
curl https://get.spellshape.com/spellshape/cli@nightly! | bash
```

This command invokes `curl` to download the installation script and pipes the output to `bash` to perform the
installation.  The `spellshape` binary is installed in `/usr/local/bin`.

To learn more or customize the installation process, see the [installer docs](https://github.com/spellshape/installer) on
GitHub.

### Write permission

Spellshape CLI installation requires write permission to the `/usr/local/bin/` directory. If the installation fails because
you do not have write permission to `/usr/local/bin/`, run the following command:

```bash
curl https://get.spellshape.com/spellshape/cli@nightly | bash
```

Then run this command to move the `spellshape` executable to `/usr/local/bin/`:

```bash
sudo mv spellshape /usr/local/bin/
```

On some machines, a permissions error occurs:

```bash
mv: rename ./spellshape to /usr/local/bin/spellshape: Permission denied
============
Error: mv failed
```

In this case, use sudo before `curl` and before `bash`:

```bash
sudo curl https://get.spellshape.com/spellshape/cli@nightly | bash
```

## Upgrading your Spellshape CLI installation

Before you install a new version of Spellshape CLI, remove all existing Spellshape CLI installations.

To remove the current Spellshape CLI installation:

1. On your terminal window, press `Ctrl+C` to stop the chain that you started with `spellshape chain serve`.
2. Remove the Spellshape CLI binary with `rm $(which spellshape)`.
   Depending on your user permissions, run the command with or without `sudo`.
3. Repeat this step until all `spellshape` installations are removed from your system.

After all existing Spellshape CLI installations are removed, follow the  [Installing Spellshape CLI](#installing-spellshape-cli)
instructions.

For details on version features and changes, see
the [changelog.md](https://github.com/spellshape/cli/blob/main/changelog.md)
in the repo.

## Build from source

To experiment with the source code, you can build from source:

```bash
git clone https://github.com/spellshape/cli --depth=1
cd cli && make install
```

## Summary

- Verify the prerequisites.
- To set up a local development environment, install Spellshape CLI locally on your computer.
- Install Spellshape CLI by fetching the binary using cURL or by building from source.
- The latest version is installed by default. You can install previous versions of the precompiled `spellshape` binary.
- Stop the chain and remove existing versions before installing a new version.
