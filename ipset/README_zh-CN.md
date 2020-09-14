# gonetx/ipset
<p align="center">
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

### 检查 ipset

### 创建新的 set

### 新增一条记录

### 删除一条记录

### 获取 set 信息
