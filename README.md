# GLOG: Simple Google Logging Implementation

## How to use
```go
package main

import "github.com/alexigasse/glog"

func main() {
  glog.Prinln("Hello world!")
  // os.Stdout {"message": "Hello world!"}
}
```