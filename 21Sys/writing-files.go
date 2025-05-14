package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dir, err := os.Getwd()
	check(err)

	path := filepath.Join(dir, "/tmp/dat1")
	os.MkdirAll(filepath.Dir(path), 0755)
	d1 := []byte("hello\ngo\n")
	err = os.WriteFile(path, d1, 0644)
	check(err)
	path = filepath.Join(dir, "/tmp/dat2")
	os.MkdirAll(filepath.Dir(path), 0755)
	f, err := os.Create(path)
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
