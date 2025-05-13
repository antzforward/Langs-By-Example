package main

//Go 有几个用于处理 文件系统中目录的有用函数。
//在os下面有一些
import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//os库下的MKdir,文件夹下面就有FileMode的设置了
	//模式为0755对应的是？
	err := os.Mkdir("subdir", 0755)
	check(err)
	//删除整个整个目录，要小心啊，类似rm -rf
	defer os.RemoveAll("subdir")
	//创建一个空文件的辅助函数，os.WriteFile 后面也有FileMode=0644
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	//创建文件是进行写入操作
	createEmptyFile("subdir/file1")
	//创建一个文件夹层级，按照顺序创建，如果存在文件会有啥err吗？不会！
	//近似mkdir -p 我不知道这个命令~
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")
	//ReadDir列出目录内容，返回对象切片os.DirEntry。
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	//改变当前工作目录，类似cd 用Chdir
	err = os.Chdir("subdir/parent/child")
	check(err)
	//获得当前路径的内容，用.表示当前路径
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	//返回，当然记得之前的路径会更简单吧
	err = os.Chdir("../../..")
	check(err)
	//我们还可以递归访问目录，包括其所有子目录。WalkDir接受回调函数来处理访问的每个文件或目录。
	//跟python下面的walk 差不多的写法，自己写func传入也行
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
