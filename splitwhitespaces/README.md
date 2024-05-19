# Instructions
Write a function that separates the words of a string and puts them in a string slice.

The separators are `spaces`, `tabs` and `newlines`.

### Expected function

```go
func splitWhiteSpaces(s string) []string {

}
```

## Usage
Here is a possible program to test your function :

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%#v\n", splitWhiteSpaces("Hello how are you?"))
}
```

And its output :

```bash
$ go run .
[]string{"Hello", "how", "are", "you?"}
$
```