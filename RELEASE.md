# Releases

## Create a release

Releases are largely automated. By tagging a commit like `v0.1.2` and pushing it, [Github Actions](https://github.com/byxorna/nycmesh-tool/actions) will create the changelog, build binaries, and attach them to the release documentation.

All that is needed is for you either:

- Create a release at https://github.com/byxorna/nycmesh-tool/releases

_or..., if you are more manually inclined:_

- `git tag -a v0.3.5`
- `git push origin --tags`
