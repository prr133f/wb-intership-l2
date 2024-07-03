#
```go
package main
 
import (
  "fmt"
)
 
func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}
 
func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}
```

Этот код выводит [3 2 3]
Происходит это потому, что функция append не изменяет передаваемый слайс, а создает копию. Получается, что после второй строчки функции modifySlice, переменная i больше не содержит исходный слайс.