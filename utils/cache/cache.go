package cache

import (
	"github.com/sillyhatxu/gocache-client"
	"time"
)

var Client *client.CacheConfig

func Initial() {
	Client = client.NewCacheConfig(12*time.Hour, 24*time.Hour)
}
