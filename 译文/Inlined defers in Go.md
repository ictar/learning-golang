原文：[Inlined defers in Go](https://rakyll.org/inlined-defers/)

---

[

# Go, the unwritten parts

](https://rakyll.org/)

Articles mostly about Go and what I am currently working on. Conventions, best practices, little known practical tips. 

by [JBD](https://twitter.com/rakyll)

  * [Home](/)
  * [Archive](/archive/)
  * [About](/about/)
  * [GitHub](https://github.com/rakyll/)



Subscribe to the [feed](/index.xml). 

This work is licensed under a [Creative Commons Attribution-ShareAlike 4.0 International License](http://creativecommons.org/licenses/by-sa/4.0/). The blog is served by the amazing [Hugo](https://gohugo.io). 

# Inlined defers in Go

Mon, Jan 20, 2020

Go's `defer` keyword allows us to schedule a function to run before a function returns. Multiple functions can be deferred from a function. `defer` is often used to cleanup resources, finish function-scoped tasks, and similar. Deferring functions are great for maintability. By deferring, for example, we reduce the risk of forgetting to close the file in the rest of the program:
    
    
    func main() {
        f, err := os.Open("hello.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
    
        // The rest of the program...
    }
    

Deferring helps us by delaying the execution of the Close method while allowing us to type it when we have the right context. This is how deferred functions also help the readability of the source code.

## How defer works

Defer handles multiple functions by stacking them hence running them in LIFO order. The more deferred functions you have, the larger the stack will be.
    
    
    func main() {
    	for i := 0; i < 5; i++ {
    		defer fmt.Printf("%v ", i)
    	}
    }
    

The above program will output "`4 3 2 1 0 `" because the last deferred function will be the first one to be executed.

When a function is deferred, the variables accessed by it are stored as its arguments. For each deferred function, compiler generates a `runtime.deferproc` call at the call site and call into `runtime.deferreturn` at the return point of the function.
    
    
    0: func run() {
    1:    defer foo()
    2:    defer bar()
    3:
    4:    fmt.Println("hello")
    5: }
    

The compiler will generate code similar to below for the program above:
    
    
    runtime.deferproc(foo) // generated for line 1
    runtime.deferproc(bar) // generated for line 2
    
    // Other code...
    
    runtime.deferreturn(bar) // generated for line 5
    runtime.deferreturn(foo) // generated for line 5
    

## Defer performance

Defer used to require two expensive runtime calls explained above. This made deferring functions to be significantly more expensive than non-deferred functions. For example, consider to lock and unlock a `sync.Mutex` deferred and not-deferred.
    
    
    var mu sync.Mutex
    mu.Lock()
    
    defer mu.Unlock()
    

The program above will work 1.7x slower than the non-deferred version. Even though it only takes ~25-30 nanoseconds to lock and unlock a mutex by deferring, it makes a difference in large scale use or in cases where a function call need to be completed under XX nanoseconds.
    
    
    BenchmarkMutexNotDeferred-8   	125341258	         9.55 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMutexDeferred-8      	45980846	        26.6 ns/op	       0 B/op	       0 allocs/op
    

This overhead is why Go developers started to avoid defers in certain cases to improve performance. Unfortunately this situation make Go developers compromise readability.

## Inlining deferred functions

In the last few versions of Go, there have been gradual improvements to defer's performance. But with Go 1.14, some common cases will see a highly significant performance improvement. The compiler will generate code to inline some of the deferred functions at return points. With this improvement, calling into some deferred functions will be only as expensive as making a regular function call.
    
    
    0: func run() {
    1:    defer foo()
    2:    defer bar()
    3:
    4:    fmt.Println("hello")
    5: }
    

With the new improvements, above code will generate:
    
    
    // Other code...
    
    bar() // generated for line 5
    foo() // generated for line 5
    

It is possible to do this improvement only in static cases. For example, in a loop where the execution is determined by the input size dynamically, the compiler doesn't have the chance to generate code to inline all the deferred functions. But in simple cases (e.g. deferring at the top of the function or in conditional blocks if they are not in loops), it is possible to inline the deferred functions. With 1.14, easy cases will be inlined and runtime coordination will be only required if the compiler cannot generate code.

I already tried the Go 1.14beta with the mutex locking/unlocking example above. Deferred and non-deferred versions perform very similarly now:
    
    
    BenchmarkMutexNotDeferred-8   	123710856	         9.64 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMutexDeferred-8      	104815354	        11.5 ns/op	       0 B/op	       0 allocs/op
    

Go 1.14 is a good time to reevaluate deferring if you avoided defers for performance gain. If you are looking for more about this improvement, see the [Low-cost defers through inline code](https://github.com/golang/proposal/blob/master/design/34481-opencoded-defers.md) proposal and [GoTime's recent episode on defer with Dan Scales](https://changelog.com/gotime/112).

* * *

Disclaimer: This article is not peer-reviewed but thanks to Dan Scales for answering my questions while I was investigating this improvement.

