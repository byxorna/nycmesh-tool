package cache

import (
	"fmt"
	"time"
)

var (
	// TODO: replace with github.com/adrg/xdg.XDG_CACHE_HOME
	cacheDir = ".cache/nycmesh-tool"

	CacheTTL     = time.Hour * 24 * 7
	ErrNotCached = fmt.Errorf("no cache found")
)

type DiskCache interface {
	Write(key string, data []byte) error
	Has(key string) bool
	Read(key string) ([]byte, error)
	Erase(key string) error
	EraseAll() error
}
