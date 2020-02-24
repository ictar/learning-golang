原文：[scheduling-in-go-part3](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)

---

### 前言

这是一个包含三个部分的系列文，它将向你提供对 Go 调度器背后对机制和语义对理解。这是第三篇。本文着重于并发。

该序列文三个部分对索引：
1) [Go 调度（第一部分）：OS 调度器](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)  
2) [Go 调度（第二部分）：Go 调度器](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)  
3) [Go 调度（第三部分）：并发](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)

### 介绍

当我在解决一个问题时，特别是在它是一个新问题的时候，最初，我并不会考虑并发是否合适。首先，我会寻找顺序解决方案，并确保它能正常工作。然后，在经过可读性和技术审查后，我将开始考虑并发是否合理可行。有时候，很明显并发是一个很好的选择，而有时并没有那么明显。

在本系列的[第一部分](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)中，我解释了操作系统调度的机制和语义，如果你计划编写多线程代码的话，我认为它们对你很重要。在[第二部分](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)，我解释了 Go 调度器的语义，我认为这对于理解如何在 Go 中编写并发代码很重要。在这篇文章中，我将把操作系统和 Go 调度器的机制和语义放在一起，以便更深入地了解并发。

本文的目标是：

  * 在确定工作负载是否适合使用并发时，你必须考虑的语义指导。
  
  * 向你展示不同类型的工作负载是如何改变语义的，从而改变你要做出的工程决策。


### 并发是什么

并发意味着“乱序”执行。对一组本应该按顺序执行的指令，找到一种乱序执行并且仍然产生相同的结果的方式。摆在你面前的问题显而易见的是，乱序执行会产生价值。这里的价值是指，为复杂性成本添加足够的性能。根据你的问题，乱序执行可能并不可能，甚至于毫无意义。

同样重要的是，要了解[并发跟并行并不相同](https://blog.golang.org/concurrency-is-not-parallelism)。并行意味着同时执行两个或更多指令。这是与并发不同的概念。只有在你拥有至少两个可用的操作系统（OS）和硬件线程，并且至少有两个 Goroutine，每个 Goroutine 在每个 OS/硬件线程上独立执行指令的时候，并行才有可能。

**图 1：并发 vs 并行**  
![](https://www.ardanlabs.com/images/goinggo/96_figure1.png)

在图 1 中，你会看到两个逻辑处理器（P），每个都有它们自己独立的 OS 线程（M）。每个 M 绑定到机器上一个独立的硬件线程（Core）。你可以看到，两个 Goroutine (G1 和 G2) 在并行执行，同时在各自的 OS/硬件线程上上执行自己的指令。在每个逻辑处理器中，三个 Goroutine 轮流共享它们自己的 OS 线程。所有这些 Goroutine 都是并发运行的，以不特定的顺序执行指令，并且共享 OS 线程上的时间。

麻烦的是，有时在没有并行的情况下利用并发实际上会减慢吞吐量。有趣的是，有时将并行与并发结合使用，并不会给你带来比你原本可以实现的更大的性能提升。

### 工作负载

要如何知道什么时候乱序执行是有可能或者有意义的呢？了解你的问题正在处理的工作负载类型是一个不错的起点。在考虑并发的时候，需要了解两种类型的工作负载。

  * **计算密集型**：这种类型的工作负载永远都不会导致 Goroutine 自然地进入和退出等待态。这类型的工作会不断地在进行计算。一个计算 Pi 到第 N 位数字的线程是计算密集型的。
  
  * **IO 密集型**： 这类型的工作负载会导致 Goroutine 自然地进入等待状态。这类型的工作包括请求访问网络资源，或者对操作系统进行系统调用，又或者等待某些事件发生。一个需要读取文件的 Goroutine 是 IO 密集型的。我将同步事件（互斥、原子）这类会导致 Goroutine 等待的工作归为此类。

对于计算密集型负载，你需要并行以利用并发。使用单个 OS/硬件线程来处理多个 Goroutine 效率并不高，因为 Goroutine 的工作负载并不包含进入和离开等待态。拥有比 OS/硬件线程更多的 Goroutine 会减慢工作执行速度，因为会有将 Goroutine 移入和移出 OS 线程的延迟时间（花费的时间）。上下文切换会为你的工作创建一个“停止（Stop The World）”事件，因为在切换期间，你想进行的工作都不会执行。

对于 IO 密集型负载，你并不需要使用并行来利用并发。单个 OS/硬件线程可以高效地处理多个 Goroutine，因为 Goroutine 的工作负载包含自然地进入和离开等待态。拥有比 OS/硬件线程更多的 Goroutine 会提高工作执行速度，因为在操作系统线程中移入移出 Goroutine 期间并不会有“停止（Stop The World）”事件。你想做的工作自然地处于停止状态，这允许其他 Goroutine 有效利用同一个 OS/硬件线程，而不是让 OS/硬件线程闲置。

要如何知道每个硬件线程要有多少个 Goroutine 才能提供最佳吞吐量呢？Goroutine 太少就会有更多的空闲事件。Goroutine 太多，则会有更多的上下文切换延迟事件。这些都是你需要考虑的东西，但是超出了这篇文章的范围。

现在，重要的是看看一些代码，以巩固你识别何时可以利用并发、何时不能利用并发以及是否需要并行的能力。

### 数字相加

我们并不需要复杂代码来可视化和理解这些语义。看看下面名为 `add` 的函数，它对一个整数序列求和。

**清单 1**  
<https://play.golang.org/p/r9LdqUsEzEz>
    
    
    36 func add(numbers []int) int {
    37     var v int
    38     for _, n := range numbers {
    39         v += n
    40     }
    41     return v
    42 }
    

在清单 1 的第 36 行中，声明了一个名为 `add` 的函数，它接收一个整数集合，然后返回该集合的总和。在第 37 行，声明了变量 `v`，表示总和。然后在第 38 行，函数线性遍历集合，并在第 39 行将每个数字加到当前的总和上。最后，在第 41 行，函数将最终的和返回给调用者。

问题：`add` 函数是适合无序执行的工作负载吗？我想，问题是肯定的。整数集合可以分成较小的列表，并且可以同时处理这些列表。一旦所有这些小列表被求和，就可以对和集求和以生成与顺序执行版本相同的答案。

但是，还有另一个问题要考虑。应该创建和处理多少个更小的列表才能获得最佳吞吐量？要回答这个问题，你必须知道 `add` 正在执行的工作负载类型。`add` 函数正在执行的是计算密集型的工作负载，因为算法正在执行纯数学运算，并且它所做的任何操作都不会导致 goroutine 自然地进入等待态。这意味着，每个 OS/硬件线程使用一个 Goroutine 就可以获得良好的吞吐量。

下面的清单 2 是关于 `add` 我的并发版本。

_注意：在编写并发版本的 add 的时候，可以采用几种方式和选项。现在不要执着于我的特定实现。如果你有可读性更高，并且性能相同甚至更好的版本，很乐意看你分享。_

**清单 2**  
<https://play.golang.org/p/r9LdqUsEzEz>
    
    
    44 func addConcurrent(goroutines int, numbers []int) int {
    45     var v int64
    46     totalNumbers := len(numbers)
    47     lastGoroutine := goroutines - 1
    48     stride := totalNumbers / goroutines
    49
    50     var wg sync.WaitGroup
    51     wg.Add(goroutines)
    52
    53     for g := 0; g < goroutines; g++ {
    54         go func(g int) {
    55             start := g * stride
    56             end := start + stride
    57             if g == lastGoroutine {
    58                 end = totalNumbers
    59             }
    60
    61             var lv int
    62             for _, n := range numbers[start:end] {
    63                 lv += n
    64             }
    65
    66             atomic.AddInt64(&v, int64(lv))
    67             wg.Done()
    68         }(g)
    69     }
    70
    71     wg.Wait()
    72
    73     return int(v)
    74 }
    

在清单 2 中，`addConcurrent` 函数给出了 `add` 函数的并发版本。相比非并发版本的 5 行代码，该并发版本使用了 26 行代码。代码很多，因此我只强调需要理解的重点行。

**第 48 行**： 每个 Goroutine 都要对属于它们自己的唯一的更小的数字列表进行求和。通过将集合大小除以 Goroutine 数来获得列表大小。

**第 53 行**： 创建 Goroutine 池来进行加法操作。

**第 57-59 行**： 最后一个 Goroutine 会对剩下的数字进行求和，数字的数目可能会比其他 Goroutine 多。

**第 66 行**： 把较小列表的和加到最终的和里。

并发版本肯定会比顺序版本更复杂，这种复杂度值得吗？回答该问题的最佳方法是创建一个基准。对于这些基准，我是用了一个大小为 10000 万的数字集合，并且关闭了垃圾收集器。有使用 `add` 函数的顺序版本和使用 `addConcurrent` 函数的并发版本。

**清单 3**
    
    
    func BenchmarkSequential(b *testing.B) {
        for i := 0; i < b.N; i++ {
            add(numbers)
        }
    }
    
    func BenchmarkConcurrent(b *testing.B) {
        for i := 0; i < b.N; i++ {
            addConcurrent(runtime.NumCPU(), numbers)
        }
    }
    

清单 3 显示了基准函数。下面是对于所有的 Goroutine 只有一个 OS/硬件线程可用时的结果。顺序版本使用 1 个 Goroutine，而并发版本使用 `runtime.NumCPU`（在我的机器上是 8）个 Goroutine。在这种情况下，并发版本利用并发而没有利用并行。

**清单 4**
    
    
    10 Million Numbers using 8 goroutines with 1 core
    2.9 GHz Intel 4 Core i7
    Concurrency WITHOUT Parallelism
    -----------------------------------------------------------------------------
    $ GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
    goos: darwin
    goarch: amd64
    pkg: github.com/ardanlabs/gotraining/topics/go/testing/benchmarks/cpu-bound
    BenchmarkSequential      	    1000	   5720764 ns/op : ~10% Faster
    BenchmarkConcurrent      	    1000	   6387344 ns/op
    BenchmarkSequentialAgain 	    1000	   5614666 ns/op : ~13% Faster
    BenchmarkConcurrentAgain 	    1000	   6482612 ns/op
    

_注意： 在本地计算机上运行基准测试很复杂。有太多可能导致基准测试不准确的因素了。确保你的计算机尽可能的空闲，并多次运行基准测试。如果想确保结果一致，那么由测试工具运行两次基准测试可以使基准测试获得最一致的结果。_

清单 4 中的基准测试表明，当所有的 Goroutine 仅使用一个 OS/硬件线程时，顺序版本比并发版本大约快 10% 到 13%。这在预料之中，因为并发版本具有在在单个 OS 线程上进行上下文切换和管理 Goroutine 的开销。

下面是每个 Goroutine 对应一个 OS/硬件线程时的结果。顺序版本使用 1 个 Goroutine，而并发版本使用 `runtime.NumCPU`（在我的机器上是 8）个 Goroutine。在这种情况下，并发版本利用了并发和并行。

**清单 5**
    
    
    10 Million Numbers using 8 goroutines with 8 cores
    2.9 GHz Intel 4 Core i7
    Concurrency WITH Parallelism
    -----------------------------------------------------------------------------
    $ GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s
    goos: darwin
    goarch: amd64
    pkg: github.com/ardanlabs/gotraining/topics/go/testing/benchmarks/cpu-bound
    BenchmarkSequential-8        	    1000	   5910799 ns/op
    BenchmarkConcurrent-8        	    2000	   3362643 ns/op : ~43% Faster
    BenchmarkSequentialAgain-8   	    1000	   5933444 ns/op
    BenchmarkConcurrentAgain-8   	    2000	   3477253 ns/op : ~41% Faster
    

清单 5 中的基准测试表明，当每个 Goroutine 对应一个 OS/硬件线程时，并发版本比顺序版本快大约 41% 到 43%。这在预料之中，因为现在所有的 Goroutine 都并行运行，八个 Goroutine 同时执行其并发工作。

### 排序

重要的是，了解并非所有的计算密集型工作负载都适合并发。当分解工作以及（或者）合并所有结果的代价非常昂贵的时候，这就特别的正确。可以看到的一个例子是使用称为冒泡排序的排序算法。请看下面用 Go 实现的冒泡排序。

**清单 6**  
<https://play.golang.org/p/S0Us1wYBqG6>
    
    
    01 package main
    02
    03 import "fmt"
    04
    05 func bubbleSort(numbers []int) {
    06     n := len(numbers)
    07     for i := 0; i < n; i++ {
    08         if !sweep(numbers, i) {
    09             return
    10         }
    11     }
    12 }
    13
    14 func sweep(numbers []int, currentPass int) bool {
    15     var idx int
    16     idxNext := idx + 1
    17     n := len(numbers)
    18     var swap bool
    19
    20     for idxNext < (n - currentPass) {
    21         a := numbers[idx]
    22         b := numbers[idxNext]
    23         if a > b {
    24             numbers[idx] = b
    25             numbers[idxNext] = a
    26             swap = true
    27         }
    28         idx++
    29         idxNext = idx + 1
    30     }
    31     return swap
    32 }
    33
    34 func main() {
    35     org := []int{1, 3, 2, 4, 8, 6, 7, 2, 3, 0}
    36     fmt.Println(org)
    37
    38     bubbleSort(org)
    39     fmt.Println(org)
    40 }
    

在清单 6 中，有一个用 Go 编写的冒泡排序的示例。这个排序算法会每次都扫描整数集合并交换值。根据列表的顺序，在排序完成之前，它可能要多次遍历集合。

问题：`bubbleSort` 函数是适合无序执行的工作负载吗？我想答案是否定的。整数集合可以被分解为较小的列表，并且可以并发地对这些列表进行排序。但是，在所有并发工作完成后，并没有一个有效的方式可以把这些较小的列表合并排序在一起。下面是冒泡排序的并发版本示例。

**清单 8**
    
    
    01 func bubbleSortConcurrent(goroutines int, numbers []int) {
    02     totalNumbers := len(numbers)
    03     lastGoroutine := goroutines - 1
    04     stride := totalNumbers / goroutines
    05
    06     var wg sync.WaitGroup
    07     wg.Add(goroutines)
    08
    09     for g := 0; g < goroutines; g++ {
    10         go func(g int) {
    11             start := g * stride
    12             end := start + stride
    13             if g == lastGoroutine {
    14                 end = totalNumbers
    15             }
    16
    17             bubbleSort(numbers[start:end])
    18             wg.Done()
    19         }(g)
    20     }
    21
    22     wg.Wait()
    23
    24     // 呃……必须再次对整个列表进行排序。
    25     bubbleSort(numbers)
    26 }
    

在清单 8 中，提供了 `bubbleSort` 函数的并发版本 `bubbleSortConcurrent` 函数。它使用多个 Goroutine 来同时对列表对各个部分进行排序。但是，得到的是一列已排序块。给定一个由 36 个数字组成的列表，将其分为 12 组，如果在第 25 行没有对整个列表再次排序的话，那么这将就是结果列表。

**清单 9**
    
    
    运行前：
      25 51 15 57 87 10 10 85 90 32 98 53
      91 82 84 97 67 37 71 94 26  2 81 79
      66 70 93 86 19 81 52 75 85 10 87 49
    
    运行后：
      10 10 15 25 32 51 53 57 85 87 90 98
       2 26 37 67 71 79 81 82 84 91 94 97
      10 19 49 52 66 70 75 81 85 86 87 93
    

由于冒泡排序的本质是遍历整个列表，因此在第 25 行对 `bubbleSort` 对调用将抵消使用并发获得对任何潜在好处。对于冒泡排序，使用并发并不会提高性能。

### 读取文件

已经看了两个计算密集型负载，那么 IO 密集型呢？当 Goroutine 自然而然地进出等待态对时候，语义是否会有所不同？那么，看一个 IO 密集型的工作负载的例子，它读取文件并进行文件搜索。

第一个版本是名为 `find` 的函数的顺序版本。

**清单 10**  
<https://play.golang.org/p/8gFe5F8zweN>
    
    
    42 func find(topic string, docs []string) int {
    43     var found int
    44     for _, doc := range docs {
    45         items, err := read(doc)
    46         if err != nil {
    47             continue
    48         }
    49         for _, item := range items {
    50             if strings.Contains(item.Description, topic) {
    51                 found++
    52             }
    53         }
    54     }
    55     return found
    56 }
    

在清单 10 中，是 `find` 函数的顺序版本。在第 43 行，声明了一个名为 `found` 的变量，它维护了在给定的文件中指定的 `topic` 被找到的次数。然后，在第 44 行，遍历文档。在第 45 行，使用 `read` 函数读取每个文档。最后，在第 49 ～ 53 行，使用 `strings` 包的 `Contains` 函数来检查是否可以在从文档读取的内容中找到 topic。如果可以找到，那么 `found` 变量加一。

下面是 `read` 函数的实现，它将被 `find` 函数调用。

**清单 11**  
<https://play.golang.org/p/8gFe5F8zweN>
    
    
    33 func read(doc string) ([]item, error) {
    34     time.Sleep(time.Millisecond) // 模拟阻塞磁盘读取。
    35     var d document
    36     if err := xml.Unmarshal([]byte(file), &d); err != nil {
    37         return nil, err
    38     }
    39     return d.Channel.Items, nil
    40 }
    

清单 11 中的 `read` 函数首先调用 `time.Sleep` 休眠 1 毫秒。利用这个调用来模拟在真实场景中通过系统调用对磁盘文件的读取可能产生的延迟。这种延迟的一致性对于准确衡量 `find` 顺序版本的性能以及并发版本的性能很重要。然后，在第 35-39 行，存储在全局变量 `file` 中的模拟 xml 文件被反序列化为结构值以进行处理。最后，在第 39 行，返回一组元素给调用者。

有了顺序版本，下面的是并发版本。

_注意：在编写 find 的并发版本的时候，可以采用几种方法和选项。现在不要执着于我的特定实现。如果你有可读性更高，并且性能相同甚至更好的版本，很乐意看你分享。_

**清单 12**  
<https://play.golang.org/p/8gFe5F8zweN>
    
    
    58 func findConcurrent(goroutines int, topic string, docs []string) int {
    59     var found int64
    60
    61     ch := make(chan string, len(docs))
    62     for _, doc := range docs {
    63         ch <- doc
    64     }
    65     close(ch)
    66
    67     var wg sync.WaitGroup
    68     wg.Add(goroutines)
    69
    70     for g := 0; g < goroutines; g++ {
    71         go func() {
    72             var lFound int64
    73             for doc := range ch {
    74                 items, err := read(doc)
    75                 if err != nil {
    76                     continue
    77                 }
    78                 for _, item := range items {
    79                     if strings.Contains(item.Description, topic) {
    80                         lFound++
    81                     }
    82                 }
    83             }
    84             atomic.AddInt64(&found, lFound)
    85             wg.Done()
    86         }()
    87     }
    88
    89     wg.Wait()
    90
    91     return int(found)
    92 }
    

在清单 12 中，`findConcurrent` 函数是 `find` 函数的并发版本。相较于非并发版本的 13 行代码，并发版本使用了 30 行代码。我在实现这个并发版本的目标是控制用来处理未知个数的文档的 Goroutine 数。我的选择是使用池模式，其中，使用一个 channel 来驱动 Goroutine 池。

代码有很多，因此我只强调要理解的重要的代码行。

**第 61-64 行：**创建一个 channel，然后将所有文档发送到该 channel 以进行处理。

**第 65 行：**关闭 channel，这样，当所有的文档都被处理时，Goroutine 池自然终止。

**第 70 行：**创建 Goroutine 池。

**Line 73-83:**池中的每个 Goroutine 从 channel 中接收一个文档，将文档读取到内存中，然后检查内容查找 topic。当出现匹配时，本地变量 `lFound` 就会加一。

**第 84 行：**将每个 Goroutine 的计数总和加在一起，成为最终计数。

并发版本肯定比顺序版本更加复杂，但值得吗？再次回答这个问题但最佳方法是创建一个基准。针对这些基准，我将使用一千个文档组成的集合，并且关闭垃圾收集器。有使用 `find` 函数的顺序版本，以及使用 `findConcurrent` 函数的并发版本。

**清单 13**
    
    
    func BenchmarkSequential(b *testing.B) {
        for i := 0; i < b.N; i++ {
            find("test", docs)
        }
    }
    
    func BenchmarkConcurrent(b *testing.B) {
        for i := 0; i < b.N; i++ {
            findConcurrent(runtime.NumCPU(), "test", docs)
        }
    }
    

清单 13 显示了基准函数。以下是所有 Goroutine 只使用一个 OS/硬件线程的结果。顺序版本使用 1 个 Goroutine，而并发版本使用 `runtime.NumCPU`（在我的机器上是 8）个 Goroutine。在这种情况下，并发版本利用了并发，但没有利用并行。

**清单 14**
    
    
    10 Thousand Documents using 8 goroutines with 1 core
    2.9 GHz Intel 4 Core i7
    Concurrency WITHOUT Parallelism
    -----------------------------------------------------------------------------
    $ GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
    goos: darwin
    goarch: amd64
    pkg: github.com/ardanlabs/gotraining/topics/go/testing/benchmarks/io-bound
    BenchmarkSequential      	       3	1483458120 ns/op
    BenchmarkConcurrent      	      20	 188941855 ns/op : ~87% Faster
    BenchmarkSequentialAgain 	       2	1502682536 ns/op
    BenchmarkConcurrentAgain 	      20	 184037843 ns/op : ~88% Faster
    

清单 14 中的基准测试表明，在只有一个 OS/硬件线程可用于所有 Goroutine 的时候，并发版本比顺序版本快了大约 87% 到 88%。这在期望之中，因为所有的 Goroutine 都有效地共享了这单个 OS/硬件线程。每个 Goroutine 在`read` 调用时都发生了自然的上下文切换，这就允许了随着事件的推移，在单个 OS/硬件线程上完成更多的工作。

下面是将并发与并行结合时的基准测试。

**清单 15**
    
    
    10 Thousand Documents using 8 goroutines with 1 core
    2.9 GHz Intel 4 Core i7
    Concurrency WITH Parallelism
    -----------------------------------------------------------------------------
    $ GOGC=off go test -run none -bench . -benchtime 3s
    goos: darwin
    goarch: amd64
    pkg: github.com/ardanlabs/gotraining/topics/go/testing/benchmarks/io-bound
    BenchmarkSequential-8        	       3	1490947198 ns/op
    BenchmarkConcurrent-8        	      20	 187382200 ns/op : ~88% Faster
    BenchmarkSequentialAgain-8   	       3	1416126029 ns/op
    BenchmarkConcurrentAgain-8   	      20	 185965460 ns/op : ~87% Faster
    

清单 15 中的基准测试表明，引入额外的 OS/硬件线程并不会提供任何性能上的提升。

### 总结

本文的目的是在你确定某个工作负载是否适合使用并发的时候，为你提供必须考虑的语义指导。我试着提供不同类型的算法和工作负载示例，以便你可以看到语义上的差异以及需要考虑的不同工程决策。

你可以很清楚地看到，对于 IO 密集型的工作负载，在提高性能方面，并行并不是必须的。这与你看到的计算密集型工作截然相反。当涉及像冒泡排序这样的算法时，并发的使用会增加复杂度，却没有带来任何实质的性能提高。确定你的工作负载是否适合并发，然后识别必须使用正确的语义的工作负载类型，这非常重要。
