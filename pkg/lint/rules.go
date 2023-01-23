package lint

import (
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type Severity int

const (
	SeverityError Severity = iota
	SeverityRecommendation
)

type Rule interface {
	Verify(*schemautils.ExtendedSchema) []string
	GetSeverity() Severity
}
