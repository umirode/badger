package response

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

func createHeaderMapFromResponse(r *http.Response) map[string]string {
	headers := map[string]string{}
	for headerKey, headerValue := range r.Header {
		headers[headerKey] = headerValue[0]
	}

	return headers
}

func getBodyFromResponse(r *http.Response) ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, err
}

func NewSerializeResponseFromResponse(r *http.Response) (*SerializedResponse, error) {
	bodyBytes, err := getBodyFromResponse(r)
	if err != nil {
		return nil, err
	}

	headers := createHeaderMapFromResponse(r)

	return &SerializedResponse{
		Header:     headers,
		Body:       bodyBytes,
		StatusCode: r.StatusCode,
	}, nil
}

func (r *SerializedResponse) writeHeadersToResponseWriter(w http.ResponseWriter) {
	for headerKey, headerValue := range r.Header {
		w.Header().Set(headerKey, headerValue)
	}
}

func (r *SerializedResponse) writeBodyToResponseWriter(w http.ResponseWriter) error {
	_, err := io.Copy(w, bytes.NewReader(r.Body))
	return err
}

func (r *SerializedResponse) WriteToResponseWriter(w http.ResponseWriter) error {
	r.writeHeadersToResponseWriter(w)

	return r.writeBodyToResponseWriter(w)
}
