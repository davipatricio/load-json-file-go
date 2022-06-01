## load-json-file

> Read and parse a JSON file

It also strips UTF-8 [byte order mark](https://en.wikipedia.org/wiki/Byte_order_mark#UTF-8) (BOM).

## Install

```
$ go get github.com/davipatricio/load-json-file-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/davipatricio/load-json-file-go"
)

func main() {
     // Returns a map[string]interface{}
    json, err := loader.LoadJsonFile("path/to/file.json")
    if err != nil {
        panic(err)
        return
    }

    // Print the parsed JSON
    fmt.Println(json)

    // Get a specific value
    fmt.Println(json["MyKey"].(string))
}

```

## Related

- [load-json-file](https://github.com/sindresorhus/load-json-file) - Read and parse a JSON file