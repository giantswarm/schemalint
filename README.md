# JSON Schema Linter

Validates whether an input is valid JSON schema.

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

### Normalization

Create a normalized (white space, sorting) representation of a JSON Schema file.

```nohighlight
$ schemalint normalize myschema.json > normalized.json
```
