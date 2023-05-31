package main

import (
	"fmt"
	"strings"
)

package main
import "fmt"
import "strings"

func main() {
	arr := []string{"Aaa","aaa", "bbb","aAa", "ccc", "fff", "bbb", "aaa"}

	multiplied := strings.Map(arr, func(s string) string { return strings.ToLower(s) })

	fmt.Println(multiplied)

}