package fortnight

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("such is")
	fmt.Println(AddNumber(2, 3))
	fmt.Fprintf(os.Stdout, "%v\n", "test")
}
