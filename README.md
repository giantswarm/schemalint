# schemalint – the Giant Swarm JSON Schema linter

schemalint helps you write great JSON schema for [app](https://docs.giantswarm.io/platform-overview/app-platform/) configuration.

## Features

- Validate whether your schema is valid JSON schema
- Normalize JSON schema (indentation, white space, sorting) and check normalization
- Experimental: deep validation for [cluster app schema requirements](https://github.com/giantswarm/rfc/pull/55)

Validates whether an input is valid JSON schema. Also helps normalizing, and checks for normalization.

## Installation

```nohighlight
go install github.com/giantswarm/schemalint@latest
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
$ schemalint verify myschema.json --rule-set cluster-app
```

**Note:** Cluster app schema validation is experimental and in development. The requirements are in discussion in [this RFC draft](https://github.com/giantswarm/rfc/pull/55).

Use `--help` to learn about more options.

### Normalization

Create a normalized (white space, sorting) representation of a JSON Schema file. This helps to avoid purely cosmetical changes to a schema.

```nohighlight
$ schemalint normalize myschema.json > normalized.json
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

## Major Releases

This repository uses [floating tags](https://github.com/giantswarm/floating-tags-action).
Other repositories that use schemalint point to major floating tag versions,
like `v1`. That means that all minor and patch releases will be automatically
rolled out to these repositories.
When doing a major release the following steps have to be completed:
1. Create a new major floating tag under "Actions -> Ensure major version tags -> Run Workflow"
2. Update all references to schemalint.
    1. devctl: `pkg/gen/input/workflows/internal/file/cluster_app_schema_validation.yaml.template`
