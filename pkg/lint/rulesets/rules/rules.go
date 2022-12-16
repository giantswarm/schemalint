package rules

import (
	"github.com/giantswarm/schemalint/pkg/lint/findings"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type RuleViolation struct {
	Reason string
}

type Rule interface {
	Verify(*jsonschema.Schema) []RuleViolation
	GetSeverity() findings.Severity
}
