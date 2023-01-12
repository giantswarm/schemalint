package findings

type Severity int

const (
	SeverityError Severity = iota
	SeverityWarning
	SeverityInfo
)

type LintFindings struct {
	Message string
	Severity
}

func (f LintFindings) String() string {
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
