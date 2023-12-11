package main

func test(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()

	defer println("c")
}

func main() {
	test(0)
}

// 多个 defer 注册，按 FILO 次序执行（先进后出）。
// 哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。
