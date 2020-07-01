package storage

import (
	"errors"
	"github.com/umirode/api-cache-proxy/response"
)

type InMemorySerializedResponseStorage struct {
	data map[string]*response.SerializedResponse
}

func NewInMemorySerializedResponseStorage() *InMemorySerializedResponseStorage {
	return &InMemorySerializedResponseStorage{
		data: map[string]*response.SerializedResponse{},
	}
}

func (i InMemorySerializedResponseStorage) Get(hash string) (*response.SerializedResponse, error) {
	r, ok := i.data[hash]
	if !ok {
		return nil, errors.New("not exists")
	}

	return r, nil
}

func (i InMemorySerializedResponseStorage) Set(hash string, response *response.SerializedResponse) error {
	i.data[hash] = response

	return nil
}
