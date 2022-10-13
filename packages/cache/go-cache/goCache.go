package gocacheCache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type GoCache interface {
	Get(key string) interface{}
	Set(key string, data interface{}, expiration time.Duration)
	Find(key string) bool
}

type Options struct {
	cache *cache.Cache
}

func NewGoCache() GoCache {
	c := cache.New(5*time.Minute, 15*time.Minute)

	opt := new(Options)
	opt.cache = c

	return opt
}

func (opt *Options) Get(key string) interface{} {
	value, _ := opt.cache.Get(key)
	return value
}

func (opt *Options) Set(key string, data interface{}, expiration time.Duration) {
	opt.cache.Set(key, data, expiration)
}

func (opt *Options) Find(key string) bool {
	_, found := opt.cache.Get(key)
	return found
}
