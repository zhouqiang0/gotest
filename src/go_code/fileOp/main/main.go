package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//O_RDONLY int = syscall.O_RDONLY // open the file read-only.
//O_WRONLY int = syscall.O_WRONLY // open the file write-only.
//O_RDWR   int = syscall.O_RDWR   // open the file read-write.
//// The remaining values may be or'ed in to control behavior.
//O_APPEND int = syscall.O_APPEND // 追加append data to the file when writing.
//O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
//O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
//O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
//O_TRUNC  int = syscall.O_TRUNC  // 覆盖truncate regular writable file when opened.

func CopyFile(destFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open err:%v", err)
	}
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	destFile, err := os.OpenFile(destFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("writer err:%v", err)
	}
	defer destFile.Close()

	writer := bufio.NewWriter(destFile)

	return io.Copy(writer, reader)
}

func main() {

	//使用带缓冲的读文件bufio
	//file, err := os.Open("C:/1test/test.txt")
	//if err != nil{
	//	fmt.Println("open file err=",err)
	//}
	//defer file.Close()
	//reader := bufio.NewReader(file)
	//for {
	//	str, err := reader.ReadString('\n')
	//	if err == io.EOF{
	//		break
	//	}
	//	fmt.Print(str)
	//}
	//fmt.Println("读取结束")

	//使用ioutil.ReadFile读取，不需要open,close。适合小文件
	//file := "C:/1test/test.txt"
	//content, err := ioutil.ReadFile(file)
	//if err != nil{
	//	fmt.Printf("err:%v", err)
	//}
	//fmt.Printf("%v", content) //[]Byte

	//filePath := "C:/1test/test1.txt"
	////file, err := os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE, 0666)
	////file, err := os.OpenFile(filePath,os.O_WRONLY | os.O_APPEND, 0666)
	//file, err := os.OpenFile(filePath,os.O_WRONLY | os.O_TRUNC, 0666)
	//if err != nil{
	//	fmt.Printf("err:%v\n",err)
	//}
	//defer file.Close()
	////写入内容 使用带缓存的writer
	//str := "new line覆盖\n"
	//writer := bufio.NewWriter(file)
	//for i := 0; i < 5 ; i++  {
	//	writer.WriteString(str)
	//}
	////调用flush方法将缓存中的数据写入
	//writer.Flush()

	//拷贝文件
	//file1 := "C:/1test/test1.txt"
	//file2 := "C:/1test/test.txt"
	//data, err := ioutil.ReadFile(file1)
	//if err != nil{
	//	fmt.Printf("读err:%v", err)
	//}
	//err = ioutil.WriteFile(file2, data, 0666)
	//if err != nil{
	//	fmt.Printf("写err:%v", err)
	//}

	//调用copy方法
	srcFile := "C:/1test/saber.jpg"
	destFile := "C:/1test/fate/saber.jpg"
	CopyFile(destFile, srcFile)

}
