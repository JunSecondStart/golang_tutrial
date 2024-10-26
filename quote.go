package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Hello())   // Hello, world.
    fmt.Println(quote.Go())      // Don't communicate by sharing memory, share memory by communicating.
    fmt.Println(quote.Glass())   // I can eat glass and it doesn’t hurt me.
    fmt.Println(quote.Opt())     // Concurrency is not parallelism.
}