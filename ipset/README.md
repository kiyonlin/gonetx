# gonetx/ipset
<p align="center">
<a href="https://github.com/kiyonlin/gonetx/tree/master/ipset/README_zh-CN.md">
    <img height="20px" src="https://img.shields.io/badge/CN-flag.svg?color=555555&style=flat&logo=data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxMjAwIDgwMCIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiPg0KPHBhdGggZmlsbD0iI2RlMjkxMCIgZD0ibTAsMGgxMjAwdjgwMGgtMTIwMHoiLz4NCjxwYXRoIGZpbGw9IiNmZmRlMDAiIGQ9Im0tMTYuNTc5Niw5OS42MDA3bDIuMzY4Ni04LjEwMzItNi45NTMtNC43ODgzIDguNDM4Ni0uMjUxNCAyLjQwNTMtOC4wOTI0IDIuODQ2Nyw3Ljk0NzkgOC40Mzk2LS4yMTMxLTYuNjc5Miw1LjE2MzQgMi44MTA2LDcuOTYwNy02Ljk3NDctNC43NTY3LTYuNzAyNSw1LjEzMzF6IiB0cmFuc2Zvcm09Im1hdHJpeCg5LjkzMzUyIC4yNzc0NyAtLjI3NzQ3IDkuOTMzNTIgMzI0LjI5MjUgLTY5NS4yNDE1KSIvPg0KPHBhdGggZmlsbD0iI2ZmZGUwMCIgaWQ9InN0YXIiIGQ9Im0zNjUuODU1MiwzMzIuNjg5NWwyOC4zMDY4LDExLjM3NTcgMTkuNjcyMi0yMy4zMTcxLTIuMDcxNiwzMC40MzY3IDI4LjI1NDksMTEuNTA0LTI5LjU4NzIsNy40MzUyLTIuMjA5NywzMC40MjY5LTE2LjIxNDItMjUuODQxNS0yOS42MjA2LDcuMzAwOSAxOS41NjYyLTIzLjQwNjEtMTYuMDk2OC0yNS45MTQ4eiIvPg0KPGcgZmlsbD0iI2ZmZGUwMCI+DQo8cGF0aCBkPSJtNTE5LjA3NzksMTc5LjMxMjlsLTMwLjA1MzQtNS4yNDE4LTE0LjM5NDUsMjYuODk3Ni00LjMwMTctMzAuMjAyMy0zMC4wMjkzLTUuMzc4MSAyNy4zOTQ4LTEzLjQyNDItNC4xNjQ3LTMwLjIyMTUgMjEuMjMyNiwyMS45MDU3IDI3LjQ1NTQtMTMuMjk5OC0xNC4yNzIzLDI2Ljk2MjcgMjEuMTMzMSwyMi4wMDE3eiIvPg0KPHBhdGggZD0ibTQ1NS4yNTkyLDMxNS45Nzk1bDkuMzczNC0yOS4wMzE0LTI0LjYzMjUtMTcuOTk3OCAzMC41MDctLjA1NjYgOS41MDUtMjguOTg4NiA5LjQ4MSwyOC45OTY0IDMwLjUwNywuMDgxOC0yNC42NDc0LDE3Ljk3NzQgOS4zNDkzLDI5LjAzOTItMjQuNzE0LTE3Ljg4NTgtMjQuNzI4OCwxNy44NjUzeiIvPg0KPC9nPg0KPHVzZSB4bGluazpocmVmPSIjc3RhciIgdHJhbnNmb3JtPSJtYXRyaXgoLjk5ODYzIC4wNTIzNCAtLjA1MjM0IC45OTg2MyAxOS40MDAwNSAtMzAwLjUzNjgxKSIvPg0KPC9zdmc+DQo=">
  </a>
  <br>
    <a href="https://pkg.go.dev/github.com/kiyonlin/gonetx/ipset?tab=doc">
      <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7&style=flat">
    </a>
    <a href="https://goreportcard.com/report/github.com/kiyonlin/gonetx">
      <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B">
    </a>
    <a href="https://gocover.io/github.com/kiyonlin/gonetx">
      <img src="https://img.shields.io/badge/%F0%9F%94%8E%20gocover-97.8%25-75C46B.svg?style=flat">
    </a>
    <a href="https://github.com/kiyonlin/gonetx/actions?query=workflow%3AGosec">
      <img src="https://img.shields.io/github/workflow/status/gofiber/fiber/Security?label=%F0%9F%94%91%20gosec&style=flat&color=75C46B">
    </a>
    <a href="https://github.com/kiyonlin/gonetx/actions?query=workflow%3ATest">
      <img src="https://img.shields.io/github/workflow/status/gofiber/fiber/Test?label=%F0%9F%A7%AA%20tests&style=flat&color=75C46B">
    </a>
</p>

This package is a almost whole `Golang` wrapper to the `ipset` user space utility. It allows `Golang` programs to easily manipulate `ipset`.

Visit [http://ipset.netfilter.org/ipset.man.html](http://ipset.netfilter.org/ipset.man.html) for more `ipset` command details. And `ipset` requires version `v6.0+`.

## Installation
Install `ipset` using the `go get` command:
```bash
go get -u github.com/kiyonlin/gonetx/ipset
```

## Quickstart

```go
package main

import (
	"log"
	"time"

	"github.com/kiyonlin/gonetx/ipset"
)

func init() {
	if err := ipset.Check(); err != nil {
		panic(err)
	}
}

func main() {
	// create test set even it's exist
	set, _ := ipset.New("test", ipset.HashIp, ipset.Exist(true))
	// output: test
	log.Println(set.Name())

	set.Flush()

	set.Add("1.1.1.1", ipset.Timeout(time.Hour))

	ok, _ := set.Test("1.1.1.1")
	// output: true
	log.Println(ok)

	ok, _ = set.Test("1.1.1.2")
	// output: false
	log.Println(ok)

	info, _ := set.List()
	// output:
	log.Println(info)

	set.Del("1.1.1.1")

	set.Destroy()
}
```

### Check ipset

### Create a new set

### Add a single entry to the set

### Delete a single entry from the set

### List entries of a set
