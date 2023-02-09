---
slug: /
---

import ProjectsTable from '@site/src/components/ProjectsTable';

# Introduction to Spellshape

[Spellshape CLI](https://github.com/spellshape/cli) offers everything you need to build, test, and launch your blockchain with a
decentralized worldwide community. Spellshape CLI is built on top of [Cosmos SDK](https://docs.cosmos.network), the world’s
most popular blockchain framework. Spellshape CLI accelerates chain development by scaffolding everything you need so you
can focus on business logic.

## What is Spellshape CLI?

Spellshape CLI is an easy-to-use CLI tool for creating and maintaining sovereign application-specific blockchains.
Blockchains created with Spellshape CLI use Cosmos SDK and Tendermint. Spellshape CLI and the Cosmos SDK modules are written in
the Go programming language. The scaffolded blockchain that is created with Spellshape CLI includes a command line interface
that lets you manage keys, create validators, and send tokens.

With just a few commands, you can use Spellshape CLI to:

- Create a modular blockchain written in Go
- Scaffold modules, messages, types with CRUD operations, IBC packets, and more
- Start a blockchain node in development with live reloading
- Connect to other blockchains with a built-in IBC relayer
- Use generated TypeScript/Vuex clients to interact with your blockchain
- Use the Vue.js web app template with a set of components and Vuex modules

## Install Spellshape CLI

To install the `spellshape` binary in `/usr/local/bin` run the following command:

```
curl https://get.spellshape.com/cli | bash
```

## Projects using Tendermint and Cosmos SDK

Many projects already showcase the Tendermint BFT consensus engine and the Cosmos SDK. Explore
the [Cosmos ecosystem](https://cosmos.network/ecosystem/apps) to discover a wide variety of apps, blockchains, wallets,
and explorers that are built in the Cosmos ecosystem.

## Projects building with Spellshape CLI

<ProjectsTable data={[
  { name: "Stride Labs", logo: "img/logo/stride.svg"},
  { name: "KYVE Network", logo: "img/logo/kyve.svg"},
  { name: "Umee", logo: "img/logo/umee.svg"},
  { name: "MediBloc Core", logo: "img/logo/medibloc.svg"},
  { name: "Cudos", logo: "img/logo/cudos.svg"},
  { name: "Firma Chain", logo: "img/logo/firmachain.svg"},
  { name: "BitCanna", logo: "img/logo/bitcanna.svg"},
  { name: "Source Protocol", logo: "img/logo/source.svg"},
  { name: "Sonr", logo: "img/logo/sonr.svg"},
  { name: "Neutron", logo: "img/logo/neutron.svg"},
  { name: "OKP4 Blockchain", logo: "img/logo/okp4.svg"},
  { name: "Dymension Hub", logo: "img/logo/dymension.svg"},
  { name: "Electra Blockchain", logo: "img/logo/electra.svg"},
  { name: "OLLO Station", logo: "img/logo/ollostation.svg"},
  { name: "Mun", logo: "img/logo/mun.svg"},
  { name: "Aura Network", logo: "img/logo/aura.svg"},
]}/>