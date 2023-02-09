---
description: Using and Developing plugins
---

# Developing Plugins

It's easy to create a plugin and use it immediately in your project. First
choose a directory outside your project and run :

```sh
$ spellshape plugin scaffold my-plugin
```

This will create a new directory `my-plugin` that contains the plugin's code,
and will output some instructions about how to use your plugin with the
`spellshape` command. Indeed, a plugin path can be a local directory, which has
several benefits:

- you don't need to use a git repository during the development of your plugin.
- the plugin is recompiled each time you run the `spellshape` binary in your
project, if the source files are older than the plugin binary.

Thus, the plugin development workflow is as simple as :

1. scaffold a plugin with `spellshape plugin scaffold my-plugin`
2. add it to your config via `spellshape plugin add -g /path/to/my-plugin`
3. update plugin code
4. run `spellshape my-plugin` binary to compile and run the plugin.
5. go back to 3.

Once your plugin is ready, you can publish it to a git repository, and the
community can use it by calling `spellshape plugin add github.com/foo/my-plugin`.

Now let's detail how to update your plugin's code.

## The plugin interface

The `spellshape` plugin system uses `github.com/hashicorp/go-plugin` under the hood,
which implies to implement a predefined interface:

```go title=spellshape/services/plugin/interface.go
// An spellshape plugin must implements the Plugin interface.
type Interface interface {
	// Manifest declares the plugin's Command(s) and Hook(s).
	Manifest() (Manifest, error)

	// Execute will be invoked by spellshape when a plugin Command is executed.
	// It is global for all commands declared in Manifest, if you have declared
	// multiple commands, use cmd.Path to distinguish them.
	Execute(cmd ExecutedCommand) error

	// ExecuteHookPre is invoked by spellshape when a command specified by the Hook
	// path is invoked.
	// It is global for all hooks declared in Manifest, if you have declared
	// multiple hooks, use hook.Name to distinguish them.
	ExecuteHookPre(hook ExecutedHook) error

	// ExecuteHookPost is invoked by spellshape when a command specified by the hook
	// path is invoked.
	// It is global for all hooks declared in Manifest, if you have declared
	// multiple hooks, use hook.Name to distinguish them.
	ExecuteHookPost(hook ExecutedHook) error

	// ExecuteHookCleanUp is invoked by spellshape when a command specified by the
	// hook path is invoked. Unlike ExecuteHookPost, it is invoked regardless of
	// execution status of the command and hooks.
	// It is global for all hooks declared in Manifest, if you have declared
	// multiple hooks, use hook.Name to distinguish them.
	ExecuteHookCleanUp(hook ExecutedHook) error
}
```

The code scaffolded already implements this interface, you just need to update
the methods' body.


## Defining plugin's manifest

Here is the `Manifest` struct :

```go title=spellshape/services/plugin/interface.go
type Manifest struct {
	Name string
	// Commands contains the commands that will be added to the list of spellshape
	// commands. Each commands are independent, for nested commands use the
	// inner Commands field.
	Commands []Command
	// Hooks contains the hooks that will be attached to the existing spellshape
	// commands.
	Hooks []Hook
	// SharedHost enables sharing a single plugin server across all running instances
	// of a plugin. Useful if a plugin adds or extends long running commands
	//
	// Example: if a plugin defines a hook on `spellshape chain serve`, a plugin server is instanciated
	// when the command is run. Now if you want to interact with that instance from commands
	// defined in that plugin, you need to enable `SharedHost`, or else the commands will just
	// instantiate separate plugin servers.
	//
	// When enabled, all plugins of the same `Path` loaded from the same configuration will
	// attach it's rpc client to a an existing rpc server.
	//
	// If a plugin instance has no other running plugin servers, it will create one and it will be the host.
	SharedHost bool `yaml:"shared_host"`
}
```

In your plugin's code, the `Manifest` method already returns a predefined
`Manifest` struct as an example. Adapt it according to your need.

If your plugin adds one or more new commands to `spellshape`, feeds the `Commands`
field.

If your plugin adds features to existing commands, feeds the `Hooks` field.

Of course a plugin can declare `Commands` *and* `Hooks`.

A plugin may also share a host process by setting `SharedHost` to `true`.
`SharedHost` is desirable if a plugin hooks into, or declares long running commands.
Commands executed from the same plugin context interact with the same plugin server. 
Allowing all executing commands to share the same server instance, giving shared execution context.

## Adding new command

Plugin commands are custom commands added to the spellshape cli by a registered
plugin. Commands can be of any path not defined already by spellshape. All plugin
commands will extend of the command root `spellshape`. 

For instance, let's say your plugin adds a new `oracle` command to `spellshape
scaffold`, the `Manifest()` method will look like :

```go
func (p) Manifest() (plugin.Manifest, error) {
	return plugin.Manifest{
		Name: "oracle",
		Commands: []plugin.Command{
			{
				Use:   "oracle [name]",
				Short: "Scaffold an oracle module",
				Long:  "Long description goes here...",
				// Optionnal flags is required
				Flags: []plugin.Flag{
					{Name: "source", Type: plugin.FlagTypeString, Usage: "the oracle source"},
				},
				// Attach the command to `scaffold`
				PlaceCommandUnder: "spellshape scaffold",
			},
		},
	}, nil
}
```

To update the plugin execution, you have to change the plugin `Execute` command,
for instance :

```go
func (p) Execute(cmd plugin.ExecutedCommand) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("oracle name missing")
	}
	var (
		name      = cmd.Args[0]
		source, _ = cmd.Flags().GetString("source")
	)
	// Read chain information
	c, err := getChain(cmd)
	if err != nil {
		return err
	}

	//...
}
```

Then, run `spellshape scaffold oracle` to execute the plugin.

## Adding hooks

Plugin `Hooks` allow existing spellshape commands to be extended with new
functionality. Hooks are useful when you want to streamline functionality
without needing to run custom scripts after or before a command has been run.
this can streamline processes that where once error prone or forgotten all
together.

The following are hooks defined which will run on a registered `spellshape` commands

| Name     | Description                                                                                                                                           |
| -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| Pre      | Runs before a commands main functionality is invoked in the `PreRun` scope                                                                            |
| Post     | Runs after a commands main functionality is invoked in the `PostRun` scope                                                                            |
| Clean Up | Runs after a commands main functionality is invoked. if the command returns an error it will run before the error is returned to guarantee execution. |

*Note*: If a hook causes an error in the pre step the command will not run
resulting in `post` and `clean up` not executing.

The following is an example of a `hook` definition.

```go
func (p) Manifest() (plugin.Manifest, error) {
	return plugin.Manifest{
		Name: "oracle",
		Hooks: []plugin.Hook{
			{
				Name:        "my-hook",
				PlaceHookOn: "spellshape chain build",
			},
		},
	}, nil
}

func (p) ExecuteHookPre(hook plugin.ExecutedHook) error {
	switch hook.Name {
	case "my-hook":
		fmt.Println("I'm executed before spellshape chain build")
	default:
		return fmt.Errorf("hook not defined")
	}
	return nil
}

func (p) ExecuteHookPost(hook plugin.ExecutedHook) error {
	switch hook.Name {
	case "my-hook":
		fmt.Println("I'm executed after spellshape chain build (if no error)")
	default:
		return fmt.Errorf("hook not defined")
	}
	return nil
}

func (p) ExecuteHookCleanUp(hook plugin.ExecutedHook) error {
	switch hook.Name {
	case "my-hook":
		fmt.Println("I'm executed after spellshape chain build (regardless errors)")
	default:
		return fmt.Errorf("hook not defined")
	}
	return nil
}
```

Above we can see a similar definition to `Command` where a hook has a `Name` and
a `PlaceHookOn`. You'll notice that the `Execute*` methods map directly to each
life cycle of the hook. All hooks defined within the plugin will invoke these
methods.
