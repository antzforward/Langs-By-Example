package main

import (
	"fmt"
	"time"
)

func main() {
	//lambda 写法真方便
	p := fmt.Println
	//当前时间
	now := time.Now()
	p(now)
	//您可以time通过提供年、月、日等来构建结构。时间始终与Location时区相关联。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	//您可以按预期提取时间值的各个组成部分。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	//周一至周日Weekday也可。
	p(then.Weekday())
	//这些方法比较两次，分别测试第一次发生在第二次之前、之后还是同时发生。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	//该Sub方法返回Duration表示两个时间之间的间隔。
	diff := now.Sub(then)
	p(diff)
	//我们可以用不同的单位来计算持续时间的长度。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	//您可以使用Add将时间提前指定的时长，或使用-将时间向后移动指定的时长。
	p(then.Add(diff))
	p(then.Add(-diff))
}

/**go run .\time.go
2025-05-06 10:14:02.4271891 +0800 CST m=+0.000000001
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
135557h39m3.775801863s
135557.65104883385
8.133459062930031e+06
4.8800754377580184e+08
488007543775801863
2025-05-06 02:14:02.4271891 +0000 UTC
1994-06-01 14:55:54.875585374 +0000 UTC*/
