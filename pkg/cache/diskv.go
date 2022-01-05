package cache

import (
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"
	diskv "github.com/peterbourgon/diskv/v3"
)

var (
	cacheItemSuffix = `.json`
)

func KeyTransform(key string) *diskv.PathKey {
	path := strings.Split(key, "/")
	last := len(path) - 1
	return &diskv.PathKey{
		Path:     path[:last],
		FileName: path[last] + cacheItemSuffix,
	}
}

func InverseKeyTransform(pathKey *diskv.PathKey) (key string) {
	txt := pathKey.FileName[len(pathKey.FileName)-4:]
	if txt != cacheItemSuffix {
		panic("Invalid file found in storage folder!")
	}
	return strings.Join(pathKey.Path, "/") + pathKey.FileName[:len(pathKey.FileName)-4]
}

func NewDiskVCache() (*diskv.Diskv, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	d := diskv.New(diskv.Options{
		BasePath:          path.Join(home, cacheDir),
		AdvancedTransform: KeyTransform,
		InverseTransform:  InverseKeyTransform,
		CacheSizeMax:      1024 * 1024 * 1024,
	})
	return d, nil
}
