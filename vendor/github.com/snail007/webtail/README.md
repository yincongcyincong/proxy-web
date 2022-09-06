# webtail

golang implement of "tail -f" unix like, which in web browser , show log file content in browser real time.

## Usage

```golang
package main

import (
    "log"
    "github.com/snail007/webtail"
)

func main() {
    address := ":8822"
    basedir = "./logs"
    listener, err := webtail.Serve(address, basedir)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("tail server on %s",(*listener).Addr())
    select {}
}
```

then access :

`http://127.0.0.1:8822/show/logfilename#width=100%&height=300px`

***logfilename*** is the log file name (no extension) in `./logs` ,

all log files in `./logs` must has extension of `.log`

width=100%&height=300px is the div width and height which show log text .

the web page will show log content in real time.

## Binary

you can also use prebuild binary , you can get it [here](../..//releases)

```text
Usage of ./webtail:
  -d string
        dir path of log files
  -l string
        listen address (default ":8100")
  -v     show version
```