package main

//import (
//	"bufio"
//	"bytes"
//	"fmt"
//	"io"
//	"os"
//	"strings"
//)
//
//func test() {
//	file, _ := os.Open("a.txt")
//	defer file.Close()
//	r2 := bufio.NewReader(file)
//	s, _ := r2.ReadString('\n')
//	fmt.Println(s)
//}
//
//func test2() {
//	s := strings.NewReader("ABCDEFGHIGKLMNOPQRSTUVWXYZ1234567890")
//	br := bufio.NewReader(s)
//	p := make([]byte, 10)
//	for {
//		n, err := br.Read(p)
//		if err == io.EOF {
//			break
//		} else {
//			fmt.Println(string(p[0:n]))
//		}
//	}
//}
//
//func test3() {
//	s := strings.NewReader("ABCDEFG")
//	br := bufio.NewReader(s)
//
//	c, _ := br.ReadByte()
//	fmt.Println(c)
//
//	c, _ = br.ReadByte()
//	fmt.Println(c)
//
//	br.UnreadByte()
//	c, _ = br.ReadByte()
//	fmt.Println(c)
//
//}
//
//func test4() {
//	s := strings.NewReader("ABC,DEF,GHI,JKL")
//	br := bufio.NewReader(s)
//
//	w, _ := br.ReadSlice(',')
//	fmt.Println(w)
//
//	w, _ = br.ReadSlice(',')
//	fmt.Println(w)
//
//	w, _ = br.ReadSlice(',')
//	fmt.Println(w)
//}
//
//func test5() {
//	s := strings.NewReader("ABCDEFGHIKLM")
//	br := bufio.NewReader(s)
//	//b := bytes.NewBuffer(make([]byte, 0))
//	file, _ := os.OpenFile("a.txt", os.O_RDWR, 0777)
//	defer file.Close()
//	br.WriteTo(file)
//	//fmt.Println(b)
//}
//
//func test6() {
//	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777)
//	writer := bufio.NewWriter(f)
//	writer.WriteString("hello world")
//	writer.Flush()
//}
//
//func test7() {
//	b := bytes.NewBuffer(make([]byte, 0))
//	bw := bufio.NewWriter(b)
//	s := strings.NewReader("123")
//	br := bufio.NewReader(s)
//	rw := bufio.NewReadWriter(br, bw)
//	p, _ := rw.ReadString('\n')
//	fmt.Println(string(p))
//	rw.WriteString("asdf")
//	rw.Flush()
//	fmt.Println(b)
//}
//
//func test8() {
//	s:=strings.NewReader("ABC DEF GHI JKL")
//	bs:= bufio.NewScanner(s)
//	bs.Split(bufio.ScanWords) // 以空格作为分隔符
//	for bs.Scan() {
//		fmt.Println(bs.Text())
//	}
//}
//
//func main() {
//
//}
