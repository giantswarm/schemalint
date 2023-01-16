package findings

import "github.com/giantswarm/schemalint/pkg/cli"

type Severity int

const (
	SeverityError Severity = iota
	SeverityRecomendation
)

type LintFindings struct {
	Message string
	Severity
}

func (f LintFindings) String() string {
	switch f.Severity {
	case SeverityError:
		return cli.SprintErrorMessage(f.Message)
	case SeverityRecomendation:
		return f.Message
	}
	return f.Message
}

func GetCount(fs []LintFindings, severity Severity) int {
	count := 0
	for _, f := range fs {
		if f.Severity == severity {
			count++
		}
	}
	return count
}
