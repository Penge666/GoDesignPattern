package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type File struct {
	filename string
}

func (f *File) Read() string {
	content, err := ioutil.ReadFile(f.filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(content)
}

type FileProxy struct {
	file *File
}

func (p *FileProxy) Read() string {
	// 在这里做访问控制或日志记录等操作
	if p.file == nil {
		p.file = &File{filename: "sample.txt"}
	}

	// 记录日志
	fmt.Println("Logging: Accessing file:", p.file.filename)

	// 进行实际的文件读取
	content := p.file.Read()

	// 读取后做一些操作
	fmt.Println("Logging: File read operation completed")

	return content
}

func main() {
	// 创建文件代理
	fileProxy := &FileProxy{}

	// 模拟文件读取
	content := fileProxy.Read()
	fmt.Println("File Content:", content)

	// 创建一个新文件用于读取
	_ = ioutil.WriteFile("sample.txt", []byte("Hello, this is a sample text file!"), os.ModePerm)
}
