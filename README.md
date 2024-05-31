# schemalint – the Giant Swarm JSON Schema linter

schemalint helps you write great JSON schema for [app](https://docs.giantswarm.io/platform-overview/app-platform/) configuration.

## Features

- Validate whether your schema is valid JSON schema
- Normalize JSON schema (indentation, white space, sorting) and check normalization
- Experimental: deep validation for [cluster app schema requirements](https://github.com/giantswarm/rfc/pull/55)

Validates whether an input is valid JSON schema. Also helps normalizing, and checks for normalization.

## Installation

```nohighlight
go install github.com/giantswarm/schemalint/v2@latest
```

## Usage

### Validation

Executing `schemalint verify` without any options will check whether a file is valid JSON Schema and whether it is normalized.

```nohighlight
$ schemalint verify myschema.json

Errors (1)

- schema is not normalized

Verification result

- [SUCCESS] Input is valid JSON Schema.
- [ERROR] Input is not normalized.
```

To validate a **cluster app** schema, apply the `--rule-set cluster-app` option.

```nohighlight
schemalint verify myschema.json --rule-set cluster-app
```

**Note:** Cluster app schema validation is experimental and in development. The requirements are in discussion in [this RFC draft](https://github.com/giantswarm/rfc/pull/55).

Use `--help` to learn about more options.

### Normalization

Create a normalized (white space, sorting) representation of a JSON Schema file. This helps to avoid purely cosmetical changes to a schema.

```nohighlight
schemalint normalize myschema.json > normalized.json
```

Use `--help` to learn about more options.

## GitHub Action

An action to run `schemalint verify` on the `values.schema.json` in app repositories in provided in `actions/verify-helm-schema`.

**Example workflow**:

```yaml
name: JSON schema validation
on:
  push: {}

jobs:
  validate:
    name: Verify values.schema.json with schemalint
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3
      - name: Run schemalint
        id: run-schemalint
        uses: giantswarm/schemalint/actions/verify-helm-schema@v1
        with:
          rule-set: 'cluster-app'
```

Note that it is possible to define the rule set to be used for the `verify` command with the `with` keyword.

```yaml
with:
  rule-set: 'RULE_SET'
```

If the rule set is not specified, no rule set will be used.

## Pre-commit Hook

This repository provides a [pre-commit hook](https://pre-commit.com/#new-hooks) that can be used to normalize JSON schema files before committing them. See `.pre-commit-hooks.yaml` for more information.

## Major Releases

This repository uses [floating tags](https://github.com/giantswarm/floating-tags-action).
Other repositories that use schemalint point to major floating tag versions,
like `v1`. That means that all minor and patch releases will be automatically
rolled out to these repositories.

When doing a major release the references to schemalint have to be manually
updated.

- devctl:
  - `pkg/gen/input/workflows/internal/file/cluster_app_schema_validation.yaml.template`
  - `pkg/gen/input/makefile/internal/file/Makefile.gen.cluster_app.mk.template`
- schema:
  - `.github/workflows/lint.yaml

## "Overriding" Properties and Understanding `PropertyAnnotationsMap`

In JSON schema it is possible to have multiple definitions for a properties
with the same location. This is possible through the use of
[refs](https://json-schema.org/understanding-json-schema/structuring.html).
Consider the following example.

```json
{
    "$defs": {
        "foo": {
            "properties": {
                "childProp": {
                    "type": "string",
                    "minLength": 2,
                    "title": "This title will be overridden"
                }
            },
            "type": "object"
        }
    },
    "properties": {
        "rootProp": {
            "$ref": "#/$defs/foo",
            "properties": {
                "childProp": {
                    "type": "string",
                    "maxLength": 4,
                    "title": "This title will be used"
                }
            },
            "type": "object"
        }
    },
    "type": "object"
}
```

Here, the property at the location `.rootProp.childProp` has two different
definitions.

One is in the original schema:

```json
{
  "type": "string",
  "maxLength": 4,
  "title": "This title will be used"
}
```

And the other one is in the referenced schema:

```json
{
  "type": "string",
  "minLength": 2,
  "title": "This title will be overridden"
}
```

In JSON schema specification is no such thing as overriding or merging. The
keywords that have actual meaning during validation will be applied
sequentially (e.g. `minLength`, `maxLength` and `type`).
In our example a payload that conforms to the given schema would need to have
a string of length 2,3 or 4 at the location `.rootProp.childProp`.

The JSON schema specification does not specify how to handle multiple
definitions for annotations like `title`
([ref](https://json-schema.org/draft/2020-12/json-schema-core.html#name-distinguishing-among-multip)).
As we use annotations for our UI, we need a clear convention when handling
multiple annotations.

To specify which annotation to use, we use _reference levels_.
A reference level describes how often a `$ref` keyword was resolved to get to
the current schema.
The root schema always has reference level 0. The resolved schema of a `$ref`
keyword increments the reference level.
The annotation, which belongs to the schema definition with the lowest
reference level is used for our UI.

In the above example what we called "original schema" has reference level 0 and
what we called "referenced schema" has reference level 1.
Therefore, the displayed title of the `childProp` would be `This title will be used`.

This logic does not only apply to our UI but also when validating annotations.
Therefore, an implementation of what annotations to use can be found in
`propertyannotations.go@BuildPropertyAnnotationsMap`.
