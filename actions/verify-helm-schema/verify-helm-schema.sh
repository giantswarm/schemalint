#!/bin/bash

RULE_SET=$1

VALUES_SCHEMA=$(find ./helm -maxdepth 2 -name values.schema.json)

echo "Running schemalint verify on ${VALUES_SCHEMA}"

SCHEMALINT_ARGS=${VALUES_SCHEMA}

if [ -z "$RULE_SET" ]; then
    echo "Using no rule set"
else
    echo "Using rule set ${RULE_SET}"
    SCHEMALINT_ARGS="${SCHEMALINT_ARGS} --rule-set ${RULE_SET}"
fi

schemalint verify ${SCHEMALINT_ARGS}
