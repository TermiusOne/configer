# Simple JSON configer for Go.

## Example

### To create config file you can use two methods:

* With creating a `configurator`:
``` go
import (
    "log"
    "github.com/TermiusOne/configer"
)

type MyConfig struct {
    DB         MyConfigDB `json:"database"`
    ServerHost string     `json:"host"`
    ServerPort int        `json:"port"`
}

type MyConfigDB struct {
    User     string `json:"user"`
    Password string `json:"password"`
    Address  string `json:"address"`
    Database string `json:"database"`
    Charset  string `json:"charset"`
}

func main() {
    conf := &MyConfig{}

    configManger := configer.New("config/config.json", conf)

    err := configManger.Create()
    if err != nil {
        log.Fatal(err)
    }
}
```

* Or using function `CreateConfig`:
``` go
func main() {
    conf := &MyConfig{}

    err := configer.CreateConfig("config/config.json", conf)
    if err != nil {
        log.Fatal(err)
    }
}
```

* You will receive this file:
``` json
{
    "database": {
        "user": "",
        "password": "",
        "address": "",
        "database": "",
        "charset": ""
    },
    "host": "",
    "port": 0
}
```
---
### You can also use two methods to read config file:

``` go
func main() {
    conf := &MyConfig{}

    configManger := configer.New("config/config.json", conf)

    err := configManger.Read()
    if err != nil {
        log.Fatal(err)
    }
}
```
* or
``` go
func main() {
    conf := &MyConfig{}

    err := configer.ReadConfig("config/config.json", conf)
    if err != nil {
        log.Fatal(err)
    }
}
```
---

### First method with create `configurator` approached if you want to create a config, add some data and use it.
