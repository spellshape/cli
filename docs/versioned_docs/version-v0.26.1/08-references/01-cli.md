---
description: Spellshape CLI docs.
---

# CLI commands

Documentation for Spellshape CLI.
## spellshape

Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain

**Synopsis**

Spellshape CLI is a tool for creating sovereign blockchains built with Cosmos SDK, the world’s
most popular modular blockchain framework. Spellshape CLI offers everything you need to scaffold,
test, build, and launch your blockchain.

To get started, create a blockchain:

	spellshape scaffold chain example


**Options**

```
  -h, --help   help for spellshape
```

**SEE ALSO**

* [spellshape account](#spellshape-account)	 - Create, delete, and show Spellshape accounts
* [spellshape chain](#spellshape-chain)	 - Build, init and start a blockchain node
* [spellshape completion](#spellshape-completion)	 - Generate the autocompletion script for the specified shell
* [spellshape docs](#spellshape-docs)	 - Show Spellshape CLI docs
* [spellshape generate](#spellshape-generate)	 - Generate clients, API docs from source code
* [spellshape network](#spellshape-network)	 - Launch a blockchain in production
* [spellshape node](#spellshape-node)	 - Make requests to a live blockchain node
* [spellshape plugin](#spellshape-plugin)	 - Handle plugins
* [spellshape relayer](#spellshape-relayer)	 - Connect blockchains with an IBC relayer
* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more
* [spellshape tools](#spellshape-tools)	 - Tools for advanced users
* [spellshape version](#spellshape-version)	 - Print the current build information


## spellshape account

Create, delete, and show Spellshape accounts

**Synopsis**

Commands for managing Spellshape accounts. An Spellshape account is a private/public
keypair stored in a keyring. Currently Spellshape accounts are used when interacting
with Spellshape relayer commands and when using "spellshape network" commands.

Note: Spellshape account commands are not for managing your chain's keys and accounts. Use
you chain's binary to manage accounts from "config.yml". For example, if your
blockchain is called "mychain", use "mychaind keys" to manage keys for the
chain.


**Options**

```
  -h, --help                     help for account
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape account create](#spellshape-account-create)	 - Create a new account
* [spellshape account delete](#spellshape-account-delete)	 - Delete an account by name
* [spellshape account export](#spellshape-account-export)	 - Export an account as a private key
* [spellshape account import](#spellshape-account-import)	 - Import an account by using a mnemonic or a private key
* [spellshape account list](#spellshape-account-list)	 - Show a list of all accounts
* [spellshape account show](#spellshape-account-show)	 - Show detailed information about a particular account


## spellshape account create

Create a new account

```
spellshape account create [name] [flags]
```

**Options**

```
  -h, --help   help for create
```

**Options inherited from parent commands**

```
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**SEE ALSO**

* [spellshape account](#spellshape-account)	 - Create, delete, and show Spellshape accounts


## spellshape account delete

Delete an account by name

```
spellshape account delete [name] [flags]
```

**Options**

```
  -h, --help   help for delete
```

**Options inherited from parent commands**

```
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**SEE ALSO**

* [spellshape account](#spellshape-account)	 - Create, delete, and show Spellshape accounts


## spellshape account export

Export an account as a private key

```
spellshape account export [name] [flags]
```

**Options**

```
  -h, --help                help for export
      --non-interactive     do not enter into interactive mode
      --passphrase string   passphrase to encrypt the exported key
      --path string         path to export private key. default: ./key_[name]
```

**Options inherited from parent commands**

```
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**SEE ALSO**

* [spellshape account](#spellshape-account)	 - Create, delete, and show Spellshape accounts


## spellshape account import

Import an account by using a mnemonic or a private key

```
spellshape account import [name] [flags]
```

**Options**

```
  -h, --help                help for import
      --non-interactive     do not enter into interactive mode
      --passphrase string   passphrase to decrypt the imported key (ignored when secret is a mnemonic)
      --secret string       Your mnemonic or path to your private key (use interactive mode instead to securely pass your mnemonic)
```

**Options inherited from parent commands**

```
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**SEE ALSO**

* [spellshape account](#spellshape-account)	 - Create, delete, and show Spellshape accounts


## spellshape account list

Show a list of all accounts

```
spellshape account list [flags]
```

**Options**

```
      --address-prefix string   account address prefix (default "cosmos")
  -h, --help                    help for list
```

**Options inherited from parent commands**

```
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**SEE ALSO**

* [spellshape account](#spellshape-account)	 - Create, delete, and show Spellshape accounts


## spellshape account show

Show detailed information about a particular account

```
spellshape account show [name] [flags]
```

**Options**

```
      --address-prefix string   account address prefix (default "cosmos")
  -h, --help                    help for show
```

**Options inherited from parent commands**

```
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**SEE ALSO**

* [spellshape account](#spellshape-account)	 - Create, delete, and show Spellshape accounts


## spellshape chain

Build, init and start a blockchain node

**Synopsis**

Commands in this namespace let you to build, initialize, and start your
blockchain node locally for development purposes.

To run these commands you should be inside the project's directory so that
Spellshape can find the source code. To ensure that you are, run "ls", you should
see the following files in the output: "go.mod", "x", "proto", "app", etc.

By default the "build" command will identify the "main" package of the project,
install dependencies if necessary, set build flags, compile the project into a
binary and install the binary. The "build" command is useful if you just want
the compiled binary, for example, to initialize and start the chain manually. It
can also be used to release your chain's binaries automatically as part of
continuous integration workflow.

The "init" command will build the chain's binary and use it to initialize a
local validator node. By default the validator node will be initialized in your
$HOME directory in a hidden directory that matches the name of your project.
This directory is called a data directory and contains a chain's genesis file
and a validator key. This command is useful if you want to quickly build and
initialize the data directory and use the chain's binary to manually start the
blockchain. The "init" command is meant only for development purposes, not
production.

The "serve" command builds, initializes, and starts your blockchain locally with
a single validator node for development purposes. "serve" also watches the
source code directory for file changes and intelligently
re-builds/initializes/starts the chain, essentially providing "code-reloading".
The "serve" command is meant only for development purposes, not production.

To distinguish between production and development consider the following.

In production, blockchains often run the same software on many validator nodes
that are run by different people and entities. To launch a blockchain in
production, the validator entities coordinate the launch process to start their
nodes simultaneously.

During development, a blockchain can be started locally on a single validator
node. This convenient process lets you restart a chain quickly and iterate
faster. Starting a chain on a single node in development is similar to starting
a traditional web application on a local server.

The "faucet" command lets you send tokens to an address from the "faucet"
account defined in "config.yml". Alternatively, you can use the chain's binary
to send token from any other account that exists on chain.

The "simulate" command helps you start a simulation testing process for your
chain.


**Options**

```
  -c, --config string   path to Spellshape config file (default: ./config.yml)
  -h, --help            help for chain
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape chain build](#spellshape-chain-build)	 - Build a node binary
* [spellshape chain debug](#spellshape-chain-debug)	 - Launch a debugger for a blockchain app
* [spellshape chain faucet](#spellshape-chain-faucet)	 - Send coins to an account
* [spellshape chain init](#spellshape-chain-init)	 - Initialize your chain
* [spellshape chain serve](#spellshape-chain-serve)	 - Start a blockchain node in development
* [spellshape chain simulate](#spellshape-chain-simulate)	 - Run simulation testing for the blockchain


## spellshape chain build

Build a node binary

**Synopsis**


The build command compiles the source code of the project into a binary and
installs the binary in the $(go env GOPATH)/bin directory.

You can customize the output directory for the binary using a flag:

	spellshape chain build --output dist

To compile the binary Spellshape first compiles protocol buffer (proto) files into
Go source code. Proto files contain required type and services definitions. If
you're using another program to compile proto files, you can use a flag to tell
Spellshape to skip the proto compilation step:

	spellshape chain build --skip-proto

Afterwards, Spellshape install dependencies specified in the go.mod file. By default
Spellshape doesn't check that dependencies of the main module stored in the module
cache have not been modified since they were downloaded. To enforce dependency
checking (essentially, running "go mod verify") use a flag:

	spellshape chain build --check-dependencies

Next, Spellshape identifies the "main" package of the project. By default the "main"
package is located in "cmd/{app}d" directory, where "{app}" is the name of the
scaffolded project and "d" stands for daemon. If your project contains more
than one "main" package, specify the path to the one that Spellshape should compile
in config.yml:

	build:
	  main: custom/path/to/main

By default the binary name will match the top-level module name (specified in
go.mod) with a suffix "d". This can be customized in config.yml:

	build:
	  binary: mychaind

You can also specify custom linker flags:

	build:
	  ldflags:
	    - "-X main.Version=development"
	    - "-X main.Date=01/05/2022T19:54"

To build binaries for a release, use the --release flag. The binaries for one or
more specified release targets are built in a "release/" directory in the
project's source directory. Specify the release targets with GOOS:GOARCH build
tags. If the optional --release.targets is not specified, a binary is created
for your current environment.

	spellshape chain build --release -t linux:amd64 -t darwin:amd64 -t darwin:arm64


```
spellshape chain build [flags]
```

**Options**

```
      --check-dependencies        verify that cached dependencies have not been modified since they were downloaded
      --clear-cache               clear the build cache (advanced)
      --debug                     build a debug binary
  -h, --help                      help for build
  -o, --output string             binary output path
  -p, --path string               path of the app (default ".")
      --release                   build for a release
      --release.prefix string     tarball prefix for each release target. Available only with --release flag
  -t, --release.targets strings   release targets. Available only with --release flag
      --skip-proto                skip file generation from proto
  -v, --verbose                   verbose output
```

**Options inherited from parent commands**

```
  -c, --config string   path to Spellshape config file (default: ./config.yml)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape chain](#spellshape-chain)	 - Build, init and start a blockchain node


## spellshape chain debug

Launch a debugger for a blockchain app

**Synopsis**

The debug command starts a debug server and launches a debugger.

Spellshape uses the Delve debugger by default. Delve enables you to interact with
your program by controlling the execution of the process, evaluating variables,
and providing information of thread / goroutine state, CPU register state and
more.

A debug server can optionally be started in cases where default terminal client
is not desirable. When the server starts it first runs the blockchain app,
attaches to it and finally waits for a client connection. It accepts both
JSON-RPC or DAP client connections.

To start a debug server use the following flag:

	spellshape chain debug --server

To start a debug server with a custom address use the following flags:

	spellshape chain debug --server --server-address 127.0.0.1:30500

The debug server stops automatically when the client connection is closed.


```
spellshape chain debug [flags]
```

**Options**

```
  -h, --help                    help for debug
  -p, --path string             path of the app (default ".")
      --server                  start a debug server
      --server-address string   debug server address (default "127.0.0.1:30500")
```

**Options inherited from parent commands**

```
  -c, --config string   path to Spellshape config file (default: ./config.yml)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape chain](#spellshape-chain)	 - Build, init and start a blockchain node


## spellshape chain faucet

Send coins to an account

```
spellshape chain faucet [address] [coin<,...>] [flags]
```

**Options**

```
  -h, --help          help for faucet
      --home string   directory where the blockchain node is initialized
  -p, --path string   path of the app (default ".")
  -v, --verbose       verbose output
```

**Options inherited from parent commands**

```
  -c, --config string   path to Spellshape config file (default: ./config.yml)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape chain](#spellshape-chain)	 - Build, init and start a blockchain node


## spellshape chain init

Initialize your chain

**Synopsis**

The init command compiles and installs the binary (like "spellshape chain build")
and uses that binary to initialize the blockchain's data directory for one
validator. To learn how the build process works, refer to "spellshape chain build
--help".

By default, the data directory will be initialized in $HOME/.mychain, where
"mychain" is the name of the project. To set a custom data directory use the
--home flag or set the value in config.yml:

	validators:
	  - name: alice
	    bonded: '100000000stake'
	    home: "~/.customdir"

The data directory contains three files in the "config" directory: app.toml,
config.toml, client.toml. These files let you customize the behavior of your
blockchain node and the client executable. When a chain is re-initialized the
data directory can be reset. To make some values in these files persistent, set
them in config.yml:

	validators:
	  - name: alice
	    bonded: '100000000stake'
	    app:
	      minimum-gas-prices: "0.025stake"
	    config:
	      consensus:
	        timeout_commit: "5s"
	        timeout_propose: "5s"
	    client:
	      output: "json"

The configuration above changes the minimum gas price of the validator (by
default the gas price is set to 0 to allow "free" transactions), sets the block
time to 5s, and changes the output format to JSON. To see what kind of values
this configuration accepts see the generated TOML files in the data directory.

As part of the initialization process Spellshape creates on-chain accounts with
token balances. By default, config.yml has two accounts in the top-level
"accounts" property. You can add more accounts and change their token balances.
Refer to config.yml guide to see which values you can set.

One of these accounts is a validator account and the amount of self-delegated
tokens can be set in the top-level "validator" property.

One of the most important components of an initialized chain is the genesis
file, the 0th block of the chain. The genesis file is stored in the data
directory "config" subdirectory and contains the initial state of the chain,
including consensus and module parameters. You can customize the values of the
genesis in config.yml:

	genesis:
	  app_state:
	    staking:
	      params:
	        bond_denom: "foo"

The example above changes the staking token to "foo". If you change the staking
denom, make sure the validator account has the right tokens.

The init command is meant to be used ONLY FOR DEVELOPMENT PURPOSES. Under the
hood it runs commands like "appd init", "appd add-genesis-account", "appd
gentx", and "appd collect-gentx". For production, you may want to run these
commands manually to ensure a production-level node initialization.


```
spellshape chain init [flags]
```

**Options**

```
      --check-dependencies   verify that cached dependencies have not been modified since they were downloaded
      --clear-cache          clear the build cache (advanced)
      --debug                build a debug binary
  -h, --help                 help for init
      --home string          directory where the blockchain node is initialized
  -p, --path string          path of the app (default ".")
      --skip-proto           skip file generation from proto
```

**Options inherited from parent commands**

```
  -c, --config string   path to Spellshape config file (default: ./config.yml)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape chain](#spellshape-chain)	 - Build, init and start a blockchain node


## spellshape chain serve

Start a blockchain node in development

**Synopsis**

The serve command compiles and installs the binary (like "spellshape chain build"),
uses that binary to initialize the blockchain's data directory for one validator
(like "spellshape chain init"), and starts the node locally for development purposes
with automatic code reloading.

Automatic code reloading means Spellshape starts watching the project directory.
Whenever a file change is detected, Spellshape automatically rebuilds, reinitializes
and restarts the node.

Whenever possible Spellshape will try to keep the current state of the chain by
exporting and importing the genesis file.

To force Spellshape to start from a clean slate even if a genesis file exists, use
the following flag:

	spellshape chain serve --reset-once

To force Spellshape to reset the state every time the source code is modified, use
the following flag:

	spellshape chain serve --force-reset

With Spellshape it's possible to start more than one blockchain from the same source
code using different config files. This is handy if you're building
inter-blockchain functionality and, for example, want to try sending packets
from one blockchain to another. To start a node using a specific config file:

	spellshape chain serve --config mars.yml

The serve command is meant to be used ONLY FOR DEVELOPMENT PURPOSES. Under the
hood, it runs "appd start", where "appd" is the name of your chain's binary. For
production, you may want to run "appd start" manually.


```
spellshape chain serve [flags]
```

**Options**

```
      --check-dependencies   verify that cached dependencies have not been modified since they were downloaded
      --clear-cache          clear the build cache (advanced)
  -f, --force-reset          force reset of the app state on start and every source change
      --generate-clients     generate code for the configured clients on reset or source code change
  -h, --help                 help for serve
      --home string          directory where the blockchain node is initialized
  -p, --path string          path of the app (default ".")
      --quit-on-fail         quit program if the app fails to start
  -r, --reset-once           reset the app state once on init
      --skip-proto           skip file generation from proto
  -v, --verbose              verbose output
```

**Options inherited from parent commands**

```
  -c, --config string   path to Spellshape config file (default: ./config.yml)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape chain](#spellshape-chain)	 - Build, init and start a blockchain node


## spellshape chain simulate

Run simulation testing for the blockchain

**Synopsis**

Run simulation testing for the blockchain. It sends many randomized-input messages of each module to a simulated node and checks if invariants break

```
spellshape chain simulate [flags]
```

**Options**

```
      --blockSize int             operations per block (default 30)
      --exportParamsHeight int    height to which export the randomly generated params
      --exportParamsPath string   custom file path to save the exported params JSON
      --exportStatePath string    custom file path to save the exported app state JSON
      --exportStatsPath string    custom file path to save the exported simulation statistics JSON
      --genesis string            custom simulation genesis file; cannot be used with params file
      --genesisTime int           override genesis UNIX time instead of using a random UNIX time
  -h, --help                      help for simulate
      --initialBlockHeight int    initial block to start the simulation (default 1)
      --lean                      lean simulation log output
      --numBlocks int             number of new blocks to simulate from the initial block height (default 200)
      --params string             custom simulation params file which overrides any random params; cannot be used with genesis
      --period uint               run slow invariants only once every period assertions
      --printAllInvariants        print all invariants if a broken invariant is found
      --seed int                  simulation random seed (default 42)
      --simulateEveryOperation    run slow invariants every operation
  -v, --verbose                   verbose log output
```

**Options inherited from parent commands**

```
  -c, --config string   path to Spellshape config file (default: ./config.yml)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape chain](#spellshape-chain)	 - Build, init and start a blockchain node


## spellshape completion

Generate the autocompletion script for the specified shell

**Synopsis**

Generate the autocompletion script for spellshape for the specified shell.
See each sub-command's help for details on how to use the generated script.


**Options**

```
  -h, --help   help for completion
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape completion bash](#spellshape-completion-bash)	 - Generate the autocompletion script for bash
* [spellshape completion fish](#spellshape-completion-fish)	 - Generate the autocompletion script for fish
* [spellshape completion powershell](#spellshape-completion-powershell)	 - Generate the autocompletion script for powershell
* [spellshape completion zsh](#spellshape-completion-zsh)	 - Generate the autocompletion script for zsh


## spellshape completion bash

Generate the autocompletion script for bash

**Synopsis**

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(spellshape completion bash)

To load completions for every new session, execute once:

**#### Linux:**

	spellshape completion bash > /etc/bash_completion.d/spellshape

**#### macOS:**

	spellshape completion bash > $(brew --prefix)/etc/bash_completion.d/spellshape

You will need to start a new shell for this setup to take effect.


```
spellshape completion bash
```

**Options**

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

**SEE ALSO**

* [spellshape completion](#spellshape-completion)	 - Generate the autocompletion script for the specified shell


## spellshape completion fish

Generate the autocompletion script for fish

**Synopsis**

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	spellshape completion fish | source

To load completions for every new session, execute once:

	spellshape completion fish > ~/.config/fish/completions/spellshape.fish

You will need to start a new shell for this setup to take effect.


```
spellshape completion fish [flags]
```

**Options**

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

**SEE ALSO**

* [spellshape completion](#spellshape-completion)	 - Generate the autocompletion script for the specified shell


## spellshape completion powershell

Generate the autocompletion script for powershell

**Synopsis**

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	spellshape completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
spellshape completion powershell [flags]
```

**Options**

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

**SEE ALSO**

* [spellshape completion](#spellshape-completion)	 - Generate the autocompletion script for the specified shell


## spellshape completion zsh

Generate the autocompletion script for zsh

**Synopsis**

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(spellshape completion zsh); compdef _spellshape spellshape

To load completions for every new session, execute once:

**#### Linux:**

	spellshape completion zsh > "${fpath[1]}/_spellshape"

**#### macOS:**

	spellshape completion zsh > $(brew --prefix)/share/zsh/site-functions/_spellshape

You will need to start a new shell for this setup to take effect.


```
spellshape completion zsh [flags]
```

**Options**

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

**SEE ALSO**

* [spellshape completion](#spellshape-completion)	 - Generate the autocompletion script for the specified shell


## spellshape docs

Show Spellshape CLI docs

```
spellshape docs [flags]
```

**Options**

```
  -h, --help   help for docs
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain


## spellshape generate

Generate clients, API docs from source code

**Synopsis**

Generate clients, API docs from source code.

Such as compiling protocol buffer files into Go or implement particular
functionality, for example, generating an OpenAPI spec.

Produced source code can be regenerated by running a command again and is not
meant to be edited by hand.


**Options**

```
      --clear-cache   clear the build cache (advanced)
  -h, --help          help for generate
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape generate composables](#spellshape-generate-composables)	 - TypeScript frontend client and Vue 3 composables
* [spellshape generate hooks](#spellshape-generate-hooks)	 - TypeScript frontend client and React hooks
* [spellshape generate openapi](#spellshape-generate-openapi)	 - OpenAPI spec for your chain
* [spellshape generate proto-go](#spellshape-generate-proto-go)	 - Compile protocol buffer files to Go source code required by Cosmos SDK
* [spellshape generate ts-client](#spellshape-generate-ts-client)	 - TypeScript frontend client
* [spellshape generate vuex](#spellshape-generate-vuex)	 - *DEPRECATED* TypeScript frontend client and Vuex stores


## spellshape generate composables

TypeScript frontend client and Vue 3 composables

```
spellshape generate composables [flags]
```

**Options**

```
  -h, --help            help for composables
  -o, --output string   Vue 3 composables output path
  -y, --yes             answers interactive yes/no questions with yes
```

**Options inherited from parent commands**

```
      --clear-cache   clear the build cache (advanced)
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [spellshape generate](#spellshape-generate)	 - Generate clients, API docs from source code


## spellshape generate hooks

TypeScript frontend client and React hooks

```
spellshape generate hooks [flags]
```

**Options**

```
  -h, --help            help for hooks
  -o, --output string   React hooks output path
  -y, --yes             answers interactive yes/no questions with yes
```

**Options inherited from parent commands**

```
      --clear-cache   clear the build cache (advanced)
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [spellshape generate](#spellshape-generate)	 - Generate clients, API docs from source code


## spellshape generate openapi

OpenAPI spec for your chain

```
spellshape generate openapi [flags]
```

**Options**

```
  -h, --help   help for openapi
  -y, --yes    answers interactive yes/no questions with yes
```

**Options inherited from parent commands**

```
      --clear-cache   clear the build cache (advanced)
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [spellshape generate](#spellshape-generate)	 - Generate clients, API docs from source code


## spellshape generate proto-go

Compile protocol buffer files to Go source code required by Cosmos SDK

```
spellshape generate proto-go [flags]
```

**Options**

```
  -h, --help   help for proto-go
  -y, --yes    answers interactive yes/no questions with yes
```

**Options inherited from parent commands**

```
      --clear-cache   clear the build cache (advanced)
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [spellshape generate](#spellshape-generate)	 - Generate clients, API docs from source code


## spellshape generate ts-client

TypeScript frontend client

**Synopsis**

Generate a framework agnostic TypeScript client for your blockchain project.

By default the TypeScript client is generated in the "ts-client/" directory. You
can customize the output directory in config.yml:

	client:
	  typescript:
	    path: new-path

Output can also be customized by using a flag:

	spellshape generate ts-client --output new-path

TypeScript client code can be automatically regenerated on reset or source code
changes when the blockchain is started with a flag:

	spellshape chain serve --generate-clients


```
spellshape generate ts-client [flags]
```

**Options**

```
  -h, --help            help for ts-client
  -o, --output string   TypeScript client output path
      --use-cache       use build cache to speed-up generation
  -y, --yes             answers interactive yes/no questions with yes
```

**Options inherited from parent commands**

```
      --clear-cache   clear the build cache (advanced)
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [spellshape generate](#spellshape-generate)	 - Generate clients, API docs from source code


## spellshape generate vuex

*DEPRECATED* TypeScript frontend client and Vuex stores

```
spellshape generate vuex [flags]
```

**Options**

```
  -h, --help            help for vuex
  -o, --output string   Vuex store output path
  -y, --yes             answers interactive yes/no questions with yes
```

**Options inherited from parent commands**

```
      --clear-cache   clear the build cache (advanced)
  -p, --path string   path of the app (default ".")
```

**SEE ALSO**

* [spellshape generate](#spellshape-generate)	 - Generate clients, API docs from source code


## spellshape network

Launch a blockchain in production

**Synopsis**


Spellshape Network commands allow to coordinate the launch of sovereign Cosmos blockchains.

To launch a Cosmos blockchain you need someone to be a coordinator and others to
be validators. These are just roles, anyone can be a coordinator or a validator.
A coordinator publishes information about a chain to be launched on the Spellshape
blockchain, approves validator requests and coordinates the launch. Validators
send requests to join a chain and start their nodes when a blockchain is ready
for launch.

To publish the information about your chain as a coordinator run the following
command (the URL should point to a repository with a Cosmos SDK chain):

	spellshape network chain publish github.com/spellshape/example

This command will return a launch identifier you will be using in the following
commands. Let's say this identifier is 42.

Next, ask validators to initialize their nodes and request to join the network
as validators. For a testnet you can use the default values suggested by the
CLI.

	spellshape network chain init 42

	spellshape network chain join 42 --amount 95000000stake

As a coordinator list all validator requests:

	spellshape network request list 42

Approve validator requests:

	spellshape network request approve 42 1,2

Once you've approved all validators you need in the validator set, announce that
the chain is ready for launch:

	spellshape network chain launch 42

Validators can now prepare their nodes for launch:

	spellshape network chain prepare 42

The output of this command will show a command that a validator would use to
launch their node, for example “exampled --home ~/.example”. After enough
validators launch their nodes, a blockchain will be live.


**Options**

```
  -h, --help                        help for network
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch
* [spellshape network coordinator](#spellshape-network-coordinator)	 - Show and update a coordinator profile
* [spellshape network profile](#spellshape-network-profile)	 - Show the address profile info
* [spellshape network project](#spellshape-network-project)	 - Handle projects
* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests
* [spellshape network reward](#spellshape-network-reward)	 - Manage network rewards
* [spellshape network tool](#spellshape-network-tool)	 - Commands to run subsidiary tools
* [spellshape network validator](#spellshape-network-validator)	 - Show and update a validator profile
* [spellshape network version](#spellshape-network-version)	 - Version of the plugin


## spellshape network chain

Publish a chain, join as a validator and prepare node for launch

**Synopsis**

The "chain" namespace features the most commonly used commands for launching
blockchains with Spellshape.

As a coordinator you "publish" your blockchain to Spellshape. When enough validators
are approved for the genesis and no changes are excepted to be made to the
genesis, a coordinator announces that the chain is ready for launch with the
"launch" command. In the case of an unsuccessful launch, the coordinator can revert it
using the "revert-launch" command.

As a validator, you "init" your node and apply to become a validator for a
blockchain with the "join" command. After the launch of the chain is announced,
validators can generate the finalized genesis and download the list of peers with the
"prepare" command.

The "install" command can be used to download, compile the source code and
install the chain's binary locally. The binary can be used, for example, to
initialize a validator node or to interact with the chain after it has been
launched.

All chains published to Spellshape can be listed by using the "list" command.


**Options**

```
  -h, --help   help for chain
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network](#spellshape-network)	 - Launch a blockchain in production
* [spellshape network chain init](#spellshape-network-chain-init)	 - Initialize a chain from a published chain ID
* [spellshape network chain install](#spellshape-network-chain-install)	 - Install chain binary for a launch
* [spellshape network chain join](#spellshape-network-chain-join)	 - Request to join a network as a validator
* [spellshape network chain launch](#spellshape-network-chain-launch)	 - Trigger the launch of a chain
* [spellshape network chain list](#spellshape-network-chain-list)	 - List published chains
* [spellshape network chain prepare](#spellshape-network-chain-prepare)	 - Prepare the chain for launch
* [spellshape network chain publish](#spellshape-network-chain-publish)	 - Publish a new chain to start a new network
* [spellshape network chain revert-launch](#spellshape-network-chain-revert-launch)	 - Revert launch of a network as a coordinator
* [spellshape network chain show](#spellshape-network-chain-show)	 - Show details of a chain


## spellshape network chain init

Initialize a chain from a published chain ID

**Synopsis**

Spellshape network chain init is a command used by validators to initialize a
validator node for a blockchain from the information stored on the Spellshape chain.

	spellshape network chain init 42

This command fetches the information about a chain with launch ID 42. The source
code of the chain is cloned in a temporary directory, and the node's binary is
compiled from the source. The binary is then used to initialize the node. By
default, Spellshape uses "~/spn/[launch-id]/" as the home directory for the blockchain.

An important part of initializing a validator node is creation of the gentx (a
transaction that adds a validator at the genesis of the chain).

The "init" command will prompt for values like self-delegation and commission.
These values will be used in the validator's gentx. You can use flags to provide
the values in non-interactive mode.

Use the "--home" flag to choose a different path for the home directory of the
blockchain:

	spellshape network chain init 42 --home ~/mychain

The end result of the "init" command is a validator home directory with a
genesis validator transaction (gentx) file.

```
spellshape network chain init [launch-id] [flags]
```

**Options**

```
      --check-dependencies                  verify that cached dependencies have not been modified since they were downloaded
      --clear-cache                         clear the build cache (advanced)
      --from string                         account name to use for sending transactions to SPN (default "default")
  -h, --help                                help for init
      --home string                         home directory used for blockchains
      --keyring-backend string              keyring backend to store your account keys (default "test")
      --keyring-dir string                  accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --validator-account string            account for the chain validator (default "default")
      --validator-details string            details about the validator
      --validator-gas-price string          validator gas price
      --validator-identity string           validator identity signature (ex. UPort or Keybase)
      --validator-moniker string            custom validator moniker
      --validator-security-contact string   validator security contact email
      --validator-self-delegation string    validator minimum self delegation
      --validator-website string            associate a website with the validator
  -y, --yes                                 answers interactive yes/no questions with yes
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch


## spellshape network chain install

Install chain binary for a launch

```
spellshape network chain install [launch-id] [flags]
```

**Options**

```
      --check-dependencies   verify that cached dependencies have not been modified since they were downloaded
      --clear-cache          clear the build cache (advanced)
      --from string          account name to use for sending transactions to SPN (default "default")
  -h, --help                 help for install
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch


## spellshape network chain join

Request to join a network as a validator

**Synopsis**

The "join" command is used by validators to send a request to join a blockchain.
The required argument is a launch ID of a blockchain. The "join" command expects
that the validator has already setup a home directory for the blockchain and has
a gentx either by running "spellshape network chain init" or initializing the data
directory manually with the chain's binary.

By default the "join" command just sends the request to join as a validator.
However, often a validator also needs to request an genesis account with a token
balance to afford self-delegation.

The following command will send a request to join blockchain with launch ID 42
as a validator and request to be added as an account with a token balance of
95000000 STAKE.

	spellshape network chain join 42 --amount 95000000stake

A request to join as a validator contains a gentx file. Spellshape looks for gentx
in a home directory used by "spellshape network chain init" by default. To use a
different directory, use the "--home" flag or pass a gentx file directly with
the  "--gentx" flag.

To join a chain as a validator, you must provide the IP address of your node so
that other validators can connect to it. The join command will ask you for the
IP address and will attempt to automatically detect and fill in the value. If
you want to manually specify the IP address, you can use the "--peer-address"
flag:

	spellshape network chain join 42 --peer-address 0.0.0.0

Since "join" broadcasts a transaction to the Spellshape blockchain, you will need an
account on the Spellshape blockchain. During the testnet phase, however, Spellshape
automatically requests tokens from a faucet.


```
spellshape network chain join [launch-id] [flags]
```

**Options**

```
      --amount string            amount of coins for account request (ignored if coordinator has fixed the account balances or if --no-acount flag is set)
      --check-dependencies       verify that cached dependencies have not been modified since they were downloaded
      --from string              account name to use for sending transactions to SPN (default "default")
      --gentx string             path to a gentx json file
  -h, --help                     help for join
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --no-account               prevent sending a request for a genesis account
      --peer-address string      peer's address
  -y, --yes                      answers interactive yes/no questions with yes
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch


## spellshape network chain launch

Trigger the launch of a chain

**Synopsis**

The launch command communicates to the world that the chain is ready to be
launched.

Only the coordinator of the chain can execute the launch command.

	spellshape network chain launch 42

After the launch command is executed no changes to the genesis are accepted. For
example, validators will no longer be able to successfully execute the "spellshape
network chain join" command to apply as a validator.

The launch command sets the date and time after which the chain will start. By
default, the current time is set. To give validators more time to prepare for
the launch, set the time with the "--launch-time" flag:

	spellshape network chain launch 42 --launch-time 2023-01-01T00:00:00Z

After the launch command is executed, validators can generate the finalized
genesis and prepare their nodes for the launch. For example, validators can run
"spellshape network chain prepare" to generate the genesis and populate the peer
list.

If you want to change the launch time or open up the genesis file for changes
you can use "spellshape network chain revert-launch" to make it possible, for
example, to accept new validators and add accounts.


```
spellshape network chain launch [launch-id] [flags]
```

**Options**

```
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for launch
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --launch-time string       timestamp the chain is effectively launched (example "2022-01-01T00:00:00Z")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch


## spellshape network chain list

List published chains

```
spellshape network chain list [flags]
```

**Options**

```
      --advanced     show advanced information about the chains
  -h, --help         help for list
      --limit uint   limit of results per page (default 100)
      --page uint    page for chain list result (default 1)
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch


## spellshape network chain prepare

Prepare the chain for launch

**Synopsis**

The prepare command prepares a validator node for the chain launch by generating
the final genesis and adding IP addresses of peers to the validator's
configuration file.

	spellshape network chain prepare 42

By default, Spellshape uses "$HOME/spn/LAUNCH_ID" as the data directory. If you used
a different data directory when initializing the node, use the "--home" flag and
set the correct path to the data directory.

Spellshape generates the genesis file in "config/genesis.json" and adds peer IPs by
modifying "config/config.toml".

The prepare command should be executed after the coordinator has triggered the
chain launch and finalized the genesis with "spellshape network chain launch". You
can force Spellshape to run the prepare command without checking if the launch has
been triggered with the "--force" flag (this is not recommended).

After the prepare command is executed the node is ready to be started.


```
spellshape network chain prepare [launch-id] [flags]
```

**Options**

```
      --check-dependencies       verify that cached dependencies have not been modified since they were downloaded
      --clear-cache              clear the build cache (advanced)
  -f, --force                    force the prepare command to run even if the chain is not launched
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for prepare
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch


## spellshape network chain publish

Publish a new chain to start a new network

**Synopsis**

To begin the process of launching a blockchain with Spellshape, a coordinator needs
to publish the information about a blockchain. The only required bit of
information is the URL of the source code of the blockchain.

The following command publishes the information about an example blockchain:

	spellshape network chain publish github.com/spellshape/example

This command fetches the source code of the blockchain, compiles the binary,
verifies that a blockchain can be started with the binary, and publishes the
information about the blockchain to Spellshape. Currently, only public repositories
are supported. The command returns an integer number that acts as an identifier
of the chain on Spellshape.

By publishing a blockchain on Spellshape you become the "coordinator" of this
blockchain. A coordinator is an account that has the authority to approve and
reject validator requests, set parameters of the blockchain and trigger the
launch of the chain.

The default Git branch is used when publishing a chain. If you want to use a
specific branch, tag or a commit hash, use "--branch", "--tag", or "--hash"
flags respectively.

The repository name is used as the default chain ID. Spellshape does not ensure that
chain IDs are unique, but they have to have a valid format: [string]-[integer].
To set a custom chain ID use the "--chain-id" flag.

	spellshape network chain publish github.com/spellshape/example --chain-id foo-1

Once the chain is published users can request accounts with coin balances to be
added to the chain's genesis. By default, users are free to request any number
of tokens. If you want all users requesting tokens to get the same amount, use
the "--account-balance" flag with a list of coins.

	spellshape network chain publish github.com/spellshape/example --account-balance 2000foocoin


```
spellshape network chain publish [source-url] [flags]
```

**Options**

```
      --account-balance string   balance for each approved genesis account for the chain
      --amount string            amount of coins for account request
      --branch string            Git branch to use for the repo
      --chain-id string          chain ID to use for this network
      --check-dependencies       verify that cached dependencies have not been modified since they were downloaded
      --clear-cache              clear the build cache (advanced)
      --from string              account name to use for sending transactions to SPN (default "default")
      --genesis-config string    name of an Spellshape config file in the repo for custom Genesis
      --genesis-url string       URL to a custom Genesis
      --hash string              Git hash to use for the repo
  -h, --help                     help for publish
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --mainnet                  initialize a mainnet project
      --metadata string          add chain metadata
      --no-check                 skip verifying chain's integrity
      --project uint             project ID to use for this network
      --reward.coins string      reward coins
      --reward.height int        last reward height
      --shares string            add shares for the project
      --tag string               Git tag to use for the repo
      --total-supply string      add a total of the mainnet of a project
  -y, --yes                      answers interactive yes/no questions with yes
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch


## spellshape network chain revert-launch

Revert launch of a network as a coordinator

**Synopsis**

The revert launch command reverts the previously scheduled launch of a chain.

Only the coordinator of the chain can execute the launch command.

	spellshape network chain revert-launch 42

After the revert launch command is executed, changes to the genesis of the chain
are allowed again. For example, validators will be able to request to join the
chain. Revert launch also resets the launch time.


```
spellshape network chain revert-launch [launch-id] [flags]
```

**Options**

```
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for revert-launch
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch


## spellshape network chain show

Show details of a chain

**Options**

```
  -h, --help   help for show
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain](#spellshape-network-chain)	 - Publish a chain, join as a validator and prepare node for launch
* [spellshape network chain show accounts](#spellshape-network-chain-show-accounts)	 - Show all vesting and genesis accounts of the chain
* [spellshape network chain show genesis](#spellshape-network-chain-show-genesis)	 - Show the chain genesis file
* [spellshape network chain show info](#spellshape-network-chain-show-info)	 - Show info details of the chain
* [spellshape network chain show peers](#spellshape-network-chain-show-peers)	 - Show peers list of the chain
* [spellshape network chain show validators](#spellshape-network-chain-show-validators)	 - Show all validators of the chain


## spellshape network chain show accounts

Show all vesting and genesis accounts of the chain

```
spellshape network chain show accounts [launch-id] [flags]
```

**Options**

```
  -h, --help            help for accounts
      --prefix string   account address prefix (default "spn")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain show](#spellshape-network-chain-show)	 - Show details of a chain


## spellshape network chain show genesis

Show the chain genesis file

```
spellshape network chain show genesis [launch-id] [flags]
```

**Options**

```
      --clear-cache   clear the build cache (advanced)
  -h, --help          help for genesis
      --out string    path to output Genesis file (default "./genesis.json")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain show](#spellshape-network-chain-show)	 - Show details of a chain


## spellshape network chain show info

Show info details of the chain

```
spellshape network chain show info [launch-id] [flags]
```

**Options**

```
  -h, --help   help for info
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain show](#spellshape-network-chain-show)	 - Show details of a chain


## spellshape network chain show peers

Show peers list of the chain

```
spellshape network chain show peers [launch-id] [flags]
```

**Options**

```
  -h, --help         help for peers
      --out string   path to output peers list (default "./peers.txt")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain show](#spellshape-network-chain-show)	 - Show details of a chain


## spellshape network chain show validators

Show all validators of the chain

```
spellshape network chain show validators [launch-id] [flags]
```

**Options**

```
  -h, --help            help for validators
      --prefix string   account address prefix (default "spn")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network chain show](#spellshape-network-chain-show)	 - Show details of a chain


## spellshape network coordinator

Show and update a coordinator profile

**Options**

```
  -h, --help   help for coordinator
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network](#spellshape-network)	 - Launch a blockchain in production
* [spellshape network coordinator set](#spellshape-network-coordinator-set)	 - Set an information in a coordinator profile
* [spellshape network coordinator show](#spellshape-network-coordinator-show)	 - Show a coordinator profile


## spellshape network coordinator set

Set an information in a coordinator profile

**Synopsis**

Coordinators on Spellshape can set a profile containing a description for the coordinator.
The coordinator set command allows to set information for the coordinator.
The following information can be set:
- details: general information about the coordinator.
- identity: a piece of information to verify the identity of the coordinator with a system like Keybase or Veramo.
- website: website of the coordinator.


```
spellshape network coordinator set details|identity|website [value] [flags]
```

**Options**

```
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for set
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network coordinator](#spellshape-network-coordinator)	 - Show and update a coordinator profile


## spellshape network coordinator show

Show a coordinator profile

```
spellshape network coordinator show [address] [flags]
```

**Options**

```
  -h, --help   help for show
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network coordinator](#spellshape-network-coordinator)	 - Show and update a coordinator profile


## spellshape network profile

Show the address profile info

```
spellshape network profile [project-id] [flags]
```

**Options**

```
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for profile
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network](#spellshape-network)	 - Launch a blockchain in production


## spellshape network project

Handle projects

**Options**

```
  -h, --help   help for project
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network](#spellshape-network)	 - Launch a blockchain in production
* [spellshape network project account](#spellshape-network-project-account)	 - Handle project accounts
* [spellshape network project create](#spellshape-network-project-create)	 - Create a project
* [spellshape network project list](#spellshape-network-project-list)	 - List published projects
* [spellshape network project show](#spellshape-network-project-show)	 - Show published project
* [spellshape network project update](#spellshape-network-project-update)	 - Update details fo the project of the project


## spellshape network project account

Handle project accounts

**Options**

```
  -h, --help   help for account
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network project](#spellshape-network-project)	 - Handle projects
* [spellshape network project account list](#spellshape-network-project-account-list)	 - Show all mainnet and mainnet vesting of the project


## spellshape network project account list

Show all mainnet and mainnet vesting of the project

```
spellshape network project account list [project-id] [flags]
```

**Options**

```
  -h, --help   help for list
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network project account](#spellshape-network-project-account)	 - Handle project accounts


## spellshape network project create

Create a project

```
spellshape network project create [name] [total-supply] [flags]
```

**Options**

```
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for create
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --metadata string          Add a metadata to the chain
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network project](#spellshape-network-project)	 - Handle projects


## spellshape network project list

List published projects

```
spellshape network project list [flags]
```

**Options**

```
  -h, --help   help for list
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network project](#spellshape-network-project)	 - Handle projects


## spellshape network project show

Show published project

```
spellshape network project show [project-id] [flags]
```

**Options**

```
  -h, --help   help for show
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network project](#spellshape-network-project)	 - Handle projects


## spellshape network project update

Update details fo the project of the project

```
spellshape network project update [project-id] [flags]
```

**Options**

```
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for update
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --metadata string          update the project metadata
      --name string              update the project name
      --total-supply string      update the total of the mainnet of a project
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network project](#spellshape-network-project)	 - Handle projects


## spellshape network request

Create, show, reject and approve requests

**Synopsis**

The "request" namespace contains commands for creating, showing, approving, and
rejecting requests.

A request is mechanism in Spellshape that allows changes to be made to the genesis
file like adding accounts with token balances and validators. Anyone can submit
a request, but only the coordinator of a chain can approve or reject a request.

Each request has a status:

* Pending: waiting for the approval of the coordinator
* Approved: approved by the coordinator, its content has been applied to the
  launch information
* Rejected: rejected by the coordinator or the request creator


**Options**

```
  -h, --help   help for request
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network](#spellshape-network)	 - Launch a blockchain in production
* [spellshape network request add-account](#spellshape-network-request-add-account)	 - Send request to add account
* [spellshape network request approve](#spellshape-network-request-approve)	 - Approve requests
* [spellshape network request change-param](#spellshape-network-request-change-param)	 - Send request to change a module param
* [spellshape network request list](#spellshape-network-request-list)	 - List all requests for a chain
* [spellshape network request reject](#spellshape-network-request-reject)	 - Reject requests
* [spellshape network request remove-account](#spellshape-network-request-remove-account)	 - Send request to remove a genesis account
* [spellshape network request remove-validator](#spellshape-network-request-remove-validator)	 - Send request to remove a validator
* [spellshape network request show](#spellshape-network-request-show)	 - Show detailed information about a request
* [spellshape network request verify](#spellshape-network-request-verify)	 - Verify the request and simulate the chain genesis from them


## spellshape network request add-account

Send request to add account

**Synopsis**

The "add account" command creates a new request to add an account with a given
address and a specified coin balance to the genesis of the chain.

The request automatically fails to be applied if a genesis account or a vesting
account with an identical address is already specified in the launch
information.

If a coordinator has specified that all genesis accounts on a chain should have
the same balance (useful for testnets, for example), the "add account" expects
only an address as an argument. Attempt to provide a token balance will result
in an error.


```
spellshape network request add-account [launch-id] [address] [coins] [flags]
```

**Options**

```
      --clear-cache              clear the build cache (advanced)
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for add-account
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests


## spellshape network request approve

Approve requests

**Synopsis**

The "approve" command is used by a chain's coordinator to approve requests.
Multiple requests can be approved using a comma-separated list and/or using a
dash syntax.

	spellshape network request approve 42 1,2,3-6,7,8

The command above approves requests with IDs from 1 to 8 included on a chain
with a launch ID 42.

When requests are approved Spellshape applies the requested changes and simulates
initializing and launching the chain locally. If the chain starts successfully,
requests are considered to be "verified" and are approved. If one or more
requested changes stop the chain from launching locally, the verification
process fails and the approval of all requests is canceled. To skip the
verification process use the "--no-verification" flag.

Note that Spellshape will try to approve requests in the same order as request IDs
are submitted to the "approve" command.

```
spellshape network request approve [launch-id] [number<,...>] [flags]
```

**Options**

```
      --clear-cache              clear the build cache (advanced)
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for approve
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --no-verification          approve the requests without verifying them
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests


## spellshape network request change-param

Send request to change a module param

```
spellshape network request change-param [launch-id] [module-name] [param-name] [value (json, string, number)] [flags]
```

**Options**

```
      --clear-cache              clear the build cache (advanced)
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for change-param
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests


## spellshape network request list

List all requests for a chain

```
spellshape network request list [launch-id] [flags]
```

**Options**

```
  -h, --help            help for list
      --prefix string   account address prefix (default "spn")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests


## spellshape network request reject

Reject requests

**Synopsis**

The "reject" command is used by a chain's coordinator to reject requests.

	spellshape network request reject 42 1,2,3-6,7,8

The syntax of the "reject" command is similar to that of the "approve" command.


```
spellshape network request reject [launch-id] [number<,...>] [flags]
```

**Options**

```
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for reject
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests


## spellshape network request remove-account

Send request to remove a genesis account

```
spellshape network request remove-account [launch-id] [address] [flags]
```

**Options**

```
      --clear-cache              clear the build cache (advanced)
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for remove-account
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests


## spellshape network request remove-validator

Send request to remove a validator

```
spellshape network request remove-validator [launch-id] [address] [flags]
```

**Options**

```
      --clear-cache              clear the build cache (advanced)
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for remove-validator
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests


## spellshape network request show

Show detailed information about a request

```
spellshape network request show [launch-id] [request-id] [flags]
```

**Options**

```
  -h, --help   help for show
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests


## spellshape network request verify

Verify the request and simulate the chain genesis from them

**Synopsis**

The "verify" command applies selected requests to the genesis of a chain locally
to verify that approving these requests will result in a valid genesis that
allows a chain to launch without issues. This command does not approve requests,
only checks them.


```
spellshape network request verify [launch-id] [number<,...>] [flags]
```

**Options**

```
      --clear-cache              clear the build cache (advanced)
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for verify
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network request](#spellshape-network-request)	 - Create, show, reject and approve requests


## spellshape network reward

Manage network rewards

**Options**

```
  -h, --help   help for reward
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network](#spellshape-network)	 - Launch a blockchain in production
* [spellshape network reward release](#spellshape-network-reward-release)	 - Connect the monitoring modules of launched chains with SPN
* [spellshape network reward set](#spellshape-network-reward-set)	 - set a network chain reward


## spellshape network reward release

Connect the monitoring modules of launched chains with SPN

```
spellshape network reward release [launch-id] [chain-rpc] [flags]
```

**Options**

```
      --create-client-only        only create the network client id
      --from string               account name to use for sending transactions to SPN (default "default")
  -h, --help                      help for release
      --keyring-backend string    keyring backend to store your account keys (default "test")
      --spn-gaslimit int          gas limit used for transactions on SPN (default 400000)
      --spn-gasprice string       gas price used for transactions on SPN (default "0.0000025uspn")
      --testnet-account string    testnet chain account (default "default")
      --testnet-faucet string     faucet address of the testnet chain
      --testnet-gaslimit int      gas limit used for transactions on testnet chain (default 400000)
      --testnet-gasprice string   gas price used for transactions on testnet chain (default "0.0000025stake")
      --testnet-prefix string     address prefix of the testnet chain (default "cosmos")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network reward](#spellshape-network-reward)	 - Manage network rewards


## spellshape network reward set

set a network chain reward

```
spellshape network reward set [launch-id] [last-reward-height] [coins] [flags]
```

**Options**

```
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for set
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network reward](#spellshape-network-reward)	 - Manage network rewards


## spellshape network tool

Commands to run subsidiary tools

**Options**

```
  -h, --help   help for tool
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network](#spellshape-network)	 - Launch a blockchain in production
* [spellshape network tool proxy-tunnel](#spellshape-network-tool-proxy-tunnel)	 - Setup a proxy tunnel via HTTP


## spellshape network tool proxy-tunnel

Setup a proxy tunnel via HTTP

**Synopsis**

Starts an HTTP proxy server and HTTP proxy clients for each node that
needs HTTP tunneling.

HTTP tunneling is activated **ONLY** if SPN_CONFIG_FILE has "tunneled_peers"
field inside with a list of tunneled peers/nodes.

If you're using SPN as coordinator and do not want to allow HTTP tunneling
feature at all, you can prevent "spn.yml" file to being generated by not
approving validator requests that has HTTP tunneling enabled instead of plain
TCP connections.

```
spellshape network tool proxy-tunnel SPN_CONFIG_FILE [flags]
```

**Options**

```
  -h, --help   help for proxy-tunnel
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network tool](#spellshape-network-tool)	 - Commands to run subsidiary tools


## spellshape network validator

Show and update a validator profile

**Options**

```
  -h, --help   help for validator
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network](#spellshape-network)	 - Launch a blockchain in production
* [spellshape network validator set](#spellshape-network-validator-set)	 - Set an information in a validator profile
* [spellshape network validator show](#spellshape-network-validator-show)	 - Show a validator profile


## spellshape network validator set

Set an information in a validator profile

**Synopsis**

Validators on Spellshape can set a profile containing a description for the validator.
The validator set command allows to set information for the validator.
The following information can be set:
- details: general information about the validator.
- identity: piece of information to verify identity of the validator with a system like Keybase of Veramo.
- website: website of the validator.
- security: security contact for the validator.


```
spellshape network validator set details|identity|website|security [value] [flags]
```

**Options**

```
      --from string              account name to use for sending transactions to SPN (default "default")
  -h, --help                     help for set
      --home string              home directory used for blockchains
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network validator](#spellshape-network-validator)	 - Show and update a validator profile


## spellshape network validator show

Show a validator profile

```
spellshape network validator show [address] [flags]
```

**Options**

```
  -h, --help   help for show
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network validator](#spellshape-network-validator)	 - Show and update a validator profile


## spellshape network version

Version of the plugin

**Synopsis**

The version of the plugin to use to interact with a chain might be specified by the coordinator.


```
spellshape network version [flags]
```

**Options**

```
  -h, --help   help for version
```

**Options inherited from parent commands**

```
      --local                       Use local SPN network
      --nightly                     Use nightly SPN network
      --spn-faucet-address string   SPN faucet address (default "https://faucet.devnet.spellshape.com:443")
      --spn-node-address string     SPN node address (default "https://rpc.devnet.spellshape.com:443")
```

**SEE ALSO**

* [spellshape network](#spellshape-network)	 - Launch a blockchain in production


## spellshape node

Make requests to a live blockchain node

**Options**

```
  -h, --help          help for node
      --node string   <host>:<port> to tendermint rpc interface for this chain (default "https://rpc.cosmos.network:443")
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape node query](#spellshape-node-query)	 - Querying subcommands
* [spellshape node tx](#spellshape-node-tx)	 - Transactions subcommands


## spellshape node query

Querying subcommands

**Options**

```
  -h, --help   help for query
```

**Options inherited from parent commands**

```
      --node string   <host>:<port> to tendermint rpc interface for this chain (default "https://rpc.cosmos.network:443")
```

**SEE ALSO**

* [spellshape node](#spellshape-node)	 - Make requests to a live blockchain node
* [spellshape node query bank](#spellshape-node-query-bank)	 - Querying commands for the bank module
* [spellshape node query tx](#spellshape-node-query-tx)	 - Query for transaction by hash


## spellshape node query bank

Querying commands for the bank module

**Options**

```
  -h, --help   help for bank
```

**Options inherited from parent commands**

```
      --node string   <host>:<port> to tendermint rpc interface for this chain (default "https://rpc.cosmos.network:443")
```

**SEE ALSO**

* [spellshape node query](#spellshape-node-query)	 - Querying subcommands
* [spellshape node query bank balances](#spellshape-node-query-bank-balances)	 - Query for account balances by account name or address


## spellshape node query bank balances

Query for account balances by account name or address

```
spellshape node query bank balances [from_account_or_address] [flags]
```

**Options**

```
      --address-prefix string    account address prefix (default "cosmos")
      --count-total              count total number of records in all balances to query for
  -h, --help                     help for balances
      --home string              directory where the blockchain node is initialized
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --limit uint               pagination limit of all balances to query for (default 100)
      --offset uint              pagination offset of all balances to query for
      --page uint                pagination page of all balances to query for. This sets offset to a multiple of limit (default 1)
      --page-key string          pagination page-key of all balances to query for
      --reverse                  results are sorted in descending order
```

**Options inherited from parent commands**

```
      --node string   <host>:<port> to tendermint rpc interface for this chain (default "https://rpc.cosmos.network:443")
```

**SEE ALSO**

* [spellshape node query bank](#spellshape-node-query-bank)	 - Querying commands for the bank module


## spellshape node query tx

Query for transaction by hash

```
spellshape node query tx [hash] [flags]
```

**Options**

```
  -h, --help   help for tx
```

**Options inherited from parent commands**

```
      --node string   <host>:<port> to tendermint rpc interface for this chain (default "https://rpc.cosmos.network:443")
```

**SEE ALSO**

* [spellshape node query](#spellshape-node-query)	 - Querying subcommands


## spellshape node tx

Transactions subcommands

**Options**

```
      --address-prefix string    account address prefix (default "cosmos")
      --fees string              fees to pay along with transaction; eg: 10uatom
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default "auto")
      --gas-prices string        gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            build an unsigned transaction and write it to STDOUT
  -h, --help                     help for tx
      --home string              directory where the blockchain node is initialized
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**Options inherited from parent commands**

```
      --node string   <host>:<port> to tendermint rpc interface for this chain (default "https://rpc.cosmos.network:443")
```

**SEE ALSO**

* [spellshape node](#spellshape-node)	 - Make requests to a live blockchain node
* [spellshape node tx bank](#spellshape-node-tx-bank)	 - Bank transaction subcommands


## spellshape node tx bank

Bank transaction subcommands

**Options**

```
  -h, --help   help for bank
```

**Options inherited from parent commands**

```
      --address-prefix string    account address prefix (default "cosmos")
      --fees string              fees to pay along with transaction; eg: 10uatom
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default "auto")
      --gas-prices string        gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            build an unsigned transaction and write it to STDOUT
      --home string              directory where the blockchain node is initialized
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "https://rpc.cosmos.network:443")
```

**SEE ALSO**

* [spellshape node tx](#spellshape-node-tx)	 - Transactions subcommands
* [spellshape node tx bank send](#spellshape-node-tx-bank-send)	 - Send funds from one account to another.


## spellshape node tx bank send

Send funds from one account to another.

```
spellshape node tx bank send [from_account_or_address] [to_account_or_address] [amount] [flags]
```

**Options**

```
  -h, --help   help for send
```

**Options inherited from parent commands**

```
      --address-prefix string    account address prefix (default "cosmos")
      --fees string              fees to pay along with transaction; eg: 10uatom
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default "auto")
      --gas-prices string        gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom)
      --generate-only            build an unsigned transaction and write it to STDOUT
      --home string              directory where the blockchain node is initialized
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "https://rpc.cosmos.network:443")
```

**SEE ALSO**

* [spellshape node tx bank](#spellshape-node-tx-bank)	 - Bank transaction subcommands


## spellshape plugin

Handle plugins

**Options**

```
  -h, --help   help for plugin
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape plugin add](#spellshape-plugin-add)	 - Adds a plugin declaration to a plugin configuration
* [spellshape plugin describe](#spellshape-plugin-describe)	 - Output information about the a registered plugin
* [spellshape plugin list](#spellshape-plugin-list)	 - List declared plugins and status
* [spellshape plugin remove](#spellshape-plugin-remove)	 - Removes a plugin declaration from a chain's plugin configuration
* [spellshape plugin scaffold](#spellshape-plugin-scaffold)	 - Scaffold a new plugin
* [spellshape plugin update](#spellshape-plugin-update)	 - Update plugins


## spellshape plugin add

Adds a plugin declaration to a plugin configuration

**Synopsis**

Adds a plugin declaration to a plugin configuration.
Respects key value pairs declared after the plugin path to be added to the
generated configuration definition.
Example:
  spellshape plugin add github.com/org/my-plugin/ foo=bar baz=qux

```
spellshape plugin add [path] [key=value]... [flags]
```

**Options**

```
  -g, --global   use global plugins configuration ($HOME/.spellshape/plugins/plugins.yml)
  -h, --help     help for add
```

**SEE ALSO**

* [spellshape plugin](#spellshape-plugin)	 - Handle plugins


## spellshape plugin describe

Output information about the a registered plugin

**Synopsis**

Output information about a registered plugins commands and hooks.

```
spellshape plugin describe [path] [flags]
```

**Options**

```
  -h, --help   help for describe
```

**SEE ALSO**

* [spellshape plugin](#spellshape-plugin)	 - Handle plugins


## spellshape plugin list

List declared plugins and status

**Synopsis**

Prints status and information of declared plugins

```
spellshape plugin list [flags]
```

**Options**

```
  -h, --help   help for list
```

**SEE ALSO**

* [spellshape plugin](#spellshape-plugin)	 - Handle plugins


## spellshape plugin remove

Removes a plugin declaration from a chain's plugin configuration

```
spellshape plugin remove [path] [flags]
```

**Options**

```
  -g, --global   use global plugins configuration ($HOME/.spellshape/plugins/plugins.yml)
  -h, --help     help for remove
```

**SEE ALSO**

* [spellshape plugin](#spellshape-plugin)	 - Handle plugins


## spellshape plugin scaffold

Scaffold a new plugin

**Synopsis**

Scaffolds a new plugin in the current directory with the given repository path configured. A git repository will be created with the given module name, unless the current directory is already a git repository.

```
spellshape plugin scaffold [github.com/org/repo] [flags]
```

**Options**

```
  -h, --help   help for scaffold
```

**SEE ALSO**

* [spellshape plugin](#spellshape-plugin)	 - Handle plugins


## spellshape plugin update

Update plugins

**Synopsis**

Updates a plugin specified by path. If no path is specified all declared plugins are updated

```
spellshape plugin update [path] [flags]
```

**Options**

```
  -h, --help   help for update
```

**SEE ALSO**

* [spellshape plugin](#spellshape-plugin)	 - Handle plugins


## spellshape relayer

Connect blockchains with an IBC relayer

**Options**

```
  -h, --help   help for relayer
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape relayer configure](#spellshape-relayer-configure)	 - Configure source and target chains for relaying
* [spellshape relayer connect](#spellshape-relayer-connect)	 - Link chains associated with paths and start relaying tx packets in between


## spellshape relayer configure

Configure source and target chains for relaying

```
spellshape relayer configure [flags]
```

**Options**

```
  -a, --advanced                  advanced configuration options for custom IBC modules
  -h, --help                      help for configure
      --keyring-backend string    keyring backend to store your account keys (default "test")
      --keyring-dir string        accounts keyring directory (default "/home/runner/.spellshape/accounts")
      --ordered                   set the channel as ordered
  -r, --reset                     reset the relayer config
      --source-account string     source Account
      --source-client-id string   use a custom client id for source
      --source-faucet string      faucet address of the source chain
      --source-gaslimit int       gas limit used for transactions on source chain
      --source-gasprice string    gas price used for transactions on source chain
      --source-port string        IBC port ID on the source chain
      --source-prefix string      address prefix of the source chain
      --source-rpc string         RPC address of the source chain
      --source-version string     module version on the source chain
      --target-account string     target Account
      --target-client-id string   use a custom client id for target
      --target-faucet string      faucet address of the target chain
      --target-gaslimit int       gas limit used for transactions on target chain
      --target-gasprice string    gas price used for transactions on target chain
      --target-port string        IBC port ID on the target chain
      --target-prefix string      address prefix of the target chain
      --target-rpc string         RPC address of the target chain
      --target-version string     module version on the target chain
```

**SEE ALSO**

* [spellshape relayer](#spellshape-relayer)	 - Connect blockchains with an IBC relayer


## spellshape relayer connect

Link chains associated with paths and start relaying tx packets in between

```
spellshape relayer connect [<path>,...] [flags]
```

**Options**

```
  -h, --help                     help for connect
      --keyring-backend string   keyring backend to store your account keys (default "test")
      --keyring-dir string       accounts keyring directory (default "/home/runner/.spellshape/accounts")
```

**SEE ALSO**

* [spellshape relayer](#spellshape-relayer)	 - Connect blockchains with an IBC relayer


## spellshape scaffold

Create a new blockchain, module, message, query, and more

**Synopsis**

Scaffolding is a quick way to generate code for major pieces of your
application.

For details on each scaffolding target (chain, module, message, etc.) run the
corresponding command with a "--help" flag, for example, "spellshape scaffold chain
--help".

The Spellshape team strongly recommends committing the code to a version control
system before running scaffolding commands. This will make it easier to see the
changes to the source code as well as undo the command if you've decided to roll
back the changes.

This blockchain you create with the chain scaffolding command uses the modular
Cosmos SDK framework and imports many standard modules for functionality like
proof of stake, token transfer, inter-blockchain connectivity, governance, and
more. Custom functionality is implemented in modules located by convention in
the "x/" directory. By default, your blockchain comes with an empty custom
module. Use the module scaffolding command to create an additional module.

An empty custom module doesn't do much, it's basically a container for logic
that is responsible for processing transactions and changing the application
state. Cosmos SDK blockchains work by processing user-submitted signed
transactions, which contain one or more messages. A message contains data that
describes a state transition. A module can be responsible for handling any
number of messages.

A message scaffolding command will generate the code for handling a new type of
Cosmos SDK message. Message fields describe the state transition that the
message is intended to produce if processed without errors.

Scaffolding messages is useful to create individual "actions" that your module
can perform. Sometimes, however, you want your blockchain to have the
functionality to create, read, update and delete (CRUD) instances of a
particular type. Depending on how you want to store the data there are three
commands that scaffold CRUD functionality for a type: list, map, and single.
These commands create four messages (one for each CRUD action), and the logic to
add, delete, and fetch the data from the store. If you want to scaffold only the
logic, for example, you've decided to scaffold messages separately, you can do
that as well with the "--no-message" flag.

Reading data from a blockchain happens with a help of queries. Similar to how
you can scaffold messages to write data, you can scaffold queries to read the
data back from your blockchain application.

You can also scaffold a type, which just produces a new protocol buffer file
with a proto message description. Note that proto messages produce (and
correspond with) Go types whereas Cosmos SDK messages correspond to proto "rpc"
in the "Msg" service.

If you're building an application with custom IBC logic, you might need to
scaffold IBC packets. An IBC packet represents the data sent from one blockchain
to another. You can only scaffold IBC packets in IBC-enabled modules scaffolded
with an "--ibc" flag. Note that the default module is not IBC-enabled.


**Options**

```
  -h, --help   help for scaffold
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape scaffold chain](#spellshape-scaffold-chain)	 - New Cosmos SDK blockchain
* [spellshape scaffold list](#spellshape-scaffold-list)	 - CRUD for data stored as an array
* [spellshape scaffold map](#spellshape-scaffold-map)	 - CRUD for data stored as key-value pairs
* [spellshape scaffold message](#spellshape-scaffold-message)	 - Message to perform state transition on the blockchain
* [spellshape scaffold module](#spellshape-scaffold-module)	 - Custom Cosmos SDK module
* [spellshape scaffold packet](#spellshape-scaffold-packet)	 - Message for sending an IBC packet
* [spellshape scaffold query](#spellshape-scaffold-query)	 - Query for fetching data from a blockchain
* [spellshape scaffold react](#spellshape-scaffold-react)	 - React web app template
* [spellshape scaffold single](#spellshape-scaffold-single)	 - CRUD for data stored in a single location
* [spellshape scaffold type](#spellshape-scaffold-type)	 - Type definition
* [spellshape scaffold vue](#spellshape-scaffold-vue)	 - Vue 3 web app template


## spellshape scaffold chain

New Cosmos SDK blockchain

**Synopsis**

Create a new application-specific Cosmos SDK blockchain.

For example, the following command will create a blockchain called "hello" in
the "hello/" directory:

	spellshape scaffold chain hello

A project name can be a simple name or a URL. The name will be used as the Go
module path for the project. Examples of project names:

	spellshape scaffold chain foo
	spellshape scaffold chain foo/bar
	spellshape scaffold chain example.org/foo
	spellshape scaffold chain github.com/username/foo
		
A new directory with source code files will be created in the current directory.
To use a different path use the "--path" flag.

Most of the logic of your blockchain is written in custom modules. Each module
effectively encapsulates an independent piece of functionality. Following the
Cosmos SDK convention, custom modules are stored inside the "x/" directory. By
default, Spellshape creates a module with a name that matches the name of the
project. To create a blockchain without a default module use the "--no-module"
flag. Additional modules can be added after a project is created with "spellshape
scaffold module" command.

Account addresses on Cosmos SDK-based blockchains have string prefixes. For
example, the Cosmos Hub blockchain uses the default "cosmos" prefix, so that
addresses look like this: "cosmos12fjzdtqfrrve7zyg9sv8j25azw2ua6tvu07ypf". To
use a custom address prefix use the "--address-prefix" flag. For example:

	spellshape scaffold chain foo --address-prefix bar

By default when compiling a blockchain's source code Spellshape creates a cache to
speed up the build process. To clear the cache when building a blockchain use
the "--clear-cache" flag. It is very unlikely you will ever need to use this
flag.

The blockchain is using the Cosmos SDK modular blockchain framework. Learn more
about Cosmos SDK on https://docs.cosmos.network


```
spellshape scaffold chain [name] [flags]
```

**Options**

```
      --address-prefix string   account address prefix (default "cosmos")
      --clear-cache             clear the build cache (advanced)
  -h, --help                    help for chain
      --no-module               create a project without a default module
  -p, --path string             create a project in a specific path (default ".")
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold list

CRUD for data stored as an array

**Synopsis**

The "list" scaffolding command is used to generate files that implement the
logic for storing and interacting with data stored as a list in the blockchain
state.

The command accepts a NAME argument that will be used as the name of a new type
of data. It also accepts a list of FIELDs that describe the type.

The interaction with the data follows the create, read, updated, and delete
(CRUD) pattern. For each type three Cosmos SDK messages are defined for writing
data to the blockchain: MsgCreate{Name}, MsgUpdate{Name}, MsgDelete{Name}. For
reading data two queries are defined: {Name} and {Name}All. The type, messages,
and queries are defined in the "proto/" directory as protocol buffer messages.
Messages and queries are mounted in the "Msg" and "Query" services respectively.

When messages are handled, the appropriate keeper methods are called. By
convention, the methods are defined in
"x/{moduleName}/keeper/msg_server_{name}.go". Helpful methods for getting,
setting, removing, and appending are defined in the same "keeper" package in
"{name}.go".

The "list" command essentially allows you to define a new type of data and
provides the logic to create, read, update, and delete instances of the type.
For example, let's review a command that generates the code to handle a list of
posts and each post has "title" and "body" fields:

	spellshape scaffold list post title body

This provides you with a "Post" type, MsgCreatePost, MsgUpdatePost,
MsgDeletePost and two queries: Post and PostAll. The compiled CLI, let's say the
binary is "blogd" and the module is "blog", has commands to query the chain (see
"blogd q blog") and broadcast transactions with the messages above (see "blogd
tx blog").

The code generated with the list command is meant to be edited and tailored to
your application needs. Consider the code to be a "skeleton" for the actual
business logic you will implement next.

By default, all fields are assumed to be strings. If you want a field of a
different type, you can specify it after a colon ":". The following types are
supported: string, bool, int, uint, coin, array.string, array.int, array.uint,
array.coin. An example of using field types:

	spellshape scaffold list pool amount:coin tags:array.string height:int

Supported types:

| Type         | Alias   | Index | Code Type | Description                     |
| ------------ | ------- | ----- | --------- | ------------------------------- |
| string       | -       | yes   | string    | Text type                       |
| array.string | strings | no    | []string  | List of text type               |
| bool         | -       | yes   | bool      | Boolean type                    |
| int          | -       | yes   | int32     | Integer type                    |
| array.int    | ints    | no    | []int32   | List of integers types          |
| uint         | -       | yes   | uint64    | Unsigned integer type           |
| array.uint   | uints   | no    | []uint64  | List of unsigned integers types |
| coin         | -       | no    | sdk.Coin  | Cosmos SDK coin type            |
| array.coin   | coins   | no    | sdk.Coins | List of Cosmos SDK coin types   |

"Index" indicates whether the type can be used as an index in
"spellshape scaffold map".

Spellshape also supports custom types:

	spellshape scaffold list product-details name desc
	spellshape scaffold list product price:coin details:ProductDetails

In the example above the "ProductDetails" type was defined first, and then used
as a custom type for the "details" field. Spellshape doesn't support arrays of
custom types yet.

Your chain will accept custom types in JSON-notation:

	exampled tx example create-product 100coin '{"name": "x", "desc": "y"}' --from alice

By default the code will be scaffolded in the module that matches your project's
name. If you have several modules in your project, you might want to specify a
different module:

	spellshape scaffold list post title body --module blog

By default, each message comes with a "creator" field that represents the
address of the transaction signer. You can customize the name of this field with
a flag:

	spellshape scaffold list post title body --signer author

It's possible to scaffold just the getter/setter logic without the CRUD
messages. This is useful when you want the methods to handle a type, but would
like to scaffold messages manually. Use a flag to skip message scaffolding:

	spellshape scaffold list post title body --no-message

The "creator" field is not generated if a list is scaffolded with the
"--no-message" flag.


```
spellshape scaffold list NAME [field]... [flags]
```

**Options**

```
      --clear-cache     clear the build cache (advanced)
  -h, --help            help for list
      --module string   specify which module to generate code in
      --no-message      skip generating message handling logic
      --no-simulation   skip simulation logic
  -p, --path string     path of the app (default ".")
      --signer string   label for the message signer (default: creator)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold map

CRUD for data stored as key-value pairs

**Synopsis**

The "map" scaffolding command is used to generate files that implement the logic
for storing and interacting with data stored as key-value pairs (or a
dictionary) in the blockchain state.

The "map" command is very similar to "spellshape scaffold list" with the main
difference in how values are indexed. With "list" values are indexed by an
incrementing integer, whereas "list" values are indexed by a user-provided value
(or multiple values).

Let's use the same blog post example:

	spellshape scaffold map post title body

This command scaffolds a "Post" type and CRUD functionality to create, read,
updated, and delete posts. However, when creating a new post with your chain's
binary (or by submitting a transaction through the chain's API) you will be
required to provide an "index":

	blogd tx blog create-post [index] [title] [body]
	blogd tx blog create-post hello "My first post" "This is the body"

This command will create a post and store it in the blockchain's state under the
"hello" index. You will be able to fetch back the value of the post by querying
for the "hello" key.

	blogd q blog show-post hello

To customize the index, use the "--index" flag. Multiple indices can be
provided, which simplifies querying values. For example:

	spellshape scaffold map product price desc --index category,guid

With this command, you would get a "Product" value indexed by both a category
and a GUID (globally unique ID). This will let you programmatically fetch
product values that have the same category but are using different GUIDs.

Since the behavior of "list" and "map" scaffolding is very similar, you can use
the "--no-message", "--module", "--signer" flags as well as the colon syntax for
custom types.


```
spellshape scaffold map NAME [field]... [flags]
```

**Options**

```
      --clear-cache     clear the build cache (advanced)
  -h, --help            help for map
      --index strings   fields that index the value (default [index])
      --module string   specify which module to generate code in
      --no-message      skip generating message handling logic
      --no-simulation   skip simulation logic
  -p, --path string     path of the app (default ".")
      --signer string   label for the message signer (default: creator)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold message

Message to perform state transition on the blockchain

**Synopsis**

Message scaffolding is useful for quickly adding functionality to your
blockchain to handle specific Cosmos SDK messages.

Messages are objects whose end goal is to trigger state transitions on the
blockchain. A message is a container for fields of data that affect how the
blockchain's state will change. You can think of messages as "actions" that a
user can perform.

For example, the bank module has a "Send" message for token transfers between
accounts. The send message has three fields: from address (sender), to address
(recipient), and a token amount. When this message is successfully processed,
the token amount will be deducted from the sender's account and added to the
recipient's account.

Spellshape's message scaffolding lets you create new types of messages and add them
to your chain. For example:

	spellshape scaffold message add-pool amount:coins denom active:bool --module dex

The command above will create a new message MsgAddPool with three fields: amount
(in tokens), denom (a string), and active (a boolean). The message will be added
to the "dex" module.

By default, the message is defined as a proto message in the
"proto/{app}/{module}/tx.proto" and registered in the "Msg" service. A CLI command to
create and broadcast a transaction with MsgAddPool is created in the module's
"cli" package. Additionally, Spellshape scaffolds a message constructor and the code
to satisfy the sdk.Msg interface and register the message in the module.

Most importantly in the "keeper" package Spellshape scaffolds an "AddPool" function.
Inside this function, you can implement message handling logic.

When successfully processed a message can return data. Use the —response flag to
specify response fields and their types. For example

	spellshape scaffold message create-post title body --response id:int,title

The command above will scaffold MsgCreatePost which returns both an ID (an
integer) and a title (a string).

Message scaffolding follows the rules as "spellshape scaffold list/map/single" and
supports fields with standard and custom types. See "spellshape scaffold list —help"
for details.


```
spellshape scaffold message [name] [field1] [field2] ... [flags]
```

**Options**

```
      --clear-cache        clear the build cache (advanced)
  -d, --desc string        description of the command
  -h, --help               help for message
      --module string      module to add the message into. Default: app's main module
      --no-simulation      disable CRUD simulation scaffolding
  -p, --path string        path of the app (default ".")
  -r, --response strings   response fields
      --signer string      label for the message signer (default: creator)
  -y, --yes                answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold module

Custom Cosmos SDK module

**Synopsis**

Scaffold a new Cosmos SDK module.

Cosmos SDK is a modular framework and each independent piece of functionality is
implemented in a separate module. By default your blockchain imports a set of
standard Cosmos SDK modules. To implement custom functionality of your
blockchain, scaffold a module and implement the logic of your application.

This command does the following:

* Creates a directory with module's protocol buffer files in "proto/"
* Creates a directory with module's boilerplate Go code in "x/"
* Imports the newly created module by modifying "app/app.go"
* Creates a file in "testutil/keeper/" that contains logic to create a keeper
  for testing purposes

This command will proceed with module scaffolding even if "app/app.go" doesn't
have the required default placeholders. If the placeholders are missing, you
will need to modify "app/app.go" manually to import the module. If you want the
command to fail if it can't import the module, use the "--require-registration"
flag.

To scaffold an IBC-enabled module use the "--ibc" flag. An IBC-enabled module is
like a regular module with the addition of IBC-specific logic and placeholders
to scaffold IBC packets with "spellshape scaffold packet".

A module can depend on one or more other modules and import their keeper
methods. To scaffold a module with a dependency use the "--dep" flag

For example, your new custom module "foo" might have functionality that requires
sending tokens between accounts. The method for sending tokens is a defined in
the "bank"'s module keeper. You can scaffold a "foo" module with the dependency
on "bank" with the following command:

	spellshape scaffold module foo --dep bank

You can then define which methods you want to import from the "bank" keeper in
"expected_keepers.go".

You can also scaffold a module with a list of dependencies that can include both
standard and custom modules (provided they exist):

	spellshape scaffold module bar --dep foo,mint,account,FeeGrant

Note: the "--dep" flag doesn't install third-party modules into your
application, it just generates extra code that specifies which existing modules
your new custom module depends on.

A Cosmos SDK module can have parameters (or "params"). Params are values that
can be set at the genesis of the blockchain and can be modified while the
blockchain is running. An example of a param is "Inflation rate change" of the
"mint" module. A module can be scaffolded with params using the "--params" flag
that accepts a list of param names. By default params are of type "string", but
you can specify a type for each param. For example:

	spellshape scaffold module foo --params baz:uint,bar:bool

Refer to Cosmos SDK documentation to learn more about modules, dependencies and
params.


```
spellshape scaffold module [name] [flags]
```

**Options**

```
      --clear-cache            clear the build cache (advanced)
      --dep strings            add a dependency on another module
  -h, --help                   help for module
      --ibc                    add IBC functionality
      --ordering string        channel ordering of the IBC module [none|ordered|unordered] (default "none")
      --params strings         add module parameters
  -p, --path string            path of the app (default ".")
      --require-registration   fail if module can't be registered
  -y, --yes                    answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold packet

Message for sending an IBC packet

**Synopsis**

Scaffold an IBC packet in a specific IBC-enabled Cosmos SDK module

```
spellshape scaffold packet [packetName] [field1] [field2] ... --module [moduleName] [flags]
```

**Options**

```
      --ack strings     custom acknowledgment type (field1,field2,...)
      --clear-cache     clear the build cache (advanced)
  -h, --help            help for packet
      --module string   IBC Module to add the packet into
      --no-message      disable send message scaffolding
  -p, --path string     path of the app (default ".")
      --signer string   label for the message signer (default: creator)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold query

Query for fetching data from a blockchain

```
spellshape scaffold query [name] [request_field1] [request_field2] ... [flags]
```

**Options**

```
      --clear-cache        clear the build cache (advanced)
  -d, --desc string        description of the CLI to broadcast a tx with the message
  -h, --help               help for query
      --module string      module to add the query into. Default: app's main module
      --paginated          define if the request can be paginated
  -p, --path string        path of the app (default ".")
  -r, --response strings   response fields
  -y, --yes                answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold react

React web app template

```
spellshape scaffold react [flags]
```

**Options**

```
  -h, --help          help for react
  -p, --path string   path to scaffold content of the React app (default "./react")
  -y, --yes           answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold single

CRUD for data stored in a single location

```
spellshape scaffold single NAME [field]... [flags]
```

**Options**

```
      --clear-cache     clear the build cache (advanced)
  -h, --help            help for single
      --module string   specify which module to generate code in
      --no-message      skip generating message handling logic
      --no-simulation   skip simulation logic
  -p, --path string     path of the app (default ".")
      --signer string   label for the message signer (default: creator)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold type

Type definition

```
spellshape scaffold type NAME [field]... [flags]
```

**Options**

```
      --clear-cache     clear the build cache (advanced)
  -h, --help            help for type
      --module string   specify which module to generate code in
      --no-message      skip generating message handling logic
      --no-simulation   skip simulation logic
  -p, --path string     path of the app (default ".")
      --signer string   label for the message signer (default: creator)
  -y, --yes             answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape scaffold vue

Vue 3 web app template

```
spellshape scaffold vue [flags]
```

**Options**

```
  -h, --help          help for vue
  -p, --path string   path to scaffold content of the Vue.js app (default "./vue")
  -y, --yes           answers interactive yes/no questions with yes
```

**SEE ALSO**

* [spellshape scaffold](#spellshape-scaffold)	 - Create a new blockchain, module, message, query, and more


## spellshape tools

Tools for advanced users

**Options**

```
  -h, --help   help for tools
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain
* [spellshape tools ibc-relayer](#spellshape-tools-ibc-relayer)	 - TypeScript implementation of an IBC relayer
* [spellshape tools ibc-setup](#spellshape-tools-ibc-setup)	 - Collection of commands to quickly setup a relayer
* [spellshape tools protoc](#spellshape-tools-protoc)	 - Execute the protoc command


## spellshape tools ibc-relayer

TypeScript implementation of an IBC relayer

```
spellshape tools ibc-relayer [--] [...] [flags]
```

**Examples**

```
spellshape tools ibc-relayer -- -h
```

**Options**

```
  -h, --help   help for ibc-relayer
```

**SEE ALSO**

* [spellshape tools](#spellshape-tools)	 - Tools for advanced users


## spellshape tools ibc-setup

Collection of commands to quickly setup a relayer

```
spellshape tools ibc-setup [--] [...] [flags]
```

**Examples**

```
spellshape tools ibc-setup -- -h
spellshape tools ibc-setup -- init --src relayer_test_1 --dest relayer_test_2
```

**Options**

```
  -h, --help   help for ibc-setup
```

**SEE ALSO**

* [spellshape tools](#spellshape-tools)	 - Tools for advanced users


## spellshape tools protoc

Execute the protoc command

**Synopsis**

The protoc command. You don't need to setup the global protoc include folder with -I, it's automatically handled

```
spellshape tools protoc [--] [...] [flags]
```

**Examples**

```
spellshape tools protoc -- --version
```

**Options**

```
  -h, --help   help for protoc
```

**SEE ALSO**

* [spellshape tools](#spellshape-tools)	 - Tools for advanced users


## spellshape version

Print the current build information

```
spellshape version [flags]
```

**Options**

```
  -h, --help   help for version
```

**SEE ALSO**

* [spellshape](#spellshape)	 - Spellshape CLI offers everything you need to scaffold, test, build, and launch your blockchain

