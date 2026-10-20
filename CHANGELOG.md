# Changelog

All notable changes to this project will be documented in this file.

The format is based on Keep a Changelog and this project adheres to Semantic Versioning. Before 1.0.0 the
module had no version tags, so `go get` resolved pseudo-versions such as
`v0.0.0-20250603083246-b586aa6ba8a2`; their history is the
[commit log](https://github.com/ConvertAPI/convertapi-library-go/commits/main).

## [1.0.0] — 2026-10-20

The first tagged release. The API is the one of the latest pseudo-version
(`v0.0.0-20250603083246-b586aa6ba8a2`), with the changes below.

### Changed
- `config.Version` is the string `"1.0.0"` (it was the number `4`), and the `User-Agent` is
  `ConvertAPI-Go/1.0.0 (<os>)`. Code that used `config.Version` as a number must change.
- The default branch is `main` (was `master`).
- Releases are now cut from a GitHub Release `v{version}` on `main`, checked by `release.yml`, which also
  makes the Go module proxy fetch the version. Every pull request is built and tested on the two latest
  Go releases.

### Fixed
- `config.New` with a nil transport no longer panics on the first request; it uses
  `http.DefaultTransport`. It also sets `Config.CaTransport`, which was nil, and no longer panics when
  `config.Default` is not set and the URL is nil.
- `param.NewFile` no longer copies a mutex (`go vet` copylocks).
- The README's full example imports the right packages.

[1.0.0]: https://github.com/ConvertAPI/convertapi-library-go/releases/tag/v1.0.0
