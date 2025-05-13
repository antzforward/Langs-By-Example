package main

//Go 支持通过基于模式的布局进行时间格式化和解析。
import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	//这是根据 RFC3339 使用相应的布局常量格式化时间的基本示例。
	//表现为2025-05-06T10:26:33+08:00
	p(t.Format(time.RFC3339))
	//细微的地方不了解内容
	//格式化2012-11-01 22:08:41 +0000 +0000
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	//用举例的方式来提供格式，输出10:26AM
	p(t.Format("3:04PM"))
	//同上，好用：Tue May  6 10:26:33 2025
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	//2025-05-06T10:26:33.97395+08:00
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	//0000-01-01 20:41:00 +0000 UTC
	p(t2)
	//2025-05-06T10:26:33-00:00
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	//parse 的时候value不匹配layout的时候就会有error 不捕捉panic了
	_, e = time.Parse(ansic, "8:41PM")
	//parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
	p(e)
}

/**go run .\time-formatting-parsing.go
2025-05-06T10:26:33+08:00
2012-11-01 22:08:41 +0000 +0000
10:26AM
Tue May  6 10:26:33 2025
2025-05-06T10:26:33.97395+08:00
0000-01-01 20:41:00 +0000 UTC
2025-05-06T10:26:33-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"*/
