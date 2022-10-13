# Fan-Out, Fan-in

## Benchmark
## With Fan-out, Fan-in (multiple workers)

```go
func main() {
    done := make(chan interface{})
    defer close(done)

    start := time.Now()
    generator := intGenerator(done)
    workerFunc := func() <-chan int {
        return searchPrime(done, generator)
    }
    workers := fanOut(runtime.NumCPU(), workerFunc)
    for result := range getResult(done, fanIn(done, workers...), 10) {
        fmt.Println(result)
    }
    fmt.Printf("took times: %v\n", time.Since(start))
    return
}
```

```shell
$ go run main.go
156728009
291741689
23352103
468293389
533230447
24406729
14777759
575314567
618038207
650942683
took times: 2.046196834s
```

### Without Fan-out, Fan-in (only one workers)
```go
func main() {
    done := make(chan interface{})
    defer close(done)

    start := time.Now()
    generator := intGenerator(done)
    for result := range getResult(done, searchPrime(done, generator), 10) {
        fmt.Println(result)
    }
    fmt.Printf("took times: %v\n", time.Since(start))
    return
}
```

```shell
$ go run main.go
9419987
315272537
306399283
559348289
784652819
32685659
955776977
314205113
914057623
224239787
took times: 6.783181631s
```
