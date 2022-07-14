# Second

[![Tests](https://github.com/iMuFeng/second/actions/workflows/test.yml/badge.svg)](https://github.com/iMuFeng/second/actions)
[![Go Version](https://img.shields.io/github/go-mod/go-version/iMuFeng/second?label=Go)](https://github.com/iMuFeng/second/blob/master/go.mod)
[![Go Ref](https://pkg.go.dev/badge/github.com/iMuFeng/second)](https://pkg.go.dev/github.com/iMuFeng/second)

[This module](https://pkg.go.dev/github.com/iMuFeng/second) is a Go implementation of [ms](https://github.com/vercel/ms).

# Install

```
go get github.com/iMuFeng/second
```

## Examples

```go
second.Parse('2 days')  // 172800
second.Parse('1d')      // 86400
second.Parse('10h')     // 36000
second.Parse('2.5 hrs') // 9000
second.Parse('2h')      // 7200
second.Parse('1m')      // 60
second.Parse('5s')      // 5
second.Parse('1y')      // 31557600
second.Parse('-3 days') // -259200
second.Parse('-1h')     // -3600
```

## License

[MIT License](./LICENSE)
