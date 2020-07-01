package main

import (
	"errors"
	r "github.com/umirode/api-cache-proxy/request"
	"github.com/umirode/api-cache-proxy/response"
	"github.com/umirode/api-cache-proxy/storage"
	"net/http"
)

func sendRequestToExternalService(request *http.Request) (*response.SerializedResponse, error) {
	externalServiceResponse, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.New("error on sending request to external service: " + err.Error())
	}

	if externalServiceResponse.StatusCode != http.StatusOK {
		return nil, errors.New("external service return not 200 response status")
	}

	serializedResponse, err := response.NewSerializeResponseFromResponse(externalServiceResponse)
	if err != nil {
		return nil, errors.New("error on serializing response: " + err.Error())
	}

	return serializedResponse, nil
}

func getSerializedResponseByRequest(serializedResponseStorage storage.ISerializedResponseStorage, request *http.Request) (*response.SerializedResponse, error) {
	requestHash, err := r.CreateHashForRequest(request)
	if err != nil {
		return nil, errors.New("error generating hash from request: " + err.Error())
	}

	serializedResponse, err := serializedResponseStorage.Get(requestHash)
	if err == nil {
		return serializedResponse, nil
	}

	serializedResponse, err = sendRequestToExternalService(request)
	if err != nil {
		return nil, errors.New("error on sending request to external service: " + err.Error())
	}

	err = serializedResponseStorage.Set(requestHash, serializedResponse)
	if err != nil {
		return nil, errors.New("error on adding response to cache: " + err.Error())
	}

	return serializedResponse, nil
}

func getServerHandler(serializedResponseStorage storage.ISerializedResponseStorage, logger ILogger) http.HandlerFunc {
	logger.LogInfo("proxy started")

	return func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		convertedRequest, err := r.ConvertRequestForExternalService(request)
		if err != nil {
			logger.LogError("error on preparing request to external service: ", err)
			return
		}
		defer convertedRequest.Body.Close()

		serializedResponse, err := getSerializedResponseByRequest(serializedResponseStorage, convertedRequest)
		if err != nil {
			logger.LogError("error on getting serialized response: ", err)
			return
		}

		err = serializedResponse.WriteToResponseWriter(writer)
		if err != nil {
			logger.LogError("error on return response: ", err)
			return
		}
	}
}
