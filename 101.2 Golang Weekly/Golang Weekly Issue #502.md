原文：[Golang Weekly Issue #502](https://golangweekly.com/issues/502)

---

🐣 简单说明一下，下周二 Go Weekly 会休息（作为复活节休息的一部分）。我们会在 4 月 16 日（星期二）回来 :-)
 
 _你的编辑，Peter Cooper_  
  

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/yctsfi9fawh7miujrzxr.jpg)](https://golangweekly.com/link/153209/web)  

[Freeze：生成代码图像和终端输出](https://golangweekly.com/link/153209/web "github.com") —— 由 Go 驱动的优秀的库和 CLI 工具（如[Bubble Tea](https://golangweekly.com/link/153210/web)、 [Huh](https://golangweekly.com/link/153211/web) 和 [Wish](https://golangweekly.com/link/153212/web)）的供应商推出了 _Freeze_，这是一种创建代码图像，以及/或者以 PNG/SVG/WebP 格式进行终端输出。

_Charm_ 


> 🤩 Charm 的那些家伙还通过[“他们是如何做到的”](https://golangweekly.com/link/153213/web)，来庆祝他们获得的第 100,000 个 GitHub star。一个真实的 Go 成功的故事！
  

[用 Go 构建一个交互式 Shell](https://golangweekly.com/link/153214/web "www.dolthub.com") —— 来自 [_Dolt_ 数据库](https://golangweekly.com/link/153215/web)（包含了一个 CLI，用于获取 SQL 查询并将其提供给数据库）的工作人员提供了创建 CLI 的一个直观的指南，这个 CLI 包括命令历史记录和自动完成等功能。[ishell](https://golangweekly.com/link/153216/web) 是他们选中的库。

_Zach Musgrave (DoltHub)_ 

  
[![](https://copm.s3.amazonaws.com/1ea6f5f1.png)](https://golangweekly.com/link/153208/web) 

[走！通过 Ardan Labs Consulting 释放您的技术潜力](https://golangweekly.com/link/153208/web "www.ardanlabs.com") —— 因技能差距、开发速度或复杂的技术挑战而苦苦挣扎？ Ardan Labs 专注于 Go、Rust、Docker 和 K8s，以加速您的软件开发、优化架构和管理技术债务。让我们为您的团队加油！


_Ardan Labs Consulting sponsor_


▶ [讨论 Go 中的调试](https://golangweekly.com/link/153217/web "changelog.com") —— 在最新一集的 _Go Time_ 中，Matt Boyle、Bill Kennedy 和 Jon Calhoun 讨论了调试技术。Bill 解释了为什么他不喜欢他的手下使用调试器，以及他更喜欢使用那些可用于生产的技术。

_Go Time Podcast podcast_


_快速了解：_

  * 🎂 12 年前，[Go 1.0 发布啦](https://golangweekly.com/link/153218/web)。

  * 🤔 谷歌的工程总监 Lars Bergstrom 表示，[“谷歌的 Rust 团队和 Go 团队的生产力一样高”](https://golangweekly.com/link/153219/web)，尽管二者的生产力都是 C++ 团队的两倍。

  * 🐞 JetBrains 的 Ruslan Akhmetzianov 展示了[GoLand 最近对数据流分析的支持](https://golangweekly.com/link/153220/web)（DFA），以及如何检测运行中的错误。

  * 🤖 [Gopher-Verse](https://golangweekly.com/link/153221/web) 是 _Creative Fabrica_ 的一个有趣的项目，它使用经过训练的模型，根据您的提示生成式创建 Go 风格的打地鼠。结果好坏参半，但很有趣。


[暴力文本搜索优化](https://golangweekly.com/link/153222/web "boyter.org") —— 如果内存中有大量文本数据，那么现在对其执行暴力搜索可能会快得惊人，特别是在进行一些优化后……

_Ben E. C. Boyter_ 
  

[为服务器优化 SQLite](https://golangweekly.com/link/153223/web "kerkour.com") ——  开发人员越来越意识到 SQLite 是可以进行 _长_ 距离扩展的，并且也可以用于许多通常使用更复杂系统的场景。Sylvain 深入研究了如何充分利用它。

_Sylvain Kerkour_ 
  

[用 Go 构建一个博客：将 Markdown 渲染为 HTML](https://golangweekly.com/link/153224/web "www.calhoun.io") —— 一个[正在进行的系列](https://golangweekly.com/link/153225/web)，介绍了如何使用 Go 构建一个简单的博客系统。

_Jon Calhoun_ 

> 💡 Jon 还发售了他内容丰富的视频课程[Test with Go](https://golangweekly.com/link/153226/web) 和 [Web Development with Go](https://golangweekly.com/link/153227/web)，截止至 4 月 5 日。  


▶ [指针性能？](https://golangweekly.com/link/153228/web "www.youtube.com") —— 从函数返回指针而不是值会提高 Go 的性能吗？_（3分钟。）_

_Bill Moran_ 


[防止敏感信息通过日志泄漏](https://golangweekly.com/link/153229/web)   
_Willem Schots_  


## 🛠 代码和工具  

[LangChain Go：用于 LLM 应用的一个 LangChain Go Port/Fork](https://golangweekly.com/link/153230/web "tmc.github.io") —— [LangChain](https://golangweekly.com/link/153231/web) 是一个流行的框架（最常与 Python 相关），用于开发语言模型驱动的应用程序。 LangChain Go 用 Go 重新实现了相同的概念。[GitHub 存储库](https://golangweekly.com/link/153232/web)。

_Travis Cline et al._ 


[Excelize：用于处理 Excel 电子表格的库](https://golangweekly.com/link/153233/web "xuri.me") —— 读写 XLAM / XLSM / XLSX / XLTM / 和 XLTX 文件。一个历史悠久的库，它正在不断发展壮大。[GitHub 存储库](https://golangweekly.com/link/153234/web)。

_QI-ANXIN GROUP_ 


[Hookdeck：Amazon EventBridge 替代方案](https://golangweekly.com/link/153235/web "hookdeck.com") —— 通过工程团队的事件网关，在 EDA 中接收、转换、过滤、路由和发布消息。

_Hookdeck sponsor_


[Ebitengine 2.7 发布：Go 的 2D 游戏引擎](https://golangweekly.com/link/153236/web "ebitengine.org") —— 这次的新功能是 `text/v2`，这是一个全新的文本渲染包，它提供了对阿拉伯语和垂直渲染日语的改进支持，[字形矢量化（glyph vectorization）](https://golangweekly.com/link/153237/web)、更多的文本对齐选项、以及改进的 OpenType 支持。

_Hajime Hoshi_ 


[Beego 2.2：RESTful APIs 和 Webapps 的后端框架](https://golangweekly.com/link/153238/web "github.com") —— 受到 Tornado、Sinatra 和 Flask 的启发。[v2.2.0](https://golangweekly.com/link/153239/web) 将 Go 版本提升至 1.20。


_beego Framework_ 


[sqlc 1.26](https://golangweekly.com/link/153240/web) - 根据 SQL 生成类型安全的 Go 代码。_（包括一个重要的安全修复，这个修复针对使用输出插件的用户。）_

[GoCV 0.36.0](https://golangweekly.com/link/153241/web) - OpenCV 4 的 Go 绑定现在支持 OpenCV 4.9.0 和 Go 1.22 啦。

[go-github 61.0](https://golangweekly.com/link/153242/web) - GitHub v3 API 的客户端库。

♪ [Oto 3.2](https://golangweekly.com/link/153243/web) - 用于播放声音的低级跨平台库。

[eBPF 0.14](https://golangweekly.com/link/153244/web) - 用于 eBPF 程序的 Pure-Go 库。

[go-sqlite 1.2](https://golangweekly.com/link/153245/web) - SQLite 3 的低级接口。

[fx 34.0](https://golangweekly.com/link/153246/web) - 强大的终端 JSON 查看器。

[Bun 1.2](https://golangweekly.com/link/153247/web) - 用于 Go 的 SQL 优先 ORM。


[Micro 4.4](https://golangweekly.com/link/153248/web)


## 😎 你知道……吗？ 

尽管我们非常喜欢制作这份时事通讯，但我们还发布了一些您可能不知道的其他内容。您可以直接在他们的主页上查看他们的最新发布：

💛 [JavaScript Weekly](https://golangweekly.com/link/153249/web) 涵盖了 JavaScript 的所有内容，以及更广泛的生态系统，包括 TypeScript、WebAssembly、Astro、HTMX、Vue.js、构建工具等。

⚛️ 不出所料，[React Status](https://golangweekly.com/link/153250/web) 涵盖了 React 世界，而 [Node Weekly](https://golangweekly.com/link/153251/web) 则深入探讨了 Node.js、npm 生态系统以及其他服务器端 JS 平台（如 Deno 和 Bun）。

👩‍💻 [Frontend Focus](https://golangweekly.com/link/153252/web) 涵盖了浏览器中出现的那些广泛的技术：CSS、HTML、可访问性、WebGL、Web API – 尽在其中。

🐘 无需对[Postgres Weekly](https://golangweekly.com/link/153253/web) 和 [Ruby Weekly](https://golangweekly.com/link/153254/web)进行过多介绍，它们正是您所期望的的周刊 ;-) 



我们将在两周后回来 - 到时见！
 
 _Peter Cooper，你的编辑_