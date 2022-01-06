#!/bin/bash
for os in linux windows darwin ; do
  GOOS=$os make
  output=bin/nycmesh-tool
  target=bin/nycmesh-tool-$os
  if [[ $os == windows ]] ; then
    target=$target.exe
  fi
  mv $output $target
  echo "$target"
done
