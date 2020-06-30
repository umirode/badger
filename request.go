package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func PrepareRequestToExternalService(r *http.Request) (*http.Request, error) {
	proxyRequestUrl := r.RequestURI

	if string(proxyRequestUrl[0]) == "/" {
		proxyRequestUrl = proxyRequestUrl[1:]
	}

	if !strings.Contains(proxyRequestUrl, "http://") {
		proxyRequestUrl = strings.ReplaceAll(proxyRequestUrl, "http:/", "http://")
	}
	if !strings.Contains(proxyRequestUrl, "https://") {
		proxyRequestUrl = strings.ReplaceAll(proxyRequestUrl, "https:/", "https://")
	}

	proxyRequestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(proxyRequestBody))

	proxyRequest, err := http.NewRequest(r.Method, proxyRequestUrl, bytes.NewReader(proxyRequestBody))
	if err != nil {
		return nil, err
	}

	logrus.Info(r.Header)

	proxyRequest.Header = make(http.Header)
	for name, value := range r.Header {
		proxyRequest.Header[name] = value
	}

	return proxyRequest, nil
}

func GenerateHashFromRequest(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	key := r.RequestURI + r.Method + r.Host + string(body) + r.URL.String()

	hash := md5.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])

	return hashString, nil
}
