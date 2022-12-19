package rules

import (
	"github.com/santhosh-tekuri/jsonschema/v5"

	"github.com/giantswarm/schemalint/pkg/lint/findings"
)

type RuleViolation struct {
	Reason string
}

type Rule interface {
	Verify(*jsonschema.Schema) []RuleViolation
	GetSeverity() findings.Severity
}
