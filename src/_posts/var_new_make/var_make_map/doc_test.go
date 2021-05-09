package doc_test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestVarMap(t *testing.T) {
	log.Print("声明变量，但没有分配空间")
	var data map[string]int
	assert.Equal(t, data, map[string]int(nil))
}
func TestDeclareEmptyMap(t *testing.T) {
	defer func() {
		log.Print(recover())
	}()
	log.Print("只声明变量，不分配空间，赋值时会panic")
	var data map[string]int
	data["age"] = 1
}
func TestMap(t *testing.T) {
	log.Print("声明变量，并初始化赋值")
	var data = map[string]int{}
	assert.Equal(t, data, map[string]int{})

	data["name"] = 1
	assert.Equal(t, data, map[string]int{"name":1})
}
func TestMakeMap(t *testing.T) {
	log.Print("通过 make 分配空间，也可以避免panic")
	var data = make(map[string]int)
	assert.Equal(t, data, map[string]int{})

	data["name"] = 1
	assert.Equal(t, data, map[string]int{"name":1})
}
