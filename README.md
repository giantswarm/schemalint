# JSON Schema linter

Validates whether an input is valid JSON schema. Also helps normalizing, and checks for normalization.

## Installation

```nohighlight
go install github.com/giantswarm/schemalint@latest
```

## Usage

### Validation

Verify whether a file is valid JSON Schema and is normalized.

```nohighlight
$ schemalint verify myschema.json
Is valid JSON schema: SUCCESS
Is normalized: SUCCESS
```

Use `--help` to learn about more options.

### Normalization

Create a normalized (white space, sorting) representation of a JSON Schema file.

```nohighlight
$ schemalint normalize myschema.json > normalized.json
```

Use `--help` to learn about more options.
