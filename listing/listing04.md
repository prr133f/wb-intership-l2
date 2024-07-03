#
```go
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}
```

Программа упадет с дедлоком.

Это произойдет потому что мейн горутина читает из канала, который никогда не закрывается. Пишущая горутина завершается, а мейн ждет.