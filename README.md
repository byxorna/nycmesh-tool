# nycmesh-tool

nycmesh-tool CLI

# Features

At the moment, the tool is pretty sparse. It provides the top level `nycmesh-tool` command, with subcommands for:

- `uisp` - Full UISP API CLI
- `meshapi` - Interact with https://github.com/meshcenter/mesh-api
  - `node` - Display nodes from https://api.nycmesh.net/v1/nodes
  - `device` - Display devices from https://api.nycmesh.net/v1/nodes
- `map` - Open up nodes on a map from the command line
- `cache` - manipulate a simple on-disk cache for things like API responses, for offline access to inventory
- `experiment` - Experimental commands. Here be dragons!
  - `devices` - display fused device data, created by joining UISP data with mesh-api data. Useful for further `jq` processing.
- `watch`
  - `logs` - easy interface to watch UISP logs in real(ish) time

# Releases

See [releases](https://github.com/byxorna/nycmesh-tool/releases) for binary downloads, and [./RELEASE.md](RELEASE.md) for more information.

# Hacking

See [./HACKING.md](HACKING.md)

# Config

By default (`--config`) we read `.nycmesh-tool.yaml` for global flags for the tool. You could pass these parameters on the CLI, or store them in `~/.nycmesh-tool.yaml` (or `.`):

```yaml
---
uisp:
  x-auth-token: xxx # get this from https://uisp.mesh/nms/user/login or scripts/uisp-user-token.sh
  hostname: uisp.mesh
  scheme: https
  skip-verify-tls: false # true needed for self-signed certs
  debug: true
```

## UISP API Commands

To test your connection to UISP works:

```
$ ./bin/nycmesh-tool uisp authorization getUser
{"userId":null,"username":"gabeconradi"}
```

Or, see how many devices are in the API:

```
$ ./bin/nycmesh-tool uisp devices getDevices --hostname=uisp.mesh --scheme=https  --debug --skip-verify-tls |jq length
1293
```


