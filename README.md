# OneCall

Пакет на Go, который координирует одновременные операции таким образом, чтобы в каждый момент времени выполнялась только одна задача определённого типа.

## Использование

```go
package main

import (
	"fmt"
	"onecall"
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
```