# [(Unofficial) setlist.fm Go API Wrapper](https://api.setlist.fm/)

This is a Go implementation of the setlisft.fm API 1.0 specification which can be found [here](https://api.setlist.fm/). This wrapper requires no external libraries.

## [Installation](#install)
```
go get github.com/jmg-duarte/setlistfm
```

## [Usage](#usage)
```go
package main
import (
    "fmt"
    "github.com/jmg-duarte/setlistfm"
)

client := setlistfm.NewClient("api-key")

func main() {
    ctx := context.Background()
    fmt.Println(client.ArtistByMBID(ctx, "3bd680c8"))
}
```
## TODO
* Finish documentation
    * Add better documentation for types
* Add more useful methods to types
* Handle errors when parsing JSON API response
* Add more usage examples
* Do some actual testing
