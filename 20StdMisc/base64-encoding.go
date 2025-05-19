package main

//Go 提供了对base64 编码/解码的内置支持。
//base64 是可以解码的sha64 是没办法解码的。
import (
	b64 "encoding/base64" //类似c#中的using b64=encoding.base64 但是c#不会整体替换包，替换函数吧
	"fmt"
)

func main() {

	data := "abc123!?$*&()'-=@~"
	//Go 支持标准 Base64 和 URL 兼容 Base64。
	// 以下是使用标准编码器进行编码的方法。
	// 编码器需要一个，[]byte因此我们将 our 转换string为该类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	//解码可能会返回错误，如果您还不知道输入是否格式正确，可以检查该错误。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	//它使用与 URL 兼容的 base64 格式进行编码/解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

/**go run .\base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~*/
