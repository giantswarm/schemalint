# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.6.1] - 2025-01-22

- Dependency updates

## [2.6.0] - 2024-06-03

### Added

- Add [pre-commit](https://pre-commit.com/) hooks

## [2.5.1] - 2023-12-19

- Update schemalint version in the GitHub action.

## [2.5.0] - 2023-12-15

- Don't check `cluster` in cluster-app ruleset, because that property should never be modified, and it is used for configuring cluster chart.
- Don't check `providerIntegration` in cluster-app ruleset, because here we have static properties that are set only in the provider-specific charts.

## [2.4.0] - 2023-11-24

- Don't check `global.apps` in cluster-app ruleset, so we can have empty objects without defined properties (until we auto-generate full schemas from actual app schemas).

## [2.3.1] - 2023-11-23

- Allow to have the same field as a both global and root-level property.

## [2.3.0] - 2023-11-16

- Support for refactored schema where root-level properties are moved to be under global property.

## [2.2.1] - 2023-11-09

## [2.2.0] - 2023-11-09

- Allow root-level object `global` for sharing values with subcharts

## [2.1.1] - 2023-04-25

## [2.1.0] - 2023-04-25

- Update major release instructions in README.
- Improve clarity of verify command output.

## [2.0.0] - 2023-04-18

- Add reference to `schema` in the major release instructions.
- Introduce a meaningful order of keywords in JSON output of normalize command.

## [1.2.0] - 2023-04-18

- Adapt common schema structure to cluster-aws.
- Only recommend examples on restricted strings.
- Add the possibility to include a URL for further reference to rule sets, which is displayed in the output.
- Hint to developer how to normalize the JSON schema file.

## [1.1.0] - 2023-04-05

- Add rule to check whether root-level specifies properties of the common schema structure.
- Add rule to check that objects are non-empty.

## [1.0.1] - 2023-03-22

- Fix module path.

## [1.0.0] - 2023-03-22

- Fix update action workflow trigger.
- Add a reusable composite GitHub action that calls `schemalint verify`.

## [0.10.0] - 2023-03-07

- Add possibility to exclude locations from rule set validation.

## [0.9.0] - 2023-03-02

- Check if normalization is already applied when calling normalize with output argument.
- Add rules to forbid infinite recursion and recursion-related keywords.
- Add possibility to output version with `schemalint -v` or `schemalint --version`.
- Add rule to check whether logical constructs (if, then & else) are not used.
- Add check that schemas only use `anyOf` and `oneOf` for specific purposes.
- Add rule to check that arrays only specify one type for their items.

## [0.8.0] - 2023-02-24

- Avoid `unevaluatedItems` and `unevaluatedProperties`.
- Prevent normalize from escaping <,> and &.
- Add rule to check whether all properties have exactly one type.
- Don't require examples for boolean properties.
- Change required draft back to 2020-12.
- Add check that every numeric property is constrained through 'minimum', 'maximum', 'exclusiveMinimum' or 'exclusiveMaximum'.
- Add check that every string property is constrained through 'pattern', 'minLength', 'maxLength', 'enum', 'constant' or 'format'.
- Add check that every property provides one or more examples with the `examples` keyword.
- Add check that every deprecated property has a comment.
- Add check that properties don't provide more than 5 examples.

## [0.7.0] - 2023-02-06

- Add check that every array specifies the schema of its items with the `items` keyword.
- Simplify output of `schemalint verify` to improve appearance in GitHub actions log.
- Add check that `additionalProperties` is disabled on all objects.

## [0.6.0] - 2023-02-02

- Add check for draft version.
- Add check that an existing title follows requirements as defined in [this RFC](https://github.com/giantswarm/rfc/pull/55).

## [0.5.0] - 2023-01-30

- Resolve schemas referenced through the `$ref` keyword while checking rules.
- Wrap all table-based tests in `t.Run`.
- Add check that an description exists on every property to `cluster-app` rule set.
- Add check that an existing description follows requirements as defined in [this RFC](https://github.com/giantswarm/rfc/pull/55).

## [0.4.0] - 2023-01-16

- Add possible values of `rule-set` to `--help` message.

## [0.3.0] - 2023-01-12

- `normalize`: add flags to write to output file path.
- Add `rule-set` flag to `verify` command, that enables additional optional rules.

## [0.2.0] - 2022-12-12

- Extend `verify` command to also check normalization.
- Fix `normalize`: avoid extra line break at end of output (breaking).

## [0.1.0] - 2022-12-09

- Move linting function from root command into `verify` command (breaking).
- Add `normalize` command.
- Avoid double error output.
- Quit with error if users gives multiple positional arguments to the `verify` and `normalize` command.

## [0.0.2] - 2022-12-05

- Return non-zero exit code in case of error.

## [0.0.1] - 2022-12-05

- Added first basic linting.

[Unreleased]: https://github.com/giantswarm/schemalint/compare/v2.6.1...HEAD
[2.6.1]: https://github.com/giantswarm/schemalint/compare/v2.6.0...v2.6.1
[2.6.0]: https://github.com/giantswarm/schemalint/compare/v2.5.1...v2.6.0
[2.5.1]: https://github.com/giantswarm/schemalint/compare/v2.5.0...v2.5.1
[2.5.0]: https://github.com/giantswarm/schemalint/compare/v2.4.0...v2.5.0
[2.4.0]: https://github.com/giantswarm/schemalint/compare/v2.3.1...v2.4.0
[2.3.1]: https://github.com/giantswarm/schemalint/compare/v2.3.0...v2.3.1
[2.3.0]: https://github.com/giantswarm/schemalint/compare/v2.2.1...v2.3.0
[2.2.1]: https://github.com/giantswarm/schemalint/compare/v2.2.0...v2.2.1
[2.2.0]: https://github.com/giantswarm/schemalint/compare/v2.1.1...v2.2.0
[2.1.1]: https://github.com/giantswarm/schemalint/compare/v2.1.0...v2.1.1
[2.1.0]: https://github.com/giantswarm/schemalint/compare/v2.0.0...v2.1.0
[2.0.0]: https://github.com/giantswarm/schemalint/compare/v1.2.0...v2.0.0
[1.2.0]: https://github.com/giantswarm/schemalint/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/giantswarm/schemalint/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/giantswarm/schemalint/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/schemalint/compare/v0.10.0...v1.0.0
[0.10.0]: https://github.com/giantswarm/schemalint/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/giantswarm/schemalint/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/giantswarm/schemalint/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/giantswarm/schemalint/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/giantswarm/schemalint/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/schemalint/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/giantswarm/schemalint/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/schemalint/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/schemalint/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/schemalint/compare/v0.0.2...v0.1.0
[0.0.2]: https://github.com/giantswarm/schemalint/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/schemalint/releases/tag/v0.0.1
