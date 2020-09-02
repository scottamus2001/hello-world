package main

import (
	"fmt"

	"github.com/scottamus2001/hello-world/locale"
)

func main() {
	fmt.Println(locale.Hello(locale.ENGLISH), "world")
	fmt.Println(locale.Hello(locale.ITALIAN), "world")
	fmt.Println(locale.Hello(locale.SPANISH), "world")
}
