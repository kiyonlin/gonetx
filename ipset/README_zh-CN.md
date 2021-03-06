# gonetx/ipset
<p align="center">
  <br>
    <a href="https://pkg.go.dev/github.com/kiyonlin/gonetx/ipset?tab=doc">
      <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7&style=flat">
    </a>
    <a href="https://goreportcard.com/report/github.com/kiyonlin/gonetx">
      <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B">
    </a>
    <a href="https://gocover.io/github.com/kiyonlin/gonetx/ipset">
      <img src="https://img.shields.io/badge/%F0%9F%94%8E%20gocover-97.8%25-75C46B.svg?style=flat">
    </a>
    <a href="https://github.com/kiyonlin/gonetx/actions?query=workflow%3ASecurity">
      <img src="https://img.shields.io/github/workflow/status/gofiber/fiber/Security?label=%F0%9F%94%91%20gosec&style=flat&color=75C46B">
    </a>
    <a href="https://github.com/kiyonlin/gonetx/actions?query=workflow%3ATest">
      <img src="https://img.shields.io/github/workflow/status/gofiber/fiber/Test?label=%F0%9F%A7%AA%20tests&style=flat&color=75C46B">
    </a>
</p>

该包是`ipset`工具的`Golang`封装。它使得`Golang`程序更简单地操作`ipset`。

访问[http://ipset.netfilter.org/ipset.man.html](http://ipset.netfilter.org/ipset.man.html)了解更多`ipset`命令文档。`ipset`需要`v6.0+`版本。

## 安装
使用`go get`安装`ipset`:
```bash
go get -u github.com/kiyonlin/gonetx/ipset
```

## 快速使用

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
	set, _ := ipset.New("test", ipset.HashIp, ipset.Exist(true), ipset.Timeout(time.Hour))
	// output: test
	log.Println(set.Name())

	_ = set.Flush()

	_ = set.Add("1.1.1.1", ipset.Timeout(time.Hour))

	ok, _ := set.Test("1.1.1.1")
	// output: true
	log.Println(ok)

	ok, _ = set.Test("1.1.1.2")
	// output: false
	log.Println(ok)

	info, _ := set.List()
	// output: &{test hash:ip 4 family inet hashsize 1024 maxelem 65536 timeout 3600 216 0 [1.1.1.1 timeout 3599]}
	log.Println(info)

	_ = set.Del("1.1.1.1")

	_ = set.Destroy()
}
```

## Check
在使用该库之前，您应该记住始终先调用`ipset.Check`并处理错误。此方法将检查`ipset`是否在`OS PATH`中存在以及其版本是否有效。

```go
func init() {
	// err will be ipset.ErrNotFound
	// or ipset.ErrVersionNotSupported
	// if check failed.
	if err := ipset.Check(); err != nil {
		panic(err)
	}
}
```

## New
使用`ipset.New`创建一个用`setname`和指定的`set`类型标识的`set`。如果指定了`ipset.Exist`选项，则当已经存在相同的`set`（`set`名称和创建参数相同）时，`ipset`将忽略该错误。

```go
set, _ := ipset.New("test", ipset.HashIp, ipset.Exist(true), ipset.Netmask(24))
```

每个`set`类型可能具有不同的创建选项，请访问[SetType](https://pkg.go.dev/github.com/kiyonlin/gonetx/ipset?tab=doc#SetType)和[Option](https://pkg.go.dev/github.com/kiyonlin/gonetx/ipset?tab=doc#Option)以获取更多详细信息。

创建`set`后，可以使用以下方法：
```go
// IPSet is abstract of ipset
type IPSet interface {
	// List dumps header data and the entries for the set to an
	// *Info instance. The Resolve option can be used to force
	// action lookups(which may be slow).
	List(options ...Option) (*Info, error)

	// List dumps header data and the entries for the set to the
	// specific file. The Resolve option can be used to force
	// action lookups(which may be slow).
	ListToFile(filename string, options ...Option) error

	// Name returns the set's name
	Name() string

	// Rename the set's action and the new action must not exist.
	Rename(newName string) error

	// Add adds a given entry to the set. If the Exist option is
	// specified, ipset ignores the error if the entry already
	// added to the set.
	Add(entry string, options ...Option) error

	// Del deletes an entry from a set. If the Exist option is
	// specified and the entry is not in the set (maybe already
	// expired), then the command ignores the error.
	Del(entry string, options ...Option) error

	// Test tests whether an entry is in a set or not.
	Test(entry string) (bool, error)

	// Flush flushed all entries from the the set.
	Flush() error

	// Destroy removes the set from kernel.
	Destroy() error

	// Save dumps the set data to a io.Reader in a format that restore
	// can read.
	Save(options ...Option) (io.Reader, error)

	// SaveToFile dumps the set data to s specific file in a format
	// that restore can read.
	SaveToFile(filename string, options ...Option) error

	// Restore restores a saved session from io.Reader generated by
	// save. Set exist to true to ignore exist error.
	Restore(r io.Reader, exist ...bool) error

	// RestoreFromFile restores a saved session from a specific file
	// generated by save. Set exist to true to ignore exist error.
	RestoreFromFile(filename string, exist ...bool) error
}
```

## Swap
使用`ipset.Swap`交换两个`set`的内容，换句话说，交换两个`set`的动作。引用的`set`必须存在，并且兼容类型的`set`才能互换。

```go
ipset.Swap("foo", "bar")
```
## Flush
使用`ipset.Flush`刷新指定`set`中的所有条目或者所有`set`。

```go
// Flush foo and bar set
ipset.Flush("foo", "bar")

// Flush all
ipset.Flush()
```

## Destroy
使用`ipset.Destroy`删除指定的`set`或所有`set`。如果`set`有被引用时，则什么也不做，也不会破坏`set`。

```go
// Destroy foo and bar set
ipset.Destroy("foo", "bar")

// Destroy all
ipset.Destroy()
```
