package utils

import (
	"github.com/giantswarm/schemalint/v2/pkg/schemautils"
)

func RecurseAll(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	if schema.IsSelfReference() {
		return
	}

	callback(schema)

	callChildren(schema, func(schema *schemautils.ExtendedSchema) {
		RecurseAll(schema, callback)
	})
}

func RecurseAllPre(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	callback(schema)

	if schema.IsSelfReference() {
		return
	}

	callChildren(schema, func(schema *schemautils.ExtendedSchema) {
		RecurseAllPre(schema, callback)
	})
}

func callChildren(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	if schema.Ref != nil {
		refSchema := schema.GetRefSchema()
		callback(refSchema)
	}

	for _, property := range schema.GetProperties() {
		callback(property)
	}

	if schema.Items2020 != nil {
		callback(schema.GetItems2020())
	}

	if schema.Items != nil {
		schemas := schema.GetItems()
		for _, itemSchema := range schemas {
			callback(itemSchema)
		}
	}

}

func RecurseProperties(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	callbackIfProperty := func(schema *schemautils.ExtendedSchema) {
		if schema.IsProperty() {
			callback(schema)
		}
	}

	RecurseAll(schema, callbackIfProperty)
}

func RecurseObjects(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	callbackIfObject := func(schema *schemautils.ExtendedSchema) {
		if schema.IsObject() {
			callback(schema)
		}
	}

	RecurseAll(schema, callbackIfObject)
}

func RecurseArrays(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	callbackIfArray := func(schema *schemautils.ExtendedSchema) {
		if schema.IsArray() {
			callback(schema)
		}
	}

	RecurseAll(schema, callbackIfArray)
}

func RecurseStrings(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	callbackIfString := func(schema *schemautils.ExtendedSchema) {
		if schema.IsString() {
			callback(schema)
		}
	}

	RecurseAll(schema, callbackIfString)
}

func RecurseNumerics(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	callbackIfString := func(schema *schemautils.ExtendedSchema) {
		if schema.IsNumeric() {
			callback(schema)
		}
	}

	RecurseAll(schema, callbackIfString)
}
