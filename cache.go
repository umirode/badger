package main

import "github.com/pkg/errors"

type IResponseCacheStorage interface {
	Exists(hash string) bool
	Get(hash string) (*SerializedResponse, error)
	Set(hash string, response *SerializedResponse) error
}

type InMemoryCacheStorage struct {
	data map[string]*SerializedResponse
}

func NewInMemoryCacheStorage() *InMemoryCacheStorage {
	return &InMemoryCacheStorage{
		data: map[string]*SerializedResponse{},
	}
}

func (i InMemoryCacheStorage) Exists(hash string) bool {
	_, ok := i.data[hash]

	return ok
}

func (i InMemoryCacheStorage) Get(hash string) (*SerializedResponse, error) {
	response, ok := i.data[hash]
	if !ok {
		return nil, errors.New("Not exists")
	}

	return response, nil
}

func (i InMemoryCacheStorage) Set(hash string, response *SerializedResponse) error {
	i.data[hash] = response

	return nil
}
