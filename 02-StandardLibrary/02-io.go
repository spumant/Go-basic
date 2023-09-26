package main

//import (
//	"bytes"
//	"fmt"
//	"io"
//	"log"
//	"os"
//	"strings"
//)
//
//func testCopy() {
//	r := strings.NewReader("hello world")
//	_, err := io.Copy(os.Stdout, r)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func testCopy2() {
//	r1 := strings.NewReader("first reader\n")
//	r2 := strings.NewReader("second reader\n")
//	buf := make([]byte, 8)
//	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
//		log.Fatal(err)
//	}
//
//	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
//		log.Fatal(err)
//	}
//}
//
//func testCopy3() {
//	r := strings.NewReader("Reader Stream\n")
//	lr := io.LimitReader(r, 4)
//
//	if _, err := io.Copy(os.Stdout, lr); err != nil {
//		log.Fatal(err)
//	}
//}
//
//func testCopy4() {
//	r1 := strings.NewReader("first reader ")
//	r2 := strings.NewReader("second reader ")
//	r3 := strings.NewReader("third reader \n")
//	r := io.MultiReader(r1, r2, r3)
//	if _, err := io.Copy(os.Stdout, r); err != nil {
//		log.Fatal(err)
//	}
//}
//
//func testCopy5() {
//	r := strings.NewReader("some io.Reader stream to be read\n")
//	var buf1, buf2 bytes.Buffer
//	w := io.MultiWriter(&buf1, &buf2)
//	if _, err := io.Copy(w, r); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(buf1.String())
//	fmt.Println(buf2.String())
//}
//
//func main() {
//
//}
