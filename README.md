# errr

A very simple go package that handles errors. There is nothing crazy going on here, just moved some code I write a million times into a separate package. 

## usage

instead of writing this 400 times
```
if err != nil {
  log.Fatal(err.Error())
}
```

you can now do this
```
package main

import "github.com/ty-e-boyd/errr"

func main() {
  d, err := MightReturnError()
  errr.Fatal(err)
}
```
