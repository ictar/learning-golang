原文：[Garbage Collection In Go : Part II - GC Traces](https://www.ardanlabs.com/blog/2019/05/garbage-collection-in-go-part2-gctraces.html)

---

### 前言

这是一个包含三个部分的系列文，它将向你提供对 Go 垃圾回收器背后对机制和语义对理解。这是第一篇。本文着重于如何生成 GC 跟踪并理解它们。

该序列文三个部分对索引： 
1) [Go 垃圾回收（第一部分）：语义](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)  
2) [Go 垃圾回收（第二部分）：GC 跟踪](https://www.ardanlabs.com/blog/2019/05/garbage-collection-in-go-part2-gctraces.html)  
3) [Go 垃圾回收（第三部分）：GC Pacing](https://www.ardanlabs.com/blog/2019/07/garbage-collection-in-go-part3-gcpacing.html)

### 介绍

在第一篇文章中，我花时间描述了垃圾回收器的行为，并展示了回收器对正在运行对应用造成的延迟。我分享了生成和理解 GC 跟踪的方式，展示了堆上的内存是如何变化的，并且解释了 GC 的不同阶段以及它们是如何影响延迟成本的。

第一篇文章的最终结论是，如果你减轻堆的压力，就会减少延迟成本，从而提高应用性能。我还指出，找到方法增加两次回收之间的时间来降低回收节奏，这并不是一个好的策略。保持一致的节奏（即使很快）会更好地让应用保持最佳性能运行。

在这篇文章中，我将带你运行一个实际的 web 应用，并向你展示如何生成 GC 跟踪和应用的 profile 信息。然后，我将展示如何理解这些工具的输出，以便找到提高应用性能的方法。

### 运行应用

看看这个我在 Go 培训中使用的 web 应用。

**图 1**  
![](https://www.ardanlabs.com/images/goinggo/101_figure1.png)

<https://github.com/ardanlabs/gotraining/tree/master/topics/go/profiling/project>

图 1 显示了这个应用的外观。该应用从不同的新闻提供商那里下载三组 rss feed，并允许用户进行搜索。构建 web 应用后，启动它。

**清单 1**
    
    
    $ go build
    $ GOGC=off ./project > /dev/null
    

清单 1 显示了如何在将 `GOGC` 变量设置为 `off`（关闭垃圾收集） 的情况下启动应用。日志被重定向到 `/dev/null`，随着应用的运行，可以将请求发送到服务器上。

**清单 2**
    
    
    $ hey -m POST -c 100 -n 10000 "http://localhost:5000/search?term=topic&cnn=on&bbc=on&nyt=on"
    

清单 2 显示了如何使用 `hey` 工具，通过与服务器的 100 个链接发送 10k 请求。一旦所有的请求发送完成，就会产生以下结果。

**图 2**  
![](https://www.ardanlabs.com/images/goinggo/101_figure2.png)

图 2 显示了在关闭垃圾回收器的情况下，处理 10k 请求的可视化表示。处理 10k 请求花费了 4,188ms，也就是说，服务器每秒处理约 2,387 个请求。

### 打开垃圾回收

如果为该应用打开垃圾回收功能，会发生什么呢？

**清单 3**
    
    
    $ GODEBUG=gctrace=1 ./project > /dev/null
    

清单 3 显示了如何在可以看到 GC 跟踪的情况下启动应用。移除 `GOGC` 变量，并且用 `GODEBUG` 变量进行替换。设置 `GODEBUG` 的话，运行时就会在每次回收时生成 GC 跟踪。现在，可以再次运行相同的 10k 请求。一旦所有的请求都发送完成，就可以分析 `hey` 工具提供的 GC 跟踪和信息。

**清单 4**
    
    
    $ GODEBUG=gctrace=1 ./project > /dev/null
    gc 3 @3.182s 0%: 0.015+0.59+0.096 ms clock, 0.19+0.10/1.3/3.0+1.1 ms cpu, 4->4->2 MB, 5 MB goal, 12 P
    .
    .
    .
    gc 2553 @8.452s 14%: 0.004+0.33+0.051 ms clock, 0.056+0.12/0.56/0.94+0.61 ms cpu, 4->4->2 MB, 5 MB goal, 12 P
    

清单 4 显示了开始运行后第三次和最后一次回收的 GC 跟踪。因为请求是在头两次回收后才进行的，因此这里我并没有列出它们。最后一次回收显示，处理这 10k 请求花了 2551 次回收操作（因为头两次回收不计在内，因此减去）。

这个跟踪中的每一个部分的细分。

**清单 5**
    
    
    gc 2553 @8.452s 14%: 0.004+0.33+0.051 ms clock, 0.056+0.12/0.56/0.94+0.61 ms cpu, 4->4->2 MB, 5 MB goal, 12 P
    
    gc 2553     : 自程序启动后运行的第 2553 次 GC
    @8.452s     : 自程序启动后的八秒
    14%         : 目前 GC 花费了 14% 的可用 CPU
    
    // wall-clock
    0.004ms     : STW        : 写屏障 —— 等待所有的 P 到达 GC 安全点。
    0.33ms      : Concurrent : 标记
    0.051ms     : STW        : 标记终止     - 关闭写屏障并清理。
    
    // CPU 时间
    0.056ms     : STW        : 写屏障
    0.12ms      : Concurrent : 标记 - 辅助时间（执行 GC 与分配一致）
    0.56ms      : Concurrent : 标记 - 后台 GC 时间
    0.94ms      : Concurrent : 标记 - 空闲 GC 时间
    0.61ms      : STW        : 标记终止
    
    4MB         : 标记开始前使用中的堆内存
    4MB         : 标记结束后使用中的堆内存
    2MB         : 标记结束后被标记为活跃的堆内存
    5MB         : 标记结束后使用中的堆内存的回收目标
    
    // 线程
    12P         : 用来运行 Goroutine 的逻辑处理器或者线程数


清单 5 显示了最后一次回收的实际数字。多亏了 `hey`，下面是运行的性能结果。

**清单 6**
    
    
    Requests            : 10,000
    ------------------------------------------------------
    Requests/sec        : 1,882 r/s   - Hey
    Total Duration      : 5,311ms     - Hey
    Percent Time in GC  : 14%         - GC Trace
    Total Collections   : 2,551       - GC Trace
    ------------------------------------------------------
    Total GC Duration   : 744.54ms    - (5,311ms * .14)
    Average Pace of GC  : ~2.08ms     - (5,311ms / 2,551)
    Requests/Collection : ~3.98 r/gc  - (10,000 / 2,511)
    

清单 6 显示了结果。下面提供了更多的可视化信息。

**图 3**  
![](https://www.ardanlabs.com/images/goinggo/101_figure3.png)

图 3 直观地显示了发生的情况。打开回收器后，必须运行大约 2.5k 次回收才能处理相同的 10k 请求。每次回收的平均启动间隔约为 2.0ms，运行所有这些回收会额外增加大约 1.1 秒的延迟。

**图 4**  
![](https://www.ardanlabs.com/images/goinggo/101_figure4.png)

图 4 显示了目前应用的两次运行情况的对比。

### 减少分配

获取堆的 profile 信息，然后看看是否有可以删除的无用分配。

**清单 7**
    
    
    go tool pprof http://localhost:5000/debug/pprof/allocs
    

清单 7 显示了使用 `pprof` 工具，调用 `/debug/pprof/allocs`，来从正在运行的应用那里获取内存 profile。由于以下代码，故而存在这个请求路径。

**清单 8**
    
    
    import _ "net/http/pprof"
    
    go func() {
        http.ListenAndServe("localhost:5000", http.DefaultServeMux)
    }()
    

清单 8 显示了如何把 `/debug/pprof/allocs` 绑定到任何应用上。添加 `net/http/pprof` 导入会将这个请求路径绑定到默认服务器上。然后使用 `http.ListenAndServer` 并传入常量 `http.DefaultServerMux` 就能使得这个请求路径kneeing

一旦启动探查器，就可以使用 `top` 命令查看分配最多的 6 个函数。

**清单 9**
    
    
    (pprof) top 6 -cum
    Showing nodes accounting for 0.56GB, 5.84% of 9.56GB total
    Dropped 80 nodes (cum <= 0.05GB)
    Showing top 6 nodes out of 51
          flat  flat%   sum%        cum   cum%
             0     0%     0%     4.96GB 51.90%  net/http.(*conn).serve
        0.49GB  5.11%  5.11%     4.93GB 51.55%  project/service.handler
             0     0%  5.11%     4.93GB 51.55%  net/http.(*ServeMux).ServeHTTP
             0     0%  5.11%     4.93GB 51.55%  net/http.HandlerFunc.ServeHTTP
             0     0%  5.11%     4.93GB 51.55%  net/http.serverHandler.ServeHTTP
        0.07GB  0.73%  5.84%     4.55GB 47.63%  project/search.rssSearch
    

清单 9 显示，在列表底部，出现了 `rssSearch` 函数。迄今为止，这个函数分配了 5.96GB 中的 4.55GB。接下来，使用 `list` 命令查看 `rssSearch` 函数的细节。

**清单 10**
    
    
    (pprof) list rssSearch
    Total: 9.56GB
    ROUTINE ======================== project/search.rssSearch in project/search/rss.go
       71.53MB     4.55GB (flat, cum) 47.63% of Total
    
    
             .          .    117:	// Capture the data we need for our results if we find ...
             .          .    118:	for _, item := range d.Channel.Items {
             .     4.48GB    119:		if strings.Contains(strings.ToLower(item.Description), strings.ToLower(term)) {
       48.53MB    48.53MB    120:			results = append(results, Result{
             .          .    121:				Engine:  engine,
             .          .    122:				Title:   item.Title,
             .          .    123:				Link:    item.Link,
             .          .    124:				Content: item.Description,
             .          .    125:			})
    

清单 10 显示了代码。第 119 行突出显示了大部分分配。

**清单 11**
    
    
             .     4.48GB    119:		if strings.Contains(strings.ToLower(item.Description), strings.ToLower(term)) {
    

清单 11 显示了有问题的代码行。到目前为止，仅该行就占据了该函数分配的 4.55GB 内存中的 4.48GB。接下来，是时候审查这行代码，看看有什么可做的了。

**清单 12**
    
    
    117 // Capture the data we need for our results if we find the search term.
    118 for _, item := range d.Channel.Items {
    119     if strings.Contains(strings.ToLower(item.Description), strings.ToLower(term)) {
    120         results = append(results, Result{
    121             Engine:  engine,
    122             Title:   item.Title,
    123             Link:    item.Link,
    124             Content: item.Description,
    125        })
    126    }
    127 }
    

清单 12 显示了，这行代码位于一个紧密循环中。`strings.ToLower` 调用正在进行分配，因为它们创建了需要在堆上分配的新字符串。这些 `strings.ToLower` 调用是不必要的，因为它们可以在循环外部完成。

可以更改第 119 行，以删除所有这些分配。

**清单 13**
    
    
    // 代码修改前。
    if strings.Contains(strings.ToLower(item.Description), strings.ToLower(term)) {
    
    // 代码修改后。
    if strings.Contains(item.Description, term) {
    

_注意： 你看不到的其他代码更改是，在将提要放入缓存之前，将 Description 修改为小写。新闻提要每 15 分钟缓存一次，将 `term` 修改为小写的操作在循环外完成。_

清单 13 显示了如何移除 `strings.ToLower` 调用。修改后再次构建项目，然后再次通过服务器运行 10k 请求。

**清单 14**
    
    
    $ go build
    $ GODEBUG=gctrace=1 ./project > /dev/null
    gc 3 @6.156s 0%: 0.011+0.72+0.068 ms clock, 0.13+0.21/1.5/3.2+0.82 ms cpu, 4->4->2 MB, 5 MB goal, 12 P
    .
    .
    .
    gc 1404 @8.808s 7%: 0.005+0.54+0.059 ms clock, 0.060+0.47/0.79/0.25+0.71 ms cpu, 4->5->2 MB, 5 MB goal, 12 P
    

清单 14 显示了，在修改代码后，现在处理相同 10k 请求需要进行 1402 次回收。下面是这两次运行的完整结果。

**清单 15**
    
    
    With Extra Allocations              Without Extra Allocations
    ======================================================================
    Requests            : 10,000        Requests            : 10,000
    ----------------------------------------------------------------------
    Requests/sec        : 1,882 r/s     Requests/sec        : 3,631 r/s
    Total Duration      : 5,311ms       Total Duration      : 2,753 ms
    Percent Time in GC  : 14%           Percent Time in GC  : 7%
    Total Collections   : 2,551         Total Collections   : 1,402
    ----------------------------------------------------------------------
    Total GC Duration   : 744.54ms      Total GC Duration   : 192.71 ms
    Average Pace of GC  : ~2.08ms       Average Pace of GC  : ~1.96ms
    Requests/Collection : ~3.98 r/gc    Requests/Collection : 7.13 r/gc
    

清单 15 显示了与上一次运行的结果对比。以下内容提供了更多的可视化信息。

**图 5**  
![](https://www.ardanlabs.com/images/goinggo/101_figure5.png)

图 5 直观地显示了这些场景。这一次，处理相同的 10k 请求，回收器少运行了 1149 次（1,402 vs 2,551）。这导致总 GC 时间的百分比从 14% 降低到 7%。这使得应用运行速度提高了 48%，而回收时间却减少了 74%。

**图 6**  
![](https://www.ardanlabs.com/images/goinggo/101_figure6.png)

图 6 显示了该应用所有不同运行之间的对比。这里，我还包含了在关闭垃圾回收器的情况下优化代码的运行情况。

### 我们学到了什么

正如我在上一篇文章中所述的，减小回收器的延迟就是减小堆内存的压力。请记住，压力可以被定义为应用在给定时间内的堆内存分配速度。当压力减小时，回收器造成的延迟将减小。是 GC 延迟拖慢了你的应用。

这不是要放慢回收节奏，而是要在每次回收之间或者回收期间完成更多的工作。通过减少堆分配的数量和分配次数，可以影响到这一点。

**清单 16**
    
    
    With Extra Allocations              Without Extra Allocations
    ======================================================================
    Requests            : 10,000        Requests            : 10,000
    ----------------------------------------------------------------------
    Requests/sec        : 1,882 r/s     Requests/sec        : 3,631 r/s
    Total Duration      : 5,311ms       Total Duration      : 2,753 ms
    Percent Time in GC  : 14%           Percent Time in GC  : 7%
    Total Collections   : 2,551         Total Collections   : 1,402
    ----------------------------------------------------------------------
    Total GC Duration   : 744.54ms      Total GC Duration   : 192.71 ms
    Average Pace of GC  : ~2.08ms       Average Pace of GC  : ~1.96ms
    Requests/Collection : ~3.98 r/gc    Requests/Collection : 7.13 r/gc
    

清单 16 显示了打开垃圾回收运行应用的两个版本。很明显，移除 4.48GB 的分配可以使应用运行得更快些。有趣的是，（对于两个版本）每次回收的平均节奏实际上是相同的，大约是 2.0ms。这两个版本之间的根本变化是，每次回收之间完成的工作量。应用的处理能力从 3.98 r/gc 变成 7.13 r/gc。完成的工作量增加了 79.1%。

在任意两次回收开始期间完成更多的工作有助于将所需的回收次数从 2,551 减少到 1,402，45% 的减少。该应用将总的 GC 时间从 745ms 减少到 193ms，减少了 74%，并且每个版本的回收总时间从 14% 减少到 7%。当你在关闭垃圾回收你的情况下运行应用的优化版本时，性能差异仅为 13%，应用花费的时间从 2,753ms 缩短到 2,398ms。

### 总结

如果你花时间在专注减少分配，那么你就是在以一名 Go 开发者的身份，尽一切努力减少垃圾回收器的延迟。你不会编写零分配的应用，因此，认识有效分配（有利于应用的分配）和无效分配（不利于应用的分配）之间的区别很重要。然后，相信垃圾回收器会保持堆健康，并让你的应用一直运行。

使用垃圾回收器是一项不错的权衡。我将承担垃圾回收的成本，因此，我就没有内存管理的负担。Go 让开发人员可以提高工作效率，同时仍然编写足够快的应用。垃圾回收器是实现这一目标的重要组成部分。在下一篇文章中，我将分享另一个程序，该程序显示了回收器可以如何很好地分析你的 Go 应用，并找到最佳的回收路径。
