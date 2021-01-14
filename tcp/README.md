This project is archived, please move to [gonetx/tcp](https://github.com/gonetx/tcp).

# Detect
Detect checks if a tcp address is connectable with specific timeout.

```go
ok, err := tcp.Detect("www.google.com:443", time.Second)
fmt.Println(ok, err)
```
