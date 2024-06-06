## lastrune
### Instructions

Write a function that returns the last `rune` of a `string`.

### Expected function
```go
func lastRune(s string) rune {

}
```

### Usage

Here is a possible program to test your function :
```go
package main

import (
	"github.com/01-edu/z01"

	"piscine"
)

func main() {
	z01.PrintRune(lastRune("Hello!"))
	z01.PrintRune(lastRune("Salut!"))
	z01.PrintRune(lastRune("Ola!"))
	z01.PrintRune('\n')
}
```

And its output :
```bash
$ go run .
!!!
$
```
Notions

- [rune-literals](https://golang.org/ref/spec#Rune_literals)
