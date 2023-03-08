#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 RULE_SET"
    exit 1
fi

RULE_SET=$1

VALUES_SCHEMA=$(find ./helm -maxdepth 2 -name values.schema.json)

schemalint verify ${VALUES_SCHEMA} --rule-set ${RULE_SET}
