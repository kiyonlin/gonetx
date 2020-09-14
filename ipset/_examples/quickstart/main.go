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
