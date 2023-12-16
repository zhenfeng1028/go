| flag参数 | 有效值 |
| --------- | ----------- |
| 字符串 | 合法字符串 |
| 整数 | 1234、0664、0x1234等类型，也可以是负数。|
| 浮点数 | 合法浮点数 |
| bool类型 | 1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False。|
| 时间段 | 任何合法的时间段字符串。如"300ms"、"-1.5h"、"2h45m"。合法的单位有"ns"、"us"、"µs"、"ms"、"s"、"m"、"h"。|


两种常用的定义命令行flag参数的方法

(1) flag.Type()

基本格式如下：

	flag.Type(flag名, 默认值, 帮助信息) *Type
	
例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义：

```go
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")
```

(2) flag.TypeVar()

基本格式如下：

	flag.TypeVar(*Type, flag名, 默认值, 帮助信息)

例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以按如下方式定义：

```go
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")
```

通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。

支持的命令行参数格式有以下几种：

	-flag xxx （使用空格，一个-符号）
	--flag xxx （使用空格，两个-符号）
	-flag=xxx （使用等号，一个-符号）
	--flag=xxx （使用等号，两个-符号）

其中，布尔类型的参数必须使用等号的方式指定。