package schema

import (
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

	s, err := compiler.Compile(url)
	if err != nil {
		return nil, err
	}
	extendedSchema := NewExtendedSchema(s)
	extendedSchema.RootFilePath = path
	return extendedSchema, nil
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
