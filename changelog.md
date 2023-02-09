## [`v0.26.0`](https://github.com/spellshape/cli/releases/tag/v0.26.0)

### Features

- [#3238](https://github.com/spellshape/cli/pull/3238) Add `Sharedhost` plugin option
- [#3214](https://github.com/spellshape/cli/pull/3214) Global plugins config.
- [#3142](https://github.com/spellshape/cli/pull/3142) Add `spellshape network request param-change` command.
- [#3181](https://github.com/spellshape/cli/pull/3181) Addition of `add` and `remove` commands for `plugins`
- [#3184](https://github.com/spellshape/cli/pull/3184) Separate `plugins.yml` config file.
- [#3038](https://github.com/spellshape/cli/pull/3038) Addition of Plugin Hooks in Plugin System
- [#3056](https://github.com/spellshape/cli/pull/3056) Add `--genesis-config` flag option to `spellshape network chain publish`
- [#2892](https://github.com/spellshape/cli/pull/2982/) Add `spellshape scaffold react` command.
- [#2892](https://github.com/spellshape/cli/pull/2982/) Add `spellshape generate composables` command.
- [#2892](https://github.com/spellshape/cli/pull/2982/) Add `spellshape generate hooks` command.
- [#2955](https://github.com/spellshape/cli/pull/2955/) Add `spellshape network request add-account` command.
- [#2877](https://github.com/spellshape/cli/pull/2877) Plugin system
- [#3060](https://github.com/spellshape/cli/pull/3060) Plugin system flag support
- [#3105](https://github.com/spellshape/cli/pull/3105) Addition of `spellshape plugin describe <path>` command
- [#2995](https://github.com/spellshape/cli/pull/2995/) Add `spellshape network request remove-validator` command.
- [#2999](https://github.com/spellshape/cli/pull/2999/) Add `spellshape network request remove-account` command.
- [#2458](https://github.com/spellshape/cli/issues/2458) New `chain serve` command UI.
- [#2992](https://github.com/spellshape/cli/issues/2992) Add `spellshape chain debug` command.
- [#2736](https://github.com/spellshape/cli/issues/2736) Add `--skip-git` flag to skip git repository initialization.

### Changes

- [#3369](https://github.com/spellshape/cli/pull/3369) Update `ibc-go` to `v6.1.0`.
- [#3306](https://github.com/spellshape/cli/pull/3306) Move network command into a plugin
- [#3305](https://github.com/spellshape/cli/pull/3305) Bump Cosmos SDK version to `v0.46.7`.
- [#3068](https://github.com/spellshape/cli/pull/3068) Add configs to generated TS code for working with JS projects
- [#3071](https://github.com/spellshape/cli/pull/3071) Refactor `spellshape/templates` package.
- [#2892](https://github.com/spellshape/cli/pull/2982/) `spellshape scaffold vue` and `spellshape scaffold react` use v0.4.2 templates
- [#2892](https://github.com/spellshape/cli/pull/2982/) `removeSigner()` method added to generated `ts-client`
- [#3035](https://github.com/spellshape/cli/pull/3035) Bump Cosmos SDK to `v0.46.4`.
- [#3037](https://github.com/spellshape/cli/pull/3037) Bump `ibc-go` to `v5.0.1`.
- [#2957](https://github.com/spellshape/cli/pull/2957) Change generate commands to print the path to the generated code.
- [#2981](https://github.com/spellshape/cli/issues/2981) Change CLI to also search chain binary in Go binary path.
- [#2958](https://github.com/spellshape/cli/pull/2958) Support absolute paths for client code generation config paths.
- [#2993](https://github.com/spellshape/cli/pull/2993) Hide `spellshape scaffold band` command and deprecate functionality.
- [#2986](https://github.com/spellshape/cli/issues/2986) Remove `--proto-all-modules` flag because it is now the default behaviour.
- [#2986](https://github.com/spellshape/cli/issues/2986) Remove automatic Vue code scaffolding from `scaffold chain` command.
- [#2986](https://github.com/spellshape/cli/issues/2986) Add `--generate-clients` to `chain serve` command for optional client code (re)generation.
- [#2998](https://github.com/spellshape/cli/pull/2998) Hide `spellshape generate dart` command and remove functionality.
- [#2991](https://github.com/spellshape/cli/pull/2991) Hide `spellshape scaffold flutter` command and remove functionality.
- [#2944](https://github.com/spellshape/cli/pull/2944) Add a new event "update" status option to `pkg/cliui`.
- [#3030](https://github.com/spellshape/cli/issues/3030) Remove colon syntax from module scaffolding `--dep` flag.
- [#3025](https://github.com/spellshape/cli/issues/3025) Improve config version error handling.
- [#3084](https://github.com/spellshape/cli/pull/3084) Add Spellshape Chain documentation.
- [#3109](https://github.com/spellshape/cli/pull/3109) Refactor scaffolding for proto files to not rely on placeholders.
- [#3106](https://github.com/spellshape/cli/pull/3106) Add zoom image plugin.
- [#3194](https://github.com/spellshape/cli/issues/3194) Move config validators check to validate only when required.
- [#3183](https://github.com/spellshape/cli/pull/3183/) Make config optional for init phase.
- [#3224](https://github.com/spellshape/cli/pull/3224) Remove `grpc_*` prefix from query files in scaffolded chains
- [#3229](https://github.com/spellshape/cli/pull/3229) Rename `campaign` to `project` in spellshape network set of commands
- [#3122](https://github.com/spellshape/cli/issues/3122) Change `generate ts-client` to ignore the cache by default.
- [#3244](https://github.com/spellshape/cli/pull/3244) Update `actions.yml` for resolving deprecation message
- [#3337](https://github.com/spellshape/cli/pull/3337) Remove `pkg/openapiconsole` import from scaffold template.
- [#3337](https://github.com/spellshape/cli/pull/3337) Register`nodeservice` gRPC in `app.go` template.

### Breaking Changes

- [#3033](https://github.com/spellshape/cli/pull/3033) Remove Cosmos SDK Launchpad version support.

### Fixes

- [#3114](https://github.com/spellshape/cli/pull/3114) Fix out of gas issue when approving many requests
- [#3068](https://github.com/spellshape/cli/pull/3068) Fix REST codegen method casing bug
- [#3031](https://github.com/spellshape/cli/pull/3031) Move keeper hooks to after all keepers initialized in `app.go` template.
- [#3098](https://github.com/spellshape/cli/issues/3098) Fix config upgrade issue that left config empty on error. 
- [#3129](https://github.com/spellshape/cli/issues/3129) Remove redundant `keyring-backend` config option.
- [#3187](https://github.com/spellshape/cli/issues/3187) Change prompt text to fit within 80 characters width.
- [#3203](https://github.com/spellshape/cli/issues/3203) Fix relayer to work with multiple paths.
- [#3320](https://github.com/spellshape/cli/pull/3320) Allow `id` and `creator` as names when scaffolding a type.
- [#3327](https://github.com/spellshape/cli/issues/3327) Scaffolding messages with same name leads to aliasing.
- [#3383](https://github.com/spellshape/cli/pull/3383) State error and info are now displayed when using serve UI.
- [#3379](https://github.com/spellshape/cli/issues/3379) Fix `spellshape docs` issue by disabling mouse support.

## [`v0.25.2`](https://github.com/spellshape/cli/releases/tag/v0.25.1)

### Changes

- [#3145](https://github.com/spellshape/cli/pull/3145) Security fix upgrading Cosmos SDK to `v0.46.6`
# Changelog

## [`v0.25.1`](https://github.com/spellshape/cli/releases/tag/v0.25.1)

### Changes

- [#2968](https://github.com/spellshape/cli/pull/2968) Dragonberry security fix upgrading Cosmos SDK to `v0.46.3`

## [`v0.25.0`](https://github.com/spellshape/cli/releases/tag/v0.25.0)

### Features

- Add `pkg/cosmostxcollector` package with support to query and save TXs and events.
- Add `spellshape network coordinator` command set.
- Add `spellshape network validator` command set.
- Deprecate `cosmoscmd` pkg and add cmd templates for scaffolding.
- Add generated TS client test support to integration tests.

### Changes

- Updated `pkg/cosmosanalysis` to discover the list of app modules when defined in variables or functions.
- Improve genesis parser for `network` commands
- Integration tests build their own spellshape binary.
- Updated `pkg/cosmosanalysis` to discover the list of app modules when defined in variables.
- Switch to broadcast mode sync in `cosmosclient`
- Updated `nodetime`: `ts-proto` to `v1.123.0`, `protobufjs` to `v7.1.1`, `swagger-typescript-api` to `v9.2.0`
- Switched codegen client to use `axios` instead of `fetch`
- Added `useKeplr()` and `useSigner()` methods to TS client. Allowed query-only instantiation.
- `nodetime` built with `vercel/pkg@5.6.0`
- Change CLI to use an events bus to print to stdout.
- Move generated proto files to `proto/{appname}/{module}`
- Update `pkg/cosmosanalysis` to detect when proto RPC services are using pagination.
- Add `--peer-address` flag to `network chain join` command.
- Change nightly tag format
- Add cosmos-sdk version in `version` command
- [#2935](https://github.com/spellshape/cli/pull/2935) Update `gobuffalo/plush` templating tool to `v4`

### Fixes

- Fix ICA controller wiring.
- Change vuex generation to use a default TS client path.
- Fix cli action org in templates.
- Seal the capability keeper in the `app.go` template.
- Change faucet to allow CORS preflight requests.
- Fix config file migration to void leaving end of file content chunks.
- Change session print loop to block until all events are handled.
- Handle "No records were found in keyring" message when checking keys.
- [#2941](https://github.com/spellshape/cli/issues/2941) Fix session to use the same spinner referece.
- [#2922](https://github.com/spellshape/cli/pull/2922) Network commands check for latest config version before building the chain binary.

## [`v0.24.1`](https://github.com/spellshape/cli/releases/tag/v0.24.1)

### Features

- Upgraded Cosmos SDK to `v0.46.2`.

## [`v0.24.0`](https://github.com/spellshape/cli/releases/tag/v0.24.0)

### Features

- Upgraded Cosmos SDK to `v0.46.0` and IBC to `v5` in CLI and scaffolding templates
- Change chain init to check that no gentx are present in the initial genesis
- Add `network rewards release` command
- Add "make mocks" target to Makefile
- Add `--skip-proto` flag to `build`, `init` and `serve` commands to build the chain without building proto files
- Add `node query tx` command to query a transaction in any chain.
- Add `node query bank` command to query an account's bank balance in any chain.
- Add `node tx bank send` command to send funds from one account to another in any chain.
- Add migration system for the config file to allow config versioning
- Add `node tx bank send` command to send funds from one account to another in any chain.
- Implement `network profile` command
- Add `generate ts-client` command to generate a stand-alone modular TypeScript client.

### Changes

- Add changelog merge strategy in `.gitattributes` to avoid conflicts.
- Refactor `templates/app` to remove `monitoringp` module from the default template
- Updated keyring dependency to match Cosmos SDK
- Speed up the integration tests
- Refactor spellshape network and fix genesis generation bug
- Make Go dependency verification optional during build by adding the `--check-dependencies` flag
  so Spellshape CLI can work in a Go workspace context.
- Temporary SPN address change for nightly
- Rename `simapp.go.plush` simulation file template to `helpers.go.plush`
- Remove campaign creation from the `network chain publish` command
- Optimized JavaScript generator to use a single typescript API generator binary
- Improve documentation and add support for protocol buffers and Go modules syntax
- Add inline documentation for CLI commands
- Change `cmd/account` to skip passphrase prompt when importing from mnemonic
- Add nodejs version in the output of spellshape version
- Removed `handler.go` from scaffolded module template
- Migrated to `cosmossdk.io` packages for and `math`
- Vuex stores from the `generate vuex` command use the new TypeScript client
- Upgraded frontend Vue template to v0.3.10

### Fixes

- Improved error handling for crypto wrapper functions
- Fix `pkg/cosmosclient` to call the faucet prior to creating the tx.
- Test and refactor `pkg/comosclient`.
- Change templates to add missing call to `RegisterMsgServer` in the default module's template to match what's specified
  in the docs
- Fix cosmoscmd appID parameter value to sign a transaction correctly
- Fix `scaffold query` command to use `GetClientQueryContext` instead of `GetClientTxContext`
- Fix flaky integration tests issue that failed with "text file busy"
- Fix default chain ID for publish
- Replace `os.Rename` with `xos.Rename`
- Fix CLI reference generation to add `spellshape completion` documentation
- Remove usage of deprecated `io/ioutil` package

## [`v0.23.0`](https://github.com/spellshape/cli/releases/tag/v0.23.0)

### Features

- Apps can now use generics

### Fixes

- Fix `pkg/cosmosanalysis` to support apps with generics
- Remove `spellshape-hq/cli` from dependency list in scaffolded chains

### Changes

- Change `pkg/cosmosgen` to allow importing IBC proto files
- Improve docs for Docker related commands
- Improve and fix documentation issues in developer tutorials
- Add migration docs for v0.22.2
- Improve `go mod download` error report in `pkg/cosmosgen`

## [`v0.22.2`](https://github.com/spellshape/cli/releases/tag/v0.22.2)

### Features

- Enable Darwin ARM 64 target for chain binary releases in CI templates

### Changes

- Rename `spellshape-hq` to `spellshape`

## [`v0.22.1`](https://github.com/spellshape/cli/releases/tag/v0.22.1)

### Fixes

- Fix IBC module scaffolding interface in templates

## [`v0.22.0`](https://github.com/spellshape/cli/releases/tag/v0.22.0)

### Features

- Optimized the build system. The `chain serve`, `chain build`, `chain generate` commands and other variants are way
  faster now
- Upgraded CLI and templates to use IBC v3

### Fixes

- Add a fix in code generation to avoid user's NodeJS configs to break TS client generation routine

## [`v0.21.2`](https://github.com/spellshape/cli/releases/tag/v0.21.2)

### Fixes

- Set min. gas to zero when running `chain` command set

## [`v0.21.1`](https://github.com/spellshape/cli/releases/tag/v0.21.1)

### Features

- Add compatibility to run chains built with Cosmos-SDK `v0.46.0-alpha1` and above
- Scaffold chains now will have `auth` module enabled by default

### Fixes

- Fixed shell completion generation
- Make sure proto package names are valid when using simple app names

## [`v0.21.0`](https://github.com/spellshape/cli/releases/tag/v0.21.0)

### Features

- Support simple app names when scaffolding chains. e.g.: `spellshape scaffold chain mars`
- Ask confirmation when scaffolding over changes that are not committed yet

## [`v0.20.4`](https://github.com/spellshape/cli/releases/tag/v0.20.4)

### Fixes

- Use `protoc` binary compiled in an older version of macOS AMD64 for backwards compatibility in code generation

## [`v0.20.3`](https://github.com/spellshape/cli/releases/tag/v0.20.3)

### Fixes

- Use the latest version of CLI in templates to fix Linux ARM support _(It's now possible to develop chains in Linux ARM
  machines and since the chain depends on the CLI in its `go.mod`, it needs to use the latest version that support ARM
  targets)_

## [`v0.20.2`](https://github.com/spellshape/cli/releases/tag/v0.20.2)

### Fixes

- Use `unsafe-reset-all` cmd under `tendermint` cmd for chains that use `=> v0.45.3` version of Cosmos SDK

## [`v0.20.1`](https://github.com/spellshape/cli/releases/tag/v0.20.1)

### Features

- Release the CLI with Linux ARM and native M1 binaries

## [`v0.20.0`](https://github.com/spellshape/cli/releases/tag/v0.20.0)

Our new name is **Spellshape CLI**!

**IMPORTANT!** This upgrade renames `starport` command to `spellshape`. From now on, use `spellshape` command to access the CLI.

### Features

- Upgraded Cosmos SDK version to `v0.45.2`
- Added support for in memory backend in `pkg/cosmosclient` package
- Improved our tutorials and documentation

## [`v0.19.5`](https://github.com/spellshape/cli/pull/2158/commits)

### Features

- Enable client code and Vuex code generation for query only modules as well.
- Upgraded the Vue template to `v0.3.5`.

### Fixes

- Fixed snake case in code generation.
- Fixed plugin installations for Go =>v1.18.

### Changes

- Dropped transpilation of TS to JS. Code generation now only produces TS files.

## `v0.19.4`

### Features

- Upgraded Vue template to `v0.3.0`.

## `v0.19.3`

### Features

- Upgraded Flutter template to `v2.0.3`

## [`v0.19.2`](https://github.com/spellshape/cli/milestone/14)

### Fixes

- Fixed race condition during faucet transfer
- Fixed account sequence mismatch issue on faucet and relayer
- Fixed templates for IBC code scaffolding

### Features

- Upgraded blockchain templates to use IBC v2.0.2

### Breaking Changes

- Deprecated the Starport Modules [tendermint/spm](https://github.com/tendermint/spm) repo and moved the contents to the
  Spellshape CLI repo [`spellshape/pkg/`](https://github.com/spellshape/cli/tree/main/spellshape/pkg/)
  in [PR 1971](https://github.com/spellshape/cli/pull/1971/files)

  Updates are required if your chain uses these packages:

  - `spm/ibckeeper` is now `pkg/cosmosibckeeper`
  - `spm/cosmoscmd` is now `pkg/cosmoscmd`
  - `spm/openapiconsole` is now `pkg/openapiconsole`
  - `testutil/sample` is now `cosmostestutil/sample`

- Updated the faucet HTTP API schema. See API changes
  in [fix: improve faucet reliability #1974](https://github.com/spellshape/cli/pull/1974/files#diff-0e157f4f60d6fbd95e695764df176c8978d85f1df61475fbfa30edef62fe35cd)

## `v0.19.1`

### Fixes

- Enabled the `scaffold flutter` command

## `v0.19.0`

### Features

- `starport scaffold` commands support `ints`, `uints`, `strings`, `coin`, `coins` as field types (#1579)
- Added simulation testing with `simapp` to the default template (#1731)
- Added `starport generate dart` to generate a Dart client from protocol buffer files
- Added `starport scaffold flutter` to scaffold a Flutter mobile app template
- Parameters can be specified with a new `--params` flag when scaffolding modules (#1716)
- Simulations can be run with `starport chain simulate`
- Set `cointype` for accounts in `config.yml` (#1663)

### Fixes

- Allow using a `creator` field when scaffolding a model with a `--no-message` flag (#1730)
- Improved error handling when generating code (#1907)
- Ensure account has funds after faucet transfer when using `cosmosclient` (#1846)
- Move from `io/ioutil` to `io` and `os` package (refactoring) (#1746)

## `v0.18.0`

### Breaking Changes

- Starport v0.18 comes with Cosmos SDK v0.44 that introduced changes that are not compatible with chains that were
  scaffolded with Starport versions lower than v0.18. After upgrading from Starport v0.17.3 to Starport v0.18, you must
  update the default blockchain template to use blockchains that were scaffolded with earlier versions.
  See [Migration](https://docs.spellshape.com/migration).

### Features

- Scaffold commands allow using previously scaffolded types as fields
- Added `--signer` flag to `message`, `list`, `map`, and `single` scaffolding to allow customizing the name of the
  signer of the message
- Added `--index` flag to `scaffold map` to provide a custom list of indices
- Added `scaffold type` to scaffold a protocol buffer definition of a type
- Automatically check for new Starport versions
- Added `starport tools completions` to generate CLI completions
- Added `starport account` commands to manage accounts (key pairs)
- `starport version` now prints detailed information about OS, Go version, and more
- Modules are scaffolded with genesis validation tests
- Types are scaffolded with tests for `ValidateBasic` methods
- `cosmosclient` has been refactored and can be used as a library for interacting with Cosmos SDK chains
- `starport relayer` uses `starport account`
- Added `--path` flag for all `scaffold`, `generate` and `chain` commands
- Added `--output` flag to the `build` command
- Configure port of gRPC web in `config.yml` with the `host.grpc-web` property
- Added `build.main` field to `config.yml` for apps to specify the path of the chain's main package. This property is
  required to be set only when an app contains multiple main packages.

### Fixes

- Scaffolding a message no longer prevents scaffolding a map, list, or single that has the same type name when using
  the `--no-message` flag
- Generate Go code from proto files only from default directories or directories specified in `config.yml`
- Fixed faucet token transfer calculation
- Removed `creator` field for types scaffolded with the `--no-message` flag
- Encode the count value in the store with `BigEndian`

## `v0.17.3`

### Fixes

- oracle: add a specific BandChain pkg version to avoid Cosmos SDK version conflicts

## `v0.17.2`

### Features

- `client.toml` is initialized and used by node's CLI, can be configured through `config.yml` with the `init.client`
  property
- Support serving Cosmos SDK `v0.43.x` based chains

## `v0.17.1`

### Fixes

- Set visibility to `public` on Gitpod's port 7575 to enable peer discovery for SPN
- Fixed GitHub action that releases blockchain node's binary
- Fixed an error in chain scaffolding due to "unknown revision"
- Fixed an error in `starport chain serve` by limiting the scope where proto files are searched for

## `v0.17`

### Features

- Added GitHub action that automatically builds and releases a binary
- The `--release` flag for the `build` command adds the ability to release binaries in a tarball with a checksum file.
- Added the flag `--no-module` to the command `starport app` to prevent scaffolding a default module when creating a new
  app
- Added `--dep` flag to specify module dependency when scaffolding a module
- Added support for multiple naming conventions for component names and field names
- Print created and modified files when scaffolding a new component
- Added `starport generate` namespace with commands to generate Go, Vuex and OpenAPI
- Added `starport chain init` command to initialize a chain without starting a node
- Scaffold a type that contains a single instance in the store
- Introduced `starport tools` command for advanced users. Existing `starport relayer lowlevel *` commands are also moved
  under `tools`
- Added `faucet.rate_limit_window` property to `config.yml`
- Simplified the `cmd` package in the template
- Added `starport scaffold band` oracle query scaffolding
- Updated TypeScript relayer to 0.2.0
- Added customizable gas limits for the relayer

### Fixes

- Use snake case for generated files
- Prevent using incorrect module name
- Fixed permissions issue when using Starport in Docker
- Ignore hidden directories when building a chain
- Fix error when scaffolding an IBC module in non-Starport chains

## `v0.16.2`

### Fix

- Prevent indirect Buf dependency

## `v0.16.1`

### Features

- Ensure that CLI operates fine even if the installation directory (bin) of Go programs is not configured properly

## `v0.16.0`

### Features

- The new `join` flag adds the ability to pass a `--genesis` file and `--peers` address list
  with `starport network chain join`
- The new `show` flag adds the ability to show `--genesis` and `--peers` list with `starport network chain show`
- `protoc` is now bundled with Spellshape CLI. You don't need to install it anymore.
- Starport is now published automatically on the Docker Hub
- `starport relayer` `configure` and `connect` commands now use
  the [confio/ts-relayer](https://github.com/confio/ts-relayer) under the hood. Also, checkout the
  new `starport relayer lowlevel` command
- An OpenAPI spec for your chain is now automatically generated with `serve` and `build` commands: a console is
  available at `localhost:1317` and spec at `localhost:1317/static/openapi.yml` by default for the newly scaffolded
  chains
- Keplr extension is supported on web apps created with Starport
- Added tests to the scaffold
- Improved reliability of scaffolding by detecting placeholders
- Added ability to scaffold modules in chains not created with Starport
- Added the ability to scaffold Cosmos SDK queries
- IBC relayer support is available on web apps created with Starport
- New types without CRUD operations can be added with the `--no-message` flag in the `type` command
- New packet without messages can be added with the `--no-message` flag in the `packet` command
- Added `docs` command to read Starport documentation on the CLI
- Published documentation on <https://docs.starport.network>
- Added `mnemonic` property to account in the `accounts` list to generate a key from a mnemonic

### Fixes

- `starport network chain join` hanging issue when creating an account
- Error when scaffolding a chain with an underscore in the repo name (thanks @bensooraj!)

### Changes

- `starport serve` no longer starts the web app in the `vue` directory (use `npm` to start it manually)
- Default scaffold no longer includes legacy REST API endpoints (thanks @bensooraj!)
- Removed support for Cosmos SDK v0.39 Launchpad

## `v0.15.0`

### Features

- IBC module scaffolding
- IBC packet scaffolding with acknowledgements
- JavaScript and Vuex client code generation for Cosmos SDK and custom modules
- Standalone relayer with `configure` and `connect` commands
- Advanced relayer options for configuring ports and versions
- Scaffold now follows `MsgServer` convention
- Message scaffolding
- Added `starport type ... --indexed` to scaffold indexed types
- Custom config file support with `starport serve -c custom.yml`
- Detailed terminal output for created accounts: name, address, mnemonic
- Added spinners to indicate progress for long-running commands
- Updated to Cosmos SDK v0.42.1

### Changes

- Replaced `packr` with Go 1.16 `embed`
- Renamed `servers` top-level property to `host`

## `v0.14.0`

### Features

- Chain state persistence between `starport serve` launches
- Integrated Stargate app's `scripts/protocgen` into Starport as a native feature. Running `starport build/serve` will
  automatically take care of building proto files without a need of script in the app's source code.
- Integrated third-party proto-files used by Cosmos SDK modules into Spellshape CLI
- Added ability to customize binary name with `build.binary` in `config.yml`
- Added ability to change path to home directory with ` .home` in `config.yml`
- Added ability to add accounts by `address` with in `config.yml`
- Added faucet functionality available on port 4500 and configurable with `faucet` in `config.yml`
- Added `starport faucet [address] [coins]` command
- Updated scaffold to Cosmos SDK v0.41.0
- Distroless multiplatform docker containers for starport that can be used for `starport serve`
- UI containers for chains scaffolded with Starport
- Use SOS-lite and Docker instead of systemD
- Arch PKGBUILD in `scripts`

### Fixes

- Support for CosmWasm on Stargate
- Bug with dashes in GitHub username breaking proto package name
- Bug with custom address prefix
- use docker buildx as a single command with multiple platforms to make multi-manifest work properly

## `v0.13.0`

### Features

- Added `starport network` commands for launching blockchains
- Added proxy (Chisel) to support launching blockchains from Gitpod
- Upgraded the template (Stargate) to Cosmos SDK v0.40.0-rc3
- Added a gRPC-Web proxy that is available under <http://localhost:12345/grpc>
- Added chain id configurability by recognizing `chain_id` from `genesis` section of `config.yml`.
- Added `config/app.toml` and `config/config.toml` configurability for appd under new `init.app` and `init.config`
  sections of `config.yml`
- Point to Stargate as default SDK version for scaffolding
- Covered CRUD operations for Stargate scaffolding
- Added docs on gopath to build from source directions
- Arch Linux Based Raspberry Pi development environment
- Calculate the necessary gas for sending transactions to SPN

### Fixes

- Routing REST API endpoints of querier on Stargate
- Evaluate `--address-prefix` option when scaffolding for Stargate
- Use a deterministic method to generate scaffolded type IDs
- Modify scaffolded type's creator type from address to string
- Copy built starport arm64 binary from tendermintdevelopment/starport:arm64 for device images
- Added git to amd64 docker image
- Comment out Gaia's seeds in the systemd unit template for downstream chains

## `v0.12.0`

### Features

- Added GitHub CLI to gitpod environment for greater ease of use
- Added `starport build` command to build and install app binaries
- Improved the first-time experience for readers of the Starport readme and parts of the Starport Handbook
- Added `starport module create` command to scaffold custom modules
- Raspberry Pi now installs, builds, and serves the Vue UI
- Improved documentation for Raspberry Pi Device Images
- Added IBC and some other modules
- Added an option to configure server addresses under `servers` section in `config.yml`

### Fixes

- `--address-prefix` will always be translated to lowercase while scaffolding with `app` command
- HTTP API: accept strings in JSON and cast them to int and bool
- Update @tendermint/vue to `v0.1.7`
- Removed "Starport Pi"
- Removed Makefile from Downstream Pi
- Fixed Downstream Pi image GitHub Action
- Prevent duplicated fields with `type` command
- Fixed handling of protobuf profiler: prof_laddr -> pprof_laddr
- Fix an error, when a Stargate `serve` cmd doesn't start if a user doesn't have a relayer installed

## `v0.11.1`

### Features

- Published on Snapcraft

## `v0.11.0`

### Features

- Added experimental [Stargate](https://stargate.cosmos.network/) scaffolding option with `--sdk-version stargate` flag
  on `starport app` command
- Pi Image Generation for chains generated with Starport
- GitHub action with capture of binary artifacts for chains generated with Starport
- Gitpod: added guidelines and changed working directory into `docs`
- Updated web scaffold with an improved sign in, balance list and a simple wallet
- Added CRUD actions for scaffolded types: delete, update, and get

## `v0.0.10`

### Features

- Add ARM64 releases
- OS Image Generation for Raspberry Pi 3 and 4
- Added `version` command
- Added support for _validator_ configuration in _config.yml_.
- Starport can be launched on Gitpod
- Added `make clean`

### Fixes

- Compile with go1.15
- Running `starport add type...` multiple times no longer breaks the app
- Running `appcli tx app create-x` now checks for all required args
- Removed unused `--denom` flag from the `app` command. It previously has moved as a prop to the `config.yml`
  under `accounts` section
- Disabled proxy server in the Vue app (this was causing to some compatibility issues) and enabled CORS
  for `appcli rest-server` instead
- `type` command supports dashes in app names

## `v0.0.10-rc.3`

### Features

- Configure `genesis.json` through `genesis` field in `config.yml`
- Initialize git repository on `app` scaffolding
- Check Go and GOPATH when running `serve`

### Changes

- verbose is --verbose, not -v, in the cli
- Renamed `frontend` directory to `vue`
- Added first E2E tests (for `app` and `add wasm` subcommands)

### Fixes

- No longer crashes when git is initialized but doesn't have commits
- Failure to start the frontend doesn't prevent Starport from running
- Changes to `config.yml` trigger reinitialization of the app
- Running `starport add wasm` multiple times no longer breaks the app

## `v0.0.10-rc.X`

### Features

- Initialize with accounts defined `config.yml`
- `starport serve --verbose` shows detailed output from every process
- Custom address prefixes with `--address-prefix` flag
- Cosmos SDK Launchpad support
- Rebuild and reinitialize on file change

## `v0.0.9`

Initial release.
