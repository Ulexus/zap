# :zap: zap [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

Blazing fast, structured, leveled logging in Go.

## Installation

`go get -u go.uber.org/zap`

## Quick Start

In contexts where performance is nice, but not critical, use the
`SugaredLogger`. It's 4-10x faster than than other structured logging libraries
and includes both structured and `printf`-style APIs.

```go
logger, _ := NewProduction()
sugar := logger.Sugar()
sugar.Infow("Failed to fetch URL.",
  // Structured context as loosely-typed key-value pairs.
  "url", url,
  "attempt", retryNum,
  "backoff", time.Second,
)
sugar.Infof("Failed to fetch URL: %s", url)
```

When performance and type safety are critical, use the `Logger`. It's even faster than
the `SugaredLogger` and allocates far less, but it only supports structured logging.

```go
logger, _ := NewProduction()
logger.Info("Failed to fetch URL.",
  // Structured context as strongly-typed Field values.
  zap.String("url", url),
  zap.Int("attempt", tryNum),
  zap.Duration("backoff", time.Second),
)
```

## Performance

For applications that log in the hot path, reflection-based serialization and
string formatting are prohibitively expensive &mdash; they're CPU-intensive and
make many small allocations. Put differently, using `encoding/json` and
`fmt.Fprintf` to log tons of `interface{}`s makes your application slow.

Zap takes a different approach. It includes a reflection-free, zero-allocation
JSON encoder, and the base `Logger` strives to avoid serialization overhead and
allocations wherever possible. By building the high-level `SugaredLogger` on
that foundation, zap lets users *choose* when they need to count every
allocation and when they'd prefer a more familiar, loosely-typed API.

As measured by its own [benchmarking suite][], not only is zap more performant
than comparable structured logging libraries &mdash; it's also faster than the
standard library. Like all benchmarks, take these with a grain of salt.<sup
id="anchor-versions">[1](#footnote-versions)</sup>

Log a message and 10 fields:

| Library | Time | Bytes Allocated | Objects Allocated |
| :--- | :---: | :---: | :---: |
| :zap: zap | 1175 ns/op | 706 B/op | 2 allocs/op |
| :zap: zap (sugared) | 1801 ns/op | 1614 B/op | 20 allocs/op |
| logrus | 7219 ns/op | 6103 B/op | 78 allocs/op |
| go-kit | 4999 ns/op | 2898 B/op | 66 allocs/op |
| log15 | 19603 ns/op | 5634 B/op | 93 allocs/op |
| apex/log | 18545 ns/op | 3836 B/op | 65 allocs/op |
| lion | 6923 ns/op | 5812 B/op | 63 allocs/op |

Log a message with a logger that already has 10 fields of context:

| Library | Time | Bytes Allocated | Objects Allocated |
| :--- | :---: | :---: | :---: |
| :zap: zap | 324 ns/op | 0 B/op | 0 allocs/op |
| :zap: zap (sugared) | 470 ns/op | 80 B/op | 2 allocs/op |
| logrus | 6512 ns/op | 4570 B/op | 63 allocs/op |
| go-kit | 4990 ns/op | 3049 B/op | 52 allocs/op |
| log15 | 15936 ns/op | 2643 B/op | 44 allocs/op |
| apex/log | 16299 ns/op | 2899 B/op | 51 allocs/op |
| lion | 4189 ns/op | 4077 B/op | 38 allocs/op |

Log a static string, without any context or `printf`-style templating:

| Library | Time | Bytes Allocated | Objects Allocated |
| :--- | :---: | :---: | :---: |
| :zap: zap | 346 ns/op | 0 B/op | 0 allocs/op |
| :zap: zap (sugared) | 446 ns/op | 80 B/op | 2 allocs/op |
| standard library | 638 ns/op | 80 B/op | 2 allocs/op |
| logrus | 1860 ns/op | 1507 B/op | 27 allocs/op |
| go-kit | 749 ns/op | 656 B/op | 13 allocs/op |
| log15 | 6783 ns/op | 1592 B/op | 26 allocs/op |
| apex/log | 3529 ns/op | 584 B/op | 11 allocs/op |
| lion | 1376 ns/op | 1225 B/op | 10 allocs/op |

## Development Status: Release Candidate 2
The current release is `v1.0.0-rc.2`. No further breaking changes are *planned*
unless wider use reveals critical bugs or usability issues, but users who need
absolute stability should wait for the 1.0.0 release.

<hr>
Released under the [MIT License](LICENSE.txt).

<sup id="footnote-versions">1</sup> In particular, keep in mind that we may be
benchmarking against slightly older versions of other libraries. Versions are
pinned in zap's [glide.lock][] file. [↩](#anchor-versions)

[doc-img]: https://godoc.org/go.uber.org/zap?status.svg
[doc]: https://godoc.org/go.uber.org/zap
[ci-img]: https://travis-ci.org/uber-go/zap.svg?branch=master
[ci]: https://travis-ci.org/uber-go/zap
[cov-img]: https://coveralls.io/repos/github/uber-go/zap/badge.svg?branch=master
[cov]: https://coveralls.io/github/uber-go/zap?branch=master
[benchmarking suite]: https://github.com/uber-go/zap/tree/master/benchmarks
[glide.lock]: https://github.com/uber-go/zap/blob/master/glide.lock
