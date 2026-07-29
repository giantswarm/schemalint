package schema

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"runtime"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v6"
)

// This processes the given input without specifying the draft to use.
// If the s provides a valid `$s` property, the one given will
// be used. If not, the latest draft will be used.
// In case of success, a string will be returned, otherwise an error.
func Compile(path string) (*ExtendedSchema, error) {
	url, err := toFileURL(path)

	if err != nil {
		return nil, err
	}

	compiler := jsonschema.NewCompiler()
	// Without this, 'Format' is left unset in draft 2020-12, where the format
	// vocabulary is annotation-only. Rules read 'format' to tell a constrained
	// string from an unconstrained one, so they need it either way.
	compiler.AssertFormat()
	compiler.UseLoader(jsonschema.SchemeURLLoader{
		"file":  jsonschema.FileLoader{},
		"http":  httpLoader{},
		"https": httpLoader{},
	})

	s, err := compiler.Compile(url)
	if err != nil {
		return nil, err
	}
	extendedSchema := NewExtendedSchema(s)
	extendedSchema.RootFilePath = path
	return extendedSchema, nil
}

// httpLoader loads a schema referenced over 'http' or 'https'. Only 'file' is
// loadable out of the box; the httploader package that used to cover the rest is
// gone as of jsonschema v6.
//
// ponytail: plain http.Get, no timeout or caching -- one fetch per remote '$ref'
// during a single lint run. Give it a client with a timeout if a hanging server
// ever becomes a problem.
type httpLoader struct{}

func (httpLoader) Load(url string) (any, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s returned status code %d", url, resp.StatusCode)
	}

	return jsonschema.UnmarshalJSON(resp.Body)
}

func toFileURL(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	path = filepath.ToSlash(path)
	if runtime.GOOS == "windows" {
		path = "/" + path
	}
	u, err := url.Parse("file://" + path)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
