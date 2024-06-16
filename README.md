# go-ethaddr

Package **ethaddr** provides tools for working with **eth-addresses**, for the Go programming language.

The **eth-address** (also known as an **ethereum-address** or **evm-address**) is a 20-byte address — and is commonly used by EVM based networks.

`ethaddr.Address` is meant to be a replacement for `go-ethereum/common.Address` from the official Ethereum golang package.

## Examples

Here is an example of loading an` ethaddr.Address` from a hexadecimal-literal stored in a Go `string`:

```golang
address, err := ethaddr.Parse("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
```

Here is an example of loading an` ethaddr.Address` from a Go `[20]byte`:

```golang
address := ethaddr.Something( [20]{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed} )
```

Here is an example of loading an` ethaddr.Address` from a Go `*big.Int`:

```golang
var bigint *big.Int

address := ethaddr.BigInt(bigint)
```

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-ethaddr

[![GoDoc](https://godoc.org/github.com/reiver/go-ethaddr?status.svg)](https://godoc.org/github.com/reiver/go-ethaddr)

## Import

To import package **ethaddr** use `import` code like the follownig:
```
import "github.com/reiver/go-ethaddr"
```

## Installation

To install package **ethaddr** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-ethaddr
```

## Author

Package **ethaddr** was written by [Charles Iliya Krempeaux](http://reiver.link)
