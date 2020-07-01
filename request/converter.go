package request

import (
	"io"
	"net/http"
	"strings"
)

func cleanTrashFromUrl(url string) string {
	if string(url[0]) == "/" {
		url = url[1:]
	}

	return url
}

func restoreHttpSlash(url string) string {
	if !strings.Contains(url, "http://") {
		url = strings.ReplaceAll(url, "http:/", "http://")
	}

	return url
}

func restoreHttpsSlash(url string) string {
	if !strings.Contains(url, "https://") {
		url = strings.ReplaceAll(url, "https:/", "https://")
	}

	return url
}

func createUrlStringFromRequest(r *http.Request) string {
	url := r.RequestURI

	url = cleanTrashFromUrl(url)
	url = restoreHttpSlash(url)
	url = restoreHttpsSlash(url)

	return url
}

func addHeadersToRequest(headers http.Header, r *http.Request) *http.Request {
	r.Header = make(http.Header)
	for name, value := range headers {
		r.Header[name] = value
	}

	return r
}

func createRequest(method string, url string, body io.Reader, header http.Header) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	request = addHeadersToRequest(header, request)

	return request, nil
}

func ConvertRequestForExternalService(r *http.Request) (*http.Request, error) {
	bodyReader, err := createReaderForRequestBody(r)
	if err != nil {
		return nil, err
	}

	url := createUrlStringFromRequest(r)

	return createRequest(r.Method, url, bodyReader, r.Header)
}
