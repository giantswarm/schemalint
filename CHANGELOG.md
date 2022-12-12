# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/schemalint/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/giantswarm/schemalint/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/schemalint/compare/v0.0.2...v0.1.0
[0.0.2]: https://github.com/giantswarm/schemalint/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/schemalint/releases/tag/v0.0.1
