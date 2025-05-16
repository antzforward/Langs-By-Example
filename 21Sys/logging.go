package main

//Go 标准库提供了用于从 Go 程序输出日志的简单工具，其中log包用于自由格式输出， log/slog包用于结构化输出。
import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {
	//Println只需从 包中调用类似 的函数log即可使用标准记录器，
	// 该记录器已预先配置，可合理地将日志输出到os.Stderr。
	// 其他方法（例如 Fatal*或 ）Panic*将在记录后退出程序。
	log.Println("standard logger")
	//可以使用标志配置记录器来设置其输出格式。
	// 默认情况下，标准记录器设置了log.Ldate和log.Ltime标志，
	// 这些标志存储在 中log.LstdFlags。例如，我们可以更改其标志，使其输出时间精度达到微秒级。
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")
	//它还支持发出log调用该函数的文件名和行。
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")
	//创建自定义记录器并传递它可能会很有用。创建新记录器时，我们可以设置前缀来区分其输出与其他记录器。
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")
	//我们可以使用该方法在现有的记录器（包括标准记录器）上设置前缀SetPrefix。
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")
	//记录器可以有自定义的输出目标；任何io.Writer工作都可以。
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	buflog.Println("hello")

	fmt.Print("from buflog:", buf.String())
	//该slog包提供 结构化的日志输出。例如，以 JSON 格式记录日志非常简单。
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	//除了消息之外，slog输出还可以包含任意数量的键=值对。
	myslog.Info("hello again", "key", "val", "age", 25)
}

/**go run .\logging.go
2025/05/15 18:52:36 standard logger
2025/05/15 18:52:36.367884 with micro
2025/05/15 18:52:36 logging.go:20: with file/line
my:2025/05/15 18:52:36 from mylog
ohmy:2025/05/15 18:52:36 from mylog
from buflog:buf:2025/05/15 18:52:36 hello
{"time":"2025-05-15T18:52:36.3678841+08:00","level":"INFO","msg":"hi there"}
{"time":"2025-05-15T18:52:36.3683932+08:00","level":"INFO","msg":"hello again","key":"val","age":25}*/
