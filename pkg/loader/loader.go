// Package loader provides means to load schema via HTTP(s) or file system.
package loader

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

type HTTPURLLoader http.Client

func (l *HTTPURLLoader) Load(url string) (any, error) {
	client := (*http.Client)(l)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("%s returned status code %d", url, resp.StatusCode)
	}
	defer resp.Body.Close()

	return jsonschema.UnmarshalJSON(resp.Body)
}

func New(insecure bool) *jsonschema.SchemeURLLoader {
	return &jsonschema.SchemeURLLoader{
		"file":  jsonschema.FileLoader{},
		"http":  NewHTTPURLLoader(insecure),
		"https": NewHTTPURLLoader(insecure),
	}
}

func NewHTTPURLLoader(insecure bool) *HTTPURLLoader {
	httpLoader := HTTPURLLoader(http.Client{
		Timeout: 15 * time.Second,
	})
	if insecure {
		httpLoader.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // nolint:gosec
		}
	}
	return &httpLoader
}
