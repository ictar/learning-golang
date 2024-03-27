原文：[Golang Weekly Issue #501](https://golangweekly.com/issues/501)

---

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/eijjb1wcqjtrhkxkmhbd.jpg)](https://golangweekly.com/link/152968/web)  

[在 2024 年学习 Go：从小白到高手](https://golangweekly.com/link/152968/web "www.bytesizego.com") —— Go 的学习资源非常丰富，涵盖从文本、视频到播客的所有媒体，以及从初学者到专家的所有技能级别。在这里，Matt 汇集了六年的书签（以及 YouTube 历史），为每个人指出一些有用的东西。

_Matt Boyle_ 


[切片的视觉指南](https://golangweekly.com/link/152969/web "sazak.io") —— 一篇包含丰富代码和图表的文章，介绍了切片的底层工作原理，涵盖切片创建、使用 `make` 和 `append` 进行操作、容量的增长方式以及从现有切片创建新切片的切片语法。

_Ozan Sazak_ 


[![](https://copm.s3.amazonaws.com/1ea6f5f1.png)](https://golangweekly.com/link/152967/web) 

[走！通过 Ardan Labs Consulting 释放您的技术潜力 ](https://golangweekly.com/link/152967/web "www.ardanlabs.com") —— 因技能差距、开发速度或复杂的技术挑战而苦苦挣扎？ Ardan Labs 专注于 Go、Rust、Docker 和 K8s，以加速您的软件开发、优化架构和管理技术债务。让我们为您的团队提供强大的支持！

_Ardan Labs Consulting sponsor_

  

[Goroutine 泄漏的一个案例](https://golangweekly.com/link/152970/web "brainbaking.com") —— 在发现 _“永无休止的 Goroutine 工厂”_ 导致应用程序出现问题后，Wouter 分享了他的发现，以便您可以避免类似的命运。

_Wouter Groeneveld_


_快速了解：_
  * [Google 推出了三个新的“安全的” Go 库](https://golangweekly.com/link/152971/web) - SafeText、SafeOpen 和 SafeArchive - 以防止使用 YAML、打开文件和处理存档文件时出现的常见安全问题。

  * 📅 [今年第二届 GopherCon Europe](https://golangweekly.com/link/152972/web) 将于今年 6 月在德国柏林举行。[完整的时间表](https://golangweekly.com/link/152973/web)现已发布。

  * [The IBM Open Enterprise SDK for Go 1.22](https://golangweekly.com/link/152974/web) 现已推出。是时候更新啦。


[\'Go 的枚举依然很恶心（Go Enums Still Suck）\'](https://golangweekly.com/link/152975/web "www.zarl.dev") —— 这篇文章的标题 _或_ 语气可能不适合你，但作者反思了他的[第一篇关于 Go“枚举”的文章](https://golangweekly.com/link/152976/web)的反馈，并认为讨论富有成效。

_Steve McCutcheon_ 


[通过 Go 应用控制 ZigBee LED 灯](https://golangweekly.com/link/152977/web "tderflinger.com") —— 公平地说，Go 应用程序主要发送 MQTT 消息并提供 HTML 页面，但这对您来说可能是一个快速、有趣的周末项目。

_Thomas Derflinger_ 


[针对 Map 和切片的泛型改进](https://golangweekly.com/link/152978/web "henry.precheur.org") —— _“Go 的日子很好，而且正在变得更好。”_

_Henry Precheur_ 


[“极速” Go Docker 构建的一种方法](https://golangweekly.com/link/152979/web)   
_Abhinav Sonkar_  
  


🛠 代码和工具  
   
[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/xlunhx4qypl7xtsq62sp.jpg)](https://golangweekly.com/link/152980/web)  


[asciigraph 0.6：ASCII 线图渲染库](https://golangweekly.com/link/152980/web "github.com") —— 一个长期存在的库，用于以 ASCII 文本渲染基本线图（例如终端的理想选择），所有这些都没有使用任何依赖。 v0.6 支持对彩色图表添加图例。

_Rohit Gupta_ 


[Charcoal：更快的 `utf8.Valid`，使用 _不带_ SIMD 的多字节处理](https://golangweekly.com/link/152981/web "github.com") —— 提供了与[标准库版本](https://golangweekly.com/link/152982/web)相同的 API，但在检查文本是否由有效的 UTF-8 符文组成时的速度提高了 50%。 

_Sugawara Yuuta_ 


[使用 OAuth 保护 Golang 应用](https://golangweekly.com/link/152983/web "fusionauth.io") —— 本教程将向您展示如何使用 OAuth 对 Golang 应用程序中的用户进行身份验证。

_FusionAuth sponsor_


[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/zamx5n18olmkduhgmyxz.jpg)](https://golangweekly.com/link/152984/web)  

[Log 0.4：色彩斑斓的 Go 日志库](https://golangweekly.com/link/152984/web "github.com") —— A library from the same folks who brought us 给我们带来 [Bubble Tea](https://golangweekly.com/link/152985/web) 和 [Gum](https://golangweekly.com/link/152986/web) 的那个人给我们又带来了这个库，这个库提供 _“可定制的人类可读彩色日志记录，包括开箱即用部分。”_ [v0.4](https://golangweekly.com/link/152987/web) 现在允许您根据自己的喜好设计自定义日志级别。

_Charm_ 


[go-toml 2.2：用于 TOML 格式的 Go 库](https://golangweekly.com/link/152988/web "github.com") —— [TOML](https://golangweekly.com/link/152989/web) 是一种配置文件格式，由 GitHub 的创始人之一（**Tom** Preston-Werner，碰巧）发明。[它看起来是这样的。](https://golangweekly.com/link/152990/web) _（这个 Go 库也是由 Tom/Thomas 编写的，这真是太棒了！）_

_Thomas Pelletier_ 


[gocron 2.2.7](https://golangweekly.com/link/152991/web) - 以预定的时间间隔运行 Go 函数。

[Katana 1.1](https://golangweekly.com/link/152992/web) - 下一代爬虫框架。

[Sessions 1.0](https://golangweekly.com/link/152993/web) - 用于会话管理的 Gin 中间件。

[Ebitengine 2.6.7](https://golangweekly.com/link/152994/web) - Go 的简单 2D 游戏引擎。

[NATS.go 1.34](https://golangweekly.com/link/152995/web) - NATS 的 Go 客户端。


---
📰 分类广告


[加入 Stickermule 的“强悍”团队！](https://golangweekly.com/link/152996/web) 我们的软件团队遍布 17 个国家/地区，而我们正在寻找更多优秀的工程师[加入我们的团队](https://golangweekly.com/link/152996/web)。


🪝 使用 [Hookdeck 的](https://golangweekly.com/link/152997/web) 开发人员工具和基础设施来构建事件驱动的应用程序。首先[免费试用 3 个月的 Hookdeck Team 计划](https://golangweekly.com/link/152997/web)。

---

🕹 还有一个好玩的……

[![](https://res.cloudinary.com/cpress/image/upload/w_1280,e_sharpen:60,q_auto/epybkk7ahetdbxce70qr.jpg)](https://golangweekly.com/link/152999/web)  

[Ebitengine 之上的一个基本的 3D 光线投射引擎](https://golangweekly.com/link/152999/web "github.com") —— [Ebitengine](https://golangweekly.com/link/153001/web) 是一个 _2D_ 游戏引擎，但正如 Wolfenstein 3D 和 Doom 等游戏在 90 年代所证明的那样，光线投射提供了一种（相对）简单的方法来使用简单的 2D 图元渲染 3D 环境。[▶️ 示例视频。](https://golangweekly.com/link/153003/web)

_Eric H_ 