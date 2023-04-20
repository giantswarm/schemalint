package recurse

import "github.com/giantswarm/schemalint/v2/pkg/schema"

func RecurseAll(
	s *schema.ExtendedSchema,
	callback func(s *schema.ExtendedSchema),
) {
	if s.IsSelfReference() {
		return
	}

	callback(s)

	callChildren(s, func(s *schema.ExtendedSchema) {
		RecurseAll(s, callback)
	})
}

func RecurseAllPre(
	s *schema.ExtendedSchema,
	callback func(s *schema.ExtendedSchema),
) {
	callback(s)

	if s.IsSelfReference() {
		return
	}

	callChildren(s, func(s *schema.ExtendedSchema) {
		RecurseAllPre(s, callback)
	})
}

func callChildren(
	s *schema.ExtendedSchema,
	callback func(s *schema.ExtendedSchema),
) {
	if s.Ref != nil {
		refSchema := s.GetRefSchema()
		callback(refSchema)
	}

	for _, property := range s.GetProperties() {
		callback(property)
	}

	if s.Items2020 != nil {
		callback(s.GetItems2020())
	}

	if s.Items != nil {
		schemas := s.GetItems()
		for _, itemSchema := range schemas {
			callback(itemSchema)
		}
	}

}

func RecurseProperties(
	s *schema.ExtendedSchema,
	callback func(s *schema.ExtendedSchema),
) {
	callbackIfProperty := func(s *schema.ExtendedSchema) {
		if s.IsProperty() {
			callback(s)
		}
	}

	RecurseAll(s, callbackIfProperty)
}

func RecurseObjects(
	s *schema.ExtendedSchema,
	callback func(s *schema.ExtendedSchema),
) {
	callbackIfObject := func(s *schema.ExtendedSchema) {
		if s.IsObject() {
			callback(s)
		}
	}

	RecurseAll(s, callbackIfObject)
}

func RecurseArrays(
	s *schema.ExtendedSchema,
	callback func(s *schema.ExtendedSchema),
) {
	callbackIfArray := func(s *schema.ExtendedSchema) {
		if s.IsArray() {
			callback(s)
		}
	}

	RecurseAll(s, callbackIfArray)
}

func RecurseStrings(
	s *schema.ExtendedSchema,
	callback func(s *schema.ExtendedSchema),
) {
	callbackIfString := func(s *schema.ExtendedSchema) {
		if s.IsString() {
			callback(s)
		}
	}

	RecurseAll(s, callbackIfString)
}

func RecurseNumerics(
	s *schema.ExtendedSchema,
	callback func(s *schema.ExtendedSchema),
) {
	callbackIfString := func(s *schema.ExtendedSchema) {
		if s.IsNumeric() {
			callback(s)
		}
	}

	RecurseAll(s, callbackIfString)
}
