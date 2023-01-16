package lint

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Severity int

const (
	SeverityError Severity = iota
	SeverityRecomendation
)

type Rule interface {
	Verify(*jsonschema.Schema) []string
	GetSeverity() Severity
}
