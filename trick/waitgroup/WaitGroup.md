正常情况下，新激活的goroutine的结束过程是不可控制的，唯一可以保证终止goroutine的行为是main goroutine的终止。也就是说，我们并不知道哪个goroutine什么时候结束。

但很多情况下，我们正需要知道goroutine是否完成。这需要借助sync包的WaitGroup来实现。

WatiGroup是sync包中的一个struct类型，用来收集需要等待执行完成的goroutine。下面是它的定义：

type WaitGroup struct {
	// Has unexported fields.
}

func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()

它有3个方法：

Add()：每次激活想要被等待完成的goroutine之前，先调用Add()，用来设置或添加要等待完成的goroutine数量
例如Add(2)或者两次调用Add(1)都会设置等待计数器的值为2，表示要等待2个goroutine完成
Done()：每次需要等待的goroutine在真正完成之前，应该调用该方法来人为表示goroutine完成了，该方法会对等待计数器减1
Wait()：在等待计数器减为0之前，Wait()会一直阻塞当前的goroutine

也就是说，Add()用来增加要等待的goroutine的数量，Done()用来表示goroutine已经完成了，减少一次计数器，Wait()用来等待所有需要等待的goroutine完成。