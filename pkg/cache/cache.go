package cache

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/mitchellh/go-homedir"
)

var (
	cacheDir = ".cache/nycmesh-tool"

	CacheTTL     = time.Hour * 24 * 7
	ErrNotCached = fmt.Errorf("no cache found")
)

type DiskCache struct {
	home string
}

func NewDiskCache() (*DiskCache, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	return &DiskCache{
		home: home,
	}, nil
}

func (c *DiskCache) HasValidCache(key string) bool {
	c.EnsureCacheDir()
	f, err := os.Stat(path.Join(c.home, cacheDir, key))
	if f != nil {
		if time.Since(f.ModTime()) > CacheTTL {
			return false
		}
	}
	return err == nil
}

func (c *DiskCache) PopulateCache(key string, contents []byte) error {
	c.EnsureCacheDir()
	log.Printf("writing cache %s (%d bytes)\n", key, len(contents))
	err := os.WriteFile(path.Join(c.home, cacheDir, key), contents, 0770)
	return err
}

func (c *DiskCache) EnsureCacheDir() {
	d := path.Join(c.home, cacheDir)
	if f, err := os.Stat(d); os.IsNotExist(err) {
		log.Printf("creating %s\n", d)
		// create dir
		os.Mkdir(d, 0750)
	} else if f != nil && !f.IsDir() {
		os.Remove(d)
		os.Mkdir(d, 0750)
	}
}

func (c *DiskCache) Load(key string) ([]byte, error) {
	c.EnsureCacheDir()
	p := path.Join(c.home, cacheDir, key)
	log.Printf("loading cache %s", p)
	dat, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return dat, nil
}
func (c *DiskCache) Clear(key string) error {
	p := path.Join(c.home, cacheDir, key)
	err := os.Remove(p)
	if err != nil {
		return fmt.Errorf("error clearing %s: %w", p, err)
	}
	return nil
}

func (c *DiskCache) ClearAll() error {
	p := path.Join(c.home, cacheDir, "*")
	cachedFiles, err := filepath.Glob(p)
	if err != nil {
		return err
	}
	for _, f := range cachedFiles {
		err := os.Remove(f)
		if err != nil {
			return fmt.Errorf("error clearing %s: %w", f, err)
		}
		log.Printf("deleted cache %s", f)
	}
	return nil
}
