package main

import (
	"fmt"

	"github.com/mxdnght0/OneCall/onecall"
)

func main() {
	group := onecall.NewGroup()
	result, err := group.Do("my-task", func() (any, error) {
		// Долгая операция
		return "done", nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
