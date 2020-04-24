package blog_go_interface_basic

import (
	gconv "github.com/og/x/conv"
	ge "github.com/og/x/error"
	"log"
	"os"
	"testing"
)
func TestInterface (t *testing.T){
	// 直接返回数据
	{
		// []int{0,1,2,3,4,5,6,7,8,9}
		log.Print("createNumbers(10)", createNumbers(10))
	}
	// 数据量大时会需要同分段返回
	{
		numbers :=  []int{}
		createNumbersUseCallbackFunc(23, func(data []int) {
			numbers = append(numbers, data...)
		})
		// []int{0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22}
		log.Print("createNumbersUseCallbackFunc", numbers)
	}
	// 定义结构体接收
	{
		m := MemoryNumbers{}
		createNumbersUseStruct(15, &m)
		// MemoryNumbers{ value: {0 1 2 3 4 5 6 7 8 9 10 11 12 13 14} }
		log.Print("createNumbersUseStruct", m)
	}
	// 定义接口接收(只要满足接口的结构体都能接收)
	{
		fileNumbers := FileNumbers{"./file_go_interface_basic.txt"}
		log.Print("createNumbersUseInterface " + fileNumbers.Name + " 写入")
		createNumbersUseInterface(12, &fileNumbers)
		memoryNumbers := MemoryNumbers{}
		createNumbersUseInterface(15, &memoryNumbers)
		// MemoryNumbers{ value: {0 1 2 3 4 5 6 7 8 9 10 11 12 13 14} }
		log.Print("createNumbersUseInterface memory", memoryNumbers)
	}
	log.Print(`注意思考: 为什么函数定义了接口,在调用函数的时候必须传递的是指针 &fileNumbers `)
}

// 直接返回所有数字
func createNumbers(n int) (numbers []int) {
	// 避免返回 []int{nil}
	numbers = []int{}
	for i:=0;i<n;i++ {
		numbers = append(numbers, i)
	}
	return
}
// 通过函数回调一段一段返回数字
func createNumbersUseCallbackFunc(n int, reader func(data []int)) {
	// 每一百个数字为一段
	chunkSize :=  10
	var chunks []int
	for i:=0;i<n;i++ {
		size := len(chunks)
		if size == chunkSize {
			reader(chunks) ; chunks = chunks[0:0] // [0:0] 是清空切片语法
		}
		chunks = append(chunks, i)
	}
	// 兜底操作,当 n = 123 时候第一次会进行 reader([]int{100,...,123})
	reader(chunks) ; chunks = chunks[0:0]
}
type MemoryNumbers struct {
	value []int
}
// 注意一定要是 *Memory (指针) ,否则修改值是没用的
func (m *MemoryNumbers) Reader(data []int){
	m.value = append(m.value, data...)
}
// 注意一定要是 *MemoryNumbers (指针) ,否则修改值是没用的
func createNumbersUseStruct(n int, m *MemoryNumbers) {
	createNumbersUseCallbackFunc(n, m.Reader)
}

//  此处不需要 m *NumbersReader 因为接口必须传入的是指针
func createNumbersUseInterface(n int, m NumbersReader) {
	createNumbersUseCallbackFunc(n, m.Reader)
}

type NumbersReader interface {
	Reader(m []int)
}
type FileNumbers struct {
	Name string
}
// 文件处理: https://github.com/nimoc/notebook/issues/5
// 注意一定要是 *Memory (指针) ,否则修改值是没用的
func (m *FileNumbers) Reader(data []int){
	byteList := []byte{}

	for i:=0;i< len(data);i++ {
		dataBytes := []byte(gconv.IntString(data[i]) + "\r\n")
		byteList = append(byteList, dataBytes...)
	}
	file, err := os.OpenFile(m.Name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) ; ge.Check(err)
	_, err = file.Write(byteList) ; ge.Check(err)
}
