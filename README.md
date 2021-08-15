# base64

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]

Faster base64 encoding as specified by RFC 4648 for Go.

## Rationale

TODO

## Features

* Drop-in replacement of `encoding/base64`.
* ~3 times faster than `encoding/base64`.
* Dependency-free.

## Install

Go version 1.16+

```
go get github.com/cristalhq/base64
```

## How to use

Replace import statement from `encoding/base64` to `github.com/cristalhq/base64`

```diff
-import "encoding/base64"
+import "github.com/cristalhq/base64"
```

# Benchmarks

```
std/Encode           674.7 ns/op      0 B/op   0 allocs/op
std/EncodeToString   943.3 ns/op   2014 B/op   2 allocs/op
std/Decode           747.8 ns/op      0 B/op   0 allocs/op
std/DecodeString      1012 ns/op   1792 B/op   2 allocs/op

own/Encode           215.3 ns/op      0 B/op   0 allocs/op
own/EncodeToString   354.6 ns/op   1024 B/op   1 allocs/op
own/Decode           306.6 ns/op      0 B/op   0 allocs/op
own/DecodeString     440.0 ns/op    768 B/op   1 allocs/op
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/base64/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/base64/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/base64
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/base64
[reportcard-img]: https://goreportcard.com/badge/cristalhq/base64
[reportcard-url]: https://goreportcard.com/report/cristalhq/base64
[coverage-img]: https://codecov.io/gh/cristalhq/base64/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/base64
