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

# Adding new commands

Create the `cmd/` entry, for flag processing and invoking the application.

```
$ cobra add set-some-feature
$ vi cmd/setSomeFeature.go
...
```

Business logic lives in `app/`, command line flag processing lives in `cmd/`.

# Releases

See [./RELEASES.md](RELEASES.md)
