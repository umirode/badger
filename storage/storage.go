package storage

import (
	"github.com/umirode/api-cache-proxy/response"
)

type ISerializedResponseStorage interface {
	Get(hash string) (*response.SerializedResponse, error)
	Set(hash string, response *response.SerializedResponse) error
}
