# Spellshape CLI

[Spellshape CLI](https://spellshape.com/cli) is the all-in-one platform to build,
launch, and maintain any crypto application on a sovereign and secured
blockchain. It is a developer-friendly interface to the [Cosmos
SDK](https://github.com/cosmos/cosmos-sdk), the world's most widely-used
blockchain application framework. Spellshape CLI generates boilerplate code for you,
so you can focus on writing business logic.

## Quick start

Open Spellshape CLI [in your web
browser](https://gitpod.io/#https://github.com/spellshape/cli/tree/v0.25.2) (or open
[nightly version](https://gitpod.io/#https://github.com/spellshape/cli)), or
[install the latest release](https://docs.spellshape.com/welcome/install).

To create and start a blockchain:

```bash
spellshape scaffold chain mars

cd mars

spellshape chain serve
```

## Documentation

To learn how to use Spellshape CLI, check out the [Spellshape CLI
docs](https://docs.spellshape.com). To learn more about how to build blockchain apps
with Spellshape CLI, see the [Spellshape CLI Developer
Tutorials](https://docs.spellshape.com/guide).

To install Spellshape CLI locally on GNU, Linux, or macOS, see [Install Spellshape
CLI](https://docs.spellshape.com/guide/install).

To learn more about building a JavaScript frontend for your Cosmos SDK
blockchain, see [spellshape/web](https://github.com/spellshape/web).

## Questions

For questions and support, join the official [Spellshape
Discord](https://discord.gg/spellshape) server. The issue list in this repo is
exclusively for bug reports and feature requests.

## Cosmos SDK compatibility

Blockchains created with Spellshape CLI use the [Cosmos
SDK](https://github.com/cosmos/cosmos-sdk) framework. To ensure the best
possible experience, use the version of Spellshape CLI that corresponds to the
version of Cosmos SDK that your blockchain is built with. Unless noted
otherwise, a row refers to a minor version and all associated patch versions.

| Spellshape CLI | Cosmos SDK  | IBC                  | Notes                                                         |
| -------------- | ----------- | -------------------- | ------------------------------------------------------------- |
| v0.26.0        | v0.46.7     | v6.1.0               |                                                               |
| v0.25.2        | v0.46.6     | v5.1.0               | Bump Tendermint version to v0.34.24                           |
| v0.25.1        | v0.46.3     | v5.0.0               | Includes  Dragonberry security fix                            |
| ~~v0.24.0~~    | ~~v0.46.0~~ | ~~v5.0.0~~           | This version is deprecated due to a security fix in `v0.25.0` |
| v0.23.0        | v0.45.5     | v3.0.1               |                                                               |
| v0.21.1        | v0.45.4     | v2.0.3               | Supports Cosmos SDK v0.46.0-alpha1 and above                  |
| v0.21.0        | v0.45.4     | v2.0.3               |                                                               |
| v0.20.0        | v0.45.3     | v2.0.3               |                                                               |
| v0.19          | v0.44       | v1.2.2               |                                                               |
| v0.18          | v0.44       | v1.2.2               | `spellshape chain serve` works with v0.44.x chains            |
| v0.17          | v0.42       | Same with Cosmos SDK |                                                               |

To upgrade your blockchain to the newer version of Cosmos SDK, see the
[Migration guide](https://docs.spellshape.com/migration).

## Plugin system

Spellshape CLI commands can be extended using plugins. A plugin is a program that
uses github.com/hashicorp/go-plugin to communicate with the spellshape binary.

#### Use a plugin

Plugins must be declared in the `config.yml` file, using the following syntax:

```yaml
plugins:
  // path can be a repository or a local path
  // the directory must contain go code under a main package.
  // For repositories you can specify a suffix @branch or @tag to target a
  // specific git reference.
  - path: github.com/org/repo/my-plugin
    // Additional parameters can be passed to the plugin
    with:
      key: value
```

Once declared, the next time the spellshape binary will be executed under this
configuration, it will fetch, build and run the plugin. As a result, more
commands should be available in the list of the spellshape commands.

`spellshape plugin` command allows to list the plugins and their status, and to
update a plugin if you need to get the latest version.

### Make a plugin

A plugin must implement `plugin.Interface`.

The easiest way to make a plugin is to use the `spellshape plugin scaffold` command.
For example:

```
$ cd /home/user/src
$ spellshape plugin scaffold github.com/foo/bar
```

It will create a folder `bar` under `/home/user/src` and generate predefined
`go.mod` and `main.go`. The code contains everything required to connect to the
spellshape binary via `hashicorp/go-plugin`. What need to be adapted is the
implementation of the `plugin.Interface` (`Commands` and `Execute` methods).

To test your plugin, you only need to declare it under a chain config, for
instance:

```yaml
plugins:
  - path: /home/user/src/bar
```

Then run `spellshape`, the plugin will compile and should be listed among the spellshape
commands. Each time `spellshape` is executed, the plugin is recompiled if the files
have changed since the last compilation. This allows fast and easy plugin
development, you only care about code and `spellshape` handles the compilation.

## Contributing

We welcome contributions from everyone. The `main` branch contains the
development version of the code. You can create a branch from `main` and
create a pull request, or maintain your own fork and submit a cross-repository
pull request.

Our [Spellshape CLI bounty program](https://docs.spellshape.com/bounty) provides
incentives for your participation and pays rewards. Track new, in-progress, and
completed bounties on the [Bounty
board](https://github.com/spellshape/cli/projects/5) in GitHub.

**Important** Before you start implementing a new Spellshape CLI feature, the first
step is to create an issue on GitHub that describes the proposed changes.

If you're not sure where to start, check out [contributing.md](contributing.md)
for our guidelines and policies for how we develop Spellshape CLI. Thank you to
everyone who has contributed to Spellshape CLI!

## Community

Spellshape CLI is a free and open source product maintained by
[Spellshape](https://spellshape.com). Here's where you can find us. Stay in touch.

* [Spellshape docs](https://docs.spellshape.com)
* [Spellshape on Twitter](https://twitter.com/spellshape)
