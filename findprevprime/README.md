## **findprevprime**

### Instructions

Write a function that returns the first prime number that is equal or inferior to the `int` passed as parameter.

If there are no primes inferior to the `int` passed as parameter the function should return 0.
### Expected function
```bash
func FindPrevPrime(nb int) int {

}
```
### Usage

Here is a possible program to test your function :
```bash
package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
```
And its output :
```bash
$ go run .
5
3
$
```