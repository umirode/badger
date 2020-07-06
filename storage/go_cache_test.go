package storage

import (
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/umirode/api-cache-proxy/response"
	"os"
	"testing"
	"time"
)

func TestGoCacheSerializedResponseStorage_Get_not_exists(t *testing.T) {
	s, err := NewGoCacheSerializedResponseStorage()
	assert.Nil(t, err)

	r, err := s.Get("key")
	assert.Error(t, err)
	assert.Nil(t, r)
}

func TestGoCacheSerializedResponseStorage_Get(t *testing.T) {
	s, err := NewGoCacheSerializedResponseStorage()
	assert.Nil(t, err)

	err = s.Set("key", &response.SerializedResponse{
		StatusCode: 225,
	})
	assert.Nil(t, err)

	r, err := s.Get("key")
	assert.Nil(t, err)
	assert.Equal(t, 225, r.StatusCode)
}

func TestGoCacheSerializedResponseStorage_Set(t *testing.T) {
	s, err := NewGoCacheSerializedResponseStorage()
	assert.Nil(t, err)

	err = s.Set("key", &response.SerializedResponse{})
	assert.Nil(t, err)
}

func TestNewGoCacheSerializedResponseStorage(t *testing.T) {
	s, err := NewGoCacheSerializedResponseStorage()
	assert.Nil(t, err)

	assert.Implements(t, (*ISerializedResponseStorage)(nil), s)
	assert.IsType(t, (*GoCacheSerializedResponseStorage)(nil), s)

}

func TestNewGoCacheSerializedResponseStorage_error(t *testing.T) {
	err := os.Setenv("CACHE_EXPIRATION_TIME", "error")
	assert.Nil(t, err)
	defer os.Unsetenv("CACHE_EXPIRATION_TIME")

	s, err := NewGoCacheSerializedResponseStorage()
	assert.Error(t, err)
	assert.Nil(t, s)
}

func Test_getCacheExpirationTime(t *testing.T) {
	t.Parallel()

	err := os.Setenv("CACHE_EXPIRATION_TIME", "10")
	assert.Nil(t, err)
	defer os.Unsetenv("CACHE_EXPIRATION_TIME")

	expTime, err := getCacheExpirationTime()
	assert.Nil(t, err)
	assert.Equal(t, time.Second*10, expTime)
}

func Test_getCacheExpirationTime_default(t *testing.T) {
	t.Parallel()

	expTime, err := getCacheExpirationTime()
	assert.Nil(t, err)
	assert.Equal(t, cache.NoExpiration, expTime)
}

func Test_getCacheExpirationTime_invalid(t *testing.T) {
	t.Parallel()

	err := os.Setenv("CACHE_EXPIRATION_TIME", "test")
	assert.Nil(t, err)
	defer os.Unsetenv("CACHE_EXPIRATION_TIME")

	expTime, err := getCacheExpirationTime()
	assert.Error(t, err)
	assert.Equal(t, cache.NoExpiration, expTime)
}
