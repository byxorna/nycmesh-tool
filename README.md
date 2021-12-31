# nycmesh-tool
nycmesh-tool CLI

# Build

```
$ make
$ ./bin/nycmesh-tool -h
```

# UISP API Commands

By default (`--config`) we read `~/.nycmesh-tool.yaml` for global flags. You could pass these parameters on the CLI, or load them in a dotfile to save keystrokes. Create your `~/.nycmesh-tool.yaml` like the following:

```yaml
x-auth-token: xxx
hostname: uisp-controller.hostname.xxx
scheme: https
```

Now, invoke the tool:

```
./bin/nycmesh-tool uisp authorization getUser
```
