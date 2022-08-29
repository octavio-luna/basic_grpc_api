package procedures

import (
	"log"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	cache *cache.Cache
}

//returns a new cache instance with no expiration
func NewCache() *Cache {
	log.Println("Creating new cache")
	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	return &Cache{
		cache: c,
	}
}
