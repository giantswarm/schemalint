package lint

import (
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

type Severity int

const (
	SeverityError Severity = iota
	SeverityRecommendation
)

type RuleViolation struct {
	message  string
	location string
}

type RuleResults struct {
	Violations []RuleViolation
}

func (r *RuleResults) Add(message string, location string) {
	r.Violations = append(r.Violations, RuleViolation{message: message, location: location})
}

// Filter out all violations whose location is in the excluded list or in any
// child locations.
func (r *RuleResults) Filter(excludedLocations []string) *RuleResults {
	filteredViolations := make([]RuleViolation, 0, len(r.Violations))
	for _, violation := range r.Violations {
		keep := true
		for _, excludedLocation := range excludedLocations {
			if violation.location == excludedLocation || schemautils.IsChildLocation(excludedLocation, violation.location) {
				keep = false
				break
			}
		}
		if keep {
			filteredViolations = append(filteredViolations, violation)
		}
	}
	return &RuleResults{Violations: filteredViolations}
}

func (r *RuleResults) GetMessages() []string {
	messages := make([]string, 0, len(r.Violations))
	for _, violation := range r.Violations {
		messages = append(messages, violation.message)
	}
	return messages
}

type Rule interface {
	Verify(*schemautils.ExtendedSchema) RuleResults
	GetSeverity() Severity
}
