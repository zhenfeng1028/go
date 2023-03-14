Scan从标准输入扫描文本，读取由空白符分隔的值传递给本函数的参数中，换行符视为空白符。
本函数返回成功扫描的数据个数和遇到的任何错误。如果读取的数据个数比提供的参数少，会返回一个错误报告原因。
函数签名

    func Scan(a ...interface{}) (n int, err error)

Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值传递给本函数的参数中。
本函数返回成功扫描的数据个数和遇到的任何错误。
函数签名

    func Scanf(format string, a ...interface{}) (n int, err error)

Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
本函数返回成功扫描的数据个数和遇到的任何错误。
函数签名

    func Scanln(a ...interface{}) (n int, err error)