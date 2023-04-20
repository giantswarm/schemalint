package rules

import (
	"github.com/giantswarm/schemalint/v2/pkg/schema"
)

type Severity int

const (
	SeverityError Severity = iota
	SeverityRecommendation
)

type Violation struct {
	Message  string
	Location string
}

type RuleResults struct {
	Violations []Violation
}

func (r *RuleResults) Add(message string, location string) {
	r.Violations = append(r.Violations, Violation{Message: message, Location: location})
}

func (r *RuleResults) Concat(other RuleResults) {
	r.Violations = append(r.Violations, other.Violations...)
}

func (r *RuleResults) IsEmpty() bool {
	return len(r.Violations) == 0
}

// Filter out all violations whose location is in the excluded list or in any
// child locations.
func (r RuleResults) Filter(excludedLocations []string) RuleResults {
	filteredViolations := make([]Violation, 0, len(r.Violations))
	for _, violation := range r.Violations {
		keep := true
		for _, excludedLocation := range excludedLocations {
			if violation.Location == excludedLocation ||
				schema.IsChildLocation(excludedLocation, violation.Location) {
				keep = false
				break
			}
		}
		if keep {
			filteredViolations = append(filteredViolations, violation)
		}
	}
	return RuleResults{Violations: filteredViolations}
}

func (r *RuleResults) GetMessages() []string {
	messages := make([]string, 0, len(r.Violations))
	for _, violation := range r.Violations {
		messages = append(messages, violation.Message)
	}
	return messages
}

type Rule interface {
	Verify(*schema.ExtendedSchema) RuleResults
	GetSeverity() Severity
}
