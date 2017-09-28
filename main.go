package main

import (
	"fmt"
	"os"

	"github.com/Choestelus/automatic-fortnight/fortnight"
)

func main() {
	fmt.Println("such is")
	fmt.Println(fortnight.AddNumber(2, 3))
	fmt.Fprintf(os.Stdout, "%v\n", "test")
}
