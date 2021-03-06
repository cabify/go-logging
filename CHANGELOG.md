# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Nothing

### Changed
- Nothing

### Deprecated
- Nothing

### Removed
- Nothing

### Fixed
- Nothing

### Security
- Nothing

## [1.5.1] - 2019-07-30
### Changed
- Make `DefaultFactory` of type `LoggerFactory` so it accepts other implementations

## [1.5.0] - 2019-07-30
### Added
- `LoggerFactory` interface abstracting `Factory` struct
- `DefaultFactory` global var that allows to replace factory used to generate loggers

## [1.4.0] - 2019-05-30
### Changed
- Baggage key is now a public string `"logctx-data-map-string-interface"` that can be set and read by anyone from any package.
The type of that baggage will be always `map[string]interface{}`

### Added
- `BaggageContextKey` constant that defines that string
- `Baggage(ctx) map[string]interface{}` that allows reading the underlying baggage. 

## [1.3.0] - 2019-03-18
### Added
- `Level` now implements the `envconfig.Decoder` interface so it can be used in config types

## [1.2.0] - 2019-02-27
### Added
- `ConfigureDefaultLogger` boilerplate for logger configuration
- `measuredLoggingHandler` to wrap loggers with metrics
- `CHANGELOG.md`: this file
