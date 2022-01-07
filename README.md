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

# Get It

See [releases](https://github.com/byxorna/nycmesh-tool/releases)

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
  x-auth-token: xxx # get this from https://uisp.mesh/nms/user/login or scripts/uisp-user-token.sh
  hostname: uisp.mesh
  scheme: https
  skip-verify-tls: false
  debug: true
```

# Hacking

See [./HACKING.md](HACKING.md)

# Releases

See [./RELEASE.md](RELEASE.md)
