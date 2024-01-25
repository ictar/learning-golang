# Gorm 包

一个处理关系数据库的 ORM 库，基于 `database/sql` 包。特性（都有测试）有：
- 支持：has-one / has-many / belongs-to / many-to-many / polymorphism（多态）/单表继承
- 钩子：before / after create / save / update / delete / find
- 预加载：`Preload` / `Joins`
- 事务 / 嵌套事务 / 保存点 / 回滚
- 上下文 / [Prepared Statement](https://gorm.io/zh_CN/docs/session.html#%E9%A2%84%E7%BC%96%E8%AF%91) 模式 / [DryRun](https://gorm.io/zh_CN/docs/session.html#DryRun) 模式

> 说明：PreparedStatement 的一个好处就是可以预防 sql 注入攻击。此外，由于 PreparedStatement 对象已预编译过，所以其执行速度要快于 Statement 对象。因此，多次执行的 SQL 语句经常创建为 PreparedStatement 对象，以提高效率。
通常批量处理时使用PreparedStatement。
>
> 说明：DryRun 模式指的是生成 SQL 但不执行，可用于准备或者测试生成的 SQL
>

- 批量插入 / [FindInBatches](https://gorm.io/zh_CN/docs/advanced_query.html#FindInBatches) / 通过Map查找以及创建 / 通过使用 SQL Expr 和 Context Valuer 进行 CRUD

> 说明：FindInBatches 用来进行批量查询和处理，在高效处理大型数据集、减少内存使用以及提高性能时特别有用
>

- SQL 构建器 / [Upsert](https://gorm.io/zh_CN/docs/create.html#upsert) / 锁 / 优化器提示（Optimizer hints） / [索引提示（Index Hints）](https://gorm.io/zh_CN/docs/advanced_query.html#Index-Hints) / 注释提示（Comment Hints）/ 命名参数 / [子查询（SubQuery）](https://gorm.io/zh_CN/docs/advanced_query.html#%E5%AD%90%E6%9F%A5%E8%AF%A2)

> 说明：upsert 是数据库插入操作的扩展，如果某个唯一字段已经存在，则将本次新增插入操作变成更新操作，否则就正常执行插入操作。
>
> 说明：优化器提示指的是建议数据库的查询优化器应该如何执行查询的指令。使用时需要导入 `gorm.io/hints`
>
> 说明：索引提示指导数据库要使用哪些索引。在查询规划器没有为查询选择最有效的索引的时候有用。使用时需要导入 `gorm.io/hints`
>
> 说明：注释提示在任意 SQL 关键字之前或之后添加 `/* */` 块注释。

- 复合主键 / 索引 / 约束
- [自动迁移](https://gorm.io/docs/migration.html#Auto-Migration)
- 日志
- 可扩展的灵活插件 API：数据库解析器（支持多种数据库、读写分离）/ Prometheus
- 开发者友好

## 教程
 
- [GORM 官方文档](https://gorm.io/docs/index.html)
- [`GORM` 包 API 文档](https://pkg.go.dev/gorm.io/gorm)

