package normalize

const (
	// Four spaces Indentation.
	indentation          = "    "
	prefix               = ""
	defaultKeyImportance = -1
)

type KeyImportanceOverride struct {
	KeyImportance map[string]int
	When          func(key string, value interface{}) bool
}

// getKeyImportanceMap returns a map where the key is a possible key in a JSON
// schema file and the value is the importance of the key. The higher the
// importance, the higher/earlier the key will be in the normalized output.
func getKeyImportanceMap() map[string]int {
	return map[string]int{
		"$id":                  11,
		"$schema":              10,
		"$ref":                 9,
		"$defs":                9,
		"type":                 8,
		"title":                7,
		"description":          6,
		"deprecated":           5,
		"$comment":             4,
		"enum":                 3,
		"examples":             3,
		"default":              2,
		"required":             2,
		"additionalProperties": 1,
		"properties":           1,
		"patternProperties":    1,
		"items":                1,
		"if":                   -2,
		"then":                 -3,
		"else":                 -4,
	}
}

// getKeyImportanceMapOverrides returns a list of KeyImportanceOverrides. Each
// override contains a function that determines if the override should be used
// and a map of key importance values that should be merged with the default
// key importance map.
func getKeyImportanceMapOverrides() []KeyImportanceOverride {
	keyImportanceWhenValueNonPrimitive := map[string]int{
		"additionalProperties": 0,
		"default":              0,
	}

	whenNonPrimitive := func(key string, value interface{}) bool {
		switch value.(type) {
		case string:
			return false
		case int:
			return false
		case float64:
			return false
		case bool:
			return false
		default:
			return true
		}
	}

	return []KeyImportanceOverride{
		{
			When:          whenNonPrimitive,
			KeyImportance: keyImportanceWhenValueNonPrimitive,
		},
	}
}
