# Releasing

`main` is the only long-lived branch. Changes reach it through pull requests, merge commits only, and
every pull request and push to `main` must pass the `build` check (`.github/workflows/build.yml`:
`go mod verify`, `go mod tidy -diff`, `go vet`, `go build` and `go test -race` on the two latest Go
releases). The tests that call the API run against the staging API (`https://stag-v2.convertapi.com`)
with the `CONVERTAPI_STAGING_TOKEN` secret; without the secret (fork pull requests) they skip and the
offline tests still run. Merging publishes nothing.

A Go module version is just a tag: there is no package to build, no registry upload and no key. The
module path is **`github.com/ConvertAPI/convertapi-go`** (the repository's old name; GitHub redirects it
to this repository, and every existing user imports it), so the version `v1.0.1` is the tag `v1.0.1` in
this repository. `v*` tags are created only by publishing a **GitHub Release** (repository admins only, by
ruleset), which `tomasr78` does:

1. In a pull request, set `Version` in `pkg/config/version.go` to the new version (it goes into the
   `User-Agent` of every request) and add the version to `CHANGELOG.md`. Merge it.
2. Wait for the `build` check of that merge on `main` to pass, then publish the release:

```bash
gh release create v1.0.1 --repo ConvertAPI/convertapi-library-go --target main --title v1.0.1 --generate-notes
```

That runs `.github/workflows/release.yml` in the `go-proxy` environment, which:

1. refuses a Release published or re-run by anyone else;
2. checks the tag is SemVer (`v1.0.1`, `v1.1.0-beta.1`), that its major version fits the module path, that
   the GitHub pre-release tick matches the version, that the tagged commit is on `main`, and that
   `config.Version` at the tag is the version;
3. runs `go vet`, `go build` and every test, the API ones included, on the latest Go against the staging API;
4. runs `go mod download github.com/ConvertAPI/convertapi-go@v1.0.1` through `proxy.golang.org`, which
   makes the proxy fetch the tag and checks it against `sum.golang.org`, then checks that the proxy serves
   the version from the tagged commit;
5. asks pkg.go.dev to fetch the version and waits up to 15 minutes for
   <https://pkg.go.dev/github.com/ConvertAPI/convertapi-go@v1.0.1> (a warning only).

The tag is public the moment it exists, so these checks cannot stop a release; a red run means the version
that is already public needs a fix. **Pre-releases** (`--prerelease`, e.g. `v1.1.0-beta.1`) work the same
way and are also tagged on `main`; `go get` installs them only when asked for by version
(`go get github.com/ConvertAPI/convertapi-go@v1.1.0-beta.1`), never as `@latest`.

Go module versions are permanent: once the proxy has fetched a tag it serves that content forever, and the
checksum database pins it, so a moved or re-created tag breaks every build that uses it. Never delete, move
or re-create a release tag. Fix a bad version by releasing the next one and say so in both Releases' notes.
To keep a broken version from being chosen, add a `retract` directive to `go.mod` in the next release.

## Major versions

Go's import compatibility rule puts v2 and later under a different module path: a `v2.0.0` release needs
`module github.com/ConvertAPI/convertapi-go/v2` in `go.mod`, and every import in the code, the tests and
the examples changes to `.../convertapi-go/v2/pkg...`. Users then upgrade by changing their imports. Both
`build.yml` and `release.yml` refuse a v2+ version on a module path without the suffix. Breaking API
changes, such as the unreleased `v3` branch (the `api.convertapi.io/v3` API with renamed types and
functions), go out as the next major version this way.

## Module path

The module path stays `github.com/ConvertAPI/convertapi-go` because changing it would break every user's
imports. It works only while GitHub redirects the old repository name here: never create a repository named
`ConvertAPI/convertapi-go`. `go get github.com/ConvertAPI/convertapi-library-go` does not work and fails
with "module declares its path as github.com/ConvertAPI/convertapi-go".
