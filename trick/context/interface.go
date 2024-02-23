package context

import (
	"time"
)

// context接口中的4个方法
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

// （1）Deadline方法是获取设置的截止时间的意思，第一个返回值是截止时间，到了这个时间点，Context会自动发起取消请求；
// 第二个返回值ok==false时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。

// （2）Done方法返回一个只读的chan，类型为struct{}，我们在goroutine中，如果该方法返回的chan可以读取，
// 则意味着parent context已经发起了取消请求，我们通过Done方法收到这个信号后，就应该做清理操作，然后退出goroutine，释放资源。

// （3）Err方法返回取消的错误原因，因为什么Context被取消。

// （4）Value方法获取该Context上绑定的值，是一个键值对，所以要通过一个Key才可以获取对应的值，这个值一般是线程安全的。

// 以上四个方法中常用的就是Done了，如果Context取消的时候，我们就可以得到一个关闭的chan，
// 关闭的chan是可以读取的，所以只要可以读取的时候，就意味着收到Context取消的信号了，以下是这个方法的经典用法。

// func Stream(ctx context.Context, out chan<- Value) error {
// 	for {
// 		v, err := DoSomething(ctx)
// 		if err != nil {
// 			return err
// 		}
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		case out <- v:
// 		}
// 	}
// }

// Context接口并不需要我们实现，Go内置已经帮我们实现了2个，
// 我们代码中最开始都是以这两个内置的作为最顶层的parent context，衍生出更多的子Context。

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}

// 一个是TODO,它目前还不知道具体的使用场景，如果我们不知道该使用什么Context的时候，可以使用这个。
// 一个是Background，主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。

// 他们两个本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。

type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

// 这就是emptyCtx实现Context接口的方法，可以看到，这些方法什么都没做，返回的都是nil或者零值。

// 有了如上的根Context，那么是如何衍生更多的子Context的呢？这就要靠context包为我们提供的With系列的函数了。
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context

// 这四个With函数，接收的都有一个partent参数，就是父Context，我们要基于这个父Context创建出
// 子Context的意思，这种方式可以理解为子Context对父Context的继承，也可以理解为基于父Context的衍生。

// 通过这些函数，就创建了一颗Context树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个。

// （1）WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。

// （2）WithDeadline函数，和WithCancel差不多，它会多传递一个截止时间参数，意味着到了这个时间点，
// 会自动取消Context，当然我们也可以不等到这个时候，可以提前通过取消函数进行取消。

// （3）WithTimeout和WithDeadline基本上一样，这个表示是超时自动取消，是多少时间后自动取消Context的意思。

// （4）WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context，这个绑定的数据可以通过Context.Value方法访问到，后面我们会专门讲。

// 大家可能留意到，前三个函数都返回一个取消函数CancelFunc，这是一个函数类型，它的定义非常简单。
type CancelFunc func()

// 这就是取消函数的类型，该函数可以取消一个Context，以及这个节点Context下所有的子Context，不管有多少层级。
