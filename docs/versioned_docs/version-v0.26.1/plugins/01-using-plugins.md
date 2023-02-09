---
description: Using and Developing plugins
---

# Using Plugins

Spellshape plugins offer a way to extend the functionality of the Spellshape CLI. There
are two core concepts within plugins : `Commands` and `Hooks`. Where `Commands`
extend the cli's functionality, and `Hooks` extend existing command
functionality.

Plugins are registered in an Spellshape scaffolded Blockchain project through the
`plugins.yml`, or globally through `$HOME/.spellshape/plugins/plugins.yml`.

To use a plugin within your project, execute the following command inside the
project directory:

```sh
spellshape plugin add github.com/project/cli-plugin
```

The plugin will be available only when running `spellshape` inside the project
directory.

To use a plugin globally on the other hand, execute the following command:

```sh
spellshape plugin add -g github.com/project/cli-plugin
```

The command will compile the plugin and make it immediately available to the
`spellshape` command lists.

## Listing installed plugins

When in an spellshape scaffolded blockchain you can use the command `spellshape plugin
list` to list all plugins and there statuses.

## Updating plugins

When a plugin in a remote repository releases updates, running `spellshape plugin
update <path/to/plugin>` will update a specific plugin declared in your
project's `config.yml`.
