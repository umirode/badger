package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

var cacheStorage = NewInMemoryCacheStorage()

var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	preparedRequest, err := PrepareRequestToExternalService(request)
	if err != nil {
		logrus.Error("error on preparing request to external service: ", err.Error())

		return
	}
	defer preparedRequest.Body.Close()

	preparedRequestHash, err := GenerateHashFromRequest(preparedRequest)
	if err != nil {
		logrus.Error("error generating hash from request: ", err.Error())

		return
	}

	var serializedResponse *SerializedResponse

	if cacheStorage.Exists(preparedRequestHash) {
		logrus.Info("use request from cache")

		serializedResponse, err = cacheStorage.Get(preparedRequestHash)
		if err != nil {
			logrus.Error("error on loading cached response: ", err.Error())
			return
		}
	} else {
		logrus.Info("sending new request")

		externalServiceResponse, err := http.DefaultClient.Do(preparedRequest)
		if err != nil {
			logrus.Error("error on sending request to external service: ", err.Error())
			return
		}

		if externalServiceResponse.StatusCode != http.StatusOK {
			logrus.Error("external service return not 200 status response")
			return
		}

		serializedResponse, err = SerializeResponse(externalServiceResponse)
		if err != nil {
			logrus.Error("error on serializing response: ", err.Error())
			return
		}

		err = cacheStorage.Set(preparedRequestHash, serializedResponse)
		if err != nil {
			logrus.Error("error on adding response to cache: ", err.Error())
			return
		}
	}

	err = WriteSerializeResponseToResponseWriter(serializedResponse, writer)
	if err != nil {
		logrus.Error("error on return response: ", err.Error())
		return
	}
}

func main() {
	logrus.Info("running proxy")
	err := http.ListenAndServe(":8082", handler)
	if err != nil {
		logrus.Error("error proxy start: ", err.Error())
	}
}
