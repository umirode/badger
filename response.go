package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type SerializedResponse struct {
	Header     map[string]string
	Body       []byte
	StatusCode int
}

func SerializeResponse(r *http.Response) (*SerializedResponse, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	headers := map[string]string{}
	for headerKey, headerValue := range r.Header {
		headers[headerKey] = headerValue[0]
	}

	sr := &SerializedResponse{
		Header:     headers,
		Body:       bodyBytes,
		StatusCode: r.StatusCode,
	}

	return sr, nil
}

func WriteSerializeResponseToResponseWriter(r *SerializedResponse, w http.ResponseWriter) error {
	for headerKey, headerValue := range r.Header {
		w.Header().Set(headerKey, headerValue)
	}

	_, err := io.Copy(w, bytes.NewReader(r.Body))
	if err != nil {
		return err
	}

	return nil
}
