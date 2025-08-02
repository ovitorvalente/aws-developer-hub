package repository

import (
	"sync"
	"time"
)

type CacheRepository struct {
	data      string
	timestamp time.Time
	mutex     sync.RWMutex
	ttl       time.Duration
}

func NewCacheRepository() *CacheRepository {
	return &CacheRepository{
		ttl: 5 * time.Minute, // Cache por 5 minutos
	}
}

func (c *CacheRepository) Get() (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	
	if time.Since(c.timestamp) > c.ttl {
		return "", false
	}
	
	return c.data, c.data != ""
}

func (c *CacheRepository) Set(data string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	
	c.data = data
	c.timestamp = time.Now()
}