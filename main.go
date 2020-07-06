package main

import (
	"errors"
	l "github.com/umirode/api-cache-proxy/logger"
	"github.com/umirode/api-cache-proxy/storage"
	"net/http"
	"os"
	"strconv"
)

func getProxyPort() (int, error) {
	envKey := "PROXY_PORT"

	proxyPortString := os.Getenv(envKey)
	if proxyPortString == "" {
		return 8080, nil
	}

	proxyPort, err := strconv.Atoi(proxyPortString)
	if err != nil {
		return 0, errors.New("invalid " + envKey + " value: " + proxyPortString)
	}

	return proxyPort, nil
}

func main() {
	logger := l.NewLogrusLogger()

	serializedResponseStorage, err := storage.NewGoCacheSerializedResponseStorage()
	if err != nil {
		logger.LogError("error on creating cache storage: ", err)
		return
	}

	proxyPort, err := getProxyPort()
	if err != nil {
		logger.LogError("invalid proxy port: ", err)
		return
	}

	err = http.ListenAndServe(":"+strconv.Itoa(proxyPort), getServerHandler(serializedResponseStorage, logger))
	if err != nil {
		logger.LogError("error proxy start: ", err)
	}
}
