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
