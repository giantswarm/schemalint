package utils

import (
	"github.com/giantswarm/schemalint/pkg/schemautils"
)

func RecurseAll(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	if schema.IsSelfReference() {
		return
	}

	callback(schema)

	if schema.Ref != nil {
		refSchema := schema.GetRefSchema()
		RecurseAll(refSchema, callback)
	}

	for _, property := range schema.GetProperties() {
		RecurseAll(property, callback)
	}

	if schema.Items2020 != nil {
		RecurseAll(schema.GetItems2020(), callback)
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

func RecursePropertiesWithDescription(schema *schemautils.ExtendedSchema, callback func(schema *schemautils.ExtendedSchema)) {
	callbackIfPropertyWithDescription := func(schema *schemautils.ExtendedSchema) {
		if schema.IsProperty() && schema.Description != "" {
			callback(schema)
		}
	}

	RecurseAll(schema, callbackIfPropertyWithDescription)
}
