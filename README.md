# nycmesh-tool

nycmesh-tool CLI


# Features

At the moment, the tool is pretty sparse. It provides the top level `nycmesh-tool` command, with subcommands for:

- `uisp` - Full UISP API CLI
- `node` - Display nodes from https://api.nycmesh.net/v1/nodes
- `setSectorFrequency` - set sector frequency (WIP)
- `cache` - manipulate a simple on-disk cache for things like nodes

Coming soon:

- suggest device frequency changes based on observed RF

# Build

To build the cli in `bin/`:

```
$ make
$ ./bin/nycmesh-tool -h
```

# UISP API Commands

By default (`--config`) we read `.nycmesh-tool.yaml` for global flags for the tool. You could pass these parameters on the CLI, or store them in `~/.nycmesh-tool.yaml` (or `.`):

```yaml
x-auth-token: xxx
hostname: uisp-controller.hostname.xxx
scheme: https
```

Now, invoke the tool:

```
./bin/nycmesh-tool uisp authorization getUser
```

# Adding new commands

```
$ cobra add set-some-feature
```
