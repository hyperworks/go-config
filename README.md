# GO-CONFIG

Simple configuration reader that utilizes struct tags and reads into a pre-defined
strongly-typed struct.

```go
import "github.com/hyperworks/go-config"
```

See the file `example/main.go` for an example.

# EXAMPLE

```sh
$ export CFG_ENV=test
$ export CFG_DRIVER=mysql
$ export CFG_DB=/tmp/mysql.sock
$
$ go run $GOPATH/src/github.com/hyperworks/go-config/example/main.go
{
	"Env": "test",
	"DBDriver": "mysql",
	"DBURL": "/tmp/mysql.sock"
}
```
