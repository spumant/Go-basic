package main

//import (
//	"fmt"
//	"io"
//	"os"
//	"time"
//)
//
//// 创建文件
//func createFile() {
//	f, err := os.Create("a.txt")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(f.Name())
//	}
//	f2, err := os.CreateTemp("", "temp")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(f2.Name())
//	}
//}
//
//// 创建目录
//func createDir() {
//	err := os.Mkdir("a", os.ModePerm)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	err2 := os.MkdirAll("a/b/c", os.ModePerm)
//	if err2 != nil {
//		fmt.Println(err2)
//	}
//}
//
//// 删除目录或者文件
//func remove() {
//	err := os.Remove("a.txt")
//	if err != nil {
//		fmt.Println(err)
//	}
//	err2 := os.RemoveAll("a")
//	if err2 != nil {
//		fmt.Println(err2)
//	}
//}
//
//// 工作目录
//func wd() {
//	dir, _ := os.Getwd()
//	fmt.Println(dir)
//	os.Chdir("d:/")
//	dir, _ = os.Getwd()
//	fmt.Println(dir)
//
//	s := os.TempDir()
//	fmt.Println(s)
//}
//
//// 重命名
//func rename() {
//	err := os.Rename("text.txt", "text2.txt")
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//// 读文件
//func readFile() {
//	file, err := os.ReadFile("text2.txt")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(string(file[:]))
//	}
//}
//
//// 写文件
//func writeFile() {
//	err := os.WriteFile("text2.txt", []byte("hello"), os.ModePerm)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//
//// 打开关闭文件
//func OpenCLoseFile() {
//	open, err := os.Open("a.txt")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(open.Name())
//		open.Close()
//	}
//
//	file, err2 := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 755)
//	if err2 != nil {
//		fmt.Println(err2)
//	} else {
//		fmt.Println(file.Name())
//		file.Close()
//	}
//}
//
//func readOps() {
//	f, _ := os.Open("a.txt")
//	for {
//		buf := make([]byte, 3)
//		n, err := f.Read(buf)
//		if err == io.EOF {
//			break
//		}
//		fmt.Println(n)
//		fmt.Println(string(buf))
//	}
//
//	buf2 := make([]byte, 4)
//	n, _ := f.ReadAt(buf2, 3)
//	fmt.Println(n)
//	fmt.Println(string(buf2))
//	f.Close()
//
//	de, _ := os.ReadDir("a")
//	for _, v := range de {
//		fmt.Println(v.IsDir())
//		fmt.Println(v.Name())
//	}
//
//	// 定位
//	f2, _ := os.Open("a.txt")
//	f.Seek(3, 0)
//	buf3 := make([]byte, 10)
//	n2, _ := f2.Read(buf3)
//	fmt.Println(n2)
//	fmt.Println(string(buf3))
//	f2.Close()
//}
//
//// 写字节数组
//func write() {
//	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0775)
//	f.Write([]byte("hello"))
//	f.Close()
//}
//
//// 写字符串
//func writeString() {
//	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0775)
//	f.WriteString("hello world")
//	f.Close()
//}
//
//// 随机写
//func writeAt() {
//	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0775)
//	f.WriteAt([]byte("aaa"), 3)
//	f.Close()
//}
//
//func main() {
//	// 获得当前正在运行的进程id
//	fmt.Println(os.Getpid())
//	// 父id
//	fmt.Println(os.Getppid())
//
//	// 设置新进程的属性
//	attr := &os.ProcAttr{
//		// Files指定新进程继承的活动文件对象
//		// 前三个分别为，标准输入、标准输出、标准错误输出
//		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
//		// 新进程的环境变量
//		Env: os.Environ(),
//	}
//	// 开始一个新进程
//	p, err := os.StartProcess("C:\\Windows\\System32\\notepad.exe", []string{"C:\\Windows\\System32\\notepad.exe", "D:\\a.txt"}, attr)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(p)
//	fmt.Println(p.Pid)
//
//	// 通过进程id查找进程
//	p2, _ := os.FindProcess(p.Pid)
//	fmt.Println(p2)
//
//	// 等待10秒，执行函数
//	time.AfterFunc(time.Second*10, func() {
//		// 向p进程发送退出信号
//		p.Signal(os.Kill)
//	})
//
//	// 等待进程p的退出，返回进程状态
//	ps, _ := p.Wait()
//	fmt.Println(ps.String())
//
//	// 获得所有环境变量
//	s := os.Environ()
//	fmt.Println(s)
//	// 获得某个环境变量
//	s2 := os.Getenv("GOPATH")
//	fmt.Println(s2)
//	// 设置环境变量
//	os.Setenv("env1", "env1")
//	//查找
//	s3, b := os.LookupEnv("env1")
//	fmt.Println(b)
//	fmt.Println(s3)
//	// 替换
//	os.Setenv("NAME", "gopher")
//	os.Setenv("BURROW", "/user/gopher")
//
//	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}"))
//
//	//清空
//	//os.Clearenv()
//}
