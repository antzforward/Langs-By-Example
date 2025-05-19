package main

//SHA256 哈希值常用于计算二进制或文本块的短身份。例如，TL​​S/SSL 证书使用 SHA256 计算证书签名。
import (
	"crypto/sha256" //Go 在各种 crypto/*包中实现了几个哈希函数。
	"fmt"
)

func main() {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s)) //Write需要字节。如果你有一个字符串s，请使用[]byte(s)将其强制转换为字节。
	//这将以字节切片的形式获取最终的哈希结果。参数Sum可用于附加到现有的字节切片：通常不需要。
	bs := h.Sum(nil)

	fmt.Println(s)
	//运行程序计算哈希值并以人类可读的十六进制格式打印它。
	fmt.Printf("%x\n", bs)
}

/**go run .\sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb*/
