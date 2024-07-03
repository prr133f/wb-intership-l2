# Что выведет программа? Объяснить вывод программы.
```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```

Этот код выведет error.

Происходит это потому, что если мы сохраним в переменную типа error структуру, реализующую интерфейс error, то этот самый интерфейс будет хранит в себе тип и значение, а следовательно никогда не будет nil.