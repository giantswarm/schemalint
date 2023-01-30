package lint

import (
	"net/url"
	"path/filepath"
	"runtime"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"

	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type draft struct {
	jsonschemaDraft *jsonschema.Draft
	name            string
}

type Linter struct {
	compilers map[string]*jsonschema.Compiler
}

var (
	// All drafts in their chonological order.
	drafts = []draft{
		{
			jsonschemaDraft: jsonschema.Draft4,
			name:            "draft4",
		},
		{
			jsonschemaDraft: jsonschema.Draft6,
			name:            "draft6",
		},
		{
			jsonschemaDraft: jsonschema.Draft7,
			name:            "draft7",
		},
		{
			jsonschemaDraft: jsonschema.Draft2019,
			name:            "draft2019",
		},
		{
			jsonschemaDraft: jsonschema.Draft2020,
			name:            "draft2020",
		},
	}
)

func New() *Linter {
	l := &Linter{
		compilers: make(map[string]*jsonschema.Compiler, 5),
	}

	// One compiler for each JSONSchema draft supported.
	for _, draft := range drafts {
		l.compilers[draft.name] = jsonschema.NewCompiler()
		l.compilers[draft.name].Draft = draft.jsonschemaDraft
		l.compilers[draft.name].ExtractAnnotations = true
	}

	return l
}

// Returns the number of supported drafts.
func (l *Linter) NumDrafts() int {
	return len(l.compilers)
}

// Validate a schema, given by URL, against all JSON Schema drafts.
// This can be used to detect general problems in the schema.
//
// Returns a list of successful drafts and a map of failed drafts with error details.
func (l *Linter) CompileAllDrafts(url string) (success []string, errors map[string]error) {
	errors = make(map[string]error)

	for _, draft := range drafts {
		_, err := l.compilers[draft.name].Compile(url)
		if err != nil {
			errors[draft.name] = err
		} else {
			success = append(success, draft.name)
		}
	}

	return success, errors
}

// This processes the given input without specifying the draft to use.
// If the schema provides a valid `$schema` property, the one given will
// be used. If not, the latest draft will be used.
// In case of success, a string will be returned, otherwise an error.
func Compile(path string) (*schemautils.ExtendedSchema, error) {
	url, err := ToFileURL(path)

	if err != nil {
		return nil, err
	}

	compiler := jsonschema.NewCompiler()
	compiler.ExtractAnnotations = true

	schema, err := compiler.Compile(url)
	if err != nil {
		return nil, err
	}
	extendedSchema := schemautils.NewExtendedSchema(schema)
	extendedSchema.RootFilePath = path
	return extendedSchema, nil
}

func ToFileURL(path string) (string, error) {
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
