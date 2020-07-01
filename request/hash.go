package request

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
)

func createKeyForRequest(r *http.Request) (string, error) {
	body, err := getBodyFromRequest(r)
	if err != nil {
		return "", err
	}

	return r.RequestURI + r.Method + r.Host + string(body) + r.URL.String(), nil
}

func CreateHashForRequest(r *http.Request) (string, error) {
	key, err := createKeyForRequest(r)
	if err != nil {
		return "", err
	}

	hash := md5.Sum([]byte(key))

	return hex.EncodeToString(hash[:]), nil
}
