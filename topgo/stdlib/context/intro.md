# Context
## 源头
在 Go http 包的 Server 中，每一个请求在都有一个对应的 goroutine 去处理。请求处理函数通常会启动额外的 goroutine 用来访问后端服务，比如数据库和 RPC 服务。用来处理一个请求的 goroutine 通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息、验证相关的 token、请求的截止时间。 当一个请求被取消或超时时，所有用来处理该请求的 goroutine 都应该迅速退出，然后系统才能释放这些 goroutine 占用的资源。

## why
1. 如何优雅地结束子 goroutine ？
2. 全局变量的方式
3. 管道的方式
4. 官方的方案（context）
5. 当子 goroutine 又开启另外一个 goroutine 时，只需要将 ctx 传入即可

## Context 初识
Go1.7 加入了一个新的标准库 context，它定义了 Context 类型，专门用来简化对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用 WithCancel、WithDeadline、WithTimeout 或 WithValue 创建的派生上下文。当一个上下文被取消时，它派生的所有上下文也被取消。

## Context 接口
context.Context 是一个接口，该接口定义了四个需要实现的方法。具体签名如下：


    type Context interface {
        Deadline() (deadline time.Time, ok bool)
        Done() <-chan struct{}
        Err() error
        Value(key interface{}) interface{}
    }

其中：
* Deadline 方法需要返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）；
* Done 方法需要返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消之后关闭，多次调用 Done 方法会返回同一个 Channel；
* Err 方法会返回当前 Context 结束的原因，它只会在 Done 返回的 Channel 被关闭时才会返回非空的值；
    * 如果当前 Context 被取消就会返回 Canceled 错误；
    * 如果当前 Context 超时就会返回 DeadlineExceeded 错误；
* Value 方法会从 Context 中返回键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法仅用于传递跨 API 和进程间跟请求域的数据；

## Background() 和 TODO()
Go 内置两个函数：Background() 和 TODO()，这两个函数分别返回一个实现了 Context 接口的 background 和 todo。我们代码中最开始都是以这两个内置的上下文对象作为最顶层的 partent context，衍生出更多的子上下文对象。

Background() 主要用于 main 函数、初始化以及测试代码中，作为 Context 这个树结构的最顶层的 Context，也就是根 Context。

TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么 Context 的时候，可以使用这个。

background 和 todo 本质上都是 emptyCtx 结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的 Context。

## With 系列函数

    func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
    func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
    func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
    func WithValue(parent Context, key, val interface{}) Context

## 使用 Context 的注意事项
* 推荐以参数的方式显示传递 Context
* 以 Context 作为参数的函数方法，应该把 Context 作为第一个参数
* 给一个函数方法传递 Context 的时候，不要传递 nil，如果不知道传递什么，就使用 context.TODO()
* Context 的 Value 相关方法应该传递请求域的必要数据，不应该用于传递可选参数
* Context 是线程安全的，可以放心的在多个 goroutine 中传递