# rand

> A random library for Go.

[![Build][build-status-image]][build-status-url] [![Codecov][codecov-image]][codecov-url] [![ReportCard][reportcard-image]][reportcard-url] [![GoDoc][godoc-image]][godoc-url] [![License][license-image]][license-url]

## Usage

```bash
$ go get -u -v github.com/LyricTian/rand
```

### Import package

```go
import "github.com/LyricTian/rand"
```

### Generate a random string

```go
rand.MustRandom(6, rand.Ldigit)
// 939430
```

#### Random flag:

- Ldigit
- LlowerCase
- LupperCase
- LlowerAndUpperCase
- LdigitAndLowerCase
- LdigitAndUpperCase
- LdigitAndLetter

### Generate a random UUID

```go
rand.MustUUID()
// a306e54f-672f-4011-889d-d09e98cbea89
```

## MIT License

    Copyright (c) 2018 Lyric

[build-status-url]: https://travis-ci.org/LyricTian/rand
[build-status-image]: https://travis-ci.org/LyricTian/rand.svg?branch=master
[codecov-url]: https://codecov.io/gh/LyricTian/rand
[codecov-image]: https://codecov.io/gh/LyricTian/rand/branch/master/graph/badge.svg
[reportcard-url]: https://goreportcard.com/report/github.com/LyricTian/rand
[reportcard-image]: https://goreportcard.com/badge/github.com/LyricTian/rand
[godoc-url]: https://godoc.org/github.com/LyricTian/rand
[godoc-image]: https://godoc.org/github.com/LyricTian/rand?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg
