# sycn.WatiGroup
> 在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：

|方法名|功能|
|----|----|
|(wg * WaitGroup) Add(delta int)|计数器+delta|
|(wg * WaitGroup) Done()|计数器-1|
|(wg * WaitGroup) Wait()|阻塞直到计数器变为0|

> sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。