# Nano ID [![GoDoc](https://godoc.org/github.com/aidarkhanov/nanoid?status.svg)](https://pkg.go.dev/github.com/aidarkhanov/nanoid/v2)

### A tiny and fast _Go_ unique string generator

* __Safe__. It uses cryptographically strong random APIs and tests distribution of symbols.
* __Compact__. It uses a larger alphabet than UUID `(A-Za-z0-9_-)`. So ID size was reduced from 36 to 21 symbols.

```go
id, err := nanoid.New() //> "i25_rX9zwDdDn7Sg-ZoaH"
if err != nil {
    log.Fatalln(err)
}
```

### Installation

Once Go is installed, run the following command to get Nano ID.

```sh
go get github.com/aidarkhanov/nanoid
```

### Documentation

The package reference is located at [pkg.go.dev/github.com/aidarkhanov/nanoid/v2](https://pkg.go.dev/github.com/aidarkhanov/nanoid/v2).

### Roadmap

* The API of this package is frozen.
* Release patches if necessary.

### License

This package is provided under MIT/Expat license. See [LICENSE.md](https://raw.githubusercontent.com/aidarkhanov/nanoid/master/LICENSE) file for details.

### Thanks to

* [Andrey Sitnik](https://github.com/ai) for [the original JavaScript algorithm](https://github.com/ai/nanoid) implementation.
* [Paul Yuan](https://github.com/puyuan) for letting me improve [the Python version](https://github.com/puyuan/py-nanoid) of the library. Without the previous work, I would not be able to fully understand the algorithm.
