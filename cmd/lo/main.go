package main

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	names := lo.Uniq([]string{"Samuel", "John", "Samuel"})

	even := lo.Filter([]int{1, 2, 3, 4}, func(x int, index int) bool {
		return x%2 == 0
	})

	sToI := lo.Map([]int64{1, 2, 3, 4}, func(x int64, index int) string {
		return strconv.FormatInt(x, 10)
	})

	fmt.Println(names)

	fmt.Println(even)

	fmt.Println(sToI)
}
