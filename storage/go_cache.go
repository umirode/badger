package storage

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"github.com/umirode/api-cache-proxy/response"
	"time"
)

type GoCacheSerializedResponseStorage struct {
	c *cache.Cache
}

func NewGoCacheSerializedResponseStorage() *GoCacheSerializedResponseStorage {
	return &GoCacheSerializedResponseStorage{
		c: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (g GoCacheSerializedResponseStorage) Get(hash string) (*response.SerializedResponse, error) {
	serializedResponse, exists := g.c.Get(hash)
	if !exists {
		return nil, errors.New("not exists")
	}

	return serializedResponse.(*response.SerializedResponse), nil
}

func (g GoCacheSerializedResponseStorage) Set(hash string, response *response.SerializedResponse) error {
	g.c.SetDefault(hash, response)

	return nil
}
