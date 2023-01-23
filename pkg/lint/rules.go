package lint

import (
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type Severity int

const (
	SeverityError Severity = iota
	SeverityRecommendation
)

type RuleResults struct {
	Violations []string
}

func (r *RuleResults) Add(violation string) {
	r.Violations = append(r.Violations, violation)
}

type Rule interface {
	Verify(*schemautils.ExtendedSchema) RuleResults
	GetSeverity() Severity
}
