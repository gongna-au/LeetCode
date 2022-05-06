package main

import "fmt"
import "os"

type LibStruct struct {
}

//回调函数1
func mainLibEntry() {

}

func mainReaderEntry(x int) {

}

func mainReturnBook(x int) {

}
func mainExit() {
	fmt.Println("退出")
	os.Exit(3)
}

//回调函数2
func libAddReader(x int) {

}

func libAddBook() {

}
func libQueryReader() {

}
func libPutBook() {

}
func libShowAll() {

}
func libExit() {

}

//reader
func readerAddBook() {

}
func readerCheckOut() {

}
func readerQueryBook() {

}
func readerListBook() {

}
func readerCart() {

}
func readerExit() {

}

func main() {
	mainExit()

}
