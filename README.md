[![GoDoc](https://godoc.org/github.com/tsne/conf?status.png)](https://godoc.org/github.com/tsne/conf)

# conf

conf is a simple configuration library for Go. It provides the functionality to import configuration files of any format.

Use `go get` to install or update the package:
```
go get -u github.com/tsne/conf
```

## Examples
Loading a JSON configuration file:
```go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/tsne/conf"
)

func main() {
	configFile := flag.String("config", "myconf.json", "configuration file")
	flag.Parse()

	f, err := os.Open(*configFile)
	if err != nil {
		panic(err)
	}

	c := conf.MustLoad(f, json.Unmarshal)
	for k, v := range c {
		fmt.Printf("%v: %v\n", k, v)
	}
}
```

Decoding a struct:
```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/tsne/conf"
)

const myConfStr = `{
	"keepalive": true,
	"server": {
		"address": "192.168.1.7",
		"port": 8080
	},
	"timeouts": {
		"requests": "5s"
	}
}`

type serverConf struct {
	Address string `config:",required"`
	Port    int    `config:",required"`
}

func main() {
	c := conf.MustLoad(bytes.NewBufferString(myConfStr), json.Unmarshal)

  var requestTimeout time.Duration
	if err := c.Decode("timeouts.requests", &requestTimeout); err != nil {
		requestTimeout = 10 * time.Second
	}

	var svConf serverConf
	if err := c.Decode("server", &svConf); err != nil {
		panic(err)
	}

	fmt.Printf("listening to %s on port %d\n", svConf.Address, svConf.Port)
	fmt.Printf("request timeout: %s\n", requestTimeout)
}
```
