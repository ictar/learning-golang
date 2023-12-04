原文：[垃圾回收 In Go : Part III - GC Pacing](https://www.ardanlabs.com/blog/2019/07/garbage-collection-in-go-part3-gcpacing.html)

---

### 前言

这是一个包含三个部分的系列文，它将向你提供对 Go 垃圾回收器背后对机制和语义对理解。这是第一篇。本文着重于 GC Pacing 本身。

该序列文三个部分对索引： 
1) [Go 垃圾回收（第一部分）：语义](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)  
2) [Go 垃圾回收（第二部分）：GC 跟踪](https://www.ardanlabs.com/blog/2019/05/garbage-collection-in-go-part2-gctraces.html)  
3) [Go 垃圾回收（第三部分）：GC Pacing](https://www.ardanlabs.com/blog/2019/07/garbage-collection-in-go-part3-gcpacing.html)

### 介绍

在第二篇文章中，我向你展示了垃圾回收器的行为，以及如何使用工具查看回收器对正在运行中的应用造成的延迟。我带你看了一个实际的 web 应用，并且向你展示了如何生成 GC 跟踪和应用 profile 信息。然后，还向你展示了如何理解这些工具的输出，以便你找到提高应用性能的方法。

那篇文章的最终结论与第一篇文章的结论相同：如果减轻了堆压力，就会减少延迟成本，从而提高应用性能。减小回收器的延迟的最佳策略就是减小每次执行的工作的分配次数和分配数量。在这篇文章中，我将展示 pacing 算法如何能够随着时间的推移，确定给定工作负载的最佳节奏。

### 并发示例代码

我将使用此链接里的代码。

<https://github.com/ardanlabs/gotraining/tree/master/topics/go/profiling/trace>

该程序确定特定主题在 RSS 新闻提要文档集中出现的频率。跟踪程序包含不同版本的查找算法，以测试不同的并发模式。我将专注于这些算法的 `freq`、`freqConcurrent` 和 `freqNumCPU` 版本。

_注意：代码运行在 Macbook Pro 上，使用带有 12 个硬件线程的 Intel i9 处理器，使用 go1.12.7。在不同的体系架构、操作系统和 Go 版本上，你会看到不同的结果。本文的核心结果应该保持不变。_

我将先从 `freq` 版本开始。它代表程序的一个非并发的顺序版本。这将为后面的并发版本提供基准。

**清单 1**
    
    
    01 func freq(topic string, docs []string) int {
    02     var found int
    03
    04     for _, doc := range docs {
    05         file := fmt.Sprintf("%s.xml", doc[:8])
    06         f, err := os.OpenFile(file, os.O_RDONLY, 0)
    07         if err != nil {
    08             log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
    09             return 0
    10         }
    11         defer f.Close()
    12
    13         data, err := ioutil.ReadAll(f)
    14         if err != nil {
    15             log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
    16             return 0
    17         }
    18
    19         var d document
    20         if err := xml.Unmarshal(data, &d); err != nil {
    21             log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
    22             return 0
    23         }
    24
    25         for _, item := range d.Channel.Items {
    26             if strings.Contains(item.Title, topic) {
    27                 found++
    28                 continue
    29             }
    30
    31             if strings.Contains(item.Description, topic) {
    32                 found++
    33             }
    34        }
    35     }
    36
    37     return found
    38 }
    

清单 1 显示了 `freq` 函数。这个顺序版本遍历文件名集合，然后执行四种操作：打开、读取、解码和搜索。每个文件执行一次，一次执行一个文件。

当我在我的机器上运行 `freq` 的这个版本时，会得到以下结果。

**清单 2**
    
    
    $ time ./trace
    2019/07/02 13:40:49 Searching 4000 files, found president 28000 times.
    ./trace  2.54s user 0.12s system 105% cpu 2.512 total
    

你可以通过 time 命令的输出看到，该程序处理 4000 个文件大约需要 2.5 秒。如果能看到垃圾回收花费的时间占了多少百分比，那就太好了。你可以通过查看该程序的跟踪来实现。由于这是一个启动之后会完成执行的程序，因此可以使用 trace 包来生成跟踪。

**清单 3**
    
    
    03 import "runtime/trace"
    04
    05 func main() {
    06     trace.Start(os.Stdout)
    07     defer trace.Stop()
    

清单 3 显示了在应用中生成跟踪所需的代码。在从标准库的 `runtime` 目录中导入 `trace` 包后，调用 `trace.Start` 和 `trace.Stop`。为了简单起见，将跟踪输出到 `os.Stdout`。

有了这个代码，现在你就可以重新构建然后再次运行程序了。不要忘了把 `stdout` 重定向到文件中。

**清单 4**
    
    
    $ go build
    $ time ./trace > t.out
    Searching 4000 files, found president 28000 times.
    ./trace > t.out  2.67s user 0.13s system 106% cpu 2.626 total
    

运行时间增加了 100 毫秒多一点，但这在预料之中。跟踪会捕获每一个函数调用（从进入函数到出函数），这是毫秒级的。重要的是，现在，有一个名为 `t.out` 的文件，它包含跟踪数据。

要查看这个跟踪信息，需要使用跟踪工具。

**清单 5**
    
    
    $ go tool trace t.out
    

运行该命令会启动 Chrome 浏览器，如下所示。

_注意： 跟踪工具使用的是 Chrome 浏览器内置的工具。此工具仅适用于 Chrome。_

**图 1**  
![](https://www.ardanlabs.com/images/goinggo/103_figure1.png)

图 1 显示了启动跟踪工具时会显示的 9 个链接。现在，重要的是第一个链接，即 `View trace`。一旦你选择了这个链接，将会看到与以下相似的内容。

**图 2**  
![](https://www.ardanlabs.com/images/goinggo/103_figure2.png?v2)

图 2 显示了在我的机器傻姑娘运行程序的完整跟踪窗口。在这篇文章中，我将重点介绍与垃圾回收器相关的部分。也就是标记为 `Heap` 的第二部分和标记为 `GC` 的第四部分。

**图 3**  
![](https://www.ardanlabs.com/images/goinggo/103_figure3.png)

图 3 详细显示了跟踪的前 200 毫秒。将注意力集中到 `Heap`（绿色和橙色区域）和 `GC`（底部的蓝色线条）。`Heap` 这一部分告诉你两件事。橙色区域是在任何给定毫秒内，堆上当前正在使用的空间。绿色区域是将会触发下一次回收的正在使用的堆内存的大小。这就是为什么每次橙色区域到达绿色区域的顶部时，就会进行垃圾回收。蓝色线条代表一次垃圾回收。

在这个版本的程序中，整个程序运行期间，堆中正在使用的内存保持在大约 4 meg。想要查看所有单个垃圾回收的统计信息，请使用选择工具，并在所有蓝色线条周围绘制一个框。

**图 4**  
![](https://www.ardanlabs.com/images/goinggo/103_figure4.png)

图 4 显示了如何使用箭头工具在蓝色线条周围绘制一个蓝色框框。你想要在每条线周围画一个框。框内的数字表示从图中所选的项目所花费的时间。在这种情况下，选择了近 316 毫秒 (ms, μs, ns) 来生成此图像。选择所有蓝线后，会提供以下统计信息。

**图 5**  
![](https://www.ardanlabs.com/images/goinggo/103_figure5.png)

图 5 显示，图中所有的蓝色线条位于 15.911 毫秒标记到 2.596 秒标记之间。有 232 次垃圾回收，花费时间 64.524 毫秒，平均每次回收花费 287.121 微秒。应用的运行时间是 2.626 秒，意味着垃圾回收仅占大约 2% 的总运行时间。从本质上来讲，垃圾回收器的运行成本对于程序整体的运行微不足道。

有了基线，就可以使用并发算法来进行相同的工作，以期加快程序运行速度。

**清单 6**
    
    
    01 func freqConcurrent(topic string, docs []string) int {
    02     var found int32
    03
    04     g := len(docs)
    05     var wg sync.WaitGroup
    06     wg.Add(g)
    07
    08     for _, doc := range docs {
    09         go func(doc string) {
    10             var lFound int32
    11             defer func() {
    12                 atomic.AddInt32(&found, lFound)
    13                 wg.Done()
    14             }()
    15
    16             file := fmt.Sprintf("%s.xml", doc[:8])
    17             f, err := os.OpenFile(file, os.O_RDONLY, 0)
    18             if err != nil {
    19                 log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
    20                 return
    21             }
    22             defer f.Close()
    23
    24             data, err := ioutil.ReadAll(f)
    25             if err != nil {
    26                 log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
    27                 return
    28             }
    29
    30             var d document
    31             if err := xml.Unmarshal(data, &d); err != nil {
    32                 log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
    33                 return
    34             }
    35
    36             for _, item := range d.Channel.Items {
    37                 if strings.Contains(item.Title, topic) {
    38                     lFound++
    39                     continue
    40                 }
    41
    42                 if strings.Contains(item.Description, topic) {
    43                     lFound++
    44                 }
    45             }
    46         }(doc)
    47     }
    48
    49     wg.Wait()
    50     return int(found)
    51 }
    

清单 6 显示了 `freq` 的一个可能的并发版本。此版本的核心设计模式是使用扇出模式（fan out pattern）。对于 `docs` 集合中列出的每个文件，都会创建一个 goroutine 来处理这个文件。如果要处理 4000 个文档，则会使用 4000 个 goroutine。这种算法的优势是，这是利用并发最简单的方法。每个 goroutine 处理并仅处理一个文件。可以使用 `WaitGroup`  等待每个文档处理完成，并且可以使用一个原子指令来保持计数器同步。

该算法的缺点是，它不能随着文档数和内核数很好地扩展。程序启动后会非常快地让所有的 goroutine 都有时间运行，这意味着会迅速消耗大量内存。在第 12 行对 `found` 变量进行增加操作仍然存在缓存一致性问题。由于每个内核都因为此变量共享相同的高速缓存线，这将导致内存崩溃。随着文件或者内核数的增加，情况会变得更加糟糕。

有了代码，现在你就可以重新构建并运行该程序了。

**清单 7**
    
    
    $ go build
    $ time ./trace > t.out
    Searching 4000 files, found president 28000 times.
    ./trace > t.out  6.49s user 2.46s system 941% cpu 0.951 total
    

你可以从清单 7 的输出中看到，处理相同 4000 个文件，程序现在需要花费 951 毫秒。性能提高了大约 64%。看一下跟踪。

**图 6**  
![](https://www.ardanlabs.com/images/goinggo/103_figure6.png)

图 6 显示了，在我的机器上，这个版本的程序使用了额外多少 CPU 容量。图的开头非常密集。这是因为，在创建所有 goroutine 时，它们会运行并开始尝试在堆中分配内存。一旦分配了前 4 meg 内存（这很快），就会启动 GC。在此 GC 期间，每个 Goroutine 都会获得运行时间，并且大多数 Goroutine 请求堆内存时会进入等待态。在这次 GC 完成时，至少有 9 个 goroutine 可以继续运行，并且将堆增长到大约 26 meg。

**图 7**  
![](https://www.ardanlabs.com/images/goinggo/103_figure7.png)

图 7 显示了，在第一次 GC 的大部分时间里，很大一部分 goroutine 处于可运行态和运行态，并且它是如何迅速重新启动的。请注意，堆的 profile 看起来是不规则的，并且回收并没有像之前那样有规律地进行。如果仔细观察，第一次 GC 之后几乎立即启动了第二次 GC。

如果你选择该图中的所有回收，那么会看到以下内容。

**图 8**  
![](https://www.ardanlabs.com/images/goinggo/103_figure8.png)


图 8 显示，图中所有的蓝色线条位于 4.828 毫秒标记到 906.939 毫秒标记之间。有 23 次垃圾回收，花费时间 284.447 毫秒，平均每次回收花费 12.367 微秒。应用的运行时间是 951 毫秒，意味着垃圾回收仅占大约 34% 的总运行时间。

与顺序版本相比，该版本在性能和 GC 时间上都存在显著差异。但是，以并行的方式运行更多的 goroutine 可以使工作的完成速度快大约 64%。成本是需要更多的机器资源。不幸的是，在高峰时期，堆上正在使用的内存约为 200 meg

有了一个并发的基准，下一个并发算法将尝试提高资源利用率。

**清单 8**
    
    
    01 func freqNumCPU(topic string, docs []string) int {
    02     var found int32
    03
    04     g := runtime.NumCPU()
    05     var wg sync.WaitGroup
    06     wg.Add(g)
    07
    08     ch := make(chan string, g)
    09
    10     for i := 0; i < g; i++ {
    11         go func() {
    12             var lFound int32
    13             defer func() {
    14                 atomic.AddInt32(&found, lFound)
    15                 wg.Done()
    16             }()
    17
    18             for doc := range ch {
    19                 file := fmt.Sprintf("%s.xml", doc[:8])
    20                 f, err := os.OpenFile(file, os.O_RDONLY, 0)
    21                 if err != nil {
    22                     log.Printf("Opening Document [%s] : ERROR : %v", doc, err)
    23                     return
    24                 }
    25
    26                 data, err := ioutil.ReadAll(f)
    27                 if err != nil {
    28                     f.Close()
    29                     log.Printf("Reading Document [%s] : ERROR : %v", doc, err)
    23                     return
    24                 }
    25                 f.Close()
    26
    27                 var d document
    28                 if err := xml.Unmarshal(data, &d); err != nil {
    29                     log.Printf("Decoding Document [%s] : ERROR : %v", doc, err)
    30                     return
    31                 }
    32
    33                 for _, item := range d.Channel.Items {
    34                     if strings.Contains(item.Title, topic) {
    35                         lFound++
    36                         continue
    37                     }
    38
    39                     if strings.Contains(item.Description, topic) {
    40                         lFound++
    41                     }
    42                 }
    43             }
    44         }()
    45     }
    46
    47     for _, doc := range docs {
    48         ch <- doc
    49     }
    50     close(ch)
    51
    52     wg.Wait()
    53     return int(found)
    54 }
    

清单 8 显示了程序的 `freqNumCPU` 版本。该版本的核心设计模式是使用池。使用基于逻辑处理器数量的 goroutine 池来处理所有的文件。如果有 12 个可用的逻辑处理器，那么就会使用 12 个 goroutine。该算法的优势在于，它使得程序的资源使用始终如一。由于使用了固定数量的 goroutine，因此只需要任意时间内这 12 个 goroutine 所需的内存。这也解决了内存抖动导致的缓存一致性问题。因为在第 14 行上对原子指令的调用只会发生固定的几次。

该算法的缺点是复杂度更高。它使用了 channel 来驱动 goroutine 池工作。在任何使用池的时间里，为池标识 goroutine 的“正确”数量都是很复杂的。通常，我会为每个逻辑处理器启动 1 个 goroutine。然后进行负载测试或者使用生产指标，然后就可以算出池的最终大小。

有了代码，现在你就可以重新构建并运行程序了。

**清单 9**
    
    
    $ go build
    $ time ./trace > t.out
    Searching 4000 files, found president 28000 times.
    ./trace > t.out  6.22s user 0.64s system 909% cpu 0.754 total
    

你可以从清单 9 的输出中看到，程序现在需要 754 毫秒来处理相同的 4000 个文件。程序快了 200 毫秒左右，对于这样小的负载而言，是非常显著的。看一下跟踪。

**图 9**  
![](https://www.ardanlabs.com/images/goinggo/103_figure9.png)

图 9 显示了这个版本的程序是如何使用我机器上所有的 CPU 容量的。如果仔细观察，程序再次具有一致的节奏。这与顺序版本非常像。

**图 10**  
![](https://www.ardanlabs.com/images/goinggo/103_figure10.png)

图 10 显示了程序的前 20 毫秒更为详细的核心指标。回收肯定比顺序版本长，但是有 12 个 goroutine 在运行。在程序的整个运行过程中，使用中的堆内存维持在 4 meg 左右。同样，再次与程序的顺序版本相同。

如果你选择图中所有的回收，将会看到以下内容。

**图 11**  
![](https://www.ardanlabs.com/images/goinggo/103_figure11.png)

图 11 显示，图中所有的蓝色线条位于 3.055 毫秒标记到 719.928 毫秒标记之间。有 467 次垃圾回收，花费时间 177.709 毫秒，平均每次回收花费 380.535 微秒。应用的运行时间是 754 毫秒，意味着垃圾回收仅占大约 25% 的总运行时间。与其他并发版本相比提高了 9%。

看起来，这个版本的并发算法可以通过更多的文件和内核来更好地扩展，在我看来，复杂度成本是值得的。可以通过切分列表来替代 channel。虽然可以减少因为 channel 引发的某些延迟成本，但这肯定会增加更多的复杂性。在更多的文件和内核的情况下，这可能很重要，但是也需要衡量复杂性成本。你可以尝试一下。

### 总结

我喜欢比较三个版本的算法的原因是看看 GC 是如何处理每种情况的。对于每个版本，处理文件所需的内存总量不会改变。改变的是程序的分配方式。

当只有一个 goroutine 的时候，只需要一个大小为 4 meg 的堆。当程序一次性运行所有的工作时，GC 采取的策略是让堆增长，减少回收次数，但是每次回收的执行时间更长。当程序控制了任意给定时间内处理的文件数目时，GC 采取的策略是再次保持较小的堆，增加回收次数，但每次回收的执行时间更短。GC 采取的每种策略本质上都会使 GC 对程序运行产生最小的影响。

    
    | 算法  | 程序 | GC 时间  | GC 百分比 | GC 次数 | 平均 GC   | 最大堆 |
    |------------|---------|----------|---------|-----------|----------|----------|
    | freq       | 2626 ms |  64.5 ms |     ~2% |       232 |   278 μs |    4 meg |
    | concurrent |  951 ms | 284.4 ms |    ~34% |        23 |  12.3 ms |  200 meg |
    | numCPU     |  754 ms | 177.7 ms |    ~25% |       467 | 380.5 μs |    4 meg |
    

`freqNumCPU` 版本还有其他好处，例如更好地处理缓存一致性。但是，每个程序的 GC 时间总量差异相当小，~284.4 ms vs ~177.7 ms。某些时候，在我的机器上运行这些程序的时候，这些数字甚至更接近。使用 1.13.beta1 版本进行一些实验的过程中，我已经看到这些算法的运行时间相同。这可能暗示着，即将到来的一些改进会使 GC 更好地预测运行方式。

所有这些都给我以信心在运行过程中处理大量的工作。例如使用 50k 个 goroutine 的 web 服务，它实际上是一种类似于第一个并发算法的扇出模式。GC 将学习工作负载，然后找出能够让服务不受限的最佳执行节奏。至少对我来说，不需要考虑任何这些事情就已经值回票价了。
