# OneCall

Пакет на Go для координации параллельных операций: гарантирует, что в каждый момент времени выполняется только одна задача определённого типа.

## Использование

```go
package main

import (
	"fmt"
	"github.com/mxdnght0/OneCall/onecall"
)

func main() {
	group := onecall.NewGroup()
	result, err := group.Do("my-task", func() (any, error) {
		// slow operation
		return "done", nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
```

```go
func (g *Group) Do(key string, f func() (any, error)) (any, error)
```

Выполняет функцию `f` только один раз для каждого значения `key`. Если для этого `key` функция уже выполняется - остальные вызовы будут ждать её завершения и получат тот же результат.
