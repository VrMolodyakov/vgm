package cache

import (
	"time"

	ttlcache "github.com/VrMolodyakov/ttl-cache"
)

func New[K comparable, V any](cleanInterval int) *ttlcache.Cache[K, V] {
	return ttlcache.New[K, V](time.Duration(cleanInterval) * time.Minute)
}
