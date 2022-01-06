# nycmesh-tool

nycmesh-tool CLI


# Features

At the moment, the tool is pretty sparse. It provides the top level `nycmesh-tool` command, with subcommands for:

- `uisp` - Full UISP API CLI
- `node` - Display nodes from https://api.nycmesh.net/v1/nodes
- `map` - Open up nodes on a map from the command line
- `cache` - manipulate a simple on-disk cache for things like API responses, for offline access to inventory

Coming soon:

- `setSectorFrequency` - set sector frequency (WIP)
  - suggest device frequency changes based on observed RF

# Build

To build the cli in `bin/`:

```
$ make
$ ./bin/nycmesh-tool -h
```

To cross-compile (i.e. for MacOS build on a linux box), pass `$GOOS` and `$GOARCH` like:

```
$ GOOS=darwin make
$ file bin/nycmesh-tool
bin/nycmesh-tool: Mach-O 64-bit x86_64 executable
```

# UISP API Commands

Invoke the tool:

```
$ ./bin/nycmesh-tool uisp authorization getUser
$ ./bin/nycmesh-tool uisp  devices getDevices --hostname=uisp.mesh --scheme=https  --debug --skip-verify-tls
```



By default (`--config`) we read `.nycmesh-tool.yaml` for global flags for the tool. You could pass these parameters on the CLI, or store them in `~/.nycmesh-tool.yaml` (or `.`):

```yaml
---
uisp:
  x-auth-token: xxx # get this from https://uisp.mesh/nms/user/login
  hostname: uisp.mesh
  scheme: https
  skip-verify-tls: false
  debug: true
```

# Adding new commands

Create the `cmd/` entry, for flag processing and invoking the application.

```
$ cobra add set-some-feature
$ vi cmd/setSomeFeature.go
...
```

Business logic lives in `app/`, command line flag processing lives in `cmd/`.
