package doc_test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestVarSlice(t *testing.T) {
	log.Print("声明并初始化切片")
	var numbers = []int{1,2,3}
	assert.Equal(t, len(numbers), 3)
}
func TestMakeSliceLen0(t *testing.T) {
			log.Print("通过 make 声明并初始化切片")
			var numbers = make([]int, 0)
			// 等同于 numbers := []int{}
			assert.Equal(t, numbers, []int{})
			assert.Equal(t, len(numbers), 0)
			numbers = append(numbers, 1,2,3)
			assert.Equal(t, len(numbers), 3)
}

func TestMakeSliceLen2(t *testing.T) {
	log.Print("通过 make 声明并初始化切片长度为2， 初始化元素为类型的 zero value (0)")
	var numbers = make([]int, 2)
	// 等同于 numbers := []int{0, 0}
	assert.Equal(t, numbers, []int{0, 0})
	assert.Equal(t, len(numbers), 2)
	numbers = append(numbers, 1,2,3)
	assert.Equal(t, numbers, []int{0, 0, 1, 2, 3})
	assert.Equal(t, len(numbers), 5)
}

func TestArray(t *testing.T) {
	log.Print("创建数组（固定长度的切片）数组元素为 zero value")
	// 此处的 2 是 cap
	var numbers = [2]int{}
	assert.Equal(t, numbers[0], 0)
	assert.Equal(t, numbers[1], 0)
	// // 如果 index(10) 超出了 cap(2) 则会发生**编译期**错误
	// log.Print(numbers[10])
	assert.Equal(t, len(numbers), 2)
}

func TestArrayLen2Cap2(t *testing.T) {
	log.Print("通过 make 创建长度2 容量2的数字，初始化元素为类型的 zero value (0)")
	var numbers = make([]int, 2, 2)
	// 等同于 numbers := []int{0,0}
	assert.Equal(t, numbers, []int{0,0})
	assert.Equal(t, len(numbers), 2)
	assert.Equal(t, cap(numbers), 2)
	numbers[0] = 9
	assert.Equal(t, numbers, []int{9,0})
	// 如果 index(10) 超出了 cap(2) 则会发生**运行时**错误
	// numbers[10] = 9
}

func TestArrayLen0Cap2(t *testing.T) {
	log.Print("通过 make 创建长度0 容量2的数组")
	var numbers = make([]int, 0, 2)
	assert.Equal(t, numbers, []int{})
	assert.Equal(t, len(numbers), 0)
	assert.Equal(t, cap(numbers), 2)

	numbers = append(numbers, 1)
	assert.Equal(t, numbers, []int{1})
	assert.Equal(t, len(numbers), 1)
	assert.Equal(t, cap(numbers), 2)

	numbers = append(numbers, 2)
	assert.Equal(t, numbers, []int{1,2})
	assert.Equal(t, len(numbers), 2)
	assert.Equal(t, cap(numbers), 2)

	numbers = append(numbers, 3)
	assert.Equal(t, numbers, []int{1,2,3})
	assert.Equal(t, len(numbers), 3)
	assert.Equal(t, cap(numbers), 4)
}

func TestMakeArrayLen0Cap2Panic(t *testing.T) {
	log.Print("大部分场景下 numbers := make([]string, len, cap) len 和 cap 设置的不一致是没有意义的")
	func() {
		defer func() {
			log.Print(recover())
		}()
		numbers := make([]int, 0, 2)
		log.Print(`因为 numbers[0] = 1 会panic`)
		numbers[0] = 1
	}()
}