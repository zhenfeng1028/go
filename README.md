一个用来系统练习 Go 语言的示例仓库，涵盖基础语法、标准库、数据结构与算法、Web 开发、并发、ORM、消息队列等多个主题。

## 仓库结构概览

- `tutorial/`：Go 入门与进阶教程代码示例，按来源分类：

  - `google/`：来自 go.dev *A Tour of Go* / 官方教程的练习与示例
  - `runoob/`：菜鸟教程相关示例
  - `topgoer/`：TopGoer 教程示例
- `pkg.go.dev/`：按标准库包名组织的示例代码，模仿 pkg.go.dev 的结构，演示常用标准库的用法，例如：

  - `bytes/`、`strings/`、`strconv/`
  - `net/http/`、`net/url/`
  - `encoding/json/`、`encoding/base64/`
- `dsa/`：数据结构与算法实现：

  - `other/`：LRU、堆排序等
  - `programiz.com/`：按专题分类的各类算法实现（图算法、动态规划、树结构、排序与查找等）
- `trick/`：各种 Go 实战技巧与小 demo，例如：

  - `context/`：上下文传递、超时控制、取消等
  - `crawler/`、`goquery/`：简单爬虫与解析示例
  - `cron/`：定时任务
  - `marshal/` / `omitempty/`：序列化相关用法
  - `http/`、`websocket/`、`sse/` 等网络编程示例
  - `mutex/`、`waitgroup/`、`spinlock/`：并发同步原语示例
  - `unsafe/`：`unsafe` 包相关示例
- `nosql/`：常用 NoSQL 的使用示例：

  - `redis/`（redigo）
  - `memcached/`
  - `mongodb/`
- `orm/`：Go ORM 框架使用示例：

  - `gorm/`
  - `xorm/`（包含 MySQL、Postgres、SQLite 以及结构体生成脚本等）
- `mq/`：消息队列相关示例：

  - `kafka/`：生产者 / 消费者 demo
  - `zookeeper/`：基础操作与客户端示例
- `gin/`：使用 Gin 框架实现的简单 HTTP 接口示例（如登录、专辑等）。
- `cgo/`：调用 C 代码的 cgo 示例，包括 C 源文件、头文件与静态库。
- `form/`：表单相关的 demo（如简单购物车、wiki 示例）。
- `tool/`：性能分析与追踪工具示例：

  - `pprof/`
  - `trace/`
- `test/`、`go_test/`、`quizzes/`：各种测试、习题和语言细节的小实验。
- `util/`：常用工具函数合集，例如：

  - Excel 生成、字符串处理、JWT/HMAC/RSA/PGP 加解密
  - worker pool、缓存、邮件发送、SFTP/SMTP、压缩解压等

## 使用方式

本仓库以示例和练习为主，没有单一入口，可按兴趣或学习主题进入对应目录阅读和运行代码。
