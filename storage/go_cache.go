package storage

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"github.com/umirode/api-cache-proxy/response"
	"os"
	"strconv"
	"time"
)

type GoCacheSerializedResponseStorage struct {
	c *cache.Cache
}

func getCacheExpirationTime() (time.Duration, error) {
	envKey := "CACHE_EXPIRATION_TIME"

	cacheExpirationTimeString := os.Getenv(envKey)
	if cacheExpirationTimeString == "" {
		return cache.NoExpiration, nil
	}

	cacheExpirationTime, err := strconv.Atoi(cacheExpirationTimeString)
	if err != nil {
		return 0, errors.New("invalid " + envKey + " value: " + cacheExpirationTimeString)
	}

	return time.Duration(cacheExpirationTime) * time.Second, nil
}

func NewGoCacheSerializedResponseStorage() (*GoCacheSerializedResponseStorage, error) {
	expTime, err := getCacheExpirationTime()
	if err != nil {
		return nil, err
	}

	return &GoCacheSerializedResponseStorage{
		c: cache.New(expTime, cache.NoExpiration),
	}, nil
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
