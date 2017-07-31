#[(Unofficial) setlist.fm Go API Wrapper](https://api.setlist.fm/)

This is a Go implementation of the API spec which can be found [here](https://api.setlist.fm/).

It is configured for version 1.0 of the API and requires no external packages.

## [Installation](#install)
```
go get github.com/jm-duarte/setlistfm
```

## [Usage](#usage)
```go
package main
import (
    "fmt"
    "github.com/jm-duarte/setlistfm"
)

SETLIST_FM_API_KEY = "<api key>"

func main() {
    fmt.Println(setlistfm.ArtistByMBID("3bd680c8")) 
}
```

#This is a test version and all help is welcome!