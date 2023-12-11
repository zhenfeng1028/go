package main

import "net/http"

func do() error {
	res, err := http.Get("http://www.google.com")
	// 应该先检查请求是否成功
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// ..code...

	return nil
}

func main() {
	do()
}

// 因为在这里我们并没有检查我们的请求是否成功执行，
// 当它失败的时候，我们访问了 Body 中的空变量 res ，因此会抛出异常
