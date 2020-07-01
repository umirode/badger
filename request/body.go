package request

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func getBodyFromRequest(r *http.Request) ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}

func createReaderForRequestBody(r *http.Request) (io.Reader, error) {
	bodyBytes, err := getBodyFromRequest(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bodyBytes), err
}
