# Nano ID

[![GoDoc](https://godoc.org/github.com/aidarkhanov/nanoid?status.svg)](https://godoc.org/github.com/aidarkhanov/nanoid) [![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/aidarkhanov/nanoid/master/LICENSE) [![Coverage](http://gocover.io/_badge/github.com/aidarkhanov/nanoid)](http://gocover.io/github.com/aidarkhanov/nanoid)

A tiny, fast and convenient _Go_ unique string ID generator.

* __Safe__. It uses cryptographically strong random APIs and tests distribution of symbols.
* __Compact__. It uses a larger alphabet than UUID (A-Za-z0-9_-). So ID size was reduced from 36 to 21 symbols.

```go
id := nanoid.New() //=> i25_rX9zwDdDn7Sg-ZoaH
```

## Installation

```sh
go get github.com/aidarkhanov/nanoid
```

## Thanks to

* [Andrey Sitnik](https://github.com/ai) for [the original JavaScript algorithm](https://github.com/ai/nanoid) implementation.
* [Paul Yuan](https://github.com/puyuan) for allowing me to improve [the Python version](https://github.com/puyuan/py-nanoid) of the library. Without the previous work, I would not be able to fully understand the algorithm.
