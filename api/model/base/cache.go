package base

import (
	"github.com/go-baa/cache"
	baa "gopkg.in/baa.v1"
)

// GetCacher 获取缓存控制
func GetCacher() cache.Cacher {
	if c := baa.Default().GetDI("cache"); c != nil {
		return c.(cache.Cacher)
	}
	return nil
}
