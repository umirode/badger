package main

import (
	"github.com/sirupsen/logrus"
	"github.com/umirode/api-cache-proxy/storage"
	"net/http"
)

func main() {
	var serializedResponseStorage = storage.NewGoCacheSerializedResponseStorage()
	var logger = NewLogrusLogger()

	err := http.ListenAndServe(":8082", getServerHandler(serializedResponseStorage, logger))
	if err != nil {
		logrus.Error("error proxy start: ", err.Error())
	}
}
