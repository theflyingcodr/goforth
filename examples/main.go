package main

import (
	"fmt"

	"github.com/theflyingcodr/goforth"
)

func main() {
	fmt.Println(goforth.Run("1 10 ADD 11 EQ HASH256"))
}
