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

# Random Notes

## Fixing/extending generated code

This project uses cobra/viper for flag and config processing, as does the goswagger generated code for UISP's CLI (`generated/go/uisp`). To inject our custom hacks into this generated code, see `internal/uisp/cli/custom.go` which is linked in at `generated/go/uisp/cli/custom.go`. 

TODO: there are some manual changes needed to make the generated go-swagger code compile. Document the patchsets, and/or automate them!

- `Sfp+1` and `Sfp1`: the generator templates just drop `+` from identifiers, so this creates a lot of duplicate functions and vars.
